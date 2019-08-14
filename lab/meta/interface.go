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

package core

import "github.com/zoumo/kubelab/lab/apps/v1"

// Interface provides access to each of this group's versions.
type Interface interface {
	V1() v1.Interface
}

type group struct {
}

// New returns a new Interface.
func New() Interface {
	return &group{}
}

// V1 returns a new v1.Interface.
func (g *group) V1() v1.Interface {
	return v1.New()
}
