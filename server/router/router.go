package router

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sidikimamsetiyawan/go-project-library-app/controller"
)

// JWT Secret Key
var jwtSecret = []byte("your_jwt_secret_key")

func SetupRoutes(app *fiber.App) {
	// Public routes
	app.Post("/register", func(c *fiber.Ctx) error {
		return controller.Register(c)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return controller.Login(c)
	})

	// Update user route
	app.Put("/users/:id", func(c *fiber.Ctx) error {
		return controller.UpdateUser(c)
	})

	// Categories route
	app.Post("/categories", func(c *fiber.Ctx) error {
		return controller.AddCategories(c)
	})

	app.Put("/categories/:id", func(c *fiber.Ctx) error {
		return controller.UpdateCategories(c)
	})

	app.Get("/categories", func(c *fiber.Ctx) error {
		return controller.ListCategories(c)
	})

	app.Delete("/categories/:id", func(c *fiber.Ctx) error {
		return controller.DeleteCategories(c)
	})

	// Books route
	app.Post("/books", func(c *fiber.Ctx) error {
		return controller.AddBooks(c)
	})

	app.Put("/books/:id", func(c *fiber.Ctx) error {
		return controller.UpdateBooks(c)
	})

	app.Get("/books/list", func(c *fiber.Ctx) error {
		return controller.ListBooks(c)
	})

	app.Delete("/books/:id", func(c *fiber.Ctx) error {
		return controller.DeleteBooks(c)
	})

	// Transactions route
	app.Post("/transactions", func(c *fiber.Ctx) error {
		return controller.AddTransactions(c)
	})

	app.Put("/transactions/:id", func(c *fiber.Ctx) error {
		return controller.UpdateTransactions(c)
	})

	app.Delete("/transactions/:id", func(c *fiber.Ctx) error {
		return controller.DeleteTransactions(c)
	})

	// Protected routes
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtSecret,
	}))

	// Admin only route
	app.Get("/admin", RoleMiddleware("admin"), func(c *fiber.Ctx) error {
		return c.SendString("Welcome, Admin!")
	})

	// User route
	app.Get("/user", RoleMiddleware("user"), func(c *fiber.Ctx) error {
		return c.SendString("Welcome, User!")
	})
}

// Middleware for Role-based Access Control
func RoleMiddleware(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		role := claims["role"].(string)

		if role != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}

		return c.Next()
	}
}
