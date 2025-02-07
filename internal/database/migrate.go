package database

import "pokeShowcase-api/internal/card"

func Migrate() {
	c := GetConnector()
	c.AutoMigrate(&card.Card{})
}
