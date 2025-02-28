package main

import (
	"github.com/Canudo319/go-file-system/src/lib"
	"github.com/Canudo319/go-file-system/src/templates/pages"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

//go:generate templ generate
//go:generate npm run compile

func main() {
	app := fiber.New()

	app.Use("/public", static.New("./public"))
	app.Get("/", func(c fiber.Ctx) error {
		return lib.Render(c, pages.Hello("Cainã"))
	})
	app.Get("/jsonRFID", func(c fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.SendString(`
{"DateTime":"2025-02-27T14:30:00Z","Task":["Picking","Pack","Store"],"Orders":[{"OrderID":12345,"CustomerID":67890,"Products":[{"ProductID":54321,"Color":"Blue","Size":"M","WMSAddress":[{"Aisle":"A12","Level":2,"Bin":"B15","Quantity":10}]},{"ProductID":98765,"Color":"Red","Size":"L","WMSAddress":[{"Aisle":"A14","Level":1,"Bin":"B22","Quantity":5}]}]},{"OrderID":67890,"CustomerID":54321,"Products":[{"ProductID":11111,"Color":"Black","Size":"S","WMSAddress":[{"Aisle":"A10","Level":3,"Bin":"B05","Quantity":20}]}]}]}
		`)
	})

	app.Listen("localhost:8080")
}
