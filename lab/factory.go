/*
Copyright 2019 zoumo (jim.zoumo@gmail.com). All rights reserved

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

package lab

import "github.com/zoumo/kubelab/lab/apps"

// Interface provides useful utils for resources in all known API group versions
type Interface interface {
	Apps() apps.Interface
}

type kubelab struct{}

// New constructs a new instance of a kubelab
func New() Interface {
	return &kubelab{}
}

func (l *kubelab) Apps() apps.Interface {
	return apps.New()
}
