package main

import (
	"fmt"
	"github.com/aminkbi/gobok/cmd/internal/data"
	"github.com/aminkbi/gobok/cmd/internal/util"
	"github.com/aminkbi/gobok/cmd/internal/validation"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func BookGroup(app *fiber.App) {

	bookGroup := app.Group("/v1/book")

	bookGroup.Get("/", GetBooks)
	bookGroup.Post("/", AddBook)
	bookGroup.Get("/:id", GetBookById)
	bookGroup.Put("/:id", UpdateBook)
	bookGroup.Delete("/:id", DeleteBook)

}

func GetBooks(ctx *fiber.Ctx) error {

	coll := GetDBCollection("books")

	// find all books
	books := make([]data.Book, 0)
	cursor, err := coll.Find(ctx.Context(), bson.M{})

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	for cursor.Next(ctx.Context()) {
		book := data.Book{}
		err := cursor.Decode(&book)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		books = append(books, book)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"data": books})
}

func AddBook(ctx *fiber.Ctx) error {

	coll := GetDBCollection("books")

	b := new(data.Book)
	if err := ctx.BodyParser(b); err != nil {
		fmt.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	validator := validation.New()

	data.ValidateBook(validator, b)

	if !validator.Valid() {
		return ctx.Status(400).JSON(validator.Errors)
	}

	one, err := coll.InsertOne(ctx.Context(), b)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			key, err := util.ExtractDupKey(err.Error())
			if err == nil {
				return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
					"error": fmt.Errorf("duplicate key: %s", key).Error(),
				})
			}
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "duplicate key",
			})
		}
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": one,
	})
}

func GetBookById(ctx *fiber.Ctx) error {

	coll := GetDBCollection("books")

	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var book data.Book

	err = coll.FindOne(ctx.Context(), bson.M{"_id": objectId}).Decode(&book)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"data": book})

}

func UpdateBook(ctx *fiber.Ctx) error {

	coll := GetDBCollection("books")

	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	b := new(data.Book)
	if err := ctx.BodyParser(b); err != nil {
		fmt.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	validator := validation.New()

	data.ValidateBook(validator, b)

	if !validator.Valid() {
		return ctx.Status(400).JSON(validator.Errors)
	}

	_, err = coll.UpdateOne(ctx.Context(), bson.M{"_id": objectId}, bson.M{"$set": b})
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error":   "Failed to update book",
			"message": err.Error(),
		})
	}

	// return the book
	return ctx.Status(200).JSON(fiber.Map{
		"result": "successfully updated document",
	})

}

func DeleteBook(ctx *fiber.Ctx) error {
	coll := GetDBCollection("books")

	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	_, err = coll.DeleteOne(ctx.Context(), bson.M{"_id": objectId})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"result": "successfully deleted document",
	})

}
