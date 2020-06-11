package main

import (
	"github.com/elbashery/tweets/routes"
	"github.com/gofiber/fiber"
)

func setUpRouts(app *fiber.App) {
	app.Get("/tweet/:id", routes.GetOne)
	app.Get("/tweets", routes.GetAll)
	app.Put("/tweet/:id", routes.Edite)
	app.Delete("/tweet/:id", routes.Remove)
}
func main() {
	app := fiber.New()
	//app.Settings.Immutable = true
	setUpRouts(app)

	app.Listen(9000)
}
