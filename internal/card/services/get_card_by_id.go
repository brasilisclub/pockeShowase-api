package services

import (
	"errors"
	"pokeShowcase-api/internal/card"
	"pokeShowcase-api/internal/database"

	"gorm.io/gorm"
)

func GetCardById(id string) (card.Card, error) {
	var dbCard card.Card

	c := database.GetConnector()

	result := c.First(&dbCard, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return dbCard, card.ErrorCardNotFounded
	}

	return dbCard, result.Error
}
