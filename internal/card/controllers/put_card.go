package controllers

import (
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

	err := ctx.ShouldBindBodyWithJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenericResponse{
			Message: fmt.Sprintf("Invalid body %s", err.Error()),
		})
		return
	}

	dbCard, err := services.UpdateCard(ctx.Param("id"), requestBody)

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
	}

	ctx.JSON(http.StatusCreated, dbCard)

}
