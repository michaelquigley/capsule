package capsule

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type Config struct {
	AttributeHandlers []AttributeHandler
}

type AttributeHandler func(string, fs.DirEntry) map[string]interface{}

func (cfg *Config) GetAttributes(path string, de fs.DirEntry) map[string]interface{} {
	merged := make(map[string]interface{})
	for _, handler := range cfg.AttributeHandlers {
		for k, v := range handler(path, de) {
			merged[k] = v
		}
	}
	return merged
}

func DefaultConfig() *Config {
	return &Config{
		AttributeHandlers: []AttributeHandler{
			filenameClassType,
			filenameRole,
		},
	}
}

func filenameRole(path string, _ fs.DirEntry) map[string]interface{} {
	base := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	if base == "story" {
		return map[string]interface{}{"role": "story"}
	}
	return nil
}

func filenameClassType(path string, _ fs.DirEntry) map[string]interface{} {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".md":
		return map[string]interface{}{"class": "document", "type": "markdown"}
	case ".png":
		return map[string]interface{}{"class": "image", "type": "png"}
	case ".jpg":
		return map[string]interface{}{"class": "image", "type": "jpeg"}
	default:
		return nil
	}
}
