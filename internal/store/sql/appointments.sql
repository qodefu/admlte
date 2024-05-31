
CREATE TABLE appointments(
	id        SERIAL PRIMARY KEY, 
	client_id   int,
	appt_time   timestamp,
	status     text,
	note      text,
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
  client_id, appt_time, Status, Note, created
) VALUES (
  $1, $2, $3, $4, $5
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

