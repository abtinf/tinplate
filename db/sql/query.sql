-- name: GetFoo :one
SELECT * FROM foo WHERE id = ?;

-- name: ListFoos :many
SELECT * FROM foo;

-- name: CreateFoo :one
INSERT INTO foo (foo, bar) VALUES (?, ?) RETURNING *;
