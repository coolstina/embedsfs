package embedsfs

import (
	"embed"
	"path/filepath"
)

// EmbedsFS custom definition structure.
type EmbedsFS struct {
	embeds embed.FS
	path   string
}

// NewEmbedsFS returns a new EmbedsFS instance
func NewEmbedsFS(embeds embed.FS, path string) *EmbedsFS {
	return &EmbedsFS{embeds: embeds, path: path}
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
	return filepath.Join(embeds.path, name)
}
