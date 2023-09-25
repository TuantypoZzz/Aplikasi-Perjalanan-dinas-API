package controller

import (
	"fmt"
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"golang-todo-app/service"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type PerdinController struct {
	service.PerdinService
	service.PerdinLaporanService
	service.DocumentService
}

func NewPerdinController(perdinService *service.PerdinService, perdinLaporanService *service.PerdinLaporanService, documentPerdinService *service.DocumentService) *PerdinController {
	return &PerdinController{PerdinService: *perdinService, PerdinLaporanService: *perdinLaporanService, DocumentService: *documentPerdinService}
}

func (controller PerdinController) Route(api fiber.Router) {
	dinas := api.Group("/perjalanan-dinas/")
	dinas.Post("/telaah", controller.CreateTelaah)
	dinas.Get("/telaah", controller.FindAllTelaah)
	dinas.Post("/spt", controller.CreateSpt)
	dinas.Get("/spt", controller.FindAllSpt)
	dinas.Post("/sppd", controller.CreateSppd)
	dinas.Get("/sppd", controller.FindAllSppd)
	dinas.Post("/rincian-perdin", controller.CreateRincianPerdin)
	dinas.Post("/laporan-perdin", controller.PerdinReport)
	dinas.Post("/dokumen-perdin", controller.DocumentPerdin)
	dinas.Get("/rekap-perdin", controller.FindAllPerdin)
}

func (controller PerdinController) CreateTelaah(ctx *fiber.Ctx) error {
	var request model.CreateTelaahModel
	errs := ctx.BodyParser(&request)
	exception.PanicLogging(errs)

	result := controller.PerdinService.CreateTelaah(ctx.Context(), request)
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Successfully Created",
		Data:    result,
	})
}

func (controller PerdinController) FindAllTelaah(ctx *fiber.Ctx) error {
	result := controller.PerdinService.FindallTelaah(ctx.Context())
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (controller PerdinController) CreateSpt(ctx *fiber.Ctx) error {
	var request model.CreateSptModel
	errs := ctx.BodyParser(&request)
	exception.PanicLogging(errs)

	controller.PerdinService.CreateSpt(ctx.Context(), request)
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Successfully Created",
	})
}

func (controller PerdinController) FindAllSpt(ctx *fiber.Ctx) error {
	result := controller.PerdinService.FindAllSpt(ctx.Context())
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (controller PerdinController) CreateSppd(ctx *fiber.Ctx) error {
	var request model.CreateSppdModel
	errs := ctx.BodyParser(&request)
	exception.PanicLogging(errs)

	controller.PerdinService.CreateSppd(ctx.Context(), request)
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Successfully Created",
	})
}

func (controller PerdinController) FindAllSppd(ctx *fiber.Ctx) error {
	result := controller.PerdinService.FindAllSppd(ctx.Context())
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (controller PerdinController) CreateRincianPerdin(ctx *fiber.Ctx) error {
	var request model.CreateDetailPerdin
	errs := ctx.BodyParser(&request)
	exception.PanicLogging(errs)

	controller.PerdinService.CreateRincianPegawaiPerdin(ctx.Context(), request)
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Successfully Created",
	})
}

// func (controller PerdinController) PerdinReport(c *fiber.Ctx) error {
// 	var request model.CreatePerdinReportModel
// 	err := c.BodyParser(&request)
// 	exception.PanicLogging(err)

// 	controller.PerdinService.CreatePerdinReport(c.Context(), request)
// 	return c.JSON(model.GeneralResponse{
// 		Code:    200,
// 		Message: "Successfully Created",
// 	})
// }

func (controller PerdinController) FindAllPerdin(ctx *fiber.Ctx) error {
	result := controller.PerdinLaporanService.FindAllPerdin(ctx.Context())
	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (controller PerdinController) PerdinReport(c *fiber.Ctx) error {
	var request model.CreatePerdinReportModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	// Upload gambar
	file, err := c.FormFile("Foto")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Failed to get image",
		})
	}

	userID := controller.PerdinService.GetPerdinReportByID(c.Context(), request.PerdinId) // Replace with your actual user ID
	uploadPath := "./uploads/dokumentasi/" + userID.PerdinId
	err = os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Failed to create upload directory",
		})
	}

	// // Mendapatkan data lama berdasarkan ID
	// oldPerdinReport := controller.PerdinService.GetPerdinReportByID(c.Context(), request.PerdinId)
	// // Menghapus gambar lama jika ada
	// if oldPerdinReport.Foto != request.Foto {
	// 	oldImagePath := "./uploads/" + oldPerdinReport.Foto
	// 	err := os.Remove(oldImagePath)
	// 	if err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
	// 			Code:    500,
	// 			Message: "Failed to remove old image",
	// 		})
	// 	}
	// }

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
	path := filepath.Join(uploadPath, fileName)

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

	fileNameOnly := filepath.Base(path)

	controller.PerdinService.CreatePerdinReport(c.Context(), request, fileNameOnly)

	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Successfully Created",
		Data:    fileNameOnly,
	})
}

func (controller PerdinController) DocumentPerdin(ctx *fiber.Ctx) error {
	var request model.DokProofPerdin
	errs := ctx.BodyParser(&request)
	exception.PanicLogging(errs)

	// Parse the multipart form
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Failed to get image",
		})
	}
	// Get the files from the form
	files := form.File["Files"]

	// Check the number of uploaded files
	if len(files) != 5 {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Exactly 5 files must be uploaded",
		})
	}

	// Create a directory with ID-based path
	userID := controller.PerdinService.GetEmployeePerdinByName(ctx.Context(), request.EmplyPerdinId) // Replace with your actual user ID
	uploadPath := "./uploads/" + userID.Nama
	err = os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Failed to create upload directory",
		})
	}

	validFilesCount := 0  // Track the number of valid files
	failedToSave := false // Track if any condition fails during file processing

	// Process each file
	for _, file := range files {
		// Open the file
		src, err := file.Open()
		if err != nil {
			failedToSave = true // Set failedToSave to true if opening file fails
			break               // Break the loop
		}
		defer src.Close()

		// Check file size
		fileSize := file.Size
		const (
			minFileSize = 100 * 1024  // 100 KB
			maxFileSize = 1024 * 1024 // 1 MB
		)
		if fileSize < minFileSize || fileSize > maxFileSize {
			src.Close()
			failedToSave = true // Set failedToSave to true if file size is invalid
			break               // Break the loop
		}

		fileName := GenerateUniqueFileName(file.Filename)
		path := filepath.Join(uploadPath, fileName)

		// Create a destination file
		dst, err := os.Create(path)
		if err != nil {
			src.Close()
			failedToSave = true // Set failedToSave to true if creating destination file fails
			break               // Break the loop
		}
		defer dst.Close()

		// Copy the file content to the destination
		_, err = io.Copy(dst, src)
		if err != nil {
			src.Close()
			dst.Close()
			failedToSave = true // Set failedToSave to true if copying file fails
			break               // Break the loop
		}
		src.Close()
		dst.Close()

		fileNameOnly := filepath.Base(path)
		controller.DocumentService.CreateDocument(ctx.Context(), request, fileNameOnly)
		validFilesCount++
	}

	// If any condition fails during the loop, delete all uploaded files and return an error response
	if validFilesCount != 5 || failedToSave {
		DeleteUploadedFiles(uploadPath) // Delete all uploaded files

		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to upload files",
		})
	}

	return ctx.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Successfully Created",
	})
}

func GenerateUniqueFileName(filename string) string {
	// Get the file extension
	extension := filepath.Ext(filename)

	// Generate a timestamp-based unique string
	// timestamp := time.Now().Format("20060102150405")

	// Remove the original extension from the filename
	fileNameWithoutExt := strings.TrimSuffix(filename, extension)

	// Construct the new unique file name by appending the timestamp and the original extension
	newFileName := fileNameWithoutExt + extension

	return newFileName
}

func DeleteUploadedFiles(uploadPath string) {
	files, err := os.ReadDir(uploadPath)
	if err != nil {
		// Handle the error if any
		fmt.Println("Failed to read the directory:", err)
		return
	}

	for _, file := range files {
		err := os.Remove(filepath.Join(uploadPath, file.Name()))
		if err != nil {
			fmt.Println("Failed to delete file:", err)
		}
	}
}
