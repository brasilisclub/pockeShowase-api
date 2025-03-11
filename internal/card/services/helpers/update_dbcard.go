package helpers

import "pokeShowcase-api/internal/card"

func UpdateDbCard(crb card.CardRequestBody, card card.Card) card.Card {

	dbCard := NewDbCard(crb)
	dbCard.ID = card.ID
	dbCard.CreatedAt = card.CreatedAt
	dbCard.DeletedAt = card.DeletedAt
	dbCard.UpdatedAt = card.UpdatedAt

	return dbCard

}
