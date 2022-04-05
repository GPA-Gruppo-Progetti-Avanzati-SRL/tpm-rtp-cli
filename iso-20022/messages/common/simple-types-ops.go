// Package common
// Do not Edit. This stuff it's been automatically generated.
// Generated at 2022-04-05 07:17:08.59722 +0200 CEST m=+0.095667126
package common

import (
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
	"regexp"
	"time"
)

/*
 * NamePrefix2Code Ops
 */

var NamePrefix2CodeEnumRestriction = []string{"DOCT", "MADM", "MISS", "MIST", "MIKS"}

// IsValid checks if NamePrefix2Code of type String is valid
func (t NamePrefix2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), NamePrefix2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t NamePrefix2Code) String() string {
	return string(t)
}

// ToNamePrefix2Code method for easy conversion with application of restrictions
func ToNamePrefix2Code(s string) (NamePrefix2Code, error) {
	if !applyEnumRestriction(s, NamePrefix2CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type NamePrefix2Code", s)
	}

	return NamePrefix2Code(s), nil
}

// MustToNamePrefix2Code method for easy conversion with application of restrictions. Panics on error.
func MustToNamePrefix2Code(s string) NamePrefix2Code {
	v, err := ToNamePrefix2Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalPaymentGroupStatus1Code Ops
 */

// IsValid checks if ExternalPaymentGroupStatus1Code of type String is valid
func (t ExternalPaymentGroupStatus1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalPaymentGroupStatus1Code) String() string {
	return string(t)
}

// ToExternalPaymentGroupStatus1Code method for easy conversion with application of restrictions
func ToExternalPaymentGroupStatus1Code(s string) (ExternalPaymentGroupStatus1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalPaymentGroupStatus1Code", s)
	}

	return ExternalPaymentGroupStatus1Code(s), nil
}

// MustToExternalPaymentGroupStatus1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalPaymentGroupStatus1Code(s string) ExternalPaymentGroupStatus1Code {
	v, err := ToExternalPaymentGroupStatus1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * AddressType2Code Ops
 */

var AddressType2CodeEnumRestriction = []string{"ADDR", "PBOX", "HOME", "BIZZ", "MLTO", "DLVY"}

// IsValid checks if AddressType2Code of type String is valid
func (t AddressType2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), AddressType2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t AddressType2Code) String() string {
	return string(t)
}

// ToAddressType2Code method for easy conversion with application of restrictions
func ToAddressType2Code(s string) (AddressType2Code, error) {
	if !applyEnumRestriction(s, AddressType2CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type AddressType2Code", s)
	}

	return AddressType2Code(s), nil
}

// MustToAddressType2Code method for easy conversion with application of restrictions. Panics on error.
func MustToAddressType2Code(s string) AddressType2Code {
	v, err := ToAddressType2Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ActiveOrHistoricCurrencyCode Ops
 */

var ActiveOrHistoricCurrencyCodePatternRestriction = regexp.MustCompile(`[A-Z]{3,3}`)

// IsValid checks if ActiveOrHistoricCurrencyCode of type String is valid
func (t ActiveOrHistoricCurrencyCode) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(t.String(), ActiveOrHistoricCurrencyCodePatternRestriction)

	return valid
}

// String method for easy conversion
func (t ActiveOrHistoricCurrencyCode) String() string {
	return string(t)
}

// ToActiveOrHistoricCurrencyCode method for easy conversion with application of restrictions
func ToActiveOrHistoricCurrencyCode(s string) (ActiveOrHistoricCurrencyCode, error) {
	if !applyPatternRestriction(s, ActiveOrHistoricCurrencyCodePatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type ActiveOrHistoricCurrencyCode", s)
	}

	return ActiveOrHistoricCurrencyCode(s), nil
}

// MustToActiveOrHistoricCurrencyCode method for easy conversion with application of restrictions. Panics on error.
func MustToActiveOrHistoricCurrencyCode(s string) ActiveOrHistoricCurrencyCode {
	v, err := ToActiveOrHistoricCurrencyCode(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalServiceLevel1Code Ops
 */

// IsValid checks if ExternalServiceLevel1Code of type String is valid
func (t ExternalServiceLevel1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalServiceLevel1Code) String() string {
	return string(t)
}

// ToExternalServiceLevel1Code method for easy conversion with application of restrictions
func ToExternalServiceLevel1Code(s string) (ExternalServiceLevel1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalServiceLevel1Code", s)
	}

	return ExternalServiceLevel1Code(s), nil
}

// MustToExternalServiceLevel1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalServiceLevel1Code(s string) ExternalServiceLevel1Code {
	v, err := ToExternalServiceLevel1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ActiveCurrencyCode Ops
 */

var ActiveCurrencyCodePatternRestriction = regexp.MustCompile(`[A-Z]{3,3}`)

// IsValid checks if ActiveCurrencyCode of type String is valid
func (t ActiveCurrencyCode) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(t.String(), ActiveCurrencyCodePatternRestriction)

	return valid
}

// String method for easy conversion
func (t ActiveCurrencyCode) String() string {
	return string(t)
}

// ToActiveCurrencyCode method for easy conversion with application of restrictions
func ToActiveCurrencyCode(s string) (ActiveCurrencyCode, error) {
	if !applyPatternRestriction(s, ActiveCurrencyCodePatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type ActiveCurrencyCode", s)
	}

	return ActiveCurrencyCode(s), nil
}

// MustToActiveCurrencyCode method for easy conversion with application of restrictions. Panics on error.
func MustToActiveCurrencyCode(s string) ActiveCurrencyCode {
	v, err := ToActiveCurrencyCode(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max15NumericText Ops
 */

var Max15NumericTextPatternRestriction = regexp.MustCompile(`[0-9]{1,15}`)

// IsValid checks if Max15NumericText of type String is valid
func (t Max15NumericText) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(t.String(), Max15NumericTextPatternRestriction)

	return valid
}

// String method for easy conversion
func (t Max15NumericText) String() string {
	return string(t)
}

// ToMax15NumericText method for easy conversion with application of restrictions
func ToMax15NumericText(s string) (Max15NumericText, error) {
	if !applyPatternRestriction(s, Max15NumericTextPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type Max15NumericText", s)
	}

	return Max15NumericText(s), nil
}

// MustToMax15NumericText method for easy conversion with application of restrictions. Panics on error.
func MustToMax15NumericText(s string) Max15NumericText {
	v, err := ToMax15NumericText(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalPersonIdentification1Code Ops
 */

// IsValid checks if ExternalPersonIdentification1Code of type String is valid
func (t ExternalPersonIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalPersonIdentification1Code) String() string {
	return string(t)
}

// ToExternalPersonIdentification1Code method for easy conversion with application of restrictions
func ToExternalPersonIdentification1Code(s string) (ExternalPersonIdentification1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalPersonIdentification1Code", s)
	}

	return ExternalPersonIdentification1Code(s), nil
}

// MustToExternalPersonIdentification1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalPersonIdentification1Code(s string) ExternalPersonIdentification1Code {
	v, err := ToExternalPersonIdentification1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max34Text Ops
 */

// IsValid checks if Max34Text of type String is valid
func (t Max34Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 34)

	return valid
}

// String method for easy conversion
func (t Max34Text) String() string {
	return string(t)
}

// ToMax34Text method for easy conversion with application of restrictions
func ToMax34Text(s string) (Max34Text, error) {
	if !applyLengthRestriction(s, 0, 1, 34) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max34Text", s)
	}

	return Max34Text(s), nil
}

// MustToMax34Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax34Text(s string) Max34Text {
	v, err := ToMax34Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ChequeType2Code Ops
 */

var ChequeType2CodeEnumRestriction = []string{"CCHQ", "CCCH", "BCHQ", "DRFT", "ELDR"}

// IsValid checks if ChequeType2Code of type String is valid
func (t ChequeType2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), ChequeType2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ChequeType2Code) String() string {
	return string(t)
}

// ToChequeType2Code method for easy conversion with application of restrictions
func ToChequeType2Code(s string) (ChequeType2Code, error) {
	if !applyEnumRestriction(s, ChequeType2CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type ChequeType2Code", s)
	}

	return ChequeType2Code(s), nil
}

// MustToChequeType2Code method for easy conversion with application of restrictions. Panics on error.
func MustToChequeType2Code(s string) ChequeType2Code {
	v, err := ToChequeType2Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ClearingChannel2Code Ops
 */

var ClearingChannel2CodeEnumRestriction = []string{"RTGS", "RTNS", "MPNS", "BOOK"}

// IsValid checks if ClearingChannel2Code of type String is valid
func (t ClearingChannel2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), ClearingChannel2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ClearingChannel2Code) String() string {
	return string(t)
}

// ToClearingChannel2Code method for easy conversion with application of restrictions
func ToClearingChannel2Code(s string) (ClearingChannel2Code, error) {
	if !applyEnumRestriction(s, ClearingChannel2CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type ClearingChannel2Code", s)
	}

	return ClearingChannel2Code(s), nil
}

// MustToClearingChannel2Code method for easy conversion with application of restrictions. Panics on error.
func MustToClearingChannel2Code(s string) ClearingChannel2Code {
	v, err := ToClearingChannel2Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalPurpose1Code Ops
 */

// IsValid checks if ExternalPurpose1Code of type String is valid
func (t ExternalPurpose1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalPurpose1Code) String() string {
	return string(t)
}

// ToExternalPurpose1Code method for easy conversion with application of restrictions
func ToExternalPurpose1Code(s string) (ExternalPurpose1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalPurpose1Code", s)
	}

	return ExternalPurpose1Code(s), nil
}

// MustToExternalPurpose1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalPurpose1Code(s string) ExternalPurpose1Code {
	v, err := ToExternalPurpose1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max2048Text Ops
 */

// IsValid checks if Max2048Text of type String is valid
func (t Max2048Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 2048)

	return valid
}

// String method for easy conversion
func (t Max2048Text) String() string {
	return string(t)
}

// ToMax2048Text method for easy conversion with application of restrictions
func ToMax2048Text(s string) (Max2048Text, error) {
	if !applyLengthRestriction(s, 0, 1, 2048) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max2048Text", s)
	}

	return Max2048Text(s), nil
}

// MustToMax2048Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax2048Text(s string) Max2048Text {
	v, err := ToMax2048Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * TaxRecordPeriod1Code Ops
 */

var TaxRecordPeriod1CodeEnumRestriction = []string{"MM01", "MM02", "MM03", "MM04", "MM05", "MM06", "MM07", "MM08", "MM09", "MM10", "MM11", "MM12", "QTR1", "QTR2", "QTR3", "QTR4", "HLF1", "HLF2"}

// IsValid checks if TaxRecordPeriod1Code of type String is valid
func (t TaxRecordPeriod1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), TaxRecordPeriod1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t TaxRecordPeriod1Code) String() string {
	return string(t)
}

// ToTaxRecordPeriod1Code method for easy conversion with application of restrictions
func ToTaxRecordPeriod1Code(s string) (TaxRecordPeriod1Code, error) {
	if !applyEnumRestriction(s, TaxRecordPeriod1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type TaxRecordPeriod1Code", s)
	}

	return TaxRecordPeriod1Code(s), nil
}

// MustToTaxRecordPeriod1Code method for easy conversion with application of restrictions. Panics on error.
func MustToTaxRecordPeriod1Code(s string) TaxRecordPeriod1Code {
	v, err := ToTaxRecordPeriod1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * RemittanceLocationMethod2Code Ops
 */

var RemittanceLocationMethod2CodeEnumRestriction = []string{"FAXI", "EDIC", "URID", "EMAL", "POST", "SMSM"}

// IsValid checks if RemittanceLocationMethod2Code of type String is valid
func (t RemittanceLocationMethod2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), RemittanceLocationMethod2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t RemittanceLocationMethod2Code) String() string {
	return string(t)
}

// ToRemittanceLocationMethod2Code method for easy conversion with application of restrictions
func ToRemittanceLocationMethod2Code(s string) (RemittanceLocationMethod2Code, error) {
	if !applyEnumRestriction(s, RemittanceLocationMethod2CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type RemittanceLocationMethod2Code", s)
	}

	return RemittanceLocationMethod2Code(s), nil
}

// MustToRemittanceLocationMethod2Code method for easy conversion with application of restrictions. Panics on error.
func MustToRemittanceLocationMethod2Code(s string) RemittanceLocationMethod2Code {
	v, err := ToRemittanceLocationMethod2Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * TransactionIndividualStatus1Code Ops
 */

var TransactionIndividualStatus1CodeEnumRestriction = []string{"ACTC", "RJCT", "PDNG", "ACCP", "ACSP", "ACSC", "ACCR", "ACWC"}

// IsValid checks if TransactionIndividualStatus1Code of type String is valid
func (t TransactionIndividualStatus1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), TransactionIndividualStatus1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t TransactionIndividualStatus1Code) String() string {
	return string(t)
}

// ToTransactionIndividualStatus1Code method for easy conversion with application of restrictions
func ToTransactionIndividualStatus1Code(s string) (TransactionIndividualStatus1Code, error) {
	if !applyEnumRestriction(s, TransactionIndividualStatus1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type TransactionIndividualStatus1Code", s)
	}

	return TransactionIndividualStatus1Code(s), nil
}

// MustToTransactionIndividualStatus1Code method for easy conversion with application of restrictions. Panics on error.
func MustToTransactionIndividualStatus1Code(s string) TransactionIndividualStatus1Code {
	v, err := ToTransactionIndividualStatus1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max4Text Ops
 */

// IsValid checks if Max4Text of type String is valid
func (t Max4Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t Max4Text) String() string {
	return string(t)
}

// ToMax4Text method for easy conversion with application of restrictions
func ToMax4Text(s string) (Max4Text, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max4Text", s)
	}

	return Max4Text(s), nil
}

// MustToMax4Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax4Text(s string) Max4Text {
	v, err := ToMax4Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max350Text Ops
 */

// IsValid checks if Max350Text of type String is valid
func (t Max350Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 350)

	return valid
}

// String method for easy conversion
func (t Max350Text) String() string {
	return string(t)
}

// ToMax350Text method for easy conversion with application of restrictions
func ToMax350Text(s string) (Max350Text, error) {
	if !applyLengthRestriction(s, 0, 1, 350) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max350Text", s)
	}

	return Max350Text(s), nil
}

// MustToMax350Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax350Text(s string) Max350Text {
	v, err := ToMax350Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ISODateTime Ops
 */

// IsValid checks if ISODateTime of type DateTime is valid
func (t ISODateTime) IsValid(optional bool) bool {

	valid := xsdt.DateTime(t).IsValid(optional)
	return valid
}

// String method for easy conversion
func (t ISODateTime) String() string {
	return string(t)
}

// ToISODateTime method for easy conversion from time.Time
func ToISODateTime(tm interface{}) (ISODateTime, error) {

	switch typedTm := tm.(type) {
	case time.Time:
		return ISODateTime(typedTm.Format(time.RFC3339)), nil
	}

	return "", fmt.Errorf("cannot convert %v to ISODateTime", tm)
}

func MustToISODateTime(tm interface{}) ISODateTime {
	d, err := ToISODateTime(tm)
	if err != nil {
		panic(err)
	}

	return d
}

/*
 * ExternalStatusReason1Code Ops
 */

// IsValid checks if ExternalStatusReason1Code of type String is valid
func (t ExternalStatusReason1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalStatusReason1Code) String() string {
	return string(t)
}

// ToExternalStatusReason1Code method for easy conversion with application of restrictions
func ToExternalStatusReason1Code(s string) (ExternalStatusReason1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalStatusReason1Code", s)
	}

	return ExternalStatusReason1Code(s), nil
}

// MustToExternalStatusReason1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalStatusReason1Code(s string) ExternalStatusReason1Code {
	v, err := ToExternalStatusReason1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max16Text Ops
 */

// IsValid checks if Max16Text of type String is valid
func (t Max16Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 16)

	return valid
}

// String method for easy conversion
func (t Max16Text) String() string {
	return string(t)
}

// ToMax16Text method for easy conversion with application of restrictions
func ToMax16Text(s string) (Max16Text, error) {
	if !applyLengthRestriction(s, 0, 1, 16) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max16Text", s)
	}

	return Max16Text(s), nil
}

// MustToMax16Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax16Text(s string) Max16Text {
	v, err := ToMax16Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalChargeType1Code Ops
 */

// IsValid checks if ExternalChargeType1Code of type String is valid
func (t ExternalChargeType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalChargeType1Code) String() string {
	return string(t)
}

// ToExternalChargeType1Code method for easy conversion with application of restrictions
func ToExternalChargeType1Code(s string) (ExternalChargeType1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalChargeType1Code", s)
	}

	return ExternalChargeType1Code(s), nil
}

// MustToExternalChargeType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalChargeType1Code(s string) ExternalChargeType1Code {
	v, err := ToExternalChargeType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * DocumentType6Code Ops
 */

var DocumentType6CodeEnumRestriction = []string{"MSIN", "CNFA", "DNFA", "CINV", "CREN", "DEBN", "HIRI", "SBIN", "CMCN", "SOAC", "DISP", "BOLD", "VCHR", "AROI", "TSUT", "PUOR"}

// IsValid checks if DocumentType6Code of type String is valid
func (t DocumentType6Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), DocumentType6CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t DocumentType6Code) String() string {
	return string(t)
}

// ToDocumentType6Code method for easy conversion with application of restrictions
func ToDocumentType6Code(s string) (DocumentType6Code, error) {
	if !applyEnumRestriction(s, DocumentType6CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type DocumentType6Code", s)
	}

	return DocumentType6Code(s), nil
}

// MustToDocumentType6Code method for easy conversion with application of restrictions. Panics on error.
func MustToDocumentType6Code(s string) DocumentType6Code {
	v, err := ToDocumentType6Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalTaxAmountType1Code Ops
 */

// IsValid checks if ExternalTaxAmountType1Code of type String is valid
func (t ExternalTaxAmountType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalTaxAmountType1Code) String() string {
	return string(t)
}

// ToExternalTaxAmountType1Code method for easy conversion with application of restrictions
func ToExternalTaxAmountType1Code(s string) (ExternalTaxAmountType1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalTaxAmountType1Code", s)
	}

	return ExternalTaxAmountType1Code(s), nil
}

// MustToExternalTaxAmountType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalTaxAmountType1Code(s string) ExternalTaxAmountType1Code {
	v, err := ToExternalTaxAmountType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * LEIIdentifier Ops
 */

var LEIIdentifierPatternRestriction = regexp.MustCompile(`[A-Z0-9]{18,18}[0-9]{2,2}`)

// IsValid checks if LEIIdentifier of type String is valid
func (t LEIIdentifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(t.String(), LEIIdentifierPatternRestriction)

	return valid
}

// String method for easy conversion
func (t LEIIdentifier) String() string {
	return string(t)
}

// ToLEIIdentifier method for easy conversion with application of restrictions
func ToLEIIdentifier(s string) (LEIIdentifier, error) {
	if !applyPatternRestriction(s, LEIIdentifierPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type LEIIdentifier", s)
	}

	return LEIIdentifier(s), nil
}

// MustToLEIIdentifier method for easy conversion with application of restrictions. Panics on error.
func MustToLEIIdentifier(s string) LEIIdentifier {
	v, err := ToLEIIdentifier(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max105Text Ops
 */

// IsValid checks if Max105Text of type String is valid
func (t Max105Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 105)

	return valid
}

// String method for easy conversion
func (t Max105Text) String() string {
	return string(t)
}

// ToMax105Text method for easy conversion with application of restrictions
func ToMax105Text(s string) (Max105Text, error) {
	if !applyLengthRestriction(s, 0, 1, 105) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max105Text", s)
	}

	return Max105Text(s), nil
}

// MustToMax105Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax105Text(s string) Max105Text {
	v, err := ToMax105Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max1025Text Ops
 */

// IsValid checks if Max1025Text of type String is valid
func (t Max1025Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 1025)

	return valid
}

// String method for easy conversion
func (t Max1025Text) String() string {
	return string(t)
}

// ToMax1025Text method for easy conversion with application of restrictions
func ToMax1025Text(s string) (Max1025Text, error) {
	if !applyLengthRestriction(s, 0, 1, 1025) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max1025Text", s)
	}

	return Max1025Text(s), nil
}

// MustToMax1025Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax1025Text(s string) Max1025Text {
	v, err := ToMax1025Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalPaymentModificationRejection1Code Ops
 */

// IsValid checks if ExternalPaymentModificationRejection1Code of type String is valid
func (t ExternalPaymentModificationRejection1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalPaymentModificationRejection1Code) String() string {
	return string(t)
}

// ToExternalPaymentModificationRejection1Code method for easy conversion with application of restrictions
func ToExternalPaymentModificationRejection1Code(s string) (ExternalPaymentModificationRejection1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalPaymentModificationRejection1Code", s)
	}

	return ExternalPaymentModificationRejection1Code(s), nil
}

// MustToExternalPaymentModificationRejection1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalPaymentModificationRejection1Code(s string) ExternalPaymentModificationRejection1Code {
	v, err := ToExternalPaymentModificationRejection1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max128Text Ops
 */

// IsValid checks if Max128Text of type String is valid
func (t Max128Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 128)

	return valid
}

// String method for easy conversion
func (t Max128Text) String() string {
	return string(t)
}

// ToMax128Text method for easy conversion with application of restrictions
func ToMax128Text(s string) (Max128Text, error) {
	if !applyLengthRestriction(s, 0, 1, 128) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max128Text", s)
	}

	return Max128Text(s), nil
}

// MustToMax128Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax128Text(s string) Max128Text {
	v, err := ToMax128Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * SequenceType3Code Ops
 */

var SequenceType3CodeEnumRestriction = []string{"FRST", "RCUR", "FNAL", "OOFF", "RPRE"}

// IsValid checks if SequenceType3Code of type String is valid
func (t SequenceType3Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), SequenceType3CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t SequenceType3Code) String() string {
	return string(t)
}

// ToSequenceType3Code method for easy conversion with application of restrictions
func ToSequenceType3Code(s string) (SequenceType3Code, error) {
	if !applyEnumRestriction(s, SequenceType3CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type SequenceType3Code", s)
	}

	return SequenceType3Code(s), nil
}

// MustToSequenceType3Code method for easy conversion with application of restrictions. Panics on error.
func MustToSequenceType3Code(s string) SequenceType3Code {
	v, err := ToSequenceType3Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * PaymentMethod7Code Ops
 */

var PaymentMethod7CodeEnumRestriction = []string{"CHK", "TRF"}

// IsValid checks if PaymentMethod7Code of type String is valid
func (t PaymentMethod7Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), PaymentMethod7CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t PaymentMethod7Code) String() string {
	return string(t)
}

// ToPaymentMethod7Code method for easy conversion with application of restrictions
func ToPaymentMethod7Code(s string) (PaymentMethod7Code, error) {
	if !applyEnumRestriction(s, PaymentMethod7CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type PaymentMethod7Code", s)
	}

	return PaymentMethod7Code(s), nil
}

// MustToPaymentMethod7Code method for easy conversion with application of restrictions. Panics on error.
func MustToPaymentMethod7Code(s string) PaymentMethod7Code {
	v, err := ToPaymentMethod7Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max10Text Ops
 */

// IsValid checks if Max10Text of type String is valid
func (t Max10Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 10)

	return valid
}

// String method for easy conversion
func (t Max10Text) String() string {
	return string(t)
}

// ToMax10Text method for easy conversion with application of restrictions
func ToMax10Text(s string) (Max10Text, error) {
	if !applyLengthRestriction(s, 0, 1, 10) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max10Text", s)
	}

	return Max10Text(s), nil
}

// MustToMax10Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax10Text(s string) Max10Text {
	v, err := ToMax10Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * CountryCode Ops
 */

var CountryCodePatternRestriction = regexp.MustCompile(`[A-Z]{2,2}`)

// IsValid checks if CountryCode of type String is valid
func (t CountryCode) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(t.String(), CountryCodePatternRestriction)

	return valid
}

// String method for easy conversion
func (t CountryCode) String() string {
	return string(t)
}

// ToCountryCode method for easy conversion with application of restrictions
func ToCountryCode(s string) (CountryCode, error) {
	if !applyPatternRestriction(s, CountryCodePatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type CountryCode", s)
	}

	return CountryCode(s), nil
}

// MustToCountryCode method for easy conversion with application of restrictions. Panics on error.
func MustToCountryCode(s string) CountryCode {
	v, err := ToCountryCode(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalCategoryPurpose1Code Ops
 */

// IsValid checks if ExternalCategoryPurpose1Code of type String is valid
func (t ExternalCategoryPurpose1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalCategoryPurpose1Code) String() string {
	return string(t)
}

// ToExternalCategoryPurpose1Code method for easy conversion with application of restrictions
func ToExternalCategoryPurpose1Code(s string) (ExternalCategoryPurpose1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalCategoryPurpose1Code", s)
	}

	return ExternalCategoryPurpose1Code(s), nil
}

// MustToExternalCategoryPurpose1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalCategoryPurpose1Code(s string) ExternalCategoryPurpose1Code {
	v, err := ToExternalCategoryPurpose1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * PaymentMethod4Code Ops
 */

var PaymentMethod4CodeEnumRestriction = []string{"CHK", "TRF", "DD", "TRA"}

// IsValid checks if PaymentMethod4Code of type String is valid
func (t PaymentMethod4Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), PaymentMethod4CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t PaymentMethod4Code) String() string {
	return string(t)
}

// ToPaymentMethod4Code method for easy conversion with application of restrictions
func ToPaymentMethod4Code(s string) (PaymentMethod4Code, error) {
	if !applyEnumRestriction(s, PaymentMethod4CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type PaymentMethod4Code", s)
	}

	return PaymentMethod4Code(s), nil
}

// MustToPaymentMethod4Code method for easy conversion with application of restrictions. Panics on error.
func MustToPaymentMethod4Code(s string) PaymentMethod4Code {
	v, err := ToPaymentMethod4Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Priority2Code Ops
 */

var Priority2CodeEnumRestriction = []string{"HIGH", "NORM"}

// IsValid checks if Priority2Code of type String is valid
func (t Priority2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), Priority2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t Priority2Code) String() string {
	return string(t)
}

// ToPriority2Code method for easy conversion with application of restrictions
func ToPriority2Code(s string) (Priority2Code, error) {
	if !applyEnumRestriction(s, Priority2CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type Priority2Code", s)
	}

	return Priority2Code(s), nil
}

// MustToPriority2Code method for easy conversion with application of restrictions. Panics on error.
func MustToPriority2Code(s string) Priority2Code {
	v, err := ToPriority2Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalLocalInstrument1Code Ops
 */

// IsValid checks if ExternalLocalInstrument1Code of type String is valid
func (t ExternalLocalInstrument1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 35)

	return valid
}

// String method for easy conversion
func (t ExternalLocalInstrument1Code) String() string {
	return string(t)
}

// ToExternalLocalInstrument1Code method for easy conversion with application of restrictions
func ToExternalLocalInstrument1Code(s string) (ExternalLocalInstrument1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 35) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalLocalInstrument1Code", s)
	}

	return ExternalLocalInstrument1Code(s), nil
}

// MustToExternalLocalInstrument1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalLocalInstrument1Code(s string) ExternalLocalInstrument1Code {
	v, err := ToExternalLocalInstrument1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * AnyBICDec2014Identifier Ops
 */

var AnyBICDec2014IdentifierPatternRestriction = regexp.MustCompile(`[A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}`)

// IsValid checks if AnyBICDec2014Identifier of type String is valid
func (t AnyBICDec2014Identifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(t.String(), AnyBICDec2014IdentifierPatternRestriction)

	return valid
}

// String method for easy conversion
func (t AnyBICDec2014Identifier) String() string {
	return string(t)
}

// ToAnyBICDec2014Identifier method for easy conversion with application of restrictions
func ToAnyBICDec2014Identifier(s string) (AnyBICDec2014Identifier, error) {
	if !applyPatternRestriction(s, AnyBICDec2014IdentifierPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type AnyBICDec2014Identifier", s)
	}

	return AnyBICDec2014Identifier(s), nil
}

// MustToAnyBICDec2014Identifier method for easy conversion with application of restrictions. Panics on error.
func MustToAnyBICDec2014Identifier(s string) AnyBICDec2014Identifier {
	v, err := ToAnyBICDec2014Identifier(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Instruction3Code Ops
 */

var Instruction3CodeEnumRestriction = []string{"CHQB", "HOLD", "PHOB", "TELB"}

// IsValid checks if Instruction3Code of type String is valid
func (t Instruction3Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), Instruction3CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t Instruction3Code) String() string {
	return string(t)
}

// ToInstruction3Code method for easy conversion with application of restrictions
func ToInstruction3Code(s string) (Instruction3Code, error) {
	if !applyEnumRestriction(s, Instruction3CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type Instruction3Code", s)
	}

	return Instruction3Code(s), nil
}

// MustToInstruction3Code method for easy conversion with application of restrictions. Panics on error.
func MustToInstruction3Code(s string) Instruction3Code {
	v, err := ToInstruction3Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * IBAN2007Identifier Ops
 */

var IBAN2007IdentifierPatternRestriction = regexp.MustCompile(`[A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}`)

// IsValid checks if IBAN2007Identifier of type String is valid
func (t IBAN2007Identifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(t.String(), IBAN2007IdentifierPatternRestriction)

	return valid
}

// String method for easy conversion
func (t IBAN2007Identifier) String() string {
	return string(t)
}

// ToIBAN2007Identifier method for easy conversion with application of restrictions
func ToIBAN2007Identifier(s string) (IBAN2007Identifier, error) {
	if !applyPatternRestriction(s, IBAN2007IdentifierPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type IBAN2007Identifier", s)
	}

	return IBAN2007Identifier(s), nil
}

// MustToIBAN2007Identifier method for easy conversion with application of restrictions. Panics on error.
func MustToIBAN2007Identifier(s string) IBAN2007Identifier {
	v, err := ToIBAN2007Identifier(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ISODate Ops
 */

// IsValid checks if ISODate of type Date is valid
func (t ISODate) IsValid(optional bool) bool {

	valid := xsdt.Date(t).IsValid(optional)
	return valid
}

// String method for easy conversion
func (t ISODate) String() string {
	return string(t)
}

// ToISODate method for easy conversion from time.Time
func ToISODate(tm interface{}) (ISODate, error) {

	switch typedTm := tm.(type) {
	case time.Time:
		return ISODate(typedTm.Format("2006-01-02")), nil
	}

	return "", fmt.Errorf("cannot convert %v to ISODate", tm)
}

func MustToISODate(tm interface{}) ISODate {
	d, err := ToISODate(tm)
	if err != nil {
		panic(err)
	}

	return d
}

/*
 * ExternalDocumentFormat1Code Ops
 */

// IsValid checks if ExternalDocumentFormat1Code of type String is valid
func (t ExternalDocumentFormat1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalDocumentFormat1Code) String() string {
	return string(t)
}

// ToExternalDocumentFormat1Code method for easy conversion with application of restrictions
func ToExternalDocumentFormat1Code(s string) (ExternalDocumentFormat1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalDocumentFormat1Code", s)
	}

	return ExternalDocumentFormat1Code(s), nil
}

// MustToExternalDocumentFormat1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalDocumentFormat1Code(s string) ExternalDocumentFormat1Code {
	v, err := ToExternalDocumentFormat1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Frequency6Code Ops
 */

var Frequency6CodeEnumRestriction = []string{"YEAR", "MNTH", "QURT", "MIAN", "WEEK", "DAIL", "ADHO", "INDA", "FRTN"}

// IsValid checks if Frequency6Code of type String is valid
func (t Frequency6Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), Frequency6CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t Frequency6Code) String() string {
	return string(t)
}

// ToFrequency6Code method for easy conversion with application of restrictions
func ToFrequency6Code(s string) (Frequency6Code, error) {
	if !applyEnumRestriction(s, Frequency6CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type Frequency6Code", s)
	}

	return Frequency6Code(s), nil
}

// MustToFrequency6Code method for easy conversion with application of restrictions. Panics on error.
func MustToFrequency6Code(s string) Frequency6Code {
	v, err := ToFrequency6Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalDocumentType1Code Ops
 */

// IsValid checks if ExternalDocumentType1Code of type String is valid
func (t ExternalDocumentType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalDocumentType1Code) String() string {
	return string(t)
}

// ToExternalDocumentType1Code method for easy conversion with application of restrictions
func ToExternalDocumentType1Code(s string) (ExternalDocumentType1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalDocumentType1Code", s)
	}

	return ExternalDocumentType1Code(s), nil
}

// MustToExternalDocumentType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalDocumentType1Code(s string) ExternalDocumentType1Code {
	v, err := ToExternalDocumentType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Exact2NumericText Ops
 */

var Exact2NumericTextPatternRestriction = regexp.MustCompile(`[0-9]{2}`)

// IsValid checks if Exact2NumericText of type String is valid
func (t Exact2NumericText) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(t.String(), Exact2NumericTextPatternRestriction)

	return valid
}

// String method for easy conversion
func (t Exact2NumericText) String() string {
	return string(t)
}

// ToExact2NumericText method for easy conversion with application of restrictions
func ToExact2NumericText(s string) (Exact2NumericText, error) {
	if !applyPatternRestriction(s, Exact2NumericTextPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type Exact2NumericText", s)
	}

	return Exact2NumericText(s), nil
}

// MustToExact2NumericText method for easy conversion with application of restrictions. Panics on error.
func MustToExact2NumericText(s string) Exact2NumericText {
	v, err := ToExact2NumericText(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalDiscountAmountType1Code Ops
 */

// IsValid checks if ExternalDiscountAmountType1Code of type String is valid
func (t ExternalDiscountAmountType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalDiscountAmountType1Code) String() string {
	return string(t)
}

// ToExternalDiscountAmountType1Code method for easy conversion with application of restrictions
func ToExternalDiscountAmountType1Code(s string) (ExternalDiscountAmountType1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalDiscountAmountType1Code", s)
	}

	return ExternalDiscountAmountType1Code(s), nil
}

// MustToExternalDiscountAmountType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalDiscountAmountType1Code(s string) ExternalDiscountAmountType1Code {
	v, err := ToExternalDiscountAmountType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalCashClearingSystem1Code Ops
 */

// IsValid checks if ExternalCashClearingSystem1Code of type String is valid
func (t ExternalCashClearingSystem1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 3)

	return valid
}

// String method for easy conversion
func (t ExternalCashClearingSystem1Code) String() string {
	return string(t)
}

// ToExternalCashClearingSystem1Code method for easy conversion with application of restrictions
func ToExternalCashClearingSystem1Code(s string) (ExternalCashClearingSystem1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 3) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalCashClearingSystem1Code", s)
	}

	return ExternalCashClearingSystem1Code(s), nil
}

// MustToExternalCashClearingSystem1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalCashClearingSystem1Code(s string) ExternalCashClearingSystem1Code {
	v, err := ToExternalCashClearingSystem1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalMandateSetupReason1Code Ops
 */

// IsValid checks if ExternalMandateSetupReason1Code of type String is valid
func (t ExternalMandateSetupReason1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalMandateSetupReason1Code) String() string {
	return string(t)
}

// ToExternalMandateSetupReason1Code method for easy conversion with application of restrictions
func ToExternalMandateSetupReason1Code(s string) (ExternalMandateSetupReason1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalMandateSetupReason1Code", s)
	}

	return ExternalMandateSetupReason1Code(s), nil
}

// MustToExternalMandateSetupReason1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalMandateSetupReason1Code(s string) ExternalMandateSetupReason1Code {
	v, err := ToExternalMandateSetupReason1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalDocumentLineType1Code Ops
 */

// IsValid checks if ExternalDocumentLineType1Code of type String is valid
func (t ExternalDocumentLineType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalDocumentLineType1Code) String() string {
	return string(t)
}

// ToExternalDocumentLineType1Code method for easy conversion with application of restrictions
func ToExternalDocumentLineType1Code(s string) (ExternalDocumentLineType1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalDocumentLineType1Code", s)
	}

	return ExternalDocumentLineType1Code(s), nil
}

// MustToExternalDocumentLineType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalDocumentLineType1Code(s string) ExternalDocumentLineType1Code {
	v, err := ToExternalDocumentLineType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * BICFIDec2014Identifier Ops
 */

var BICFIDec2014IdentifierPatternRestriction = regexp.MustCompile(`[A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}`)

// IsValid checks if BICFIDec2014Identifier of type String is valid
func (t BICFIDec2014Identifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(t.String(), BICFIDec2014IdentifierPatternRestriction)

	return valid
}

// String method for easy conversion
func (t BICFIDec2014Identifier) String() string {
	return string(t)
}

// ToBICFIDec2014Identifier method for easy conversion with application of restrictions
func ToBICFIDec2014Identifier(s string) (BICFIDec2014Identifier, error) {
	if !applyPatternRestriction(s, BICFIDec2014IdentifierPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type BICFIDec2014Identifier", s)
	}

	return BICFIDec2014Identifier(s), nil
}

// MustToBICFIDec2014Identifier method for easy conversion with application of restrictions. Panics on error.
func MustToBICFIDec2014Identifier(s string) BICFIDec2014Identifier {
	v, err := ToBICFIDec2014Identifier(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalAccountIdentification1Code Ops
 */

// IsValid checks if ExternalAccountIdentification1Code of type String is valid
func (t ExternalAccountIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalAccountIdentification1Code) String() string {
	return string(t)
}

// ToExternalAccountIdentification1Code method for easy conversion with application of restrictions
func ToExternalAccountIdentification1Code(s string) (ExternalAccountIdentification1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalAccountIdentification1Code", s)
	}

	return ExternalAccountIdentification1Code(s), nil
}

// MustToExternalAccountIdentification1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalAccountIdentification1Code(s string) ExternalAccountIdentification1Code {
	v, err := ToExternalAccountIdentification1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalFinancialInstitutionIdentification1Code Ops
 */

// IsValid checks if ExternalFinancialInstitutionIdentification1Code of type String is valid
func (t ExternalFinancialInstitutionIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalFinancialInstitutionIdentification1Code) String() string {
	return string(t)
}

// ToExternalFinancialInstitutionIdentification1Code method for easy conversion with application of restrictions
func ToExternalFinancialInstitutionIdentification1Code(s string) (ExternalFinancialInstitutionIdentification1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalFinancialInstitutionIdentification1Code", s)
	}

	return ExternalFinancialInstitutionIdentification1Code(s), nil
}

// MustToExternalFinancialInstitutionIdentification1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalFinancialInstitutionIdentification1Code(s string) ExternalFinancialInstitutionIdentification1Code {
	v, err := ToExternalFinancialInstitutionIdentification1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max35Text Ops
 */

// IsValid checks if Max35Text of type String is valid
func (t Max35Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 35)

	return valid
}

// String method for easy conversion
func (t Max35Text) String() string {
	return string(t)
}

// ToMax35Text method for easy conversion with application of restrictions
func ToMax35Text(s string) (Max35Text, error) {
	if !applyLengthRestriction(s, 0, 1, 35) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max35Text", s)
	}

	return Max35Text(s), nil
}

// MustToMax35Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax35Text(s string) Max35Text {
	v, err := ToMax35Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalInvestigationExecutionConfirmation1Code Ops
 */

// IsValid checks if ExternalInvestigationExecutionConfirmation1Code of type String is valid
func (t ExternalInvestigationExecutionConfirmation1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalInvestigationExecutionConfirmation1Code) String() string {
	return string(t)
}

// ToExternalInvestigationExecutionConfirmation1Code method for easy conversion with application of restrictions
func ToExternalInvestigationExecutionConfirmation1Code(s string) (ExternalInvestigationExecutionConfirmation1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalInvestigationExecutionConfirmation1Code", s)
	}

	return ExternalInvestigationExecutionConfirmation1Code(s), nil
}

// MustToExternalInvestigationExecutionConfirmation1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalInvestigationExecutionConfirmation1Code(s string) ExternalInvestigationExecutionConfirmation1Code {
	v, err := ToExternalInvestigationExecutionConfirmation1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ChequeDelivery1Code Ops
 */

var ChequeDelivery1CodeEnumRestriction = []string{"MLDB", "MLCD", "MLFA", "CRDB", "CRCD", "CRFA", "PUDB", "PUCD", "PUFA", "RGDB", "RGCD", "RGFA"}

// IsValid checks if ChequeDelivery1Code of type String is valid
func (t ChequeDelivery1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), ChequeDelivery1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ChequeDelivery1Code) String() string {
	return string(t)
}

// ToChequeDelivery1Code method for easy conversion with application of restrictions
func ToChequeDelivery1Code(s string) (ChequeDelivery1Code, error) {
	if !applyEnumRestriction(s, ChequeDelivery1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type ChequeDelivery1Code", s)
	}

	return ChequeDelivery1Code(s), nil
}

// MustToChequeDelivery1Code method for easy conversion with application of restrictions. Panics on error.
func MustToChequeDelivery1Code(s string) ChequeDelivery1Code {
	v, err := ToChequeDelivery1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Exact4AlphaNumericText Ops
 */

var Exact4AlphaNumericTextPatternRestriction = regexp.MustCompile(`[a-zA-Z0-9]{4}`)

// IsValid checks if Exact4AlphaNumericText of type String is valid
func (t Exact4AlphaNumericText) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(t.String(), Exact4AlphaNumericTextPatternRestriction)

	return valid
}

// String method for easy conversion
func (t Exact4AlphaNumericText) String() string {
	return string(t)
}

// ToExact4AlphaNumericText method for easy conversion with application of restrictions
func ToExact4AlphaNumericText(s string) (Exact4AlphaNumericText, error) {
	if !applyPatternRestriction(s, Exact4AlphaNumericTextPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type Exact4AlphaNumericText", s)
	}

	return Exact4AlphaNumericText(s), nil
}

// MustToExact4AlphaNumericText method for easy conversion with application of restrictions. Panics on error.
func MustToExact4AlphaNumericText(s string) Exact4AlphaNumericText {
	v, err := ToExact4AlphaNumericText(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalCancellationReason1Code Ops
 */

// IsValid checks if ExternalCancellationReason1Code of type String is valid
func (t ExternalCancellationReason1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalCancellationReason1Code) String() string {
	return string(t)
}

// ToExternalCancellationReason1Code method for easy conversion with application of restrictions
func ToExternalCancellationReason1Code(s string) (ExternalCancellationReason1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalCancellationReason1Code", s)
	}

	return ExternalCancellationReason1Code(s), nil
}

// MustToExternalCancellationReason1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalCancellationReason1Code(s string) ExternalCancellationReason1Code {
	v, err := ToExternalCancellationReason1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalProxyAccountType1Code Ops
 */

// IsValid checks if ExternalProxyAccountType1Code of type String is valid
func (t ExternalProxyAccountType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalProxyAccountType1Code) String() string {
	return string(t)
}

// ToExternalProxyAccountType1Code method for easy conversion with application of restrictions
func ToExternalProxyAccountType1Code(s string) (ExternalProxyAccountType1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalProxyAccountType1Code", s)
	}

	return ExternalProxyAccountType1Code(s), nil
}

// MustToExternalProxyAccountType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalProxyAccountType1Code(s string) ExternalProxyAccountType1Code {
	v, err := ToExternalProxyAccountType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * RegulatoryReportingType1Code Ops
 */

var RegulatoryReportingType1CodeEnumRestriction = []string{"CRED", "DEBT", "BOTH"}

// IsValid checks if RegulatoryReportingType1Code of type String is valid
func (t RegulatoryReportingType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), RegulatoryReportingType1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t RegulatoryReportingType1Code) String() string {
	return string(t)
}

// ToRegulatoryReportingType1Code method for easy conversion with application of restrictions
func ToRegulatoryReportingType1Code(s string) (RegulatoryReportingType1Code, error) {
	if !applyEnumRestriction(s, RegulatoryReportingType1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type RegulatoryReportingType1Code", s)
	}

	return RegulatoryReportingType1Code(s), nil
}

// MustToRegulatoryReportingType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToRegulatoryReportingType1Code(s string) RegulatoryReportingType1Code {
	v, err := ToRegulatoryReportingType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max140Text Ops
 */

// IsValid checks if Max140Text of type String is valid
func (t Max140Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 140)

	return valid
}

// String method for easy conversion
func (t Max140Text) String() string {
	return string(t)
}

// ToMax140Text method for easy conversion with application of restrictions
func ToMax140Text(s string) (Max140Text, error) {
	if !applyLengthRestriction(s, 0, 1, 140) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max140Text", s)
	}

	return Max140Text(s), nil
}

// MustToMax140Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax140Text(s string) Max140Text {
	v, err := ToMax140Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalGarnishmentType1Code Ops
 */

// IsValid checks if ExternalGarnishmentType1Code of type String is valid
func (t ExternalGarnishmentType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalGarnishmentType1Code) String() string {
	return string(t)
}

// ToExternalGarnishmentType1Code method for easy conversion with application of restrictions
func ToExternalGarnishmentType1Code(s string) (ExternalGarnishmentType1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalGarnishmentType1Code", s)
	}

	return ExternalGarnishmentType1Code(s), nil
}

// MustToExternalGarnishmentType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalGarnishmentType1Code(s string) ExternalGarnishmentType1Code {
	v, err := ToExternalGarnishmentType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * UUIDv4Identifier Ops
 */

var UUIDv4IdentifierPatternRestriction = regexp.MustCompile(`[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}`)

// IsValid checks if UUIDv4Identifier of type String is valid
func (t UUIDv4Identifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(t.String(), UUIDv4IdentifierPatternRestriction)

	return valid
}

// String method for easy conversion
func (t UUIDv4Identifier) String() string {
	return string(t)
}

// ToUUIDv4Identifier method for easy conversion with application of restrictions
func ToUUIDv4Identifier(s string) (UUIDv4Identifier, error) {
	if !applyPatternRestriction(s, UUIDv4IdentifierPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type UUIDv4Identifier", s)
	}

	return UUIDv4Identifier(s), nil
}

// MustToUUIDv4Identifier method for easy conversion with application of restrictions. Panics on error.
func MustToUUIDv4Identifier(s string) UUIDv4Identifier {
	v, err := ToUUIDv4Identifier(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * PhoneNumber Ops
 */

var PhoneNumberPatternRestriction = regexp.MustCompile(`\+[0-9]{1,3}-[0-9()+\-]{1,30}`)

// IsValid checks if PhoneNumber of type String is valid
func (t PhoneNumber) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(t.String(), PhoneNumberPatternRestriction)

	return valid
}

// String method for easy conversion
func (t PhoneNumber) String() string {
	return string(t)
}

// ToPhoneNumber method for easy conversion with application of restrictions
func ToPhoneNumber(s string) (PhoneNumber, error) {
	if !applyPatternRestriction(s, PhoneNumberPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type PhoneNumber", s)
	}

	return PhoneNumber(s), nil
}

// MustToPhoneNumber method for easy conversion with application of restrictions. Panics on error.
func MustToPhoneNumber(s string) PhoneNumber {
	v, err := ToPhoneNumber(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * PreferredContactMethod1Code Ops
 */

var PreferredContactMethod1CodeEnumRestriction = []string{"LETT", "MAIL", "PHON", "FAXX", "CELL"}

// IsValid checks if PreferredContactMethod1Code of type String is valid
func (t PreferredContactMethod1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), PreferredContactMethod1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t PreferredContactMethod1Code) String() string {
	return string(t)
}

// ToPreferredContactMethod1Code method for easy conversion with application of restrictions
func ToPreferredContactMethod1Code(s string) (PreferredContactMethod1Code, error) {
	if !applyEnumRestriction(s, PreferredContactMethod1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type PreferredContactMethod1Code", s)
	}

	return PreferredContactMethod1Code(s), nil
}

// MustToPreferredContactMethod1Code method for easy conversion with application of restrictions. Panics on error.
func MustToPreferredContactMethod1Code(s string) PreferredContactMethod1Code {
	v, err := ToPreferredContactMethod1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * GroupCancellationStatus1Code Ops
 */

var GroupCancellationStatus1CodeEnumRestriction = []string{"PACR", "RJCR", "ACCR", "PDCR"}

// IsValid checks if GroupCancellationStatus1Code of type String is valid
func (t GroupCancellationStatus1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), GroupCancellationStatus1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t GroupCancellationStatus1Code) String() string {
	return string(t)
}

// ToGroupCancellationStatus1Code method for easy conversion with application of restrictions
func ToGroupCancellationStatus1Code(s string) (GroupCancellationStatus1Code, error) {
	if !applyEnumRestriction(s, GroupCancellationStatus1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type GroupCancellationStatus1Code", s)
	}

	return GroupCancellationStatus1Code(s), nil
}

// MustToGroupCancellationStatus1Code method for easy conversion with application of restrictions. Panics on error.
func MustToGroupCancellationStatus1Code(s string) GroupCancellationStatus1Code {
	v, err := ToGroupCancellationStatus1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * CreditDebitCode Ops
 */

var CreditDebitCodeEnumRestriction = []string{"CRDT", "DBIT"}

// IsValid checks if CreditDebitCode of type String is valid
func (t CreditDebitCode) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), CreditDebitCodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t CreditDebitCode) String() string {
	return string(t)
}

// ToCreditDebitCode method for easy conversion with application of restrictions
func ToCreditDebitCode(s string) (CreditDebitCode, error) {
	if !applyEnumRestriction(s, CreditDebitCodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type CreditDebitCode", s)
	}

	return CreditDebitCode(s), nil
}

// MustToCreditDebitCode method for easy conversion with application of restrictions. Panics on error.
func MustToCreditDebitCode(s string) CreditDebitCode {
	v, err := ToCreditDebitCode(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalCashAccountType1Code Ops
 */

// IsValid checks if ExternalCashAccountType1Code of type String is valid
func (t ExternalCashAccountType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalCashAccountType1Code) String() string {
	return string(t)
}

// ToExternalCashAccountType1Code method for easy conversion with application of restrictions
func ToExternalCashAccountType1Code(s string) (ExternalCashAccountType1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalCashAccountType1Code", s)
	}

	return ExternalCashAccountType1Code(s), nil
}

// MustToExternalCashAccountType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalCashAccountType1Code(s string) ExternalCashAccountType1Code {
	v, err := ToExternalCashAccountType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * DocumentType3Code Ops
 */

var DocumentType3CodeEnumRestriction = []string{"RADM", "RPIN", "FXDR", "DISP", "PUOR", "SCOR"}

// IsValid checks if DocumentType3Code of type String is valid
func (t DocumentType3Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), DocumentType3CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t DocumentType3Code) String() string {
	return string(t)
}

// ToDocumentType3Code method for easy conversion with application of restrictions
func ToDocumentType3Code(s string) (DocumentType3Code, error) {
	if !applyEnumRestriction(s, DocumentType3CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type DocumentType3Code", s)
	}

	return DocumentType3Code(s), nil
}

// MustToDocumentType3Code method for easy conversion with application of restrictions. Panics on error.
func MustToDocumentType3Code(s string) DocumentType3Code {
	v, err := ToDocumentType3Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalOrganisationIdentification1Code Ops
 */

// IsValid checks if ExternalOrganisationIdentification1Code of type String is valid
func (t ExternalOrganisationIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalOrganisationIdentification1Code) String() string {
	return string(t)
}

// ToExternalOrganisationIdentification1Code method for easy conversion with application of restrictions
func ToExternalOrganisationIdentification1Code(s string) (ExternalOrganisationIdentification1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalOrganisationIdentification1Code", s)
	}

	return ExternalOrganisationIdentification1Code(s), nil
}

// MustToExternalOrganisationIdentification1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalOrganisationIdentification1Code(s string) ExternalOrganisationIdentification1Code {
	v, err := ToExternalOrganisationIdentification1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalPaymentCompensationReason1Code Ops
 */

// IsValid checks if ExternalPaymentCompensationReason1Code of type String is valid
func (t ExternalPaymentCompensationReason1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalPaymentCompensationReason1Code) String() string {
	return string(t)
}

// ToExternalPaymentCompensationReason1Code method for easy conversion with application of restrictions
func ToExternalPaymentCompensationReason1Code(s string) (ExternalPaymentCompensationReason1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalPaymentCompensationReason1Code", s)
	}

	return ExternalPaymentCompensationReason1Code(s), nil
}

// MustToExternalPaymentCompensationReason1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalPaymentCompensationReason1Code(s string) ExternalPaymentCompensationReason1Code {
	v, err := ToExternalPaymentCompensationReason1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalPaymentTransactionStatus1Code Ops
 */

// IsValid checks if ExternalPaymentTransactionStatus1Code of type String is valid
func (t ExternalPaymentTransactionStatus1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalPaymentTransactionStatus1Code) String() string {
	return string(t)
}

// ToExternalPaymentTransactionStatus1Code method for easy conversion with application of restrictions
func ToExternalPaymentTransactionStatus1Code(s string) (ExternalPaymentTransactionStatus1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalPaymentTransactionStatus1Code", s)
	}

	return ExternalPaymentTransactionStatus1Code(s), nil
}

// MustToExternalPaymentTransactionStatus1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalPaymentTransactionStatus1Code(s string) ExternalPaymentTransactionStatus1Code {
	v, err := ToExternalPaymentTransactionStatus1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max70Text Ops
 */

// IsValid checks if Max70Text of type String is valid
func (t Max70Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 70)

	return valid
}

// String method for easy conversion
func (t Max70Text) String() string {
	return string(t)
}

// ToMax70Text method for easy conversion with application of restrictions
func ToMax70Text(s string) (Max70Text, error) {
	if !applyLengthRestriction(s, 0, 1, 70) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max70Text", s)
	}

	return Max70Text(s), nil
}

// MustToMax70Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax70Text(s string) Max70Text {
	v, err := ToMax70Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalClearingSystemIdentification1Code Ops
 */

// IsValid checks if ExternalClearingSystemIdentification1Code of type String is valid
func (t ExternalClearingSystemIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 5)

	return valid
}

// String method for easy conversion
func (t ExternalClearingSystemIdentification1Code) String() string {
	return string(t)
}

// ToExternalClearingSystemIdentification1Code method for easy conversion with application of restrictions
func ToExternalClearingSystemIdentification1Code(s string) (ExternalClearingSystemIdentification1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 5) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalClearingSystemIdentification1Code", s)
	}

	return ExternalClearingSystemIdentification1Code(s), nil
}

// MustToExternalClearingSystemIdentification1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalClearingSystemIdentification1Code(s string) ExternalClearingSystemIdentification1Code {
	v, err := ToExternalClearingSystemIdentification1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * CancellationIndividualStatus1Code Ops
 */

var CancellationIndividualStatus1CodeEnumRestriction = []string{"RJCR", "ACCR", "PDCR"}

// IsValid checks if CancellationIndividualStatus1Code of type String is valid
func (t CancellationIndividualStatus1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), CancellationIndividualStatus1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t CancellationIndividualStatus1Code) String() string {
	return string(t)
}

// ToCancellationIndividualStatus1Code method for easy conversion with application of restrictions
func ToCancellationIndividualStatus1Code(s string) (CancellationIndividualStatus1Code, error) {
	if !applyEnumRestriction(s, CancellationIndividualStatus1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type CancellationIndividualStatus1Code", s)
	}

	return CancellationIndividualStatus1Code(s), nil
}

// MustToCancellationIndividualStatus1Code method for easy conversion with application of restrictions. Panics on error.
func MustToCancellationIndividualStatus1Code(s string) CancellationIndividualStatus1Code {
	v, err := ToCancellationIndividualStatus1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * SettlementMethod1Code Ops
 */

var SettlementMethod1CodeEnumRestriction = []string{"INDA", "INGA", "COVE", "CLRG"}

// IsValid checks if SettlementMethod1Code of type String is valid
func (t SettlementMethod1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), SettlementMethod1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t SettlementMethod1Code) String() string {
	return string(t)
}

// ToSettlementMethod1Code method for easy conversion with application of restrictions
func ToSettlementMethod1Code(s string) (SettlementMethod1Code, error) {
	if !applyEnumRestriction(s, SettlementMethod1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type SettlementMethod1Code", s)
	}

	return SettlementMethod1Code(s), nil
}

// MustToSettlementMethod1Code method for easy conversion with application of restrictions. Panics on error.
func MustToSettlementMethod1Code(s string) SettlementMethod1Code {
	v, err := ToSettlementMethod1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalClaimNonReceiptRejection1Code Ops
 */

// IsValid checks if ExternalClaimNonReceiptRejection1Code of type String is valid
func (t ExternalClaimNonReceiptRejection1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalClaimNonReceiptRejection1Code) String() string {
	return string(t)
}

// ToExternalClaimNonReceiptRejection1Code method for easy conversion with application of restrictions
func ToExternalClaimNonReceiptRejection1Code(s string) (ExternalClaimNonReceiptRejection1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalClaimNonReceiptRejection1Code", s)
	}

	return ExternalClaimNonReceiptRejection1Code(s), nil
}

// MustToExternalClaimNonReceiptRejection1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalClaimNonReceiptRejection1Code(s string) ExternalClaimNonReceiptRejection1Code {
	v, err := ToExternalClaimNonReceiptRejection1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalPaymentCancellationRejection1Code Ops
 */

// IsValid checks if ExternalPaymentCancellationRejection1Code of type String is valid
func (t ExternalPaymentCancellationRejection1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(t.String(), 0, 1, 4)

	return valid
}

// String method for easy conversion
func (t ExternalPaymentCancellationRejection1Code) String() string {
	return string(t)
}

// ToExternalPaymentCancellationRejection1Code method for easy conversion with application of restrictions
func ToExternalPaymentCancellationRejection1Code(s string) (ExternalPaymentCancellationRejection1Code, error) {
	if !applyLengthRestriction(s, 0, 1, 4) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalPaymentCancellationRejection1Code", s)
	}

	return ExternalPaymentCancellationRejection1Code(s), nil
}

// MustToExternalPaymentCancellationRejection1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalPaymentCancellationRejection1Code(s string) ExternalPaymentCancellationRejection1Code {
	v, err := ToExternalPaymentCancellationRejection1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ChargeBearerType1Code Ops
 */

var ChargeBearerType1CodeEnumRestriction = []string{"DEBT", "CRED", "SHAR", "SLEV"}

// IsValid checks if ChargeBearerType1Code of type String is valid
func (t ChargeBearerType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(t.String(), ChargeBearerType1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ChargeBearerType1Code) String() string {
	return string(t)
}

// ToChargeBearerType1Code method for easy conversion with application of restrictions
func ToChargeBearerType1Code(s string) (ChargeBearerType1Code, error) {
	if !applyEnumRestriction(s, ChargeBearerType1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type ChargeBearerType1Code", s)
	}

	return ChargeBearerType1Code(s), nil
}

// MustToChargeBearerType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToChargeBearerType1Code(s string) ChargeBearerType1Code {
	v, err := ToChargeBearerType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max10MbBinary Ops
 */

// IsValid checks if Max10MbBinary of type Base64Binary is valid
func (t Max10MbBinary) IsValid(optional bool) bool {

	valid := xsdt.Base64Binary(t).IsValid(optional)
	if 10485760 != 0 {
		valid = valid && len(t) < 10485760
	}

	return valid
}

// String method for easy conversion
func (t Max10MbBinary) String() string {
	return string(t)
}
