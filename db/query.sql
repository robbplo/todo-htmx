-- name: GetTodo :one
SELECT * FROM todos
WHERE id = ? LIMIT 1;

-- name: GetTodos :many
SELECT * FROM todos
ORDER BY id;

-- name: CreateTodo :one
INSERT INTO todos (id, task, done)
VALUES (NULL, ?, ?)
RETURNING *;

-- name: UpdateTodo :exec
UPDATE todos SET task = ?, done = ?
WHERE id = ?;

-- name: DeleteAllDone :exec
DELETE FROM todos
WHERE done = 1;
