package config

import (
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	*toml.Tree
}

const DefaultFileName = "config.toml"

var config *Config

func init() {
	config = DefaultConfig()
}

func CFG() *Config {
	return config
}

func NewConfig(name string) *Config {
	file, err := os.OpenFile(name, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil
	}
	tree, err := toml.LoadReader(file)
	if err != nil {
		return nil
	}
	return &Config{
		Tree: tree,
	}
}

func DefaultConfig() *Config {
	return NewConfig(DefaultFileName)
}

func (c *Config) GetTree(name string) *toml.Tree {
	if c == nil {
		return nil
	}
	if v := c.Get(name); v != nil {
		if tree, b := v.(*toml.Tree); b {
			return tree
		}
	}
	return nil
}

func (c *Config) GetSub(name string) *Config {
	if v := c.GetTree(name); v != nil {
		return &Config{
			Tree: v,
		}
	}
	return (*Config)(nil)
}

func (c *Config) GetString(name string) string {
	if v := c.Get(name); v != nil {
		if v1, b := v.(string); b {
			return v1
		}
	}
	return ""
}

func (c *Config) GetStringD(name string, def string) string {
	if c == nil {
		return def
	}
	if v := c.Get(name); v != nil {
		if v1, b := v.(string); b {
			return v1
		}
	}
	return def
}

func (c *Config) GetBool(name string) bool {
	if c == nil {
		return false
	}
	if v := c.Get(name); v != nil {
		if v1, b := v.(bool); b {
			return v1
		}
	}
	return false
}

func (c *Config) GetBoolD(name string, def bool) bool {
	if v := c.Get(name); v != nil {
		if v1, b := v.(bool); b {
			return v1
		}
	}
	return def
}
