package controller

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sidikimamsetiyawan/go-project-library-app/database"
	"github.com/sidikimamsetiyawan/go-project-library-app/model"
)

// Add a transactions into database
func AddTransactions(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Transaction Added Successfully",
	}

	var transactions []model.Transactions

	// Parse the incoming JSON body into the slice of Category structs
	if err := c.BodyParser(&transactions); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	i := 0
	for _, transaction := range transactions {
		transactions[i].CreatedBy = transaction.CreatedBy
		transactions[i].CreatedDate = time.Now()
		transactions[i].ModifiedBy = transaction.CreatedBy
		transactions[i].ModifiedDate = time.Now()
		i++
	}

	result := database.DBConn.Create(transactions)

	if result.Error != nil {
		log.Println("error in saving data.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong"
	}

	context["data"] = transactions

	c.Status(201)
	return c.JSON(context)

}

func UpdateTransactions(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Updated Transaction Successfully",
	}

	id := c.Params("id")

	var record []model.Transactions

	database.DBConn.First(&record, id)

	if record[0].TransactionID == 0 {
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

	context["data"] = record

	c.Status(200)
	return c.JSON(context)
}

func DeleteTransactions(c *fiber.Ctx) error {

	c.Status(400)

	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record []model.Transactions

	database.DBConn.First(&record, id)

	if record[0].TransactionID == 0 {
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
	context["msg"] = "Transactions deleted successfully."

	c.Status(200)
	return c.JSON(context)
}
