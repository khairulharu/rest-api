package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/restapi/domain"
	"github.com/khairulharu/restapi/dto"
)

type apiUser struct {
	userService domain.UserService
}

func NewApp(app *fiber.App, userService domain.UserService) {
	api := apiUser{
		userService: userService,
	}

	app.Post("/user", api.SaveUser)
	app.Get("/user", api.GetUser)
	app.Post("/user/update", api.Updating)
	app.Get("users", api.GetAllUsers)
	app.Delete("/user", api.DeleteUser)
}

func (a apiUser) SaveUser(ctx *fiber.Ctx) error {
	var userBody dto.UserData
	err := ctx.BodyParser(&userBody)
	if err != nil {
		ctx.Status(401).JSON(dto.ApiResponse{
			Message: "handler body parser error",
			Error:   err.Error(),
		})
	}
	response := a.userService.Save(ctx.Context(), userBody)
	if response.Code == "400" {
		return ctx.Status(400).JSON(response)
	}
	return ctx.Status(200).JSON(response)
}

func (a apiUser) GetUser(ctx *fiber.Ctx) error {
	var user dto.UserData
	username := ctx.Query("username")
	user.Username = username
	result := a.userService.FindUserByUsername(ctx.Context(), user)
	if result.Code == "400" {
		return ctx.Status(400).JSON(result)
	}
	return ctx.Status(200).JSON(result)
}

func (a apiUser) Updating(ctx *fiber.Ctx) error {
	var userData dto.UserData

	idString := ctx.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return ctx.Status(401).JSON(err.Error())
	}
	userData.ID = id
	err = ctx.BodyParser(&userData)
	if err != nil {
		return ctx.Status(401).JSON(err.Error())
	}

	res := a.userService.SaveUpdate(ctx.Context(), userData)

	return ctx.Status(200).JSON(res)
}

func (a apiUser) GetAllUsers(ctx *fiber.Ctx) error {
	res := a.userService.GetAll(ctx.Context())
	return ctx.JSON(res)
}

func (a apiUser) DeleteUser(ctx *fiber.Ctx) error {
	idString := ctx.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}
	res := a.userService.Deleting(ctx.Context(), int64(id))
	return ctx.Status(200).JSON(res)
}
