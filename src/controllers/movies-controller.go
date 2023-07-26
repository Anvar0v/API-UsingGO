package controllers

import (
	"movies_api/dtos"
	"movies_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MoviesController struct {
	DB *gorm.DB
}

//injecting instance of gorm.DB
func NewMoviesController(DB *gorm.DB) MoviesController {
	return MoviesController{DB}
}

// Endpoint for creating a new movie
func (mc *MoviesController) CreateMovie(ctx *gin.Context) {
	var createMovieDto *dtos.CreateMovieDto

	if err := ctx.ShouldBindJSON(&createMovieDto); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newMovie := models.Movie{
		Name:        createMovieDto.Name,
		ReleaseDate: createMovieDto.ReleaseDate,
		AuthorName:  createMovieDto.AuthorName,
	}

	createMovieResult := mc.DB.Create(&newMovie)

	if createMovieResult.Error != nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": createMovieResult.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newMovie})
}

// Retrieving the list of movies
func (mc *MoviesController) GetMovies(ctx *gin.Context) {
	page := ctx.DefaultQuery("p	age", "1")
	limit := ctx.DefaultQuery("limit", "10")

	//Atoi parses string to int(similar to ParseInt)
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var movies []models.Movie
	getMoviesResult := mc.DB.Limit(intLimit).Offset(offset).Find(&movies)

	if getMoviesResult.Error != nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": getMoviesResult.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": movies})
}

// Retrieving certain movie by id
func (mc *MoviesController) GetMovieById(ctx *gin.Context) {
	movieId := ctx.Param("movieId")

	var movie models.Movie
	getByIdResult := mc.DB.Find(&movie, "id = ?", movieId)

	if getByIdResult.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "No movie with that Id exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": movie})
}

// Deleting particular movie
func (mc *MoviesController) DeleteMovie(ctx *gin.Context) {
	movieId := ctx.Param("movieId")

	deleteResult := mc.DB.Delete(&models.Movie{}, "id = ?", movieId)

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "No movie with that Id exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (mc *MoviesController) UpdateMovie(ctx *gin.Context) {
	movieId := ctx.Param("movieId")

	var movie *dtos.UpdateMovieDto

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var updatedMovie models.Movie
	getByIdResult := mc.DB.First(&updatedMovie, "id = ?", movieId)

	if getByIdResult.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "No such movie with that Id exists"})
		return
	}
}
