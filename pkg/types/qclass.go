// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package types define all the DNS types used by the server
package types

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// QClass type
type QClass struct {
	Name string
	Code int
}

// QClass constants
var (
	IN QClass = QClass{Name: "IN", Code: 1}
	CS QClass = QClass{Name: "CS", Code: 2}
	CH QClass = QClass{Name: "CH", Code: 3}
	HS QClass = QClass{Name: "HS", Code: 4}
)

// UnmarshalYAML - Function to Unmarshal to YAML
func (qclass *QClass) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := unmarshal(&qclass.Name)

	if err != nil {
		log.Printf("Error when unmarshal YAML QClass")
	}

	return err
}

// MarshalYAML - Function to Marshal from YAML
func (qclass QClass) MarshalYAML() (interface{}, error) {
	return &qclass.Name, nil
}

// UnmarshalJSON - Function to Unmarshal to JSON
func (qclass QClass) UnmarshalJSON(bytes []byte) error {
	qtypeName, err := strconv.Unquote(string(bytes))

	if err != nil {
		log.Printf("Error when unmarshal JSON QClass: %s", err)
	}

	qclass, err = QClassFromString(qtypeName)

	return err
}

// MarshalJSON - Function to Marshal from JSON
func (qclass QClass) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(fmt.Sprintf("\"%s\"", qclass.Name))

	return buffer.Bytes(), nil
}

func (qCLass QClass) String() string {
	return qCLass.Name
}

// QClassFromCode - Generate an QClass struct from a numeric code
// choose one between 1 and 4
func QClassFromCode(code int) (QClass, error) {
	switch code {
	case IN.Code:
		return IN, nil
	case CS.Code:
		return CS, nil
	case CH.Code:
		return CH, nil
	case HS.Code:
		return HS, nil
	}

	return QClass{}, fmt.Errorf("Code %d not available, choose one between 1 and 4", code)
}

// QClassFromString - Generate an OpCode struct form a Name
// choose one of "IN", "CS", "CH" or "HS"
func QClassFromString(name string) (QClass, error) {
	switch strings.ToUpper(name) {
	case IN.Name:
		return IN, nil
	case CS.Name:
		return CS, nil
	case CH.Name:
		return CH, nil
	case HS.Name:
		return HS, nil
	}

	return QClass{}, fmt.Errorf("Name %s not available, choose one of \"IN\", \"CS\", \"CH\", \"HS\"", name)
}
