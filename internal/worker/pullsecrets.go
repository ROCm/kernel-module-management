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
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/go-logr/logr"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/authn/kubernetes"
	v1 "k8s.io/api/core/v1"
)

func ReadKubernetesSecrets(ctx context.Context, rootDir string, logger logr.Logger) (authn.Keychain, error) {
	var secrets []v1.Secret

	err := filepath.WalkDir(rootDir, func(path string, de fs.DirEntry, err error) error {
		if err != nil || de.IsDir() {
			return err
		}

		var (
			sKey  string
			sType v1.SecretType
		)

		switch sKey = filepath.Base(path); sKey {
		case v1.DockerConfigKey:
			sType = v1.SecretTypeDockercfg
		case v1.DockerConfigJsonKey:
			sType = v1.SecretTypeDockerConfigJson
		default:
			logger.Info("Unhandled file name; ignoring", "path", path)
			return nil
		}

		logger.Info("Reading file", "path", path, "type", sType)

		b, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("could not read %s: %v", path, err)
		}

		s := v1.Secret{
			Type: sType,
			Data: map[string][]byte{sKey: b},
		}

		secrets = append(secrets, s)

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error while walking %q: %v", rootDir, err)
	}

	return kubernetes.NewFromPullSecrets(ctx, secrets)
}
