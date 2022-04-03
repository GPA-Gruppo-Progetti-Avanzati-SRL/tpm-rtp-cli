// Package common
// Do not Edit. This stuff it's been automatically generated.
// Generated at 2022-04-04 00:49:38.620577 +0200 CEST m=+0.050820167
package common

// IsValid checks if ClearingSystemMemberIdentification2 is valid
func (s ClearingSystemMemberIdentification2) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.ClrSysId != nil && s.ClrSysId.IsValid(true)) || (s.ClrSysId == nil && true))

	valid = valid && s.MmbId.IsValid(false)

	return valid
}

// IsValid checks if ClearingSystemIdentification2Choice is valid
func (s ClearingSystemIdentification2Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if ActiveOrHistoricCurrencyAndAmount is valid
func (s ActiveOrHistoricCurrencyAndAmount) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Ccy.IsValid(false)
	valid = valid && s.Value.IsValid(false)

	return valid
}

// IsValid checks if GenericIdentification1 is valid
func (s GenericIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && s.SchmeNm.IsValid(true)
	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if FinancialIdentificationSchemeName1Choice is valid
func (s FinancialIdentificationSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if TaxRecord2 is valid
func (s TaxRecord2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp.IsValid(true)
	valid = valid && s.Ctgy.IsValid(true)
	valid = valid && s.CtgyDtls.IsValid(true)
	valid = valid && s.DbtrSts.IsValid(true)
	valid = valid && s.CertId.IsValid(true)
	valid = valid && s.FrmsCd.IsValid(true)
	valid = valid && ((s.Prd != nil && s.Prd.IsValid(true)) || (s.Prd == nil && true))

	valid = valid && ((s.TaxAmt != nil && s.TaxAmt.IsValid(true)) || (s.TaxAmt == nil && true))

	valid = valid && s.AddtlInf.IsValid(true)

	return valid
}

// IsValid checks if TaxAmountType1Choice is valid
func (s TaxAmountType1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if SkipPayload is valid
func (s SkipPayload) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Item.IsValid(false)

	return valid
}

// IsValid checks if AccountSchemeName1Choice is valid
func (s AccountSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if SupplementaryData1 is valid
func (s SupplementaryData1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.PlcAndNm.IsValid(true)
	valid = valid && s.Envlp.IsValid(false)

	return valid
}

// IsValid checks if AmountType4Choice is valid
func (s AmountType4Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.InstdAmt != nil && s.InstdAmt.IsValid(true)) || (s.InstdAmt == nil && true))

	valid = valid && ((s.EqvtAmt != nil && s.EqvtAmt.IsValid(true)) || (s.EqvtAmt == nil && true))

	return valid
}

// IsValid checks if TaxInformation7 is valid
func (s TaxInformation7) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.Cdtr != nil && s.Cdtr.IsValid(true)) || (s.Cdtr == nil && true))

	valid = valid && ((s.Dbtr != nil && s.Dbtr.IsValid(true)) || (s.Dbtr == nil && true))

	valid = valid && ((s.UltmtDbtr != nil && s.UltmtDbtr.IsValid(true)) || (s.UltmtDbtr == nil && true))

	valid = valid && s.AdmstnZone.IsValid(true)
	valid = valid && s.RefNb.IsValid(true)
	valid = valid && s.Mtd.IsValid(true)
	valid = valid && ((s.TtlTaxblBaseAmt != nil && s.TtlTaxblBaseAmt.IsValid(true)) || (s.TtlTaxblBaseAmt == nil && true))

	valid = valid && ((s.TtlTaxAmt != nil && s.TtlTaxAmt.IsValid(true)) || (s.TtlTaxAmt == nil && true))

	valid = valid && s.Dt.IsValid(true)
	valid = valid && s.SeqNb.IsValid(true)
	for j := 0; j < len(s.Rcrd); j++ {
		valid = valid && s.Rcrd[j].IsValid(true)
	}

	return valid
}

// IsValid checks if BranchAndFinancialInstitutionIdentification6 is valid
func (s BranchAndFinancialInstitutionIdentification6) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.FinInstnId.IsValid(false)
	valid = valid && ((s.BrnchId != nil && s.BrnchId.IsValid(true)) || (s.BrnchId == nil && true))

	return valid
}

// IsValid checks if CreditorReferenceType2 is valid
func (s CreditorReferenceType2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.CdOrPrtry.IsValid(false)
	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if ReferredDocumentType3Choice is valid
func (s ReferredDocumentType3Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if GenericIdentification30 is valid
func (s GenericIdentification30) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && s.Issr.IsValid(false)
	valid = valid && s.SchmeNm.IsValid(true)

	return valid
}

// IsValid checks if TaxRecordDetails2 is valid
func (s TaxRecordDetails2) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.Prd != nil && s.Prd.IsValid(true)) || (s.Prd == nil && true))

	valid = valid && s.Amt.IsValid(false)

	return valid
}

// IsValid checks if Document12 is valid
func (s Document12) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp.IsValid(false)
	valid = valid && s.Id.IsValid(false)
	valid = valid && s.IsseDt.IsValid(false)
	valid = valid && s.Nm.IsValid(true)
	valid = valid && s.LangCd.IsValid(true)
	valid = valid && s.Frmt.IsValid(false)
	valid = valid && s.FileNm.IsValid(true)
	valid = valid && ((s.DgtlSgntr != nil && s.DgtlSgntr.IsValid(true)) || (s.DgtlSgntr == nil && true))

	valid = valid && s.Nclsr.IsValid(false)

	return valid
}

// IsValid checks if OrganisationIdentification29 is valid
func (s OrganisationIdentification29) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.AnyBIC.IsValid(true)
	valid = valid && s.LEI.IsValid(true)
	for j := 0; j < len(s.Othr); j++ {
		valid = valid && s.Othr[j].IsValid(true)
	}

	return valid
}

// IsValid checks if LocalInstrument2Choice is valid
func (s LocalInstrument2Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if CreditorReferenceType1Choice is valid
func (s CreditorReferenceType1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if DiscountAmountType1Choice is valid
func (s DiscountAmountType1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if OrganisationIdentificationSchemeName1Choice is valid
func (s OrganisationIdentificationSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if OtherContact1 is valid
func (s OtherContact1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.ChanlTp.IsValid(false)
	valid = valid && s.Id.IsValid(true)

	return valid
}

// IsValid checks if TaxAuthorisation1 is valid
func (s TaxAuthorisation1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Titl.IsValid(true)
	valid = valid && s.Nm.IsValid(true)

	return valid
}

// IsValid checks if TaxPeriod2 is valid
func (s TaxPeriod2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Yr.IsValid(true)
	valid = valid && s.Tp.IsValid(true)
	valid = valid && ((s.FrToDt != nil && s.FrToDt.IsValid(true)) || (s.FrToDt == nil && true))

	return valid
}

// IsValid checks if GenericAccountIdentification1 is valid
func (s GenericAccountIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && ((s.SchmeNm != nil && s.SchmeNm.IsValid(true)) || (s.SchmeNm == nil && true))

	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if ProxyAccountType1Choice is valid
func (s ProxyAccountType1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if ReferredDocumentInformation7 is valid
func (s ReferredDocumentInformation7) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.Tp != nil && s.Tp.IsValid(true)) || (s.Tp == nil && true))

	valid = valid && s.Nb.IsValid(true)
	valid = valid && s.RltdDt.IsValid(true)
	for j := 0; j < len(s.LineDtls); j++ {
		valid = valid && s.LineDtls[j].IsValid(true)
	}

	return valid
}

// IsValid checks if GarnishmentType1 is valid
func (s GarnishmentType1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.CdOrPrtry.IsValid(false)
	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if GenericOrganisationIdentification1 is valid
func (s GenericOrganisationIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && ((s.SchmeNm != nil && s.SchmeNm.IsValid(true)) || (s.SchmeNm == nil && true))

	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if PersonIdentification13 is valid
func (s PersonIdentification13) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.DtAndPlcOfBirth != nil && s.DtAndPlcOfBirth.IsValid(true)) || (s.DtAndPlcOfBirth == nil && true))

	for j := 0; j < len(s.Othr); j++ {
		valid = valid && s.Othr[j].IsValid(true)
	}

	return valid
}

// IsValid checks if TaxAmountAndType1 is valid
func (s TaxAmountAndType1) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.Tp != nil && s.Tp.IsValid(true)) || (s.Tp == nil && true))

	valid = valid && s.Amt.IsValid(false)

	return valid
}

// IsValid checks if PartyAndSignature3 is valid
func (s PartyAndSignature3) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Pty.IsValid(false)
	valid = valid && s.Sgntr.IsValid(false)

	return valid
}

// IsValid checks if Contact4 is valid
func (s Contact4) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.NmPrfx.IsValid(true)
	valid = valid && s.Nm.IsValid(true)
	valid = valid && s.PhneNb.IsValid(true)
	valid = valid && s.MobNb.IsValid(true)
	valid = valid && s.FaxNb.IsValid(true)
	valid = valid && s.EmailAdr.IsValid(true)
	valid = valid && s.EmailPurp.IsValid(true)
	valid = valid && s.JobTitl.IsValid(true)
	valid = valid && s.Rspnsblty.IsValid(true)
	valid = valid && s.Dept.IsValid(true)
	for j := 0; j < len(s.Othr); j++ {
		valid = valid && s.Othr[j].IsValid(true)
	}

	valid = valid && s.PrefrdMtd.IsValid(true)

	return valid
}

// IsValid checks if Garnishment3 is valid
func (s Garnishment3) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp.IsValid(false)
	valid = valid && ((s.Grnshee != nil && s.Grnshee.IsValid(true)) || (s.Grnshee == nil && true))

	valid = valid && ((s.GrnshmtAdmstr != nil && s.GrnshmtAdmstr.IsValid(true)) || (s.GrnshmtAdmstr == nil && true))

	valid = valid && s.RefNb.IsValid(true)
	valid = valid && s.Dt.IsValid(true)
	valid = valid && ((s.RmtdAmt != nil && s.RmtdAmt.IsValid(true)) || (s.RmtdAmt == nil && true))

	valid = valid && s.FmlyMdclInsrncInd.IsValid(true)
	valid = valid && s.MplyeeTermntnInd.IsValid(true)

	return valid
}

// IsValid checks if GarnishmentType1Choice is valid
func (s GarnishmentType1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if ReferredDocumentType4 is valid
func (s ReferredDocumentType4) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.CdOrPrtry.IsValid(false)
	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if ActiveCurrencyAndAmount is valid
func (s ActiveCurrencyAndAmount) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Ccy.IsValid(false)
	valid = valid && s.Value.IsValid(false)

	return valid
}

// IsValid checks if CashAccountType2Choice is valid
func (s CashAccountType2Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if BranchData3 is valid
func (s BranchData3) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(true)
	valid = valid && s.LEI.IsValid(true)
	valid = valid && s.Nm.IsValid(true)
	valid = valid && ((s.PstlAdr != nil && s.PstlAdr.IsValid(true)) || (s.PstlAdr == nil && true))

	return valid
}

// IsValid checks if TaxParty2 is valid
func (s TaxParty2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.TaxId.IsValid(true)
	valid = valid && s.RegnId.IsValid(true)
	valid = valid && s.TaxTp.IsValid(true)
	valid = valid && ((s.Authstn != nil && s.Authstn.IsValid(true)) || (s.Authstn == nil && true))

	return valid
}

// IsValid checks if Party38Choice is valid
func (s Party38Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.OrgId != nil && s.OrgId.IsValid(true)) || (s.OrgId == nil && true))

	valid = valid && ((s.PrvtId != nil && s.PrvtId.IsValid(true)) || (s.PrvtId == nil && true))

	return valid
}

// IsValid checks if EquivalentAmount2 is valid
func (s EquivalentAmount2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Amt.IsValid(false)
	valid = valid && s.CcyOfTrf.IsValid(false)

	return valid
}

// IsValid checks if PartyIdentification135 is valid
func (s PartyIdentification135) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Nm.IsValid(true)
	valid = valid && ((s.PstlAdr != nil && s.PstlAdr.IsValid(true)) || (s.PstlAdr == nil && true))

	valid = valid && ((s.Id != nil && s.Id.IsValid(true)) || (s.Id == nil && true))

	valid = valid && s.CtryOfRes.IsValid(true)
	valid = valid && ((s.CtctDtls != nil && s.CtctDtls.IsValid(true)) || (s.CtctDtls == nil && true))

	return valid
}

// IsValid checks if StructuredRemittanceInformation16 is valid
func (s StructuredRemittanceInformation16) IsValid(optional bool) bool {

	valid := true
	for j := 0; j < len(s.RfrdDocInf); j++ {
		valid = valid && s.RfrdDocInf[j].IsValid(true)
	}

	valid = valid && ((s.RfrdDocAmt != nil && s.RfrdDocAmt.IsValid(true)) || (s.RfrdDocAmt == nil && true))

	valid = valid && ((s.CdtrRefInf != nil && s.CdtrRefInf.IsValid(true)) || (s.CdtrRefInf == nil && true))

	valid = valid && ((s.Invcr != nil && s.Invcr.IsValid(true)) || (s.Invcr == nil && true))

	valid = valid && ((s.Invcee != nil && s.Invcee.IsValid(true)) || (s.Invcee == nil && true))

	valid = valid && ((s.TaxRmt != nil && s.TaxRmt.IsValid(true)) || (s.TaxRmt == nil && true))

	valid = valid && ((s.GrnshmtRmt != nil && s.GrnshmtRmt.IsValid(true)) || (s.GrnshmtRmt == nil && true))

	for j := 0; j < len(s.AddtlRmtInf); j++ {
		valid = valid && s.AddtlRmtInf[j].IsValid(true)
	}

	return valid
}

// IsValid checks if DocumentLineType1Choice is valid
func (s DocumentLineType1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if CashAccount38 is valid
func (s CashAccount38) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && ((s.Tp != nil && s.Tp.IsValid(true)) || (s.Tp == nil && true))

	valid = valid && s.Ccy.IsValid(true)
	valid = valid && s.Nm.IsValid(true)
	valid = valid && ((s.Prxy != nil && s.Prxy.IsValid(true)) || (s.Prxy == nil && true))

	return valid
}

// IsValid checks if AccountIdentification4Choice is valid
func (s AccountIdentification4Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.IBAN.IsValid(true)
	valid = valid && ((s.Othr != nil && s.Othr.IsValid(true)) || (s.Othr == nil && true))

	return valid
}

// IsValid checks if ProxyAccountIdentification1 is valid
func (s ProxyAccountIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.Tp != nil && s.Tp.IsValid(true)) || (s.Tp == nil && true))

	valid = valid && s.Id.IsValid(false)

	return valid
}

// IsValid checks if DocumentLineIdentification1 is valid
func (s DocumentLineIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.Tp != nil && s.Tp.IsValid(true)) || (s.Tp == nil && true))

	valid = valid && s.Nb.IsValid(true)
	valid = valid && s.RltdDt.IsValid(true)

	return valid
}

// IsValid checks if PostalAddress24 is valid
func (s PostalAddress24) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.AdrTp != nil && s.AdrTp.IsValid(true)) || (s.AdrTp == nil && true))

	valid = valid && s.Dept.IsValid(true)
	valid = valid && s.SubDept.IsValid(true)
	valid = valid && s.StrtNm.IsValid(true)
	valid = valid && s.BldgNb.IsValid(true)
	valid = valid && s.BldgNm.IsValid(true)
	valid = valid && s.Flr.IsValid(true)
	valid = valid && s.PstBx.IsValid(true)
	valid = valid && s.Room.IsValid(true)
	valid = valid && s.PstCd.IsValid(true)
	valid = valid && s.TwnNm.IsValid(true)
	valid = valid && s.TwnLctnNm.IsValid(true)
	valid = valid && s.DstrctNm.IsValid(true)
	valid = valid && s.CtrySubDvsn.IsValid(true)
	valid = valid && s.Ctry.IsValid(true)
	for j := 0; j < len(s.AdrLine); j++ {
		valid = valid && s.AdrLine[j].IsValid(true)
	}

	return valid
}

// IsValid checks if DateAndPlaceOfBirth1 is valid
func (s DateAndPlaceOfBirth1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.BirthDt.IsValid(false)
	valid = valid && s.PrvcOfBirth.IsValid(true)
	valid = valid && s.CityOfBirth.IsValid(false)
	valid = valid && s.CtryOfBirth.IsValid(false)

	return valid
}

// IsValid checks if CategoryPurpose1Choice is valid
func (s CategoryPurpose1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if RemittanceAmount2 is valid
func (s RemittanceAmount2) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.DuePyblAmt != nil && s.DuePyblAmt.IsValid(true)) || (s.DuePyblAmt == nil && true))

	for j := 0; j < len(s.DscntApldAmt); j++ {
		valid = valid && s.DscntApldAmt[j].IsValid(true)
	}

	valid = valid && ((s.CdtNoteAmt != nil && s.CdtNoteAmt.IsValid(true)) || (s.CdtNoteAmt == nil && true))

	for j := 0; j < len(s.TaxAmt); j++ {
		valid = valid && s.TaxAmt[j].IsValid(true)
	}

	for j := 0; j < len(s.AdjstmntAmtAndRsn); j++ {
		valid = valid && s.AdjstmntAmtAndRsn[j].IsValid(true)
	}

	valid = valid && ((s.RmtdAmt != nil && s.RmtdAmt.IsValid(true)) || (s.RmtdAmt == nil && true))

	return valid
}

// IsValid checks if CreditorReferenceInformation2 is valid
func (s CreditorReferenceInformation2) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.Tp != nil && s.Tp.IsValid(true)) || (s.Tp == nil && true))

	valid = valid && s.Ref.IsValid(true)

	return valid
}

// IsValid checks if DocumentType1Choice is valid
func (s DocumentType1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && ((s.Prtry != nil && s.Prtry.IsValid(true)) || (s.Prtry == nil && true))

	return valid
}

// IsValid checks if PaymentTypeInformation26 is valid
func (s PaymentTypeInformation26) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.InstrPrty.IsValid(true)
	for j := 0; j < len(s.SvcLvl); j++ {
		valid = valid && s.SvcLvl[j].IsValid(true)
	}

	valid = valid && ((s.LclInstrm != nil && s.LclInstrm.IsValid(true)) || (s.LclInstrm == nil && true))

	valid = valid && ((s.CtgyPurp != nil && s.CtgyPurp.IsValid(true)) || (s.CtgyPurp == nil && true))

	return valid
}

// IsValid checks if DocumentFormat1Choice is valid
func (s DocumentFormat1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && ((s.Prtry != nil && s.Prtry.IsValid(true)) || (s.Prtry == nil && true))

	return valid
}

// IsValid checks if AmountOrRate1Choice is valid
func (s AmountOrRate1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.Amt != nil && s.Amt.IsValid(true)) || (s.Amt == nil && true))

	valid = valid && s.Rate.IsValid(true)

	return valid
}

// IsValid checks if PaymentCondition1 is valid
func (s PaymentCondition1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.AmtModAllwd.IsValid(false)
	valid = valid && s.EarlyPmtAllwd.IsValid(false)
	valid = valid && s.DelyPnlty.IsValid(true)
	valid = valid && ((s.ImdtPmtRbt != nil && s.ImdtPmtRbt.IsValid(true)) || (s.ImdtPmtRbt == nil && true))

	valid = valid && s.GrntedPmtReqd.IsValid(false)

	return valid
}

// IsValid checks if DocumentLineInformation1 is valid
func (s DocumentLineInformation1) IsValid(optional bool) bool {

	valid := true
	if len(s.Id) == 0 {
		valid = false
	}
	for j := 0; j < len(s.Id); j++ {
		valid = valid && s.Id[j].IsValid(false)
	}

	valid = valid && s.Desc.IsValid(true)
	valid = valid && ((s.Amt != nil && s.Amt.IsValid(true)) || (s.Amt == nil && true))

	return valid
}

// IsValid checks if DiscountAmountAndType1 is valid
func (s DiscountAmountAndType1) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.Tp != nil && s.Tp.IsValid(true)) || (s.Tp == nil && true))

	valid = valid && s.Amt.IsValid(false)

	return valid
}

// IsValid checks if DocumentAdjustment1 is valid
func (s DocumentAdjustment1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Amt.IsValid(false)
	valid = valid && s.CdtDbtInd.IsValid(true)
	valid = valid && s.Rsn.IsValid(true)
	valid = valid && s.AddtlInf.IsValid(true)

	return valid
}

// IsValid checks if FinancialInstitutionIdentification18 is valid
func (s FinancialInstitutionIdentification18) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.BICFI.IsValid(true)
	valid = valid && ((s.ClrSysMmbId != nil && s.ClrSysMmbId.IsValid(true)) || (s.ClrSysMmbId == nil && true))

	valid = valid && s.LEI.IsValid(true)
	valid = valid && s.Nm.IsValid(true)
	valid = valid && ((s.PstlAdr != nil && s.PstlAdr.IsValid(true)) || (s.PstlAdr == nil && true))

	valid = valid && ((s.Othr != nil && s.Othr.IsValid(true)) || (s.Othr == nil && true))

	return valid
}

// IsValid checks if DocumentLineType1 is valid
func (s DocumentLineType1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.CdOrPrtry.IsValid(false)
	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if SupplementaryDataEnvelope1 is valid
func (s SupplementaryDataEnvelope1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Item.IsValid(false)

	return valid
}

// IsValid checks if DatePeriod2 is valid
func (s DatePeriod2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.FrDt.IsValid(false)
	valid = valid && s.ToDt.IsValid(false)

	return valid
}

// IsValid checks if AddressType3Choice is valid
func (s AddressType3Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && ((s.Prtry != nil && s.Prtry.IsValid(true)) || (s.Prtry == nil && true))

	return valid
}

// IsValid checks if GenericPersonIdentification1 is valid
func (s GenericPersonIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && ((s.SchmeNm != nil && s.SchmeNm.IsValid(true)) || (s.SchmeNm == nil && true))

	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if ServiceLevel8Choice is valid
func (s ServiceLevel8Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if DateAndDateTime2Choice is valid
func (s DateAndDateTime2Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Dt.IsValid(true)
	valid = valid && s.DtTm.IsValid(true)

	return valid
}

// IsValid checks if GenericFinancialIdentification1 is valid
func (s GenericFinancialIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && ((s.SchmeNm != nil && s.SchmeNm.IsValid(true)) || (s.SchmeNm == nil && true))

	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if TaxAmount2 is valid
func (s TaxAmount2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Rate.IsValid(true)
	valid = valid && ((s.TaxblBaseAmt != nil && s.TaxblBaseAmt.IsValid(true)) || (s.TaxblBaseAmt == nil && true))

	valid = valid && ((s.TtlAmt != nil && s.TtlAmt.IsValid(true)) || (s.TtlAmt == nil && true))

	for j := 0; j < len(s.Dtls); j++ {
		valid = valid && s.Dtls[j].IsValid(true)
	}

	return valid
}

// IsValid checks if RemittanceInformation16 is valid
func (s RemittanceInformation16) IsValid(optional bool) bool {

	valid := true
	for j := 0; j < len(s.Ustrd); j++ {
		valid = valid && s.Ustrd[j].IsValid(true)
	}

	for j := 0; j < len(s.Strd); j++ {
		valid = valid && s.Strd[j].IsValid(true)
	}

	return valid
}

// IsValid checks if RemittanceAmount3 is valid
func (s RemittanceAmount3) IsValid(optional bool) bool {

	valid := true
	valid = valid && ((s.DuePyblAmt != nil && s.DuePyblAmt.IsValid(true)) || (s.DuePyblAmt == nil && true))

	for j := 0; j < len(s.DscntApldAmt); j++ {
		valid = valid && s.DscntApldAmt[j].IsValid(true)
	}

	valid = valid && ((s.CdtNoteAmt != nil && s.CdtNoteAmt.IsValid(true)) || (s.CdtNoteAmt == nil && true))

	for j := 0; j < len(s.TaxAmt); j++ {
		valid = valid && s.TaxAmt[j].IsValid(true)
	}

	for j := 0; j < len(s.AdjstmntAmtAndRsn); j++ {
		valid = valid && s.AdjstmntAmtAndRsn[j].IsValid(true)
	}

	valid = valid && ((s.RmtdAmt != nil && s.RmtdAmt.IsValid(true)) || (s.RmtdAmt == nil && true))

	return valid
}

// IsValid checks if TaxParty1 is valid
func (s TaxParty1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.TaxId.IsValid(true)
	valid = valid && s.RegnId.IsValid(true)
	valid = valid && s.TaxTp.IsValid(true)

	return valid
}

// IsValid checks if PersonIdentificationSchemeName1Choice is valid
func (s PersonIdentificationSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}
