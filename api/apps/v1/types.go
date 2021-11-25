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

type Helm struct {
	Name       string `json:"name,omitempty"`
	Namespace  string `json:"namespace,omitempty"`
	Revision   string `json:"revision,omitempty"`
	Updated    string `json:"updated,omitempty"`
	Status     string `json:"status,omitempty"`
	Chart      string `json:"chart,omitempty"`
	AppVersion string `json:"app_version,omitempty"`
}

type HelmList struct {
	// Items is the list of Helms.
	Items []Helm `json:"items"`
}

type SearchHub struct {
	Url           string `json:"url,omitempty"`
	AppVersion    string `json:"app_version,omitempty"`
	Description   string `json:"description,omitempty"`
}
type SearchHubList struct {
	Items []SearchHub `json:"items"`
}

type SearchRepoList struct {
	Items []SearchRepo `json:"items"`
}

type SearchRepo struct {
	Name string `json:"name,omitempty"`
	AppVersion    string `json:"app_version,omitempty"`
	Description   string `json:"description,omitempty"`
}
