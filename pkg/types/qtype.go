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

// QType type - Query type
type QType struct {
	Name string
	Code int
}

// QType constants
var (
	A          QType = QType{Name: "A", Code: 1}
	AAAA       QType = QType{Name: "AAAA", Code: 28}
	AFSDB      QType = QType{Name: "AFSDB", Code: 18}
	APL        QType = QType{Name: "APL", Code: 42}
	CAA        QType = QType{Name: "CAA", Code: 257}
	CDNSKEY    QType = QType{Name: "CDNSKEY", Code: 60}
	CDS        QType = QType{Name: "CDS", Code: 59}
	CERT       QType = QType{Name: "CERT", Code: 37}
	CNAME      QType = QType{Name: "CNAME", Code: 5}
	DHCID      QType = QType{Name: "DHCID", Code: 49}
	DLV        QType = QType{Name: "DLV", Code: 32769}
	DNSKEY     QType = QType{Name: "DNSKEY", Code: 48}
	DS         QType = QType{Name: "DS", Code: 43}
	IPSECKEY   QType = QType{Name: "IPSECKEY", Code: 45}
	KEY        QType = QType{Name: "KEY", Code: 25}
	KX         QType = QType{Name: "KX", Code: 36}
	LOC        QType = QType{Name: "LOC", Code: 29}
	MD         QType = QType{Name: "MD", Code: 3}
	MF         QType = QType{Name: "MF", Code: 4}
	MB         QType = QType{Name: "MB", Code: 7}
	MG         QType = QType{Name: "MG", Code: 8}
	MR         QType = QType{Name: "MR", Code: 9}
	MX         QType = QType{Name: "MX", Code: 15}
	NAPTR      QType = QType{Name: "NAPTR", Code: 35}
	NS         QType = QType{Name: "NS", Code: 2}
	NSEC       QType = QType{Name: "NSEC", Code: 47}
	NSEC3      QType = QType{Name: "NSEC3", Code: 50}
	NSEC3PARAM QType = QType{Name: "NSEC3PARAM", Code: 51}
	NULL       QType = QType{Name: "NULL", Code: 10}
	PTR        QType = QType{Name: "PTR", Code: 12}
	RRSIG      QType = QType{Name: "RRSIG", Code: 46}
	RP         QType = QType{Name: "RP", Code: 17}
	SIG        QType = QType{Name: "SIG", Code: 24}
	SOA        QType = QType{Name: "SOA", Code: 6}
	SRV        QType = QType{Name: "SRV", Code: 33}
	SSHFP      QType = QType{Name: "SSHFP", Code: 44}
	TA         QType = QType{Name: "TA", Code: 32768}
	TKEY       QType = QType{Name: "TKEY", Code: 249}
	TLSA       QType = QType{Name: "TLSA", Code: 52}
	TSIG       QType = QType{Name: "TSIG", Code: 250}
	TXT        QType = QType{Name: "TXT", Code: 16}
	DNAME      QType = QType{Name: "DNAME", Code: 39}
	WKS        QType = QType{Name: "WKS", Code: 11}
	HINFO      QType = QType{Name: "HINFO", Code: 13}
	MINFO      QType = QType{Name: "MINFO", Code: 14}
	QTYPEALL   QType = QType{Name: "QTYPE_ALL", Code: 255}
	AXFR       QType = QType{Name: "AXFR", Code: 252}
	IXFR       QType = QType{Name: "IXFR", Code: 251}
	OPT        QType = QType{Name: "OPT", Code: 41}
	MAILB      QType = QType{Name: "MAILB", Code: 253}
	MAILA      QType = QType{Name: "MAILA", Code: 254}
)

// UnmarshalYAML - Function to Unmarshal to YAML
func (qtype *QType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := unmarshal(&qtype.Name)

	if err != nil {
		log.Printf("Error when unmarshal YAML QType")
	}

	return err
}

// MarshalYAML - Function to Marshal from YAML
func (qtype QType) MarshalYAML() (interface{}, error) {
	return &qtype.Name, nil
}

// UnmarshalJSON - Function to Unmarshal to JSON
func (qtype QType) UnmarshalJSON(bytes []byte) error {
	qtypeName, err := strconv.Unquote(string(bytes))

	if err != nil {
		log.Printf("Error when unmarshal JSON QType: %s", err)
	}

	qtype, err = QTypeFromString(qtypeName)

	return err
}

// MarshalJSON - Function to Marshal from JSON
func (qtype QType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(fmt.Sprintf("\"%s\"", qtype.Name))

	return buffer.Bytes(), nil
}

func (qtype QType) String() string {
	return qtype.Name
}

// QTypeFromCode - Generate an QType struct from a numeric code
// choose one between 1 and 4
func QTypeFromCode(code int) (QType, error) {
	switch code {
	case A.Code:
		return A, nil
	case AAAA.Code:
		return AAAA, nil
	case AFSDB.Code:
		return AFSDB, nil
	case APL.Code:
		return APL, nil
	case CAA.Code:
		return CAA, nil
	case CDNSKEY.Code:
		return CDNSKEY, nil
	case CDS.Code:
		return CDS, nil
	case CERT.Code:
		return CERT, nil
	case CNAME.Code:
		return CNAME, nil
	case DHCID.Code:
		return DHCID, nil
	case DLV.Code:
		return DLV, nil
	case DNSKEY.Code:
		return DNSKEY, nil
	case DS.Code:
		return DS, nil
	case IPSECKEY.Code:
		return IPSECKEY, nil
	case KEY.Code:
		return KEY, nil
	case KX.Code:
		return KX, nil
	case LOC.Code:
		return LOC, nil
	case MD.Code:
		return MD, nil
	case MF.Code:
		return MF, nil
	case MB.Code:
		return MB, nil
	case MG.Code:
		return MG, nil
	case MR.Code:
		return MR, nil
	case MX.Code:
		return MX, nil
	case NAPTR.Code:
		return NAPTR, nil
	case NS.Code:
		return NS, nil
	case NSEC.Code:
		return NSEC, nil
	case NSEC3.Code:
		return NSEC3, nil
	case NSEC3PARAM.Code:
		return NSEC3PARAM, nil
	case NULL.Code:
		return NULL, nil
	case PTR.Code:
		return PTR, nil
	case RRSIG.Code:
		return RRSIG, nil
	case RP.Code:
		return RP, nil
	case SIG.Code:
		return SIG, nil
	case SOA.Code:
		return SOA, nil
	case SRV.Code:
		return SRV, nil
	case SSHFP.Code:
		return SSHFP, nil
	case TA.Code:
		return TA, nil
	case TKEY.Code:
		return TKEY, nil
	case TLSA.Code:
		return TLSA, nil
	case TSIG.Code:
		return TSIG, nil
	case TXT.Code:
		return TXT, nil
	case DNAME.Code:
		return DNAME, nil
	case WKS.Code:
		return WKS, nil
	case HINFO.Code:
		return HINFO, nil
	case MINFO.Code:
		return MINFO, nil
	case QTYPEALL.Code:
		return QTYPEALL, nil
	case AXFR.Code:
		return AXFR, nil
	case IXFR.Code:
		return IXFR, nil
	case OPT.Code:
		return OPT, nil
	case MAILB.Code:
		return MAILB, nil
	case MAILA.Code:
		return MAILA, nil
	}

	return QType{}, fmt.Errorf("Code %d not available, choose one of 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 24, 25, 28, 29, 33, 35, 36, 37, 39, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 59, 60, 249, 250, 251, 252, 253, 254, 255, 257, 32768 or 32769", code)
}

// QTypeFromString - Generate an OpCode struct form a name
// choose one between 1 and 4
func QTypeFromString(name string) (QType, error) {
	switch name {
	case A.Name:
		return A, nil
	case AAAA.Name:
		return AAAA, nil
	case AFSDB.Name:
		return AFSDB, nil
	case APL.Name:
		return APL, nil
	case CAA.Name:
		return CAA, nil
	case CDNSKEY.Name:
		return CDNSKEY, nil
	case CDS.Name:
		return CDS, nil
	case CERT.Name:
		return CERT, nil
	case CNAME.Name:
		return CNAME, nil
	case DHCID.Name:
		return DHCID, nil
	case DLV.Name:
		return DLV, nil
	case DNSKEY.Name:
		return DNSKEY, nil
	case DS.Name:
		return DS, nil
	case IPSECKEY.Name:
		return IPSECKEY, nil
	case KEY.Name:
		return KEY, nil
	case KX.Name:
		return KX, nil
	case LOC.Name:
		return LOC, nil
	case MD.Name:
		return MD, nil
	case MF.Name:
		return MF, nil
	case MB.Name:
		return MB, nil
	case MG.Name:
		return MG, nil
	case MR.Name:
		return MR, nil
	case MX.Name:
		return MX, nil
	case NAPTR.Name:
		return NAPTR, nil
	case NS.Name:
		return NS, nil
	case NSEC.Name:
		return NSEC, nil
	case NSEC3.Name:
		return NSEC3, nil
	case NSEC3PARAM.Name:
		return NSEC3PARAM, nil
	case NULL.Name:
		return NULL, nil
	case PTR.Name:
		return PTR, nil
	case RRSIG.Name:
		return RRSIG, nil
	case RP.Name:
		return RP, nil
	case SIG.Name:
		return SIG, nil
	case SOA.Name:
		return SOA, nil
	case SRV.Name:
		return SRV, nil
	case SSHFP.Name:
		return SSHFP, nil
	case TA.Name:
		return TA, nil
	case TKEY.Name:
		return TKEY, nil
	case TLSA.Name:
		return TLSA, nil
	case TSIG.Name:
		return TSIG, nil
	case TXT.Name:
		return TXT, nil
	case DNAME.Name:
		return DNAME, nil
	case WKS.Name:
		return WKS, nil
	case HINFO.Name:
		return HINFO, nil
	case MINFO.Name:
		return MINFO, nil
	case QTYPEALL.Name:
		return QTYPEALL, nil
	case AXFR.Name:
		return AXFR, nil
	case IXFR.Name:
		return IXFR, nil
	case OPT.Name:
		return OPT, nil
	case MAILB.Name:
		return MAILB, nil
	case MAILA.Name:
		return MAILA, nil
	}

	return QType{}, fmt.Errorf("Name %s not available, choose one of \"A\", \"AAAA\", \"AFSDB\", \"APL\", \"CAA\", \"CDNSKEY\", \"CDS\", \"CERT\", \"CNAME\", \"DHCID\", \"DLV\", \"DNSKEY\", \"DS\", \"IPSECKEY\", \"KEY\", \"KX\", \"LOC\", \"MD\", \"MF\", \"MB\", \"MG\", \"MR\", \"MX\", \"NAPTR\", \"NS\", \"NSEC\", \"NSEC3\", \"NSEC3PARAM\", \"NULL\", \"PTR\", \"RRSIG\", \"RP\", \"SIG\", \"SOA\", \"SRV\", \"SSHFP\", \"TA\", \"TKEY\", \"TLSA\", \"TSIG\", \"TXT\", \"DNAME\", \"WKS\", \"HINFO\", \"MINFO\", \"QTYPE_ALL\", \"AXFR\", \"IXFR\", \"OPT\", \"MAILB\", \"MAILA\", ", name)
}
