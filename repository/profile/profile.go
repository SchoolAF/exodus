package profile

import (
	"fmt"
	"github.com/SchoolAF/exodus/repository/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserID(c *fiber.Ctx) (string, error) {
	userToken, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return "", fmt.Errorf("could not get JWT token from context")
	}

	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("could not get claims from JWT token")
	}

	userID, ok := claims["userid"].(string)
	if !ok {
		return "", fmt.Errorf("could not get userID from claims")
	}

	return userID, nil
}

func GetMyRole(userid string) (string, error) {
	filter := bson.M{"id": userid}
	account, err := db.GetOneAccountFilter(filter)
	if err != nil {
		return "", err
	}
	return account.Role, nil
}

func GetUsernameByID(userid string) (string, error) {
	filter := bson.M{"id": userid}
	account, err := db.GetOneAccountFilter(filter)
	if err != nil {
		return "", err
	}
	return account.Username, nil
}

func AuthorizeAdmin(c *fiber.Ctx) (string, error) {
	userID, err := GetUserID(c)
	if err != nil {
		return "", err
	}

	myRole, err := GetMyRole(userID)
	if err != nil {
		return "", err
	}

	if myRole != "admin" {
		return "", fiber.ErrUnauthorized
	}

	return userID, nil
}
