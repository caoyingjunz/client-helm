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
)

type Interface interface {
	Delete(namespace string, name string) error
	Get(namespace string, name string) ([]byte, error)
	List(namespace string) ([]byte, error)
}

const (
	cmdHelm string = "helm"
)

type operation string

const (
	opList   operation = "list"
	opDelete operation = "delete"
	opCreate operation = "create"
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

func (runner *runner) Delete(namespace string, name string) error {
	runner.mu.Lock()
	defer runner.mu.Unlock()

	trace := utiltrace.New("helm delete")
	defer trace.LogIfLong(2 * time.Second)

	fullArgs := makeFullArgs(namespace, name)
	klog.V(4).Infof("running %s %v", cmdHelm, fullArgs)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	out, err := runner.runContext(ctx, opDelete, fullArgs)
	if ctx.Err() == context.DeadlineExceeded {
		return fmt.Errorf("timed out while delete release %s", name)
	}
	if err == nil {
		return nil
	}

	return fmt.Errorf("error delete release: %v: %s", err, out)
}

func (runner *runner) Get(namespace string, name string) ([]byte, error) {
	runner.mu.Lock()
	defer runner.mu.Unlock()

	trace := utiltrace.New("helm get")
	defer trace.LogIfLong(2 * time.Second)

	fullArgs := makeFullArgs(namespace, "-f", fmt.Sprintf("^%s$", name))
	fullArgs = append(fullArgs, []string{"-o", "json"}...)

	klog.V(4).Infof("running %s %v", cmdHelm, fullArgs)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	out, err := runner.runContext(ctx, opList, fullArgs)
	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("timed out while get release")
	}
	if err == nil {
		return out, nil
	}

	return nil, fmt.Errorf("error get release: %v: %s", err, out)
}

func (runner *runner) List(namespace string) ([]byte, error) {
	runner.mu.Lock()
	defer runner.mu.Unlock()

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

func (runner *runner)makeFullArgs(namespace string, args ...string) []string {
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
