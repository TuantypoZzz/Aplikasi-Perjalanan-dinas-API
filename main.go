package main

import (
	"golang-todo-app/configuration"
	"golang-todo-app/controller"
	"golang-todo-app/exception"
	"golang-todo-app/middleware"
	"golang-todo-app/repository"
	"golang-todo-app/service"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database := configuration.NewDatabase()

	//repository
	todoRepository := repository.NewTodoRepositoryImpl(database)
	userRepository := repository.NewUserRepositoryImpl(database)
	roleRepository := repository.NewRoleRepositoryImpl(database)
	DepartmentRepository := repository.NewDepartmentRepositoryImpl(database)
	HandleRepository := repository.NewHandleRepositoryImpl(database)
	employeeRepository := repository.NewEmployeeRepositoryImpl(database)
	lumpsumRepository := repository.NewLumpsumRepositoryImpl(database)
	activityRepository := repository.NewActivityRepositoryImpl(database)
	perdinRepository := repository.NewPerdinRepositoryImpl(database)
	telaahRepository := repository.NewTelaahRepositoryImpl(database)
	sptRepository := repository.NewSptRepositoryImpl(database)
	sppdRepository := repository.NewSppdRepositoryImpl(database)
	transportRepository := repository.NewTransportRepositoryImpl(database)
	hotelRepository := repository.NewAccomodationRepositoryImpl(database)
	EmplyRepository := repository.NewEmployeePerdinRepositoryImpl(database)
	LaporanRepository := repository.NewPerdinLaporanRepositoryImpl(database)
	documentRepository := repository.NewDocumentRepositoryImpl(database)

	//service
	todoService := service.NewTodoServiceImpl(&todoRepository)
	authService := service.NewAuthServiceImpl(&userRepository, &roleRepository)
	DepartmentService := service.NewDepartmentRepositoryImpl(&DepartmentRepository)
	HandleService := service.NewHandleRepositoryImpl(&HandleRepository)
	employeeService := service.NewEmployeeRepositoryImpl(&employeeRepository)
	lumpsumService := service.NewLumpsumRepositoryImpl(&lumpsumRepository)
	activityService := service.NewActivityServiceImpl(&activityRepository)
	perdinService := service.NewPerdinServiceImpl(&perdinRepository, &telaahRepository, &sptRepository, &sppdRepository, &transportRepository, &hotelRepository, &EmplyRepository)
	laporanService := service.NewPerdinLaporanServiceImpl(&LaporanRepository, &perdinRepository)
	documentService := service.NewDocumentImpl(&documentRepository, &EmplyRepository, &perdinRepository)

	//controller
	todoController := controller.NewTodoController(&todoService)
	authController := controller.NewAuthController(&authService)
	DepartmentController := controller.NewDepartmentController(&DepartmentService)
	pangkatDepartmentController := controller.NewPangkatgGolController(&HandleService)
	employeeController := controller.NewEmployeeController(&employeeService)
	lumpsumController := controller.NewLumpsumController(&lumpsumService)
	activityController := controller.NewActivityController(&activityService)
	perdinController := controller.NewPerdinController(&perdinService, &laporanService, &documentService)

	//setup fiber
	app := fiber.New(configuration.NewFiberConfig())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New(configuration.NewLoggerConfig()))

	//routing
	api := app.Group("/api/v1")
	authController.Route(api)
	api.Use(middleware.JwtCustomStrategy(userRepository))
	todoController.Route(api)
	DepartmentController.Route(api)
	pangkatDepartmentController.Route(api)
	employeeController.Route(api)
	lumpsumController.Route(api)
	activityController.Route(api)
	perdinController.Route(api)

	err := app.Listen(os.Getenv("SERVER_PORT"))
	exception.PanicLogging(err)

	configuration.NewLogger().Info("log info")
}

/*
ref:
- https://github.com/RizkiMufrizal/gofiber-clean-architecture
- https://project-awesome.org/gofiber/awesome-fiber
*/
