// Package pain_013_001_07
// Do not Edit. This stuff it's been automatically generated.
package pain_013_001_07

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
)

func (d *Document) With_CdtrPmtActvtnReq(_cdtrPmtActvtnReq CreditorPaymentActivationRequestV07) {
	d.CdtrPmtActvtnReq = _cdtrPmtActvtnReq
}

func (d *Document) With_CdtrPmtActvtnReq_GrpHdr(_grpHdr GroupHeader78) {
	d.CdtrPmtActvtnReq.GrpHdr = _grpHdr
}

func (d *Document) With_CdtrPmtActvtnReq_GrpHdr_InitgPty(_initgPty common.PartyIdentification135) {
	d.CdtrPmtActvtnReq.GrpHdr.InitgPty = _initgPty
}

func (d *Document) With_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr(_pstlAdr *common.PostalAddress24) {
	d.CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id(_id *common.Party38Choice) {
	d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id = _id
}

func (d *Document) With_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId(_orgId *common.OrganisationIdentification29) {
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId = _orgId
}

func (d *Document) With_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr(_othr []common.GenericOrganisationIdentification1) {
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId = &common.OrganisationIdentification29{}
	}
	d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm(_schmeNm *common.OrganisationIdentificationSchemeName1Choice) {
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId = &common.OrganisationIdentification29{}
	}
	if len(d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr) == 0 {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr = append(d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr, common.GenericOrganisationIdentification1{})
	}
	d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId(_prvtId *common.PersonIdentification13) {
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId = _prvtId
}

func (d *Document) With_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth(_dtAndPlcOfBirth *common.DateAndPlaceOfBirth1) {
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth = _dtAndPlcOfBirth
}

func (d *Document) With_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr(_othr []common.GenericPersonIdentification1) {
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm(_schmeNm *common.PersonIdentificationSchemeName1Choice) {
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId = &common.PersonIdentification13{}
	}
	if len(d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.Othr) == 0 {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.Othr = append(d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.Othr, common.GenericPersonIdentification1{})
	}
	d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls(_ctctDtls *common.Contact4) {
	d.CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls = _ctctDtls
}

func (d *Document) With_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Othr(_othr []common.OtherContact1) {
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls = &common.Contact4{}
	}
	d.CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf(_pmtInf []PaymentInstruction31) {
	d.CdtrPmtActvtnReq.PmtInf = _pmtInf
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_PmtTpInf(_pmtTpInf *common.PaymentTypeInformation26) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].PmtTpInf = _pmtTpInf
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_PmtTpInf_SvcLvl(_svcLvl []common.ServiceLevel8Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].PmtTpInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].PmtTpInf = &common.PaymentTypeInformation26{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].PmtTpInf.SvcLvl = _svcLvl
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_PmtTpInf_LclInstrm(_lclInstrm *common.LocalInstrument2Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].PmtTpInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].PmtTpInf = &common.PaymentTypeInformation26{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].PmtTpInf.LclInstrm = _lclInstrm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_PmtTpInf_CtgyPurp(_ctgyPurp *common.CategoryPurpose1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].PmtTpInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].PmtTpInf = &common.PaymentTypeInformation26{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].PmtTpInf.CtgyPurp = _ctgyPurp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_ReqdExctnDt(_reqdExctnDt common.DateAndDateTime2Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].ReqdExctnDt = _reqdExctnDt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_XpryDt(_xpryDt *common.DateAndDateTime2Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].XpryDt = _xpryDt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_PmtCond(_pmtCond *common.PaymentCondition1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].PmtCond = _pmtCond
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_PmtCond_ImdtPmtRbt(_imdtPmtRbt *common.AmountOrRate1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].PmtCond == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].PmtCond = &common.PaymentCondition1{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].PmtCond.ImdtPmtRbt = _imdtPmtRbt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_PmtCond_ImdtPmtRbt_Amt(_amt *common.ActiveCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].PmtCond == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].PmtCond = &common.PaymentCondition1{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].PmtCond.ImdtPmtRbt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].PmtCond.ImdtPmtRbt = &common.AmountOrRate1Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].PmtCond.ImdtPmtRbt.Amt = _amt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_Dbtr(_dbtr common.PartyIdentification135) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].Dbtr = _dbtr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_Dbtr_Id(_id *common.Party38Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id = _id
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId(_orgId *common.OrganisationIdentification29) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.OrgId = _orgId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr(_othr []common.GenericOrganisationIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.OrgId = &common.OrganisationIdentification29{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.OrgId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm(_schmeNm *common.OrganisationIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.OrgId = &common.OrganisationIdentification29{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.OrgId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.OrgId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.OrgId.Othr, common.GenericOrganisationIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.OrgId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId(_prvtId *common.PersonIdentification13) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.PrvtId = _prvtId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth(_dtAndPlcOfBirth *common.DateAndPlaceOfBirth1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.PrvtId.DtAndPlcOfBirth = _dtAndPlcOfBirth
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr(_othr []common.GenericPersonIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.PrvtId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm(_schmeNm *common.PersonIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.PrvtId = &common.PersonIdentification13{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.PrvtId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.PrvtId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.PrvtId.Othr, common.GenericPersonIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.Id.PrvtId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls(_ctctDtls *common.Contact4) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.CtctDtls = _ctctDtls
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Othr(_othr []common.OtherContact1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.CtctDtls == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.CtctDtls = &common.Contact4{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].Dbtr.CtctDtls.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAcct(_dbtrAcct *common.CashAccount38) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct = _dbtrAcct
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id(_id common.AccountIdentification4Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct = &common.CashAccount38{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct.Id = _id
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr(_othr *common.GenericAccountIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct = &common.CashAccount38{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct.Id.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr_SchmeNm(_schmeNm *common.AccountSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct = &common.CashAccount38{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct.Id.Othr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct.Id.Othr = &common.GenericAccountIdentification1{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct.Id.Othr.SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Tp(_tp *common.CashAccountType2Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct = &common.CashAccount38{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct.Tp = _tp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Prxy(_prxy *common.ProxyAccountIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct = &common.CashAccount38{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct.Prxy = _prxy
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Prxy_Tp(_tp *common.ProxyAccountType1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct = &common.CashAccount38{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct.Prxy == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct.Prxy = &common.ProxyAccountIdentification1{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAcct.Prxy.Tp = _tp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAgt(_dbtrAgt common.BranchAndFinancialInstitutionIdentification6) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt = _dbtrAgt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId(_finInstnId common.FinancialInstitutionIdentification18) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId = _finInstnId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId(_clrSysMmbId *common.ClearingSystemMemberIdentification2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.ClrSysMmbId = _clrSysMmbId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId(_clrSysId *common.ClearingSystemIdentification2Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.ClrSysMmbId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.ClrSysMmbId = &common.ClearingSystemMemberIdentification2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId = _clrSysId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr(_othr *common.GenericFinancialIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm(_schmeNm *common.FinancialIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.Othr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.Othr = &common.GenericFinancialIdentification1{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.Othr.SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId(_brnchId *common.BranchData3) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId = _brnchId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId = &common.BranchData3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId = &common.BranchData3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId = &common.BranchData3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.BrnchId.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_UltmtDbtr(_ultmtDbtr *common.PartyIdentification135) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr = _ultmtDbtr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id(_id *common.Party38Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id = _id
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId(_orgId *common.OrganisationIdentification29) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.OrgId = _orgId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr(_othr []common.GenericOrganisationIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.OrgId = &common.OrganisationIdentification29{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.OrgId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm(_schmeNm *common.OrganisationIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.OrgId = &common.OrganisationIdentification29{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.OrgId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.OrgId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.OrgId.Othr, common.GenericOrganisationIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.OrgId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId(_prvtId *common.PersonIdentification13) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.PrvtId = _prvtId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth(_dtAndPlcOfBirth *common.DateAndPlaceOfBirth1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth = _dtAndPlcOfBirth
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr(_othr []common.GenericPersonIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.PrvtId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm(_schmeNm *common.PersonIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.PrvtId = &common.PersonIdentification13{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.PrvtId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.PrvtId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.PrvtId.Othr, common.GenericPersonIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.Id.PrvtId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls(_ctctDtls *common.Contact4) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.CtctDtls = _ctctDtls
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Othr(_othr []common.OtherContact1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.CtctDtls == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.CtctDtls = &common.Contact4{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].UltmtDbtr.CtctDtls.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx(_cdtTrfTx []CreditTransferTransaction35) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = _cdtTrfTx
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId(_pmtId PaymentIdentification6) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtId = _pmtId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf(_pmtTpInf *common.PaymentTypeInformation26) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf = _pmtTpInf
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_SvcLvl(_svcLvl []common.ServiceLevel8Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf = &common.PaymentTypeInformation26{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf.SvcLvl = _svcLvl
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_LclInstrm(_lclInstrm *common.LocalInstrument2Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf = &common.PaymentTypeInformation26{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf.LclInstrm = _lclInstrm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_CtgyPurp(_ctgyPurp *common.CategoryPurpose1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf = &common.PaymentTypeInformation26{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf.CtgyPurp = _ctgyPurp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond(_pmtCond *common.PaymentCondition1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtCond = _pmtCond
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_ImdtPmtRbt(_imdtPmtRbt *common.AmountOrRate1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtCond == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtCond = &common.PaymentCondition1{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtCond.ImdtPmtRbt = _imdtPmtRbt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_ImdtPmtRbt_Amt(_amt *common.ActiveCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtCond == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtCond = &common.PaymentCondition1{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtCond.ImdtPmtRbt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtCond.ImdtPmtRbt = &common.AmountOrRate1Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtCond.ImdtPmtRbt.Amt = _amt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt(_amt common.AmountType4Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Amt = _amt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt(_instdAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Amt.InstdAmt = _instdAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_EqvtAmt(_eqvtAmt *common.EquivalentAmount2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Amt.EqvtAmt = _eqvtAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_EqvtAmt_Amt(_amt common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Amt.EqvtAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Amt.EqvtAmt = &common.EquivalentAmount2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Amt.EqvtAmt.Amt = _amt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr(_chqInstr *Cheque11) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr = _chqInstr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr(_chqFr *NameAndAddress16) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr = &Cheque11{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.ChqFr = _chqFr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr(_adr common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr = &Cheque11{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.ChqFr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.ChqFr = &NameAndAddress16{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.ChqFr.Adr = _adr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr = &Cheque11{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.ChqFr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.ChqFr = &NameAndAddress16{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.ChqFr.Adr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr = &Cheque11{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.ChqFr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.ChqFr = &NameAndAddress16{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.ChqFr.Adr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.ChqFr.Adr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.ChqFr.Adr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvryMtd(_dlvryMtd *ChequeDeliveryMethod1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr = &Cheque11{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.DlvryMtd = _dlvryMtd
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo(_dlvrTo *NameAndAddress16) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr = &Cheque11{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.DlvrTo = _dlvrTo
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr(_adr common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr = &Cheque11{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.DlvrTo == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.DlvrTo = &NameAndAddress16{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.DlvrTo.Adr = _adr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr = &Cheque11{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.DlvrTo == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.DlvrTo = &NameAndAddress16{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.DlvrTo.Adr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr = &Cheque11{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.DlvrTo == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.DlvrTo = &NameAndAddress16{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.DlvrTo.Adr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.DlvrTo.Adr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChqInstr.DlvrTo.Adr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr(_ultmtDbtr *common.PartyIdentification135) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr = _ultmtDbtr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id(_id *common.Party38Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id = _id
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId(_orgId *common.OrganisationIdentification29) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.OrgId = _orgId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr(_othr []common.GenericOrganisationIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.OrgId = &common.OrganisationIdentification29{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.OrgId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_SchmeNm(_schmeNm *common.OrganisationIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.OrgId = &common.OrganisationIdentification29{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.OrgId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.OrgId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.OrgId.Othr, common.GenericOrganisationIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.OrgId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId(_prvtId *common.PersonIdentification13) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.PrvtId = _prvtId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth(_dtAndPlcOfBirth *common.DateAndPlaceOfBirth1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth = _dtAndPlcOfBirth
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr(_othr []common.GenericPersonIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.PrvtId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_SchmeNm(_schmeNm *common.PersonIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.PrvtId = &common.PersonIdentification13{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.PrvtId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.PrvtId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.PrvtId.Othr, common.GenericPersonIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.Id.PrvtId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls(_ctctDtls *common.Contact4) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.CtctDtls = _ctctDtls
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Othr(_othr []common.OtherContact1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.CtctDtls == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.CtctDtls = &common.Contact4{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtDbtr.CtctDtls.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1(_intrmyAgt1 *common.BranchAndFinancialInstitutionIdentification6) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 = _intrmyAgt1
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId(_finInstnId common.FinancialInstitutionIdentification18) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId = _finInstnId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_ClrSysMmbId(_clrSysMmbId *common.ClearingSystemMemberIdentification2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.ClrSysMmbId = _clrSysMmbId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId(_clrSysId *common.ClearingSystemIdentification2Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.ClrSysMmbId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.ClrSysMmbId = &common.ClearingSystemMemberIdentification2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.ClrSysMmbId.ClrSysId = _clrSysId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr(_othr *common.GenericFinancialIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr_SchmeNm(_schmeNm *common.FinancialIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.Othr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.Othr = &common.GenericFinancialIdentification1{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.FinInstnId.Othr.SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId(_brnchId *common.BranchData3) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId = _brnchId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId = &common.BranchData3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId = &common.BranchData3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId = &common.BranchData3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt1.BrnchId.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2(_intrmyAgt2 *common.BranchAndFinancialInstitutionIdentification6) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 = _intrmyAgt2
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId(_finInstnId common.FinancialInstitutionIdentification18) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId = _finInstnId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_ClrSysMmbId(_clrSysMmbId *common.ClearingSystemMemberIdentification2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.ClrSysMmbId = _clrSysMmbId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId(_clrSysId *common.ClearingSystemIdentification2Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.ClrSysMmbId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.ClrSysMmbId = &common.ClearingSystemMemberIdentification2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.ClrSysMmbId.ClrSysId = _clrSysId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr(_othr *common.GenericFinancialIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr_SchmeNm(_schmeNm *common.FinancialIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.Othr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.Othr = &common.GenericFinancialIdentification1{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.FinInstnId.Othr.SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId(_brnchId *common.BranchData3) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId = _brnchId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId = &common.BranchData3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId = &common.BranchData3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId = &common.BranchData3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt2.BrnchId.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3(_intrmyAgt3 *common.BranchAndFinancialInstitutionIdentification6) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 = _intrmyAgt3
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId(_finInstnId common.FinancialInstitutionIdentification18) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId = _finInstnId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_ClrSysMmbId(_clrSysMmbId *common.ClearingSystemMemberIdentification2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.ClrSysMmbId = _clrSysMmbId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId(_clrSysId *common.ClearingSystemIdentification2Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.ClrSysMmbId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.ClrSysMmbId = &common.ClearingSystemMemberIdentification2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.ClrSysMmbId.ClrSysId = _clrSysId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr(_othr *common.GenericFinancialIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr_SchmeNm(_schmeNm *common.FinancialIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.Othr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.Othr = &common.GenericFinancialIdentification1{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.FinInstnId.Othr.SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId(_brnchId *common.BranchData3) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId = _brnchId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId = &common.BranchData3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId = &common.BranchData3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3 = &common.BranchAndFinancialInstitutionIdentification6{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId = &common.BranchData3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].IntrmyAgt3.BrnchId.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt(_cdtrAgt common.BranchAndFinancialInstitutionIdentification6) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt = _cdtrAgt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId(_finInstnId common.FinancialInstitutionIdentification18) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId = _finInstnId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_ClrSysMmbId(_clrSysMmbId *common.ClearingSystemMemberIdentification2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.ClrSysMmbId = _clrSysMmbId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId(_clrSysId *common.ClearingSystemIdentification2Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.ClrSysMmbId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.ClrSysMmbId = &common.ClearingSystemMemberIdentification2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId = _clrSysId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr(_othr *common.GenericFinancialIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr_SchmeNm(_schmeNm *common.FinancialIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.Othr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.Othr = &common.GenericFinancialIdentification1{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.Othr.SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId(_brnchId *common.BranchData3) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId = _brnchId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId = &common.BranchData3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId = &common.BranchData3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId = &common.BranchData3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.BrnchId.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr(_cdtr common.PartyIdentification135) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr = _cdtr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id(_id *common.Party38Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id = _id
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId(_orgId *common.OrganisationIdentification29) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.OrgId = _orgId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr(_othr []common.GenericOrganisationIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.OrgId = &common.OrganisationIdentification29{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.OrgId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_SchmeNm(_schmeNm *common.OrganisationIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.OrgId = &common.OrganisationIdentification29{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.OrgId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.OrgId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.OrgId.Othr, common.GenericOrganisationIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.OrgId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId(_prvtId *common.PersonIdentification13) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.PrvtId = _prvtId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_DtAndPlcOfBirth(_dtAndPlcOfBirth *common.DateAndPlaceOfBirth1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.PrvtId.DtAndPlcOfBirth = _dtAndPlcOfBirth
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr(_othr []common.GenericPersonIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.PrvtId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_SchmeNm(_schmeNm *common.PersonIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.PrvtId = &common.PersonIdentification13{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.PrvtId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.PrvtId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.PrvtId.Othr, common.GenericPersonIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Id.PrvtId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls(_ctctDtls *common.Contact4) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.CtctDtls = _ctctDtls
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Othr(_othr []common.OtherContact1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.CtctDtls == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.CtctDtls = &common.Contact4{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.CtctDtls.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct(_cdtrAcct *common.CashAccount38) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct = _cdtrAcct
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id(_id common.AccountIdentification4Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct = &common.CashAccount38{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct.Id = _id
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr(_othr *common.GenericAccountIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct = &common.CashAccount38{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct.Id.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr_SchmeNm(_schmeNm *common.AccountSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct = &common.CashAccount38{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct.Id.Othr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct.Id.Othr = &common.GenericAccountIdentification1{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct.Id.Othr.SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Tp(_tp *common.CashAccountType2Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct = &common.CashAccount38{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct.Tp = _tp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Prxy(_prxy *common.ProxyAccountIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct = &common.CashAccount38{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct.Prxy = _prxy
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Prxy_Tp(_tp *common.ProxyAccountType1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct = &common.CashAccount38{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct.Prxy == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct.Prxy = &common.ProxyAccountIdentification1{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct.Prxy.Tp = _tp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr(_ultmtCdtr *common.PartyIdentification135) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr = _ultmtCdtr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id(_id *common.Party38Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id = _id
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId(_orgId *common.OrganisationIdentification29) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.OrgId = _orgId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr(_othr []common.GenericOrganisationIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.OrgId = &common.OrganisationIdentification29{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.OrgId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_SchmeNm(_schmeNm *common.OrganisationIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.OrgId = &common.OrganisationIdentification29{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.OrgId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.OrgId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.OrgId.Othr, common.GenericOrganisationIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.OrgId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId(_prvtId *common.PersonIdentification13) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.PrvtId = _prvtId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth(_dtAndPlcOfBirth *common.DateAndPlaceOfBirth1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth = _dtAndPlcOfBirth
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr(_othr []common.GenericPersonIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.PrvtId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_SchmeNm(_schmeNm *common.PersonIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.PrvtId = &common.PersonIdentification13{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.PrvtId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.PrvtId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.PrvtId.Othr, common.GenericPersonIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.Id.PrvtId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls(_ctctDtls *common.Contact4) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.CtctDtls = _ctctDtls
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Othr(_othr []common.OtherContact1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.CtctDtls == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.CtctDtls = &common.Contact4{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].UltmtCdtr.CtctDtls.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_InstrForCdtrAgt(_instrForCdtrAgt []InstructionForCreditorAgent1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].InstrForCdtrAgt = _instrForCdtrAgt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Purp(_purp *common.Purpose2Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Purp = _purp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg(_rgltryRptg []RegulatoryReporting3) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg = _rgltryRptg
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Authrty(_authrty *RegulatoryAuthority2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg, RegulatoryReporting3{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg[0].Authrty = _authrty
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls(_dtls []StructuredRegulatoryReporting3) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg, RegulatoryReporting3{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg[0].Dtls = _dtls
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Amt(_amt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg, RegulatoryReporting3{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg[0].Dtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg[0].Dtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg[0].Dtls, StructuredRegulatoryReporting3{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RgltryRptg[0].Dtls[0].Amt = _amt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax(_tax *TaxInformation8) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = _tax
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Cdtr(_cdtr *common.TaxParty1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = &TaxInformation8{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Cdtr = _cdtr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr(_dbtr *common.TaxParty2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = &TaxInformation8{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Dbtr = _dbtr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_Authstn(_authstn *common.TaxAuthorisation1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = &TaxInformation8{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Dbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Dbtr = &common.TaxParty2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Dbtr.Authstn = _authstn
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxblBaseAmt(_ttlTaxblBaseAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = &TaxInformation8{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.TtlTaxblBaseAmt = _ttlTaxblBaseAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxAmt(_ttlTaxAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = &TaxInformation8{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.TtlTaxAmt = _ttlTaxAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd(_rcrd []common.TaxRecord2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = &TaxInformation8{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd = _rcrd
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd(_prd *common.TaxPeriod2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = &TaxInformation8{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd, common.TaxRecord2{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].Prd = _prd
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd_FrToDt(_frToDt *common.DatePeriod2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = &TaxInformation8{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd, common.TaxRecord2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].Prd == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].Prd = &common.TaxPeriod2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].Prd.FrToDt = _frToDt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt(_taxAmt *common.TaxAmount2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = &TaxInformation8{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd, common.TaxRecord2{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt = _taxAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TaxblBaseAmt(_taxblBaseAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = &TaxInformation8{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd, common.TaxRecord2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt = &common.TaxAmount2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.TaxblBaseAmt = _taxblBaseAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TtlAmt(_ttlAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = &TaxInformation8{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd, common.TaxRecord2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt = &common.TaxAmount2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.TtlAmt = _ttlAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls(_dtls []common.TaxRecordDetails2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = &TaxInformation8{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd, common.TaxRecord2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt = &common.TaxAmount2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.Dtls = _dtls
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd(_prd *common.TaxPeriod2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = &TaxInformation8{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd, common.TaxRecord2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt = &common.TaxAmount2{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.Dtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.Dtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.Dtls, common.TaxRecordDetails2{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.Dtls[0].Prd = _prd
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt(_frToDt *common.DatePeriod2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = &TaxInformation8{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd, common.TaxRecord2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt = &common.TaxAmount2{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.Dtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.Dtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.Dtls, common.TaxRecordDetails2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.Dtls[0].Prd == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.Dtls[0].Prd = &common.TaxPeriod2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.Dtls[0].Prd.FrToDt = _frToDt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Amt(_amt common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax = &TaxInformation8{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd, common.TaxRecord2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt = &common.TaxAmount2{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.Dtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.Dtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.Dtls, common.TaxRecordDetails2{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Tax.Rcrd[0].TaxAmt.Dtls[0].Amt = _amt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf(_rltdRmtInf []RemittanceLocation7) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf = _rltdRmtInf
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls(_rmtLctnDtls []RemittanceLocationData1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf, RemittanceLocation7{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls = _rmtLctnDtls
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr(_pstlAdr *NameAndAddress16) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf, RemittanceLocation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls, RemittanceLocationData1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls[0].PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr(_adr common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf, RemittanceLocation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls, RemittanceLocationData1{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls[0].PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls[0].PstlAdr = &NameAndAddress16{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls[0].PstlAdr.Adr = _adr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf, RemittanceLocation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls, RemittanceLocationData1{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls[0].PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls[0].PstlAdr = &NameAndAddress16{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls[0].PstlAdr.Adr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf, RemittanceLocation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls, RemittanceLocationData1{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls[0].PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls[0].PstlAdr = &NameAndAddress16{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls[0].PstlAdr.Adr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls[0].PstlAdr.Adr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RltdRmtInf[0].RmtLctnDtls[0].PstlAdr.Adr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf(_rmtInf *common.RemittanceInformation16) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = _rmtInf
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd(_strd []common.StructuredRemittanceInformation16) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = _strd
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf(_rfrdDocInf []common.ReferredDocumentInformation7) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = _rfrdDocInf
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Tp(_tp *common.ReferredDocumentType4) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].Tp = _tp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry(_cdOrPrtry common.ReferredDocumentType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].Tp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].Tp = &common.ReferredDocumentType4{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].Tp.CdOrPrtry = _cdOrPrtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls(_lineDtls []common.DocumentLineInformation1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = _lineDtls
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id(_id []common.DocumentLineIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls, common.DocumentLineInformation1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Id = _id
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp(_tp *common.DocumentLineType1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls, common.DocumentLineInformation1{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Id) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Id = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Id, common.DocumentLineIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Id[0].Tp = _tp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry(_cdOrPrtry common.DocumentLineType1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls, common.DocumentLineInformation1{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Id) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Id = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Id, common.DocumentLineIdentification1{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Id[0].Tp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Id[0].Tp = &common.DocumentLineType1{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Id[0].Tp.CdOrPrtry = _cdOrPrtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt(_amt *common.RemittanceAmount3) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls, common.DocumentLineInformation1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt = _amt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt(_duePyblAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls, common.DocumentLineInformation1{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt = &common.RemittanceAmount3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.DuePyblAmt = _duePyblAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt(_dscntApldAmt []common.DiscountAmountAndType1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls, common.DocumentLineInformation1{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt = &common.RemittanceAmount3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.DscntApldAmt = _dscntApldAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp(_tp *common.DiscountAmountType1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls, common.DocumentLineInformation1{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt = &common.RemittanceAmount3{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.DscntApldAmt) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.DscntApldAmt = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.DscntApldAmt, common.DiscountAmountAndType1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.DscntApldAmt[0].Tp = _tp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt(_amt common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls, common.DocumentLineInformation1{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt = &common.RemittanceAmount3{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.DscntApldAmt) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.DscntApldAmt = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.DscntApldAmt, common.DiscountAmountAndType1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.DscntApldAmt[0].Amt = _amt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt(_cdtNoteAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls, common.DocumentLineInformation1{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt = &common.RemittanceAmount3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.CdtNoteAmt = _cdtNoteAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt(_taxAmt []common.TaxAmountAndType1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls, common.DocumentLineInformation1{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt = &common.RemittanceAmount3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.TaxAmt = _taxAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp(_tp *common.TaxAmountType1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls, common.DocumentLineInformation1{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt = &common.RemittanceAmount3{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.TaxAmt) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.TaxAmt = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.TaxAmt, common.TaxAmountAndType1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.TaxAmt[0].Tp = _tp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt(_amt common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls, common.DocumentLineInformation1{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt = &common.RemittanceAmount3{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.TaxAmt) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.TaxAmt = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.TaxAmt, common.TaxAmountAndType1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.TaxAmt[0].Amt = _amt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn(_adjstmntAmtAndRsn []common.DocumentAdjustment1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls, common.DocumentLineInformation1{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt = &common.RemittanceAmount3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.AdjstmntAmtAndRsn = _adjstmntAmtAndRsn
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt(_amt common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls, common.DocumentLineInformation1{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt = &common.RemittanceAmount3{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.AdjstmntAmtAndRsn) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.AdjstmntAmtAndRsn = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.AdjstmntAmtAndRsn, common.DocumentAdjustment1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.AdjstmntAmtAndRsn[0].Amt = _amt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt(_rmtdAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf, common.ReferredDocumentInformation7{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls, common.DocumentLineInformation1{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt = &common.RemittanceAmount3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].LineDtls[0].Amt.RmtdAmt = _rmtdAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt(_rfrdDocAmt *common.RemittanceAmount2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt = _rfrdDocAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DuePyblAmt(_duePyblAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt = &common.RemittanceAmount2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.DuePyblAmt = _duePyblAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt(_dscntApldAmt []common.DiscountAmountAndType1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt = &common.RemittanceAmount2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.DscntApldAmt = _dscntApldAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp(_tp *common.DiscountAmountType1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt = &common.RemittanceAmount2{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.DscntApldAmt) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.DscntApldAmt = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.DscntApldAmt, common.DiscountAmountAndType1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.DscntApldAmt[0].Tp = _tp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt(_amt common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt = &common.RemittanceAmount2{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.DscntApldAmt) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.DscntApldAmt = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.DscntApldAmt, common.DiscountAmountAndType1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.DscntApldAmt[0].Amt = _amt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt(_cdtNoteAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt = &common.RemittanceAmount2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.CdtNoteAmt = _cdtNoteAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt(_taxAmt []common.TaxAmountAndType1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt = &common.RemittanceAmount2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.TaxAmt = _taxAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp(_tp *common.TaxAmountType1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt = &common.RemittanceAmount2{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.TaxAmt) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.TaxAmt = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.TaxAmt, common.TaxAmountAndType1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.TaxAmt[0].Tp = _tp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt(_amt common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt = &common.RemittanceAmount2{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.TaxAmt) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.TaxAmt = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.TaxAmt, common.TaxAmountAndType1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.TaxAmt[0].Amt = _amt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn(_adjstmntAmtAndRsn []common.DocumentAdjustment1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt = &common.RemittanceAmount2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.AdjstmntAmtAndRsn = _adjstmntAmtAndRsn
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt(_amt common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt = &common.RemittanceAmount2{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.AdjstmntAmtAndRsn) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.AdjstmntAmtAndRsn = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.AdjstmntAmtAndRsn, common.DocumentAdjustment1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.AdjstmntAmtAndRsn[0].Amt = _amt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_RmtdAmt(_rmtdAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt = &common.RemittanceAmount2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocAmt.RmtdAmt = _rmtdAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf(_cdtrRefInf *common.CreditorReferenceInformation2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].CdtrRefInf = _cdtrRefInf
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Tp(_tp *common.CreditorReferenceType2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].CdtrRefInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].CdtrRefInf = &common.CreditorReferenceInformation2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].CdtrRefInf.Tp = _tp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry(_cdOrPrtry common.CreditorReferenceType1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].CdtrRefInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].CdtrRefInf = &common.CreditorReferenceInformation2{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].CdtrRefInf.Tp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].CdtrRefInf.Tp = &common.CreditorReferenceType2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].CdtrRefInf.Tp.CdOrPrtry = _cdOrPrtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr(_invcr *common.PartyIdentification135) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr = _invcr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id(_id *common.Party38Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id = _id
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId(_orgId *common.OrganisationIdentification29) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.OrgId = _orgId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr(_othr []common.GenericOrganisationIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.OrgId = &common.OrganisationIdentification29{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.OrgId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm(_schmeNm *common.OrganisationIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.OrgId = &common.OrganisationIdentification29{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.OrgId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.OrgId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.OrgId.Othr, common.GenericOrganisationIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.OrgId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId(_prvtId *common.PersonIdentification13) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.PrvtId = _prvtId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth(_dtAndPlcOfBirth *common.DateAndPlaceOfBirth1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.PrvtId.DtAndPlcOfBirth = _dtAndPlcOfBirth
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr(_othr []common.GenericPersonIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.PrvtId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm(_schmeNm *common.PersonIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.PrvtId = &common.PersonIdentification13{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.PrvtId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.PrvtId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.PrvtId.Othr, common.GenericPersonIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.Id.PrvtId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls(_ctctDtls *common.Contact4) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.CtctDtls = _ctctDtls
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Othr(_othr []common.OtherContact1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.CtctDtls == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.CtctDtls = &common.Contact4{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcr.CtctDtls.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee(_invcee *common.PartyIdentification135) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee = _invcee
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id(_id *common.Party38Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id = _id
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId(_orgId *common.OrganisationIdentification29) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.OrgId = _orgId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr(_othr []common.GenericOrganisationIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.OrgId = &common.OrganisationIdentification29{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.OrgId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm(_schmeNm *common.OrganisationIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.OrgId = &common.OrganisationIdentification29{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.OrgId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.OrgId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.OrgId.Othr, common.GenericOrganisationIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.OrgId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId(_prvtId *common.PersonIdentification13) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.PrvtId = _prvtId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth(_dtAndPlcOfBirth *common.DateAndPlaceOfBirth1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.PrvtId.DtAndPlcOfBirth = _dtAndPlcOfBirth
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr(_othr []common.GenericPersonIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.PrvtId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm(_schmeNm *common.PersonIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.PrvtId = &common.PersonIdentification13{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.PrvtId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.PrvtId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.PrvtId.Othr, common.GenericPersonIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.Id.PrvtId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls(_ctctDtls *common.Contact4) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.CtctDtls = _ctctDtls
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Othr(_othr []common.OtherContact1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.CtctDtls == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.CtctDtls = &common.Contact4{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].Invcee.CtctDtls.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt(_taxRmt *common.TaxInformation7) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = _taxRmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Cdtr(_cdtr *common.TaxParty1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Cdtr = _cdtr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr(_dbtr *common.TaxParty2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Dbtr = _dbtr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_Authstn(_authstn *common.TaxAuthorisation1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Dbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Dbtr = &common.TaxParty2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Dbtr.Authstn = _authstn
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr(_ultmtDbtr *common.TaxParty2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.UltmtDbtr = _ultmtDbtr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn(_authstn *common.TaxAuthorisation1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.UltmtDbtr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.UltmtDbtr = &common.TaxParty2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.UltmtDbtr.Authstn = _authstn
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt(_ttlTaxblBaseAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.TtlTaxblBaseAmt = _ttlTaxblBaseAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxAmt(_ttlTaxAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.TtlTaxAmt = _ttlTaxAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd(_rcrd []common.TaxRecord2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd = _rcrd
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd(_prd *common.TaxPeriod2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd, common.TaxRecord2{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].Prd = _prd
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt(_frToDt *common.DatePeriod2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd, common.TaxRecord2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].Prd == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].Prd = &common.TaxPeriod2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].Prd.FrToDt = _frToDt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt(_taxAmt *common.TaxAmount2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd, common.TaxRecord2{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt = _taxAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt(_taxblBaseAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd, common.TaxRecord2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt = &common.TaxAmount2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.TaxblBaseAmt = _taxblBaseAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt(_ttlAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd, common.TaxRecord2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt = &common.TaxAmount2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.TtlAmt = _ttlAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls(_dtls []common.TaxRecordDetails2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd, common.TaxRecord2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt = &common.TaxAmount2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.Dtls = _dtls
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd(_prd *common.TaxPeriod2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd, common.TaxRecord2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt = &common.TaxAmount2{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.Dtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.Dtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.Dtls, common.TaxRecordDetails2{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.Dtls[0].Prd = _prd
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt(_frToDt *common.DatePeriod2) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd, common.TaxRecord2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt = &common.TaxAmount2{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.Dtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.Dtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.Dtls, common.TaxRecordDetails2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.Dtls[0].Prd == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.Dtls[0].Prd = &common.TaxPeriod2{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.Dtls[0].Prd.FrToDt = _frToDt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt(_amt common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt = &common.TaxInformation7{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd, common.TaxRecord2{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt = &common.TaxAmount2{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.Dtls) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.Dtls = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.Dtls, common.TaxRecordDetails2{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].TaxAmt.Dtls[0].Amt = _amt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt(_grnshmtRmt *common.Garnishment3) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = _grnshmtRmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Tp(_tp common.GarnishmentType1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Tp = _tp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry(_cdOrPrtry common.GarnishmentType1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Tp.CdOrPrtry = _cdOrPrtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee(_grnshee *common.PartyIdentification135) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee = _grnshee
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id(_id *common.Party38Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id = _id
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId(_orgId *common.OrganisationIdentification29) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.OrgId = _orgId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr(_othr []common.GenericOrganisationIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.OrgId = &common.OrganisationIdentification29{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.OrgId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm(_schmeNm *common.OrganisationIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.OrgId = &common.OrganisationIdentification29{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.OrgId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.OrgId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.OrgId.Othr, common.GenericOrganisationIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.OrgId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId(_prvtId *common.PersonIdentification13) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.PrvtId = _prvtId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth(_dtAndPlcOfBirth *common.DateAndPlaceOfBirth1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.PrvtId.DtAndPlcOfBirth = _dtAndPlcOfBirth
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr(_othr []common.GenericPersonIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.PrvtId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm(_schmeNm *common.PersonIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.PrvtId = &common.PersonIdentification13{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.PrvtId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.PrvtId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.PrvtId.Othr, common.GenericPersonIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.Id.PrvtId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls(_ctctDtls *common.Contact4) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.CtctDtls = _ctctDtls
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr(_othr []common.OtherContact1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.CtctDtls == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.CtctDtls = &common.Contact4{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.Grnshee.CtctDtls.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr(_grnshmtAdmstr *common.PartyIdentification135) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr = _grnshmtAdmstr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id(_id *common.Party38Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id = _id
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId(_orgId *common.OrganisationIdentification29) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId = _orgId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr(_othr []common.GenericOrganisationIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId = &common.OrganisationIdentification29{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm(_schmeNm *common.OrganisationIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId = &common.OrganisationIdentification29{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr, common.GenericOrganisationIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId(_prvtId *common.PersonIdentification13) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId = _prvtId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth(_dtAndPlcOfBirth *common.DateAndPlaceOfBirth1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.DtAndPlcOfBirth = _dtAndPlcOfBirth
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr(_othr []common.GenericPersonIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm(_schmeNm *common.PersonIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId = &common.PersonIdentification13{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr, common.GenericPersonIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls(_ctctDtls *common.Contact4) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr = &common.PartyIdentification135{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.CtctDtls = _ctctDtls
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr(_othr []common.OtherContact1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr = &common.PartyIdentification135{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.CtctDtls == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.CtctDtls = &common.Contact4{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_RmtdAmt(_rmtdAmt *common.ActiveOrHistoricCurrencyAndAmount) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf = &common.RemittanceInformation16{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd, common.StructuredRemittanceInformation16{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt = &common.Garnishment3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].GrnshmtRmt.RmtdAmt = _rmtdAmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile(_nclsdFile []common.Document12) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = _nclsdFile
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp(_tp common.DocumentType1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].Tp = _tp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp_Prtry(_prtry *common.GenericIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].Tp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_IsseDt(_isseDt common.DateAndDateTime2Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].IsseDt = _isseDt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt(_frmt common.DocumentFormat1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].Frmt = _frmt
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt_Prtry(_prtry *common.GenericIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].Frmt.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr(_dgtlSgntr *common.PartyAndSignature3) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = _dgtlSgntr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty(_pty common.PartyIdentification135) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = &common.PartyAndSignature3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty = _pty
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr(_pstlAdr *common.PostalAddress24) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = &common.PartyAndSignature3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.PstlAdr = _pstlAdr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp(_adrTp *common.AddressType3Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = &common.PartyAndSignature3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.PstlAdr = &common.PostalAddress24{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.PstlAdr.AdrTp = _adrTp
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry(_prtry *common.GenericIdentification30) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = &common.PartyAndSignature3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.PstlAdr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.PstlAdr = &common.PostalAddress24{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.PstlAdr.AdrTp == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.PstlAdr.AdrTp = &common.AddressType3Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.PstlAdr.AdrTp.Prtry = _prtry
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id(_id *common.Party38Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = &common.PartyAndSignature3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id = _id
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId(_orgId *common.OrganisationIdentification29) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = &common.PartyAndSignature3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.OrgId = _orgId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr(_othr []common.GenericOrganisationIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = &common.PartyAndSignature3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.OrgId = &common.OrganisationIdentification29{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.OrgId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_SchmeNm(_schmeNm *common.OrganisationIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = &common.PartyAndSignature3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.OrgId = &common.OrganisationIdentification29{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.OrgId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.OrgId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.OrgId.Othr, common.GenericOrganisationIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.OrgId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId(_prvtId *common.PersonIdentification13) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = &common.PartyAndSignature3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id = &common.Party38Choice{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.PrvtId = _prvtId
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth(_dtAndPlcOfBirth *common.DateAndPlaceOfBirth1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = &common.PartyAndSignature3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.PrvtId.DtAndPlcOfBirth = _dtAndPlcOfBirth
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr(_othr []common.GenericPersonIdentification1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = &common.PartyAndSignature3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.PrvtId = &common.PersonIdentification13{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.PrvtId.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_SchmeNm(_schmeNm *common.PersonIdentificationSchemeName1Choice) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = &common.PartyAndSignature3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.PrvtId == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.PrvtId = &common.PersonIdentification13{}
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.PrvtId.Othr) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.PrvtId.Othr = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.PrvtId.Othr, common.GenericPersonIdentification1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.Id.PrvtId.Othr[0].SchmeNm = _schmeNm
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls(_ctctDtls *common.Contact4) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = &common.PartyAndSignature3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.CtctDtls = _ctctDtls
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Othr(_othr []common.OtherContact1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = &common.PartyAndSignature3{}
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.CtctDtls == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.CtctDtls = &common.Contact4{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Pty.CtctDtls.Othr = _othr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Sgntr(_sgntr common.SkipPayload) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile, common.Document12{})
	}
	if d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr == nil {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr = &common.PartyAndSignature3{}
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].NclsdFile[0].DgtlSgntr.Sgntr = _sgntr
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_SplmtryData(_splmtryData []common.SupplementaryData1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].SplmtryData = _splmtryData
}

func (d *Document) With_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_SplmtryData_Envlp(_envlp common.SupplementaryDataEnvelope1) {
	if len(d.CdtrPmtActvtnReq.PmtInf) == 0 {
		d.CdtrPmtActvtnReq.PmtInf = append(d.CdtrPmtActvtnReq.PmtInf, PaymentInstruction31{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx, CreditTransferTransaction35{})
	}
	if len(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].SplmtryData) == 0 {
		d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].SplmtryData = append(d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].SplmtryData, common.SupplementaryData1{})
	}
	d.CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].SplmtryData[0].Envlp = _envlp
}

func (d *Document) With_CdtrPmtActvtnReq_SplmtryData(_splmtryData []common.SupplementaryData1) {
	d.CdtrPmtActvtnReq.SplmtryData = _splmtryData
}

func (d *Document) With_CdtrPmtActvtnReq_SplmtryData_Envlp(_envlp common.SupplementaryDataEnvelope1) {
	if len(d.CdtrPmtActvtnReq.SplmtryData) == 0 {
		d.CdtrPmtActvtnReq.SplmtryData = append(d.CdtrPmtActvtnReq.SplmtryData, common.SupplementaryData1{})
	}
	d.CdtrPmtActvtnReq.SplmtryData[0].Envlp = _envlp
}
