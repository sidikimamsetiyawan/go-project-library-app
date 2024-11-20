package controller

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sidikimamsetiyawan/go-project-library-app/database"
	"github.com/sidikimamsetiyawan/go-project-library-app/model"
)

/*

	AddCategories
	UpdateCategories
	ListCategories
	DeleteCategories
*/

func AddCategories(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a category data into database",
	}

	var categories []model.Categories

	// Parse the incoming JSON body into the slice of Category structs
	if err := c.BodyParser(&categories); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	// Insert categories into the database
	result := database.DBConn.Create(categories)

	if result.Error != nil {
		log.Println("error in saving data.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong"
	}

	context["msg"] = "Record is saved successfully."
	context["data"] = categories

	c.Status(201)
	return c.JSON(context)
}

func UpdateCategories(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update Blog",
	}

	id := c.Params("id")

	var record []model.Categories

	database.DBConn.First(&record, id)

	if record[0].CategoryID == 0 {
		log.Println("Record not found.")
		context[""] = ""
		context["msg"] = "Record not found."
		c.Status(400)
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
	}

	result := database.DBConn.Save(record)

	if result.Error != nil {
		log.Println("Error saving data.")
	}

	context["msg"] = "Record update successfuly."
	context["data"] = record

	c.Status(200)
	return c.JSON(context)
}

func ListCategories(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Category List",
	}

	db := database.DBConn

	var records []model.Categories

	db.Find(&records)

	fmt.Println("Category List Data : ")
	fmt.Println(records)

	context["category_records"] = records

	c.Status(200)
	return c.JSON(context)

}

func DeleteCategories(c *fiber.Ctx) error {

	c.Status(400)

	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record []model.Categories

	database.DBConn.First(&record, id)

	if record[0].CategoryID == 0 {
		log.Println("Record not found.")
		context["msg"] = "Record not found."

		c.Status(400)
		return c.JSON(context)
	}

	result := database.DBConn.Delete(record)

	if result.Error != nil {
		context["msg"] = "Something went wrong."
		return c.JSON(context)
	}

	context["statusText"] = "Ok."
	context["msg"] = "Record deleted successfully."

	c.Status(200)
	return c.JSON(context)

}
