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
	"context"

	"github.com/caoyingjunz/client-helm/api/apps/v1"
	metav1 "github.com/caoyingjunz/client-helm/api/meta/v1"
)

// A group's client should implement this interface.
type HelmGetter interface {
	Helm(namespace string) HelmInterface
}

// HelmInterface has methods to work with Helm resources.
type HelmInterface interface {
	Create(ctx context.Context, opts metav1.CreateOptions) error
	List(ctx context.Context) (*v1.HelmList, error)
	Get(ctx context.Context) (*v1.Helm, error)

	HelmExpansion
}

// helm implements HelmInterface
type helm struct {
	ns string
}

func newHelm(namespace string) *helm {
	return &helm{
		ns: namespace,
	}
}

func (c *helm) Create(ctx context.Context, opts metav1.CreateOptions) error {
	// TODO
	return nil
}

// List returns the list of Helms that match those ns
func (c *helm) List(ctx context.Context) (*v1.HelmList, error) {
	// TODO
	return nil, nil
}

func (c *helm) Get(ctx context.Context) (*v1.Helm, error) {
	// TODO
	return nil, nil
}