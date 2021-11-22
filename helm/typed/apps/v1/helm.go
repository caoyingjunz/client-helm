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
	"encoding/json"
	"fmt"

	"github.com/caoyingjunz/client-helm/api/apps/v1"
	metav1 "github.com/caoyingjunz/client-helm/api/meta/v1"
	utilhelm "github.com/caoyingjunz/client-helm/pkg/util/helm"
)

const (
	defaultNamespace = "default"
)

// A group's client should implement this interface.
type HelmsGetter interface {
	Helms(namespace string) HelmInterface
}

// HelmInterface has methods to work with Helm resources.
type HelmInterface interface {
	Create(ctx context.Context, opts metav1.CreateOptions) error
	List(ctx context.Context, opts metav1.ListOptions) (*v1.HelmList, error)
	Get(ctx context.Context, opts metav1.GetOptions) (*v1.Helm, error)

	HelmExpansion
}

// helm implements HelmInterface
type helm struct {
	kubeConfig string
	helmClient utilhelm.Interface
	ns         string
}

// newHelms returns s Helms
func newHelms(cc *AppsV1Client, namespace string) *helm {
	if len(namespace) == 0 {
		namespace = defaultNamespace
	}

	client := cc.HelmClient()
	return &helm{
		kubeConfig: client.GetConfig(),
		helmClient: client.GetClient(),
		ns:         namespace,
	}
}

func (c *helm) Create(ctx context.Context, opts metav1.CreateOptions) error {
	// TODO
	return nil
}

// List returns the list of Helms that match those ns
func (c *helm) List(ctx context.Context, opts metav1.ListOptions) (*v1.HelmList, error) {
	out, err := c.helmClient.List(c.ns)
	if err != nil {
		return nil, err
	}

	var hs []v1.Helm
	if err = json.Unmarshal(out, &hs); err != nil {
		return nil, fmt.Errorf("unmarshal to helms failed %v", err)
	}

	return &v1.HelmList{
		Items: hs,
	}, nil
}

func (c *helm) Get(ctx context.Context, opts metav1.GetOptions) (*v1.Helm, error) {
	// TODO
	return nil, nil
}
