package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.GET("/", controllers.GetUsers);
	r.POST("/", controllers.CreateUser);
}