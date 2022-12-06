package apps

import (
	"e-invocing/database"
	"e-invocing/models"
	"github.com/gofiber/fiber/v2"
)

func CreateAppField(c *fiber.Ctx) error {
	var model models.InvoiceAppFields

	if err := c.BodyParser(&model); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Fields cannot be empty",
		})
	}

	if model.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Name  cannot be empty",
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
func AllAppFields(c *fiber.Ctx) error {
	var model []models.InvoiceAppFields
	database.DB.Preload("Category").Preload("InvoiceApps").Find(&model)
	return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    model,
		"Message": "Record Found",
	})
}

func GetAppFields(c *fiber.Ctx) error {
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
func GetAppCatFields(c *fiber.Ctx) error {
	id := c.Params("id")
	catId := c.Params("catId")
	var model []models.InvoiceAppFields

	database.DB.Find(&model, "invoice_app_id = ? AND category_id = ?", id,catId)
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
func UpdateAppField(c *fiber.Ctx) error {
	var body models.InvoiceAppFields

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Fields cannot be empty",
		})
	}

	var model models.InvoiceAppFields

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

func DeleteAppField(c *fiber.Ctx) error {
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
