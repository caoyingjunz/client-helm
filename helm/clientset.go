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
	v1 "github.com/caoyingjunz/client-helm/helm/typed/apps/v1"
	"github.com/caoyingjunz/client-helm/rest"
)

type Interface interface {
	Helms() v1.HelmInterface
}

// Clientset contains the clients for groups. Each group maybe has exactly one
// version included in a Clientset.
type Clientset struct {
	restConfig *rest.Config

	helm v1.AppsV1Interface
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	var cs Clientset
	var err error
	cs.helm, err = v1.NewForConfig(c)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}
