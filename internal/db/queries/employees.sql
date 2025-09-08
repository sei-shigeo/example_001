-- name: GetEmployees :many
SELECT * FROM employees
ORDER BY id;

-- name: GetEmployeeByID :one
SELECT * FROM employees
WHERE id = $1;

-- name: CreateEmployee :one
INSERT INTO employees (name, email)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateEmployee :one
UPDATE employees
SET name = $2, email = $3
WHERE id = $1
RETURNING *;

-- name: DeleteEmployee :exec
DELETE FROM employees
WHERE id = $1;