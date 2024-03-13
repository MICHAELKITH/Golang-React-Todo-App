package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct{
	ID int `json:"id"`
	Title  string`json:"title"`
	Done bool`json:"done"`
	Body string `json:"body"`
}

func main() {
	fmt.Println("hello")

	app := fiber.New()

	todo := []Todo{}

	app.Get("/pro", func(c *fiber.Ctx) error {
		return c.SendString("Ok pro mike ")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo :=&Todo{}

		c.BodyParser(todo)
	})

	log.Fatal(app.Listen(":4000"))

}
