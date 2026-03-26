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

package cmd

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/go-logr/logr"
)

func FatalError(l logr.Logger, err error, msg string, fields ...interface{}) {
	l.Error(err, msg, fields...)
	os.Exit(1)
}

func GetEnvOrFatalError(name string, logger logr.Logger) string {
	val := os.Getenv(name)
	if val == "" {
		FatalError(logger, errors.New("empty value"), "Could not get the environment variable", "name", name)
	}

	return val
}

func GitCommit() (string, error) {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return "", errors.New("build info is not available")
	}

	const vcsRevisionKey = "vcs.revision"

	for _, s := range bi.Settings {
		if s.Key == vcsRevisionKey {
			return s.Value, nil
		}
	}

	return "", fmt.Errorf("%s not found in build info settings", vcsRevisionKey)
}
