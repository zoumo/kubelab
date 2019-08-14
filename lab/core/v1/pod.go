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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	k8spodutil "k8s.io/kubernetes/pkg/api/pod"
	k8score "k8s.io/kubernetes/pkg/apis/core"
)

var (
	DefaultPodLab = &podImpl{}
)

type PodLab interface {
	DropDisabledAlphaFields(in *corev1.PodSpec)
}

type podImpl struct{}

// DropDisabledAlphaFields removes disabled fields from the pod spec.
// This should be called from PrepareForCreate/PrepareForUpdate for all resources containing a pod spec.
//
// TODO: the feature in utilfeature.DefaultFeatureGate must be the same as apiserver
func (l *podImpl) DropDisabledAlphaFields(in *corev1.PodSpec) {
	out := &k8score.PodSpec{}
	legacyscheme.Scheme.Convert(in, out, nil)
	// drop disabled alpha fields in podSpec
	k8spodutil.DropDisabledAlphaFields(out)
	legacyscheme.Scheme.Convert(out, in, nil)
}
