// Package common
// Do not Edit. This stuff it's been automatically generated.
// Generated at 2022-04-05 15:30:32.624602 +0200 CEST m=+0.108817042
package common

import (
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
	"regexp"
	"time"
)

/*
 * IBAN2007Identifier Ops
 */

const (
	IBAN2007IdentifierSample = "Value"
)

var IBAN2007IdentifierPatternRestriction = regexp.MustCompile(`[A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}`)

// IsValid checks if IBAN2007Identifier of type String is valid
func (t IBAN2007Identifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isPatternRestrictionValid(t.String(), IBAN2007IdentifierPatternRestriction)

	return valid
}

// String method for easy conversion
func (t IBAN2007Identifier) String() string {
	return string(t)
}

// ToIBAN2007Identifier method for easy conversion with application of restrictions
func ToIBAN2007Identifier(s string) (IBAN2007Identifier, error) {
	if !isPatternRestrictionValid(s, IBAN2007IdentifierPatternRestriction) {
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

// IBAN2007IdentifierExample method for generation of valid sample data
func IBAN2007IdentifierExample() IBAN2007Identifier {
	return IBAN2007Identifier(generateSampleDataWithPatternRestriction(IBAN2007IdentifierPatternRestriction))

}

/*
 * CancellationIndividualStatus1Code Ops
 */

const (
	CancellationIndividualStatus1CodeSample = "Value"
	CancellationIndividualStatus1CodeRJCR   = "RJCR"
	CancellationIndividualStatus1CodeACCR   = "ACCR"
	CancellationIndividualStatus1CodePDCR   = "PDCR"
)

var CancellationIndividualStatus1CodeEnumRestriction = []string{CancellationIndividualStatus1CodeRJCR, CancellationIndividualStatus1CodeACCR, CancellationIndividualStatus1CodePDCR}

// IsValid checks if CancellationIndividualStatus1Code of type String is valid
func (t CancellationIndividualStatus1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), CancellationIndividualStatus1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t CancellationIndividualStatus1Code) String() string {
	return string(t)
}

// ToCancellationIndividualStatus1Code method for easy conversion with application of restrictions
func ToCancellationIndividualStatus1Code(s string) (CancellationIndividualStatus1Code, error) {
	if !isEnumRestrictionValid(s, CancellationIndividualStatus1CodeEnumRestriction) {
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

// CancellationIndividualStatus1CodeExample method for generation of valid sample data
func CancellationIndividualStatus1CodeExample() CancellationIndividualStatus1Code {
	return CancellationIndividualStatus1Code(generateSampleDataWithEnumRestriction(CancellationIndividualStatus1CodeEnumRestriction))

}

/*
 * ExternalClaimNonReceiptRejection1Code Ops
 */

const (
	ExternalClaimNonReceiptRejection1CodeSample    = "Value"
	ExternalClaimNonReceiptRejection1CodeLength    = 0
	ExternalClaimNonReceiptRejection1CodeMinLength = 1
	ExternalClaimNonReceiptRejection1CodeMaxLength = 4
)

// IsValid checks if ExternalClaimNonReceiptRejection1Code of type String is valid
func (t ExternalClaimNonReceiptRejection1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalClaimNonReceiptRejection1CodeLength, ExternalClaimNonReceiptRejection1CodeMinLength, ExternalClaimNonReceiptRejection1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalClaimNonReceiptRejection1Code) String() string {
	return string(t)
}

// ToExternalClaimNonReceiptRejection1Code method for easy conversion with application of restrictions
func ToExternalClaimNonReceiptRejection1Code(s string) (ExternalClaimNonReceiptRejection1Code, error) {
	if !isLengthRestrictionValid(s, ExternalClaimNonReceiptRejection1CodeLength, ExternalClaimNonReceiptRejection1CodeMinLength, ExternalClaimNonReceiptRejection1CodeMaxLength) {
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

// ExternalClaimNonReceiptRejection1CodeExample method for generation of valid sample data
func ExternalClaimNonReceiptRejection1CodeExample() ExternalClaimNonReceiptRejection1Code {
	return ExternalClaimNonReceiptRejection1Code(generateSampleDataWithLengthRestriction(ExternalClaimNonReceiptRejection1CodeLength, ExternalClaimNonReceiptRejection1CodeMinLength, ExternalClaimNonReceiptRejection1CodeMaxLength))

}

/*
 * ExternalCancellationReason1Code Ops
 */

const (
	ExternalCancellationReason1CodeSample    = "Value"
	ExternalCancellationReason1CodeLength    = 0
	ExternalCancellationReason1CodeMinLength = 1
	ExternalCancellationReason1CodeMaxLength = 4
)

// IsValid checks if ExternalCancellationReason1Code of type String is valid
func (t ExternalCancellationReason1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalCancellationReason1CodeLength, ExternalCancellationReason1CodeMinLength, ExternalCancellationReason1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalCancellationReason1Code) String() string {
	return string(t)
}

// ToExternalCancellationReason1Code method for easy conversion with application of restrictions
func ToExternalCancellationReason1Code(s string) (ExternalCancellationReason1Code, error) {
	if !isLengthRestrictionValid(s, ExternalCancellationReason1CodeLength, ExternalCancellationReason1CodeMinLength, ExternalCancellationReason1CodeMaxLength) {
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

// ExternalCancellationReason1CodeExample method for generation of valid sample data
func ExternalCancellationReason1CodeExample() ExternalCancellationReason1Code {
	return ExternalCancellationReason1Code(generateSampleDataWithLengthRestriction(ExternalCancellationReason1CodeLength, ExternalCancellationReason1CodeMinLength, ExternalCancellationReason1CodeMaxLength))

}

/*
 * ExternalStatusReason1Code Ops
 */

const (
	ExternalStatusReason1CodeSample    = "Value"
	ExternalStatusReason1CodeLength    = 0
	ExternalStatusReason1CodeMinLength = 1
	ExternalStatusReason1CodeMaxLength = 4
)

// IsValid checks if ExternalStatusReason1Code of type String is valid
func (t ExternalStatusReason1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalStatusReason1CodeLength, ExternalStatusReason1CodeMinLength, ExternalStatusReason1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalStatusReason1Code) String() string {
	return string(t)
}

// ToExternalStatusReason1Code method for easy conversion with application of restrictions
func ToExternalStatusReason1Code(s string) (ExternalStatusReason1Code, error) {
	if !isLengthRestrictionValid(s, ExternalStatusReason1CodeLength, ExternalStatusReason1CodeMinLength, ExternalStatusReason1CodeMaxLength) {
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

// ExternalStatusReason1CodeExample method for generation of valid sample data
func ExternalStatusReason1CodeExample() ExternalStatusReason1Code {
	return ExternalStatusReason1Code(generateSampleDataWithLengthRestriction(ExternalStatusReason1CodeLength, ExternalStatusReason1CodeMinLength, ExternalStatusReason1CodeMaxLength))

}

/*
 * Max35Text Ops
 */

const (
	Max35TextSample    = "Value"
	Max35TextLength    = 0
	Max35TextMinLength = 1
	Max35TextMaxLength = 35
)

// IsValid checks if Max35Text of type String is valid
func (t Max35Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), Max35TextLength, Max35TextMinLength, Max35TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max35Text) String() string {
	return string(t)
}

// ToMax35Text method for easy conversion with application of restrictions
func ToMax35Text(s string) (Max35Text, error) {
	if !isLengthRestrictionValid(s, Max35TextLength, Max35TextMinLength, Max35TextMaxLength) {
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

// Max35TextExample method for generation of valid sample data
func Max35TextExample() Max35Text {
	return Max35Text(generateSampleDataWithLengthRestriction(Max35TextLength, Max35TextMinLength, Max35TextMaxLength))

}

/*
 * ExternalDiscountAmountType1Code Ops
 */

const (
	ExternalDiscountAmountType1CodeSample    = "Value"
	ExternalDiscountAmountType1CodeLength    = 0
	ExternalDiscountAmountType1CodeMinLength = 1
	ExternalDiscountAmountType1CodeMaxLength = 4
)

// IsValid checks if ExternalDiscountAmountType1Code of type String is valid
func (t ExternalDiscountAmountType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalDiscountAmountType1CodeLength, ExternalDiscountAmountType1CodeMinLength, ExternalDiscountAmountType1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalDiscountAmountType1Code) String() string {
	return string(t)
}

// ToExternalDiscountAmountType1Code method for easy conversion with application of restrictions
func ToExternalDiscountAmountType1Code(s string) (ExternalDiscountAmountType1Code, error) {
	if !isLengthRestrictionValid(s, ExternalDiscountAmountType1CodeLength, ExternalDiscountAmountType1CodeMinLength, ExternalDiscountAmountType1CodeMaxLength) {
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

// ExternalDiscountAmountType1CodeExample method for generation of valid sample data
func ExternalDiscountAmountType1CodeExample() ExternalDiscountAmountType1Code {
	return ExternalDiscountAmountType1Code(generateSampleDataWithLengthRestriction(ExternalDiscountAmountType1CodeLength, ExternalDiscountAmountType1CodeMinLength, ExternalDiscountAmountType1CodeMaxLength))

}

/*
 * ExternalFinancialInstitutionIdentification1Code Ops
 */

const (
	ExternalFinancialInstitutionIdentification1CodeSample    = "Value"
	ExternalFinancialInstitutionIdentification1CodeLength    = 0
	ExternalFinancialInstitutionIdentification1CodeMinLength = 1
	ExternalFinancialInstitutionIdentification1CodeMaxLength = 4
)

// IsValid checks if ExternalFinancialInstitutionIdentification1Code of type String is valid
func (t ExternalFinancialInstitutionIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalFinancialInstitutionIdentification1CodeLength, ExternalFinancialInstitutionIdentification1CodeMinLength, ExternalFinancialInstitutionIdentification1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalFinancialInstitutionIdentification1Code) String() string {
	return string(t)
}

// ToExternalFinancialInstitutionIdentification1Code method for easy conversion with application of restrictions
func ToExternalFinancialInstitutionIdentification1Code(s string) (ExternalFinancialInstitutionIdentification1Code, error) {
	if !isLengthRestrictionValid(s, ExternalFinancialInstitutionIdentification1CodeLength, ExternalFinancialInstitutionIdentification1CodeMinLength, ExternalFinancialInstitutionIdentification1CodeMaxLength) {
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

// ExternalFinancialInstitutionIdentification1CodeExample method for generation of valid sample data
func ExternalFinancialInstitutionIdentification1CodeExample() ExternalFinancialInstitutionIdentification1Code {
	return ExternalFinancialInstitutionIdentification1Code(generateSampleDataWithLengthRestriction(ExternalFinancialInstitutionIdentification1CodeLength, ExternalFinancialInstitutionIdentification1CodeMinLength, ExternalFinancialInstitutionIdentification1CodeMaxLength))

}

/*
 * PaymentMethod4Code Ops
 */

const (
	PaymentMethod4CodeSample = "Value"
	PaymentMethod4CodeCHK    = "CHK"
	PaymentMethod4CodeTRF    = "TRF"
	PaymentMethod4CodeDD     = "DD"
	PaymentMethod4CodeTRA    = "TRA"
)

var PaymentMethod4CodeEnumRestriction = []string{PaymentMethod4CodeCHK, PaymentMethod4CodeTRF, PaymentMethod4CodeDD, PaymentMethod4CodeTRA}

// IsValid checks if PaymentMethod4Code of type String is valid
func (t PaymentMethod4Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), PaymentMethod4CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t PaymentMethod4Code) String() string {
	return string(t)
}

// ToPaymentMethod4Code method for easy conversion with application of restrictions
func ToPaymentMethod4Code(s string) (PaymentMethod4Code, error) {
	if !isEnumRestrictionValid(s, PaymentMethod4CodeEnumRestriction) {
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

// PaymentMethod4CodeExample method for generation of valid sample data
func PaymentMethod4CodeExample() PaymentMethod4Code {
	return PaymentMethod4Code(generateSampleDataWithEnumRestriction(PaymentMethod4CodeEnumRestriction))

}

/*
 * ExternalPaymentCompensationReason1Code Ops
 */

const (
	ExternalPaymentCompensationReason1CodeSample    = "Value"
	ExternalPaymentCompensationReason1CodeLength    = 0
	ExternalPaymentCompensationReason1CodeMinLength = 1
	ExternalPaymentCompensationReason1CodeMaxLength = 4
)

// IsValid checks if ExternalPaymentCompensationReason1Code of type String is valid
func (t ExternalPaymentCompensationReason1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalPaymentCompensationReason1CodeLength, ExternalPaymentCompensationReason1CodeMinLength, ExternalPaymentCompensationReason1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalPaymentCompensationReason1Code) String() string {
	return string(t)
}

// ToExternalPaymentCompensationReason1Code method for easy conversion with application of restrictions
func ToExternalPaymentCompensationReason1Code(s string) (ExternalPaymentCompensationReason1Code, error) {
	if !isLengthRestrictionValid(s, ExternalPaymentCompensationReason1CodeLength, ExternalPaymentCompensationReason1CodeMinLength, ExternalPaymentCompensationReason1CodeMaxLength) {
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

// ExternalPaymentCompensationReason1CodeExample method for generation of valid sample data
func ExternalPaymentCompensationReason1CodeExample() ExternalPaymentCompensationReason1Code {
	return ExternalPaymentCompensationReason1Code(generateSampleDataWithLengthRestriction(ExternalPaymentCompensationReason1CodeLength, ExternalPaymentCompensationReason1CodeMinLength, ExternalPaymentCompensationReason1CodeMaxLength))

}

/*
 * Max4Text Ops
 */

const (
	Max4TextSample    = "Value"
	Max4TextLength    = 0
	Max4TextMinLength = 1
	Max4TextMaxLength = 4
)

// IsValid checks if Max4Text of type String is valid
func (t Max4Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), Max4TextLength, Max4TextMinLength, Max4TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max4Text) String() string {
	return string(t)
}

// ToMax4Text method for easy conversion with application of restrictions
func ToMax4Text(s string) (Max4Text, error) {
	if !isLengthRestrictionValid(s, Max4TextLength, Max4TextMinLength, Max4TextMaxLength) {
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

// Max4TextExample method for generation of valid sample data
func Max4TextExample() Max4Text {
	return Max4Text(generateSampleDataWithLengthRestriction(Max4TextLength, Max4TextMinLength, Max4TextMaxLength))

}

/*
 * ExternalDocumentLineType1Code Ops
 */

const (
	ExternalDocumentLineType1CodeSample    = "Value"
	ExternalDocumentLineType1CodeLength    = 0
	ExternalDocumentLineType1CodeMinLength = 1
	ExternalDocumentLineType1CodeMaxLength = 4
)

// IsValid checks if ExternalDocumentLineType1Code of type String is valid
func (t ExternalDocumentLineType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalDocumentLineType1CodeLength, ExternalDocumentLineType1CodeMinLength, ExternalDocumentLineType1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalDocumentLineType1Code) String() string {
	return string(t)
}

// ToExternalDocumentLineType1Code method for easy conversion with application of restrictions
func ToExternalDocumentLineType1Code(s string) (ExternalDocumentLineType1Code, error) {
	if !isLengthRestrictionValid(s, ExternalDocumentLineType1CodeLength, ExternalDocumentLineType1CodeMinLength, ExternalDocumentLineType1CodeMaxLength) {
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

// ExternalDocumentLineType1CodeExample method for generation of valid sample data
func ExternalDocumentLineType1CodeExample() ExternalDocumentLineType1Code {
	return ExternalDocumentLineType1Code(generateSampleDataWithLengthRestriction(ExternalDocumentLineType1CodeLength, ExternalDocumentLineType1CodeMinLength, ExternalDocumentLineType1CodeMaxLength))

}

/*
 * ExternalClearingSystemIdentification1Code Ops
 */

const (
	ExternalClearingSystemIdentification1CodeSample    = "Value"
	ExternalClearingSystemIdentification1CodeLength    = 0
	ExternalClearingSystemIdentification1CodeMinLength = 1
	ExternalClearingSystemIdentification1CodeMaxLength = 5
)

// IsValid checks if ExternalClearingSystemIdentification1Code of type String is valid
func (t ExternalClearingSystemIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalClearingSystemIdentification1CodeLength, ExternalClearingSystemIdentification1CodeMinLength, ExternalClearingSystemIdentification1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalClearingSystemIdentification1Code) String() string {
	return string(t)
}

// ToExternalClearingSystemIdentification1Code method for easy conversion with application of restrictions
func ToExternalClearingSystemIdentification1Code(s string) (ExternalClearingSystemIdentification1Code, error) {
	if !isLengthRestrictionValid(s, ExternalClearingSystemIdentification1CodeLength, ExternalClearingSystemIdentification1CodeMinLength, ExternalClearingSystemIdentification1CodeMaxLength) {
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

// ExternalClearingSystemIdentification1CodeExample method for generation of valid sample data
func ExternalClearingSystemIdentification1CodeExample() ExternalClearingSystemIdentification1Code {
	return ExternalClearingSystemIdentification1Code(generateSampleDataWithLengthRestriction(ExternalClearingSystemIdentification1CodeLength, ExternalClearingSystemIdentification1CodeMinLength, ExternalClearingSystemIdentification1CodeMaxLength))

}

/*
 * Max1025Text Ops
 */

const (
	Max1025TextSample    = "Value"
	Max1025TextLength    = 0
	Max1025TextMinLength = 1
	Max1025TextMaxLength = 1025
)

// IsValid checks if Max1025Text of type String is valid
func (t Max1025Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), Max1025TextLength, Max1025TextMinLength, Max1025TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max1025Text) String() string {
	return string(t)
}

// ToMax1025Text method for easy conversion with application of restrictions
func ToMax1025Text(s string) (Max1025Text, error) {
	if !isLengthRestrictionValid(s, Max1025TextLength, Max1025TextMinLength, Max1025TextMaxLength) {
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

// Max1025TextExample method for generation of valid sample data
func Max1025TextExample() Max1025Text {
	return Max1025Text(generateSampleDataWithLengthRestriction(Max1025TextLength, Max1025TextMinLength, Max1025TextMaxLength))

}

/*
 * Max10Text Ops
 */

const (
	Max10TextSample    = "Value"
	Max10TextLength    = 0
	Max10TextMinLength = 1
	Max10TextMaxLength = 10
)

// IsValid checks if Max10Text of type String is valid
func (t Max10Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), Max10TextLength, Max10TextMinLength, Max10TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max10Text) String() string {
	return string(t)
}

// ToMax10Text method for easy conversion with application of restrictions
func ToMax10Text(s string) (Max10Text, error) {
	if !isLengthRestrictionValid(s, Max10TextLength, Max10TextMinLength, Max10TextMaxLength) {
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

// Max10TextExample method for generation of valid sample data
func Max10TextExample() Max10Text {
	return Max10Text(generateSampleDataWithLengthRestriction(Max10TextLength, Max10TextMinLength, Max10TextMaxLength))

}

/*
 * CountryCode Ops
 */

const (
	CountryCodeSample = "Value"
)

var CountryCodePatternRestriction = regexp.MustCompile(`[A-Z]{2,2}`)

// IsValid checks if CountryCode of type String is valid
func (t CountryCode) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isPatternRestrictionValid(t.String(), CountryCodePatternRestriction)

	return valid
}

// String method for easy conversion
func (t CountryCode) String() string {
	return string(t)
}

// ToCountryCode method for easy conversion with application of restrictions
func ToCountryCode(s string) (CountryCode, error) {
	if !isPatternRestrictionValid(s, CountryCodePatternRestriction) {
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

// CountryCodeExample method for generation of valid sample data
func CountryCodeExample() CountryCode {
	return CountryCode(generateSampleDataWithPatternRestriction(CountryCodePatternRestriction))

}

/*
 * PaymentMethod7Code Ops
 */

const (
	PaymentMethod7CodeSample = "Value"
	PaymentMethod7CodeCHK    = "CHK"
	PaymentMethod7CodeTRF    = "TRF"
)

var PaymentMethod7CodeEnumRestriction = []string{PaymentMethod7CodeCHK, PaymentMethod7CodeTRF}

// IsValid checks if PaymentMethod7Code of type String is valid
func (t PaymentMethod7Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), PaymentMethod7CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t PaymentMethod7Code) String() string {
	return string(t)
}

// ToPaymentMethod7Code method for easy conversion with application of restrictions
func ToPaymentMethod7Code(s string) (PaymentMethod7Code, error) {
	if !isEnumRestrictionValid(s, PaymentMethod7CodeEnumRestriction) {
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

// PaymentMethod7CodeExample method for generation of valid sample data
func PaymentMethod7CodeExample() PaymentMethod7Code {
	return PaymentMethod7Code(generateSampleDataWithEnumRestriction(PaymentMethod7CodeEnumRestriction))

}

/*
 * Max140Text Ops
 */

const (
	Max140TextSample    = "Value"
	Max140TextLength    = 0
	Max140TextMinLength = 1
	Max140TextMaxLength = 140
)

// IsValid checks if Max140Text of type String is valid
func (t Max140Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), Max140TextLength, Max140TextMinLength, Max140TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max140Text) String() string {
	return string(t)
}

// ToMax140Text method for easy conversion with application of restrictions
func ToMax140Text(s string) (Max140Text, error) {
	if !isLengthRestrictionValid(s, Max140TextLength, Max140TextMinLength, Max140TextMaxLength) {
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

// Max140TextExample method for generation of valid sample data
func Max140TextExample() Max140Text {
	return Max140Text(generateSampleDataWithLengthRestriction(Max140TextLength, Max140TextMinLength, Max140TextMaxLength))

}

/*
 * CreditDebitCode Ops
 */

const (
	CreditDebitCodeSample = "Value"
	CreditDebitCodeCRDT   = "CRDT"
	CreditDebitCodeDBIT   = "DBIT"
)

var CreditDebitCodeEnumRestriction = []string{CreditDebitCodeCRDT, CreditDebitCodeDBIT}

// IsValid checks if CreditDebitCode of type String is valid
func (t CreditDebitCode) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), CreditDebitCodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t CreditDebitCode) String() string {
	return string(t)
}

// ToCreditDebitCode method for easy conversion with application of restrictions
func ToCreditDebitCode(s string) (CreditDebitCode, error) {
	if !isEnumRestrictionValid(s, CreditDebitCodeEnumRestriction) {
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

// CreditDebitCodeExample method for generation of valid sample data
func CreditDebitCodeExample() CreditDebitCode {
	return CreditDebitCode(generateSampleDataWithEnumRestriction(CreditDebitCodeEnumRestriction))

}

/*
 * ExternalPurpose1Code Ops
 */

const (
	ExternalPurpose1CodeSample    = "Value"
	ExternalPurpose1CodeLength    = 0
	ExternalPurpose1CodeMinLength = 1
	ExternalPurpose1CodeMaxLength = 4
)

// IsValid checks if ExternalPurpose1Code of type String is valid
func (t ExternalPurpose1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalPurpose1CodeLength, ExternalPurpose1CodeMinLength, ExternalPurpose1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalPurpose1Code) String() string {
	return string(t)
}

// ToExternalPurpose1Code method for easy conversion with application of restrictions
func ToExternalPurpose1Code(s string) (ExternalPurpose1Code, error) {
	if !isLengthRestrictionValid(s, ExternalPurpose1CodeLength, ExternalPurpose1CodeMinLength, ExternalPurpose1CodeMaxLength) {
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

// ExternalPurpose1CodeExample method for generation of valid sample data
func ExternalPurpose1CodeExample() ExternalPurpose1Code {
	return ExternalPurpose1Code(generateSampleDataWithLengthRestriction(ExternalPurpose1CodeLength, ExternalPurpose1CodeMinLength, ExternalPurpose1CodeMaxLength))

}

/*
 * Max128Text Ops
 */

const (
	Max128TextSample    = "Value"
	Max128TextLength    = 0
	Max128TextMinLength = 1
	Max128TextMaxLength = 128
)

// IsValid checks if Max128Text of type String is valid
func (t Max128Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), Max128TextLength, Max128TextMinLength, Max128TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max128Text) String() string {
	return string(t)
}

// ToMax128Text method for easy conversion with application of restrictions
func ToMax128Text(s string) (Max128Text, error) {
	if !isLengthRestrictionValid(s, Max128TextLength, Max128TextMinLength, Max128TextMaxLength) {
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

// Max128TextExample method for generation of valid sample data
func Max128TextExample() Max128Text {
	return Max128Text(generateSampleDataWithLengthRestriction(Max128TextLength, Max128TextMinLength, Max128TextMaxLength))

}

/*
 * ExternalInvestigationExecutionConfirmation1Code Ops
 */

const (
	ExternalInvestigationExecutionConfirmation1CodeSample    = "Value"
	ExternalInvestigationExecutionConfirmation1CodeLength    = 0
	ExternalInvestigationExecutionConfirmation1CodeMinLength = 1
	ExternalInvestigationExecutionConfirmation1CodeMaxLength = 4
)

// IsValid checks if ExternalInvestigationExecutionConfirmation1Code of type String is valid
func (t ExternalInvestigationExecutionConfirmation1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalInvestigationExecutionConfirmation1CodeLength, ExternalInvestigationExecutionConfirmation1CodeMinLength, ExternalInvestigationExecutionConfirmation1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalInvestigationExecutionConfirmation1Code) String() string {
	return string(t)
}

// ToExternalInvestigationExecutionConfirmation1Code method for easy conversion with application of restrictions
func ToExternalInvestigationExecutionConfirmation1Code(s string) (ExternalInvestigationExecutionConfirmation1Code, error) {
	if !isLengthRestrictionValid(s, ExternalInvestigationExecutionConfirmation1CodeLength, ExternalInvestigationExecutionConfirmation1CodeMinLength, ExternalInvestigationExecutionConfirmation1CodeMaxLength) {
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

// ExternalInvestigationExecutionConfirmation1CodeExample method for generation of valid sample data
func ExternalInvestigationExecutionConfirmation1CodeExample() ExternalInvestigationExecutionConfirmation1Code {
	return ExternalInvestigationExecutionConfirmation1Code(generateSampleDataWithLengthRestriction(ExternalInvestigationExecutionConfirmation1CodeLength, ExternalInvestigationExecutionConfirmation1CodeMinLength, ExternalInvestigationExecutionConfirmation1CodeMaxLength))

}

/*
 * Max350Text Ops
 */

const (
	Max350TextSample    = "Value"
	Max350TextLength    = 0
	Max350TextMinLength = 1
	Max350TextMaxLength = 350
)

// IsValid checks if Max350Text of type String is valid
func (t Max350Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), Max350TextLength, Max350TextMinLength, Max350TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max350Text) String() string {
	return string(t)
}

// ToMax350Text method for easy conversion with application of restrictions
func ToMax350Text(s string) (Max350Text, error) {
	if !isLengthRestrictionValid(s, Max350TextLength, Max350TextMinLength, Max350TextMaxLength) {
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

// Max350TextExample method for generation of valid sample data
func Max350TextExample() Max350Text {
	return Max350Text(generateSampleDataWithLengthRestriction(Max350TextLength, Max350TextMinLength, Max350TextMaxLength))

}

/*
 * BICFIDec2014Identifier Ops
 */

const (
	BICFIDec2014IdentifierSample = "Value"
)

var BICFIDec2014IdentifierPatternRestriction = regexp.MustCompile(`[A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}`)

// IsValid checks if BICFIDec2014Identifier of type String is valid
func (t BICFIDec2014Identifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isPatternRestrictionValid(t.String(), BICFIDec2014IdentifierPatternRestriction)

	return valid
}

// String method for easy conversion
func (t BICFIDec2014Identifier) String() string {
	return string(t)
}

// ToBICFIDec2014Identifier method for easy conversion with application of restrictions
func ToBICFIDec2014Identifier(s string) (BICFIDec2014Identifier, error) {
	if !isPatternRestrictionValid(s, BICFIDec2014IdentifierPatternRestriction) {
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

// BICFIDec2014IdentifierExample method for generation of valid sample data
func BICFIDec2014IdentifierExample() BICFIDec2014Identifier {
	return BICFIDec2014Identifier(generateSampleDataWithPatternRestriction(BICFIDec2014IdentifierPatternRestriction))

}

/*
 * GroupCancellationStatus1Code Ops
 */

const (
	GroupCancellationStatus1CodeSample = "Value"
	GroupCancellationStatus1CodePACR   = "PACR"
	GroupCancellationStatus1CodeRJCR   = "RJCR"
	GroupCancellationStatus1CodeACCR   = "ACCR"
	GroupCancellationStatus1CodePDCR   = "PDCR"
)

var GroupCancellationStatus1CodeEnumRestriction = []string{GroupCancellationStatus1CodePACR, GroupCancellationStatus1CodeRJCR, GroupCancellationStatus1CodeACCR, GroupCancellationStatus1CodePDCR}

// IsValid checks if GroupCancellationStatus1Code of type String is valid
func (t GroupCancellationStatus1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), GroupCancellationStatus1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t GroupCancellationStatus1Code) String() string {
	return string(t)
}

// ToGroupCancellationStatus1Code method for easy conversion with application of restrictions
func ToGroupCancellationStatus1Code(s string) (GroupCancellationStatus1Code, error) {
	if !isEnumRestrictionValid(s, GroupCancellationStatus1CodeEnumRestriction) {
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

// GroupCancellationStatus1CodeExample method for generation of valid sample data
func GroupCancellationStatus1CodeExample() GroupCancellationStatus1Code {
	return GroupCancellationStatus1Code(generateSampleDataWithEnumRestriction(GroupCancellationStatus1CodeEnumRestriction))

}

/*
 * Exact4AlphaNumericText Ops
 */

const (
	Exact4AlphaNumericTextSample = "Value"
)

var Exact4AlphaNumericTextPatternRestriction = regexp.MustCompile(`[a-zA-Z0-9]{4}`)

// IsValid checks if Exact4AlphaNumericText of type String is valid
func (t Exact4AlphaNumericText) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isPatternRestrictionValid(t.String(), Exact4AlphaNumericTextPatternRestriction)

	return valid
}

// String method for easy conversion
func (t Exact4AlphaNumericText) String() string {
	return string(t)
}

// ToExact4AlphaNumericText method for easy conversion with application of restrictions
func ToExact4AlphaNumericText(s string) (Exact4AlphaNumericText, error) {
	if !isPatternRestrictionValid(s, Exact4AlphaNumericTextPatternRestriction) {
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

// Exact4AlphaNumericTextExample method for generation of valid sample data
func Exact4AlphaNumericTextExample() Exact4AlphaNumericText {
	return Exact4AlphaNumericText(generateSampleDataWithPatternRestriction(Exact4AlphaNumericTextPatternRestriction))

}

/*
 * Max16Text Ops
 */

const (
	Max16TextSample    = "Value"
	Max16TextLength    = 0
	Max16TextMinLength = 1
	Max16TextMaxLength = 16
)

// IsValid checks if Max16Text of type String is valid
func (t Max16Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), Max16TextLength, Max16TextMinLength, Max16TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max16Text) String() string {
	return string(t)
}

// ToMax16Text method for easy conversion with application of restrictions
func ToMax16Text(s string) (Max16Text, error) {
	if !isLengthRestrictionValid(s, Max16TextLength, Max16TextMinLength, Max16TextMaxLength) {
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

// Max16TextExample method for generation of valid sample data
func Max16TextExample() Max16Text {
	return Max16Text(generateSampleDataWithLengthRestriction(Max16TextLength, Max16TextMinLength, Max16TextMaxLength))

}

/*
 * ExternalDocumentFormat1Code Ops
 */

const (
	ExternalDocumentFormat1CodeSample    = "Value"
	ExternalDocumentFormat1CodeLength    = 0
	ExternalDocumentFormat1CodeMinLength = 1
	ExternalDocumentFormat1CodeMaxLength = 4
)

// IsValid checks if ExternalDocumentFormat1Code of type String is valid
func (t ExternalDocumentFormat1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalDocumentFormat1CodeLength, ExternalDocumentFormat1CodeMinLength, ExternalDocumentFormat1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalDocumentFormat1Code) String() string {
	return string(t)
}

// ToExternalDocumentFormat1Code method for easy conversion with application of restrictions
func ToExternalDocumentFormat1Code(s string) (ExternalDocumentFormat1Code, error) {
	if !isLengthRestrictionValid(s, ExternalDocumentFormat1CodeLength, ExternalDocumentFormat1CodeMinLength, ExternalDocumentFormat1CodeMaxLength) {
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

// ExternalDocumentFormat1CodeExample method for generation of valid sample data
func ExternalDocumentFormat1CodeExample() ExternalDocumentFormat1Code {
	return ExternalDocumentFormat1Code(generateSampleDataWithLengthRestriction(ExternalDocumentFormat1CodeLength, ExternalDocumentFormat1CodeMinLength, ExternalDocumentFormat1CodeMaxLength))

}

/*
 * DocumentType3Code Ops
 */

const (
	DocumentType3CodeSample = "Value"
	DocumentType3CodeRADM   = "RADM"
	DocumentType3CodeRPIN   = "RPIN"
	DocumentType3CodeFXDR   = "FXDR"
	DocumentType3CodeDISP   = "DISP"
	DocumentType3CodePUOR   = "PUOR"
	DocumentType3CodeSCOR   = "SCOR"
)

var DocumentType3CodeEnumRestriction = []string{DocumentType3CodeRADM, DocumentType3CodeRPIN, DocumentType3CodeFXDR, DocumentType3CodeDISP, DocumentType3CodePUOR, DocumentType3CodeSCOR}

// IsValid checks if DocumentType3Code of type String is valid
func (t DocumentType3Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), DocumentType3CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t DocumentType3Code) String() string {
	return string(t)
}

// ToDocumentType3Code method for easy conversion with application of restrictions
func ToDocumentType3Code(s string) (DocumentType3Code, error) {
	if !isEnumRestrictionValid(s, DocumentType3CodeEnumRestriction) {
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

// DocumentType3CodeExample method for generation of valid sample data
func DocumentType3CodeExample() DocumentType3Code {
	return DocumentType3Code(generateSampleDataWithEnumRestriction(DocumentType3CodeEnumRestriction))

}

/*
 * Max10MbBinary Ops
 */

const (
	Max10MbBinarySample    = "Value"
	Max10MbBinaryLength    = 0
	Max10MbBinaryMinLength = 1
	Max10MbBinaryMaxLength = 10485760
)

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

// ToMax10MbBinary method for easy conversion with application of restrictions
func ToMax10MbBinary(s string) (Max10MbBinary, error) {
	if !isLengthRestrictionValid(s, Max10MbBinaryLength, Max10MbBinaryMinLength, Max10MbBinaryMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max10MbBinary", s)
	}

	return Max10MbBinary(s), nil
}

// MustToMax10MbBinary method for easy conversion with application of restrictions. Panics on error.
func MustToMax10MbBinary(s string) Max10MbBinary {
	v, err := ToMax10MbBinary(s)
	if err != nil {
		panic(err)
	}

	return v
}

// Max10MbBinaryExample method for generation of valid sample data
func Max10MbBinaryExample() Max10MbBinary {
	return Max10MbBinary(generateB64SampleData())
}

/*
 * ISODate Ops
 */

const (
	ISODateSample = "Value"
)

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

// ISODateExample method for generation of valid sample data
func ISODateExample() ISODate {
	return ISODate(time.Now().Format("2006-01-02"))
}

/*
 * ExternalGarnishmentType1Code Ops
 */

const (
	ExternalGarnishmentType1CodeSample    = "Value"
	ExternalGarnishmentType1CodeLength    = 0
	ExternalGarnishmentType1CodeMinLength = 1
	ExternalGarnishmentType1CodeMaxLength = 4
)

// IsValid checks if ExternalGarnishmentType1Code of type String is valid
func (t ExternalGarnishmentType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalGarnishmentType1CodeLength, ExternalGarnishmentType1CodeMinLength, ExternalGarnishmentType1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalGarnishmentType1Code) String() string {
	return string(t)
}

// ToExternalGarnishmentType1Code method for easy conversion with application of restrictions
func ToExternalGarnishmentType1Code(s string) (ExternalGarnishmentType1Code, error) {
	if !isLengthRestrictionValid(s, ExternalGarnishmentType1CodeLength, ExternalGarnishmentType1CodeMinLength, ExternalGarnishmentType1CodeMaxLength) {
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

// ExternalGarnishmentType1CodeExample method for generation of valid sample data
func ExternalGarnishmentType1CodeExample() ExternalGarnishmentType1Code {
	return ExternalGarnishmentType1Code(generateSampleDataWithLengthRestriction(ExternalGarnishmentType1CodeLength, ExternalGarnishmentType1CodeMinLength, ExternalGarnishmentType1CodeMaxLength))

}

/*
 * ActiveCurrencyCode Ops
 */

const (
	ActiveCurrencyCodeSample = "Value"
)

var ActiveCurrencyCodePatternRestriction = regexp.MustCompile(`[A-Z]{3,3}`)

// IsValid checks if ActiveCurrencyCode of type String is valid
func (t ActiveCurrencyCode) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isPatternRestrictionValid(t.String(), ActiveCurrencyCodePatternRestriction)

	return valid
}

// String method for easy conversion
func (t ActiveCurrencyCode) String() string {
	return string(t)
}

// ToActiveCurrencyCode method for easy conversion with application of restrictions
func ToActiveCurrencyCode(s string) (ActiveCurrencyCode, error) {
	if !isPatternRestrictionValid(s, ActiveCurrencyCodePatternRestriction) {
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

// ActiveCurrencyCodeExample method for generation of valid sample data
func ActiveCurrencyCodeExample() ActiveCurrencyCode {
	return ActiveCurrencyCode(generateSampleDataWithPatternRestriction(ActiveCurrencyCodePatternRestriction))

}

/*
 * ChargeBearerType1Code Ops
 */

const (
	ChargeBearerType1CodeSample = "Value"
	ChargeBearerType1CodeDEBT   = "DEBT"
	ChargeBearerType1CodeCRED   = "CRED"
	ChargeBearerType1CodeSHAR   = "SHAR"
	ChargeBearerType1CodeSLEV   = "SLEV"
)

var ChargeBearerType1CodeEnumRestriction = []string{ChargeBearerType1CodeDEBT, ChargeBearerType1CodeCRED, ChargeBearerType1CodeSHAR, ChargeBearerType1CodeSLEV}

// IsValid checks if ChargeBearerType1Code of type String is valid
func (t ChargeBearerType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), ChargeBearerType1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ChargeBearerType1Code) String() string {
	return string(t)
}

// ToChargeBearerType1Code method for easy conversion with application of restrictions
func ToChargeBearerType1Code(s string) (ChargeBearerType1Code, error) {
	if !isEnumRestrictionValid(s, ChargeBearerType1CodeEnumRestriction) {
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

// ChargeBearerType1CodeExample method for generation of valid sample data
func ChargeBearerType1CodeExample() ChargeBearerType1Code {
	return ChargeBearerType1Code(generateSampleDataWithEnumRestriction(ChargeBearerType1CodeEnumRestriction))

}

/*
 * Max2048Text Ops
 */

const (
	Max2048TextSample    = "Value"
	Max2048TextLength    = 0
	Max2048TextMinLength = 1
	Max2048TextMaxLength = 2048
)

// IsValid checks if Max2048Text of type String is valid
func (t Max2048Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), Max2048TextLength, Max2048TextMinLength, Max2048TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max2048Text) String() string {
	return string(t)
}

// ToMax2048Text method for easy conversion with application of restrictions
func ToMax2048Text(s string) (Max2048Text, error) {
	if !isLengthRestrictionValid(s, Max2048TextLength, Max2048TextMinLength, Max2048TextMaxLength) {
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

// Max2048TextExample method for generation of valid sample data
func Max2048TextExample() Max2048Text {
	return Max2048Text(generateSampleDataWithLengthRestriction(Max2048TextLength, Max2048TextMinLength, Max2048TextMaxLength))

}

/*
 * ISODateTime Ops
 */

const (
	ISODateTimeSample = "Value"
)

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

// ISODateTimeExample method for generation of valid sample data
func ISODateTimeExample() ISODateTime {
	return ISODateTime(time.Now().Format(time.RFC3339))
}

/*
 * RegulatoryReportingType1Code Ops
 */

const (
	RegulatoryReportingType1CodeSample = "Value"
	RegulatoryReportingType1CodeCRED   = "CRED"
	RegulatoryReportingType1CodeDEBT   = "DEBT"
	RegulatoryReportingType1CodeBOTH   = "BOTH"
)

var RegulatoryReportingType1CodeEnumRestriction = []string{RegulatoryReportingType1CodeCRED, RegulatoryReportingType1CodeDEBT, RegulatoryReportingType1CodeBOTH}

// IsValid checks if RegulatoryReportingType1Code of type String is valid
func (t RegulatoryReportingType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), RegulatoryReportingType1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t RegulatoryReportingType1Code) String() string {
	return string(t)
}

// ToRegulatoryReportingType1Code method for easy conversion with application of restrictions
func ToRegulatoryReportingType1Code(s string) (RegulatoryReportingType1Code, error) {
	if !isEnumRestrictionValid(s, RegulatoryReportingType1CodeEnumRestriction) {
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

// RegulatoryReportingType1CodeExample method for generation of valid sample data
func RegulatoryReportingType1CodeExample() RegulatoryReportingType1Code {
	return RegulatoryReportingType1Code(generateSampleDataWithEnumRestriction(RegulatoryReportingType1CodeEnumRestriction))

}

/*
 * ExternalServiceLevel1Code Ops
 */

const (
	ExternalServiceLevel1CodeSample    = "Value"
	ExternalServiceLevel1CodeLength    = 0
	ExternalServiceLevel1CodeMinLength = 1
	ExternalServiceLevel1CodeMaxLength = 4
)

// IsValid checks if ExternalServiceLevel1Code of type String is valid
func (t ExternalServiceLevel1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalServiceLevel1CodeLength, ExternalServiceLevel1CodeMinLength, ExternalServiceLevel1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalServiceLevel1Code) String() string {
	return string(t)
}

// ToExternalServiceLevel1Code method for easy conversion with application of restrictions
func ToExternalServiceLevel1Code(s string) (ExternalServiceLevel1Code, error) {
	if !isLengthRestrictionValid(s, ExternalServiceLevel1CodeLength, ExternalServiceLevel1CodeMinLength, ExternalServiceLevel1CodeMaxLength) {
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

// ExternalServiceLevel1CodeExample method for generation of valid sample data
func ExternalServiceLevel1CodeExample() ExternalServiceLevel1Code {
	return ExternalServiceLevel1Code(generateSampleDataWithLengthRestriction(ExternalServiceLevel1CodeLength, ExternalServiceLevel1CodeMinLength, ExternalServiceLevel1CodeMaxLength))

}

/*
 * Frequency6Code Ops
 */

const (
	Frequency6CodeSample = "Value"
	Frequency6CodeYEAR   = "YEAR"
	Frequency6CodeMNTH   = "MNTH"
	Frequency6CodeQURT   = "QURT"
	Frequency6CodeMIAN   = "MIAN"
	Frequency6CodeWEEK   = "WEEK"
	Frequency6CodeDAIL   = "DAIL"
	Frequency6CodeADHO   = "ADHO"
	Frequency6CodeINDA   = "INDA"
	Frequency6CodeFRTN   = "FRTN"
)

var Frequency6CodeEnumRestriction = []string{Frequency6CodeYEAR, Frequency6CodeMNTH, Frequency6CodeQURT, Frequency6CodeMIAN, Frequency6CodeWEEK, Frequency6CodeDAIL, Frequency6CodeADHO, Frequency6CodeINDA, Frequency6CodeFRTN}

// IsValid checks if Frequency6Code of type String is valid
func (t Frequency6Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), Frequency6CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t Frequency6Code) String() string {
	return string(t)
}

// ToFrequency6Code method for easy conversion with application of restrictions
func ToFrequency6Code(s string) (Frequency6Code, error) {
	if !isEnumRestrictionValid(s, Frequency6CodeEnumRestriction) {
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

// Frequency6CodeExample method for generation of valid sample data
func Frequency6CodeExample() Frequency6Code {
	return Frequency6Code(generateSampleDataWithEnumRestriction(Frequency6CodeEnumRestriction))

}

/*
 * PhoneNumber Ops
 */

const (
	PhoneNumberSample = "Value"
)

var PhoneNumberPatternRestriction = regexp.MustCompile(`\+[0-9]{1,3}-[0-9()+\-]{1,30}`)

// IsValid checks if PhoneNumber of type String is valid
func (t PhoneNumber) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isPatternRestrictionValid(t.String(), PhoneNumberPatternRestriction)

	return valid
}

// String method for easy conversion
func (t PhoneNumber) String() string {
	return string(t)
}

// ToPhoneNumber method for easy conversion with application of restrictions
func ToPhoneNumber(s string) (PhoneNumber, error) {
	if !isPatternRestrictionValid(s, PhoneNumberPatternRestriction) {
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

// PhoneNumberExample method for generation of valid sample data
func PhoneNumberExample() PhoneNumber {
	return PhoneNumber(generateSampleDataWithPatternRestriction(PhoneNumberPatternRestriction))

}

/*
 * Max34Text Ops
 */

const (
	Max34TextSample    = "Value"
	Max34TextLength    = 0
	Max34TextMinLength = 1
	Max34TextMaxLength = 34
)

// IsValid checks if Max34Text of type String is valid
func (t Max34Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), Max34TextLength, Max34TextMinLength, Max34TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max34Text) String() string {
	return string(t)
}

// ToMax34Text method for easy conversion with application of restrictions
func ToMax34Text(s string) (Max34Text, error) {
	if !isLengthRestrictionValid(s, Max34TextLength, Max34TextMinLength, Max34TextMaxLength) {
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

// Max34TextExample method for generation of valid sample data
func Max34TextExample() Max34Text {
	return Max34Text(generateSampleDataWithLengthRestriction(Max34TextLength, Max34TextMinLength, Max34TextMaxLength))

}

/*
 * ExternalDocumentType1Code Ops
 */

const (
	ExternalDocumentType1CodeSample    = "Value"
	ExternalDocumentType1CodeLength    = 0
	ExternalDocumentType1CodeMinLength = 1
	ExternalDocumentType1CodeMaxLength = 4
)

// IsValid checks if ExternalDocumentType1Code of type String is valid
func (t ExternalDocumentType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalDocumentType1CodeLength, ExternalDocumentType1CodeMinLength, ExternalDocumentType1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalDocumentType1Code) String() string {
	return string(t)
}

// ToExternalDocumentType1Code method for easy conversion with application of restrictions
func ToExternalDocumentType1Code(s string) (ExternalDocumentType1Code, error) {
	if !isLengthRestrictionValid(s, ExternalDocumentType1CodeLength, ExternalDocumentType1CodeMinLength, ExternalDocumentType1CodeMaxLength) {
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

// ExternalDocumentType1CodeExample method for generation of valid sample data
func ExternalDocumentType1CodeExample() ExternalDocumentType1Code {
	return ExternalDocumentType1Code(generateSampleDataWithLengthRestriction(ExternalDocumentType1CodeLength, ExternalDocumentType1CodeMinLength, ExternalDocumentType1CodeMaxLength))

}

/*
 * ExternalPaymentCancellationRejection1Code Ops
 */

const (
	ExternalPaymentCancellationRejection1CodeSample    = "Value"
	ExternalPaymentCancellationRejection1CodeLength    = 0
	ExternalPaymentCancellationRejection1CodeMinLength = 1
	ExternalPaymentCancellationRejection1CodeMaxLength = 4
)

// IsValid checks if ExternalPaymentCancellationRejection1Code of type String is valid
func (t ExternalPaymentCancellationRejection1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalPaymentCancellationRejection1CodeLength, ExternalPaymentCancellationRejection1CodeMinLength, ExternalPaymentCancellationRejection1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalPaymentCancellationRejection1Code) String() string {
	return string(t)
}

// ToExternalPaymentCancellationRejection1Code method for easy conversion with application of restrictions
func ToExternalPaymentCancellationRejection1Code(s string) (ExternalPaymentCancellationRejection1Code, error) {
	if !isLengthRestrictionValid(s, ExternalPaymentCancellationRejection1CodeLength, ExternalPaymentCancellationRejection1CodeMinLength, ExternalPaymentCancellationRejection1CodeMaxLength) {
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

// ExternalPaymentCancellationRejection1CodeExample method for generation of valid sample data
func ExternalPaymentCancellationRejection1CodeExample() ExternalPaymentCancellationRejection1Code {
	return ExternalPaymentCancellationRejection1Code(generateSampleDataWithLengthRestriction(ExternalPaymentCancellationRejection1CodeLength, ExternalPaymentCancellationRejection1CodeMinLength, ExternalPaymentCancellationRejection1CodeMaxLength))

}

/*
 * ExternalMandateSetupReason1Code Ops
 */

const (
	ExternalMandateSetupReason1CodeSample    = "Value"
	ExternalMandateSetupReason1CodeLength    = 0
	ExternalMandateSetupReason1CodeMinLength = 1
	ExternalMandateSetupReason1CodeMaxLength = 4
)

// IsValid checks if ExternalMandateSetupReason1Code of type String is valid
func (t ExternalMandateSetupReason1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalMandateSetupReason1CodeLength, ExternalMandateSetupReason1CodeMinLength, ExternalMandateSetupReason1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalMandateSetupReason1Code) String() string {
	return string(t)
}

// ToExternalMandateSetupReason1Code method for easy conversion with application of restrictions
func ToExternalMandateSetupReason1Code(s string) (ExternalMandateSetupReason1Code, error) {
	if !isLengthRestrictionValid(s, ExternalMandateSetupReason1CodeLength, ExternalMandateSetupReason1CodeMinLength, ExternalMandateSetupReason1CodeMaxLength) {
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

// ExternalMandateSetupReason1CodeExample method for generation of valid sample data
func ExternalMandateSetupReason1CodeExample() ExternalMandateSetupReason1Code {
	return ExternalMandateSetupReason1Code(generateSampleDataWithLengthRestriction(ExternalMandateSetupReason1CodeLength, ExternalMandateSetupReason1CodeMinLength, ExternalMandateSetupReason1CodeMaxLength))

}

/*
 * ExternalCategoryPurpose1Code Ops
 */

const (
	ExternalCategoryPurpose1CodeSample    = "Value"
	ExternalCategoryPurpose1CodeLength    = 0
	ExternalCategoryPurpose1CodeMinLength = 1
	ExternalCategoryPurpose1CodeMaxLength = 4
)

// IsValid checks if ExternalCategoryPurpose1Code of type String is valid
func (t ExternalCategoryPurpose1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalCategoryPurpose1CodeLength, ExternalCategoryPurpose1CodeMinLength, ExternalCategoryPurpose1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalCategoryPurpose1Code) String() string {
	return string(t)
}

// ToExternalCategoryPurpose1Code method for easy conversion with application of restrictions
func ToExternalCategoryPurpose1Code(s string) (ExternalCategoryPurpose1Code, error) {
	if !isLengthRestrictionValid(s, ExternalCategoryPurpose1CodeLength, ExternalCategoryPurpose1CodeMinLength, ExternalCategoryPurpose1CodeMaxLength) {
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

// ExternalCategoryPurpose1CodeExample method for generation of valid sample data
func ExternalCategoryPurpose1CodeExample() ExternalCategoryPurpose1Code {
	return ExternalCategoryPurpose1Code(generateSampleDataWithLengthRestriction(ExternalCategoryPurpose1CodeLength, ExternalCategoryPurpose1CodeMinLength, ExternalCategoryPurpose1CodeMaxLength))

}

/*
 * ActiveOrHistoricCurrencyCode Ops
 */

const (
	ActiveOrHistoricCurrencyCodeSample = "Value"
)

var ActiveOrHistoricCurrencyCodePatternRestriction = regexp.MustCompile(`[A-Z]{3,3}`)

// IsValid checks if ActiveOrHistoricCurrencyCode of type String is valid
func (t ActiveOrHistoricCurrencyCode) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isPatternRestrictionValid(t.String(), ActiveOrHistoricCurrencyCodePatternRestriction)

	return valid
}

// String method for easy conversion
func (t ActiveOrHistoricCurrencyCode) String() string {
	return string(t)
}

// ToActiveOrHistoricCurrencyCode method for easy conversion with application of restrictions
func ToActiveOrHistoricCurrencyCode(s string) (ActiveOrHistoricCurrencyCode, error) {
	if !isPatternRestrictionValid(s, ActiveOrHistoricCurrencyCodePatternRestriction) {
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

// ActiveOrHistoricCurrencyCodeExample method for generation of valid sample data
func ActiveOrHistoricCurrencyCodeExample() ActiveOrHistoricCurrencyCode {
	return ActiveOrHistoricCurrencyCode(generateSampleDataWithPatternRestriction(ActiveOrHistoricCurrencyCodePatternRestriction))

}

/*
 * ChequeDelivery1Code Ops
 */

const (
	ChequeDelivery1CodeSample = "Value"
	ChequeDelivery1CodeMLDB   = "MLDB"
	ChequeDelivery1CodeMLCD   = "MLCD"
	ChequeDelivery1CodeMLFA   = "MLFA"
	ChequeDelivery1CodeCRDB   = "CRDB"
	ChequeDelivery1CodeCRCD   = "CRCD"
	ChequeDelivery1CodeCRFA   = "CRFA"
	ChequeDelivery1CodePUDB   = "PUDB"
	ChequeDelivery1CodePUCD   = "PUCD"
	ChequeDelivery1CodePUFA   = "PUFA"
	ChequeDelivery1CodeRGDB   = "RGDB"
	ChequeDelivery1CodeRGCD   = "RGCD"
	ChequeDelivery1CodeRGFA   = "RGFA"
)

var ChequeDelivery1CodeEnumRestriction = []string{ChequeDelivery1CodeMLDB, ChequeDelivery1CodeMLCD, ChequeDelivery1CodeMLFA, ChequeDelivery1CodeCRDB, ChequeDelivery1CodeCRCD, ChequeDelivery1CodeCRFA, ChequeDelivery1CodePUDB, ChequeDelivery1CodePUCD, ChequeDelivery1CodePUFA, ChequeDelivery1CodeRGDB, ChequeDelivery1CodeRGCD, ChequeDelivery1CodeRGFA}

// IsValid checks if ChequeDelivery1Code of type String is valid
func (t ChequeDelivery1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), ChequeDelivery1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ChequeDelivery1Code) String() string {
	return string(t)
}

// ToChequeDelivery1Code method for easy conversion with application of restrictions
func ToChequeDelivery1Code(s string) (ChequeDelivery1Code, error) {
	if !isEnumRestrictionValid(s, ChequeDelivery1CodeEnumRestriction) {
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

// ChequeDelivery1CodeExample method for generation of valid sample data
func ChequeDelivery1CodeExample() ChequeDelivery1Code {
	return ChequeDelivery1Code(generateSampleDataWithEnumRestriction(ChequeDelivery1CodeEnumRestriction))

}

/*
 * PreferredContactMethod1Code Ops
 */

const (
	PreferredContactMethod1CodeSample = "Value"
	PreferredContactMethod1CodeLETT   = "LETT"
	PreferredContactMethod1CodeMAIL   = "MAIL"
	PreferredContactMethod1CodePHON   = "PHON"
	PreferredContactMethod1CodeFAXX   = "FAXX"
	PreferredContactMethod1CodeCELL   = "CELL"
)

var PreferredContactMethod1CodeEnumRestriction = []string{PreferredContactMethod1CodeLETT, PreferredContactMethod1CodeMAIL, PreferredContactMethod1CodePHON, PreferredContactMethod1CodeFAXX, PreferredContactMethod1CodeCELL}

// IsValid checks if PreferredContactMethod1Code of type String is valid
func (t PreferredContactMethod1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), PreferredContactMethod1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t PreferredContactMethod1Code) String() string {
	return string(t)
}

// ToPreferredContactMethod1Code method for easy conversion with application of restrictions
func ToPreferredContactMethod1Code(s string) (PreferredContactMethod1Code, error) {
	if !isEnumRestrictionValid(s, PreferredContactMethod1CodeEnumRestriction) {
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

// PreferredContactMethod1CodeExample method for generation of valid sample data
func PreferredContactMethod1CodeExample() PreferredContactMethod1Code {
	return PreferredContactMethod1Code(generateSampleDataWithEnumRestriction(PreferredContactMethod1CodeEnumRestriction))

}

/*
 * TransactionIndividualStatus1Code Ops
 */

const (
	TransactionIndividualStatus1CodeSample = "Value"
	TransactionIndividualStatus1CodeACTC   = "ACTC"
	TransactionIndividualStatus1CodeRJCT   = "RJCT"
	TransactionIndividualStatus1CodePDNG   = "PDNG"
	TransactionIndividualStatus1CodeACCP   = "ACCP"
	TransactionIndividualStatus1CodeACSP   = "ACSP"
	TransactionIndividualStatus1CodeACSC   = "ACSC"
	TransactionIndividualStatus1CodeACCR   = "ACCR"
	TransactionIndividualStatus1CodeACWC   = "ACWC"
)

var TransactionIndividualStatus1CodeEnumRestriction = []string{TransactionIndividualStatus1CodeACTC, TransactionIndividualStatus1CodeRJCT, TransactionIndividualStatus1CodePDNG, TransactionIndividualStatus1CodeACCP, TransactionIndividualStatus1CodeACSP, TransactionIndividualStatus1CodeACSC, TransactionIndividualStatus1CodeACCR, TransactionIndividualStatus1CodeACWC}

// IsValid checks if TransactionIndividualStatus1Code of type String is valid
func (t TransactionIndividualStatus1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), TransactionIndividualStatus1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t TransactionIndividualStatus1Code) String() string {
	return string(t)
}

// ToTransactionIndividualStatus1Code method for easy conversion with application of restrictions
func ToTransactionIndividualStatus1Code(s string) (TransactionIndividualStatus1Code, error) {
	if !isEnumRestrictionValid(s, TransactionIndividualStatus1CodeEnumRestriction) {
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

// TransactionIndividualStatus1CodeExample method for generation of valid sample data
func TransactionIndividualStatus1CodeExample() TransactionIndividualStatus1Code {
	return TransactionIndividualStatus1Code(generateSampleDataWithEnumRestriction(TransactionIndividualStatus1CodeEnumRestriction))

}

/*
 * SequenceType3Code Ops
 */

const (
	SequenceType3CodeSample = "Value"
	SequenceType3CodeFRST   = "FRST"
	SequenceType3CodeRCUR   = "RCUR"
	SequenceType3CodeFNAL   = "FNAL"
	SequenceType3CodeOOFF   = "OOFF"
	SequenceType3CodeRPRE   = "RPRE"
)

var SequenceType3CodeEnumRestriction = []string{SequenceType3CodeFRST, SequenceType3CodeRCUR, SequenceType3CodeFNAL, SequenceType3CodeOOFF, SequenceType3CodeRPRE}

// IsValid checks if SequenceType3Code of type String is valid
func (t SequenceType3Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), SequenceType3CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t SequenceType3Code) String() string {
	return string(t)
}

// ToSequenceType3Code method for easy conversion with application of restrictions
func ToSequenceType3Code(s string) (SequenceType3Code, error) {
	if !isEnumRestrictionValid(s, SequenceType3CodeEnumRestriction) {
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

// SequenceType3CodeExample method for generation of valid sample data
func SequenceType3CodeExample() SequenceType3Code {
	return SequenceType3Code(generateSampleDataWithEnumRestriction(SequenceType3CodeEnumRestriction))

}

/*
 * Instruction3Code Ops
 */

const (
	Instruction3CodeSample = "Value"
	Instruction3CodeCHQB   = "CHQB"
	Instruction3CodeHOLD   = "HOLD"
	Instruction3CodePHOB   = "PHOB"
	Instruction3CodeTELB   = "TELB"
)

var Instruction3CodeEnumRestriction = []string{Instruction3CodeCHQB, Instruction3CodeHOLD, Instruction3CodePHOB, Instruction3CodeTELB}

// IsValid checks if Instruction3Code of type String is valid
func (t Instruction3Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), Instruction3CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t Instruction3Code) String() string {
	return string(t)
}

// ToInstruction3Code method for easy conversion with application of restrictions
func ToInstruction3Code(s string) (Instruction3Code, error) {
	if !isEnumRestrictionValid(s, Instruction3CodeEnumRestriction) {
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

// Instruction3CodeExample method for generation of valid sample data
func Instruction3CodeExample() Instruction3Code {
	return Instruction3Code(generateSampleDataWithEnumRestriction(Instruction3CodeEnumRestriction))

}

/*
 * Max15NumericText Ops
 */

const (
	Max15NumericTextSample = "Value"
)

var Max15NumericTextPatternRestriction = regexp.MustCompile(`[0-9]{1,15}`)

// IsValid checks if Max15NumericText of type String is valid
func (t Max15NumericText) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isPatternRestrictionValid(t.String(), Max15NumericTextPatternRestriction)

	return valid
}

// String method for easy conversion
func (t Max15NumericText) String() string {
	return string(t)
}

// ToMax15NumericText method for easy conversion with application of restrictions
func ToMax15NumericText(s string) (Max15NumericText, error) {
	if !isPatternRestrictionValid(s, Max15NumericTextPatternRestriction) {
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

// Max15NumericTextExample method for generation of valid sample data
func Max15NumericTextExample() Max15NumericText {
	return Max15NumericText(generateSampleDataWithPatternRestriction(Max15NumericTextPatternRestriction))

}

/*
 * Max70Text Ops
 */

const (
	Max70TextSample    = "Value"
	Max70TextLength    = 0
	Max70TextMinLength = 1
	Max70TextMaxLength = 70
)

// IsValid checks if Max70Text of type String is valid
func (t Max70Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), Max70TextLength, Max70TextMinLength, Max70TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max70Text) String() string {
	return string(t)
}

// ToMax70Text method for easy conversion with application of restrictions
func ToMax70Text(s string) (Max70Text, error) {
	if !isLengthRestrictionValid(s, Max70TextLength, Max70TextMinLength, Max70TextMaxLength) {
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

// Max70TextExample method for generation of valid sample data
func Max70TextExample() Max70Text {
	return Max70Text(generateSampleDataWithLengthRestriction(Max70TextLength, Max70TextMinLength, Max70TextMaxLength))

}

/*
 * LEIIdentifier Ops
 */

const (
	LEIIdentifierSample = "Value"
)

var LEIIdentifierPatternRestriction = regexp.MustCompile(`[A-Z0-9]{18,18}[0-9]{2,2}`)

// IsValid checks if LEIIdentifier of type String is valid
func (t LEIIdentifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isPatternRestrictionValid(t.String(), LEIIdentifierPatternRestriction)

	return valid
}

// String method for easy conversion
func (t LEIIdentifier) String() string {
	return string(t)
}

// ToLEIIdentifier method for easy conversion with application of restrictions
func ToLEIIdentifier(s string) (LEIIdentifier, error) {
	if !isPatternRestrictionValid(s, LEIIdentifierPatternRestriction) {
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

// LEIIdentifierExample method for generation of valid sample data
func LEIIdentifierExample() LEIIdentifier {
	return LEIIdentifier(generateSampleDataWithPatternRestriction(LEIIdentifierPatternRestriction))

}

/*
 * ExternalCashClearingSystem1Code Ops
 */

const (
	ExternalCashClearingSystem1CodeSample    = "Value"
	ExternalCashClearingSystem1CodeLength    = 0
	ExternalCashClearingSystem1CodeMinLength = 1
	ExternalCashClearingSystem1CodeMaxLength = 3
)

// IsValid checks if ExternalCashClearingSystem1Code of type String is valid
func (t ExternalCashClearingSystem1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalCashClearingSystem1CodeLength, ExternalCashClearingSystem1CodeMinLength, ExternalCashClearingSystem1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalCashClearingSystem1Code) String() string {
	return string(t)
}

// ToExternalCashClearingSystem1Code method for easy conversion with application of restrictions
func ToExternalCashClearingSystem1Code(s string) (ExternalCashClearingSystem1Code, error) {
	if !isLengthRestrictionValid(s, ExternalCashClearingSystem1CodeLength, ExternalCashClearingSystem1CodeMinLength, ExternalCashClearingSystem1CodeMaxLength) {
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

// ExternalCashClearingSystem1CodeExample method for generation of valid sample data
func ExternalCashClearingSystem1CodeExample() ExternalCashClearingSystem1Code {
	return ExternalCashClearingSystem1Code(generateSampleDataWithLengthRestriction(ExternalCashClearingSystem1CodeLength, ExternalCashClearingSystem1CodeMinLength, ExternalCashClearingSystem1CodeMaxLength))

}

/*
 * Exact2NumericText Ops
 */

const (
	Exact2NumericTextSample = "Value"
)

var Exact2NumericTextPatternRestriction = regexp.MustCompile(`[0-9]{2}`)

// IsValid checks if Exact2NumericText of type String is valid
func (t Exact2NumericText) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isPatternRestrictionValid(t.String(), Exact2NumericTextPatternRestriction)

	return valid
}

// String method for easy conversion
func (t Exact2NumericText) String() string {
	return string(t)
}

// ToExact2NumericText method for easy conversion with application of restrictions
func ToExact2NumericText(s string) (Exact2NumericText, error) {
	if !isPatternRestrictionValid(s, Exact2NumericTextPatternRestriction) {
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

// Exact2NumericTextExample method for generation of valid sample data
func Exact2NumericTextExample() Exact2NumericText {
	return Exact2NumericText(generateSampleDataWithPatternRestriction(Exact2NumericTextPatternRestriction))

}

/*
 * SettlementMethod1Code Ops
 */

const (
	SettlementMethod1CodeSample = "Value"
	SettlementMethod1CodeINDA   = "INDA"
	SettlementMethod1CodeINGA   = "INGA"
	SettlementMethod1CodeCOVE   = "COVE"
	SettlementMethod1CodeCLRG   = "CLRG"
)

var SettlementMethod1CodeEnumRestriction = []string{SettlementMethod1CodeINDA, SettlementMethod1CodeINGA, SettlementMethod1CodeCOVE, SettlementMethod1CodeCLRG}

// IsValid checks if SettlementMethod1Code of type String is valid
func (t SettlementMethod1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), SettlementMethod1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t SettlementMethod1Code) String() string {
	return string(t)
}

// ToSettlementMethod1Code method for easy conversion with application of restrictions
func ToSettlementMethod1Code(s string) (SettlementMethod1Code, error) {
	if !isEnumRestrictionValid(s, SettlementMethod1CodeEnumRestriction) {
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

// SettlementMethod1CodeExample method for generation of valid sample data
func SettlementMethod1CodeExample() SettlementMethod1Code {
	return SettlementMethod1Code(generateSampleDataWithEnumRestriction(SettlementMethod1CodeEnumRestriction))

}

/*
 * ExternalAccountIdentification1Code Ops
 */

const (
	ExternalAccountIdentification1CodeSample    = "Value"
	ExternalAccountIdentification1CodeLength    = 0
	ExternalAccountIdentification1CodeMinLength = 1
	ExternalAccountIdentification1CodeMaxLength = 4
)

// IsValid checks if ExternalAccountIdentification1Code of type String is valid
func (t ExternalAccountIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalAccountIdentification1CodeLength, ExternalAccountIdentification1CodeMinLength, ExternalAccountIdentification1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalAccountIdentification1Code) String() string {
	return string(t)
}

// ToExternalAccountIdentification1Code method for easy conversion with application of restrictions
func ToExternalAccountIdentification1Code(s string) (ExternalAccountIdentification1Code, error) {
	if !isLengthRestrictionValid(s, ExternalAccountIdentification1CodeLength, ExternalAccountIdentification1CodeMinLength, ExternalAccountIdentification1CodeMaxLength) {
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

// ExternalAccountIdentification1CodeExample method for generation of valid sample data
func ExternalAccountIdentification1CodeExample() ExternalAccountIdentification1Code {
	return ExternalAccountIdentification1Code(generateSampleDataWithLengthRestriction(ExternalAccountIdentification1CodeLength, ExternalAccountIdentification1CodeMinLength, ExternalAccountIdentification1CodeMaxLength))

}

/*
 * TaxRecordPeriod1Code Ops
 */

const (
	TaxRecordPeriod1CodeSample = "Value"
	TaxRecordPeriod1CodeMM01   = "MM01"
	TaxRecordPeriod1CodeMM02   = "MM02"
	TaxRecordPeriod1CodeMM03   = "MM03"
	TaxRecordPeriod1CodeMM04   = "MM04"
	TaxRecordPeriod1CodeMM05   = "MM05"
	TaxRecordPeriod1CodeMM06   = "MM06"
	TaxRecordPeriod1CodeMM07   = "MM07"
	TaxRecordPeriod1CodeMM08   = "MM08"
	TaxRecordPeriod1CodeMM09   = "MM09"
	TaxRecordPeriod1CodeMM10   = "MM10"
	TaxRecordPeriod1CodeMM11   = "MM11"
	TaxRecordPeriod1CodeMM12   = "MM12"
	TaxRecordPeriod1CodeQTR1   = "QTR1"
	TaxRecordPeriod1CodeQTR2   = "QTR2"
	TaxRecordPeriod1CodeQTR3   = "QTR3"
	TaxRecordPeriod1CodeQTR4   = "QTR4"
	TaxRecordPeriod1CodeHLF1   = "HLF1"
	TaxRecordPeriod1CodeHLF2   = "HLF2"
)

var TaxRecordPeriod1CodeEnumRestriction = []string{TaxRecordPeriod1CodeMM01, TaxRecordPeriod1CodeMM02, TaxRecordPeriod1CodeMM03, TaxRecordPeriod1CodeMM04, TaxRecordPeriod1CodeMM05, TaxRecordPeriod1CodeMM06, TaxRecordPeriod1CodeMM07, TaxRecordPeriod1CodeMM08, TaxRecordPeriod1CodeMM09, TaxRecordPeriod1CodeMM10, TaxRecordPeriod1CodeMM11, TaxRecordPeriod1CodeMM12, TaxRecordPeriod1CodeQTR1, TaxRecordPeriod1CodeQTR2, TaxRecordPeriod1CodeQTR3, TaxRecordPeriod1CodeQTR4, TaxRecordPeriod1CodeHLF1, TaxRecordPeriod1CodeHLF2}

// IsValid checks if TaxRecordPeriod1Code of type String is valid
func (t TaxRecordPeriod1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), TaxRecordPeriod1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t TaxRecordPeriod1Code) String() string {
	return string(t)
}

// ToTaxRecordPeriod1Code method for easy conversion with application of restrictions
func ToTaxRecordPeriod1Code(s string) (TaxRecordPeriod1Code, error) {
	if !isEnumRestrictionValid(s, TaxRecordPeriod1CodeEnumRestriction) {
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

// TaxRecordPeriod1CodeExample method for generation of valid sample data
func TaxRecordPeriod1CodeExample() TaxRecordPeriod1Code {
	return TaxRecordPeriod1Code(generateSampleDataWithEnumRestriction(TaxRecordPeriod1CodeEnumRestriction))

}

/*
 * RemittanceLocationMethod2Code Ops
 */

const (
	RemittanceLocationMethod2CodeSample = "Value"
	RemittanceLocationMethod2CodeFAXI   = "FAXI"
	RemittanceLocationMethod2CodeEDIC   = "EDIC"
	RemittanceLocationMethod2CodeURID   = "URID"
	RemittanceLocationMethod2CodeEMAL   = "EMAL"
	RemittanceLocationMethod2CodePOST   = "POST"
	RemittanceLocationMethod2CodeSMSM   = "SMSM"
)

var RemittanceLocationMethod2CodeEnumRestriction = []string{RemittanceLocationMethod2CodeFAXI, RemittanceLocationMethod2CodeEDIC, RemittanceLocationMethod2CodeURID, RemittanceLocationMethod2CodeEMAL, RemittanceLocationMethod2CodePOST, RemittanceLocationMethod2CodeSMSM}

// IsValid checks if RemittanceLocationMethod2Code of type String is valid
func (t RemittanceLocationMethod2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), RemittanceLocationMethod2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t RemittanceLocationMethod2Code) String() string {
	return string(t)
}

// ToRemittanceLocationMethod2Code method for easy conversion with application of restrictions
func ToRemittanceLocationMethod2Code(s string) (RemittanceLocationMethod2Code, error) {
	if !isEnumRestrictionValid(s, RemittanceLocationMethod2CodeEnumRestriction) {
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

// RemittanceLocationMethod2CodeExample method for generation of valid sample data
func RemittanceLocationMethod2CodeExample() RemittanceLocationMethod2Code {
	return RemittanceLocationMethod2Code(generateSampleDataWithEnumRestriction(RemittanceLocationMethod2CodeEnumRestriction))

}

/*
 * AddressType2Code Ops
 */

const (
	AddressType2CodeSample = "Value"
	AddressType2CodeADDR   = "ADDR"
	AddressType2CodePBOX   = "PBOX"
	AddressType2CodeHOME   = "HOME"
	AddressType2CodeBIZZ   = "BIZZ"
	AddressType2CodeMLTO   = "MLTO"
	AddressType2CodeDLVY   = "DLVY"
)

var AddressType2CodeEnumRestriction = []string{AddressType2CodeADDR, AddressType2CodePBOX, AddressType2CodeHOME, AddressType2CodeBIZZ, AddressType2CodeMLTO, AddressType2CodeDLVY}

// IsValid checks if AddressType2Code of type String is valid
func (t AddressType2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), AddressType2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t AddressType2Code) String() string {
	return string(t)
}

// ToAddressType2Code method for easy conversion with application of restrictions
func ToAddressType2Code(s string) (AddressType2Code, error) {
	if !isEnumRestrictionValid(s, AddressType2CodeEnumRestriction) {
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

// AddressType2CodeExample method for generation of valid sample data
func AddressType2CodeExample() AddressType2Code {
	return AddressType2Code(generateSampleDataWithEnumRestriction(AddressType2CodeEnumRestriction))

}

/*
 * NamePrefix2Code Ops
 */

const (
	NamePrefix2CodeSample = "Value"
	NamePrefix2CodeDOCT   = "DOCT"
	NamePrefix2CodeMADM   = "MADM"
	NamePrefix2CodeMISS   = "MISS"
	NamePrefix2CodeMIST   = "MIST"
	NamePrefix2CodeMIKS   = "MIKS"
)

var NamePrefix2CodeEnumRestriction = []string{NamePrefix2CodeDOCT, NamePrefix2CodeMADM, NamePrefix2CodeMISS, NamePrefix2CodeMIST, NamePrefix2CodeMIKS}

// IsValid checks if NamePrefix2Code of type String is valid
func (t NamePrefix2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), NamePrefix2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t NamePrefix2Code) String() string {
	return string(t)
}

// ToNamePrefix2Code method for easy conversion with application of restrictions
func ToNamePrefix2Code(s string) (NamePrefix2Code, error) {
	if !isEnumRestrictionValid(s, NamePrefix2CodeEnumRestriction) {
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

// NamePrefix2CodeExample method for generation of valid sample data
func NamePrefix2CodeExample() NamePrefix2Code {
	return NamePrefix2Code(generateSampleDataWithEnumRestriction(NamePrefix2CodeEnumRestriction))

}

/*
 * ExternalChargeType1Code Ops
 */

const (
	ExternalChargeType1CodeSample    = "Value"
	ExternalChargeType1CodeLength    = 0
	ExternalChargeType1CodeMinLength = 1
	ExternalChargeType1CodeMaxLength = 4
)

// IsValid checks if ExternalChargeType1Code of type String is valid
func (t ExternalChargeType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalChargeType1CodeLength, ExternalChargeType1CodeMinLength, ExternalChargeType1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalChargeType1Code) String() string {
	return string(t)
}

// ToExternalChargeType1Code method for easy conversion with application of restrictions
func ToExternalChargeType1Code(s string) (ExternalChargeType1Code, error) {
	if !isLengthRestrictionValid(s, ExternalChargeType1CodeLength, ExternalChargeType1CodeMinLength, ExternalChargeType1CodeMaxLength) {
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

// ExternalChargeType1CodeExample method for generation of valid sample data
func ExternalChargeType1CodeExample() ExternalChargeType1Code {
	return ExternalChargeType1Code(generateSampleDataWithLengthRestriction(ExternalChargeType1CodeLength, ExternalChargeType1CodeMinLength, ExternalChargeType1CodeMaxLength))

}

/*
 * Priority2Code Ops
 */

const (
	Priority2CodeSample = "Value"
	Priority2CodeHIGH   = "HIGH"
	Priority2CodeNORM   = "NORM"
)

var Priority2CodeEnumRestriction = []string{Priority2CodeHIGH, Priority2CodeNORM}

// IsValid checks if Priority2Code of type String is valid
func (t Priority2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), Priority2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t Priority2Code) String() string {
	return string(t)
}

// ToPriority2Code method for easy conversion with application of restrictions
func ToPriority2Code(s string) (Priority2Code, error) {
	if !isEnumRestrictionValid(s, Priority2CodeEnumRestriction) {
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

// Priority2CodeExample method for generation of valid sample data
func Priority2CodeExample() Priority2Code {
	return Priority2Code(generateSampleDataWithEnumRestriction(Priority2CodeEnumRestriction))

}

/*
 * AnyBICDec2014Identifier Ops
 */

const (
	AnyBICDec2014IdentifierSample = "Value"
)

var AnyBICDec2014IdentifierPatternRestriction = regexp.MustCompile(`[A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}`)

// IsValid checks if AnyBICDec2014Identifier of type String is valid
func (t AnyBICDec2014Identifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isPatternRestrictionValid(t.String(), AnyBICDec2014IdentifierPatternRestriction)

	return valid
}

// String method for easy conversion
func (t AnyBICDec2014Identifier) String() string {
	return string(t)
}

// ToAnyBICDec2014Identifier method for easy conversion with application of restrictions
func ToAnyBICDec2014Identifier(s string) (AnyBICDec2014Identifier, error) {
	if !isPatternRestrictionValid(s, AnyBICDec2014IdentifierPatternRestriction) {
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

// AnyBICDec2014IdentifierExample method for generation of valid sample data
func AnyBICDec2014IdentifierExample() AnyBICDec2014Identifier {
	return AnyBICDec2014Identifier(generateSampleDataWithPatternRestriction(AnyBICDec2014IdentifierPatternRestriction))

}

/*
 * UUIDv4Identifier Ops
 */

const (
	UUIDv4IdentifierSample = "Value"
)

var UUIDv4IdentifierPatternRestriction = regexp.MustCompile(`[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}`)

// IsValid checks if UUIDv4Identifier of type String is valid
func (t UUIDv4Identifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isPatternRestrictionValid(t.String(), UUIDv4IdentifierPatternRestriction)

	return valid
}

// String method for easy conversion
func (t UUIDv4Identifier) String() string {
	return string(t)
}

// ToUUIDv4Identifier method for easy conversion with application of restrictions
func ToUUIDv4Identifier(s string) (UUIDv4Identifier, error) {
	if !isPatternRestrictionValid(s, UUIDv4IdentifierPatternRestriction) {
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

// UUIDv4IdentifierExample method for generation of valid sample data
func UUIDv4IdentifierExample() UUIDv4Identifier {
	return UUIDv4Identifier(generateSampleDataWithPatternRestriction(UUIDv4IdentifierPatternRestriction))

}

/*
 * DocumentType6Code Ops
 */

const (
	DocumentType6CodeSample = "Value"
	DocumentType6CodeMSIN   = "MSIN"
	DocumentType6CodeCNFA   = "CNFA"
	DocumentType6CodeDNFA   = "DNFA"
	DocumentType6CodeCINV   = "CINV"
	DocumentType6CodeCREN   = "CREN"
	DocumentType6CodeDEBN   = "DEBN"
	DocumentType6CodeHIRI   = "HIRI"
	DocumentType6CodeSBIN   = "SBIN"
	DocumentType6CodeCMCN   = "CMCN"
	DocumentType6CodeSOAC   = "SOAC"
	DocumentType6CodeDISP   = "DISP"
	DocumentType6CodeBOLD   = "BOLD"
	DocumentType6CodeVCHR   = "VCHR"
	DocumentType6CodeAROI   = "AROI"
	DocumentType6CodeTSUT   = "TSUT"
	DocumentType6CodePUOR   = "PUOR"
)

var DocumentType6CodeEnumRestriction = []string{DocumentType6CodeMSIN, DocumentType6CodeCNFA, DocumentType6CodeDNFA, DocumentType6CodeCINV, DocumentType6CodeCREN, DocumentType6CodeDEBN, DocumentType6CodeHIRI, DocumentType6CodeSBIN, DocumentType6CodeCMCN, DocumentType6CodeSOAC, DocumentType6CodeDISP, DocumentType6CodeBOLD, DocumentType6CodeVCHR, DocumentType6CodeAROI, DocumentType6CodeTSUT, DocumentType6CodePUOR}

// IsValid checks if DocumentType6Code of type String is valid
func (t DocumentType6Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), DocumentType6CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t DocumentType6Code) String() string {
	return string(t)
}

// ToDocumentType6Code method for easy conversion with application of restrictions
func ToDocumentType6Code(s string) (DocumentType6Code, error) {
	if !isEnumRestrictionValid(s, DocumentType6CodeEnumRestriction) {
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

// DocumentType6CodeExample method for generation of valid sample data
func DocumentType6CodeExample() DocumentType6Code {
	return DocumentType6Code(generateSampleDataWithEnumRestriction(DocumentType6CodeEnumRestriction))

}

/*
 * ExternalPaymentTransactionStatus1Code Ops
 */

const (
	ExternalPaymentTransactionStatus1CodeSample    = "Value"
	ExternalPaymentTransactionStatus1CodeLength    = 0
	ExternalPaymentTransactionStatus1CodeMinLength = 1
	ExternalPaymentTransactionStatus1CodeMaxLength = 4
)

// IsValid checks if ExternalPaymentTransactionStatus1Code of type String is valid
func (t ExternalPaymentTransactionStatus1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalPaymentTransactionStatus1CodeLength, ExternalPaymentTransactionStatus1CodeMinLength, ExternalPaymentTransactionStatus1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalPaymentTransactionStatus1Code) String() string {
	return string(t)
}

// ToExternalPaymentTransactionStatus1Code method for easy conversion with application of restrictions
func ToExternalPaymentTransactionStatus1Code(s string) (ExternalPaymentTransactionStatus1Code, error) {
	if !isLengthRestrictionValid(s, ExternalPaymentTransactionStatus1CodeLength, ExternalPaymentTransactionStatus1CodeMinLength, ExternalPaymentTransactionStatus1CodeMaxLength) {
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

// ExternalPaymentTransactionStatus1CodeExample method for generation of valid sample data
func ExternalPaymentTransactionStatus1CodeExample() ExternalPaymentTransactionStatus1Code {
	return ExternalPaymentTransactionStatus1Code(generateSampleDataWithLengthRestriction(ExternalPaymentTransactionStatus1CodeLength, ExternalPaymentTransactionStatus1CodeMinLength, ExternalPaymentTransactionStatus1CodeMaxLength))

}

/*
 * ExternalTaxAmountType1Code Ops
 */

const (
	ExternalTaxAmountType1CodeSample    = "Value"
	ExternalTaxAmountType1CodeLength    = 0
	ExternalTaxAmountType1CodeMinLength = 1
	ExternalTaxAmountType1CodeMaxLength = 4
)

// IsValid checks if ExternalTaxAmountType1Code of type String is valid
func (t ExternalTaxAmountType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalTaxAmountType1CodeLength, ExternalTaxAmountType1CodeMinLength, ExternalTaxAmountType1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalTaxAmountType1Code) String() string {
	return string(t)
}

// ToExternalTaxAmountType1Code method for easy conversion with application of restrictions
func ToExternalTaxAmountType1Code(s string) (ExternalTaxAmountType1Code, error) {
	if !isLengthRestrictionValid(s, ExternalTaxAmountType1CodeLength, ExternalTaxAmountType1CodeMinLength, ExternalTaxAmountType1CodeMaxLength) {
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

// ExternalTaxAmountType1CodeExample method for generation of valid sample data
func ExternalTaxAmountType1CodeExample() ExternalTaxAmountType1Code {
	return ExternalTaxAmountType1Code(generateSampleDataWithLengthRestriction(ExternalTaxAmountType1CodeLength, ExternalTaxAmountType1CodeMinLength, ExternalTaxAmountType1CodeMaxLength))

}

/*
 * ChequeType2Code Ops
 */

const (
	ChequeType2CodeSample = "Value"
	ChequeType2CodeCCHQ   = "CCHQ"
	ChequeType2CodeCCCH   = "CCCH"
	ChequeType2CodeBCHQ   = "BCHQ"
	ChequeType2CodeDRFT   = "DRFT"
	ChequeType2CodeELDR   = "ELDR"
)

var ChequeType2CodeEnumRestriction = []string{ChequeType2CodeCCHQ, ChequeType2CodeCCCH, ChequeType2CodeBCHQ, ChequeType2CodeDRFT, ChequeType2CodeELDR}

// IsValid checks if ChequeType2Code of type String is valid
func (t ChequeType2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), ChequeType2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ChequeType2Code) String() string {
	return string(t)
}

// ToChequeType2Code method for easy conversion with application of restrictions
func ToChequeType2Code(s string) (ChequeType2Code, error) {
	if !isEnumRestrictionValid(s, ChequeType2CodeEnumRestriction) {
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

// ChequeType2CodeExample method for generation of valid sample data
func ChequeType2CodeExample() ChequeType2Code {
	return ChequeType2Code(generateSampleDataWithEnumRestriction(ChequeType2CodeEnumRestriction))

}

/*
 * ExternalProxyAccountType1Code Ops
 */

const (
	ExternalProxyAccountType1CodeSample    = "Value"
	ExternalProxyAccountType1CodeLength    = 0
	ExternalProxyAccountType1CodeMinLength = 1
	ExternalProxyAccountType1CodeMaxLength = 4
)

// IsValid checks if ExternalProxyAccountType1Code of type String is valid
func (t ExternalProxyAccountType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalProxyAccountType1CodeLength, ExternalProxyAccountType1CodeMinLength, ExternalProxyAccountType1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalProxyAccountType1Code) String() string {
	return string(t)
}

// ToExternalProxyAccountType1Code method for easy conversion with application of restrictions
func ToExternalProxyAccountType1Code(s string) (ExternalProxyAccountType1Code, error) {
	if !isLengthRestrictionValid(s, ExternalProxyAccountType1CodeLength, ExternalProxyAccountType1CodeMinLength, ExternalProxyAccountType1CodeMaxLength) {
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

// ExternalProxyAccountType1CodeExample method for generation of valid sample data
func ExternalProxyAccountType1CodeExample() ExternalProxyAccountType1Code {
	return ExternalProxyAccountType1Code(generateSampleDataWithLengthRestriction(ExternalProxyAccountType1CodeLength, ExternalProxyAccountType1CodeMinLength, ExternalProxyAccountType1CodeMaxLength))

}

/*
 * ClearingChannel2Code Ops
 */

const (
	ClearingChannel2CodeSample = "Value"
	ClearingChannel2CodeRTGS   = "RTGS"
	ClearingChannel2CodeRTNS   = "RTNS"
	ClearingChannel2CodeMPNS   = "MPNS"
	ClearingChannel2CodeBOOK   = "BOOK"
)

var ClearingChannel2CodeEnumRestriction = []string{ClearingChannel2CodeRTGS, ClearingChannel2CodeRTNS, ClearingChannel2CodeMPNS, ClearingChannel2CodeBOOK}

// IsValid checks if ClearingChannel2Code of type String is valid
func (t ClearingChannel2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), ClearingChannel2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ClearingChannel2Code) String() string {
	return string(t)
}

// ToClearingChannel2Code method for easy conversion with application of restrictions
func ToClearingChannel2Code(s string) (ClearingChannel2Code, error) {
	if !isEnumRestrictionValid(s, ClearingChannel2CodeEnumRestriction) {
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

// ClearingChannel2CodeExample method for generation of valid sample data
func ClearingChannel2CodeExample() ClearingChannel2Code {
	return ClearingChannel2Code(generateSampleDataWithEnumRestriction(ClearingChannel2CodeEnumRestriction))

}

/*
 * ExternalPaymentGroupStatus1Code Ops
 */

const (
	ExternalPaymentGroupStatus1CodeSample    = "Value"
	ExternalPaymentGroupStatus1CodeLength    = 0
	ExternalPaymentGroupStatus1CodeMinLength = 1
	ExternalPaymentGroupStatus1CodeMaxLength = 4
)

// IsValid checks if ExternalPaymentGroupStatus1Code of type String is valid
func (t ExternalPaymentGroupStatus1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalPaymentGroupStatus1CodeLength, ExternalPaymentGroupStatus1CodeMinLength, ExternalPaymentGroupStatus1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalPaymentGroupStatus1Code) String() string {
	return string(t)
}

// ToExternalPaymentGroupStatus1Code method for easy conversion with application of restrictions
func ToExternalPaymentGroupStatus1Code(s string) (ExternalPaymentGroupStatus1Code, error) {
	if !isLengthRestrictionValid(s, ExternalPaymentGroupStatus1CodeLength, ExternalPaymentGroupStatus1CodeMinLength, ExternalPaymentGroupStatus1CodeMaxLength) {
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

// ExternalPaymentGroupStatus1CodeExample method for generation of valid sample data
func ExternalPaymentGroupStatus1CodeExample() ExternalPaymentGroupStatus1Code {
	return ExternalPaymentGroupStatus1Code(generateSampleDataWithLengthRestriction(ExternalPaymentGroupStatus1CodeLength, ExternalPaymentGroupStatus1CodeMinLength, ExternalPaymentGroupStatus1CodeMaxLength))

}

/*
 * ExternalOrganisationIdentification1Code Ops
 */

const (
	ExternalOrganisationIdentification1CodeSample    = "Value"
	ExternalOrganisationIdentification1CodeLength    = 0
	ExternalOrganisationIdentification1CodeMinLength = 1
	ExternalOrganisationIdentification1CodeMaxLength = 4
)

// IsValid checks if ExternalOrganisationIdentification1Code of type String is valid
func (t ExternalOrganisationIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalOrganisationIdentification1CodeLength, ExternalOrganisationIdentification1CodeMinLength, ExternalOrganisationIdentification1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalOrganisationIdentification1Code) String() string {
	return string(t)
}

// ToExternalOrganisationIdentification1Code method for easy conversion with application of restrictions
func ToExternalOrganisationIdentification1Code(s string) (ExternalOrganisationIdentification1Code, error) {
	if !isLengthRestrictionValid(s, ExternalOrganisationIdentification1CodeLength, ExternalOrganisationIdentification1CodeMinLength, ExternalOrganisationIdentification1CodeMaxLength) {
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

// ExternalOrganisationIdentification1CodeExample method for generation of valid sample data
func ExternalOrganisationIdentification1CodeExample() ExternalOrganisationIdentification1Code {
	return ExternalOrganisationIdentification1Code(generateSampleDataWithLengthRestriction(ExternalOrganisationIdentification1CodeLength, ExternalOrganisationIdentification1CodeMinLength, ExternalOrganisationIdentification1CodeMaxLength))

}

/*
 * ExternalPersonIdentification1Code Ops
 */

const (
	ExternalPersonIdentification1CodeSample    = "Value"
	ExternalPersonIdentification1CodeLength    = 0
	ExternalPersonIdentification1CodeMinLength = 1
	ExternalPersonIdentification1CodeMaxLength = 4
)

// IsValid checks if ExternalPersonIdentification1Code of type String is valid
func (t ExternalPersonIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalPersonIdentification1CodeLength, ExternalPersonIdentification1CodeMinLength, ExternalPersonIdentification1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalPersonIdentification1Code) String() string {
	return string(t)
}

// ToExternalPersonIdentification1Code method for easy conversion with application of restrictions
func ToExternalPersonIdentification1Code(s string) (ExternalPersonIdentification1Code, error) {
	if !isLengthRestrictionValid(s, ExternalPersonIdentification1CodeLength, ExternalPersonIdentification1CodeMinLength, ExternalPersonIdentification1CodeMaxLength) {
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

// ExternalPersonIdentification1CodeExample method for generation of valid sample data
func ExternalPersonIdentification1CodeExample() ExternalPersonIdentification1Code {
	return ExternalPersonIdentification1Code(generateSampleDataWithLengthRestriction(ExternalPersonIdentification1CodeLength, ExternalPersonIdentification1CodeMinLength, ExternalPersonIdentification1CodeMaxLength))

}

/*
 * ExternalPaymentModificationRejection1Code Ops
 */

const (
	ExternalPaymentModificationRejection1CodeSample    = "Value"
	ExternalPaymentModificationRejection1CodeLength    = 0
	ExternalPaymentModificationRejection1CodeMinLength = 1
	ExternalPaymentModificationRejection1CodeMaxLength = 4
)

// IsValid checks if ExternalPaymentModificationRejection1Code of type String is valid
func (t ExternalPaymentModificationRejection1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalPaymentModificationRejection1CodeLength, ExternalPaymentModificationRejection1CodeMinLength, ExternalPaymentModificationRejection1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalPaymentModificationRejection1Code) String() string {
	return string(t)
}

// ToExternalPaymentModificationRejection1Code method for easy conversion with application of restrictions
func ToExternalPaymentModificationRejection1Code(s string) (ExternalPaymentModificationRejection1Code, error) {
	if !isLengthRestrictionValid(s, ExternalPaymentModificationRejection1CodeLength, ExternalPaymentModificationRejection1CodeMinLength, ExternalPaymentModificationRejection1CodeMaxLength) {
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

// ExternalPaymentModificationRejection1CodeExample method for generation of valid sample data
func ExternalPaymentModificationRejection1CodeExample() ExternalPaymentModificationRejection1Code {
	return ExternalPaymentModificationRejection1Code(generateSampleDataWithLengthRestriction(ExternalPaymentModificationRejection1CodeLength, ExternalPaymentModificationRejection1CodeMinLength, ExternalPaymentModificationRejection1CodeMaxLength))

}

/*
 * ExternalLocalInstrument1Code Ops
 */

const (
	ExternalLocalInstrument1CodeSample    = "Value"
	ExternalLocalInstrument1CodeLength    = 0
	ExternalLocalInstrument1CodeMinLength = 1
	ExternalLocalInstrument1CodeMaxLength = 35
)

// IsValid checks if ExternalLocalInstrument1Code of type String is valid
func (t ExternalLocalInstrument1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalLocalInstrument1CodeLength, ExternalLocalInstrument1CodeMinLength, ExternalLocalInstrument1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalLocalInstrument1Code) String() string {
	return string(t)
}

// ToExternalLocalInstrument1Code method for easy conversion with application of restrictions
func ToExternalLocalInstrument1Code(s string) (ExternalLocalInstrument1Code, error) {
	if !isLengthRestrictionValid(s, ExternalLocalInstrument1CodeLength, ExternalLocalInstrument1CodeMinLength, ExternalLocalInstrument1CodeMaxLength) {
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

// ExternalLocalInstrument1CodeExample method for generation of valid sample data
func ExternalLocalInstrument1CodeExample() ExternalLocalInstrument1Code {
	return ExternalLocalInstrument1Code(generateSampleDataWithLengthRestriction(ExternalLocalInstrument1CodeLength, ExternalLocalInstrument1CodeMinLength, ExternalLocalInstrument1CodeMaxLength))

}

/*
 * ExternalCashAccountType1Code Ops
 */

const (
	ExternalCashAccountType1CodeSample    = "Value"
	ExternalCashAccountType1CodeLength    = 0
	ExternalCashAccountType1CodeMinLength = 1
	ExternalCashAccountType1CodeMaxLength = 4
)

// IsValid checks if ExternalCashAccountType1Code of type String is valid
func (t ExternalCashAccountType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), ExternalCashAccountType1CodeLength, ExternalCashAccountType1CodeMinLength, ExternalCashAccountType1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalCashAccountType1Code) String() string {
	return string(t)
}

// ToExternalCashAccountType1Code method for easy conversion with application of restrictions
func ToExternalCashAccountType1Code(s string) (ExternalCashAccountType1Code, error) {
	if !isLengthRestrictionValid(s, ExternalCashAccountType1CodeLength, ExternalCashAccountType1CodeMinLength, ExternalCashAccountType1CodeMaxLength) {
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

// ExternalCashAccountType1CodeExample method for generation of valid sample data
func ExternalCashAccountType1CodeExample() ExternalCashAccountType1Code {
	return ExternalCashAccountType1Code(generateSampleDataWithLengthRestriction(ExternalCashAccountType1CodeLength, ExternalCashAccountType1CodeMinLength, ExternalCashAccountType1CodeMaxLength))

}

/*
 * Max105Text Ops
 */

const (
	Max105TextSample    = "Value"
	Max105TextLength    = 0
	Max105TextMinLength = 1
	Max105TextMaxLength = 105
)

// IsValid checks if Max105Text of type String is valid
func (t Max105Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isLengthRestrictionValid(t.String(), Max105TextLength, Max105TextMinLength, Max105TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max105Text) String() string {
	return string(t)
}

// ToMax105Text method for easy conversion with application of restrictions
func ToMax105Text(s string) (Max105Text, error) {
	if !isLengthRestrictionValid(s, Max105TextLength, Max105TextMinLength, Max105TextMaxLength) {
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

// Max105TextExample method for generation of valid sample data
func Max105TextExample() Max105Text {
	return Max105Text(generateSampleDataWithLengthRestriction(Max105TextLength, Max105TextMinLength, Max105TextMaxLength))

}
