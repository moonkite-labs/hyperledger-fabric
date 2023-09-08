package utils

import (
	"fabric-gateway/config"

	"github.com/joho/godotenv"
)

const ENV_FILE = "../.env.test"

func SetupEnv(env_file string) (*config.Config, error) {
	err := godotenv.Load(env_file)
	if err != nil {
		return nil, err
	}

	var cfg config.Config

	if err := cfg.ParseEnv(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
