// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source Code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package types define all the DNS types used by the server
package types

import (
	"bytes"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// OpCode type
type OpCode struct {
	Name string
	Code int
}

// OpCode constants
var (
	Query  OpCode = OpCode{Name: "Query", Code: 0}
	IQuery OpCode = OpCode{Name: "IQuery", Code: 1}
	Status OpCode = OpCode{Name: "Status", Code: 2}
	Notify OpCode = OpCode{Name: "Notify", Code: 4}
	Update OpCode = OpCode{Name: "Update", Code: 5}
)

// UnmarshalYAML - Function to Unmarshal to YAML
func (opCode *OpCode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := unmarshal(&opCode.Name)

	if err != nil {
		log.Printf("Error when unmarshal YAML OpCode")
	}

	return err
}

// MarshalYAML - Function to Marshal from YAML
func (opCode OpCode) MarshalYAML() (interface{}, error) {
	return &opCode.Name, nil
}

// UnmarshalJSON - Function to Unmarshal to JSON
func (opCode OpCode) UnmarshalJSON(bytes []byte) error {
	qtypeName, err := strconv.Unquote(string(bytes))

	if err != nil {
		log.Printf("Error when unmarshal JSON OpCode: %s", err)
	}

	opCode, err = OpCodeFromString(qtypeName)

	return err
}

// MarshalJSON - Function to Marshal from JSON
func (opCode OpCode) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(fmt.Sprintf("\"%s\"", opCode.Name))

	return buffer.Bytes(), nil
}

func (opCode OpCode) String() string {
	return opCode.Name
}

// OpCodeFromCode - Generate an OpCode struct from a numeric code
// choose one between 1 and 15
func OpCodeFromCode(code int) (OpCode, error) {
	switch {
	case code == Query.Code:
		return Query, nil
	case code == IQuery.Code:
		return IQuery, nil
	case code == Status.Code:
		return Status, nil
	case code == Notify.Code:
		return Notify, nil
	case code == Update.Code:
		return Update, nil
	case code == 3, code >= 6 && code <= 15:
		return OpCode{Name: "AvailableForAssignment", Code: code}, nil
	}

	return OpCode{}, fmt.Errorf("Code %d not available, choose one between 1 and 15", code)
}

// OpCodeFromString - Generate an OpCode struct form a name
// choose one of "Query", "IQuery", "Status", "Notify", "Update" or "AvailableForAssignment"
func OpCodeFromString(name string) (OpCode, error) {
	switch name {
	case Query.Name:
		return Query, nil
	case IQuery.Name:
		return IQuery, nil
	case Status.Name:
		return Status, nil
	case Notify.Name:
		return Notify, nil
	case Update.Name:
		return Update, nil
	case "AvailableForAssignment":
		return OpCode{Name: "AvailableForAssignment", Code: 3}, nil
	}

	return OpCode{}, fmt.Errorf("Name %s not available, choose one of \"Query\", \"IQuery\", \"Status\", \"Notify\", \"Update\" or \"AvailableForAssignment\"", name)
}
