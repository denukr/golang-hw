package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Не удалось загрузить переменную окружения")
	}
	return &Config{
		Key: os.Getenv("KEY"),
	}
}
