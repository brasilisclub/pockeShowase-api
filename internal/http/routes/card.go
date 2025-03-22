package routes

import (
	"pokeShowcase-api/internal/card/controllers"

	"github.com/gin-gonic/gin"
)

func CardRoutes(r *gin.Engine) {
	r.DELETE("/card/:id", controllers.DeleteCard)
	r.GET("/cards", controllers.GetCards)
	r.GET("/card/:id", controllers.GetCard)
	r.POST("/card", controllers.PostCard)
	r.PUT("/card/:id", controllers.PutCard)
}
