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

package config

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/utils/ptr"
)

var _ = Describe("ParseFile", func() {
	It("should return an error if the file does not exist", func() {
		_, err := ParseFile("/non/existent/path")
		Expect(err).To(HaveOccurred())
	})

	It("should parse the file correctly", func() {
		expected := &Config{
			HealthProbeBindAddress: ":8081",
			Job: Job{
				GCDelay: time.Hour,
			},
			LeaderElection: LeaderElection{
				Enabled:    true,
				ResourceID: "some-resource-id",
			},
			Metrics: Metrics{
				BindAddress:      "0.0.0.0:8443",
				EnableAuthnAuthz: true,
				SecureServing:    true,
			},
			WebhookPort: 9443,
			Worker: Worker{
				RunAsUser:        ptr.To[int64](1234),
				SELinuxType:      "mySELinuxType",
				FirmwareHostPath: ptr.To("/some/path"),
			},
		}

		Expect(
			ParseFile("testdata/config.yaml"),
		).To(
			Equal(expected),
		)
	})
})

var _ = Describe("ManagerOptions", func() {
	DescribeTable(
		"should enable authn/authz if configured",
		func(enabled bool) {
			c := &Config{
				Metrics: Metrics{EnableAuthnAuthz: enabled},
			}

			mo := c.ManagerOptions()

			if enabled {
				Expect(mo.Metrics.FilterProvider).NotTo(BeNil())
			} else {
				Expect(mo.Metrics.FilterProvider).To(BeNil())
			}
		},
		Entry(nil, false),
		Entry(nil, true),
	)
})
