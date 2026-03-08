package config

import (
	"github.com/BurntSushi/toml"
)

type Endpoint struct {
	Path   string            `toml:"path"`
	Method string            `toml:"method"`
	Body   map[string]string `toml:"body"`
}

type Config struct {
	URL       string              `toml:"url"`
	Endpoints map[string]Endpoint `toml:"req"`
}

// NOTE: Reads Fluxfile
func Parsefile(path string) (Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
