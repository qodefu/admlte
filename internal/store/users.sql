
CREATE TABLE users(
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  email  text,
  password text,
  created timestamp
);


-- name: GetAuthor :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM  users 
ORDER BY $1;

-- name: CreateAuthor :one
INSERT INTO  users(
  name, email, password, created
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateAuthor :exec
UPDATE users
  set name = $2,
  email = $3,
  password = $4
WHERE id = $1;

-- name: DeleteAuthor :exec
DELETE FROM users
WHERE id = $1;

