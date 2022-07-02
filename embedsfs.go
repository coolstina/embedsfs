// Copyright 2022 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package embedsfs

import (
	"embed"
	"path/filepath"
	"strings"
)

// EmbedsFS custom definition structure.
type EmbedsFS struct {
	embeds embed.FS
	path   string
}

// NewEmbedsFS returns a new EmbedsFS instance
func NewEmbedsFS(embeds embed.FS, options ...Option) *EmbedsFS {
	ops := &option{}
	for _, o := range options {
		o.apply(ops)
	}

	return &EmbedsFS{embeds: embeds, path: ops.path}
}

// Embeds returns a embeds filesystem structure
func (embeds *EmbedsFS) Embeds() embed.FS {
	return embeds.embeds
}

// Content returns embed filename contains embed.FS fill path
func (embeds *EmbedsFS) Content(name string) ([]byte, error) {
	data, err := embeds.embeds.ReadFile(embeds.Filename(name))
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Filename returns embed filename contains embed.FS fill path
func (embeds *EmbedsFS) Filename(name string) string {
	var full = name
	if embeds.path != "" {
		has := strings.HasPrefix(full, embeds.path)
		if !has {
			full = filepath.Join(embeds.path, full)
		}
	}

	return full
}
