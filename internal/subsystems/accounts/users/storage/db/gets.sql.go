/*
Copyright 2023 Francisco Simões Braço-Forte

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: gets.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/twpayne/go-geom"
)

const getAccount = `-- name: GetAccount :one

SELECT
    aid,
    email,
    phone_prefix,
    phone_number
FROM users_info.accounts
WHERE username = $1
`

type GetAccountRow struct {
	Aid         uuid.UUID
	Email       string
	PhonePrefix sql.NullInt32
	PhoneNumber sql.NullInt32
}

// Accounts related queries.
func (q *Queries) GetAccount(ctx context.Context, username string) (GetAccountRow, error) {
	row := q.db.QueryRow(ctx, getAccount, username)
	var i GetAccountRow
	err := row.Scan(
		&i.Aid,
		&i.Email,
		&i.PhonePrefix,
		&i.PhoneNumber,
	)
	return i, err
}

const getAccountInternals = `-- name: GetAccountInternals :one
SELECT aid, email
FROM users_info.accounts
WHERE username = $1
`

type GetAccountInternalsRow struct {
	Aid   uuid.UUID
	Email string
}

func (q *Queries) GetAccountInternals(ctx context.Context, username string) (GetAccountInternalsRow, error) {
	row := q.db.QueryRow(ctx, getAccountInternals, username)
	var i GetAccountInternalsRow
	err := row.Scan(&i.Aid, &i.Email)
	return i, err
}

const getAccountInternalsAuth = `-- name: GetAccountInternalsAuth :one
SELECT aid, email
FROM users_info.accounts
WHERE username = $1 AND
      password = $2
`

type GetAccountInternalsAuthParams struct {
	Username string
	Password string
}

type GetAccountInternalsAuthRow struct {
	Aid   uuid.UUID
	Email string
}

func (q *Queries) GetAccountInternalsAuth(ctx context.Context, arg GetAccountInternalsAuthParams) (GetAccountInternalsAuthRow, error) {
	row := q.db.QueryRow(ctx, getAccountInternalsAuth, arg.Username, arg.Password)
	var i GetAccountInternalsAuthRow
	err := row.Scan(&i.Aid, &i.Email)
	return i, err
}

const getAccountPhone = `-- name: GetAccountPhone :one
SELECT phone_prefix, phone_number
FROM users_info.accounts
WHERE aid = $1
`

type GetAccountPhoneRow struct {
	PhonePrefix sql.NullInt32
	PhoneNumber sql.NullInt32
}

func (q *Queries) GetAccountPhone(ctx context.Context, aid uuid.UUID) (GetAccountPhoneRow, error) {
	row := q.db.QueryRow(ctx, getAccountPhone, aid)
	var i GetAccountPhoneRow
	err := row.Scan(&i.PhonePrefix, &i.PhoneNumber)
	return i, err
}

const getProfile = `-- name: GetProfile :one

SELECT
    aid,
    description,
    first_name,
    middle_name,
    surname,
    location
FROM users_info.profiles
WHERE aid = $1
`

type GetProfileRow struct {
	Aid         uuid.NullUUID
	Description sql.NullString
	FirstName   string
	MiddleName  sql.NullString
	Surname     string
	Location    geom.MultiPolygon
}

// Profiles related queries.
func (q *Queries) GetProfile(ctx context.Context, aid uuid.NullUUID) (GetProfileRow, error) {
	row := q.db.QueryRow(ctx, getProfile, aid)
	var i GetProfileRow
	err := row.Scan(
		&i.Aid,
		&i.Description,
		&i.FirstName,
		&i.MiddleName,
		&i.Surname,
		&i.Location,
	)
	return i, err
}

const getProfileDescription = `-- name: GetProfileDescription :one
SELECT description
FROM users_info.profiles
WHERE aid = $1
`

func (q *Queries) GetProfileDescription(ctx context.Context, aid uuid.NullUUID) (sql.NullString, error) {
	row := q.db.QueryRow(ctx, getProfileDescription, aid)
	var description sql.NullString
	err := row.Scan(&description)
	return description, err
}

const getProfileLocation = `-- name: GetProfileLocation :one
SELECT location
FROM users_info.profiles
WHERE aid = $1
`

func (q *Queries) GetProfileLocation(ctx context.Context, aid uuid.NullUUID) (geom.MultiPolygon, error) {
	row := q.db.QueryRow(ctx, getProfileLocation, aid)
	var location geom.MultiPolygon
	err := row.Scan(&location)
	return location, err
}

const getProfileName = `-- name: GetProfileName :one
SELECT first_name, middle_name, surname
FROM users_info.profiles
WHERE aid = $1
`

type GetProfileNameRow struct {
	FirstName  string
	MiddleName sql.NullString
	Surname    string
}

func (q *Queries) GetProfileName(ctx context.Context, aid uuid.NullUUID) (GetProfileNameRow, error) {
	row := q.db.QueryRow(ctx, getProfileName, aid)
	var i GetProfileNameRow
	err := row.Scan(&i.FirstName, &i.MiddleName, &i.Surname)
	return i, err
}
