// Package pain_013_001_07
// Do not Edit. This stuff it's been automatically generated.
package pain_013_001_07

import (
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
	"reflect"
)

func (d *Document) Set(path string, src interface{}) error {

	/*
		path = strings.TrimPrefix(path, "/Doc/")
		path = strings.Replace(path, "*", "", -1)
		path = strings.Replace(path, "[]", "", -1)
		path = strings.Replace(path, "/", ".", -1)
	*/

	v := reflect.ValueOf(d)
	fields := d.mapper.TraversalsByName(v.Type(), []string{path})

	values := make([]interface{}, 1)
	err := fieldsByTraversal(v, fields, values, true)
	if err != nil {
		return err
	}

	return copy2Dest(path, values[0], src)
}

/*
func convertAssignRows(dest, src interface{}) error {

	switch d := dest.(type) {
	case *common.Max35Text:
		if d == nil {
			return errors.New("nil pointer... in unmarshalling Max35Text data")
		}
		*d = common.MustToMax35Text(src)
		return nil
	}
	return nil
}
*/

func copy2Dest(docPath string, dest, src interface{}) error {

	var err error
	switch docPath {

	case Path_CdtrPmtActvtnReq_GrpHdr_MsgId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_CreDtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_NbOfTxs:
		d, _ := dest.(*common.Max15NumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max15NumericText data for" + docPath)
		}

		*d, err = common.ToMax15NumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_CtrlSum:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtInfId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtMtd:
		d, _ := dest.(*common.PaymentMethod7Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PaymentMethod7Code data for" + docPath)
		}

		*d, err = common.ToPaymentMethod7Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_InstrPrty:
		d, _ := dest.(*common.Priority2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Priority2Code data for" + docPath)
		}

		*d, err = common.ToPriority2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_SvcLvl_Cd:
		d, _ := dest.(*common.ExternalServiceLevel1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalServiceLevel1Code data for" + docPath)
		}

		*d, err = common.ToExternalServiceLevel1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_SvcLvl_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_LclInstrm_Cd:
		d, _ := dest.(*common.ExternalLocalInstrument1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalLocalInstrument1Code data for" + docPath)
		}

		*d, err = common.ToExternalLocalInstrument1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_LclInstrm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_CtgyPurp_Cd:
		d, _ := dest.(*common.ExternalCategoryPurpose1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCategoryPurpose1Code data for" + docPath)
		}

		*d, err = common.ToExternalCategoryPurpose1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_CtgyPurp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_ReqdExctnDt_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_ReqdExctnDt_DtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_XpryDt_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_XpryDt_DtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtCond_AmtModAllwd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtCond_EarlyPmtAllwd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtCond_DelyPnlty:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtCond_ImdtPmtRbt_Amt_Ccy:
		d, _ := dest.(*common.ActiveCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtCond_ImdtPmtRbt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtCond_ImdtPmtRbt_Rate:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_PmtCond_GrntedPmtReqd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_IBAN:
		d, _ := dest.(*common.IBAN2007Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*d, err = common.ToIBAN2007Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr_Id:
		d, _ := dest.(*common.Max34Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*d, err = common.ToMax34Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalAccountIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalAccountIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Tp_Cd:
		d, _ := dest.(*common.ExternalCashAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashAccountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Nm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Prxy_Tp_Cd:
		d, _ := dest.(*common.ExternalProxyAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalProxyAccountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Prxy_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Prxy_Id:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_ChrgBr:
		d, _ := dest.(*common.ChargeBearerType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ChargeBearerType1Code data for" + docPath)
		}

		*d, err = common.ToChargeBearerType1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId_InstrId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId_EndToEndId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId_UETR:
		d, _ := dest.(*common.UUIDv4Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.UUIDv4Identifier data for" + docPath)
		}

		*d, err = common.ToUUIDv4Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_InstrPrty:
		d, _ := dest.(*common.Priority2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Priority2Code data for" + docPath)
		}

		*d, err = common.ToPriority2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_SvcLvl_Cd:
		d, _ := dest.(*common.ExternalServiceLevel1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalServiceLevel1Code data for" + docPath)
		}

		*d, err = common.ToExternalServiceLevel1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_SvcLvl_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_LclInstrm_Cd:
		d, _ := dest.(*common.ExternalLocalInstrument1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalLocalInstrument1Code data for" + docPath)
		}

		*d, err = common.ToExternalLocalInstrument1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_LclInstrm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_CtgyPurp_Cd:
		d, _ := dest.(*common.ExternalCategoryPurpose1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCategoryPurpose1Code data for" + docPath)
		}

		*d, err = common.ToExternalCategoryPurpose1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_CtgyPurp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_AmtModAllwd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_EarlyPmtAllwd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_DelyPnlty:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_ImdtPmtRbt_Amt_Ccy:
		d, _ := dest.(*common.ActiveCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_ImdtPmtRbt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_ImdtPmtRbt_Rate:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_GrntedPmtReqd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_EqvtAmt_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_EqvtAmt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_EqvtAmt_CcyOfTrf:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChrgBr:
		d, _ := dest.(*common.ChargeBearerType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ChargeBearerType1Code data for" + docPath)
		}

		*d, err = common.ToChargeBearerType1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqTp:
		d, _ := dest.(*common.ChequeType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ChequeType2Code data for" + docPath)
		}

		*d, err = common.ToChequeType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqNb:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvryMtd_Cd:
		d, _ := dest.(*common.ChequeDelivery1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ChequeDelivery1Code data for" + docPath)
		}

		*d, err = common.ToChequeDelivery1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvryMtd_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_InstrPrty:
		d, _ := dest.(*common.Priority2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Priority2Code data for" + docPath)
		}

		*d, err = common.ToPriority2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqMtrtyDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_FrmsCd:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_MemoFld:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_RgnlClrZone:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_PrtLctn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_Sgntr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_IBAN:
		d, _ := dest.(*common.IBAN2007Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*d, err = common.ToIBAN2007Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr_Id:
		d, _ := dest.(*common.Max34Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*d, err = common.ToMax34Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalAccountIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalAccountIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Tp_Cd:
		d, _ := dest.(*common.ExternalCashAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashAccountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Nm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Prxy_Tp_Cd:
		d, _ := dest.(*common.ExternalProxyAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalProxyAccountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Prxy_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Prxy_Id:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_InstrForCdtrAgt_Cd:
		d, _ := dest.(*common.Instruction3Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Instruction3Code data for" + docPath)
		}

		*d, err = common.ToInstruction3Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_InstrForCdtrAgt_InstrInf:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Purp_Cd:
		d, _ := dest.(*common.ExternalPurpose1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPurpose1Code data for" + docPath)
		}

		*d, err = common.ToExternalPurpose1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Purp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_DbtCdtRptgInd:
		d, _ := dest.(*common.RegulatoryReportingType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.RegulatoryReportingType1Code data for" + docPath)
		}

		*d, err = common.ToRegulatoryReportingType1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Authrty_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Authrty_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Tp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Cd:
		d, _ := dest.(*common.Max10Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max10Text data for" + docPath)
		}

		*d, err = common.ToMax10Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Inf:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Cdtr_TaxId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Cdtr_RegnId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Cdtr_TaxTp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_TaxId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_RegnId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_TaxTp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_Authstn_Titl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_Authstn_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_AdmstnZone:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_RefNb:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Mtd:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxblBaseAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxblBaseAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_SeqNb:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Tp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Ctgy:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_CtgyDtls:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_DbtrSts:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_CertId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_FrmsCd:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd_Yr:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd_Tp:
		d, _ := dest.(*common.TaxRecordPeriod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.TaxRecordPeriod1Code data for" + docPath)
		}

		*d, err = common.ToTaxRecordPeriod1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd_FrToDt_FrDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd_FrToDt_ToDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Rate:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TtlAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TtlAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd_Yr:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd_Tp:
		d, _ := dest.(*common.TaxRecordPeriod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.TaxRecordPeriod1Code data for" + docPath)
		}

		*d, err = common.ToTaxRecordPeriod1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_AddtlInf:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_Mtd:
		d, _ := dest.(*common.RemittanceLocationMethod2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.RemittanceLocationMethod2Code data for" + docPath)
		}

		*d, err = common.ToRemittanceLocationMethod2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_ElctrncAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Ustrd:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Cd:
		d, _ := dest.(*common.DocumentType6Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.DocumentType6Code data for" + docPath)
		}

		*d, err = common.ToDocumentType6Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Tp_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Nb:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_RltdDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry_Cd:
		d, _ := dest.(*common.ExternalDocumentLineType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDocumentLineType1Code data for" + docPath)
		}

		*d, err = common.ToExternalDocumentLineType1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Nb:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_RltdDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Desc:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp_Cd:
		d, _ := dest.(*common.ExternalDiscountAmountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDiscountAmountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalDiscountAmountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp_Cd:
		d, _ := dest.(*common.ExternalTaxAmountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalTaxAmountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalTaxAmountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_CdtDbtInd:
		d, _ := dest.(*common.CreditDebitCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CreditDebitCode data for" + docPath)
		}

		*d, err = common.ToCreditDebitCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Rsn:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_AddtlInf:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp_Cd:
		d, _ := dest.(*common.ExternalDiscountAmountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDiscountAmountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalDiscountAmountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp_Cd:
		d, _ := dest.(*common.ExternalTaxAmountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalTaxAmountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalTaxAmountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_CdtDbtInd:
		d, _ := dest.(*common.CreditDebitCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CreditDebitCode data for" + docPath)
		}

		*d, err = common.ToCreditDebitCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Rsn:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_AddtlInf:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Cd:
		d, _ := dest.(*common.DocumentType3Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.DocumentType3Code data for" + docPath)
		}

		*d, err = common.ToDocumentType3Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Tp_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Ref:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Cdtr_TaxId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Cdtr_RegnId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Cdtr_TaxTp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_TaxId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_RegnId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_TaxTp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_Authstn_Titl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_Authstn_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_TaxId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_RegnId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_TaxTp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn_Titl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_AdmstnZone:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_RefNb:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Mtd:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_SeqNb:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Tp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Ctgy:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_CtgyDtls:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_DbtrSts:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_CertId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_FrmsCd:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd_Yr:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd_Tp:
		d, _ := dest.(*common.TaxRecordPeriod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.TaxRecordPeriod1Code data for" + docPath)
		}

		*d, err = common.ToTaxRecordPeriod1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt_FrDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt_ToDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Rate:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_Yr:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_Tp:
		d, _ := dest.(*common.TaxRecordPeriod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.TaxRecordPeriod1Code data for" + docPath)
		}

		*d, err = common.ToTaxRecordPeriod1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_AddtlInf:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry_Cd:
		d, _ := dest.(*common.ExternalGarnishmentType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalGarnishmentType1Code data for" + docPath)
		}

		*d, err = common.ToExternalGarnishmentType1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Tp_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_RefNb:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_RmtdAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_RmtdAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_FmlyMdclInsrncInd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_MplyeeTermntnInd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_AddtlRmtInf:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp_Cd:
		d, _ := dest.(*common.ExternalDocumentType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDocumentType1Code data for" + docPath)
		}

		*d, err = common.ToExternalDocumentType1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp_Prtry_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_IsseDt_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_IsseDt_DtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_LangCd:
		d, _ := dest.(*xsdt.String)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.String data for" + docPath)
		}

		*d, err = xsdt.ToString(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt_Cd:
		d, _ := dest.(*common.ExternalDocumentFormat1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDocumentFormat1Code data for" + docPath)
		}

		*d, err = common.ToExternalDocumentFormat1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt_Prtry_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_FileNm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Sgntr_Item:
		d, _ := dest.(*xsdt.AnyType)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.AnyType data for" + docPath)
		}

		*d, err = xsdt.ToAnyType(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Nclsr:
		d, _ := dest.(*common.Max10MbBinary)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max10MbBinary data for" + docPath)
		}

		*d, err = common.ToMax10MbBinary(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_SplmtryData_PlcAndNm:
		d, _ := dest.(*common.Max350Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max350Text data for" + docPath)
		}

		*d, err = common.ToMax350Text(src)
		return err

	case Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_SplmtryData_Envlp_Item:
		d, _ := dest.(*xsdt.AnyType)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.AnyType data for" + docPath)
		}

		*d, err = xsdt.ToAnyType(src)
		return err

	case Path_CdtrPmtActvtnReq_SplmtryData_PlcAndNm:
		d, _ := dest.(*common.Max350Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max350Text data for" + docPath)
		}

		*d, err = common.ToMax350Text(src)
		return err

	case Path_CdtrPmtActvtnReq_SplmtryData_Envlp_Item:
		d, _ := dest.(*xsdt.AnyType)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.AnyType data for" + docPath)
		}

		*d, err = xsdt.ToAnyType(src)
		return err
	default:
		return fmt.Errorf("could not find path to node %s", docPath)
	}

	return nil
}
