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

package storage

import (
	"fmt"
)

// UserPhone represents an optional user
// phone containing its country code and number
type UserPhone struct {
	Prefix int
	Number int
}

func (p *UserPhone) String() string {
	return fmt.Sprintf(
		"UserPhone[Prefix: %v, Number: %v]",
		p.Prefix, p.Number)
}

// UserName represents the name elements
type UserName struct {
	First   string
	Middle  string
	Surname string
}

func (un *UserName) String() string {
	return fmt.Sprintf(
		"UserName[First: %v, Middle: %v, Surname: %v]",
		un.First, un.Middle, un.Surname)
}

// UserRegistration represents the
// required elements to register a user
type UserRegistration struct {
	Username string
	Email    string
	Password string

	Phone *UserPhone

	Name *UserName
}

func (ur *UserRegistration) String() string {
	return fmt.Sprintf(
		"UserRegistration[Username: %v, Email: %v, Password: %v, Phone: %v, Name: %v]",
		ur.Username, ur.Email, ur.Password, ur.Phone, ur.Name)
}

// UserLocation represents an optional location
// where the user is currently living. Expects
// to follow the SRID 4326 (WGS84 long/lat)
type UserLocation struct {
	Longitude float64
	Latitude  float64
}

func (ul *UserLocation) String() string {
	return fmt.Sprintf(
		"UserLocation[Longitude: %v, Latitude: %v]",
		ul.Longitude, ul.Latitude)
}

// UserInfo represents all info
// about a registered user
type UserInfo struct {
	Username string
	Email    string

	Phone *UserPhone

	Name *UserName

	Location *UserLocation
}

func (uf *UserInfo) String() string {
	return fmt.Sprintf(
		"UserInfo[Username: %v, Email: %v, Phone: %v, Name: %v, Location: %v]",
		uf.Username, uf.Email, uf.Phone, uf.Name, uf.Location)
}

// UsersRepository represents all user operations
// (get/delete/insert/update) in the database. Each
// one may return an error if something went wrong
type UsersRepository interface {
	GetUser(username string) *UserInfo
	SetUser(user *UserRegistration) error
	DeleteUser(username string) error

	GetEmail(username string) (string, error)
	SetEmail(username, email string) error

	GetPhone(username string) (*UserPhone, error)
	SetPhone(username string, phone *UserPhone) error

	GetName(username string) (*UserName, error)
	SetName(username string, name *UserName) error

	GetDescription(username string) (string, error)
	SetDescription(username, description string)

	GetLocation(username string) (*UserLocation, error)
	SetLocation(username string, location *UserLocation) error

	MatchesPassword(username, password string) (bool, error)
	SetPassword(username, password string) error
}

// UsersStore represents all user operations
// (get/delete/insert/update) in the cache. Each
// one may return an error if something went wrong
type UsersStore interface {
	GetUser(username string) (*UserInfo, error)
	SetUser(user *UserInfo) error
	DeleteUser(username string) error

	GetEmail(username string) (string, error)
	SetEmail(username, email string) error
	DeleteEmail(username string) error

	GetPhone(username string) (*UserPhone, error)
	SetPhone(username string, phone *UserPhone) error
	DeletePhone(username string) error

	GetName(username string) (*UserName, error)
	SetName(username string, name *UserName) error
	DeleteName(username string) error

	GetDescription(username string) (string, error)
	SetDescription(username, description string)
	DeleteDescription(username string) error

	GetLocation(username string) (*UserLocation, error)
	SetLocation(username string, location *UserLocation) error
	DeleteLocation(username string) error
}

// Users aggregates database
// and cache interfaces
type Users struct {
	store      UsersStore
	repository UsersRepository
}

// NewUsers Creates a new users aggregator
func NewUsers(store UsersStore, repository UsersRepository) *Users {
	return &Users{store: store, repository: repository}
}
