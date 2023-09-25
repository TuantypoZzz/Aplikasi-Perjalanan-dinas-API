package controller

import (
	"golang-todo-app/entity"
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"golang-todo-app/service"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type EmployeeController struct {
	service.EmployeeService
	service.PerdinService
}

func NewEmployeeController(employeeService *service.EmployeeService) *EmployeeController {
	return &EmployeeController{EmployeeService: *employeeService}

}

func (controller EmployeeController) Route(api fiber.Router) {
	employee := api.Group("/employees")
	employee.Post("/", controller.Create)
	employee.Put("/:id", controller.Update)
	employee.Get("/:id", controller.FindById)
	employee.Get("/", controller.FindAll)
	employee.Delete("/:id", controller.Delete)
}

func (controller EmployeeController) Create(c *fiber.Ctx) error {
	var request model.CreateEmployeeModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	user := c.Locals("userData").(entity.User)

	file, err := c.FormFile("Photo")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Failed to get image",
		})
	}
	fileNameOnly := filepath.Base(file.Filename)

	err = UploadImage(c, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Failed to upload image",
		})
	}

	controller.EmployeeService.Create(c.Context(), request, user, fileNameOnly)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Data Has been Created",
	})
}

func (controller EmployeeController) Update(c *fiber.Ctx) error {
	var request model.UpdateEmployee
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	user := c.Locals("userData").(entity.User)

	id := c.Params("id")

	//Mendapatkan data lama berdasarkan ID
	oldPerdinReport := controller.EmployeeService.FindById(c.Context(), id)

	// Menghapus gambar lama jika ada
	if oldPerdinReport.Foto != request.Foto {
		oldImagePath := "./uploads/profile/" + oldPerdinReport.Foto
		err := os.Remove(oldImagePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
				Code:    500,
				Message: "Failed to remove old image" + err.Error(),
			})
		}
	}

	// Upload gambar
	file, err := c.FormFile("Photo")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Failed to get image",
		})
	}

	fileNameOnly := filepath.Base(file.Filename)

	err = UploadImage(c, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Failed to upload image",
		})
	}

	controller.EmployeeService.Update(c.Context(), request, id, user, fileNameOnly)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Employee has been Updated",
	})
}

func (c EmployeeController) Delete(ct *fiber.Ctx) error {
	id := ct.Params("id")

	c.EmployeeService.Delete(ct.Context(), id)
	return ct.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Employee has been Deleted",
	})
}

func (controller EmployeeController) FindById(ct *fiber.Ctx) error {
	id := ct.Params("id")

	result := controller.EmployeeService.FindById(ct.Context(), id)
	return ct.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})

}

func (c EmployeeController) FindAll(ct *fiber.Ctx) error {
	//user := ct.Locals("userData").(entity.User)
	result := c.EmployeeService.FindAll(ct.Context())
	return ct.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
func UploadImage(c *fiber.Ctx, file *multipart.FileHeader) error {

	// Mengecek ekstensi file
	allowedExtensions := []string{".png", ".jpg", ".jpeg"}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	isAllowed := false
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid image format. Only PNG and JPG/JPEG are allowed.",
		})
	}

	mimeType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(mimeType, "image/") {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Uploaded file is not an image",
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Failed to read image file",
		})
	}
	defer src.Close()

	fileName := GenerateUniqueFileName(file.Filename)
	path := "./uploads/profile/" + fileName

	// Cek apakah file sudah ada
	if _, err := os.Stat(path); err == nil {
		// File dengan nama yang sama sudah ada
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "File with the same name already exists",
		})
	}

	dst, err := os.Create(path)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Failed to create destination file",
		})
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Failed to save image file",
		})
	}
	// Jika berhasil mengunggah gambar
	return c.JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Image uploaded successfully!",
	})
}
