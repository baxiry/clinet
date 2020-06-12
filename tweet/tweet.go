package tweet

import (
	//	"fmt"

	"github.com/elbashery/tweets/dbs"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Tweet struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

func GetOne(c *fiber.Ctx) {
	//c.Accepts("application/json")
	id := c.Params("id")
	db := dbs.Conn
	var tweet Tweet
	db.Find(&tweet, id)

	c.JSON(tweet)
}
func Edite(c *fiber.Ctx) {

	t := &Tweet{}
	t.Title = "first"         //c.Params("title")
	t.Body = "Hello evry one" // c.Params("tweet")}
	db := dbs.Conn
	db.Create(t)
	c.JSON(t)
}

func Remove(c *fiber.Ctx) {
	var tweet Tweet
	db := dbs.Conn
	id := c.Params("id")
	db.First(&tweet, id)

	if tweet.Title == "" {
		c.Status(500).JSON("no tweet find with given ID ")
		return
	}
	db.Delete(&tweet)
	c.JSON("tweet is deleted")
}

func GetAll(c *fiber.Ctx) {
	db := dbs.Conn
	var tweets []Tweet
	db.Find(&tweets)

	c.JSON(tweets)
}
