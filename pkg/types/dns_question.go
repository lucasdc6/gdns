// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package types define all the DNS types used by the server
package types

// DNSQuestion - DNS question type
// Question structure from RFC 1035
//
//                                     1  1  1  1  1  1
//       0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
//     +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//     |                                               |
//     /                     QNAME                     /
//     /                                               /
//     +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//     |                     QTYPE                     |
//     +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//     |                     QCLASS                    |
//     +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//
// where:
//
// QNAME           a domain name represented as a sequence of labels, where
//                 each label consists of a length octet followed by that
//                 number of octets.  The domain name terminates with the
//                 zero length octet for the null label of the root.  Note
//                 that this field may be an odd number of octets; no
//                 padding is used.
//
// QTYPE           a two octet code which specifies the type of the query.
//                 The values for this field include all codes valid for a
//                 TYPE field, together with some more general codes which
//                 can match more than one type of RR.
// QCLASS          a two octet code that specifies the class of the query.
//                 For example, the QCLASS field is IN for the Internet.
type DNSQuestion struct {
	Name  string `json:"name"`
	Type  QType  `json:"type"`
	Class QClass `json:"class"`
}
