package routes

import (
	"movies_api/controllers"

	"github.com/gin-gonic/gin"
)

type MovieRouteController struct {
	movieController controllers.MoviesController
}

func NewMovieRouteController(movieController controllers.MoviesController) MovieRouteController {
	return MovieRouteController{movieController}
}

func (mc *MovieRouteController) MoviesRoute(routeGroup *gin.RouterGroup) {
	router := routeGroup.Group("movies")
	router.POST("/", mc.movieController.CreateMovie)
	router.GET("/", mc.movieController.GetMovies)
	router.GET("/:movieId", mc.movieController.GetMovieById)
	router.PUT("/:movieId", mc.movieController.UpdateMovie)
	router.DELETE("/:movieId", mc.movieController.DeleteMovie)
}
