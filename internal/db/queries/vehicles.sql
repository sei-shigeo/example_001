-- name: GetVehicles :many
SELECT * FROM vehicles
ORDER BY id;

-- name: GetVehicleByID :one
SELECT * FROM vehicles
WHERE id = $1;

-- name: CreateVehicle :one
INSERT INTO vehicles (number)
VALUES ($1)
RETURNING *; 

-- name: UpdateVehicle :one
UPDATE vehicles
SET number = $2
WHERE id = $1
RETURNING *;

-- name: DeleteVehicle :exec
DELETE FROM vehicles
WHERE id = $1;