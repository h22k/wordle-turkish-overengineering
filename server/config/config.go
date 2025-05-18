package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppKey  string
	AppName string
	AppEnv  string

	ServerPort string

	DBConnection string
	DbHost       string
	DbPort       string
	DbUser       string
	DbPassword   string
	DbName       string

	MaxDbConns        int32
	MinDbConns        int32
	MaxDbIdleTime     time.Duration
	MaxDbConnLifeTime time.Duration

	AllowOrigins []string
	CookieDomain string

	DbUrl string

	PyroscopeUrl string
}

func LoadConfig() Config {
	envFile := os.Getenv("ENV_FILE")

	if err := godotenv.Load(envFile); err != nil {
		log.Println("Error loading .env file")
	}

	conf := Config{
		AppKey:       getValue("APP_KEY", ""),
		AppName:      getValue("APP_NAME", "wordle"),
		AppEnv:       getValue("APP_ENV", "development"),
		ServerPort:   getValue("PORT", ":8080"),
		DBConnection: getValue("DB_CONNECTION", "postgres"),
		DbHost:       getValue("DB_HOST", "localhost"),
		DbPort:       getValue("DB_PORT", "3306"),
		DbUser:       getValue("DB_USERNAME", "root"),
		DbPassword:   getValue("DB_PASSWORD", "password"),
		DbName:       getValue("DB_DATABASE", "wordle"),
		MaxDbConns:   getValue("MAX_DB_CONNECTIONS", int32(10)),
		MinDbConns:   getValue("MIN_DB_CONNECTIONS", int32(1)),
		PyroscopeUrl: getValue("PYROSCOPE_URL", ""),
		AllowOrigins: strToArraySeparatedByComma(getValue("ALLOW_ORIGINS", "")),
		CookieDomain: getValue("COOKIE_DOMAIN", ""),
	}

	conf.DbUrl = conf.DBConnection + "://" + conf.DbUser + ":" + conf.DbPassword + "@" + conf.DbHost + ":" + conf.DbPort + "/" + conf.DbName
	conf.MaxDbConnLifeTime = time.Hour
	conf.MaxDbIdleTime = 10 * time.Minute

	return conf
}

func (c Config) IsProd() bool {
	return c.AppEnv == "production"
}

func strToArraySeparatedByComma(str string) []string {
	if str == "" {
		return nil
	}

	result := make([]string, 0)
	for _, s := range strings.Split(str, ",") {
		result = append(result, strings.TrimSpace(s))
	}

	return result
}

func getValue[T string | int | int32](key string, def T) T {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return def
	}

	var result any
	var err error

	switch any(def).(type) {
	case string:
		result = value
	case int32:
		var i int64
		i, err = strconv.ParseInt(value, 10, 64)
		result = int32(i)
	default:
		panic("unsupported type")
	}

	if err != nil {
		panic(fmt.Sprintf("invalid env value for %s: %v", key, err))
	}

	return result.(T)
}
