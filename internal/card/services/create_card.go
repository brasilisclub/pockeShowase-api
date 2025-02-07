package services

import (
	"pokeShowcase-api/internal/card"
	"pokeShowcase-api/internal/database"
)

func CreateCard(crb card.CardRequestBody) (card.Card, error) {
	var dbCard card.Card

	dbCard.Name = crb.Name
	dbCard.Collection = crb.Collection
	dbCard.Rarity = crb.Rarity
	dbCard.ImageUrl = crb.ImageUrl

	c := database.GetConnector()

	err := c.Create(&dbCard).Error
	if err != nil {
		return card.Card{}, err
	}

	return dbCard, nil
}
