package helpers

import (
	"fmt"
	"pokeShowcase-api/internal/card"
)

func GetCardPath(dbCard card.Card) string {
	return fmt.Sprintf("%s/%s-%s.jpg", dbCard.Collection, dbCard.CardId, dbCard.Name)
}
