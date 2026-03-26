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
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"github.com/go-logr/logr"
)

//go:generate mockgen -source=filesystem_helper.go -package=utils -destination=mock_filesystem_helper.go

type FSHelper interface {
	RemoveSrcFilesFromDst(srcDir, dstDir string) error
	FileExists(root, fileRegex string) (bool, error)
}

type fsHelper struct {
	logger logr.Logger
}

func NewFSHelper(logger logr.Logger) FSHelper {
	return &fsHelper{
		logger: logger,
	}
}

func (fh *fsHelper) RemoveSrcFilesFromDst(srcDir, dstDir string) error {
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relPath, err := filepath.Rel(srcDir, path)
			if err != nil {
				fh.logger.Info(WarnString("failed to get relative path"), "srcDir", srcDir, "path", path, "error", err)
				return nil
			}
			fileToRemove := filepath.Join(dstDir, relPath)
			fh.logger.Info("Removing dst file", "file", fileToRemove)
			err = os.Remove(fileToRemove)
			if err != nil {
				fh.logger.Info(WarnString("failed to delete file"), "file", fileToRemove, "error", err)
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to remove files %s/* from %s\n", srcDir, dstDir)
	}
	return nil
}

func (fh *fsHelper) FileExists(root, fileRegex string) (bool, error) {
	regex, err := regexp.Compile(fileRegex)
	if err != nil {
		return false, fmt.Errorf("failed to compile regex %s: %v", fileRegex, err)
	}

	found := false
	// Walk through the directory
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Match file names against the regex
		if !d.IsDir() && regex.MatchString(d.Name()) {
			found = true
			return fs.SkipAll
		}
		return nil
	})

	return found, err
}
