package controllers

import (
	"movies_api/dtos"
	"movies_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MoviesController struct {
	DB *gorm.DB
}

func NewMoviesController(DB *gorm.DB) MoviesController {
	return MoviesController{DB}
}

func (mc *MoviesController) CreateMovie(ctx gin.Context) {
	var createMovieDto *dtos.CreateMovieDto

	if err := ctx.ShouldBindJSON(&createMovieDto); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newMovie := models.Movie{
		Name:        createMovieDto.Name,
		ReleaseDate: createMovieDto.ReleaseDate,
		AuthorId:    createMovieDto.AuthorId,
	}

	result := mc.DB.Create(&newMovie)

	if result.Error != nil {

	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newMovie})
}
