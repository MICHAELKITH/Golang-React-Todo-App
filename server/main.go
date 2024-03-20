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

	todos:= []Todo{}

	app.Get("/pro", func(c *fiber.Ctx) error {
		return c.SendString("Ok pro mike ")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo :=&Todo{}

		 if err := c.BodyParser(todo); err != nil{
			return nil
		 }
		 todo.ID =len(todos) +1;

		 todos = append(todos, *todo)

		 return c.JSON(todos)
	});


	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id");

		if err != nil{
			return c.Status(401).SendString("Invalid id")

		}

		for i, t := range todos{
			if t.ID == id {
				todos[i].Done = true

				break
			}
		}

		return c.JSON(todos)
	})

	app.Get("api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))

}
