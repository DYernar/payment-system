package main

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Db   struct {
			Dsn          string `yaml:"dsn"`
			MaxOpenConns int    `yaml:"maxOpenConns"`
			MaxIdleConns int    `yaml:"maxIdleConns"`
			MaxIdleTime  string `yaml:"maxIdleTime"`
		} `yaml:"db"`
		Timeout struct {
			Server time.Duration `yaml:"server"`
			Write  time.Duration `yaml:"write"`
			Read   time.Duration `yaml:"read"`
			Idle   time.Duration `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`
}

func NewConfig(configPath string) (*config, error) {
	config := &config{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
