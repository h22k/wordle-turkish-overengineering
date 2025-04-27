package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppKey string

	DBConnection string
	DbHost       string
	DbPort       string
	DbUser       string
	DbPassword   string
	DbName       string

	DbUrl string
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	conf := Config{
		AppKey:       getValue("APP_KEY", ""),
		DBConnection: getValue("DB_CONNECTION", "postgres"),
		DbHost:       getValue("DB_HOST", "localhost"),
		DbPort:       getValue("DB_PORT", "3306"),
		DbUser:       getValue("DB_USERNAME", "root"),
		DbPassword:   getValue("DB_PASSWORD", "password"),
		DbName:       getValue("DB_DATABASE", "wordle"),
	}

	conf.DbUrl = conf.DBConnection + "://" + conf.DbUser + ":" + conf.DbPassword + "@" + conf.DbHost + ":" + conf.DbPort + "/" + conf.DbName

	return conf
}

func getValue(key, def string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}

	if def == "" {
		panic("No value for " + key)
	}

	return def
}
