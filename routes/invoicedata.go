package routes

import (
	"e-invocing/controllers/invoice"

	"github.com/gofiber/fiber/v2"
)

func InvoiceDataRoute(app *fiber.App) {

	//status routes
	app.Post("/api/v1/invoicedata", invoice.CreateInvoice)
	app.Get("/api/v1/invoicedata/:id", invoice.GetAppInvoiceData)
	app.Get("/api/v1/invoicedata", invoice.AllInvoiceData)
}
