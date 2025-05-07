package middleware

import (
	"encoding/json"
	"golang-todo-app/entity"
	"golang-todo-app/model"
	"golang-todo-app/repository"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func JwtCustomStrategy(cacheRepository repository.CacheRepository) func(*fiber.Ctx) error {
	jwtSecret := os.Getenv("JWT_SECRET_TOKEN")
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Locals("user").(*jwt.Token)
			claims := token.Claims.(jwt.MapClaims)

			cachedData, err := cacheRepository.Get(claims["username"].(string))
			if err != nil || cachedData == nil {
				return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
					Code:    401,
					Message: "Unauthorized",
					Data:    "User not found in cache",
				})
			}

			// First unmarshal into CachedUser
			var cachedUser model.CachedUser
			err = json.Unmarshal(cachedData, &cachedUser)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
					Code:    500,
					Message: "Error parsing user data",
					Data:    err.Error(),
				})
			}

			// Convert CachedUser to entity.User
			user := entity.User{
				Username:    cachedUser.Username,
				Roles:       make([]entity.Role, len(cachedUser.Roles)),
				Permissions: make([]entity.Permission, len(cachedUser.Permissions)),
			}

			// Convert role strings to Role objects
			for i, roleName := range cachedUser.Roles {
				user.Roles[i] = entity.Role{
					Name: roleName,
				}
			}

			// Convert permission strings to Permission objects
			for i, permName := range cachedUser.Permissions {
				user.Permissions[i] = entity.Permission{
					Name: permName,
				}
			}

			c.Locals("userData", user)
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				// panic(exception.ValidationError{
				// 	Message: "Missing or malformed JWT",
				//  })
				return c.
					Status(fiber.StatusBadRequest).
					JSON(model.GeneralResponse{
						Code:    400,
						Message: "Bad Request",
						Data:    "Missing or malformed JWT",
					})
			} else {
				return c.
					Status(fiber.StatusUnauthorized).
					JSON(model.GeneralResponse{
						Code:    401,
						Message: "Unauthorized",
						Data:    "Invalid or expired JWT",
					})
			}
		},
	})
}
