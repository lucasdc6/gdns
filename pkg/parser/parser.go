// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package parser define the function used to parse the
// DNS package
package parser

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"

	log "github.com/sirupsen/logrus"

	"os"

	"github.com/lucasdc6/gdns/pkg/errors"
	"github.com/lucasdc6/gdns/pkg/types"
)

func mapBytesToString(r rune) rune {
	if r == 3 || r == 2 || r == 6 {
		return '.'
	}
	return r
}

func getName(chunk []byte) (string, []byte, error) {
	chunks := bytes.SplitN(chunk, []byte{0}, 2)
	name := bytes.Map(mapBytesToString, chunks[0])

	return string(name), chunks[1], nil
}

func parseDNSHeader(query []byte) (header types.DNSHeader, err error) {
	log.Trace("Parsing header")

	qr := query[2] >> 7
	opCode := query[2] >> 3 &^ 16
	truncatedMessage := query[2] >> 1 &^ 252
	recursionDesired := query[2] &^ 254
	recursionAvailable := query[3] >> 7
	z := query[3] >> 6 &^ 254
	ad := query[3] >> 5 &^ 254
	cd := query[3] >> 4 &^ 254
	rCode := query[3] &^ 240

	opCodeStruct, err := types.OpCodeFromCode(int(opCode))

	if err != nil {
		log.Errorf("Error parsing DNS header: %s", err)
		os.Exit(errors.OpCodeNotFound)
	}

	rCodeStruct, err := types.RCodeFromCode(int(rCode))

	if err != nil {
		log.Errorf("Error parsing DNS header: %s", err)
		os.Exit(errors.RCodeNotFound)
	}

	header = types.DNSHeader{
		Identifier:         binary.BigEndian.Uint16(query[0:2]),
		QR:                 !(qr == 0),
		OpCode:             opCodeStruct,
		TruncatedMessage:   !(truncatedMessage == 0),
		RecursionDesired:   !(recursionDesired == 0),
		RecursionAvailable: !(recursionAvailable == 0),
		Z:                  !(z == 0),
		AD:                 !(ad == 0),
		CD:                 !(cd == 0),
		RCode:              rCodeStruct,
		QDcount:            binary.BigEndian.Uint16(query[4:6]),
		ANcount:            binary.BigEndian.Uint16(query[6:8]),
		NScount:            binary.BigEndian.Uint16(query[8:10]),
		ARcount:            binary.BigEndian.Uint16(query[10:12]),
	}

	return header, nil
}

func parseDNSQuestions(query []byte, questionsCount uint16) (questions []types.DNSQuestion, last uint16, err error) {
	log.Trace("Parsing questions")
	chunk := query
	last = 0

	if questionsCount == 0 {
		return questions, last, nil
	}

	for i := 0; i < int(questionsCount); i++ {
		name := ""
		/*
		 * RFC 1035 - Section 4.1.4 Message compression
		 * Check for compression
		 */
		if uint16(chunk[0]&192) > 0 {
			log.Trace("Follow pointer")
			nameOffset := (binary.BigEndian.Uint16(chunk[0:2]) & 16383) - 12
			name, chunk, _ = getName(query[nameOffset:])
		} else {
			log.Trace("  Get name")
			name, chunk, _ = getName(chunk)
		}

		log.Debug("Name: %s", name)
		last += uint16(len(name))

		qtypeChunk := binary.BigEndian.Uint16(chunk[0:2])
		qtype, err := types.QTypeFromCode(int(qtypeChunk))

		if err != nil {
			log.Errorf("Error parsing DNS questions #%d: %s", i, err)
			os.Exit(errors.QTypeNotFound)
		}

		classChunk := binary.BigEndian.Uint16(chunk[2:4])
		class, err := types.QClassFromCode(int(classChunk))

		if err != nil {
			log.Errorf("Error parsing DNS questions #%d: %s", i, err)
			os.Exit(errors.QClassNotFound)
		}
		last += 4

		questions = append(questions, types.DNSQuestion{
			Name:  string(name),
			Type:  qtype,
			Class: class,
		})
	}

	return questions, last, nil
}

func parseDNSAnswers(query []byte, offset, answersCount uint16) (answers []types.DNSResource, last uint16, err error) {
	log.WithFields(log.Fields{
		"relativeOffset": offset,
		"absoluteOffset": offset + 12,
	}).Debug("Parsing answers")
	last = offset
	chunk := query[offset:]

	if answersCount == 0 {
		return answers, last, nil
	}

	for i := 0; i < int(answersCount); i++ {
		log.Debug("Data:\n%s", hex.Dump(chunk))
		name := ""

		/*
		 * RFC 1035 - Section 4.1.4 Message compression
		 * Check for compression
		 */
		if uint16(chunk[0]&192) > 0 {
			log.Trace("Follow pointer")
			nameOffset := (binary.BigEndian.Uint16(chunk[0:2]) & 16383) - 12
			name, chunk, _ = getName(query[nameOffset:])
		} else {
			log.Trace("Get name")
			name, chunk, _ = getName(chunk)
		}
		log.Debugf("name: %s", name)
		nameAddress := uint16(chunk[0])
		num := binary.BigEndian.Uint16(chunk[0:2])
		log.Debugf("op %b AND 11000000 = %d", nameAddress, nameAddress&192)
		log.Printf("op %b AND 00111111 11111111 = %d", num, num&16383)

	}
	return answers, last, nil
}

func parseDNSAuthority(query []byte, authorityCount uint16) (authority []types.DNSResource, err error) {
	authority = []types.DNSResource{}

	return authority, nil
}

func parseDNSAdditional(query []byte, additionalCount uint16) (additional []types.DNSResource, err error) {
	additional = []types.DNSResource{}

	return additional, nil
}

// ParseDNSQuery - Parse the query and return a DNSMessage
func ParseDNSQuery(query []byte) types.DNSMessage {
	headerBytes := query[0:12]
	header, err := parseDNSHeader(headerBytes)

	if err != nil {
		log.Fatalf("Error parsing header: %s", err)
		os.Exit(errors.ParsingHeader)
	}

	questions, last, err := parseDNSQuestions(query[12:], header.QDcount)

	if err != nil {
		log.Fatalf("Error parsing questions: %s", err)
		os.Exit(errors.ParsingQuestions)
	}
	last++

	answers, _, err := parseDNSAnswers(query[12:], last, header.ANcount)

	if err != nil {
		log.Fatalf("Error parsing answers: %s", err)
		os.Exit(errors.ParsingAnswers)
	}

	authority, err := parseDNSAuthority(query[12:], header.NScount)

	if err != nil {
		log.Fatalf("Error parsing authority: %s", err)
		os.Exit(errors.ParsingAuthority)
	}

	additional, err := parseDNSAdditional(query[12:], header.NScount)

	if err != nil {
		log.Fatalf("Error parsing additional: %s", err)
		os.Exit(errors.ParsingAdditional)
	}

	message := types.DNSMessage{
		Header:     header,
		Questions:  questions,
		Answers:    answers,
		Authority:  authority,
		Additional: additional,
	}

	//messageJSON, err := json.MarshalIndent(message, "", "    ")
	//messageJSON, err := json.Marshal(message)
	log.Printf("Data parse: %+v\n", message)

	return message
}
