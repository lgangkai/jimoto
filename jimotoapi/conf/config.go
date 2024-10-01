package conf

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server      Server      `yaml:"server"`
	Micro       Micro       `yaml:"micro"`
	Etcd        Etcd        `yaml:"etcd"`
	ImageServer ImageServer `yaml:"image-server"`
}

type Server struct {
	Addr   string `yaml:"addr"`
	Scheme string `yaml:"scheme"`
}

type Micro struct {
	Commodity Commodity `yaml:"commodity"`
	Account   Account   `yaml:"account"`
}

type Commodity struct {
	Name string `yaml:"name"`
}

type Account struct {
	Name string `yaml:"name"`
}

type Etcd struct {
	Addr string `yaml:"addr"`
}

type ImageServer struct {
	Local     bool   `yaml:"local"`
	LocalPath string `yaml:"local-path"`
}

func LoadConfig(confPath string) (*Config, error) {
	config := &Config{}
	data, err := os.ReadFile(confPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
