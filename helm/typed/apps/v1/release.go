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
type ReleasesGetter interface {
	Releases(namespace string) ReleaseInterface
}

// ReleaseInterface has methods to work with release resources.
type ReleaseInterface interface {
	Create(ctx context.Context, opts metav1.CreateOptions) error
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Release, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ReleaseList, error)

	ReleaseExpansion
}

// release implements ReleaseInterface
type release struct {
	client utilhelm.Interface
	ns     string
}

// newReleases returns s release
func newReleases(cc *AppsV1Client, namespace string) *release {
	if len(namespace) == 0 {
		namespace = defaultNamespace
	}

	c := cc.Client()
	return &release{
		client: c.GetClient(),
		ns:     namespace,
	}
}

func (c *release) Create(ctx context.Context, opts metav1.CreateOptions) error {
	// TODO
	return nil
}

func (c *release) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete(c.ns, name)
}

func (c *release) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Release, error) {
	out, err := c.client.Get(c.ns, name)
	if err != nil {
		return nil, err
	}

	var hs []v1.Release
	if err = json.Unmarshal(out, &hs); err != nil {
		return nil, fmt.Errorf("unmarshal to release failed %v", err)
	}
	if len(hs) == 0 {
		return nil, utilhelm.ErrReleaseNotFound
	}

	return &hs[0], nil
}

// List returns the list of Helms that match those ns
func (c *release) List(ctx context.Context, opts metav1.ListOptions) (*v1.ReleaseList, error) {
	out, err := c.client.List(c.ns)
	if err != nil {
		return nil, err
	}

	var hs []v1.Release
	if err = json.Unmarshal(out, &hs); err != nil {
		return nil, fmt.Errorf("unmarshal to release failed %v", err)
	}

	return &v1.ReleaseList{
		Items: hs,
	}, nil
}
