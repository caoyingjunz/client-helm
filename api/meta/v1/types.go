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

// CreateOptions may be provided when creating an API object.
type CreateOptions struct{}

type InstallOptions struct {
	ChartReference string `json:"chartReference,omitempty"`
	// Create the release namespace if not present
	// +optional
	CreateNamespace bool `json:"createNamespace,omitempty"`

	// generate the name (and omit the NAME parameter)
	// +optional
	GenerateName bool `json:"generateName,omitempty"`

	// Specify the exact chart version to use. If this is not specified, the latest version is used
	// +optional
	Version *string `json:"version,omitempty"`
	// if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment,
	// StatefulSet, or ReplicaSet are in a ready state before marking the release as successful.
	Wait bool `json:"wait"`

	// Specify values in a YAML file or a URL (can specify multiple)
	// +optional
	ValuesFiles []string `json:"valuesFiles,omitempty"`

	// Set values on the command line
	// +optional
	ValuesSets map[string]string `json:"valueSets,omitempty"`
}

type DeleteOptions struct{}

// GetOptions is the standard query options to the standard REST get call.
type GetOptions struct {
	ResourceVersion string `json:"resourceVersion,omitempty"`
}

type ListOptions struct{}

type RepoListOptions struct{}

type HubListOptions struct{}
