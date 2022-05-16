// Package pain_014_001_07
// Do not Edit. This stuff it's been automatically generated.
package pain_014_001_07

import (
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"reflect"
	"strings"
	"sync"
)

const (
	Iso20022MsgName = "pain.014.001.07"

	RTPStatusRejected            = "RJCT" // applies to: DS04, DS08, DS09 (negative)
	RTPStatusTechnicallyAccepted = "ACTC" // applies to: DS05, DS06
	RTPStatusAccepted            = "ACCP" // applies to: DS08, DS09 (positive)
	RTPStatusAcceptedWithChange  = "ACWC" // applies to: DS08, DS09 (positive)

	RTPStatusReasonAlreadyExpiredRTP         = "AEXR" // applies to: DS17 StatusUpdate pain.014
	RTPStatusReasonAlreadyAcceptedRTP        = "ALAC" // applies to: DS17 StatusUpdate pain.014
	RTPStatusReasonAlreadyRefusedRTP         = "ARFR" // applies to: DS17 StatusUpdate pain.014
	RTPStatusReasonAlreadyRejectedRTP        = "ARJR" // applies to: DS17 StatusUpdate pain.014
	RTPStatusReasonInitialRTPNeverReceived   = "IRNR" // applies to: DS17 StatusUpdate pain.014
	RTPStatusReasonRTPReceivedCanBeProcessed = "REPR" // applies to: DS17 StatusUpdate pain.014

	RTPRejectReasonInvalidDebtorAccountNumber            = "AC02" // applies to: DS04
	RTPRejectReasonNotAllowedCurrency                    = "AM03" // applies to: DS04, DS08, DS09
	RTPRejectReasonDuplication                           = "AM05" // applies to: DS04, DS08, DS09
	RTPRejectReasonWrongAmount                           = "AM09" // applies to: DS08, DS09
	RTPRejectReasonAttachmentsNotSupported               = "ATNS" // applies to: DS04
	RTPRejectReasonInvalidDebtorIdentificationCode       = "BE16" // applies to: DS04
	RTPRejectReasonExpiryDateTooLong                     = "EDTL" // applies to: DS04
	RTPRejectReasonExpiryDateTimeReached                 = "EDTR" // applies to: DS04
	RTPRejectReasonInvalidFileFormat                     = "FF01" // applies to: DS04
	RTPRejectReasonFraudulentOrigin                      = "FRAD" // applies to: DS04
	RTPRejectReasonIncorrectExpiryDateTime               = "IEDT" // applies to: DS08, DS09
	RTPRejectReasonNotSpecifiedReasonCustomerGenerated   = "MS02" // applies to: DS08, DS09
	RTPRejectReasonNotSpecifiedReasonAgentGenerated      = "MS03" // applies to: DS04
	RTPRejectReasonNonAgreedRTP                          = "NOAR" // applies to: DS08, DS09
	RTPRejectReasonPayerOrPayerRTPSPNotReachable         = "NRCH" // applies to: DS04
	RTPRejectReasonTypeOfPaymentInstrumentNotSupported   = "PINS" // applies to: DS04
	RTPRejectReasonRegulatoryReason                      = "RR04" // applies to: DS04, DS10b, DS13
	RTPRejectReasonRTPNotSupportedForDebtor              = "RTNS" // applies to: DS04
	RTPRejectReasonRTPServiceProviderIdentifierIncorrect = "SPII" // applies to: DS04
	RTPRejectReasonUnknownCreditor                       = "UCRD" // applies to: DS08, DS09

)

// Document type definition
type Document struct {
	XMLName                xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.07 Document"`
	CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV07 `xml:"CdtrPmtActvtnReqStsRpt"`
	mapper                 *common.Mapper                                  `xml:"-"`
}

func NewDocument() Document {
	d := Document{
		mapper: mapper(),
	}

	return d
}

func (d *Document) WithMapper() {

	if d.mapper == nil {
		d.mapper = mapper()
	}

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
func (d Document) IsValid(optional bool) bool {

	valid := true
	valid = valid && d.CdtrPmtActvtnReqStsRpt.IsValid(false)

	return valid
}

/*
 * Document reflection funcs
 */

// NameMapper is used to map column names to struct field names.  By default,
// it uses strings.ToLower to lowercase struct field names.  It can be set
// to whatever you want, but it is encouraged to be set before use
// as name-to-field mappings are cached after first use on a type.
var NameMapper = NameMapperFunc

func NameMapperFunc(s string) string {
	return strings.TrimPrefix(s, "urn:iso:std:iso:20022:tech:xsd:pain.013.001.07")
}

// Rather than creating on init, this is created when necessary so that
// importers have time to customize the NameMapper.
var mpr *common.Mapper

// mprMu protects mpr.
var mprMu sync.Mutex

// mapper returns a valid mapper using the configured NameMapper func.
func mapper() *common.Mapper {
	mprMu.Lock()
	defer mprMu.Unlock()

	if mpr == nil {
		mpr = common.NewMapperFunc("xml", NameMapper)
	}
	return mpr
}

// fieldsByName fills a values interface with fields from the passed value based
// on the traversals in int.  If ptrs is true, return addresses instead of values.
// We write this instead of using FieldsByName to save allocations and map lookups
// when iterating over many rows.  Empty traversals will get an interface pointer.
// Because of the necessity of requesting ptrs or values, it's considered a bit too
// specialized for inclusion in reflectx itself.
func fieldsByTraversal(v reflect.Value, traversals [][]int, values []interface{}, ptrs bool) error {
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		return errors.New("argument not a struct")
	}

	for i, traversal := range traversals {
		if len(traversal) == 0 {
			values[i] = new(interface{})
			continue
		}
		f := common.FieldByIndexes(v, traversal)
		if ptrs {
			values[i] = f.Addr().Interface()
		} else {
			values[i] = f.Interface()
		}
	}
	return nil
}
