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

func (cfg *Config) PropertyType(path string, de fs.DirEntry) Attributes {
	merged := Attributes{}
	for _, handler := range cfg.AttributeHandlers {
		attrs := handler(path, de)
		if attrs.Class != "" {
			merged.Class = attrs.Class
		}
		if attrs.Role != "" {
			merged.Role = attrs.Role
		}
		if attrs.Type != "" {
			merged.Type = attrs.Type
		}
	}
	return merged
}

func filenameRole(path string, _ fs.DirEntry) Attributes {
	base := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	if base == "story" {
		return Attributes{Role: "story"}
	}
	return Attributes{}
}

func filenameClassType(path string, _ fs.DirEntry) Attributes {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".md":
		return Attributes{Class: "document", Type: "markdown"}
	case ".png":
		return Attributes{Class: "image", Type: "png"}
	case ".jpg":
		return Attributes{Class: "image", Type: "jpeg"}
	default:
		return Attributes{}
	}
}
