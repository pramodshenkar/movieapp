package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	r := gin.Default()
	r.GET("/movies", GetMovieHandler)
	// r.GET("/movie", GetMovieByIdHandler)
	r.POST("/movie", PostMovieHandler)
	// r.PUT("/movie", PutMovieHandler)
	// r.DELETE("/movie", DeleteMovieHandler)
	r.Run(":8000")
}

func PostMovieHandler(c *gin.Context) {

	client, ctx, cancel := getConnection()

	defer cancel()
	defer client.Disconnect(ctx)

	var movie Movie
	c.Bind(&movie)

	collection := client.Database(dbname).Collection("movies")
	res, err := collection.InsertOne(context.Background(), bson.D{{Key: "_id", Value: movie.Id}, {Key: "name", Value: movie.Name}, {Key: "budget", Value: movie.Budget}, {Key: "director", Value: movie.Director}, {Key: "actor", Value: movie.Actor}})
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, movie)
	fmt.Println("\nInserted ID: ", *res)

}

func GetMovieHandler(c *gin.Context) {

	client, ctx, cancel := getConnection()

	defer cancel()
	defer client.Disconnect(ctx)

	var movies []*Movie

	collection := client.Database(dbname).Collection("movies")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result Movie
		err := cur.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, &result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	c.JSON(200, movies)

}

// func GetMovieByIdHandler(c *gin.Context) {

// 	client, ctx, cancel := getConnection()

// 	defer cancel()
// 	defer client.Disconnect(ctx)

// 	var movie Movie

// 	collection := client.Database(dbname).Collection("movies")
// 	err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: movie.Id}}).Decode(&movie)
// 	if err != nil {
// 		log.Println(err)
// 		c.String(404, "Document not found")
// 	}

// 	c.JSON(200, movie)

// }

// func DeleteMovieHandler(c *gin.Context) {

// 	client, ctx, cancel := getConnection()

// 	defer cancel()
// 	defer client.Disconnect(ctx)

// 	var deleteId int

// 	collection := client.Database(dbname).Collection("movies")
// 	res, err := collection.DeleteOne(context.Background(), bson.D{{Key: "budget", Value: deleteId}})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	c.JSON(200, *res)
// }
