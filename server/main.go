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

		 if err := c.BodyParser(todo); err != nil{
			return nil
		 }
		 todo.ID =len(todos)+1;

		 todos = append(todos, *todo)

		 return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))

}
