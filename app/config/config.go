package config

import "os"

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Host     string
	Name     string
	Port     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Username: "ldp",
			Password: os.Getenv("LDP_PW"),
			Host:     "glintcore.net",
			Name:     "ldp_folio_release",
			Port:     "5432",
		},
	}
}
