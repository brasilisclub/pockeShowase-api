package configs

import "github.com/kelseyhightower/envconfig"

type envs struct {
	ENVIRONMENT string `env:"ENVIRONMENT" default:"local"`
	PORT        string `env:"PORTS" default:"8080"`
	LOG_LEVEL   string `env:"LOG_LEVEL" default:"info"`
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
