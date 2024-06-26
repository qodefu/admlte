

-- name: GetUser :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * 
FROM users 
WHERE name like '%' || @name::text || '%'
  OR email like '%' || @email::text || '%'
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
