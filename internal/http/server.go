package http

import (
	"pokeShowcase-api/internal/http/routes"

	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	r := gin.Default()
	routes.Load(r)
	return r
}
