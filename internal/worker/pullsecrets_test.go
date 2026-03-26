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

package worker

import (
	"context"

	"github.com/google/go-containerregistry/pkg/name"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReadKubernetesSecrets", func() {
	It("should work as expected", func() {
		kc, err := ReadKubernetesSecrets(context.TODO(), "testdata/pull-secrets", GinkgoLogr)
		Expect(err).NotTo(HaveOccurred())

		By("dockercfg")

		dockerCfgTag, err := name.NewTag("dockercfg.registry/repo/image")
		Expect(err).NotTo(HaveOccurred())

		dockerCfgAuthenticator, err := kc.Resolve(dockerCfgTag)
		Expect(err).NotTo(HaveOccurred())

		dockerCfgAuthconfig, err := dockerCfgAuthenticator.Authorization()
		Expect(err).NotTo(HaveOccurred())

		Expect(dockerCfgAuthconfig.Username).To(Equal("username"))
		Expect(dockerCfgAuthconfig.Password).To(Equal("dockercfg"))

		By("dockerconfigjson")

		dockerConfigJsonTag, err := name.NewTag("dockerconfigjson.registry/repo/image")
		Expect(err).NotTo(HaveOccurred())

		dockerConfigJsonAuthenticator, err := kc.Resolve(dockerConfigJsonTag)
		Expect(err).NotTo(HaveOccurred())

		dockerConfigJsonAuthconfig, err := dockerConfigJsonAuthenticator.Authorization()
		Expect(err).NotTo(HaveOccurred())

		Expect(dockerConfigJsonAuthconfig.Username).To(Equal("username"))
		Expect(dockerConfigJsonAuthconfig.Password).To(Equal("dockerconfigjson"))

	})
})
