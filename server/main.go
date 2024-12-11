package main

import (
	"os"
	"server/configs"
	"server/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		println(configs.Red, err, configs.Reset);
	}

	// Connect database
	configs.InitDatabase();

	GIN_MODE := os.Getenv("GIN_MODE"); // Server port

	if (GIN_MODE == "release") {
		gin.SetMode(gin.ReleaseMode);
	} else {
		gin.SetMode(gin.DebugMode);
	}

	router := gin.Default(); // Gin router

	PORT := os.Getenv("PORT"); // Server port
	CLIENT_URL := os.Getenv("CLIENT_URL"); // Client url

	// CORS config
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{CLIENT_URL},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}));

	// Server router
	routes.SetupRoutes(router);

	println(configs.Green, ">>>>> Jira Clone server run successfully at port:" + PORT, configs.Reset);

	router.Run(":" + PORT);
}
