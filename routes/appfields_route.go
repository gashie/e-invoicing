package routes

import (
	"e-invocing/controllers/apps"

	"github.com/gofiber/fiber/v2"
)

func AppFieldsRoute(app *fiber.App) {

	//status routes
	app.Post("/api/v1/appfield", apps.CreateAppField)
	app.Get("/api/v1/appfield/:id", apps.GetAppFields)
	app.Get("/api/v1/appfield/:id/:catId", apps.GetAppCatFields)

	app.Get("/api/v1/appfield", apps.AllAppFields)
	app.Put("/api/v1/appfield/:id", apps.UpdateAppField)
	app.Delete("/api/v1/appfield/:id", apps.DeleteAppField)
}
