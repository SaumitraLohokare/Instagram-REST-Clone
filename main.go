package main

import (
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB    *mongo.Database
	users *mongo.Collection
	posts *mongo.Collection
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	host := os.Getenv("MONGO_URL")
	dbName := os.Getenv("DB_NAME")
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(host))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	err = mongoClient.Connect(ctx)
	if err != nil {
		panic("Could not connect: " + err.Error())
	}
	defer cancel()

	DB = mongoClient.Database(dbName)
	users = DB.Collection("Users")
	posts = DB.Collection("Posts")
}

func main() {
	router := gin.Default()

	router.POST("/users", createUser)
	router.GET("/users/:id", getUser)
	router.POST("/posts", createPost)
	router.GET("/posts/:id", getPost)
	router.GET("/posts/users/:id/*page", getUserPosts)

	router.Run("localhost:5000")
}
