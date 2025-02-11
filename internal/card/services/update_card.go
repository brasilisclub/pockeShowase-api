package services

import (
	"pokeShowcase-api/internal/card"
	"pokeShowcase-api/internal/database"
)

func UpdateCard(id string, crb card.CardRequestBody) (card.Card, error) {
	var dbProduct card.Card

	dbProduct, err := GetCardById(id)
	if err != nil {
		return dbProduct, err
	}

	dbProduct.Name = crb.Name
	dbProduct.Collection = crb.Collection
	dbProduct.Rarity = crb.Rarity
	dbProduct.ImageUrl = crb.ImageUrl

	c := database.GetConnector()

	err = c.Save(&dbProduct).Error

	return dbProduct, err
}
