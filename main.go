package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetMovies(c *gin.Context) {
	var movies []Movie
	var movie Movie
	movie.Name = "3 idiots"
	movie.Budget = 10

	movies = append(movies, movie)
	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

func main() {
	r := gin.Default()
	r.GET("/movie", handleGetMovies)
	r.Run()
}
