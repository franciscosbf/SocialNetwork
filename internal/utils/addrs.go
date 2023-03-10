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
	// Host may contain the host name
	// or the ip. This field can be
	// empty, if it wasn't specified
	Host string
	// Port is always present
	Port string
}

// Addrs represents a collection
// of addresses
type Addrs struct {
	Bucket []*Address
}

// InvalidAddrsListError represents a bad formatted list of addresses
var InvalidAddrsListError = errors.New(
	"invalid format. expects host1:port1,host2:port2")

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
// host1:port1,host2:port2,(...). If the host is an ipv6,
// then the address should have the following schema with
// square brackets: [host]:port. Returns InvalidAddrsListError
// if the format is invalid, DuplicatedAddrError if there's
// a defined address more than once and InvalidAddrError if
// an address doesn't match the format host:port. Keep in mind
// that this doesn't check if the port is valid or the host ip
// has a valid address
func ParseAddrs(addrsList string) (*Addrs, error) {
	addrsList = PolishString(addrsList)

	if addrsList == "" || strings.HasPrefix(addrsList, ",") ||
		strings.HasSuffix(addrsList, ",") {
		return nil, InvalidAddrsListError
	}

	// Cache to look for duplicates
	var parsed = NewSet[string]()

	var addrs []*Address

	for _, rawAddr := range strings.Split(addrsList, ",") {
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
