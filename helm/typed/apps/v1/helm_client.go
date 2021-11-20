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

package v1

import "github.com/caoyingjunz/client-helm/rest"

type AppsV1Interface interface {
	HelmsGetter
}

type AppsV1Client struct {
	restConfig rest.Config
}

func (c *AppsV1Client) Helms(namespace string) HelmInterface {
	return newHelms(c.restConfig, namespace)
}

// NewForConfig creates a new Helm AppsV1Client for the given config.
func NewForConfig(config *rest.Config) (*AppsV1Client, error) {
	return &AppsV1Client{*config}, nil
}

// NewForConfigOrDie creates a new AppsV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(config *rest.Config) *AppsV1Client {
	client, err := NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return client
}
