package utils

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

// Address contains
// its host and port
type Address struct {
	Host string
	Port string
}

// Addrs represents a collection
// of addresses
type Addrs struct {
	Bucket []*Address
}

// InvalidAddrsListError represents a bad formatted list of addresses
var InvalidAddrsListError = errors.New(
	"invalid format. expects host1:port1;host2:port2")

// DuplicatedAddrError represents
// a repeated address
type DuplicatedAddrError struct {
	rawAddr string
}

func (e *DuplicatedAddrError) Error() string {
	return fmt.Sprintf("duplicated address %v", e.rawAddr)
}

// InvalidAddrError represents an
// address with an invalid format
type InvalidAddrError struct {
	rawAddr string
	origin  error
}

func (e *InvalidAddrError) Error() string {
	return fmt.Sprintf(
		"address %v has invalid format; error: %v",
		e.rawAddr, e.origin)
}

// ParseAddrs parses a list of addresses with the format
// host1:port1;host2:port2;(...). If the host is an ipv6,
// then the address should have the following schema with
// square brackets: [host]:port. Returns InvalidAddrsListError
// if the format is invalid, DuplicatedAddrError if there's
// a defined address more than once and InvalidAddrError if
// an address doesn't match the format host:port
func ParseAddrs(addrsList string) (*Addrs, error) {
	addrsList = PolishString(addrsList)

	if addrsList == "" || strings.HasPrefix(addrsList, ";") ||
		strings.HasSuffix(addrsList, ";") {
		return nil, InvalidAddrsListError
	}

	// Cache to look for duplicates
	var parsed = NewSet[string]()

	var addrs []*Address

	for _, rawAddr := range strings.Split(addrsList, ";") {
		rawAddr = PolishString(rawAddr)

		if rawAddr == "" {
			return nil, InvalidAddrsListError
		}

		if parsed.Contains(rawAddr) {
			return nil, &DuplicatedAddrError{
				rawAddr: rawAddr,
			}
		}

		host, port, err := net.SplitHostPort(rawAddr)
		if err != nil {
			return nil, &InvalidAddrError{
				rawAddr: rawAddr,
				origin:  err,
			}
		}

		parsed.Put(rawAddr)

		addrs = append(addrs, &Address{
			Host: host,
			Port: port,
		})
	}

	return &Addrs{
		Bucket: addrs,
	}, nil
}
