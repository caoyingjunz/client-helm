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

import (
	"github.com/caoyingjunz/client-helm/rest"
)

type AppsV1Interface interface {
	HelmsGetter
	SearchGetter
}

type AppsV1Client struct {
	helmClient rest.Interface
}

func (c *AppsV1Client) Helms(namespace string) HelmInterface {
	return newHelms(c, namespace)
}

func (c *AppsV1Client) Search(namespace string) SearchInterface {
	return newSearch(c, namespace)
}

// HelmClient returns a HelmClient that is used to communicate
// with helm server by this client implementation.
func (c *AppsV1Client) HelmClient() rest.Interface {
	return c.helmClient
}

// NewForConfig creates a new Helm AppsV1Client for the given config.
func NewForConfig(client rest.Interface) (*AppsV1Client, error) {
	return &AppsV1Client{helmClient: client}, nil
}
