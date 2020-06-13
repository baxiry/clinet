package main

import (
	"fmt"

	"github.com/bashery/tweets/dbs"
	"github.com/bashery/tweets/routs"
	"github.com/bashery/tweets/tweet"
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

func main() {
	app := fiber.New()
	initdb()
	defer dbs.Conn.Close()
	routs.SetUpRouts(app)

	app.Listen(9000)
}
