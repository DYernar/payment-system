package main

import (
	"flag"
	"fmt"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")

	flag.Parse()
	config, err := NewConfig(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	app := NewApplication(config)

	app.run()
}
