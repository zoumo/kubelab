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
	"encoding/json"
	"fmt"
	"reflect"

	libcorev1 "github.com/zoumo/kubelab/lab/core/v1"
	libmetav1 "github.com/zoumo/kubelab/lab/meta/v1"

	"github.com/InVisionApp/conjungo"
	"github.com/mattbaird/jsonpatch"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/klog"
	k8sappsv1 "k8s.io/kubernetes/pkg/apis/apps/v1"
)

// DeploymentLab contains some utils for Deployments
type DeploymentLab interface {
	// Merge merges the following fields from src to dst
	// - ObjectMeta.Labels
	// - ObjectMeta.Annotations
	// - Spec
	Merge(dst, src *appsv1.Deployment) error
	// IsEqual checks if the given two deployments are equal
	//
	// If ignoreFields is provided, the function will be call on each
	// deployment's deepcopy(be free to mutate it) before comparing to
	// ignore specified fields.
	// You can mutate object in the function like:
	// func (in *appsv1.Deployment) {
	//    in.Spec.Replicas = nil
	// }
	IsEqual(a, b *appsv1.Deployment, ignoreFields func(*appsv1.Deployment)) bool
}

func newOptions() *conjungo.Options {
	o := conjungo.NewOptions()
	o.SetKindMergeFunc(reflect.Slice, func(t, s reflect.Value, o *conjungo.Options) (reflect.Value, error) {
		// Merges two slices of the same type by appending source to target.
		if t.Type() != s.Type() {
			return reflect.Value{}, fmt.Errorf("slices must have same type: T: %v S: %v", t.Type(), s.Type())
		}
		if !t.CanSet() {
			return reflect.Value{}, fmt.Errorf("the target value can not be set")
		}
		if o.Overwrite {
			// overwrite it no matter what it is
			t.Set(s)
		} else if t.Len() == 0 && s.Len() > 0 {
			// without overwrite
			// only change target when it is empty but src is not
			t.Set(s)
		}
		return t, nil
	})

	return o
}

type deploymentImpl struct{}

func (l *deploymentImpl) Merge(dst, src *appsv1.Deployment) error {
	setObjectDefaultsDeployments(src)
	libmetav1.DefaultObjectMetaLab.Merge(&dst.ObjectMeta, &src.ObjectMeta)
	// merge spec
	err := conjungo.Merge(&dst.Spec, src.Spec, newOptions())
	if err != nil {
		return err
	}
	return nil
}

func (l *deploymentImpl) IsEqual(a, b *appsv1.Deployment, ignoreFields func(*appsv1.Deployment)) bool {
	acopy := a.DeepCopy()
	bcopy := b.DeepCopy()

	setObjectDefaultsDeployments(acopy)
	setObjectDefaultsDeployments(bcopy)

	if ignoreFields != nil {
		ignoreFields(acopy)
		ignoreFields(bcopy)
	}

	if !libmetav1.DefaultObjectMetaLab.IsEqual(&acopy.ObjectMeta, &bcopy.ObjectMeta) {
		klog.V(2).Infof("deployment %v/%v metadata changed", a.Namespace, a.Name)
		return false
	}

	aSpecBytes, _ := json.Marshal(acopy.Spec)
	bSpecBytes, _ := json.Marshal(bcopy.Spec)

	if !reflect.DeepEqual(aSpecBytes, bSpecBytes) {
		if klog.V(2) {
			if patch, err := jsonpatch.CreatePatch(aSpecBytes, bSpecBytes); err == nil {
				klog.Infof("deployment %v/%v spec changed, the patch is: %v", a.Namespace, a.Name, patch)
			}
		}
		klog.V(5).Infof("deployment %v/%v spec changed\na => %v\nb => %v", a.Namespace, a.Name, string(aSpecBytes), string(bSpecBytes))
		return false
	}
	return true
}

func setObjectDefaultsDeployments(in *appsv1.Deployment) {
	k8sappsv1.SetObjectDefaults_Deployment(in)
	libcorev1.DefaultPodLab.DropDisabledAlphaFields(&in.Spec.Template.Spec)
}
