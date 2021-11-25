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

package main

import (
	"context"
	"fmt"
	"k8s.io/client-go/util/homedir"
	"path"

	metav1 "github.com/caoyingjunz/client-helm/api/meta/v1"
	"github.com/caoyingjunz/client-helm/helm"
	"github.com/caoyingjunz/client-helm/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags(path.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		panic(err)
	}
	clientSet, err := helm.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	releases, err := clientSet.AppsV1().Search("default").RepoList(context.TODO(), "redis", metav1.RepoListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(fmt.Sprintf("%+v", releases.Items))

	release, err := clientSet.AppsV1().Search("default").HubList(context.TODO(), "redis", metav1.HubListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(fmt.Sprintf("%+v", release.Items))
}