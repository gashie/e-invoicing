package invoice

import (
	"e-invocing/database"
	"e-invocing/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateInvoiceDetail(c *fiber.Ctx) error {
	var model models.InvoiceDetails

	if err := c.BodyParser(&model); err != nil {
		fmt.Println(err)

		fmt.Println(c.BodyParser(&model))
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Fields cannot be empty",
		})
	}


	// result := database.DB.Create(&model)
	if err:= database.DB.Create(&model).Error;  err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":0,
			"Message":err.Error(),
		}) 
	   }
	// if result.RowsAffected != 1 {

	// 	return c.Status(200).JSON(fiber.Map{
	// 		"Status":  0,
	// 		"Message": "Record Not Saved",
	// 	})

	// }
	return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    model,
		"Message": "Record Saved",
	})

}
func AllInvoiceDetail(c *fiber.Ctx) error {
	var model []models.InvoiceAppFields
	database.DB.Preload("Category").Preload("InvoiceApps").Find(&model)
	return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    model,
		"Message": "Record Found",
	})
}

func GetAppInvoiceDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var model []models.InvoiceAppFields

	database.DB.Find(&model, "invoice_app_id = ?", id)
	if len(model) == 0 {
		return c.Status(200).JSON(fiber.Map{
			"Status":  0,
			"Message": "No Record Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    model,
		"Message": "Record Found",
	})
}
