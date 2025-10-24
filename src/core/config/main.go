package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/golobby/dotenv"
)

var DB DBConfig
var App AppConfig

type configs struct {
	*DBConfig
	*AppConfig
}

func init() {
	envFile := ".env"

	set(nil, &configs{&DB, &App}) // load default value first

	// It should still work even if the .env file does not exist
	if file, err := os.Open(envFile); err == nil {
		set(file, &configs{&DB, &App})
		defer file.Close()
	}
}

func set(file *os.File, structure interface{}) {
	if err := env.Parse(structure); err != nil {
		fmt.Printf("%+v\n", err)
	}

	if err := dotenv.NewDecoder(file).Decode(structure); err != nil {
	}
}
