package v1

import (
	"context"
	"encoding/json"
	"fmt"

	v1 "github.com/caoyingjunz/client-helm/api/apps/v1"
	metav1 "github.com/caoyingjunz/client-helm/api/meta/v1"
	utilhelm "github.com/caoyingjunz/client-helm/pkg/util/helm"
)

type SearchGetter interface {
	Search(namespace string) SearchInterface
}

type SearchInterface interface {
	HubList(ctx context.Context, name string, opts metav1.HubListOptions) (*v1.SearchHubList, error)
	RepoList(ctx context.Context, name string, opts metav1.RepoListOptions) (*v1.SearchRepoList, error)
}

type search struct {
	client utilhelm.Interface
	ns     string
}

func newSearch(cc *AppsV1Client, namespace string) *search {
	if len(namespace) == 0 {
		namespace = defaultNamespace
	}

	client := cc.Client()
	return &search{
		client: client.GetClient(),
		ns:     namespace,
	}
}

func (c *search) HubList(ctx context.Context, name string, opts metav1.HubListOptions) (*v1.SearchHubList, error) {
	out, err := c.client.HubList(c.ns, name)
	if err != nil {
		return nil, err
	}

	var hs []v1.SearchHub
	if err = json.Unmarshal(out, &hs); err != nil {
		return nil, fmt.Errorf("unmarshal to helms failed %v", err)
	}

	return &v1.SearchHubList{
		Items: hs,
	}, nil
}

func (c *search) RepoList(ctx context.Context, name string, opts metav1.RepoListOptions) (*v1.SearchRepoList, error) {
	out, err := c.client.RepoList(c.ns, name)
	if err != nil {
		return nil, err
	}

	var hs []v1.SearchRepo
	if err = json.Unmarshal(out, &hs); err != nil {
		return nil, fmt.Errorf("unmarshal to helms failed %v", err)
	}

	return &v1.SearchRepoList{
		Items: hs,
	}, nil
}
