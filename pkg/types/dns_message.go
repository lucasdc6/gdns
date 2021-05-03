// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package types define all the DNS types used by the server
package types

// DNSMessage - DNS message type
// Message structure from RFC 1035
//
// +---------------------+
// |       Header        |
// +---------------------+
// |       Question      | the question for the name server
// +---------------------+
// |       Answer        | RRs answering the question
// +---------------------+
// |       Authority     | RRs pointing toward an authority
// +---------------------+
// |       Additional    | RRs holding additional information
// +---------------------+
type DNSMessage struct {
	Header     DNSHeader     `json:"header"`
	Questions  []DNSQuestion `json:"questions"`
	Answers    []DNSResource `json:"answers"`
	Authority  []DNSResource `json:"authority"`
	Additional []DNSResource `json:"additional"`
}
