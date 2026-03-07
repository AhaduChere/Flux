package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Endpoint struct {
	Path   string
	Method string
}

type Config struct {
	URL       string
	Endpoints map[string]Endpoint
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
