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

	app.Listen("localhost:8080")
}
