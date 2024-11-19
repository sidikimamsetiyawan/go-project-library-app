package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidikimamsetiyawan/go-project-library-app/controller"
)

// Setup routing information
func SetupRoutes(app *fiber.App) {
	// List =>
	// Add => Post
	// Update => Put
	// Delete => Delete

	app.Get("/", controller.CategoryList)
}
