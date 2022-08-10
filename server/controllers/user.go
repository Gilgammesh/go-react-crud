package controllers

import (
	"context"

	"github.com/Gilgammesh/go-react-crud/database"
	"github.com/Gilgammesh/go-react-crud/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers(c *fiber.Ctx) error {
	db := database.Connection()
	coll := db.Collection("users")

	var users []models.User

	results, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	for results.Next(context.TODO()) {
		var user models.User
		results.Decode(&user)
		users = append(users, user)
	}

	return c.JSON(fiber.Map{
		"data":      users,
		"registers": len(users),
	})
}

func GetUser(c *fiber.Ctx) error {
	db := database.Connection()
	coll := db.Collection("users")

	id, errId := primitive.ObjectIDFromHex(c.Params("id"))
	if errId != nil {
		panic(errId)
	}

	var result bson.M
	err := coll.FindOne(
		context.TODO(),
		bson.D{{Key: "_id", Value: id}},
	).Decode(&result)
	if err != nil {
		panic(err)
	}

	return c.JSON(&fiber.Map{
		"data": result,
	})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.Connection()
	coll := db.Collection("users")

	user := new(models.User)
	c.BodyParser(user)

	result, err := coll.InsertOne(
		context.TODO(),
		bson.D{
			{Key: "firstname", Value: user.Firstname},
			{Key: "lastname", Value: user.Lastname},
			{Key: "email", Value: user.Email},
			{Key: "phone", Value: user.Phone},
		},
	)
	if err != nil {
		panic(err)
	}

	return c.JSON(&fiber.Map{
		"data": result,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.Connection()
	coll := db.Collection("users")

	id, errId := primitive.ObjectIDFromHex(c.Params("id"))
	if errId != nil {
		panic(errId)
	}

	user := new(models.User)
	c.BodyParser(user)

	result, err := coll.UpdateOne(
		context.TODO(),
		bson.D{{Key: "_id", Value: id}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "firstname", Value: user.Firstname},
			{Key: "lastname", Value: user.Lastname},
			{Key: "email", Value: user.Email},
			{Key: "phone", Value: user.Phone},
		}}},
	)
	if err != nil {
		panic(err)
	}

	return c.JSON(&fiber.Map{
		"data": result,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.Connection()
	coll := db.Collection("users")

	id, errId := primitive.ObjectIDFromHex(c.Params("id"))
	if errId != nil {
		panic(errId)
	}

	result, err := coll.DeleteOne(
		context.TODO(),
		bson.D{{Key: "_id", Value: id}},
	)
	if err != nil {
		panic(err)
	}

	return c.JSON(&fiber.Map{
		"data": result,
	})
}
