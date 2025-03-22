package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pokeShowcase-api/internal/card"
	"pokeShowcase-api/internal/card/services"
	"pokeShowcase-api/utils"

	"github.com/gin-gonic/gin"
)

func PostCard(ctx *gin.Context) {

	var card card.CardRequestBody

	cardData := ctx.PostForm("card")
	err := json.Unmarshal([]byte(cardData), &card)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenericResponse{
			Message: fmt.Sprintf("Invalid body, %s", err.Error()),
		})
		return
	}

	file, err := ctx.FormFile("image")

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.GenericResponse{
				Message: fmt.Sprintf("Error getting the file %s", err.Error()),
			},
		)
		return
	}

	createdCard, err := services.CreateCard(card, file)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.GenericResponse{
			Message: fmt.Sprintf("Internal error trying to create card %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdCard)

}
