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
	"os"

	v1 "k8s.io/api/core/v1"
)

var (
	proxyEnvList = []string{
		"HTTP_PROXY", "HTTPS_PROXY", "NO_PROXY",
		"http_proxy", "https_proxy", "no_proxy",
	}
)

func GetProxyEnvVars() []v1.EnvVar {
	envVarList := []v1.EnvVar{}
	for _, env := range proxyEnvList {
		if val, ok := os.LookupEnv(env); ok {
			envVarList = append(envVarList, v1.EnvVar{
				Name:  env,
				Value: val,
			})
		}
	}
	if len(envVarList) > 0 {
		return envVarList
	}
	return nil
}
