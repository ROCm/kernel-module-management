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

package registry

import (
	"archive/tar"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type uncompressedLayer struct {
	diffID    v1.Hash
	mediaType types.MediaType
	content   []byte
}

func (ul *uncompressedLayer) DiffID() (v1.Hash, error) {
	return ul.diffID, nil
}

func (ul *uncompressedLayer) Uncompressed() (io.ReadCloser, error) {
	return io.NopCloser(bytes.NewBuffer(ul.content)), nil
}

func (ul *uncompressedLayer) MediaType() (types.MediaType, error) {
	return ul.mediaType, nil
}

func prepareLayer(fileName string, data []byte) (v1.Layer, error) {
	// Hash the contents as we write it out to the buffer.
	var b bytes.Buffer
	hasher := sha256.New()
	mw := io.MultiWriter(&b, hasher)

	// Write a single file with a random name and random contents.
	tw := tar.NewWriter(mw)
	if err := tw.WriteHeader(&tar.Header{
		Name:     fileName,
		Size:     int64(len(data)),
		Typeflag: tar.TypeReg,
	}); err != nil {
		return nil, err
	}
	if _, err := io.Copy(tw, bytes.NewReader(data)); err != nil {
		return nil, err
	}
	if err := tw.Close(); err != nil {
		return nil, err
	}

	h := v1.Hash{
		Algorithm: "sha256",
		Hex:       hex.EncodeToString(hasher.Sum(make([]byte, 0, hasher.Size()))),
	}

	return partial.UncompressedToLayer(&uncompressedLayer{
		diffID:    h,
		mediaType: types.DockerLayer,
		content:   b.Bytes(),
	})
}
