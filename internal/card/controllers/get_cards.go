package controllers

import (
	"fmt"
	"net/http"
	"pokeShowcase-api/internal/card"
	"pokeShowcase-api/internal/card/services"
	"pokeShowcase-api/utils"

	"github.com/gin-gonic/gin"
)

func GetCards(ctx *gin.Context) {

	cards, err := services.GetAllCards()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			utils.GenericResponse{
				Message: fmt.Sprintf("Internal error trying to get cards %s", err.Error()),
			})
		return
	}

	ctx.JSON(http.StatusOK, card.Cards{
		Cards: cards,
	})

}
