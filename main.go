package main

import (
	"fmt"

	"github.com/elbashery/tweets/db"
	"github.com/elbashery/tweets/routes"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initdb() {
	var err error
	db.DBconn, err = gorm.Open("sqlite3", "tweets.db")
	if err != nil {
		panic("error when open database")
	}
	fmt.Println("database successfully connection")
	db.DBconn.AutoMigrate(&routes.Tweet{})
	fmt.Println("database megrated")
}
func setUpRouts(app *fiber.App) {
	app.Get("/tweet/:id", routes.GetOne)
	app.Get("/tweets", routes.GetAll)
	app.Put("/tweet/:id", routes.Edite)
	app.Post("/tweet/:id", routes.Edite)
	app.Delete("/tweet/:id", routes.Remove)
}
func main() {
	app := fiber.New()
	//app.Settings.Immutable = true
	initdb()
	defer db.DBconn.Close()
	setUpRouts(app)

	app.Listen(9000)
}
