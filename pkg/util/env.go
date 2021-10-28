package util

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var EnvPrefix = "staff_management"

func ReadEnvVars(args ...interface{}) error {
	if len(args) == 1 {
		return envconfig.Process(EnvPrefix, args[0])
	}

	if len(args) != 2 {
		return errors.New("invalid env params")
	}

	if prefix, ok := args[1].(string); ok {
		return envconfig.Process(prefix, args[0])
	}

	return errors.New("invalid env params")
}


func LoadEnv(file string) error {
	if file == "" {
		return godotenv.Load()
	}

	return godotenv.Load(file)
}
