/*
Copyright 2021 The Pixiu Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package helm

import (
	"context"
	"fmt"
	"sync"
	"time"

	"k8s.io/klog/v2"
	utilexec "k8s.io/utils/exec"
	utiltrace "k8s.io/utils/trace"

	metav1 "github.com/caoyingjunz/client-helm/api/meta/v1"
)

type Interface interface {
	Install(namespace string, name string, opts metav1.InstallOptions) error
	Delete(namespace string, name string, opts metav1.DeleteOptions) error
	Get(namespace string, name string) ([]byte, error)
	List(namespace string) ([]byte, error)
	HubList(namespace string, name string) ([]byte, error)
	RepoList(namespace string, name string) ([]byte, error)
}

const (
	cmdHelm string = "helm"
)

type operation string

const (
	opInstall    operation = "install"
	opList       operation = "list"
	opDelete     operation = "delete"
	opCreate     operation = "create"
	opSearch     operation = "search"
	opSearchHub  operation = "hub"
	opSearchRepo operation = "repo"
)

// Namespace represents different ns for helm (k8s)
type Namespace string

// runner implements Interface in terms of exec("helm").
type runner struct {
	mu         sync.Mutex
	exec       utilexec.Interface
	kubeConfig string
}

func New(exec utilexec.Interface, kubeconfig string) Interface {
	return &runner{
		exec:       exec,
		kubeConfig: kubeconfig,
	}
}

func (runner *runner) Install(namespace string, name string, opts metav1.InstallOptions) error {
	trace := utiltrace.New("helm install")
	defer trace.LogIfLong(2 * time.Second)

	if len(name) == 0 {
		return fmt.Errorf("name can not be empty when install release")
	}
	// TODO: only supported install chart relase by reference for now
	if opts.ChartReference == "" {
		return fmt.Errorf("chart reference can not be empty when install release")
	}

	// setup args
	args := []string{name, opts.ChartReference}
	if opts.CreateNamespace {
		args = append(args, "--create-namespace")
	}
	if opts.Version != nil {
		args = append(args, []string{"--version", *opts.Version}...)
	}
	if opts.Wait {
		args = append(args, "--wait")
	}
	if len(opts.ValuesFiles) != 0 {
		for _, valuesFile := range opts.ValuesFiles {
			// TODO: To ensure the yaml file exists
			args = append(args, []string{"-f", valuesFile}...)
		}
	}
	if len(opts.ValuesSets) != 0 {
		for k, v := range opts.ValuesSets {
			args = append(args, []string{"--set", fmt.Sprintf("%s=%s", k, v)}...)
		}
	}

	fullArgs := runner.makeFullArgs(namespace, args...)
	if out, err := runner.runContext(context.TODO(), opInstall, fullArgs); err != nil {
		return fmt.Errorf("error install release: %v: %s", err, out)
	}

	return nil
}

func (runner *runner) Delete(namespace string, name string, opts metav1.DeleteOptions) error {
	trace := utiltrace.New("helm delete")
	defer trace.LogIfLong(2 * time.Second)

	fullArgs := runner.makeFullArgs(namespace, name)
	klog.V(4).Infof("running %s %v", cmdHelm, fullArgs)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	out, err := runner.runContext(ctx, opDelete, fullArgs)
	if ctx.Err() == context.DeadlineExceeded {
		return fmt.Errorf("timed out while delete release %s", name)
	}
	if err != nil {
		return fmt.Errorf("error delete release: %v: %s", err, out)
	}

	return nil
}

func (runner *runner) Get(namespace string, name string) ([]byte, error) {
	trace := utiltrace.New("helm get")
	defer trace.LogIfLong(2 * time.Second)

	// setup args
	fullArgs := runner.makeFullArgs(namespace, []string{"-f", fmt.Sprintf("^%s$", name)}...)
	fullArgs = append(fullArgs, []string{"-o", "json"}...)
	klog.V(4).Infof("running %s %v", cmdHelm, fullArgs)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	out, err := runner.runContext(ctx, opList, fullArgs)
	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("timed out while get release")
	}
	if err != nil {
		return nil, fmt.Errorf("error get release: %v: %s", err, out)
	}

	return out, nil
}

func (runner *runner) List(namespace string) ([]byte, error) {
	//runner.mu.Lock()
	//defer runner.mu.Unlock()
	trace := utiltrace.New("helm list")
	defer trace.LogIfLong(2 * time.Second)

	fullArgs := runner.makeFullArgs(namespace)
	fullArgs = append(fullArgs, []string{"-o", "json"}...)

	klog.V(4).Infof("running %s %v", cmdHelm, fullArgs)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	out, err := runner.runContext(ctx, opList, fullArgs)
	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("timed out while list release")
	}
	if err == nil {
		return out, nil
	}

	return nil, fmt.Errorf("error list release: %v: %s", err, out)
}

func (runner *runner) HubList(namespace string, name string) ([]byte, error) {
	runner.mu.Lock()
	defer runner.mu.Unlock()

	trace := utiltrace.New("helm search")
	defer trace.LogIfLong(2 * time.Second)

	fullArgs := []string{string(opSearchHub)}
	fullArgs = append(fullArgs, runner.makeFullArgs(namespace, name)...)
	fullArgs = append(fullArgs, []string{"-o", "json"}...)

	klog.V(4).Infof("running %s %v", cmdHelm, fullArgs)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	out, err := runner.runContext(ctx, opSearch, fullArgs)
	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("timed out while list release")
	}
	if err == nil {
		return out, nil
	}

	return nil, fmt.Errorf("error search hub release: %v: %s", err, out)
}

func (runner *runner) RepoList(namespace string, name string) ([]byte, error) {
	runner.mu.Lock()
	defer runner.mu.Unlock()

	trace := utiltrace.New("helm search")
	defer trace.LogIfLong(2 * time.Second)

	fullArgs := []string{string(opSearchRepo)}
	fullArgs = append(fullArgs, runner.makeFullArgs(namespace, name)...)
	fullArgs = append(fullArgs, []string{"-o", "json"}...)

	klog.V(4).Infof("running %s %v", cmdHelm, fullArgs)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	out, err := runner.runContext(ctx, opSearch, fullArgs)
	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("timed out while list release")
	}
	if err == nil {
		return out, nil
	}

	return nil, fmt.Errorf("error search repo release: %v: %s", err, out)
}

func (runner *runner) makeFullArgs(namespace string, args ...string) []string {
	if len(runner.kubeConfig) != 0 {
		args = append(args, []string{"--kubeconfig", runner.kubeConfig}...)
	}
	return append(args, []string{"-n", namespace}...)
}

func (runner *runner) run(op operation, args []string) ([]byte, error) {
	return runner.runContext(context.TODO(), op, args)
}

func (runner *runner) runContext(ctx context.Context, op operation, args []string) ([]byte, error) {
	fullArgs := []string{string(op)}
	fullArgs = append(fullArgs, args...)

	klog.V(5).Infof("running helm: %s %v", cmdHelm, fullArgs)
	if ctx == nil {
		return runner.exec.Command(cmdHelm, fullArgs...).CombinedOutput()
	}

	return runner.exec.CommandContext(ctx, cmdHelm, fullArgs...).CombinedOutput()
}
