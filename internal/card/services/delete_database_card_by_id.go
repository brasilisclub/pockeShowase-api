package services

import (
	"fmt"
	"pokeShowcase-api/internal/card"
	"pokeShowcase-api/internal/database"
)

func DeleteDatabaseCardById(id string) error {
	c := database.GetConnector()

	result := c.Delete(&card.Card{}, id)

	if result.Error != nil {
		return fmt.Errorf("error deleting ordine: %s", result.Error)
	}

	if result.RowsAffected == 0 {
		return card.ErrorCardNotFounded
	}
	return nil
}
