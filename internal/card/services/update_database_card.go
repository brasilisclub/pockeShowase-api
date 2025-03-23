package services

import (
	"pokeShowcase-api/internal/card"
	"pokeShowcase-api/internal/card/services/helpers"
	"pokeShowcase-api/internal/database"
)

func UpdateDatabaseCard(id string, crb card.CardRequestBody) (card.Card, error) {
	var dbCard card.Card

	dbCard, err := GetCardById(id)
	if err != nil {
		return dbCard, err
	}

	dbCard = helpers.UpdateDbCard(crb, dbCard)

	c := database.GetConnector()

	err = c.Save(&dbCard).Error

	return dbCard, err
}
