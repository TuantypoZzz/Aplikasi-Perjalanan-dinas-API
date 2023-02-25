package controller

import (
	"golang-todo-app/entity"
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"golang-todo-app/service"

	"github.com/gofiber/fiber/v2"
)

type EmployeeController struct {
	service.EmployeeService
}

func NewEmployeeController(employeeService *service.EmployeeService) *EmployeeController {
	return &EmployeeController{EmployeeService: *employeeService}

}

func (controller EmployeeController) Route(api fiber.Router) {
	employee := api.Group("/employees")
	employee.Post("/", controller.Create)
}

func (controller EmployeeController) Create(c *fiber.Ctx) error {
	var request model.CreateEmployeeModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	user := c.Locals("userData").(entity.User)

	response := controller.EmployeeService.Create(c.Context(), request, user)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Data Has been Created",
		Data:    response,
	})
}
