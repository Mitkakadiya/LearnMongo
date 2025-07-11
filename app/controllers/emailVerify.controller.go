package controllers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
	"mongo_db/app/models"
	"mongo_db/config"
	"mongo_db/pkg/utills"
	"mongo_db/pkg/validators"
)

func EmailVerify(c *fiber.Ctx) error {
	var inputEmail = models.InputEmail{}
	if err := c.BodyParser(&inputEmail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "Enter valid email",
		})
	}

	if err := validators.ValidateStruct(&inputEmail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "Enter valid email",
		})
	}

	var token, err = models.GenerateToken(16)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err,
			"message": "Token generate error",
		})
	}

	var emailVerification = models.EmailVerification{}
	emailVerification.Email = inputEmail.Email
	emailVerification.EmailVerificationToken = token

	if _, err := config.EmailCollection.InsertOne(context.Background(), &emailVerification); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "Enter valid email",
		})
	}

	htmlBody, err := utills.RenderVerificationEmail("http://127.0.0.1:8081/auth/email/verify?token=" + token)
	if err != nil {
		fmt.Println("Template rendering error:", err)
	}

	if err := utills.SendVerificationEmail(inputEmail.Email, htmlBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "Faild to send verification email",
		})
	}

	return c.JSON(fiber.Map{
		"success": fiber.StatusOK,
		"message": "Verify Email",
	})
}

func TokenVerification(c *fiber.Ctx) error {
	token := c.Query("token")

	if token == "" {
		return c.JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "Please provide valid token",
		})
	}

	res, err := config.EmailCollection.UpdateOne(context.Background(), bson.M{"email_verification_token": token}, bson.M{
		"$set": bson.M{"is_verified": true, "email_verification_token": ""},
	})

	if err != nil {
		return c.JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "Please provide valid token",
		})
	}

	if res.MatchedCount == 0 {
		return c.JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "Invalid or missing token",
		})
	}

	return c.JSON(fiber.Map{
		"success": fiber.StatusOK,
		"message": "Verify Token",
	})
}
