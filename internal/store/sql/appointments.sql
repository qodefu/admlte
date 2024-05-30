
CREATE TABLE appointments(
	Id        SERIAL PRIMARY KEY, 
	ClientId   int,
	ApptTime   timestamp,
	Status     text,
	Note      text,
  created timestamp 
);


-- name: GetAppointment :one
SELECT * FROM appointments 
WHERE id = $1 LIMIT 1;

-- name: ListAppt :many
SELECT * FROM  appointments 
ORDER BY $1;

-- name: CreateAppt :one
INSERT INTO appointments(
  clientId, ApptTime, Status, Note, created
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateAppt :exec
UPDATE appointments 
  set clientId= $2,
  apptTime = $3,
  status = $4,
  note = $5
WHERE id = $1;

-- name: DeleteAppt :exec
DELETE FROM appointments 
WHERE id = $1;

