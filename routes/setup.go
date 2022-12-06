package routes

import (
	"e-invocing/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/v1/health", controllers.Health)

	//app routes
	AppRoute(app)

	//category routes
	CategoryRoute(app)

	//appfields routes
	AppFieldsRoute(app)

	//invoice data routes
	InvoiceDataRoute(app)

	//invoice detail routes
	InvoiceDetailsRoute(app)
}
