package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func MainRoutes(r *gin.RouterGroup) {
	r.GET("/", controllers.HandleRunServer);
}