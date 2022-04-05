// Package pacs_028_001_03
// Do not Edit. This stuff it's been automatically generated.
package pacs_028_001_03

import (
	"bytes"
	"encoding/xml"
)

// Document type definition
type Document struct {
	XMLName         xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 Document"`
	FIToFIPmtStsReq FIToFIPaymentStatusRequestV03 `xml:"FIToFIPmtStsReq"`
}

func (d *Document) ToXML() ([]byte, error) {
	w := &bytes.Buffer{}
	w.Write([]byte(xml.Header))

	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	err := enc.Encode(d)
	if err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

func NewDocumentFromXML(b []byte) (*Document, error) {
	d := &Document{}
	err := xml.Unmarshal(b, d)
	return d, err
}

// IsValid checks if Document is valid
func (s Document) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.FIToFIPmtStsReq.IsValid(false)

	return valid
}
