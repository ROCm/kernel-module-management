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
	"github.com/kubernetes-sigs/kernel-module-management/api/v1beta2"
	"github.com/kubernetes-sigs/kernel-module-management/internal/preflight"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const namespace = "namespace"

var _ = Describe("DeleteModuleStatus", func() {
	It("should do nothing if the corresponding status is not found", func() {
		statuses := []v1beta2.PreflightValidationModuleStatus{
			{Name: "module-1", Namespace: namespace},
			{Name: "module-2", Namespace: namespace},
		}

		DeleteModuleStatus(&statuses, types.NamespacedName{Name: "module-3", Namespace: namespace})

		Expect(statuses).To(HaveLen(2))
	})

	It("should delete an element if it is found", func() {
		const moduleName = "module-1"

		statuses := []v1beta2.PreflightValidationModuleStatus{
			{Name: moduleName, Namespace: namespace},
			{Name: "module-2", Namespace: namespace},
		}

		DeleteModuleStatus(&statuses, types.NamespacedName{Name: moduleName, Namespace: namespace})

		Expect(statuses).To(HaveLen(1))
	})
})

var _ = Describe("UpsertModuleStatus", func() {
	It("should return an error if name or namespace are not set", func() {
		Expect(
			UpsertModuleStatus(nil, v1beta2.PreflightValidationModuleStatus{Name: "name"}),
		).To(
			HaveOccurred(),
		)

		Expect(
			UpsertModuleStatus(nil, v1beta2.PreflightValidationModuleStatus{Namespace: namespace}),
		).To(
			HaveOccurred(),
		)
	})

	It("should update a status if it already exists", func() {
		const moduleName = "module-name"

		oldStatus := v1beta2.PreflightValidationModuleStatus{
			Name:      moduleName,
			Namespace: namespace,
			CRBaseStatus: v1beta2.CRBaseStatus{
				VerificationStatus: "True",
				StatusReason:       "some-reason",
				VerificationStage:  "some-stage",
				LastTransitionTime: metav1.Now(),
			},
		}

		newStatus := v1beta2.PreflightValidationModuleStatus{
			Name:      moduleName,
			Namespace: namespace,
			CRBaseStatus: v1beta2.CRBaseStatus{
				VerificationStatus: "False",
				StatusReason:       "some-other-reason",
				VerificationStage:  "some-other-stage",
				LastTransitionTime: metav1.Now(),
			},
		}

		Expect(oldStatus).NotTo(Equal(newStatus))

		statuses := []v1beta2.PreflightValidationModuleStatus{
			oldStatus,
			{
				Name:      "other-module",
				Namespace: namespace,
				CRBaseStatus: v1beta2.CRBaseStatus{
					VerificationStatus: "True",
					StatusReason:       "some-third-reason",
					VerificationStage:  "some-third-stage",
					LastTransitionTime: metav1.Now(),
				},
			},
		}

		Expect(
			UpsertModuleStatus(&statuses, newStatus),
		).NotTo(
			HaveOccurred(),
		)

		Expect(statuses).To(HaveLen(2))

		res, ok := preflight.FindModuleStatus(statuses, types.NamespacedName{Name: moduleName, Namespace: namespace})
		Expect(ok).To(BeTrue())
		Expect(*res).To(BeComparableTo(newStatus))
	})
})
