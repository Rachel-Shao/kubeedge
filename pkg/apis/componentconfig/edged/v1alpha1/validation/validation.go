/*
Copyright 2019 The KubeEdge Authors.

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

package validation

import (
	"github.com/kubeedge/kubeedge/pkg/apis/componentconfig/edged/v1alpha1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/klog/v2"
)

// ValidateEdgedConfiguration validates `c` and returns an errorList if it is invalid
func ValidateEdgedConfiguration(c *v1alpha1.EdgedConfig) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, ValidateModuleEdged(*c.Modules.Edged)...)
	return allErrs
}


// ValidateModuleEdged validates `e` and returns an errorList if it is invalid
func ValidateModuleEdged(e v1alpha1.Edged) field.ErrorList {
	if !e.Enable {
		return field.ErrorList{}
	}
	allErrs := field.ErrorList{}
	if e.NodeIP == "" {
		klog.Warningf("NodeIP is empty , use default ip which can connect to cloud.")
	}
	switch e.CGroupDriver {
	case v1alpha1.CGroupDriverCGroupFS, v1alpha1.CGroupDriverSystemd:
	default:
		allErrs = append(allErrs, field.Invalid(field.NewPath("CGroupDriver"), e.CGroupDriver,
			"CGroupDriver value error"))
	}
	return allErrs
}

