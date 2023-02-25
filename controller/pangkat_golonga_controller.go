package controller

import (
	"golang-todo-app/entity"
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"golang-todo-app/service"

	"github.com/gofiber/fiber/v2"
)

type HandleController struct {
	service.HandleService
}

func NewPangkatgGolController(pangkatgolservice *service.HandleService) *HandleController {
	return &HandleController{HandleService: *pangkatgolservice}
}

func (c HandleController) Route(api fiber.Router) {
	pangkatgol := api.Group("/handle")
	pangkatgol.Post("/", c.Create)
	pangkatgol.Put("/:id", c.Update)
	pangkatgol.Delete("/:id", c.Delete)
	pangkatgol.Get("/:id", c.FindById)
	pangkatgol.Get("/", c.FindAll)
}

func (c HandleController) Create(ctx *fiber.Ctx) error {
	var request model.CreateHandle
	err := ctx.BodyParser(&request)
	exception.PanicLogging(err)

	user := ctx.Locals("userData").(entity.User)

	response := c.HandleService.Create(ctx.Context(), request, user)
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Data has been created successfully",
		Data:    response,
	})
}

func (c HandleController) Update(ctx *fiber.Ctx) error {
	var request model.UpdateHandle
	err := ctx.BodyParser(&request)
	exception.PanicLogging(err)

	user := ctx.Locals("userData").(entity.User)
	id := ctx.Params("id")

	response := c.HandleService.Update(ctx.Context(), request, id, user)
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Pangkat Golongan has been updated successfully",
		Data:    response,
	})
}

func (c HandleController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	c.HandleService.Delete(ctx.Context(), id)
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Department has been deleted successfully",
	})
}

func (c HandleController) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	result := c.HandleService.FindById(ctx.Context(), id)
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Successfully found",
		Data:    result,
	})
}

func (c HandleController) FindAll(ctx *fiber.Ctx) error {
	result := c.HandleService.FindAll(ctx.Context())
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Successfully Found",
		Data:    result,
	})
}
