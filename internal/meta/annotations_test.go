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

package meta

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var _ = Describe("SetAnnotation", func() {
	const key = "test-key"

	DescribeTable(
		"should work as expected",
		func(annotations map[string]string, key, value string) {
			obj := &unstructured.Unstructured{}

			obj.SetAnnotations(annotations)

			SetAnnotation(obj, key, value)

			Expect(
				obj.GetAnnotations(),
			).To(
				HaveKeyWithValue(key, value),
			)
		},
		Entry("nil annotations", nil, key, "test value"),
		Entry("empty annotations", make(map[string]string), key, "test value"),
		Entry("existing annotation", map[string]string{key: "some-other-value"}, key, "test value"),
	)
})
