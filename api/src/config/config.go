package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvVariable interface {
	string | int
}

func ConfigStr(key string) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}

func ConfigInt(key string) (int, error) {
	if err := godotenv.Load(".env"); err != nil {
		return 0, err
	}

	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return 0, err
	}

	return val, nil
}
