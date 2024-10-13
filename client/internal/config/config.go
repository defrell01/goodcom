package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Directory  string   `yaml:"directory`
	Extensions []string `yaml:extensions`
}

func LoadConfig(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
