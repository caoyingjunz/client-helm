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

package clientcmd

import (
	"k8s.io/klog/v2"

	restclient "github.com/caoyingjunz/client-helm/rest"
)

// BuildConfigFromFlags is a helper function that builds configs from a kubeconfig filepath.
// These are passed in as command line flags for cluster components.
// Warnings should reflect this usage. If kubeconfigPath
// are passed in we fallback to ~/.kube/config.
func BuildConfigFromFlags(kubeconfigPath string) (*restclient.Config, error) {
	if kubeconfigPath == "" {
		klog.Warningf("--kubeconfig was not specified.  This might not work.")
	}

	return &restclient.Config{KubeConfig: kubeconfigPath}, nil
}
