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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

var (
	DefaultObjectMetaLab = &objectMetaImpl{}
)

// ObjectMetaLab contains some utils for ObjectMeta
type ObjectMetaLab interface {
	// Merge only merges Labels and Annotations
	Merge(dst, src *metav1.ObjectMeta)
	// IsEqual only checks Labels and Annotations
	IsEqual(a, b *metav1.ObjectMeta) bool
}

type objectMetaImpl struct{}

func (l *objectMetaImpl) Merge(dst, src *metav1.ObjectMeta) {
	if len(src.Labels) > 0 && dst.Labels == nil {
		dst.Labels = make(map[string]string)
	}
	for k, v := range src.Labels {
		dst.Labels[k] = v
	}
	if len(src.Annotations) > 0 && dst.Labels == nil {
		dst.Annotations = make(map[string]string)
	}
	for k, v := range src.Annotations {
		dst.Annotations[k] = v
	}
	dst.OwnerReferences = src.OwnerReferences
}

func (l *objectMetaImpl) IsEqual(a, b *metav1.ObjectMeta) bool {
	if !reflect.DeepEqual(a.Labels, b.Labels) {
		klog.V(5).Infof("%v/%v labels changed, a.Labels: %v, b.Labels: %v", a.Namespace, b.Name, a.Labels, b.Labels)
		return false
	}
	if !reflect.DeepEqual(a.Annotations, b.Annotations) {
		klog.V(5).Infof("%v/%v Annotations changed, a.Annotations: %v, b.Annotations: %v", a.Namespace, b.Name, a.Annotations, b.Annotations)
		return false
	}
	return true
}
