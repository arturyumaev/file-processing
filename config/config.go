package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Mode                      string `yaml:"mode"`
	ApplicationHandlerTimeout uint   `yaml:"application_handler_timeout"`
	Server                    struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
}

func Read(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) IsProduction() bool {
	return c.Mode == "production"
}
