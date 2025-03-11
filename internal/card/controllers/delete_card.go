package controllers

import (
	"errors"
	"net/http"
	"pokeShowcase-api/internal/card"
	"pokeShowcase-api/internal/card/services"
	"pokeShowcase-api/utils"

	"github.com/gin-gonic/gin"
)

func DeleteCard(ctx *gin.Context) {

	err := services.DeleteCardById(ctx.Param("id"))

	if err != nil {
		if errors.Is(card.ErrorCardNotFounded, err) {
			ctx.JSON(http.StatusNotFound, utils.GenericResponse{
				Message: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, utils.GenericResponse{
			Message: err.Error(),
		})
		return

	}

	ctx.JSON(http.StatusOK, utils.GenericResponse{
		Message: "card deleted successfully",
	})

}
