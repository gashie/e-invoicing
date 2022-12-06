package routes

import (
	"e-invocing/controllers/apps"

	"github.com/gofiber/fiber/v2"
)

func AppRoute(app *fiber.App) {

	//status routes
	app.Post("/api/v1/app", apps.CreateApp)
	app.Get("/api/v1/app/:id", apps.GetApp)
	app.Get("/api/v1/app", apps.AllApps)
	app.Put("/api/v1/app/:id", apps.UpdateApp)
	app.Delete("/api/v1/app/:id", apps.DeleteApp)
}
