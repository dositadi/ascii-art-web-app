package config

import "os"

type Config struct {
	DBName   string
	Addr     string
	User     string
	Password string
	DBUrl    string
}

func (c *Config) GetDBConfig() Config {
	return Config{
		DBName:   os.Getenv("ASCII_DB_NAME"),
		Addr:     os.Getenv("ASCII_DB_ADDR"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBUrl:    os.Getenv("ASCII_DB_URL"),
	}
}
