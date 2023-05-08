package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Port string `yaml:"port"`
}

func (conf *Conf) GetConf(path string) (*Conf, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
