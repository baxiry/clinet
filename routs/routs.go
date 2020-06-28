package routs

import (
	//"github.com/bashery/tweets/join"
	"github.com/bashery/tweets/login"
	"github.com/bashery/tweets/tweet"
	"github.com/bashery/tweets/user"
	"github.com/gofiber/fiber"
)

func SetUpRouts(app *fiber.App) {
	// tweet opiration
	app.Get("/tweet/:id", tweet.GetOne)
	app.Get("/tweets", tweet.GetAll)
	app.Post("/tweet/new/:title/:body", tweet.New)
	app.Put("/tweet/:id/:body", tweet.Update)
	app.Delete("/tweet/:id", tweet.Remove)

	// Profile opiration
	app.Get("/user/:id", user.Get)
	app.Post("/user/new", user.New)
	app.Put("/user/:id", user.Update)
	app.Get("/user/:id", user.Remove)

	// login
	//app.Get("/login", login.Get)
	app.Post("/login", login.Submit)

	// register
	//app.Get("/join", join.Get)
	//app.Post("/join", join.Submit)

	// logout ??
}
