package database

import (
	"fmt"
	"pokeShowcase-api/configs"
)

func getDsn() string {
	if configs.IsApplicationInTestMode() {
		return ":memory:"
	}

	user := configs.Envs.DATABASE_USER
	password := configs.Envs.DATABASE_PASSWORD
	host := configs.Envs.DATABASE_HOST
	port := configs.Envs.DATABASE_PORT

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/ordine?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port)
	return dsn
}
