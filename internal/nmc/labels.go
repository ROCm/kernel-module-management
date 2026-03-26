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
	"fmt"
	"regexp"
)

var (
	reConfiguredLabel = regexp.MustCompile(`^beta\.kmm\.node\.kubernetes\.io/([a-zA-Z0-9-]+)\.([a-zA-Z0-9-]+)\.module-configured$`)
	reInUseLabel      = regexp.MustCompile(`^beta\.kmm\.node\.kubernetes\.io/([a-zA-Z0-9-]+)\.([a-zA-Z0-9-]+)\.module-in-use$`)
)

func IsModuleConfiguredLabel(s string) (bool, string, string) {
	res := reConfiguredLabel.FindStringSubmatch(s)

	if len(res) != 3 {
		return false, "", ""
	}

	return true, res[1], res[2]
}

func IsModuleInUseLabel(s string) (bool, string, string) {
	res := reInUseLabel.FindStringSubmatch(s)

	if len(res) != 3 {
		return false, "", ""
	}

	return true, res[1], res[2]
}

func ModuleConfiguredLabel(namespace, name string) string {
	return fmt.Sprintf("beta.kmm.node.kubernetes.io/%s.%s.module-configured", namespace, name)
}

func ModuleInUseLabel(namespace, name string) string {
	return fmt.Sprintf("beta.kmm.node.kubernetes.io/%s.%s.module-in-use", namespace, name)
}
