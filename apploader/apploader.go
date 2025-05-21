package apploader

import (
	"cloud-sek/cache"
	"cloud-sek/globals"
	"cloud-sek/models"
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func Init() {
	cache.Init()
	configPath := getConfigPath()
	LoadConfig(configPath)
}

func getConfigPath() string {
	var path string
	flag.StringVar(&path, "config", "config.yaml", "path to config file")
	flag.Parse()
	if path == "" {
		log.Fatal("config file path cannot be empty")
	}
	return path
}

func LoadConfig(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	config := &models.Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}
	globals.Config = config
}
