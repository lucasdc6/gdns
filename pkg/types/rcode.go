// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package types define all the DNS types used by the server
package types

import (
	"bytes"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// RCode type
type RCode struct {
	Name string
	Code int
}

// RCode constants
var (
	NoError          RCode = RCode{Name: "NoError", Code: 0}
	FormatError      RCode = RCode{Name: "FormatError", Code: 1}
	ServerFailure    RCode = RCode{Name: "ServerFailure", Code: 2}
	NXDomain         RCode = RCode{Name: "NXDomain", Code: 3}
	NotImplemented   RCode = RCode{Name: "NotImplemented", Code: 4}
	Refuced          RCode = RCode{Name: "Refuced", Code: 5}
	YXDomain         RCode = RCode{Name: "YXDomain", Code: 6}
	YXRRSet          RCode = RCode{Name: "YXRRSet", Code: 7}
	NXRRSet          RCode = RCode{Name: "NXRRSet", Code: 8}
	NotAuthoritative RCode = RCode{Name: "NotAuthoritative", Code: 9}
	NotZone          RCode = RCode{Name: "NotZone", Code: 10}
	BadOptVersion    RCode = RCode{Name: "BadOptVersion", Code: 16}
	BADSIG           RCode = RCode{Name: "BADSIG", Code: 16}
	BADKEY           RCode = RCode{Name: "BADKEY", Code: 17}
	BADTIME          RCode = RCode{Name: "BADTIME", Code: 18}
	BADMODE          RCode = RCode{Name: "BADMODE", Code: 19}
	BADNAME          RCode = RCode{Name: "BADNAME", Code: 20}
	BADALG           RCode = RCode{Name: "BADALG", Code: 21}
)

// UnmarshalYAML - Function to Unmarshal to YAML
func (rcode *RCode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := unmarshal(&rcode.Name)

	if err != nil {
		log.Printf("Error when unmarshal YAML RCode")
	}

	return err
}

// MarshalYAML - Function to Marshal from YAML
func (rcode RCode) MarshalYAML() (interface{}, error) {
	return &rcode.Name, nil
}

// UnmarshalJSON - Function to Unmarshal to JSON
func (rcode RCode) UnmarshalJSON(bytes []byte) error {
	qtypeName, err := strconv.Unquote(string(bytes))

	if err != nil {
		log.Printf("Error when unmarshal JSON RCode: %s", err)
	}

	rcode, err = RCodeFromString(qtypeName)

	return err
}

// MarshalJSON - Function to Marshal from JSON
func (rcode RCode) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(fmt.Sprintf("\"%s\"", rcode.Name))

	return buffer.Bytes(), nil
}

func (rcode RCode) String() string {
	return rcode.Name
}

// RCodeFromCode - Generate an RCode struct from a numeric code
// choose one between 1 and 4
func RCodeFromCode(code int) (RCode, error) {
	switch {
	case code == NoError.Code:
		return NoError, nil
	case code == FormatError.Code:
		return FormatError, nil
	case code == ServerFailure.Code:
		return ServerFailure, nil
	case code == NXDomain.Code:
		return NXDomain, nil
	case code == NotImplemented.Code:
		return NotImplemented, nil
	case code == Refuced.Code:
		return Refuced, nil
	case code == YXDomain.Code:
		return YXDomain, nil
	case code == YXRRSet.Code:
		return YXRRSet, nil
	case code == NXRRSet.Code:
		return NXRRSet, nil
	case code == NotAuthoritative.Code:
		return NotAuthoritative, nil
	case code == NotZone.Code:
		return NotZone, nil
	case code == BadOptVersion.Code:
		return BadOptVersion, nil
	case code == BADKEY.Code:
		return BADKEY, nil
	case code == BADTIME.Code:
		return BADTIME, nil
	case code == BADMODE.Code:
		return BADMODE, nil
	case code == BADNAME.Code:
		return BADNAME, nil
	case code == BADALG.Code:
		return BADALG, nil
	case code >= 11 && code <= 15, code >= 22 && code <= 3840, code >= 4096 && code <= 65535:
		return RCode{Name: "AvailableForAssignment", Code: code}, nil
	case code >= 3841 && code <= 4095:
		return RCode{Name: "PrivateUse", Code: code}, nil
	}

	return RCode{}, fmt.Errorf("Code %d not available, choose one between 1 and 65535", code)
}

// RCodeFromString - Generate an RCode struct from a numeric code
// choose one between 1 and 4
func RCodeFromString(name string) (RCode, error) {
	switch name {
	case NoError.Name:
		return NoError, nil
	case FormatError.Name:
		return FormatError, nil
	case ServerFailure.Name:
		return ServerFailure, nil
	case NXDomain.Name:
		return NXDomain, nil
	case NotImplemented.Name:
		return NotImplemented, nil
	case Refuced.Name:
		return Refuced, nil
	case YXDomain.Name:
		return YXDomain, nil
	case YXRRSet.Name:
		return YXRRSet, nil
	case NXRRSet.Name:
		return NXRRSet, nil
	case NotAuthoritative.Name:
		return NotAuthoritative, nil
	case NotZone.Name:
		return NotZone, nil
	case BadOptVersion.Name:
		return BadOptVersion, nil
	case BADKEY.Name:
		return BADKEY, nil
	case BADTIME.Name:
		return BADTIME, nil
	case BADMODE.Name:
		return BADMODE, nil
	case BADNAME.Name:
		return BADNAME, nil
	case BADALG.Name:
		return BADALG, nil
	case "AvailableForAssignment":
		return RCode{Name: "AvailableForAssignment", Code: 65535}, nil
	case "PrivateUse":
		return RCode{Name: "PrivateUse", Code: 4095}, nil
	}

	return RCode{}, fmt.Errorf("Name %s not available, choose one of NoError, FormatError, ServerFailure, NXDomain, NotImplemented, Refuced, YXDomain, YXRRSet, NXRRSet, NotAuthoritative, NotZone, BadOptVersion, BADSIG, BADKEY, BADTIME, BADMODE, BADNAME, BADALG, PrivateUse or AvailableForAssignment", name)
}
