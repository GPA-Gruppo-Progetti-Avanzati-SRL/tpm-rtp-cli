// Package pain_014_001_07
// Do not Edit. This stuff it's been automatically generated.
// Generated at 2022-04-05 22:58:59.627106 +0200 CEST m=+0.110044876
package pain_014_001_07

import (
	"bytes"
	"encoding/xml"
)

// Document type definition
type Document struct {
	XMLName                xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.07 Document"`
	CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV07 `xml:"CdtrPmtActvtnReqStsRpt"`
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
	valid = valid && s.CdtrPmtActvtnReqStsRpt.IsValid(false)

	return valid
}