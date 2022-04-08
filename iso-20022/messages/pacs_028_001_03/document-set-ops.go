// Package pacs_028_001_03
// Do not Edit. This stuff it's been automatically generated.
package pacs_028_001_03

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

	case Path_FIToFIPmtStsReq_GrpHdr_MsgId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_CreDtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_OrgnlGrpInf_OrgnlMsgId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_OrgnlGrpInf_OrgnlMsgNmId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_OrgnlGrpInf_OrgnlCreDtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_FIToFIPmtStsReq_OrgnlGrpInf_OrgnlNbOfTxs:
		d, _ := dest.(*common.Max15NumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max15NumericText data for" + docPath)
		}

		*d, err = common.ToMax15NumericText(src)
		return err

	case Path_FIToFIPmtStsReq_OrgnlGrpInf_OrgnlCtrlSum:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_StsReqId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlGrpInf_OrgnlMsgId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlGrpInf_OrgnlMsgNmId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlGrpInf_OrgnlCreDtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlInstrId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlEndToEndId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlUETR:
		d, _ := dest.(*common.UUIDv4Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.UUIDv4Identifier data for" + docPath)
		}

		*d, err = common.ToUUIDv4Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_AccptncDtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_ClrSysRef:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstgAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_InstdAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_IntrBkSttlmAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_IntrBkSttlmAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Amt_InstdAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Amt_InstdAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Amt_EqvtAmt_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Amt_EqvtAmt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Amt_EqvtAmt_CcyOfTrf:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_IntrBkSttlmDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_ReqdColltnDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_ReqdExctnDt_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_ReqdExctnDt_DtTm:
		d, _ := dest.(*common.ISODateTime)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*d, err = common.ToISODateTime(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrSchmeId_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_SttlmMtd:
		d, _ := dest.(*common.SettlementMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.SettlementMethod1Code data for" + docPath)
		}

		*d, err = common.ToSettlementMethod1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_SttlmAcct_Id_IBAN:
		d, _ := dest.(*common.IBAN2007Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*d, err = common.ToIBAN2007Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr_Id:
		d, _ := dest.(*common.Max34Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*d, err = common.ToMax34Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalAccountIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalAccountIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_SttlmAcct_Tp_Cd:
		d, _ := dest.(*common.ExternalCashAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_SttlmAcct_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_SttlmAcct_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_SttlmAcct_Nm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_SttlmAcct_Prxy_Tp_Cd:
		d, _ := dest.(*common.ExternalProxyAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalProxyAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_SttlmAcct_Prxy_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_SttlmAcct_Prxy_Id:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ClrSys_Cd:
		d, _ := dest.(*common.ExternalCashClearingSystem1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashClearingSystem1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashClearingSystem1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ClrSys_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_IBAN:
		d, _ := dest.(*common.IBAN2007Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*d, err = common.ToIBAN2007Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr_Id:
		d, _ := dest.(*common.Max34Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*d, err = common.ToMax34Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalAccountIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalAccountIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Tp_Cd:
		d, _ := dest.(*common.ExternalCashAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Nm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Prxy_Tp_Cd:
		d, _ := dest.(*common.ExternalProxyAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalProxyAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Prxy_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Prxy_Id:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_IBAN:
		d, _ := dest.(*common.IBAN2007Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*d, err = common.ToIBAN2007Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr_Id:
		d, _ := dest.(*common.Max34Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*d, err = common.ToMax34Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalAccountIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalAccountIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Tp_Cd:
		d, _ := dest.(*common.ExternalCashAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Nm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Prxy_Tp_Cd:
		d, _ := dest.(*common.ExternalProxyAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalProxyAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Prxy_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Prxy_Id:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_IBAN:
		d, _ := dest.(*common.IBAN2007Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*d, err = common.ToIBAN2007Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr_Id:
		d, _ := dest.(*common.Max34Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*d, err = common.ToMax34Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalAccountIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalAccountIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Tp_Cd:
		d, _ := dest.(*common.ExternalCashAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Nm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Prxy_Tp_Cd:
		d, _ := dest.(*common.ExternalProxyAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalProxyAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Prxy_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Prxy_Id:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_PmtTpInf_InstrPrty:
		d, _ := dest.(*common.Priority2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Priority2Code data for" + docPath)
		}

		*d, err = common.ToPriority2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_PmtTpInf_ClrChanl:
		d, _ := dest.(*common.ClearingChannel2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ClearingChannel2Code data for" + docPath)
		}

		*d, err = common.ToClearingChannel2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_PmtTpInf_SvcLvl_Cd:
		d, _ := dest.(*common.ExternalServiceLevel1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalServiceLevel1Code data for" + docPath)
		}

		*d, err = common.ToExternalServiceLevel1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_PmtTpInf_SvcLvl_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_PmtTpInf_LclInstrm_Cd:
		d, _ := dest.(*common.ExternalLocalInstrument1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalLocalInstrument1Code data for" + docPath)
		}

		*d, err = common.ToExternalLocalInstrument1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_PmtTpInf_LclInstrm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_PmtTpInf_SeqTp:
		d, _ := dest.(*common.SequenceType3Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.SequenceType3Code data for" + docPath)
		}

		*d, err = common.ToSequenceType3Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_PmtTpInf_CtgyPurp_Cd:
		d, _ := dest.(*common.ExternalCategoryPurpose1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCategoryPurpose1Code data for" + docPath)
		}

		*d, err = common.ToExternalCategoryPurpose1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_PmtTpInf_CtgyPurp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_PmtMtd:
		d, _ := dest.(*common.PaymentMethod4Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PaymentMethod4Code data for" + docPath)
		}

		*d, err = common.ToPaymentMethod4Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_MndtId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_DtOfSgntr:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlMndtId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_IBAN:
		d, _ := dest.(*common.IBAN2007Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*d, err = common.ToIBAN2007Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_Id:
		d, _ := dest.(*common.Max34Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*d, err = common.ToMax34Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalAccountIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalAccountIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Tp_Cd:
		d, _ := dest.(*common.ExternalCashAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Nm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Prxy_Tp_Cd:
		d, _ := dest.(*common.ExternalProxyAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalProxyAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Prxy_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Prxy_Id:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_IBAN:
		d, _ := dest.(*common.IBAN2007Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*d, err = common.ToIBAN2007Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_Id:
		d, _ := dest.(*common.Max34Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*d, err = common.ToMax34Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalAccountIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalAccountIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Tp_Cd:
		d, _ := dest.(*common.ExternalCashAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Nm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Prxy_Tp_Cd:
		d, _ := dest.(*common.ExternalProxyAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalProxyAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Prxy_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Prxy_Id:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_IBAN:
		d, _ := dest.(*common.IBAN2007Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*d, err = common.ToIBAN2007Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_Id:
		d, _ := dest.(*common.Max34Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*d, err = common.ToMax34Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalAccountIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalAccountIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Tp_Cd:
		d, _ := dest.(*common.ExternalCashAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Nm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Prxy_Tp_Cd:
		d, _ := dest.(*common.ExternalProxyAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalProxyAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Prxy_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Prxy_Id:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlFnlColltnDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlFrqcy_Tp:
		d, _ := dest.(*common.Frequency6Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Frequency6Code data for" + docPath)
		}

		*d, err = common.ToFrequency6Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlFrqcy_Prd_Tp:
		d, _ := dest.(*common.Frequency6Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Frequency6Code data for" + docPath)
		}

		*d, err = common.ToFrequency6Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlFrqcy_Prd_CntPerPrd:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlFrqcy_PtInTm_Tp:
		d, _ := dest.(*common.Frequency6Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Frequency6Code data for" + docPath)
		}

		*d, err = common.ToFrequency6Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlFrqcy_PtInTm_PtInTm:
		d, _ := dest.(*common.Exact2NumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact2NumericText data for" + docPath)
		}

		*d, err = common.ToExact2NumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlRsn_Cd:
		d, _ := dest.(*common.ExternalMandateSetupReason1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalMandateSetupReason1Code data for" + docPath)
		}

		*d, err = common.ToExternalMandateSetupReason1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlRsn_Prtry:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlTrckgDays:
		d, _ := dest.(*common.Exact2NumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact2NumericText data for" + docPath)
		}

		*d, err = common.ToExact2NumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_ElctrncSgntr:
		d, _ := dest.(*common.Max1025Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max1025Text data for" + docPath)
		}

		*d, err = common.ToMax1025Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_FrstColltnDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_FnlColltnDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_Frqcy_Tp:
		d, _ := dest.(*common.Frequency6Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Frequency6Code data for" + docPath)
		}

		*d, err = common.ToFrequency6Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_Frqcy_Prd_Tp:
		d, _ := dest.(*common.Frequency6Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Frequency6Code data for" + docPath)
		}

		*d, err = common.ToFrequency6Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_Frqcy_Prd_CntPerPrd:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_Frqcy_PtInTm_Tp:
		d, _ := dest.(*common.Frequency6Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Frequency6Code data for" + docPath)
		}

		*d, err = common.ToFrequency6Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_Frqcy_PtInTm_PtInTm:
		d, _ := dest.(*common.Exact2NumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact2NumericText data for" + docPath)
		}

		*d, err = common.ToExact2NumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_Rsn_Cd:
		d, _ := dest.(*common.ExternalMandateSetupReason1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalMandateSetupReason1Code data for" + docPath)
		}

		*d, err = common.ToExternalMandateSetupReason1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_Rsn_Prtry:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_MndtRltdInf_TrckgDays:
		d, _ := dest.(*common.Exact2NumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact2NumericText data for" + docPath)
		}

		*d, err = common.ToExact2NumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Ustrd:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Cd:
		d, _ := dest.(*common.DocumentType6Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.DocumentType6Code data for" + docPath)
		}

		*d, err = common.ToDocumentType6Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Nb:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_RltdDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry_Cd:
		d, _ := dest.(*common.ExternalDocumentLineType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDocumentLineType1Code data for" + docPath)
		}

		*d, err = common.ToExternalDocumentLineType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Nb:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Id_RltdDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Desc:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp_Cd:
		d, _ := dest.(*common.ExternalDiscountAmountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDiscountAmountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalDiscountAmountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp_Cd:
		d, _ := dest.(*common.ExternalTaxAmountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalTaxAmountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalTaxAmountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_CdtDbtInd:
		d, _ := dest.(*common.CreditDebitCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CreditDebitCode data for" + docPath)
		}

		*d, err = common.ToCreditDebitCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Rsn:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_AddtlInf:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp_Cd:
		d, _ := dest.(*common.ExternalDiscountAmountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDiscountAmountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalDiscountAmountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp_Cd:
		d, _ := dest.(*common.ExternalTaxAmountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalTaxAmountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalTaxAmountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_CdtDbtInd:
		d, _ := dest.(*common.CreditDebitCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CreditDebitCode data for" + docPath)
		}

		*d, err = common.ToCreditDebitCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Rsn:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_AddtlInf:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Cd:
		d, _ := dest.(*common.DocumentType3Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.DocumentType3Code data for" + docPath)
		}

		*d, err = common.ToDocumentType3Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Ref:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Cdtr_TaxId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Cdtr_RegnId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Cdtr_TaxTp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Dbtr_TaxId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Dbtr_RegnId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Dbtr_TaxTp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Dbtr_Authstn_Titl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Dbtr_Authstn_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_UltmtDbtr_TaxId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_UltmtDbtr_RegnId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_UltmtDbtr_TaxTp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn_Titl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_AdmstnZone:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_RefNb:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Mtd:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_TtlTaxAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_TtlTaxAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_SeqNb:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_Tp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_Ctgy:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_CtgyDtls:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_DbtrSts:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_CertId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_FrmsCd:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_Prd_Yr:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_Prd_Tp:
		d, _ := dest.(*common.TaxRecordPeriod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.TaxRecordPeriod1Code data for" + docPath)
		}

		*d, err = common.ToTaxRecordPeriod1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt_FrDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt_ToDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Rate:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_Yr:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_Tp:
		d, _ := dest.(*common.TaxRecordPeriod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.TaxRecordPeriod1Code data for" + docPath)
		}

		*d, err = common.ToTaxRecordPeriod1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_TaxRmt_Rcrd_AddtlInf:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry_Cd:
		d, _ := dest.(*common.ExternalGarnishmentType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalGarnishmentType1Code data for" + docPath)
		}

		*d, err = common.ToExternalGarnishmentType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Tp_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_RefNb:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_Dt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_RmtdAmt_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_RmtdAmt_Value:
		d, _ := dest.(*xsdt.Decimal)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*d, err = xsdt.ToDecimal(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_FmlyMdclInsrncInd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_GrnshmtRmt_MplyeeTermntnInd:
		d, _ := dest.(*xsdt.Boolean)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*d, err = xsdt.ToBoolean(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Strd_AddtlRmtInf:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Pty_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtDbtr_Agt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Agt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAcct_Id_IBAN:
		d, _ := dest.(*common.IBAN2007Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*d, err = common.ToIBAN2007Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAcct_Id_Othr_Id:
		d, _ := dest.(*common.Max34Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*d, err = common.ToMax34Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAcct_Id_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalAccountIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalAccountIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAcct_Id_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAcct_Id_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAcct_Tp_Cd:
		d, _ := dest.(*common.ExternalCashAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAcct_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAcct_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAcct_Nm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAcct_Prxy_Tp_Cd:
		d, _ := dest.(*common.ExternalProxyAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalProxyAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAcct_Prxy_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAcct_Prxy_Id:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgtAcct_Id_IBAN:
		d, _ := dest.(*common.IBAN2007Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*d, err = common.ToIBAN2007Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgtAcct_Id_Othr_Id:
		d, _ := dest.(*common.Max34Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*d, err = common.ToMax34Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgtAcct_Id_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalAccountIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalAccountIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgtAcct_Id_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgtAcct_Id_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgtAcct_Tp_Cd:
		d, _ := dest.(*common.ExternalCashAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgtAcct_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgtAcct_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgtAcct_Nm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgtAcct_Prxy_Tp_Cd:
		d, _ := dest.(*common.ExternalProxyAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalProxyAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgtAcct_Prxy_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_DbtrAgtAcct_Prxy_Id:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgtAcct_Id_IBAN:
		d, _ := dest.(*common.IBAN2007Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*d, err = common.ToIBAN2007Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgtAcct_Id_Othr_Id:
		d, _ := dest.(*common.Max34Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*d, err = common.ToMax34Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgtAcct_Id_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalAccountIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalAccountIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgtAcct_Id_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgtAcct_Id_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgtAcct_Tp_Cd:
		d, _ := dest.(*common.ExternalCashAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgtAcct_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgtAcct_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgtAcct_Nm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgtAcct_Prxy_Tp_Cd:
		d, _ := dest.(*common.ExternalProxyAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalProxyAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgtAcct_Prxy_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAgtAcct_Prxy_Id:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Pty_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAcct_Id_IBAN:
		d, _ := dest.(*common.IBAN2007Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*d, err = common.ToIBAN2007Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAcct_Id_Othr_Id:
		d, _ := dest.(*common.Max34Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*d, err = common.ToMax34Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAcct_Id_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalAccountIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalAccountIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAcct_Id_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAcct_Id_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAcct_Tp_Cd:
		d, _ := dest.(*common.ExternalCashAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalCashAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAcct_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAcct_Ccy:
		d, _ := dest.(*common.ActiveOrHistoricCurrencyCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*d, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAcct_Nm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAcct_Prxy_Tp_Cd:
		d, _ := dest.(*common.ExternalProxyAccountType1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*d, err = common.ToExternalProxyAccountType1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAcct_Prxy_Tp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAcct_Prxy_Id:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_Id_OrgId_AnyBIC:
		d, _ := dest.(*common.AnyBICDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToAnyBICDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_Id_OrgId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_Id_OrgId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_Id_OrgId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalOrganisationIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalOrganisationIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_Id_OrgId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_Id_OrgId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:
		d, _ := dest.(*common.ISODate)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*d, err = common.ToISODate(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_Id_PrvtId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_Id_PrvtId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalPersonIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalPersonIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_Id_PrvtId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_Id_PrvtId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_CtryOfRes:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_CtctDtls_NmPrfx:
		d, _ := dest.(*common.NamePrefix2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*d, err = common.ToNamePrefix2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_CtctDtls_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_CtctDtls_PhneNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_CtctDtls_MobNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_CtctDtls_FaxNb:
		d, _ := dest.(*common.PhoneNumber)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*d, err = common.ToPhoneNumber(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_CtctDtls_EmailAdr:
		d, _ := dest.(*common.Max2048Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*d, err = common.ToMax2048Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_CtctDtls_EmailPurp:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_CtctDtls_JobTitl:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_CtctDtls_Rspnsblty:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_CtctDtls_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_CtctDtls_Othr_ChanlTp:
		d, _ := dest.(*common.Max4Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*d, err = common.ToMax4Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_CtctDtls_Othr_Id:
		d, _ := dest.(*common.Max128Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*d, err = common.ToMax128Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Pty_CtctDtls_PrefrdMtd:
		d, _ := dest.(*common.PreferredContactMethod1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*d, err = common.ToPreferredContactMethod1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_BICFI:
		d, _ := dest.(*common.BICFIDec2014Identifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*d, err = common.ToBICFIDec2014Identifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:
		d, _ := dest.(*common.ExternalClearingSystemIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_ClrSysMmbId_MmbId:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_Othr_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_Othr_SchmeNm_Cd:
		d, _ := dest.(*common.ExternalFinancialInstitutionIdentification1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*d, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_Othr_SchmeNm_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_FinInstnId_Othr_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_Id:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_LEI:
		d, _ := dest.(*common.LEIIdentifier)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*d, err = common.ToLEIIdentifier(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_Nm:
		d, _ := dest.(*common.Max140Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*d, err = common.ToMax140Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_AdrTp_Cd:
		d, _ := dest.(*common.AddressType2Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*d, err = common.ToAddressType2Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_AdrTp_Prtry_Id:
		d, _ := dest.(*common.Exact4AlphaNumericText)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*d, err = common.ToExact4AlphaNumericText(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_Dept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_SubDept:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_StrtNm:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_BldgNb:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_BldgNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_Flr:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_PstBx:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_Room:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_PstCd:
		d, _ := dest.(*common.Max16Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*d, err = common.ToMax16Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_TwnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_TwnLctnNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_DstrctNm:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_CtrySubDvsn:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_Ctry:
		d, _ := dest.(*common.CountryCode)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*d, err = common.ToCountryCode(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_UltmtCdtr_Agt_BrnchId_PstlAdr_AdrLine:
		d, _ := dest.(*common.Max70Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*d, err = common.ToMax70Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Purp_Cd:
		d, _ := dest.(*common.ExternalPurpose1Code)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPurpose1Code data for" + docPath)
		}

		*d, err = common.ToExternalPurpose1Code(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Purp_Prtry:
		d, _ := dest.(*common.Max35Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*d, err = common.ToMax35Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_SplmtryData_PlcAndNm:
		d, _ := dest.(*common.Max350Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max350Text data for" + docPath)
		}

		*d, err = common.ToMax350Text(src)
		return err

	case Path_FIToFIPmtStsReq_TxInf_SplmtryData_Envlp_Item:
		d, _ := dest.(*xsdt.AnyType)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.AnyType data for" + docPath)
		}

		*d, err = xsdt.ToAnyType(src)
		return err

	case Path_FIToFIPmtStsReq_SplmtryData_PlcAndNm:
		d, _ := dest.(*common.Max350Text)
		if d == nil {
			return errors.New("nil pointer... in unmarshalling common.Max350Text data for" + docPath)
		}

		*d, err = common.ToMax350Text(src)
		return err

	case Path_FIToFIPmtStsReq_SplmtryData_Envlp_Item:
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
