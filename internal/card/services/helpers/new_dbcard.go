package helpers

import (
	"fmt"
	"pokeShowcase-api/configs"
	"pokeShowcase-api/internal/card"
)

func NewDbCard(crb card.CardRequestBody) card.Card {
	return card.Card{
		CardId:     crb.CardId,
		Name:       crb.Name,
		Collection: crb.Collection,
		Rarity:     crb.Rarity,
		ImageUrl:   fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s/%s-%s.jpg", configs.Envs.S3_BUCKET_NAME, configs.Envs.AWS_REGION, crb.Collection, crb.CardId, crb.Name),
	}
}
