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

package nmc

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ModuleConfiguredLabel", func() {
	It("should work as expected", func() {

		Expect(
			ModuleConfiguredLabel("a", "b"),
		).To(
			Equal("beta.kmm.node.kubernetes.io/a.b.module-configured"),
		)
	})
})

var _ = Describe("ModuleInUseLabel", func() {
	It("should work as expected", func() {

		Expect(
			ModuleInUseLabel("a", "b"),
		).To(
			Equal("beta.kmm.node.kubernetes.io/a.b.module-in-use"),
		)
	})
})

var _ = Describe("IsModuleConfiguredLabel", func() {
	DescribeTable(
		"should work as expected",
		func(input string, expectedOK bool, expectedNS, expectedName string) {
			ok, ns, name := IsModuleConfiguredLabel(input)

			if !expectedOK {
				Expect(ok).To(BeFalse())
				return
			}

			Expect(ok).To(BeTrue())
			Expect(ns).To(Equal(expectedNS))
			Expect(name).To(Equal(expectedName))
		},
		Entry(nil, "a.b.module-in-use", false, "", ""),
		Entry(nil, "beta.kmm.node.kubernetes.io/a.b.module-configured", true, "a", "b"),
		Entry(nil, "beta.kmm.node.kubernetes.io/..module-configured", false, "", ""),
		Entry(nil, "beta.kmm.node.kubernetes.io/a123.b456.module-configured", true, "a123", "b456"),
		Entry(nil, "beta.kmm.node.kubernetes.io/with-hypen.withouthypen.module-configured", true, "with-hypen", "withouthypen"),
	)
})

var _ = Describe("IsModuleInUseLabel", func() {
	DescribeTable(
		"should work as expected",
		func(input string, expectedOK bool, expectedNS, expectedName string) {
			ok, ns, name := IsModuleInUseLabel(input)

			if !expectedOK {
				Expect(ok).To(BeFalse())
				return
			}

			Expect(ok).To(BeTrue())
			Expect(ns).To(Equal(expectedNS))
			Expect(name).To(Equal(expectedName))
		},
		Entry(nil, "a.b.module-in-use", false, "", ""),
		Entry(nil, "beta.kmm.node.kubernetes.io/a.b.module-in-use", true, "a", "b"),
		Entry(nil, "beta.kmm.node.kubernetes.io/..module-in-use", false, "", ""),
		Entry(nil, "beta.kmm.node.kubernetes.io/a123.b456.module-in-use", true, "a123", "b456"),
		Entry(nil, "beta.kmm.node.kubernetes.io/with-hypen.withouthypen.module-in-use", true, "with-hypen", "withouthypen"),
	)
})
