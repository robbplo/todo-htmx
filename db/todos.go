package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Todo struct {
	Id   int
	Task string
	Done bool
}

var db = initDB()

func initDB() *sql.DB {
	// Create a new sqlite database in directory `db`
	db, err := sql.Open("sqlite3", "./db/storage.db")
	if err != nil {
		log.Fatal(err)
	}

	// Yes, i write migrations in my code. Try and stop me.
	db.Exec("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY, task VARCHAR(255), done BOOLEAN)")
	return db
}

func (t *Todo) Create() error {
	_, err := db.Exec("INSERT INTO todos (id, task, done) VALUES (?, ?, ?)", nil, t.Task, t.Done)
	return err
}

func (t *Todo) Update() error {
	_, err := db.Exec("UPDATE todos SET (task, done) = (?, ?) WHERE id = ?", t.Task, t.Done, t.Id)
	return err
}

func Find(id string) (Todo, error) {
	var todo Todo
	err := db.QueryRow("SELECT id, task, done FROM todos WHERE id = ?", id).Scan(&todo.Id, &todo.Task, &todo.Done)
	return todo, err
}

func AllTodos() ([]Todo, error) {
	rows, err := db.Query("SELECT id, task, done FROM todos")
	if err != nil {
		return nil, err
	}

	var todos []Todo
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.Id, &todo.Task, &todo.Done)
		todos = append(todos, todo)
	}

	return todos, nil
}

func DeleteDone() error {
	_, err := db.Exec("DELETE FROM todos WHERE done = ?", true)
	return err
}
