package summary

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var summaryCollection *mongo.Collection = GetCollection(DB, "summary")
var validate = validator.New()

func CreateSummary(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var meta Summary
	defer cancel()

	//validate request body
	if err := c.BodyParser(&meta); err != nil {
		return c.Status(http.StatusBadRequest).JSON(MetaResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}
	if validationErr := validate.Struct(&meta); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(MetaResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newSummary := Summary{
		Id:             primitive.NewObjectID(),
		Link:           meta.Link,
		Prompt:         meta.Prompt,
		GptSummary:     meta.GptSummary,
		HindiTranslate: meta.HindiTranslate,
		CreatedAt:      meta.CreatedAt,
	}
	result, err := summaryCollection.InsertOne(ctx, newSummary)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(MetaResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}

	return c.Status(http.StatusCreated).JSON(MetaResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})

}
