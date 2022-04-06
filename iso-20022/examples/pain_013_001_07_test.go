package examples_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/pain_013_001_07"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
	"github.com/stretchr/testify/require"
	"io/fs"
	"io/ioutil"
	"os"
	"time"

	"testing"
)

const Example_Pain_013_001_07_1 = "example-document-pain_013_001_07-1.xml"
const Example_Pain_013_001_07_2 = "example-document-pain_013_001_07-2.xml"

func TestDocumentPain_013_001_07(t *testing.T) {

	d := pain_013_001_07.Document{
		CdtrPmtActvtnReq: pain_013_001_07.CreditorPaymentActivationRequestV07{

			GrpHdr: pain_013_001_07.GroupHeader78{
				MsgId:   common.MustToMax35Text("pain013-DS01-20220322"),
				CreDtTm: common.MustToISODateTime(time.Now()),
				NbOfTxs: common.MustToMax15NumericText("1"),
				InitgPty: common.PartyIdentification135{
					Nm: common.MustToMax140Text("Poste Italiane"),
					Id: &common.Party38Choice{
						OrgId: &common.OrganisationIdentification29{
							Othr: []common.GenericOrganisationIdentification1{
								{
									Id:   common.MustToMax35Text("01114601006"),
									Issr: common.MustToMax35Text("POSTE ITALIANE"),
								},
							},
						},
					},
				},
			},
			PmtInf: []pain_013_001_07.PaymentInstruction31{
				{
					PmtInfId: common.MustToMax35Text("Fatt2022-000001-2022-03-22-11:50:45"),
					PmtMtd:   common.MustToPaymentMethod7Code(common.PaymentMethod7CodeTRF),
					PmtTpInf: &common.PaymentTypeInformation26{
						SvcLvl: []common.ServiceLevel8Choice{
							{
								Cd: common.MustToExternalServiceLevel1Code("SRTP"),
							},
						},
						LclInstrm: &common.LocalInstrument2Choice{
							Prtry: common.MustToMax35Text("NOTPROVIDED"),
						},
						CtgyPurp: &common.CategoryPurpose1Choice{
							Cd: common.MustToExternalCategoryPurpose1Code("OTHR"),
						},
					},
					ReqdExctnDt: common.DateAndDateTime2Choice{
						Dt: common.MustToISODate(time.Now()),
					},
					XpryDt: &common.DateAndDateTime2Choice{
						Dt: common.MustToISODate(time.Now()),
					},
					Dbtr: common.PartyIdentification135{
						Nm: common.MustToMax140Text("LoremIpsumSPA"),
						PstlAdr: &common.PostalAddress24{
							AdrLine: []common.Max70Text{
								common.MustToMax70Text("ViaLoremIpsum 30 Roma"),
							},
						},
						Id: &common.Party38Choice{
							PrvtId: &common.PersonIdentification13{
								Othr: []common.GenericPersonIdentification1{
									{
										Id: common.MustToMax35Text("123456789"),
										SchmeNm: &common.PersonIdentificationSchemeName1Choice{
											Cd: common.MustToExternalPersonIdentification1Code("POID"),
										},
									},
								},
							},
						},
					},
					DbtrAcct: &common.CashAccount38{
						Id: common.AccountIdentification4Choice{
							IBAN: common.MustToIBAN2007Identifier("IT60X0760111101000000123456"),
						},
					},
					DbtrAgt: common.BranchAndFinancialInstitutionIdentification6{
						FinInstnId: common.FinancialInstitutionIdentification18{
							BICFI: common.MustToBICFIDec2014Identifier("BPPIITRRXXX"),
						},
					},
					CdtTrfTx: []pain_013_001_07.CreditTransferTransaction35{
						{
							PmtId: pain_013_001_07.PaymentIdentification6{
								EndToEndId: common.MustToMax35Text("fatt2022-000001"),
							},
							PmtTpInf: &common.PaymentTypeInformation26{
								SvcLvl: []common.ServiceLevel8Choice{
									{
										Cd: common.MustToExternalServiceLevel1Code("SRTP"),
									},
								},
								LclInstrm: &common.LocalInstrument2Choice{
									Prtry: common.MustToMax35Text("NOTPROVIDED"),
								},
								CtgyPurp: &common.CategoryPurpose1Choice{
									Cd: common.MustToExternalCategoryPurpose1Code("OTHR"),
								},
							},
							PmtCond: &common.PaymentCondition1{
								AmtModAllwd:   false,
								EarlyPmtAllwd: false,
								GrntedPmtReqd: false,
							},
							Amt: common.AmountType4Choice{
								InstdAmt: &common.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode("EUR"),
									Value: xsdt.MustToDecimal(535.35),
								},
							},
							ChrgBr: common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSLEV),
							CdtrAgt: common.BranchAndFinancialInstitutionIdentification6{
								FinInstnId: common.FinancialInstitutionIdentification18{
									BICFI: common.MustToBICFIDec2014Identifier("BPPIITRRXXX"),
								},
							},
							Cdtr: common.PartyIdentification135{
								Nm: common.MustToMax140Text("AFC Poste Italiane"),
								PstlAdr: &common.PostalAddress24{
									AdrLine: []common.Max70Text{
										common.MustToMax70Text("Via Del Creditore 75 Roma"),
									},
								},
								Id: &common.Party38Choice{
									OrgId: &common.OrganisationIdentification29{
										Othr: []common.GenericOrganisationIdentification1{
											{
												Id: common.MustToMax35Text("0468651441"),
												SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
													Cd: common.MustToExternalOrganisationIdentification1Code("BOID"),
												},
											},
										},
									},
								},
							},
							CdtrAcct: &common.CashAccount38{
								Id: common.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier("IT60X0760111101000004545561"),
								},
							},
							RmtInf: &common.RemittanceInformation16{
								Ustrd: []common.Max140Text{
									common.MustToMax140Text("AT41/fatt2022-000001/AT05/pagamento fattura"),
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

	err = ioutil.WriteFile(Example_Pain_013_001_07_1, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_Pain_013_001_07_1)

	d2, err := pain_013_001_07.NewDocumentFromXML(b)
	require.NoError(t, err)

	b2, err := d2.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_Pain_013_001_07_2, b2, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_Pain_013_001_07_2)

}

func TestDocumentPain_013_001_07_SetOps(t *testing.T) {
	d := pain_013_001_07.Document{}
	d.With_CdtrPmtActvtnReq_GrpHdr_MsgId(common.MustToMax35Text("pain013-DS01-20220322"))

	b, err := d.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_Pain_013_001_07_1, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_Pain_013_001_07_1)

}
