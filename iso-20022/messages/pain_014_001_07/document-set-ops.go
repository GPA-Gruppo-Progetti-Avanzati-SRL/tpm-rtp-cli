// Package pain_014_001_07
// Do not Edit. This stuff it's been automatically generated.
package pain_014_001_07

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

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_MsgId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CreDtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_InitgPty_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_OrgnlMsgId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_OrgnlMsgNmId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_OrgnlCreDtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_OrgnlNbOfTxs:
		d, _ := dest.(*common.Max15NumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max15NumericText data for" + docPath)
		}

		*d, err = common.ToMax15NumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_OrgnlCtrlSum:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_GrpSts:
		d, _ := dest.(*common.ExternalPaymentGroupStatus1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPaymentGroupStatus1Code data for" + docPath)
		}

		*d, err = common.ToExternalPaymentGroupStatus1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Rsn_Cd:
		d, _ := dest.(*common.ExternalStatusReason1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalStatusReason1Code data for" + docPath)
		}

		*d, err = common.ToExternalStatusReason1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Rsn_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_StsRsnInf_AddtlInf:
		d, _ := dest.(*common.Max105Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max105Text data for" + docPath)
		}

		*d, err = common.ToMax105Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_NbOfTxsPerSts_DtldNbOfTxs:
		d, _ := dest.(*common.Max15NumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max15NumericText data for" + docPath)
		}

		*d, err = common.ToMax15NumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_NbOfTxsPerSts_DtldSts:
		d, _ := dest.(*common.ExternalPaymentTransactionStatus1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPaymentTransactionStatus1Code data for" + docPath)
		}

		*d, err = common.ToExternalPaymentTransactionStatus1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlGrpInfAndSts_NbOfTxsPerSts_DtldCtrlSum:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_OrgnlPmtInfId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_OrgnlNbOfTxs:
		d, _ := dest.(*common.Max15NumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max15NumericText data for" + docPath)
		}

		*d, err = common.ToMax15NumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_OrgnlCtrlSum:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_PmtInfSts:
		d, _ := dest.(*common.ExternalPaymentGroupStatus1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPaymentGroupStatus1Code data for" + docPath)
		}

		*d, err = common.ToExternalPaymentGroupStatus1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Rsn_Cd:
		d, _ := dest.(*common.ExternalStatusReason1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalStatusReason1Code data for" + docPath)
		}

		*d, err = common.ToExternalStatusReason1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Rsn_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_StsRsnInf_AddtlInf:
		d, _ := dest.(*common.Max105Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max105Text data for" + docPath)
		}

		*d, err = common.ToMax105Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_NbOfTxsPerSts_DtldNbOfTxs:
		d, _ := dest.(*common.Max15NumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max15NumericText data for" + docPath)
		}

		*d, err = common.ToMax15NumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_NbOfTxsPerSts_DtldSts:
		d, _ := dest.(*common.ExternalPaymentTransactionStatus1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPaymentTransactionStatus1Code data for" + docPath)
		}

		*d, err = common.ToExternalPaymentTransactionStatus1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_NbOfTxsPerSts_DtldCtrlSum:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlInstrId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlEndToEndId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlUETR:
		d, _ := dest.(*common.UUIDv4Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.UUIDv4Identifier data for" + docPath)
		}

		*d, err = common.ToUUIDv4Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_TxSts:
		d, _ := dest.(*common.ExternalPaymentTransactionStatus1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPaymentTransactionStatus1Code data for" + docPath)
		}

		*d, err = common.ToExternalPaymentTransactionStatus1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Rsn_Cd:
		d, _ := dest.(*common.ExternalStatusReason1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalStatusReason1Code data for" + docPath)
		}

		*d, err = common.ToExternalStatusReason1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Rsn_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_AddtlInf:
		d, _ := dest.(*common.Max105Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max105Text data for" + docPath)
		}

		*d, err = common.ToMax105Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_PmtCondSts_AccptdAmt_Ccy:
		d, _ := dest.(*common.ActiveCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_PmtCondSts_AccptdAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_PmtCondSts_GrntedPmt:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_PmtCondSts_EarlyPmt:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Agt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_DbtrDcsnDtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_AccptncDtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_AcctSvcrRef:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ClrSysRef:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_InstdAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_InstdAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_EqvtAmt_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_EqvtAmt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_EqvtAmt_CcyOfTrf:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_ReqdExctnDt_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_ReqdExctnDt_DtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_XpryDt_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_XpryDt_DtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtCond_AmtModAllwd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtCond_EarlyPmtAllwd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtCond_DelyPnlty:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtCond_ImdtPmtRbt_Amt_Ccy:
		d, _ := dest.(*common.ActiveCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtCond_ImdtPmtRbt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtCond_ImdtPmtRbt_Rate:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtCond_GrntedPmtReqd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_InstrPrty:
		d, _ := dest.(*common.Priority2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Priority2Code data for" + docPath)
		}

		*d, err = common.ToPriority2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_SvcLvl_Cd:
		d, _ := dest.(*common.ExternalServiceLevel1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalServiceLevel1Code data for" + docPath)
		}

		*d, err = common.ToExternalServiceLevel1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_SvcLvl_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_LclInstrm_Cd:
		d, _ := dest.(*common.ExternalLocalInstrument1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalLocalInstrument1Code data for" + docPath)
		}

		*d, err = common.ToExternalLocalInstrument1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_LclInstrm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_CtgyPurp_Cd:
		d, _ := dest.(*common.ExternalCategoryPurpose1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCategoryPurpose1Code data for" + docPath)
		}

		*d, err = common.ToExternalCategoryPurpose1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_CtgyPurp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtMtd:
		d, _ := dest.(*common.PaymentMethod4Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PaymentMethod4Code data for" + docPath)
		}

		*d, err = common.ToPaymentMethod4Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Ustrd:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Cd:
		d, _ := dest.(*common.DocumentType6Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.DocumentType6Code data for" + docPath)
		}

		*d, err = common.ToDocumentType6Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Nb:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_RltdDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry_Cd:
		d, _ := dest.(*common.ExternalDocumentLineType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDocumentLineType1Code data for" + docPath)
		}

		*d, err = common.ToExternalDocumentLineType1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Nb:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Id_RltdDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Desc:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp_Cd:
		d, _ := dest.(*common.ExternalDiscountAmountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDiscountAmountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalDiscountAmountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp_Cd:
		d, _ := dest.(*common.ExternalTaxAmountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalTaxAmountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalTaxAmountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_CdtDbtInd:
		d, _ := dest.(*common.CreditDebitCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CreditDebitCode data for" + docPath)
		}

		*d, err = common.ToCreditDebitCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Rsn:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_AddtlInf:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp_Cd:
		d, _ := dest.(*common.ExternalDiscountAmountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDiscountAmountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalDiscountAmountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp_Cd:
		d, _ := dest.(*common.ExternalTaxAmountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalTaxAmountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalTaxAmountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_CdtDbtInd:
		d, _ := dest.(*common.CreditDebitCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CreditDebitCode data for" + docPath)
		}

		*d, err = common.ToCreditDebitCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Rsn:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_AddtlInf:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Cd:
		d, _ := dest.(*common.DocumentType3Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.DocumentType3Code data for" + docPath)
		}

		*d, err = common.ToDocumentType3Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Ref:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Cdtr_TaxId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Cdtr_RegnId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Cdtr_TaxTp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Dbtr_TaxId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Dbtr_RegnId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Dbtr_TaxTp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Dbtr_Authstn_Titl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Dbtr_Authstn_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_UltmtDbtr_TaxId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_UltmtDbtr_RegnId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_UltmtDbtr_TaxTp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn_Titl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_AdmstnZone:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_RefNb:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Mtd:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_TtlTaxAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_TtlTaxAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_SeqNb:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_Tp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_Ctgy:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_CtgyDtls:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_DbtrSts:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_CertId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_FrmsCd:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_Prd_Yr:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_Prd_Tp:
		d, _ := dest.(*common.TaxRecordPeriod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.TaxRecordPeriod1Code data for" + docPath)
		}

		*d, err = common.ToTaxRecordPeriod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt_FrDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt_ToDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Rate:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_Yr:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_Tp:
		d, _ := dest.(*common.TaxRecordPeriod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.TaxRecordPeriod1Code data for" + docPath)
		}

		*d, err = common.ToTaxRecordPeriod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_AddtlInf:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry_Cd:
		d, _ := dest.(*common.ExternalGarnishmentType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalGarnishmentType1Code data for" + docPath)
		}

		*d, err = common.ToExternalGarnishmentType1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Tp_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_RefNb:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_RmtdAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_RmtdAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_FmlyMdclInsrncInd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_MplyeeTermntnInd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_AddtlRmtInf:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_Tp_Cd:
		d, _ := dest.(*common.ExternalDocumentType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDocumentType1Code data for" + docPath)
		}

		*d, err = common.ToExternalDocumentType1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_Tp_Prtry_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_Tp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_Tp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_IsseDt_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_IsseDt_DtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_LangCd:
		d, _ := dest.(*xsdt.String)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.String data for" + docPath)
		}

		*d, err = xsdt.ToString(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_Frmt_Cd:
		d, _ := dest.(*common.ExternalDocumentFormat1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDocumentFormat1Code data for" + docPath)
		}

		*d, err = common.ToExternalDocumentFormat1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_Frmt_Prtry_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_Frmt_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_Frmt_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_FileNm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Pty_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_DgtlSgntr_Sgntr_Item:
		d, _ := dest.(*xsdt.AnyType)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.AnyType data for" + docPath)
		}

		*d, err = xsdt.ToAnyType(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_NclsdFile_Nclsr:
		d, _ := dest.(*common.Max10MbBinary)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max10MbBinary data for" + docPath)
		}

		*d, err = common.ToMax10MbBinary(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_IBAN:
		d, _ := dest.(*common.IBAN2007Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*d, err = common.ToIBAN2007Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr_Id:
		d, _ := dest.(*common.Max34Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*d, err = common.ToMax34Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalAccountIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalAccountIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Tp_Cd:
		d, _ := dest.(*common.ExternalCashAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashAccountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Nm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Prxy_Tp_Cd:
		d, _ := dest.(*common.ExternalProxyAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalProxyAccountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Prxy_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Prxy_Id:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_IBAN:
		d, _ := dest.(*common.IBAN2007Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*d, err = common.ToIBAN2007Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr_Id:
		d, _ := dest.(*common.Max34Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*d, err = common.ToMax34Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalAccountIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalAccountIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Tp_Cd:
		d, _ := dest.(*common.ExternalCashAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashAccountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Nm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Prxy_Tp_Cd:
		d, _ := dest.(*common.ExternalProxyAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalProxyAccountType1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Prxy_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Prxy_Id:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_Tp_Cd:
		d, _ := dest.(*common.ExternalDocumentType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDocumentType1Code data for" + docPath)
		}

		*d, err = common.ToExternalDocumentType1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_Tp_Prtry_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_Tp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_Tp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_IsseDt_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_IsseDt_DtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_LangCd:
		d, _ := dest.(*xsdt.String)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.String data for" + docPath)
		}

		*d, err = xsdt.ToString(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_Frmt_Cd:
		d, _ := dest.(*common.ExternalDocumentFormat1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDocumentFormat1Code data for" + docPath)
		}

		*d, err = common.ToExternalDocumentFormat1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_Frmt_Prtry_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_Frmt_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_Frmt_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_FileNm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Pty_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_DgtlSgntr_Sgntr_Item:
		d, _ := dest.(*xsdt.AnyType)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.AnyType data for" + docPath)
		}

		*d, err = xsdt.ToAnyType(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_NclsdFile_Nclsr:
		d, _ := dest.(*common.Max10MbBinary)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max10MbBinary data for" + docPath)
		}

		*d, err = common.ToMax10MbBinary(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_SplmtryData_PlcAndNm:
		d, _ := dest.(*common.Max350Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max350Text data for" + docPath)
		}

		*d, err = common.ToMax350Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_SplmtryData_Envlp_Item:
		d, _ := dest.(*xsdt.AnyType)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.AnyType data for" + docPath)
		}

		*d, err = xsdt.ToAnyType(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_SplmtryData_PlcAndNm:
		d, _ := dest.(*common.Max350Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max350Text data for" + docPath)
		}

		*d, err = common.ToMax350Text(src)
		return err

	case Path_CdtrPmtActvtnReqStsRpt_SplmtryData_Envlp_Item:
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
