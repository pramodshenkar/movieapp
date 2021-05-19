package main

type Movie struct {
	Id       int      `bson:"_id"`
	Name     string   `bson:"name"`
	Budget   int      `bson:"budget"`
	Director string   `bson:"director"`
	Actor    []string `bson:"actor"`
}
