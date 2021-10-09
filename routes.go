package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createUser(c *gin.Context) {
	id, _ := c.GetPostForm("Id")
	name, _ := c.GetPostForm("Name")
	email, _ := c.GetPostForm("Email")
	password, _ := c.GetPostForm("Password")
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashPassword := hex.EncodeToString(hasher.Sum(nil))
	user := User{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: hashPassword,
	}
	// _ is result
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := users.InsertOne(ctx, user)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func getUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id := c.Param("id")
	var result User
	err := users.FindOne(ctx, bson.M{"Id": id}).Decode(&result)
	if err != nil {
		panic(err.Error())
	}
	// User json
	c.IndentedJSON(http.StatusOK, result)
}

func createPost(c *gin.Context) {
	id, _ := c.GetPostForm("Id")
	userId, _ := c.GetPostForm("UserId")
	caption, _ := c.GetPostForm("Caption")
	imageURL, _ := c.GetPostForm("ImageURL")
	timestamp, _ := c.GetPostForm("Timestamp")
	post := Post{
		Id:        id,
		UserId:    userId,
		Caption:   caption,
		ImageURL:  imageURL,
		Timestamp: timestamp,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := posts.InsertOne(ctx, post)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func getPost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id := c.Param("id")
	var result Post
	err := posts.FindOne(ctx, bson.M{"Id": id}).Decode(&result)
	if err != nil {
		panic(err.Error())
	}
	// Post json
	c.IndentedJSON(http.StatusOK, result)
}

func getUserPosts(c *gin.Context) {
	id := c.Param("id")
	page, _ := strconv.ParseInt(c.Param("page")[1:], 10, 32)
	paginationLimit := 10

	if page < 1 {
		c.String(http.StatusInternalServerError, "Invalid Page Number")
		return
	}

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * int64(paginationLimit))
	findOptions.SetLimit(int64(paginationLimit))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := posts.Find(ctx, bson.M{"UserId": id}, findOptions)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	cursorCtx, cursorCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cursorCancel()
	var entries []Post
	if err = cursor.All(cursorCtx, &entries); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// List Posts json
	c.IndentedJSON(http.StatusOK, entries)
}
