// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package types define all the tests for the types package
package types_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/lucasdc6/gdns/pkg/types"
)

func TestDNSMessageShow(t *testing.T) {
	dnsHeader := &types.DNSMessage{
		Header: types.DNSHeader{
			Identifier:          1,
			QR:                  true,
			OpCode:              types.Query,
			AuthoritativeAnswer: true,
			TruncatedMessage:    false,
			RecursionDesired:    true,
			RecursionAvailable:  true,
			Z:                   false,
			RCode:               types.NoError,
			QDcount:             1,
			ANcount:             0,
			NScount:             0,
			ARcount:             0,
		},
		Questions: []types.DNSQuestion{
			types.DNSQuestion{
				Name:  "www.google.com",
				Type:  types.A,
				Class: types.IN,
			},
		},
		Answers:    []types.DNSResource{},
		Authority:  []types.DNSResource{},
		Additional: []types.DNSResource{},
	}

	dnsHeaderJSON, _ := json.Marshal(dnsHeader)

	fmt.Printf("%+v\n", dnsHeader)
	fmt.Printf("%s", dnsHeaderJSON)
}
