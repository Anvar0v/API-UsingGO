package main

import (
	"fmt"
	"movies_api/config"
	"movies_api/controllers"
	"movies_api/models"
	"movies_api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server               *gin.Engine
	MoviesController     controllers.MoviesController
	MovieRouteController routes.MovieRouteController
)

func init() {
	config.ConnectToDB()

	MoviesController = controllers.NewMoviesController(config.DB)
	MovieRouteController = routes.NewMovieRouteController(MoviesController)
	server = gin.Default()
}

func main() {
	config.DB.AutoMigrate(&models.Movie{})
	fmt.Println("? Migration completed")

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8080"}

	//implementing cors configuration
	server.Use(cors.New(corsConfig))
	router := server.Group("/api")
	MovieRouteController.MoviesRoute(router)
	//running application in port 8080
	server.Run(":8080")
}
