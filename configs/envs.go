package configs

import "github.com/kelseyhightower/envconfig"

type envs struct {
	ENVIRONMENT       string `env:"ENVIRONMENT" default:"local"`
	DATABASE_HOST     string `json:"DATABASE_HOST" env:"DATABASE_HOST" default:"database"`
	DATABASE_PASSWORD string `json:"DATABASE_PASSWORD" env:"DATABASE_PASSWORD" default:"root"`
	DATABASE_PORT     string `json:"DATABASE_PORT" env:"DATABASE_PORT" default:"3306"`
	DATABASE_USER     string `json:"DATABASE_USER" env:"DATABASE_USER" default:"root"`
	PORT              string `env:"PORTS" default:"8080"`
	S3_BUCKET_NAME    string `env:"S3_BUCKET_NAME" default:"pokeshowcase-api"`
	AWS_REGION        string `env:"AWS_REGION" default:"us-east-1"`
	LOG_LEVEL         string `env:"LOG_LEVEL" default:"info"`
}

var (
	Envs = &envs{}
)

const Testing = "testing"

func IsApplicationInTestMode() bool {
	return Envs.ENVIRONMENT == Testing
}

func LoadEnvs() {
	envconfig.Process("", Envs)
}

func init() {
	LoadEnvs()
}
