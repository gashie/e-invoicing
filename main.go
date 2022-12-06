package main

import (
	"fmt"
	"log"
	"os"
	"e-invocing/database"
	"e-invocing/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
)

func main() {
	fastergoding.Run() // hot reload
	database.Connect()
	app := fiber.New()
	routes.Setup(app)

	errListen := app.Listen(":3003")
	if errListen != nil {
		log.Println("Fail to listen go fiber server")
		fmt.Println(errListen)
		os.Exit(1)
	}
}
