package main

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/docmapper"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/pacs_028_001_03"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
)

func Gen_AFC_DS15_Pacs028_001_03(fn string) error {

	// In the normal case use the NewMapppingClass and avoid the WithFuncMap(nil) call that sets the builtins...
	dm, err := docmapper.NewMapperClass("pacs.028.001.03", "AFC-DS15-pacs.028.001.03")
	if err != nil {
		return err
	}

	dm.Rules = []docmapper.MappingRule{
		// GrpHdr
		{
			Name:        "GrpHdr_MsgId",
			SourceValue: `{$._ids.msg_id}`,
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_GrpHdr_MsgId,
		},
		{
			Name:        "GrpHdr_CreDtTm",
			SourceValue: "{$._ids.cre_dt_tm}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_GrpHdr_CreDtTm,
		},
		{
			Name:        "GrpHdr_InstgAgt_FinInstnId_BICFI",
			SourceValue: "{$.payee-psp-bic}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_GrpHdr_InstgAgt_FinInstnId_BICFI,
		},
		{
			Name:        "GrpHdr_InstdAgt_FinInstnId_BICFI",
			SourceValue: "{$.payer-psp-id}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_GrpHdr_InstdAgt_FinInstnId_BICFI,
		},

		// OrgnlGrpInf
		{
			Name:        "OrgnlGrpInf_OrgnlMsgId",
			SourceValue: "{$._org_ids.msg_id}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_OrgnlGrpInf_OrgnlMsgId,
		},
		{
			Name:        "OrgnlGrpInf_OrgnlMsgNmId",
			SourceValue: "pain.013.001.07",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_OrgnlGrpInf_OrgnlMsgNmId,
		},
		{
			Name:        "OrgnlGrpInf_OrgnlCtrlSum",
			SourceValue: `decimal("{$.amount}")`,
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_OrgnlGrpInf_OrgnlCtrlSum,
		},
		{
			Name:        "OrgnlGrpInf_OrgnlNbOfTxs",
			SourceValue: "1",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_OrgnlGrpInf_OrgnlNbOfTxs,
		},
		{
			Name:        "OrgnlGrpInf_OrgnlCreDtTm",
			SourceValue: "{$._org_ids.cre_dt_tm}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_OrgnlGrpInf_OrgnlCreDtTm,
		},

		// TxInf
		{
			Name:        "TxInf_StsReqId",
			SourceValue: `{$._ids.lnk_id}`,
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_StsReqId,
		},
		{
			Name:        "TxInf_OrgnlEndToEndId",
			SourceValue: "{$.payee-e2e-rtp-ref}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlEndToEndId,
		},
		{
			Name:        "TxInf_OrgnlInstrId",
			SourceValue: "{$._org_ids.instr_id}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlInstrId,
		},

		{
			Name:        "TxInf_InstgAgt_FinInstnId_BICFI",
			SourceValue: "{$.payee-psp-bic}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_InstgAgt_FinInstnId_BICFI,
		},
		{
			Name:        "TxInf_InstdAgt_FinInstnId_BICFI",
			SourceValue: "{$.payer-psp-id}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_InstdAgt_FinInstnId_BICFI,
		},

		// TxInf_OrgnlTxRef
		{
			Name:        "TxInf_OrgnlTxRef_Amt_InstdAmt_Value",
			SourceValue: `decimal("{$.amount}")`,
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Amt_InstdAmt_Value,
		},
		{
			Name:        "TxInf_OrgnlTxRef_Amt_InstdAmt_Ccy",
			SourceValue: "EUR",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Amt_InstdAmt_Ccy,
		},
		{
			Name:        "TxInf_OrgnlTxRef_ReqdExctnDt_Dt",
			SourceValue: "{$.pmt-req-exec-date}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_ReqdExctnDt_Dt,
		},

		{
			Name:        "TxInf_OrgnlTxRef_PmtTpInf_SvcLvl_Cd",
			SourceValue: "SRTP",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_PmtTpInf_SvcLvl_Cd,
		},
		{
			Name:        "TxInf_OrgnlTxRef_PmtTpInf_LclInstrm_Prtry",
			SourceValue: "{$.pmt-instrument}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_PmtTpInf_LclInstrm_Prtry,
		},
		{
			Name:        "TxInf_OrgnlTxRef_PmtTpInf_CtgyPurp_Cd",
			SourceValue: "OTHR",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_PmtTpInf_CtgyPurp_Cd,
		},
		{
			Name:        "TxInf_OrgnlTxRef_RmtInf_Ustrd",
			SourceValue: "AT41/{$.payee-e2e-rtp-ref}",
			Guard:       `"{$.rmt-inf-1}" == ""`,
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Ustrd,
		},
		{
			Name:        "TxInf_OrgnlTxRef_RmtInf_Ustrd",
			SourceValue: "AT41/{$.payee-e2e-rtp-ref}/AT05/{$.rmt-inf-1}",
			Guard:       "{$.rmt-inf-1}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Ustrd,
		},
		{
			Name:        "TxInf_OrgnlTxRef_RmtInf_Ustrd",
			SourceValue: "AT77/{$.rtp-expiry-date}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_RmtInf_Ustrd,
		},

		{
			Name:        "TxInf_OrgnlTxRef_Dbtr_Pty_Id_PrvtId_Othr_Id",
			SourceValue: "{$.payer-id}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_PrvtId_Othr_Id,
		},
		{
			Name:        "TxInf_OrgnlTxRef_Dbtr_Pty_Id_PrvtId_Othr_SchmeNm_Cd",
			SourceValue: "POID",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Dbtr_Pty_Id_PrvtId_Othr_SchmeNm_Cd,
		},
		{
			Name:        "TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_Nm",
			SourceValue: "{$.payee-name}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_Nm,
		},
		{
			Name:        "TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_BICFI",
			SourceValue: "{$.payee-psp-bic}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_BICFI,
		},

		{
			Name:        "TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_Othr_Id",
			SourceValue: "{$.payee-id}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_Othr_Id,
		},
		{
			Name:        "TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_Othr_SchmeNm_Prtry",
			SourceValue: "BOID",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_Othr_SchmeNm_Prtry,
		},
		{
			Name:        "TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_Othr_Issr",
			SourceValue: "{$.payee-company-name} {$.payee-name}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_Cdtr_Agt_FinInstnId_Othr_Issr,
		},

		{
			Name:        "TxInf_OrgnlTxRef_CdtrAcct_Id_IBAN",
			SourceValue: "{$.payee-iban}",
			TargetPath:  pacs_028_001_03.Path_FIToFIPmtStsReq_TxInf_OrgnlTxRef_CdtrAcct_Id_IBAN,
		},
	}

	b, err := yaml.Marshal(dm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fn, b, fs.ModePerm)
	return err
}
