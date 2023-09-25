package controller

import (
	"golang-todo-app/entity"
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"golang-todo-app/service"

	"github.com/gofiber/fiber/v2"
)

type LumpsumController struct {
	service.LumpsumService
}

func NewLumpsumController(lumpsumService *service.LumpsumService) *LumpsumController {
	return &LumpsumController{LumpsumService: *lumpsumService}
}

func (controller LumpsumController) Route(api fiber.Router) {
	lumpsum := api.Group("/lumpsum")
	lumpsum.Post("/", controller.Create)
	lumpsum.Put("/:id", controller.Update)
	lumpsum.Delete("/:id", controller.Delete)
	lumpsum.Get("/", controller.FindAll)
	lumpsum.Get("/:id", controller.FindById)
}

func (controller LumpsumController) Create(ctx *fiber.Ctx) error {
	var request model.CreateLumpsum
	err := ctx.BodyParser(&request)
	exception.PanicLogging(err)

	user := ctx.Locals("userData").(entity.User)

	controller.LumpsumService.Create(ctx.Context(), request, user)
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Data has been created",
	})
}

func (controller LumpsumController) Update(ctx *fiber.Ctx) error {
	var request model.UpdateLumpsum
	err := ctx.BodyParser(&request)
	exception.PanicLogging(err)

	user := ctx.Locals("userData").(entity.User)

	id := ctx.Params("id")
	controller.LumpsumService.Update(ctx.Context(), request, id, user)
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Lumpsum has been updated",
	})
}

func (controller LumpsumController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	controller.LumpsumService.Delete(ctx.Context(), id)
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Lumpsum has been deleted",
	})
}

func (controller LumpsumController) FindAll(ctx *fiber.Ctx) error {
	result := controller.LumpsumService.FindAll(ctx.Context())
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (controller LumpsumController) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	result := controller.LumpsumService.FindById(ctx.Context(), id)
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
