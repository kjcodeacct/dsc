package config

import (
	_ "gopkg.in/yaml.v2"
)

type Config struct {
	Name   string `json:"name"`
	Email  string `json:"Email"`
	Editor string `json:"editor"`
}

func Get() (*Config, error) {

	newConfig := &Config{
		Name:   "test",
		Email:  "email@example.com",
		Editor: "subl",
	}

	return newConfig, nil
}
