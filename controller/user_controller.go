package controller

import (
	"net/http"

	"github.com/HiteshKumarMeghwar/L-M-S/domain"
	"github.com/HiteshKumarMeghwar/L-M-S/model"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (n *UserController) CreateUser(ctx *fiber.Ctx) error {
	var userRequest model.User
	var response model.Response

	// handle the request
	if err := ctx.BodyParser(&userRequest); err != nil {
		response = model.Response{StatusCode: http.StatusBadRequest, Message: err.Error()}
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	// check if the request was bad/null
	if userRequest.Author == "" || userRequest.Name == "" || userRequest.Description == "" {
		response = model.Response{StatusCode: http.StatusBadRequest, Message: "Request cannot be empty"}
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	// save into database
	if err := n.userUsecase.CreateUser(userRequest); err != nil {
		response = model.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response = model.Response{StatusCode: http.StatusOK, Message: "Inserted Successfully"}
	return ctx.Status(http.StatusOK).JSON(response)
}
