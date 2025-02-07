package routes

import (
	"pokeShowcase-api/internal/card/controllers"

	"github.com/gin-gonic/gin"
)

func CardRoutes(r *gin.Engine) {
	r.POST("/card", controllers.PostProduct)
}
