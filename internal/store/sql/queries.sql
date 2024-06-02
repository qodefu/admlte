
-- name: GetAppointment :one
SELECT * FROM appointments 
WHERE id = $1 LIMIT 1;

-- name: ListAppt :many
SELECT * FROM  appointments 
ORDER BY $1;

-- name: CreateAppt :one
INSERT INTO appointments(
  client_id, appt_time, Status, Note
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateAppt :exec
UPDATE appointments 
  set client_id= $2,
  appt_time = $3,
  status = $4,
  note = $5
WHERE id = $1;

-- name: DeleteAppt :exec
DELETE FROM appointments 
WHERE id = $1;


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



-- name: GetUser :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users 
ORDER BY id ASC OFFSET $2 LIMIT $1;

-- name: CreateUser :one
INSERT INTO  users(
  name, email, password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
  set name = $2,
  email = $3,
  password = $4
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetUserCount :one
SELECT count(*)
FROM users;
