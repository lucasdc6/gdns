// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errors define all the relationated with
// error managment
package errors

// List of errors constants
const (
	StartingServer              = 1
	RetrivingUDPPackage         = 2
	EstablishingTCPConn         = 3
	RetrivingTCPData            = 4
	ParsingHeader               = 5
	ParsingQuestions            = 6
	ParsingAnswers              = 7
	ParsingAuthority            = 8
	ParsingAdditional           = 9
	ModuleDocumentationNotFound = 10
	ModuleNotFound              = 11
	ReadingYAMLConfiguration    = 12
	ReadingJSONConfiguration    = 13
	ReadingXMLConfiguration     = 14
	OpeningConfigurationFile    = 15
	OpCodeNotFound              = 16
	RCodeNotFound               = 17
	QTypeNotFound               = 18
	QClassNotFound              = 19
)
