package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"pokeShowcase-api/internal/card"
	"pokeShowcase-api/internal/card/services"
	"pokeShowcase-api/utils"

	"github.com/gin-gonic/gin"
)

func PutCard(ctx *gin.Context) {
	var requestBody card.CardRequestBody

	cardData := ctx.PostForm("card")
	err := json.Unmarshal([]byte(cardData), &requestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenericResponse{
			Message: fmt.Sprintf("Invalid body %s", err.Error()),
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

	dbCard, err := services.UpdateCard(ctx.Param("id"), requestBody, file)

	if err != nil {
		if errors.Is(err, card.ErrorCardNotFounded) {
			ctx.JSON(http.StatusNotFound, utils.GenericResponse{
				Message: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.GenericResponse{
			Message: fmt.Sprintf("internal server error %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dbCard)

}
