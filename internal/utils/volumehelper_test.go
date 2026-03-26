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

package utils

import (
	kmmv1beta1 "github.com/kubernetes-sigs/kernel-module-management/api/v1beta1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
)

var _ = Describe("MakeSecretVolumeMount", func() {
	It("should return a valid volumeMount", func() {
		signConfig := &kmmv1beta1.Sign{
			CertSecret: &v1.LocalObjectReference{Name: "securebootcert"},
		}
		secretMount := v1.VolumeMount{
			Name:      "secret-securebootcert",
			ReadOnly:  true,
			MountPath: "/signingcert",
		}

		volMount := MakeSecretVolumeMount(signConfig.CertSecret, "/signingcert", true)
		Expect(volMount).To(Equal(secretMount))
	})

	It("should return an empty volumeMount if signConfig is empty", func() {
		Expect(
			MakeSecretVolumeMount(nil, "/signingcert", true),
		).To(
			Equal(v1.VolumeMount{}),
		)
	})
})
