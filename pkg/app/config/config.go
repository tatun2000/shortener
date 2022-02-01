package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	HttpServerHost string `toml:"HTTP_SERVER_HOST"`
	DbName         string `toml:"DB_NAME"`
	BucketName     string `toml:"BUCKET_NAME"`
	RedirectNum    int    `toml:"REDIRECT_NUM"`
}

func GetConf() *Config {
	var conf Config
	_, err := toml.DecodeFile(os.Args[1], &conf)
	if err != nil {
		log.Fatal(err)
	}
	return &conf
}
