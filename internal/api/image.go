package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/restapi/domain"
	"github.com/khairulharu/restapi/dto"
)

type apiImage struct {
	imageService domain.ImageService
}

func NewImage(app *fiber.App, imageService domain.ImageService) {
	h := apiImage{
		imageService: imageService,
	}

	app.Get("image/", func(ctx *fiber.Ctx) error {
		res := h.imageService.GetAll(ctx.Context())
		return ctx.JSON(res)
	})

	app.Post("image/insert", h.SaveImage)
	app.Delete("image/delete", h.Delete)
}

func (a apiImage) SaveImage(ctx *fiber.Ctx) error {
	var imagePars dto.ImageData

	err := ctx.BodyParser(&imagePars)
	if err != nil {
		return ctx.Status(401).JSON(err.Error())
	}

	res := a.imageService.SaveData(ctx.Context(), imagePars)

	return ctx.JSON(res)
}

func (a apiImage) Delete(ctx *fiber.Ctx) error {
	queries := ctx.Queries()

	idString := queries["id"]
	id, _ := strconv.Atoi(idString)

	userId := queries["user"]
	user, _ := strconv.Atoi(userId)

	res := a.imageService.Save(ctx.Context(), int64(id), int64(user))

	return ctx.JSON(res)
}
