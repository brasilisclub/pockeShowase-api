package helpers

import (
	"fmt"
	"pokeShowcase-api/internal/card"
)

func NewDbCard(crb card.CardRequestBody) card.Card {
	return card.Card{
		CardId:     crb.CardId,
		Name:       crb.Name,
		Collection: crb.Collection,
		Rarity:     crb.Rarity,
		ImageUrl:   fmt.Sprintf("%s/%s-%s.jpg", crb.Collection, crb.CardId, crb.Name),
	}
}
