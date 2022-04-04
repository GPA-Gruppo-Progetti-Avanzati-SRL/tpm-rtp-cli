package examples_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/camt_055_001_08"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
	"github.com/stretchr/testify/require"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

const Example_Camt_055_001_08_1 = "example-document-camt_055_001_08.xml"

func TestDocumentCamt_055_001_08(t *testing.T) {

	d := camt_055_001_08.Document{
		CstmrPmtCxlReq: camt_055_001_08.CustomerPaymentCancellationRequestV08{
			Assgnmt: common.CaseAssignment5{
				Id: common.Max35Text("str12345"),
				Assgnr: common.Party40Choice{
					Pty: &common.PartyIdentification135{
						Nm: common.MustToMax140Text("AFC Poste Italiane"),
						Id: &common.Party38Choice{
							OrgId: &common.OrganisationIdentification29{
								Othr: []common.GenericOrganisationIdentification1{
									{
										Id: common.Max35Text("0468651441"),
										SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
											Cd: common.ExternalOrganisationIdentification1Code("BOID"),
										},
									},
								},
							},
						},
					},
				},
				Assgne: common.Party40Choice{
					Agt: &common.BranchAndFinancialInstitutionIdentification6{
						FinInstnId: common.FinancialInstitutionIdentification18{
							BICFI: common.BICFIDec2014Identifier("BPPIITRRXXX"),
						},
					},
				},
				CreDtTm: common.MustToISODateTime(time.Now()),
			},
			Undrlyg: []camt_055_001_08.UnderlyingTransaction24{
				{
					OrgnlGrpInfAndCxl: &camt_055_001_08.OriginalGroupHeader15{
						GrpCxlId:     common.Max35Text("camt055-DS10-20220330"),
						OrgnlMsgId:   common.Max35Text("pacs013-DS01-20220322"),
						OrgnlMsgNmId: common.Max35Text("pain.013.001.07"),
						OrgnlCreDtTm: common.MustToISODateTime(time.Now()),
						NbOfTxs:      common.Max15NumericText("1"),
						CtrlSum:      xsdt.ToDecimal(535.25),
					},
					OrgnlPmtInfAndCxl: []camt_055_001_08.OriginalPaymentInstruction34{
						{
							OrgnlPmtInfId: common.Max35Text("Fatt2022-000001-2022-03-22-11:50:45"),
							TxInf: []camt_055_001_08.PaymentTransaction109{
								{
									CxlId:           common.Max35Text("CancellationId1234"),
									OrgnlEndToEndId: common.Max35Text("fatt2022-000001"),
									CxlRsnInf: []camt_055_001_08.PaymentCancellationReason5{
										{
											Orgtr: &common.PartyIdentification135{
												Nm: common.Max140Text("AFC Poste Italiane"),
												Id: &common.Party38Choice{
													OrgId: &common.OrganisationIdentification29{
														Othr: []common.GenericOrganisationIdentification1{
															{
																Id: common.Max35Text("0468651441"),
																SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
																	Cd: common.ExternalOrganisationIdentification1Code("BOID"),
																},
															},
														},
													},
												},
											},
											Rsn: &camt_055_001_08.CancellationReason33Choice{
												Cd: common.ExternalCancellationReason1Code("errore nel precedente importo della fattura"),
											},
										},
									},
									OrgnlTxRef: &common.OriginalTransactionReference28{
										Amt: &common.AmountType4Choice{
											InstdAmt: &common.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.ActiveOrHistoricCurrencyCode("EUR"),
												Value: xsdt.ToDecimal(535.25),
											},
										},
										ReqdExctnDt: &common.DateAndDateTime2Choice{
											Dt: common.MustToISODate(time.Now()),
										},
										PmtTpInf: &common.PaymentTypeInformation27{
											SvcLvl: []common.ServiceLevel8Choice{
												{
													Cd: common.ExternalServiceLevel1Code("SEPA"),
												},
											},
											LclInstrm: &common.LocalInstrument2Choice{
												Prtry: common.Max35Text("NOTPROVIDED"),
											},
										},
										RmtInf: &common.RemittanceInformation16{
											Ustrd: []common.Max140Text{
												common.Max140Text("AT41/fatt2022-000001/AT05/pagamento fattura"),
											},
										},
										Cdtr: &common.Party40Choice{
											Pty: &common.PartyIdentification135{
												Nm: common.Max140Text("AFC Poste Italiane"),
												Id: &common.Party38Choice{
													OrgId: &common.OrganisationIdentification29{
														Othr: []common.GenericOrganisationIdentification1{
															{
																Id: common.Max35Text("0468651441"),
																SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
																	Cd: common.ExternalOrganisationIdentification1Code("BOID"),
																},
															},
														},
													},
												},
											},
										},
										CdtrAcct: &common.CashAccount38{
											Id: common.AccountIdentification4Choice{
												IBAN: common.IBAN2007Identifier("IT60X0760111101000004545561"),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	b, err := d.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_Camt_055_001_08_1, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_Camt_055_001_08_1)

}
