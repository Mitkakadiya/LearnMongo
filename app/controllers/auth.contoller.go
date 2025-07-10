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

	if err := config.UserCollection.FindOne(context.Background(), models.User{Mobile: loginInput.Mobile, CountryCode: loginInput.CountryCode}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			var newUser = models.User{
				Mobile:      loginInput.Mobile,
				CountryCode: loginInput.CountryCode,
			}

			if _, err := config.UserCollection.InsertOne(context.Background(), newUser); err != nil {
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
		if _, err := config.UserCollection.UpdateOne(context.Background(), bson.M{"mobile": loginInput.Mobile}, updateUser); err != nil {
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

func DeleteUser(c *fiber.Ctx) error {
	paramsId := c.Params("id")
	if paramsId == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Please provide valid id",
		})
	}
	oid, err := bson.ObjectIDFromHex(paramsId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Please provide valid id",
		})
	}
	if _, err := config.UserCollection.DeleteOne(context.Background(), bson.M{"_id": oid}); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Error while deleting user",
		})
	}

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "delete User Successfully",
	})
}
