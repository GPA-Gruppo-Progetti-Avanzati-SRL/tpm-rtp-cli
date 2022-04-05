// Package camt_055_001_08
// Do not Edit. This stuff it's been automatically generated.
// Generated at 2022-04-05 08:23:15.82911 +0200 CEST m=+0.108518417
package camt_055_001_08

import (
	"bytes"
	"encoding/xml"
)

// IsValid checks if ControlData1 is valid
func (s ControlData1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.NbOfTxs.IsValid(false)
	valid = valid && s.CtrlSum.IsValid(true)

	return valid
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
	valid = valid && s.CstmrPmtCxlReq.IsValid(false)

	return valid
}

// IsValid checks if CancellationReason33Choice is valid
func (s CancellationReason33Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if OriginalPaymentInstruction34 is valid
func (s OriginalPaymentInstruction34) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.PmtCxlId.IsValid(true)
	valid = valid && (s.Case == nil || (s.Case != nil && s.Case.IsValid(true)))

	valid = valid && s.OrgnlPmtInfId.IsValid(false)
	valid = valid && (s.OrgnlGrpInf == nil || (s.OrgnlGrpInf != nil && s.OrgnlGrpInf.IsValid(true)))

	valid = valid && s.NbOfTxs.IsValid(true)
	valid = valid && s.CtrlSum.IsValid(true)
	valid = valid && s.PmtInfCxl.IsValid(true)
	for j := 0; j < len(s.CxlRsnInf); j++ {
		valid = valid && s.CxlRsnInf[j].IsValid(true)
	}

	for j := 0; j < len(s.TxInf); j++ {
		valid = valid && s.TxInf[j].IsValid(true)
	}

	return valid
}

// IsValid checks if PaymentTransaction109 is valid
func (s PaymentTransaction109) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.CxlId.IsValid(true)
	valid = valid && (s.Case == nil || (s.Case != nil && s.Case.IsValid(true)))

	valid = valid && s.OrgnlInstrId.IsValid(true)
	valid = valid && s.OrgnlEndToEndId.IsValid(true)
	valid = valid && s.OrgnlUETR.IsValid(true)
	valid = valid && (s.OrgnlInstdAmt == nil || (s.OrgnlInstdAmt != nil && s.OrgnlInstdAmt.IsValid(true)))

	valid = valid && (s.OrgnlReqdExctnDt == nil || (s.OrgnlReqdExctnDt != nil && s.OrgnlReqdExctnDt.IsValid(true)))

	valid = valid && s.OrgnlReqdColltnDt.IsValid(true)
	for j := 0; j < len(s.CxlRsnInf); j++ {
		valid = valid && s.CxlRsnInf[j].IsValid(true)
	}

	valid = valid && (s.OrgnlTxRef == nil || (s.OrgnlTxRef != nil && s.OrgnlTxRef.IsValid(true)))

	for j := 0; j < len(s.SplmtryData); j++ {
		valid = valid && s.SplmtryData[j].IsValid(true)
	}

	return valid
}

// IsValid checks if CustomerPaymentCancellationRequestV08 is valid
func (s CustomerPaymentCancellationRequestV08) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Assgnmt.IsValid(false)
	valid = valid && (s.Case == nil || (s.Case != nil && s.Case.IsValid(true)))

	valid = valid && (s.CtrlData == nil || (s.CtrlData != nil && s.CtrlData.IsValid(true)))

	if len(s.Undrlyg) == 0 {
		valid = false
	}
	for j := 0; j < len(s.Undrlyg); j++ {
		valid = valid && s.Undrlyg[j].IsValid(false)
	}

	for j := 0; j < len(s.SplmtryData); j++ {
		valid = valid && s.SplmtryData[j].IsValid(true)
	}

	return valid
}

// IsValid checks if OriginalGroupHeader15 is valid
func (s OriginalGroupHeader15) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.GrpCxlId.IsValid(true)
	valid = valid && (s.Case == nil || (s.Case != nil && s.Case.IsValid(true)))

	valid = valid && s.OrgnlMsgId.IsValid(false)
	valid = valid && s.OrgnlMsgNmId.IsValid(false)
	valid = valid && s.OrgnlCreDtTm.IsValid(true)
	valid = valid && s.NbOfTxs.IsValid(true)
	valid = valid && s.CtrlSum.IsValid(true)
	valid = valid && s.GrpCxl.IsValid(true)
	for j := 0; j < len(s.CxlRsnInf); j++ {
		valid = valid && s.CxlRsnInf[j].IsValid(true)
	}

	return valid
}

// IsValid checks if UnderlyingTransaction24 is valid
func (s UnderlyingTransaction24) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.OrgnlGrpInfAndCxl == nil || (s.OrgnlGrpInfAndCxl != nil && s.OrgnlGrpInfAndCxl.IsValid(true)))

	for j := 0; j < len(s.OrgnlPmtInfAndCxl); j++ {
		valid = valid && s.OrgnlPmtInfAndCxl[j].IsValid(true)
	}

	return valid
}

// IsValid checks if PaymentCancellationReason5 is valid
func (s PaymentCancellationReason5) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Orgtr == nil || (s.Orgtr != nil && s.Orgtr.IsValid(true)))

	valid = valid && (s.Rsn == nil || (s.Rsn != nil && s.Rsn.IsValid(true)))

	for j := 0; j < len(s.AddtlInf); j++ {
		valid = valid && s.AddtlInf[j].IsValid(true)
	}

	return valid
}
