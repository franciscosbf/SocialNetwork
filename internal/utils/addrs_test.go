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
	"fmt"
	"testing"
)

func TestValidAddrsList(t *testing.T) {
	addrs, err := ParseAddrs("		\tlocalhost:123		\n\r,  127.0.0.1:123  ")

	if err != nil {
		t.Errorf("Unexpected error %v", err)
		return
	}

	if addrs == nil {
		t.Error("Addrs is nil")
		return
	}

	if len(addrs.Bucket) != 2 {
		t.Errorf("Got %v, not exactly 2 with localhost:123;127.0.0.1:123", len(addrs.Bucket))
		return
	}

	first := fmt.Sprintf("%v:%v", addrs.Bucket[0].Host, addrs.Bucket[0].Port)
	if first != "localhost:123" {
		t.Errorf("Expecting first addr to be localhost:123, got %v", first)
	}

	second := fmt.Sprintf("%v:%v", addrs.Bucket[1].Host, addrs.Bucket[1].Port)
	if second != "127.0.0.1:123" {
		t.Errorf("Expecting second addr to be 127.0.0.1:123, got %v", second)
	}
}

func TestInvalidAddrsList(t *testing.T) {
	testBattery := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "TestBadFmtPrefix",
			test: func(t *testing.T) {
				addrs, err := ParseAddrs(",localhost:123,127.0.0.1:123")

				if err != InvalidAddrsListError {
					t.Errorf("Unexpected error %v", err)
					return
				}

				if addrs != nil {
					t.Error("Expecting addrs to be nil")
					return
				}
			},
		},
		{
			name: "TestBadFmtSuffix",
			test: func(t *testing.T) {
				addrs, err := ParseAddrs("localhost:123,127.0.0.1:123,")

				if err != InvalidAddrsListError {
					t.Errorf("Unexpected error %v", err)
					return
				}

				if addrs != nil {
					t.Error("Expecting addrs to be nil")
					return
				}
			},
		},
		{
			name: "TestBadAddrFmt",
			test: func(t *testing.T) {
				addrs, err := ParseAddrs("localhost:123,  \n  \t		,127.0.0.1:123")

				if err != InvalidAddrsListError {
					t.Errorf("Unexpected error %v", err)
					return
				}

				if addrs != nil {
					t.Error("Expecting addrs to be nil")
					return
				}
			},
		},
		{
			name: "TestWithDuplicatedAddr",
			test: func(t *testing.T) {
				addrs, err := ParseAddrs("localhost:123,127.0.0.1:123,127.0.0.1:123")

				if _, ok := err.(*DuplicatedAddrError); !ok {
					t.Errorf("Unexpected error %v", err)
					return
				}

				if addrs != nil {
					t.Error("Expecting addrs to be nil")
					return
				}
			},
		},
		{
			name: "TestWithInvalidAddr",
			test: func(t *testing.T) {
				addrs, err := ParseAddrs("[:123,127.0.0.1:123")

				if _, ok := err.(*InvalidAddrError); !ok {
					t.Errorf("Unexpected error %v", err)
					return
				}

				if addrs != nil {
					t.Error("Expecting addrs to be nil")
					return
				}
			},
		},
	}

	for _, pair := range testBattery {
		t.Run(pair.name, pair.test)
	}
}
