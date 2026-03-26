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
	"bufio"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"sync"

	"github.com/go-logr/logr"
)

type CommandLogger struct {
	logger         logr.Logger
	stdErr, stdOut io.Reader
	wg             *sync.WaitGroup
}

func NewCommandLogger(cmd *exec.Cmd, logger logr.Logger) (*CommandLogger, error) {
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("could not obtain a pipe to stderr")
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("could not obtain a pipe to stdout")
	}

	cl := CommandLogger{
		logger: logger,
		stdErr: stderr,
		stdOut: stdout,
		wg:     &sync.WaitGroup{},
	}

	return &cl, nil
}

func (cl *CommandLogger) Wait() error {
	const goroutines = 2

	cl.wg.Add(goroutines)
	errs := make(chan error, goroutines)

	go cl.write(cl.stdErr, "stderr", errs)
	go cl.write(cl.stdOut, "stdout", errs)

	cl.wg.Wait()
	close(errs)

	chErrs := make([]error, 0, goroutines)

	for err := range errs {
		chErrs = append(chErrs, err)
	}

	return errors.Join(chErrs...)
}

func (cl *CommandLogger) write(r io.Reader, name string, errs chan<- error) {
	defer cl.wg.Done()

	logger := cl.logger.WithName(name)

	s := bufio.NewScanner(r)

	for s.Scan() {
		logger.Info(s.Text())
	}

	if err := s.Err(); err != nil {
		errs <- err
	}
}
