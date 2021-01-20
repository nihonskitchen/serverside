package controllers

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	repositories "github.com/nihonskitchen/serverside/repositories"
)

// GetAllUsers is called by GET /api/users
func GetAllUsers(ctx *fiber.Ctx) error {
	users := repositories.FindAllUsers()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Got All Users",
		"data": fiber.Map{
			"users": users,
		},
	})
}

// GetUserByID is called by GET /api/users/:id
func GetUserByID(ctx *fiber.Ctx) error {
	//ドキュメントIDを渡す必要がある
	docID := ctx.Params("id")
	if docID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "No Document ID",
		})
	}
	user := repositories.FindUserByID(ctx, docID)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Got user by ID",
		"data": fiber.Map{
			"users": user,
		},
	})
}

// CreateUser is called by POST /api/users
func CreateUser(ctx *fiber.Ctx) error {
	params := new(struct {
		UID  string `json:"uid"`
		Name string `json:"name"`
	})

	err := ctx.BodyParser(&params)

	// if error
	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	if len(params.UID) == 0 || len(params.Name) == 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "ID or name not specified.",
		})
	}

	targetUser := repositories.User{UID: params.UID, Name: params.Name}
	createdUser, err := repositories.SaveUser(targetUser)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":     "Created New User",
		"createdUser": createdUser,
	})
}

// PutUser is called by PUT /api/users/:id
func PutUser(ctx *fiber.Ctx) error {
	//ドキュメントIDを渡す必要がある
	docID := ctx.Params("id")
	if docID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "No Document ID",
		})
	}

	params := new(struct {
		UID  string
		Name string
	})

	err := ctx.BodyParser(&params)

	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	//TODO 現状は0値が入っていてても更新されてしまう
	if params.UID == "" || params.Name == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Sorry now you need to pass all data to update (;^;)",
		})
	}

	targetUser := repositories.User{UID: params.UID, Name: params.Name}
	updatedUser, err := repositories.UpdateUser(docID, targetUser)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":     "Updated User",
		"updatedUser": updatedUser,
	})
}

// DeleteUser is called by DELETE /api/users/:id
func DeleteUser(ctx *fiber.Ctx) error {
	//ドキュメントIDを渡す必要がある
	docID := ctx.Params("id")
	if docID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "No Document ID",
		})
	}
	err := repositories.DeleteUserByID(docID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Deleted User",
	})
}
