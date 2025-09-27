-- name: GetCustomers :many
SELECT * FROM customers
ORDER BY id;

-- name: GetCustomerByID :one
SELECT * FROM customers
WHERE id = $1;

-- name: CreateCustomer :one
INSERT INTO customers (name, email)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateCustomer :one
UPDATE customers
SET name = $2, email = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE id = $1;