package handler

import (
	"app/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultMovies struct {
	db     []internal.Movie
	lastID int
}

func NewDefaultMovies() *DefaultMovies {
	return &DefaultMovies{
		db:     make([]internal.Movie, 0),
		lastID: 0,
	}
}

type MovieJSON struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Director  string `json:"director"`
	InTheater bool   `json:"in_theater"`
}

type MovieRequest struct {
	Title     *string `json:"title"`
	Year      *int    `json:"year"`
	Director  *string `json:"director"`
	InTheater *bool   `json:"in_theater"`
}

func (d *DefaultMovies) CreateMovie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//err := json.NewDecoder(ctx.Request.Body).Decode(&body)
		var body MovieRequest
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
			return
		}

		if body.Title == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "title is required"})
			return
		}
		if body.Year == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "year must be greater than 0"})
			return
		}
		if body.Director == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "director is required"})
			return
		}
		if body.InTheater == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "in_theater is required"})
			return
		}

		mv := internal.Movie{
			Title:     *body.Title,
			Year:      *body.Year,
			Director:  *body.Director,
			InTheater: *body.InTheater,
		}
		d.lastID++
		mv.ID = d.lastID
		d.db = append(d.db, mv)

		ctx.JSON(http.StatusCreated, MovieJSON{
			ID:        mv.ID,
			Title:     mv.Title,
			Year:      mv.Year,
			Director:  mv.Director,
			InTheater: mv.InTheater,
		})
	}
}
