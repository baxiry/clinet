package login

import (
	//"fmt"

	//"github.com/bashery/tweets/dbs"
	"github.com/gofiber/fiber"
	//"github.com/jinzhu/gorm"
)

type User struct {
	Name string
	Pass string
}

func Submit(c *fiber.Ctx) {

	//db := dbs.Conn
	user := &User{}
	if err := c.BodyParser(user); err != nil {
		c.Status(503).JSON(err)
		//fmt.Println("here error")
		return
	}
	//db.Create(profile)
	if user.Name == "adam" && user.Pass == "1234" {
		c.JSON(user) // or c.Send("success")
		return
	}
	c.Send("wrron pass")
}

/*
func Update(c *fiber.Ctx) {

	db := dbs.Conn
//	var profile Profile

	db.First(&profile, c.Params("id"))

	// TODO check params is no zero val and handel it
	db.Model(&profile).Update("fullname", c.Params("fullname")) // &tweet.Body)

	c.Send("success") // or c.JSON("success")
}

func Remove(c *fiber.Ctx) {
	//var profile Profile
	db := dbs.Conn
	id := c.Params("id")
	//db.First(&profile, id)

	if profile.Fullname == "" {
		c.Status(500).JSON("no profile founded")
		return
	}
	db.Delete(&profile)
	c.JSON("Profile is deleted")
}

func Get(c *fiber.Ctx) {
	//c.Accepts("application/json")
	//id := c.Params("id")
	//c.send(profile)
}
*/
