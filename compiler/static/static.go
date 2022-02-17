package static

import (
	"github.com/michaelquigley/capsule"
)

type Config struct {
	BuildPath    string
	ResourcePath string
}

type compiler struct {
	cfg *Config
}

func New(cfg *Config) *compiler {
	return &compiler{cfg: cfg}
}

func (c *compiler) Compile(m *capsule.Model) error {
	return nil
}
