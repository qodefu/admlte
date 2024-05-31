
-- name: GetUserClient :one
SELECT u.id, c.name FROM clients c 
JOIN users u
on c.id = u.id 
WHERE u.id = $1 LIMIT 1;

-- name: GetClient :one
SELECT * FROM clients 
WHERE id = $1 LIMIT 1;

-- name: ListClients :many
SELECT * FROM  clients 
ORDER BY $1;

-- name: CreateClient :one
INSERT INTO  clients(
  name
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateClient :exec
UPDATE clients 
  set name = $2
WHERE id = $1;

-- name: DeleteClient :exec
DELETE FROM clients 
WHERE id = $1;

