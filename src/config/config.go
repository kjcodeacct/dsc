package config

import (
	_ "gopkg.in/yaml.v2"
)

type Config struct {
	Name   string `yaml:"name"`
	Email  string `yaml:"Email"`
	Editor string `yaml:"editor"`
}

func Get() (*Config, error) {

	newConfig := &Config{
		Name:   "test",
		Email:  "email@example.com",
		Editor: "subl",
	}

	return newConfig, nil
}
