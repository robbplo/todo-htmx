package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/robbplo/todo-htmx/components"
	"github.com/robbplo/todo-htmx/db"

	"log"
)

// go:embed db/schema.sql
var ddl string

db, err := sql.Open("sqlite3", "file::memory:?cache=shared")

func indexHandler(ctx *fiber.Ctx) error {
	todos, err := db.AllTodos()
	if err != nil {
		return err
	}

	ctx.Set("Content-Type", "text/html")
	component := components.Homepage(todos)
	return component.Render(ctx.Context(), ctx)
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

func deleteDoneHandler(ctx *fiber.Ctx) error {
	err := db.DeleteDone()
	if err != nil {
		return err
	}

	todos, err := db.AllTodos()
	if err != nil {
		return err
	}
	component := components.TodoList(todos)
	return component.Render(ctx.Context(), ctx)
}

func main() {
	app := fiber.New()
	app.Get("/", indexHandler).Name("index")
	app.Get("/todos", todosHandler).Name("todos")
	app.Post("/todos", createHandler).Name("create")
	app.Put("/todos/:id", updateHandler).Name("update")
	app.Delete("/todos/done", deleteDoneHandler).Name("deleteDone")

	log.Fatal(app.Listen(":8080"))
}
