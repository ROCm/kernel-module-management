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

package controllers

import "github.com/kubernetes-sigs/kernel-module-management/internal/utils"

const (
	labelMissing   = "missing"
	labelPresent   = "present"
	labelDifferent = "different"

	addAction    = "add"
	deleteAction = "delete"
	noneAction   = "none"
)

type labelActionKey struct {
	module       string
	workerPod    string
	devicePlugin string
}

type labelAction struct {
	getLabelName func(string, string) string
	action       string
}

var labelActionTable = map[labelActionKey]labelAction{
	labelActionKey{
		module:       labelMissing,
		workerPod:    labelMissing,
		devicePlugin: labelMissing}: labelAction{getLabelName: nil, action: noneAction},

	labelActionKey{
		module:       labelMissing,
		workerPod:    labelPresent,
		devicePlugin: labelPresent}: labelAction{getLabelName: utils.GetDevicePluginVersionLabelName, action: deleteAction},

	labelActionKey{
		module:       labelMissing,
		workerPod:    labelPresent,
		devicePlugin: labelMissing}: labelAction{getLabelName: utils.GetWorkerPodVersionLabelName, action: deleteAction},

	labelActionKey{
		module:       labelPresent,
		workerPod:    labelMissing,
		devicePlugin: labelMissing}: labelAction{getLabelName: utils.GetWorkerPodVersionLabelName, action: addAction},

	labelActionKey{
		module:       labelPresent,
		workerPod:    labelPresent,
		devicePlugin: labelMissing}: labelAction{getLabelName: utils.GetDevicePluginVersionLabelName, action: addAction},

	labelActionKey{
		module:       labelPresent,
		workerPod:    labelPresent,
		devicePlugin: labelPresent}: labelAction{getLabelName: nil, action: noneAction},

	labelActionKey{
		module:       labelPresent,
		workerPod:    labelDifferent,
		devicePlugin: labelDifferent}: labelAction{getLabelName: utils.GetDevicePluginVersionLabelName, action: deleteAction},

	labelActionKey{
		module:       labelPresent,
		workerPod:    labelDifferent,
		devicePlugin: labelMissing}: labelAction{getLabelName: utils.GetWorkerPodVersionLabelName, action: deleteAction},
}
