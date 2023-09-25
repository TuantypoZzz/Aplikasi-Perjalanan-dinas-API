package controller

import (
	"golang-todo-app/entity"
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"golang-todo-app/service"

	"github.com/gofiber/fiber/v2"
)

type activityController struct {
	service.ActivityService
}

func NewActivityController(employeeService *service.ActivityService) *activityController {
	return &activityController{ActivityService: *employeeService}

}

func (controller activityController) Route(api fiber.Router) {
	activities := api.Group("/kegiatan")
	activities.Post("/", controller.Create)
	activities.Put("/:id", controller.Update)
	activities.Get("/:id", controller.FindById)
	activities.Get("/", controller.FindAll)
	activities.Delete("/:id", controller.Delete)
}

func (controller activityController) Create(c *fiber.Ctx) error {
	var request model.CreateActivityModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	user := c.Locals("userData").(entity.User)

	controller.ActivityService.Create(c.Context(), request, user)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Data Has been Created",
	})
}

func (controller activityController) Update(ctx *fiber.Ctx) error {
	var request model.UpdateActivityModel
	errs := ctx.BodyParser(&request)
	exception.PanicLogging(errs)

	id := ctx.Params("id")
	controller.ActivityService.Update(ctx.Context(), request, id)
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Activity has been Updated",
	})
}

func (c activityController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	c.ActivityService.Delete(ctx.Context(), id)
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Activity has been Deleted",
	})
}

func (controller activityController) FindById(ct *fiber.Ctx) error {
	id := ct.Params("id")

	result := controller.ActivityService.FindById(ct.Context(), id)
	return ct.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (c activityController) FindAll(ct *fiber.Ctx) error {
	//user := ct.Locals("userData").(entity.User)
	result := c.ActivityService.FindAll(ct.Context())
	return ct.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
