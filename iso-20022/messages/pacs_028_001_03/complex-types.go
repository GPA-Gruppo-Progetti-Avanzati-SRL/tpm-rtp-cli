// Package pacs_028_001_03
// Do not Edit. This stuff it's been automatically generated.
// Generated at 2022-04-05 22:58:59.627106 +0200 CEST m=+0.110044876
package pacs_028_001_03

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
)

// PaymentTransaction113 type definition
type PaymentTransaction113 struct {
	StsReqId        common.Max35Text                                     `xml:"StsReqId,omitempty"`
	OrgnlGrpInf     *common.OriginalGroupInformation29                   `xml:"OrgnlGrpInf,omitempty"`
	OrgnlInstrId    common.Max35Text                                     `xml:"OrgnlInstrId,omitempty"`
	OrgnlEndToEndId common.Max35Text                                     `xml:"OrgnlEndToEndId,omitempty"`
	OrgnlTxId       common.Max35Text                                     `xml:"OrgnlTxId,omitempty"`
	OrgnlUETR       common.UUIDv4Identifier                              `xml:"OrgnlUETR,omitempty"`
	AccptncDtTm     common.ISODateTime                                   `xml:"AccptncDtTm,omitempty"`
	ClrSysRef       common.Max35Text                                     `xml:"ClrSysRef,omitempty"`
	InstgAgt        *common.BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty"`
	InstdAgt        *common.BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty"`
	OrgnlTxRef      *common.OriginalTransactionReference28               `xml:"OrgnlTxRef,omitempty"`
	SplmtryData     []common.SupplementaryData1                          `xml:"SplmtryData,omitempty"`
}

// FIToFIPaymentStatusRequestV03 type definition
type FIToFIPaymentStatusRequestV03 struct {
	GrpHdr      GroupHeader91                `xml:"GrpHdr"`
	OrgnlGrpInf []OriginalGroupInformation27 `xml:"OrgnlGrpInf,omitempty"`
	TxInf       []PaymentTransaction113      `xml:"TxInf,omitempty"`
	SplmtryData []common.SupplementaryData1  `xml:"SplmtryData,omitempty"`
}

// OriginalGroupInformation27 type definition
type OriginalGroupInformation27 struct {
	OrgnlMsgId   common.Max35Text        `xml:"OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text        `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm common.ISODateTime      `xml:"OrgnlCreDtTm,omitempty"`
	OrgnlNbOfTxs common.Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum xsdt.Decimal            `xml:"OrgnlCtrlSum,omitempty"`
}

// GroupHeader91 type definition
type GroupHeader91 struct {
	MsgId    common.Max35Text                                     `xml:"MsgId"`
	CreDtTm  common.ISODateTime                                   `xml:"CreDtTm"`
	InstgAgt *common.BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty"`
	InstdAgt *common.BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty"`
}
