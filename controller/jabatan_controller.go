package controller

import (
	"golang-todo-app/entity"
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"golang-todo-app/service"

	"github.com/gofiber/fiber/v2"
)

type DepartmentController struct {
	service.DepartmentService
}

func NewDepartmentController(DepartmentService *service.DepartmentService) *DepartmentController {
	return &DepartmentController{DepartmentService: *DepartmentService}
}

func (controller DepartmentController) Route(api fiber.Router) {
	Department := api.Group("/department")
	Department.Post("/", controller.Create)
	Department.Put("/:id", controller.Update)
	Department.Delete("/:id", controller.Delete)
	Department.Get("/:id", controller.FindById)
	Department.Get("/", controller.FindAll)
}

func (controller DepartmentController) Create(c *fiber.Ctx) error {
	var request model.CreateDepartment
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	user := c.Locals("userData").(entity.User)

	response := controller.DepartmentService.Create(c.Context(), request, user)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Data Has been Created",
		Data:    response,
	})
}

func (c DepartmentController) Update(ct *fiber.Ctx) error {
	var request model.UpdateDepartment
	err := ct.BodyParser(&request)
	exception.PanicLogging(err)

	user := ct.Locals("userData").(entity.User)

	id := ct.Params("id")
	response := c.DepartmentService.Update(ct.Context(), request, id, user)
	return ct.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Department has been Updated",
		Data:    response,
	})
}

func (c DepartmentController) Delete(ct *fiber.Ctx) error {
	id := ct.Params("id")

	c.DepartmentService.Delete(ct.Context(), id)
	return ct.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Department has been Deleted",
	})
}

func (c DepartmentController) FindById(ct *fiber.Ctx) error {
	id := ct.Params("id")

	result := c.DepartmentService.FindById(ct.Context(), id)
	return ct.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})

}

func (c DepartmentController) FindAll(ct *fiber.Ctx) error {
	//user := ct.Locals("userData").(entity.User)
	result := c.DepartmentService.FindAll(ct.Context())
	return ct.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
