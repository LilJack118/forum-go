package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Config(key string, t string) (interface{}, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	envVar := os.Getenv(key)

	switch t {
	case "string":
		return envVar, nil
	case "int":
		envVarInt, err := strconv.Atoi(envVar)
		if err != nil {
			return nil, err
		}
		return envVarInt, nil
	default:
		return nil, fmt.Errorf("unsupported environment variable type: %s", t)
	}
}
