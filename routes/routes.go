package routes

import (
	"fmt"

	"github.com/gofiber/fiber"
)

type tweet struct {
	Data string
	Id   string
}

func GetOne(c *fiber.Ctx) {
	//t := &tweet{data: "Get one tweet", id: c.Params("id")}
	c.Accepts("application/json")

	t := &tweet{
		Data: "a nice Tweet",
		Id:   c.Params("id"),
	}

	fmt.Println(*t)

	c.JSON(*t)
}
func GetAll(c *fiber.Ctx) {
	c.JSON("this is all tweets")
}
func Edite(c *fiber.Ctx) {
	t := tweet{Data: "Edit this tweet", Id: c.Params("id")}

	c.JSON(t)
}

func Remove(c *fiber.Ctx) {
	c.JSON("remove tweet")
}
