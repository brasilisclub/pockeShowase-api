package routes

import (
	"pokeShowcase-api/internal/card/controllers"

	"github.com/gin-gonic/gin"
)

func CardRoutes(r *gin.Engine) {
	r.GET("/cards", controllers.GetCards)
	r.POST("/card", controllers.PostProduct)
	r.PUT("/card/:id", controllers.PutCard)
}
