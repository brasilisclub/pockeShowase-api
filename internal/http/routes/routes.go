package routes

import "github.com/gin-gonic/gin"

func Load(r *gin.Engine) {
	CardRoutes(r)
}
