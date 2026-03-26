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
	"os/exec"

	"github.com/go-logr/logr"
)

//go:generate mockgen -source=modprobe.go -package=worker -destination=mock_modprobe.go

type ModprobeRunner interface {
	Run(ctx context.Context, args ...string) error
}

type modprobeRunnerImpl struct {
	logger logr.Logger
}

func NewModprobeRunner(logger logr.Logger) ModprobeRunner {
	return &modprobeRunnerImpl{logger: logger}
}

func (mr *modprobeRunnerImpl) Run(ctx context.Context, args ...string) error {
	cmd := exec.CommandContext(ctx, "modprobe", args...)

	cl, err := NewCommandLogger(cmd, mr.logger.WithName("modprobe"))
	if err != nil {
		return fmt.Errorf("could not create a command logger: %v", err)
	}

	mr.logger.Info("Running modprobe", "command", cmd.String())

	if err = cmd.Start(); err != nil {
		return fmt.Errorf("could not start modprobe: %v", err)
	}

	if err = cl.Wait(); err != nil {
		return fmt.Errorf("error while waiting on the command logger: %v", err)
	}

	if err = cmd.Wait(); err != nil {
		return fmt.Errorf("error while waiting on the command: %v", err)
	}

	return nil
}
