package capsule

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func DefaultConfig() *Config {
	return &Config{
		TypeHandlers: []PropertyTypeHandler{
			FileExtensionPropertyType,
		},
	}
}

func (cfg *Config) PropertyType(path string, de fs.DirEntry) (string, bool) {
	for _, handler := range cfg.TypeHandlers {
		typeId, found := handler(path, de)
		if found {
			return typeId, true
		}
	}
	return "", false
}

type Config struct {
	TypeHandlers []PropertyTypeHandler
}

type PropertyTypeHandler func(string, fs.DirEntry) (string, bool)

func FileExtensionPropertyType(path string, de fs.DirEntry) (string, bool) {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".md":
		return "story; markdown", true
	case ".png":
		return "figure; png", true
	case ".jpg":
		return "figure; jpeg", true
	default:
		return "unknown", true
	}
}
