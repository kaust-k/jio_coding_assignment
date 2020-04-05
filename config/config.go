package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config stores values of parsed environment variables
type Config struct {
	HTTPPort          string        `envconfig:"http_port" default:"3000"`
	DatabaseUser      string        `envconfig:"database_user" required:"true"`
	DatabasePassword  string        `envconfig:"database_password" required:"true"`
	DatabaseHost      string        `envconfig:"database_host" required:"true"`
	DatabasePort      string        `envconfig:"database_port" required:"true"`
	DatabaseName      string        `envconfig:"database_name" required:"true"`
	RedisAddress      string        `envconfig:"redis_address" required:"true"`
	RedisPassword     string        `envconfig:"redis_password"`
	RedisDB           int           `envconfig:"redis_db" default:"0"`
	JwtSigningSecret  string        `envconfig:"jwt_signing_secret" required:"true"`
	JwtExpiryDuration time.Duration `envconfig:"jwt_expiry_duration" default:"30m"`
}

var config Config

func init() {
	parseEnvVars()
}

func parseEnvVars() {
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// GetConfig to access values of parsed environment variables
func GetConfig() Config {
	return config
}
