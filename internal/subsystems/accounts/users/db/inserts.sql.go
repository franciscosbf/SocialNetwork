// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: inserts.sql

package db

import (
	"context"
	"database/sql"
)

const insertNewAccount = `-- name: InsertNewAccount :one
INSERT INTO users_info.accounts (username, email, password, phone_prefix, phone_number)
VALUES ($1, $2, $3, $4, $5)
RETURNING aid
`

type InsertNewAccountParams struct {
	Username    string
	Email       string
	Password    string
	PhonePrefix sql.NullInt32
	PhoneNumber sql.NullInt32
}

func (q *Queries) InsertNewAccount(ctx context.Context, arg InsertNewAccountParams) (int64, error) {
	row := q.db.QueryRow(ctx, insertNewAccount,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.PhonePrefix,
		arg.PhoneNumber,
	)
	var aid int64
	err := row.Scan(&aid)
	return aid, err
}

const insertNewProfile = `-- name: InsertNewProfile :exec
INSERT INTO users_info.profiles (aid, first_name, middle_name, surname)
VALUES ($1, $2, $3, $4)
`

type InsertNewProfileParams struct {
	Aid        sql.NullInt64
	FirstName  string
	MiddleName sql.NullString
	Surname    string
}

func (q *Queries) InsertNewProfile(ctx context.Context, arg InsertNewProfileParams) error {
	_, err := q.db.Exec(ctx, insertNewProfile,
		arg.Aid,
		arg.FirstName,
		arg.MiddleName,
		arg.Surname,
	)
	return err
}
