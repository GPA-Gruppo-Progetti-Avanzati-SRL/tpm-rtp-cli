// Package common
// Do not Edit. This stuff it's been automatically generated.
// Generated at 2022-04-04 00:49:38.620577 +0200 CEST m=+0.050820167
package common

import (
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
	"regexp"
)

// DocumentType6Code funcs and properties -----------------

var DocumentType6CodeEnumRestriction = []string{"MSIN", "CNFA", "DNFA", "CINV", "CREN", "DEBN", "HIRI", "SBIN", "CMCN", "SOAC", "DISP", "BOLD", "VCHR", "AROI", "TSUT", "PUOR"}

// IsValid checks if DocumentType6Code of type String is valid
func (t DocumentType6Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), DocumentType6CodeEnumRestriction)

	return valid
}

// Exact4AlphaNumericText funcs and properties -----------------

var Exact4AlphaNumericTextPatternRestriction = regexp.MustCompile(`[a-zA-Z0-9]{4}`)

// IsValid checks if Exact4AlphaNumericText of type String is valid
func (t Exact4AlphaNumericText) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(fmt.Sprintf("%v", t), Exact4AlphaNumericTextPatternRestriction)

	return valid
}

// AnyBICDec2014Identifier funcs and properties -----------------

var AnyBICDec2014IdentifierPatternRestriction = regexp.MustCompile(`[A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}`)

// IsValid checks if AnyBICDec2014Identifier of type String is valid
func (t AnyBICDec2014Identifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(fmt.Sprintf("%v", t), AnyBICDec2014IdentifierPatternRestriction)

	return valid
}

// AddressType2Code funcs and properties -----------------

var AddressType2CodeEnumRestriction = []string{"ADDR", "PBOX", "HOME", "BIZZ", "MLTO", "DLVY"}

// IsValid checks if AddressType2Code of type String is valid
func (t AddressType2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), AddressType2CodeEnumRestriction)

	return valid
}

// Max128Text funcs and properties -----------------

// IsValid checks if Max128Text of type String is valid
func (t Max128Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 128)

	return valid
}

// Max10MbBinary funcs and properties -----------------

// IsValid checks if Max10MbBinary of type Base64Binary is valid
func (t Max10MbBinary) IsValid(optional bool) bool {

	valid := xsdt.Base64Binary(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 10485760)

	return valid
}

// DocumentType3Code funcs and properties -----------------

var DocumentType3CodeEnumRestriction = []string{"RADM", "RPIN", "FXDR", "DISP", "PUOR", "SCOR"}

// IsValid checks if DocumentType3Code of type String is valid
func (t DocumentType3Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), DocumentType3CodeEnumRestriction)

	return valid
}

// RegulatoryReportingType1Code funcs and properties -----------------

var RegulatoryReportingType1CodeEnumRestriction = []string{"CRED", "DEBT", "BOTH"}

// IsValid checks if RegulatoryReportingType1Code of type String is valid
func (t RegulatoryReportingType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), RegulatoryReportingType1CodeEnumRestriction)

	return valid
}

// ChequeDelivery1Code funcs and properties -----------------

var ChequeDelivery1CodeEnumRestriction = []string{"MLDB", "MLCD", "MLFA", "CRDB", "CRCD", "CRFA", "PUDB", "PUCD", "PUFA", "RGDB", "RGCD", "RGFA"}

// IsValid checks if ChequeDelivery1Code of type String is valid
func (t ChequeDelivery1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), ChequeDelivery1CodeEnumRestriction)

	return valid
}

// Instruction3Code funcs and properties -----------------

var Instruction3CodeEnumRestriction = []string{"CHQB", "HOLD", "PHOB", "TELB"}

// IsValid checks if Instruction3Code of type String is valid
func (t Instruction3Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), Instruction3CodeEnumRestriction)

	return valid
}

// ActiveOrHistoricCurrencyCode funcs and properties -----------------

var ActiveOrHistoricCurrencyCodePatternRestriction = regexp.MustCompile(`[A-Z]{3,3}`)

// IsValid checks if ActiveOrHistoricCurrencyCode of type String is valid
func (t ActiveOrHistoricCurrencyCode) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(fmt.Sprintf("%v", t), ActiveOrHistoricCurrencyCodePatternRestriction)

	return valid
}

// ExternalAccountIdentification1Code funcs and properties -----------------

// IsValid checks if ExternalAccountIdentification1Code of type String is valid
func (t ExternalAccountIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// Max350Text funcs and properties -----------------

// IsValid checks if Max350Text of type String is valid
func (t Max350Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 350)

	return valid
}

// ExternalPaymentGroupStatus1Code funcs and properties -----------------

// IsValid checks if ExternalPaymentGroupStatus1Code of type String is valid
func (t ExternalPaymentGroupStatus1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// PaymentMethod4Code funcs and properties -----------------

var PaymentMethod4CodeEnumRestriction = []string{"CHK", "TRF", "DD", "TRA"}

// IsValid checks if PaymentMethod4Code of type String is valid
func (t PaymentMethod4Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), PaymentMethod4CodeEnumRestriction)

	return valid
}

// ISODateTime funcs and properties -----------------

// IsValid checks if ISODateTime of type DateTime is valid
func (t ISODateTime) IsValid(optional bool) bool {

	valid := xsdt.DateTime(t).IsValid(optional)
	return valid
}

// ISODate funcs and properties -----------------

// IsValid checks if ISODate of type Date is valid
func (t ISODate) IsValid(optional bool) bool {

	valid := xsdt.Date(t).IsValid(optional)
	return valid
}

// PhoneNumber funcs and properties -----------------

var PhoneNumberPatternRestriction = regexp.MustCompile(`\+[0-9]{1,3}-[0-9()+\-]{1,30}`)

// IsValid checks if PhoneNumber of type String is valid
func (t PhoneNumber) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(fmt.Sprintf("%v", t), PhoneNumberPatternRestriction)

	return valid
}

// ExternalProxyAccountType1Code funcs and properties -----------------

// IsValid checks if ExternalProxyAccountType1Code of type String is valid
func (t ExternalProxyAccountType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// Max15NumericText funcs and properties -----------------

var Max15NumericTextPatternRestriction = regexp.MustCompile(`[0-9]{1,15}`)

// IsValid checks if Max15NumericText of type String is valid
func (t Max15NumericText) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(fmt.Sprintf("%v", t), Max15NumericTextPatternRestriction)

	return valid
}

// Max16Text funcs and properties -----------------

// IsValid checks if Max16Text of type String is valid
func (t Max16Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 16)

	return valid
}

// ExternalServiceLevel1Code funcs and properties -----------------

// IsValid checks if ExternalServiceLevel1Code of type String is valid
func (t ExternalServiceLevel1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// ExternalCategoryPurpose1Code funcs and properties -----------------

// IsValid checks if ExternalCategoryPurpose1Code of type String is valid
func (t ExternalCategoryPurpose1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// TaxRecordPeriod1Code funcs and properties -----------------

var TaxRecordPeriod1CodeEnumRestriction = []string{"MM01", "MM02", "MM03", "MM04", "MM05", "MM06", "MM07", "MM08", "MM09", "MM10", "MM11", "MM12", "QTR1", "QTR2", "QTR3", "QTR4", "HLF1", "HLF2"}

// IsValid checks if TaxRecordPeriod1Code of type String is valid
func (t TaxRecordPeriod1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), TaxRecordPeriod1CodeEnumRestriction)

	return valid
}

// ExternalTaxAmountType1Code funcs and properties -----------------

// IsValid checks if ExternalTaxAmountType1Code of type String is valid
func (t ExternalTaxAmountType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// ExternalStatusReason1Code funcs and properties -----------------

// IsValid checks if ExternalStatusReason1Code of type String is valid
func (t ExternalStatusReason1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// Max2048Text funcs and properties -----------------

// IsValid checks if Max2048Text of type String is valid
func (t Max2048Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 2048)

	return valid
}

// ChequeType2Code funcs and properties -----------------

var ChequeType2CodeEnumRestriction = []string{"CCHQ", "CCCH", "BCHQ", "DRFT", "ELDR"}

// IsValid checks if ChequeType2Code of type String is valid
func (t ChequeType2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), ChequeType2CodeEnumRestriction)

	return valid
}

// ExternalDocumentLineType1Code funcs and properties -----------------

// IsValid checks if ExternalDocumentLineType1Code of type String is valid
func (t ExternalDocumentLineType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// CreditDebitCode funcs and properties -----------------

var CreditDebitCodeEnumRestriction = []string{"CRDT", "DBIT"}

// IsValid checks if CreditDebitCode of type String is valid
func (t CreditDebitCode) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), CreditDebitCodeEnumRestriction)

	return valid
}

// Max70Text funcs and properties -----------------

// IsValid checks if Max70Text of type String is valid
func (t Max70Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 70)

	return valid
}

// ChargeBearerType1Code funcs and properties -----------------

var ChargeBearerType1CodeEnumRestriction = []string{"DEBT", "CRED", "SHAR", "SLEV"}

// IsValid checks if ChargeBearerType1Code of type String is valid
func (t ChargeBearerType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), ChargeBearerType1CodeEnumRestriction)

	return valid
}

// ExternalDocumentFormat1Code funcs and properties -----------------

// IsValid checks if ExternalDocumentFormat1Code of type String is valid
func (t ExternalDocumentFormat1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// Max140Text funcs and properties -----------------

// IsValid checks if Max140Text of type String is valid
func (t Max140Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 140)

	return valid
}

// PaymentMethod7Code funcs and properties -----------------

var PaymentMethod7CodeEnumRestriction = []string{"CHK", "TRF"}

// IsValid checks if PaymentMethod7Code of type String is valid
func (t PaymentMethod7Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), PaymentMethod7CodeEnumRestriction)

	return valid
}

// NamePrefix2Code funcs and properties -----------------

var NamePrefix2CodeEnumRestriction = []string{"DOCT", "MADM", "MISS", "MIST", "MIKS"}

// IsValid checks if NamePrefix2Code of type String is valid
func (t NamePrefix2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), NamePrefix2CodeEnumRestriction)

	return valid
}

// Max10Text funcs and properties -----------------

// IsValid checks if Max10Text of type String is valid
func (t Max10Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 10)

	return valid
}

// ExternalDiscountAmountType1Code funcs and properties -----------------

// IsValid checks if ExternalDiscountAmountType1Code of type String is valid
func (t ExternalDiscountAmountType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// ExternalLocalInstrument1Code funcs and properties -----------------

// IsValid checks if ExternalLocalInstrument1Code of type String is valid
func (t ExternalLocalInstrument1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 35)

	return valid
}

// RemittanceLocationMethod2Code funcs and properties -----------------

var RemittanceLocationMethod2CodeEnumRestriction = []string{"FAXI", "EDIC", "URID", "EMAL", "POST", "SMSM"}

// IsValid checks if RemittanceLocationMethod2Code of type String is valid
func (t RemittanceLocationMethod2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), RemittanceLocationMethod2CodeEnumRestriction)

	return valid
}

// ActiveCurrencyCode funcs and properties -----------------

var ActiveCurrencyCodePatternRestriction = regexp.MustCompile(`[A-Z]{3,3}`)

// IsValid checks if ActiveCurrencyCode of type String is valid
func (t ActiveCurrencyCode) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(fmt.Sprintf("%v", t), ActiveCurrencyCodePatternRestriction)

	return valid
}

// ExternalClearingSystemIdentification1Code funcs and properties -----------------

// IsValid checks if ExternalClearingSystemIdentification1Code of type String is valid
func (t ExternalClearingSystemIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 5)

	return valid
}

// UUIDv4Identifier funcs and properties -----------------

var UUIDv4IdentifierPatternRestriction = regexp.MustCompile(`[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}`)

// IsValid checks if UUIDv4Identifier of type String is valid
func (t UUIDv4Identifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(fmt.Sprintf("%v", t), UUIDv4IdentifierPatternRestriction)

	return valid
}

// ExternalGarnishmentType1Code funcs and properties -----------------

// IsValid checks if ExternalGarnishmentType1Code of type String is valid
func (t ExternalGarnishmentType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// ExternalOrganisationIdentification1Code funcs and properties -----------------

// IsValid checks if ExternalOrganisationIdentification1Code of type String is valid
func (t ExternalOrganisationIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// Max4Text funcs and properties -----------------

// IsValid checks if Max4Text of type String is valid
func (t Max4Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// BICFIDec2014Identifier funcs and properties -----------------

var BICFIDec2014IdentifierPatternRestriction = regexp.MustCompile(`[A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}`)

// IsValid checks if BICFIDec2014Identifier of type String is valid
func (t BICFIDec2014Identifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(fmt.Sprintf("%v", t), BICFIDec2014IdentifierPatternRestriction)

	return valid
}

// ExternalFinancialInstitutionIdentification1Code funcs and properties -----------------

// IsValid checks if ExternalFinancialInstitutionIdentification1Code of type String is valid
func (t ExternalFinancialInstitutionIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// ExternalPaymentTransactionStatus1Code funcs and properties -----------------

// IsValid checks if ExternalPaymentTransactionStatus1Code of type String is valid
func (t ExternalPaymentTransactionStatus1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// IBAN2007Identifier funcs and properties -----------------

var IBAN2007IdentifierPatternRestriction = regexp.MustCompile(`[A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}`)

// IsValid checks if IBAN2007Identifier of type String is valid
func (t IBAN2007Identifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(fmt.Sprintf("%v", t), IBAN2007IdentifierPatternRestriction)

	return valid
}

// Max34Text funcs and properties -----------------

// IsValid checks if Max34Text of type String is valid
func (t Max34Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 34)

	return valid
}

// ExternalCashAccountType1Code funcs and properties -----------------

// IsValid checks if ExternalCashAccountType1Code of type String is valid
func (t ExternalCashAccountType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// ExternalDocumentType1Code funcs and properties -----------------

// IsValid checks if ExternalDocumentType1Code of type String is valid
func (t ExternalDocumentType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// CountryCode funcs and properties -----------------

var CountryCodePatternRestriction = regexp.MustCompile(`[A-Z]{2,2}`)

// IsValid checks if CountryCode of type String is valid
func (t CountryCode) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(fmt.Sprintf("%v", t), CountryCodePatternRestriction)

	return valid
}

// ExternalPurpose1Code funcs and properties -----------------

// IsValid checks if ExternalPurpose1Code of type String is valid
func (t ExternalPurpose1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// LEIIdentifier funcs and properties -----------------

var LEIIdentifierPatternRestriction = regexp.MustCompile(`[A-Z0-9]{18,18}[0-9]{2,2}`)

// IsValid checks if LEIIdentifier of type String is valid
func (t LEIIdentifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyPatternRestriction(fmt.Sprintf("%v", t), LEIIdentifierPatternRestriction)

	return valid
}

// PreferredContactMethod1Code funcs and properties -----------------

var PreferredContactMethod1CodeEnumRestriction = []string{"LETT", "MAIL", "PHON", "FAXX", "CELL"}

// IsValid checks if PreferredContactMethod1Code of type String is valid
func (t PreferredContactMethod1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), PreferredContactMethod1CodeEnumRestriction)

	return valid
}

// Priority2Code funcs and properties -----------------

var Priority2CodeEnumRestriction = []string{"HIGH", "NORM"}

// IsValid checks if Priority2Code of type String is valid
func (t Priority2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyEnumRestriction(fmt.Sprintf("%v", t), Priority2CodeEnumRestriction)

	return valid
}

// Max35Text funcs and properties -----------------

// IsValid checks if Max35Text of type String is valid
func (t Max35Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 35)

	return valid
}

// ExternalPersonIdentification1Code funcs and properties -----------------

// IsValid checks if ExternalPersonIdentification1Code of type String is valid
func (t ExternalPersonIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 4)

	return valid
}

// Max105Text funcs and properties -----------------

// IsValid checks if Max105Text of type String is valid
func (t Max105Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && applyLengthRestriction(fmt.Sprintf("%v", t), 0, 1, 105)

	return valid
}
