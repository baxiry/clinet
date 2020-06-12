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
	//	t := &Tweet{Body: "Edit this tweet", Id: c.Params("id")}
	//	c.JSON(t)
}

func Remove(c *fiber.Ctx) {
	//	t := Tweet{Body: "Delet this tweet", Id: c.Params("id")}
	//	c.JSON(t)
}

func GetAll(c *fiber.Ctx) {
	db := dbs.Conn
	var tweets []Tweet
	db.Find(&tweets)

	c.JSON(tweets)
}
