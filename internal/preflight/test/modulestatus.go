/*
Copyright (c) Advanced Micro Devices, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the \"License\");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an \"AS IS\" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package test

import (
	"errors"
	"slices"

	"github.com/kubernetes-sigs/kernel-module-management/api/v1beta2"
	"github.com/kubernetes-sigs/kernel-module-management/internal/preflight"
	"k8s.io/apimachinery/pkg/types"
)

func DeleteModuleStatus(statuses *[]v1beta2.PreflightValidationModuleStatus, nsn types.NamespacedName) {
	*statuses = slices.DeleteFunc(*statuses, func(status v1beta2.PreflightValidationModuleStatus) bool {
		return status.Namespace == nsn.Namespace && status.Name == nsn.Name
	})
}

func UpsertModuleStatus(statuses *[]v1beta2.PreflightValidationModuleStatus, s v1beta2.PreflightValidationModuleStatus) error {
	if s.Name == "" || s.Namespace == "" {
		return errors.New("name and namespace may not be empty")
	}

	found, ok := preflight.FindModuleStatus(*statuses, types.NamespacedName{Name: s.Name, Namespace: s.Namespace})
	if ok {
		found.CRBaseStatus = s.CRBaseStatus
	} else {
		*statuses = append(*statuses, s)
	}

	return nil
}
