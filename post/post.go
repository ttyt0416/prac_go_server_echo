package post

import (
	"net/http"
	"time"

	"github.com/ttyt0416/learngo/configs"
	"github.com/ttyt0416/learngo/models"
	"github.com/ttyt0416/learngo/responses"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var (
	postCollection *mongo.Collection = configs.GetCollection(configs.DB, "posts")
	validate                         = validator.New()
)

func CreatePost(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var post models.Post
	defer cancel()

	//validate the request body
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&post); validationErr != nil {
		return c.JSON(http.StatusBadRequest, validationErr)
	}

	newPost := models.Post{
		Id:          primitive.NewObjectID(),
		Writer:      post.Writer,
		Title:       post.Title,
		Description: post.Description,
		CreatedAt:   primitive.DateTime(time.Now().UTC().UnixNano() / 1e6),
		UpdatedAt:   primitive.DateTime(time.Now().UTC().UnixNano() / 1e6),
	}

	result, err := postCollection.InsertOne(ctx, newPost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}

func GetPost(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	postId := c.Param("postId")
	var post models.Post
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)

	err := postCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&post)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": post}})
}

func UpdatePost(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	postId := c.Param("postId")
	var post models.Post
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)

	//validate the request body
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&post); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"title": post.Title, "Description": post.Description, "updatedAt": primitive.DateTime(time.Now().UTC().UnixNano() / 1e6)}

	result, err := postCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//get updated user details
	var updatedUser models.Post
	if result.MatchedCount == 1 {
		err := postCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": updatedUser}})
}

func DeletePost(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	postId := c.Param("postId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)

	result, err := postCollection.DeleteOne(ctx, bson.M{"id": objId})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: &echo.Map{"data": "User with specified ID not found!"}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": "User successfully deleted!"}})
}

func GetAllPosts(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var posts []models.Post
	defer cancel()

	results, err := postCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singlePost models.Post
		if err = results.Decode(&singlePost); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}

		posts = append(posts, singlePost)
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": posts}})
}
