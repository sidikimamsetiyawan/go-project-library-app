package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sidikimamsetiyawan/go-project-library-app/database"
	"github.com/sidikimamsetiyawan/go-project-library-app/model"
)

func AddBooks(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add books data successfully.",
	}

	var books []model.Books

	// Parse the incoming JSON body into the slice of Category structs
	if err := c.BodyParser(&books); err != nil {
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	// Print the parsed data to the console

	// Insert books into the database
	result := database.DBConn.Create(books)

	if result.Error != nil {
		context["statusText"] = ""
		context["msg"] = "Something went wrong"
	}

	context["data"] = books

	c.Status(201)
	return c.JSON(context)
}

func UpdateBooks(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update books data successfully.",
	}

	id := c.Params("id")

	var books []model.Books

	database.DBConn.First(&books, id)

	if books[0].BookID == 0 {
		context[""] = ""
		context["msg"] = "books not found."
		c.Status(400)
		return c.JSON(context)
	}

	if err := c.BodyParser(&books); err != nil {
		log.Println("Error in parsing request.")
	}

	result := database.DBConn.Save(books)

	if result.Error != nil {
		log.Println("Error saving data.")
	}

	context["data"] = books

	c.Status(200)
	return c.JSON(context)
}

func DeleteBooks(c *fiber.Ctx) error {
	c.Status(400)

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Delete books data successfully.",
	}

	id := c.Params("id")

	var books []model.Books

	database.DBConn.First(&books, id)

	if books[0].CategoryID == 0 {
		log.Println("books not found.")
		context["msg"] = "Record not found."

		c.Status(400)
		return c.JSON(context)
	}

	result := database.DBConn.Delete(books)

	if result.Error != nil {
		context["msg"] = "Something went wrong."
		return c.JSON(context)
	}

	c.Status(200)
	return c.JSON(context)
}

func ListBooks(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Books List",
	}

	db := database.DBConn

	var records []model.ListBooks

	db.Table("books").
		Select(`books.book_id as book_id
		, books.title as title
		, categories.category_name as category_name
		, books.author as author
		, books.published_year as published_year
		, books.total_copies as total_copies
		, books.available_copies as available_copies`).
		Joins("left join categories on categories.category_id = books.category_id").
		// Where("orders.amount > ?", 100).
		Find(&records)

	context["product_records"] = records

	c.Status(200)
	return c.JSON(context)
}
