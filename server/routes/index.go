package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	MainRoutes(r.Group("/"));
	UserRoutes(r.Group("/user"));
}
