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

package labels

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetModuleReadyAndDevicePluginReadyLabels", func() {
	It("module ready label", func() {
		res := GetKernelModuleReadyNodeLabel("some-namespace", "some-module")
		Expect(res).To(Equal("kmm.node.kubernetes.io/some-namespace.some-module.ready"))
	})

	It("device-plugin ready label", func() {
		res := GetDevicePluginNodeLabel("some-namespace", "some-module")
		Expect(res).To(Equal("kmm.node.kubernetes.io/some-namespace.some-module.device-plugin-ready"))
	})
})
