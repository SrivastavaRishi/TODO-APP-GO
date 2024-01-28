package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var items = []string{"Do Coding"}

func printItems() {
	for _, val := range items {
		fmt.Printf("Value is %s", val)
	}
}

type Item struct {
	Name string `json:"name"`
}

func main() {
	app := fiber.New()
	printItems()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(items)
	})
	app.Post("/", func(c *fiber.Ctx) error {
		newItem := new(Item)
		if err := c.BodyParser(newItem); err != nil {
			fmt.Println("error = ", err)
			c.SendStatus(417)
		}
		items = append(items, newItem.Name)
		return c.JSON(items)
	})
	app.Delete("/deleteItem/:id", func(c *fiber.Ctx) error {
		index := c.Params("id")
		x, _ := strconv.ParseInt(index, 10, 64)
		items = append(items[:x], items[x+1:]...)
		return c.JSON(items)
	})

	app.Listen(":3000")
}
