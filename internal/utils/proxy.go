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
