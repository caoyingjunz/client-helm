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
	"bytes"
	"sync"
	"time"

	"k8s.io/klog/v2"
	utilexec "k8s.io/utils/exec"
	utiltrace "k8s.io/utils/trace"
)

type Interface interface {
	List(namespace Namespace, buffer *bytes.Buffer) error
}

const (
	cmdHelm string = "helm"
)

// Namespace represents different ns for helm (k8s)
type Namespace string

// runner implements Interface in terms of exec("helm").
type runner struct {
	mu   sync.Mutex
	exec utilexec.Interface
}

func New(exec utilexec.Interface) Interface {
	return &runner{
		exec: exec,
	}
}

func (runner *runner) List(namespace Namespace, buffer *bytes.Buffer) error {
	runner.mu.Lock()
	defer runner.mu.Unlock()

	trace := utiltrace.New("helm list")
	defer trace.LogIfLong(2 * time.Second)

	// run and return
	args := []string{"-n", string(namespace)}
	klog.V(4).Infof("running %s %v", cmdHelm, args)
	cmd := runner.exec.Command(cmdHelm, args...)
	cmd.SetStdout(buffer)
	stderrBuffer := bytes.NewBuffer(nil)
	cmd.SetStderr(stderrBuffer)

	err := cmd.Run()
	if err != nil {
		stderrBuffer.WriteTo(buffer)
	}
	return err
}
