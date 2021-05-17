package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Budget   int      `json:"budget"`
	Director string   `json:"director"`
	Actor    []string `json:"actor"`
}

func main() {
	r := gin.Default()
	r.POST("/movie", handle)
	r.Run(":8000")
}

func handle(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}
	defer func() {
		err = client.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}()

	var movie Movie
	c.Bind(&movie)

	response := client.Database("dempmovieapp").Collection("movies")
	res, err := response.InsertOne(context.Background(), bson.D{{Key: "_id", Value: movie.Id}, {Key: "name", Value: movie.Name}, {Key: "budget", Value: movie.Budget}, {Key: "director", Value: movie.Director}, {Key: "actor", Value: movie.Actor}})
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, movie)
	fmt.Println("\nInserted ID: ", *res)

}
