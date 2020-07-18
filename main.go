package main

import (
	"fmt"

	"github.com/bashery/tweets/dbs"
	"github.com/bashery/tweets/routs"
	"github.com/bashery/tweets/tweet"
	"github.com/bashery/tweets/user"
	//"github.com/bashery/tweets/user"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initdb() {
	var err error
	dbs.Conn, err = gorm.Open("sqlite3", "../.tweets.db")
	if err != nil {
		fmt.Println(err)
		panic("error when open database")
	}
	//fmt.Println("database successfully connection")
	dbs.Conn.AutoMigrate(&tweet.Tweet{})
	dbs.Conn.AutoMigrate(&user.Profile{})
	//fmt.Println("database megrated")
}

func main() {
	auth()
	app := fiber.New()
	app.Static("/", "./assets")
	initdb()
	defer dbs.Conn.Close()
	routs.SetUpRouts(app)

	app.Listen(8000)
}
