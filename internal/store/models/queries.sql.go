// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAppt = `-- name: CreateAppt :one
INSERT INTO appointments(
  client_id, appt_time, Status, Note
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, client_id, appt_time, status, note, created
`

type CreateApptParams struct {
	ClientID pgtype.Int8
	ApptTime pgtype.Timestamp
	Status   pgtype.Text
	Note     pgtype.Text
}

func (q *Queries) CreateAppt(ctx context.Context, arg CreateApptParams) (Appointment, error) {
	row := q.db.QueryRow(ctx, createAppt,
		arg.ClientID,
		arg.ApptTime,
		arg.Status,
		arg.Note,
	)
	var i Appointment
	err := row.Scan(
		&i.ID,
		&i.ClientID,
		&i.ApptTime,
		&i.Status,
		&i.Note,
		&i.Created,
	)
	return i, err
}

const createClient = `-- name: CreateClient :one
INSERT INTO  clients(
  name
) VALUES (
  $1
)
RETURNING id, name, created
`

func (q *Queries) CreateClient(ctx context.Context, name string) (Client, error) {
	row := q.db.QueryRow(ctx, createClient, name)
	var i Client
	err := row.Scan(&i.ID, &i.Name, &i.Created)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO  users(
  name, email, password
) VALUES (
  $1, $2, $3
)
RETURNING id, name, email, password, created
`

type CreateUserParams struct {
	Name     string
	Email    pgtype.Text
	Password pgtype.Text
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Name, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Created,
	)
	return i, err
}

const deleteAppt = `-- name: DeleteAppt :exec
DELETE FROM appointments 
WHERE id = $1
`

func (q *Queries) DeleteAppt(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAppt, id)
	return err
}

const deleteClient = `-- name: DeleteClient :exec
DELETE FROM clients 
WHERE id = $1
`

func (q *Queries) DeleteClient(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteClient, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getAppointment = `-- name: GetAppointment :one
SELECT c.name, a.id, a.client_id, a.appt_time, a.status, a.note, a.created  
FROM  appointments a
JOIN clients c
  ON a.client_id = c.id
WHERE a.id = $1 LIMIT 1
`

type GetAppointmentRow struct {
	Name     string
	ID       int64
	ClientID pgtype.Int8
	ApptTime pgtype.Timestamp
	Status   pgtype.Text
	Note     pgtype.Text
	Created  pgtype.Timestamp
}

func (q *Queries) GetAppointment(ctx context.Context, id int64) (GetAppointmentRow, error) {
	row := q.db.QueryRow(ctx, getAppointment, id)
	var i GetAppointmentRow
	err := row.Scan(
		&i.Name,
		&i.ID,
		&i.ClientID,
		&i.ApptTime,
		&i.Status,
		&i.Note,
		&i.Created,
	)
	return i, err
}

const getAppointmentCount = `-- name: GetAppointmentCount :one
SELECT count(id) FROM appointments
`

func (q *Queries) GetAppointmentCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, getAppointmentCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getClient = `-- name: GetClient :one
SELECT id, name, created FROM clients 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetClient(ctx context.Context, id int64) (Client, error) {
	row := q.db.QueryRow(ctx, getClient, id)
	var i Client
	err := row.Scan(&i.ID, &i.Name, &i.Created)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, name, email, password, created FROM users 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Created,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, password, created FROM users 
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email pgtype.Text) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Created,
	)
	return i, err
}

const getUserClient = `-- name: GetUserClient :one
SELECT u.id, c.name FROM clients c 
JOIN users u
on c.id = u.id 
WHERE u.id = $1 LIMIT 1
`

type GetUserClientRow struct {
	ID   int64
	Name string
}

func (q *Queries) GetUserClient(ctx context.Context, id int64) (GetUserClientRow, error) {
	row := q.db.QueryRow(ctx, getUserClient, id)
	var i GetUserClientRow
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getUserCount = `-- name: GetUserCount :one
SELECT count(*)
FROM users
`

func (q *Queries) GetUserCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, getUserCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const listAppt = `-- name: ListAppt :many
SELECT c.name, a.id, a.client_id, a.appt_time, a.status, a.note, a.created  
FROM  appointments a
JOIN clients c
  ON a.client_id = c.id
ORDER BY a.id ASC 
OFFSET $2 LIMIT $1
`

type ListApptParams struct {
	Limit  int32
	Offset int32
}

type ListApptRow struct {
	Name     string
	ID       int64
	ClientID pgtype.Int8
	ApptTime pgtype.Timestamp
	Status   pgtype.Text
	Note     pgtype.Text
	Created  pgtype.Timestamp
}

func (q *Queries) ListAppt(ctx context.Context, arg ListApptParams) ([]ListApptRow, error) {
	rows, err := q.db.Query(ctx, listAppt, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListApptRow
	for rows.Next() {
		var i ListApptRow
		if err := rows.Scan(
			&i.Name,
			&i.ID,
			&i.ClientID,
			&i.ApptTime,
			&i.Status,
			&i.Note,
			&i.Created,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listClients = `-- name: ListClients :many
SELECT id, name, created FROM  clients 
ORDER BY id
`

func (q *Queries) ListClients(ctx context.Context) ([]Client, error) {
	rows, err := q.db.Query(ctx, listClients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Client
	for rows.Next() {
		var i Client
		if err := rows.Scan(&i.ID, &i.Name, &i.Created); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, email, password, created FROM users 
ORDER BY id ASC OFFSET $2 LIMIT $1
`

type ListUsersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Created,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAppt = `-- name: UpdateAppt :exec
UPDATE appointments 
  set client_id= $2,
  appt_time = $3,
  status = $4,
  note = $5
WHERE id = $1
`

type UpdateApptParams struct {
	ID       int64
	ClientID pgtype.Int8
	ApptTime pgtype.Timestamp
	Status   pgtype.Text
	Note     pgtype.Text
}

func (q *Queries) UpdateAppt(ctx context.Context, arg UpdateApptParams) error {
	_, err := q.db.Exec(ctx, updateAppt,
		arg.ID,
		arg.ClientID,
		arg.ApptTime,
		arg.Status,
		arg.Note,
	)
	return err
}

const updateClient = `-- name: UpdateClient :exec
UPDATE clients 
  set name = $2
WHERE id = $1
`

type UpdateClientParams struct {
	ID   int64
	Name string
}

func (q *Queries) UpdateClient(ctx context.Context, arg UpdateClientParams) error {
	_, err := q.db.Exec(ctx, updateClient, arg.ID, arg.Name)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
  set name = $2,
  email = $3,
  password = $4
WHERE id = $1
`

type UpdateUserParams struct {
	ID       int64
	Name     string
	Email    pgtype.Text
	Password pgtype.Text
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
	)
	return err
}
