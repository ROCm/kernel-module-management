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
	"fmt"
	"os"

	kmmv1beta1 "github.com/kubernetes-sigs/kernel-module-management/api/v1beta1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

//go:generate mockgen -source=config.go -package=worker -destination=mock_config.go

type ConfigHelper interface {
	ReadConfigFile(path string) (*kmmv1beta1.ModuleConfig, error)
}

type configHelper struct{}

func NewConfigHelper() ConfigHelper {
	return &configHelper{}
}

func (c *configHelper) ReadConfigFile(path string) (*kmmv1beta1.ModuleConfig, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read the configuration file %s: %v", path, err)
	}

	cfg := kmmv1beta1.ModuleConfig{}

	if err = yaml.UnmarshalStrict(b, &cfg); err != nil {
		return nil, fmt.Errorf("could not decode the configuration from %s: %v", path, err)
	}

	return &cfg, nil
}
