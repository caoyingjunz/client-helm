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

	utilhelm "github.com/caoyingjunz/client-helm/pkg/util/helm"
)

// A group's client should implement this interface.
type ReposGetter interface {
	Repos(namespace string) RepoInterface
}

type RepoInterface interface {
	Add(ctx context.Context) error    // add a chart repository
	Index(ctx context.Context) error  // generate an index file given a directory containing packaged charts
	List(ctx context.Context) error   // list chart repositories
	Remove(ctx context.Context) error // remove one or more chart repositories
	Update(ctx context.Context) error // update information of available charts locally from chart repositories

	RepoExpansion
}

type repo struct {
	client utilhelm.Interface
	ns     string
}

func newRepos(cc *AppsV1Client, namespace string) *repo {
	c := cc.Client()
	return &repo{
		client: c.GetClient(),
		ns:     namespace,
	}
}

func (c *repo) Add(ctx context.Context) error {
	return nil
}

func (c *repo) Index(ctx context.Context) error {
	return nil
}

func (c *repo) List(ctx context.Context) error {
	return nil
}

func (c *repo) Remove(ctx context.Context) error {
	return nil
}

func (c *repo) Update(ctx context.Context) error {
	return nil
}
