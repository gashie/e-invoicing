package routes

import (
	"e-invocing/controllers/invoice"

	"github.com/gofiber/fiber/v2"
)

func InvoiceDetailsRoute(app *fiber.App) {

	//status routes
	app.Post("/api/v1/invoicedetails", invoice.CreateInvoiceDetail)
	app.Get("/api/v1/invoicedetails/:id", invoice.GetAppInvoiceDetail)
	app.Get("/api/v1/invoicedetails", invoice.AllInvoiceDetail)
}
