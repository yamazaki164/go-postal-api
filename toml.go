package main

import (
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port     int    `toml:"port"`
	Endpoint string `toml:"endpoint"`
	JsonDir  string `toml:"json_dir"`
}

func (c *Config) isValidDir(dir string) bool {
	st, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return st.IsDir()
}

func (c *Config) isValidPort(p int) bool {
	return 0 < p && p <= 65535
}

func (c *Config) isValidEndpoint(u string) bool {
	if _, e := url.Parse(u); e != nil {
		return false
	}
	return true
}

func (c *Config) IsValidConfig() bool {
	return c.isValidDir(c.JsonDir) && c.isValidPort(c.Port) && c.isValidEndpoint(c.Endpoint)
}

func (c *Config) BindAddress() string {
	return ":" + strconv.Itoa(c.Port)
}

func (c *Config) JsonFile(code string) string {
	return filepath.Join(c.JsonDir, code+".json")
}

func LoadToml(file string) (*Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(file, &conf); err != nil {
		return &conf, err
	}

	return &conf, nil
}
