package user

import (
	"fmt"

	"github.com/bashery/tweets/dbs"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Password string `json:"password"` // ??
	Avatar   string `json:"avatar"`   // link to place stor avatar
}

func Get(c *fiber.Ctx) {
	//c.Accepts("application/json")
	id := c.Params("id")
	db := dbs.Conn
	var profile Profile
	db.Find(&profile, id)

	if profile.Fullname == "" {
		c.Send("Profile is no found")
		return
	}

	c.JSON(profile)
}

func New(c *fiber.Ctx) {

	db := dbs.Conn
	profile := &Profile{}
	if err := c.BodyParser(profile); err != nil {
		c.Status(503).JSON(err)
		fmt.Println("here error")
		return
	}
	db.Create(profile)
	c.JSON(profile) // or c.Send("success")
}

func Update(c *fiber.Ctx) {

	db := dbs.Conn
	var profile Profile

	db.First(&profile, c.Params("id"))

	/* // I will use params instead BodyParser just for semple and performece
	if err := c.BodyParser(&tweet); err != nil {
		c.Status(503).Send(err)
		fmt.Println(err)
		return
	}*/

	// TODO check params is no zero val and handel it
	db.Model(&profile).Update("fullname", c.Params("fullname")) // &tweet.Body)

	c.Send("success") // or c.JSON("success")
}

func Remove(c *fiber.Ctx) {
	var profile Profile
	db := dbs.Conn
	id := c.Params("id")
	db.First(&profile, id)

	if profile.Fullname == "" {
		c.Status(500).JSON("no profile founded")
		return
	}
	db.Delete(&profile)
	c.JSON("Profile is deleted")
}

/*
func GetAll(c *fiber.Ctx) {
	db := dbs.Conn
	var tweets []Tweet
	db.Find(&tweets)
	c.JSON(tweets)
}
*/
