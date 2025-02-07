package controllers

import (
	"fmt"
	"net/http"
	"pokeShowcase-api/internal/card"
	"pokeShowcase-api/internal/card/services"
	"pokeShowcase-api/utils"

	"github.com/gin-gonic/gin"
)

func PostProduct(ctx *gin.Context) {

	var card card.CardRequestBody
	err := ctx.ShouldBindBodyWithJSON(&card)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenericResponse{
			Message: fmt.Sprintf("Invalid body, %s", err.Error()),
		})
		return
	}

	createdCard, err := services.CreateCard(card)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.GenericResponse{
			Message: fmt.Sprintf("Internal error trying to create card %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdCard)

}
