package docmapper_test

import (
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/docmapper"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/pain_013_001_07"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

const (
	ISO20022MessageFile = "iso20022-message.xml"
)

func TestMap2Pain013Doc(t *testing.T) {

	// In the normal case use the NewMapppingClass and avoid the WithFuncMap(nil) call that sets the builtins...
	dm, err := docmapper.NewMapperClass()
	require.NoError(t, err)

	dm.Rules = []docmapper.MappingRule{
		{
			Name:        "msg-id",
			SourceValue: "{$.msg-id}",
			MapFunc:     "trimSpace",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_MsgId,
		},
		{
			Name:        "from-name",
			SourceValue: "{$.header.from.name}",
			MapFunc:     "trimSpace",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm,
		},
		{
			Name:        "creation-date",
			SourceValue: "{$.header.creation-date}",
			MapFunc:     "trimSpace",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_CreDtTm,
		},
		{
			Name:        "number-of-txs",
			SourceValue: "2",
			MapFunc:     "trimSpace",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_NbOfTxs,
		},
	}
	t.Log(dm)

	m := docmapper.MappableDocMap{
		"msg-id": "pain013-DS01-20220322",
		"header": map[string]interface{}{
			"creation-date": time.Now().Format(time.RFC3339),
			"number-of-txs": "1",
			"from": map[string]interface{}{
				"name": "ACME Corp.",
			},
		},
	}

	d := pain_013_001_07.NewDocument()
	err = dm.Map(m, &d)
	require.NoError(t, err)

	t.Log("is document valid? --> ", d.IsValid(false))

	b1, err := d.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(ISO20022MessageFile, b1, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(ISO20022MessageFile)
}

func TestPain013Doc2Map(t *testing.T) {

	// In the normal case use the NewMapppingClass and avoid the WithFuncMap(nil) call that sets the builtins...
	dm, err := docmapper.NewMapperClass()
	require.NoError(t, err)

	dm.Rules = []docmapper.MappingRule{
		{
			Name:        "msg-id",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_MsgId),
			MapFunc:     "trimSpace",
			TargetPath:  "msg-id",
		},
		{
			Name:        "from-name",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm),
			MapFunc:     "trimSpace",
			TargetPath:  "header.from.name",
		},
		{
			Name:        "creation-date",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_CreDtTm),
			MapFunc:     "trimSpace",
			TargetPath:  "header.creation-date",
		},
		{
			Name:        "number-of-txs",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_NbOfTxs),
			MapFunc:     "trimSpace",
			TargetPath:  "header.number-of-txs",
		},
		{
			Name:        "pmt-info-amount",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Ccy),
			MapFunc:     "trimSpace",
			TargetPath:  "message.pmtinfo.amount.currency",
		},
		{
			Name:        "pmt-info-amount",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Value),
			MapFunc:     "trimSpace",
			TargetPath:  "message.pmtinfo.amount.value",
		},
	}

	d := pain_013_001_07.NewDocument()
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_MsgId, common.MustToMax35Text("pain013-DS01-20220322"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_CreDtTm, common.MustToISODateTime(time.Now()))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_NbOfTxs, common.MustToMax15NumericText("7"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm, common.MustToMax140Text("ACME Corp."))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Value, "535.90")
	require.NoError(t, err)

	m := docmapper.MappableDocMap(make(map[string]interface{}))

	err = dm.Map(&d, m)
	require.NoError(t, err)

	b, err := jsoniter.MarshalIndent(m, "", " ")
	require.NoError(t, err)

	t.Log(string(b))
}
