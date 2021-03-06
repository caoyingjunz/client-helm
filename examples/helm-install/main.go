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
	"path"

	"k8s.io/client-go/util/homedir"

	metav1 "github.com/caoyingjunz/client-helm/api/meta/v1"
	"github.com/caoyingjunz/client-helm/helm"
	"github.com/caoyingjunz/client-helm/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags(path.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		panic(err)
	}
	helmClient, err := helm.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	if err = helmClient.AppsV1().Releases("kubez-sysns").Install(context.TODO(), "nginx", metav1.InstallOptions{
		ChartReference:  "bitnami/nginx",
		CreateNamespace: true,
	}); err != nil {
		panic(err)
	}
}
