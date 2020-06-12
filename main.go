package main

import (
	"fmt"

	"github.com/elbashery/tweets/dbs"
	"github.com/elbashery/tweets/tweet"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initdb() {
	fmt.Println("hello")
	var err error
	dbs.Conn, err = gorm.Open("sqlite3", "tweets.db")
	if err != nil {
		panic("error when open database")
	}
	fmt.Println("database successfully connection")
	dbs.Conn.AutoMigrate(&tweet.Tweet{})
	fmt.Println("database megrated")
}
func setUpRouts(app *fiber.App) {
	app.Get("/tweet/:id", tweet.GetOne)
	app.Get("/tweets", tweet.GetAll)
	app.Put("/tweet/:id", tweet.Edite)
	app.Post("/tweet/:id", tweet.Edite)
	app.Delete("/tweet/:id", tweet.Remove)
}

func main() {
	app := fiber.New()
	//app.Settings.Immutable = true
	initdb()
	defer dbs.Conn.Close()
	setUpRouts(app)

	app.Listen(9000)
}
