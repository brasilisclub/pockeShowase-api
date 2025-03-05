package controllers

import (
	"errors"
	"net/http"
	"pokeShowcase-api/internal/card"
	"pokeShowcase-api/internal/card/services"
	"pokeShowcase-api/utils"

	"github.com/gin-gonic/gin"
)

func GetCard(ctx *gin.Context) {

	var cardId string

	cardId = ctx.Param("id")

	Card, err := services.GetCardById(cardId)
	if err != nil {
		if errors.Is(err, card.ErrorCardNotFounded) {

			ctx.JSON(
				http.StatusNotFound,
				utils.GenericResponse{
					Message: err.Error(),
				},
			)
			return
		}
		ctx.JSON(
			http.StatusInternalServerError,
			utils.GenericResponse{
				Message: err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		Card,
	)
}
