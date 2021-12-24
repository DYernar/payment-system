package main

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type config struct {
	Server struct {
		Host          string        `yaml:"host"`
		Port          int           `yaml:"port"`
		AccessSecret  string        `yaml:"accessSecret"`
		RefreshSecret string        `yaml:"refreshSecret"`
		AccessTtl     time.Duration `yaml:"accessTtl"`
		RefreshTtl    time.Duration `yaml:"refreshTtl"`
		Db            struct {
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
		Kafka struct {
			Broker string `yaml:"broker"`
			Group  string `yaml:"group"`
			Topic  string `yaml:"topic"`
		}
		Redis struct {
			Addr string `yaml:"addr"`
		}
		Grpc struct {
			Port int `yaml:"port"`
		} `yaml:"grpc"`
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
