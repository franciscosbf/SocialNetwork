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
	Password string

	Phone *UserPhone

	Name *UserName

	Location *UserLocation
}

func (uf *UserInfo) String() string {
	return fmt.Sprintf(
		"UserInfo[Username: %v, Email: %v, Password: %v, Phone: %v, Name: %v, Location: %v]",
		uf.Username, uf.Email, uf.Password, uf.Phone, uf.Name, uf.Location)
}

// UsersRepository represents all user operations
// (get/delete/insert/update) in the database. Each
// one may return an error if something went wrong
type UsersRepository interface {
	// RegisterUser Inserts a new user. UserPhone
	// and UserName.Middle are optional fields
	RegisterUser(user *UserRegistration) error

	// DeleteUser removes a user from the system
	DeleteUser(username string) error

	GetEmail(username string) (string, error)
	SetEmail(username, email string) error

	MatchesPassword(username, password string) (bool, error)
	SetPassword(username, password string) error

	GetPhone(username string) (*UserPhone, error)
	SetPhone(username string, phone *UserPhone) error

	GetName(username string) (*UserName, error)
	SetName(username string, name *UserName) error

	GetDescription(username string) (string, error)
	SetDescription(username, description string)

	GetLocation(username string) (*UserLocation, error)
	SetLocation(username string, location *UserLocation) error
}

// UsersStore represents all user operations
// (get/delete/insert/update) in the cache. Each
// one may return an error if something went wrong
type UsersStore interface {
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
}

// Users aggregates database
// and cache interfaces
type Users struct {
	store      UsersStore
	repository UsersRepository
}
