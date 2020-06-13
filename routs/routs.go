package routs

import (
	"github.com/bashery/tweets/tweet"
	"github.com/gofiber/fiber"
)

func SetUpRouts(app *fiber.App) {
	app.Get("/tweet/:id", tweet.GetOne)
	app.Get("/tweets", tweet.GetAll)
	app.Post("/tweet/new", tweet.New)
	app.Put("/tweet/:id/:body", tweet.Update)
	app.Delete("/tweet/:id", tweet.Remove)
}
