package config

import (
	"log"
	"time"

	"github.com/jinzhu/configor"
)

// Config ...
type Config struct {
	Boilerplate struct {
		URL string `default:"https://github.com/takaaki-mizuno/go-boilerplate/archive/main.zip"`
	}
}

// LoadConfig ...
func LoadConfig() (*Config, error) {
	var config Config

	err := configor.
		New(&configor.Config{AutoReload: true, AutoReloadInterval: time.Minute}).
		Load(&config)

	if err != nil {
		log.Println(err)
		log.Fatal("Error loading config")
	}

	return &config, err
}
