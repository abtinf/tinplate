-- name: TryMigrationLock :one
SELECT pg_try_advisory_lock(1);

-- name: AwaitMigrationLock :exec
SELECT pg_advisory_lock(1);

-- name: ReleaseMigrationLock :one
SELECT pg_advisory_unlock(1);

-- name: InsertMigrationIfTableEmpty :exec
INSERT INTO migration (name, query) SELECT $1, $2 WHERE NOT EXISTS (SELECT 1 FROM migration);

-- name: InsertMigration :one
INSERT INTO migration (name, query) VALUES ($1, $2) RETURNING *;

-- name: MostRecentMigration :one
SELECT * FROM migration ORDER BY id DESC LIMIT 1;

-- name: ListAllMigrations :many
SELECT * FROM migration ORDER BY id;



-- name: GetFoo :one
SELECT
	*
FROM
	foo
WHERE
	id = $1;

-- name: ListFoos :many
SELECT
	*
FROM
	foo;

-- name: CreateFoo :one
INSERT INTO foo (foo, bar)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateFoo :exec
UPDATE foo
SET
	foo = $1,
	bar = $2
WHERE
	id = $3;

-- name: DeleteFoo :exec
DELETE FROM foo
WHERE
	id = $1;
