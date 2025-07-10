package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"mongo_db/app/models"
	"mongo_db/config"
	"mongo_db/pkg/validators"
	"net/http"
	"os"
	"time"
)

func Login(c *fiber.Ctx) error {

	var loginInput models.LoginInput

	if err := c.BodyParser(&loginInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Please provide valid mobile number and phone code",
		})
	}

	if err := validators.ValidateStruct(&loginInput); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"error":  err,
		})
	}

	var user models.User

	collection := config.DB.Database("Testing").Collection("Users")
	if err := collection.FindOne(context.Background(), models.User{Mobile: loginInput.Mobile, CountryCode: loginInput.CountryCode}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			var newUser = models.User{
				Mobile:      loginInput.Mobile,
				CountryCode: loginInput.CountryCode,
			}

			if _, err := collection.InsertOne(context.Background(), newUser); err != nil {
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{
					"status":  http.StatusBadRequest,
					"message": "Error while inserting user",
				})
			}
		} else {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"status":  http.StatusBadRequest,
				"message": "Error while inserting user",
			})
		}
	}

	if os.Getenv("ENV") != "production" {
		updateUser := bson.M{
			"$set": bson.M{
				"otp":        "000000",
				"expired_at": time.Time.UnixMilli(time.Now()) + (5 * 60 * 1000),
			},
		}
		if _, err := collection.UpdateOne(context.Background(), bson.M{"mobile": loginInput.Mobile}, updateUser); err != nil {
			fmt.Println(err.Error())
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"status":  http.StatusBadRequest,
				"message": "Error while updating user",
			})
		}
	} else {

	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "OTP sent successfully",
	})
}
