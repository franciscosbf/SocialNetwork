// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/twpayne/go-geom"
)

type UsersInfoAccount struct {
	Aid         uuid.UUID
	Username    string
	Email       string
	Password    string
	PhonePrefix sql.NullInt32
	PhoneNumber sql.NullInt32
}

type UsersInfoProfile struct {
	Pid         uuid.UUID
	Aid         uuid.NullUUID
	Description sql.NullString
	FirstName   string
	MiddleName  sql.NullString
	Surname     string
	Location    geom.MultiPolygon
}