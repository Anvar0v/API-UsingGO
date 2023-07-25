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

func NewMoviesController(DB *gorm.DB) MoviesController {
	return MoviesController{DB}
}

func (mc *MoviesController) CreateMovie(ctx *gin.Context) {
	var createMovieDto *dtos.CreateMovieDto

	if err := ctx.ShouldBindJSON(&createMovieDto); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newMovie := models.Movie{
		Name:        createMovieDto.Name,
		ReleaseDate: createMovieDto.ReleaseDate,
		AuthorName:      createMovieDto.AuthorName,
	}

	createMovieResult := mc.DB.Create(&newMovie)

	if createMovieResult.Error != nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": createMovieResult.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newMovie})
}

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
