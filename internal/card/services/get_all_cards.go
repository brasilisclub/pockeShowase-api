package services

import (
	"pokeShowcase-api/internal/card"
	"pokeShowcase-api/internal/database"
)

func GetAllCards() ([]card.Card, error) {

	var cards []card.Card

	c := database.GetConnector()

	result := c.Find(&cards)
	if result.Error != nil {
		return []card.Card{}, result.Error
	}

	return cards, nil
}
