package capsule

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type Config struct {
	AttributeHandlers []AttributeHandler
}

type AttributeHandler func(string, fs.DirEntry) Attributes

func DefaultConfig() *Config {
	return &Config{
		AttributeHandlers: []AttributeHandler{
			filenameClassType,
			filenameRole,
		},
	}
}

func (cfg *Config) GetAttributes(path string, de fs.DirEntry) Attributes {
	merged := make(Attributes)
	for _, handler := range cfg.AttributeHandlers {
		for k, v := range handler(path, de) {
			merged[k] = v
		}
	}
	return merged
}

func filenameRole(path string, _ fs.DirEntry) Attributes {
	base := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	if base == "story" {
		return Attributes{"role": "story"}
	}
	return nil
}

func filenameClassType(path string, _ fs.DirEntry) Attributes {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".md":
		return Attributes{"class": "document", "type": "markdown"}
	case ".png":
		return Attributes{"class": "image", "type": "png"}
	case ".jpg":
		return Attributes{"class": "image", "type": "jpeg"}
	default:
		return nil
	}
}
