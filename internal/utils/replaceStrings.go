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
	"fmt"
	"regexp"
	"strings"

	"github.com/a8m/envsubst/parse"
)

const (
	kernelVersionMajorIdx = 0
	kernelVersionMinorIdx = 1
	kernelVersionPatchIdx = 2
)

var kernelRegexp = regexp.MustCompile("[.,-]")

func KernelComponentsAsEnvVars(kernel string) []string {
	osConfigFieldsList := kernelRegexp.Split(kernel, -1)

	envvars := []string{
		"KERNEL_FULL_VERSION=" + kernel,
		"KERNEL_VERSION=" + kernel,
		"KERNEL_XYZ=" + strings.Join(osConfigFieldsList[:kernelVersionPatchIdx+1], "."),
		"KERNEL_X=" + osConfigFieldsList[kernelVersionMajorIdx],
		"KERNEL_Y=" + osConfigFieldsList[kernelVersionMinorIdx],
		"KERNEL_Z=" + osConfigFieldsList[kernelVersionPatchIdx],
	}

	return envvars
}

func ReplaceInTemplates(envvars []string, templates ...string) ([]string, error) {
	parser := parse.New("mapping", envvars, &parse.Restrictions{})

	replacedStrings := make([]string, 0, len(templates))

	for _, v := range templates {
		resultString, err := parser.Parse(v)
		if err != nil {
			return nil, fmt.Errorf("failed to substitute %q: %v", v, err)
		}

		replacedStrings = append(replacedStrings, resultString)
	}
	return replacedStrings, nil
}
