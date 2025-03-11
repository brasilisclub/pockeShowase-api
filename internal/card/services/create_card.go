package services

import (
	"pokeShowcase-api/internal/card"
	"pokeShowcase-api/internal/card/services/helpers"
	"pokeShowcase-api/internal/database"
)

func CreateCard(crb card.CardRequestBody) (card.Card, error) {
	var dbCard card.Card

	dbCard = helpers.NewDbCard(crb)

	c := database.GetConnector()

	err := c.Create(&dbCard).Error
	if err != nil {
		return card.Card{}, err
	}

	return dbCard, nil
}
