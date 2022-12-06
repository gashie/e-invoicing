package apps

import (
	"e-invocing/database"
	"e-invocing/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateApp(c *fiber.Ctx) error {
	var model models.InvoiceApps

	if err := c.BodyParser(&model); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Fields cannot be empty",
		})
	}

	if model.AppName == "" {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Appname  cannot be empty",
		})
	}

	if model.Description == "" {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Description  cannot be empty",
		})
	}
	result := database.DB.Create(&model)
	if result.RowsAffected != 1 {

		return c.Status(200).JSON(fiber.Map{
			"Status":  0,
			"Message": "Record Not Saved",
		})

	}
	return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    model,
		"Message": "Record Saved",
	})

}
func AllApps(c *fiber.Ctx) error {
	var model []models.InvoiceApps
	database.DB.Find(&model)
	return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    model,
		"Message": "Record Found",
	})
}

func GetApp(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)
	var model models.InvoiceApps
	fmt.Println(model.ID)

	database.DB.Find(&model, "id = ?", id)
	if model.AppName == "" {
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
func UpdateApp(c *fiber.Ctx) error {
	var body models.InvoiceApps

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Fields cannot be empty",
		})
	}

	var model models.InvoiceApps

	Id := c.Params("id")
	// CHECK AVAILABLE CATEGORY
	err := database.DB.First(&model, "id = ?", Id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Message": "Record not found",
			"Status":  0,
		})
	}
	// UPDATE USER DATA
	if body.AppName != "" {
		model.AppName = body.AppName
	}

	errUpdate := database.DB.Model(&model).Updates(body).Error //save new changes

	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Failed to update record",
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Record Updated",
		"Data":    model,
	})
}

func DeleteApp(c *fiber.Ctx) error {
	Id := c.Params("id")
	var model models.InvoiceApps

	//	CHECK AVAILABLE USER
	err := database.DB.Debug().First(&model, "id=?", Id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Message": "Record not found",
			"Status":  0,
		})
	}

	errDelete := database.DB.Debug().Delete(&model).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Failed to update record",
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Record Removed",
		"Status":  1,
	})
}
