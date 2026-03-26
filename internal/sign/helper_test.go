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

package sign

import (
	"strings"

	"github.com/google/go-cmp/cmp"
	kmmv1beta1 "github.com/kubernetes-sigs/kernel-module-management/api/v1beta1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
)

var _ = Describe("GetRelevantSign", func() {

	const (
		unsignedImage = "my.registry/my/image"
		keySecret     = "securebootkey"
		certSecret    = "securebootcert"
		filesToSign   = "/modules/simple-kmod.ko:/modules/simple-procfs-kmod.ko"
		kernelVersion = "1.2.3"
	)

	var (
		h Helper
	)

	BeforeEach(func() {
		h = NewSignerHelper()
	})

	expected := &kmmv1beta1.Sign{
		UnsignedImage: unsignedImage,
		KeySecret:     &v1.LocalObjectReference{Name: keySecret},
		CertSecret:    &v1.LocalObjectReference{Name: certSecret},
		FilesToSign:   strings.Split(filesToSign, ":"),
	}

	DescribeTable("should set fields correctly", func(moduleSign *kmmv1beta1.Sign, mappingSign *kmmv1beta1.Sign) {
		actual, err := h.GetRelevantSign(moduleSign, mappingSign, kernelVersion)
		Expect(err).NotTo(HaveOccurred())
		Expect(
			cmp.Diff(expected, actual),
		).To(
			BeEmpty(),
		)
	},
		Entry(
			"no km.Sign",
			&kmmv1beta1.Sign{
				UnsignedImage: unsignedImage,
				KeySecret:     &v1.LocalObjectReference{Name: keySecret},
				CertSecret:    &v1.LocalObjectReference{Name: certSecret},
				FilesToSign:   strings.Split(filesToSign, ":"),
			},
			nil,
		),
		Entry(
			"no container.Sign",
			nil,
			&kmmv1beta1.Sign{
				UnsignedImage: unsignedImage,
				KeySecret:     &v1.LocalObjectReference{Name: keySecret},
				CertSecret:    &v1.LocalObjectReference{Name: certSecret},
				FilesToSign:   strings.Split(filesToSign, ":"),
			},
		),
		Entry(
			"default UnsignedImage",
			&kmmv1beta1.Sign{
				UnsignedImage: unsignedImage,
			},
			&kmmv1beta1.Sign{
				KeySecret:   &v1.LocalObjectReference{Name: keySecret},
				CertSecret:  &v1.LocalObjectReference{Name: certSecret},
				FilesToSign: strings.Split(filesToSign, ":"),
			},
		),
		Entry(
			"default UnsignedImage and KeySecret",
			&kmmv1beta1.Sign{
				UnsignedImage: unsignedImage,
				KeySecret:     &v1.LocalObjectReference{Name: keySecret},
			},
			&kmmv1beta1.Sign{
				CertSecret:  &v1.LocalObjectReference{Name: certSecret},
				FilesToSign: strings.Split(filesToSign, ":"),
			},
		),
		Entry(
			"default UnsignedImage, KeySecret, and CertSecret",
			&kmmv1beta1.Sign{
				UnsignedImage: unsignedImage,
				KeySecret:     &v1.LocalObjectReference{Name: keySecret},
				CertSecret:    &v1.LocalObjectReference{Name: certSecret},
			},
			&kmmv1beta1.Sign{
				FilesToSign: strings.Split(filesToSign, ":"),
			},
		),
		Entry(
			"default FilesToSign only",
			&kmmv1beta1.Sign{
				FilesToSign: strings.Split(filesToSign, ":"),
			},
			&kmmv1beta1.Sign{
				UnsignedImage: unsignedImage,
				KeySecret:     &v1.LocalObjectReference{Name: keySecret},
				CertSecret:    &v1.LocalObjectReference{Name: certSecret},
			},
		),
	)

})
var _ = Describe("GetRelevantSign", func() {

	const (
		unsignedImage = "my.registry/my/image"
		keySecret     = "securebootkey"
		certSecret    = "securebootcert"
		filesToSign   = "/modules/${KERNEL_VERSION}/simple-kmod.ko:/modules/${KERNEL_VERSION}/simple-procfs-kmod.ko"
		kernelVersion = "1.2.3"
	)

	var (
		h Helper
	)

	BeforeEach(func() {
		h = NewSignerHelper()
	})

	expected := &kmmv1beta1.Sign{
		UnsignedImage: unsignedImage + ":" + kernelVersion,
		KeySecret:     &v1.LocalObjectReference{Name: keySecret},
		CertSecret:    &v1.LocalObjectReference{Name: certSecret},
		FilesToSign:   strings.Split("/modules/"+kernelVersion+"/simple-kmod.ko:/modules/"+kernelVersion+"/simple-procfs-kmod.ko", ":"),
	}

	DescribeTable("should set fields correctly", func(moduleSign *kmmv1beta1.Sign, mappingSign *kmmv1beta1.Sign) {
		actual, _ := h.GetRelevantSign(moduleSign, mappingSign, kernelVersion)
		Expect(
			cmp.Diff(expected, actual),
		).To(
			BeEmpty(),
		)
	},
		Entry(
			"no km.Sign",
			&kmmv1beta1.Sign{
				UnsignedImage: unsignedImage + ":${KERNEL_VERSION}",
				KeySecret:     &v1.LocalObjectReference{Name: keySecret},
				CertSecret:    &v1.LocalObjectReference{Name: certSecret},
				FilesToSign:   strings.Split(filesToSign, ":"),
			},
			nil,
		),
		Entry(
			"no container.Sign",
			nil,
			&kmmv1beta1.Sign{
				UnsignedImage: unsignedImage + ":${KERNEL_VERSION}",
				KeySecret:     &v1.LocalObjectReference{Name: keySecret},
				CertSecret:    &v1.LocalObjectReference{Name: certSecret},
				FilesToSign:   strings.Split(filesToSign, ":"),
			},
		),
	)
})
