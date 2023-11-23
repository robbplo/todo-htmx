package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/robbplo/todo-htmx/components"
	"github.com/robbplo/todo-htmx/db"
	"log"
)

func indexHandler(ctx *fiber.Ctx) error {
	todos, err := db.AllTodos()
	if err != nil {
		return err
	}

	ctx.Set("Content-Type", "text/html")
	component := components.Homepage(todos)
	component.Render(ctx.Context(), ctx)
	return nil
}

func todosHandler(ctx *fiber.Ctx) error {
	todos, err := db.AllTodos()
	if err != nil {
		return err
	}

	component := components.TodoList(todos)
	return component.Render(ctx.Context(), ctx)
}

func createHandler(ctx *fiber.Ctx) error {
	todo := db.Todo{Task: ctx.FormValue("task"), Done: false}
	err := todo.Create()
	if err != nil {
		return err
	}

	return ctx.Redirect("/todos")
}

func updateHandler(ctx *fiber.Ctx) error {
	todo, err := db.Find(ctx.Params("id"))
	if err != nil {
		return err
	}

	todo.Done = ctx.FormValue("done") == "on"

	err = todo.Update()
	if err != nil {
		return err
	}

	component := components.Todo(todo)
	return component.Render(ctx.Context(), ctx)
}

func main() {
	app := fiber.New()
	app.Get("/", indexHandler)
	app.Get("/todos", todosHandler)
	app.Post("/todos", createHandler)
	app.Put("/todos/:id", updateHandler)

	println("Listening on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
