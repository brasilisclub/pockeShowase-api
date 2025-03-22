package services

import (
	"pokeShowcase-api/internal/aws/s3"
	"pokeShowcase-api/internal/card/services/helpers"
)

func DeleteCardById(id string) error {
	card, err := GetCardById(id)
	if err != nil {
		return err
	}

	err = DeleteDatabaseCardById(id)
	if err != nil {
		return err
	}

	err = s3.DeleteBucketObject(helpers.GetCardPath(card))
	return err
}
