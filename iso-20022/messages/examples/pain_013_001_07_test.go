package examples_test

import (
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/pain_013_001_07"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
	"github.com/stretchr/testify/require"
	"io/fs"
	"io/ioutil"
	"os"

	"testing"
	"time"
)

func TestDocumentPain_013_001_07(t *testing.T) {

	d := pain_013_001_07.Document{
		CdtrPmtActvtnReq: pain_013_001_07.CreditorPaymentActivationRequestV07{

			GrpHdr: pain_013_001_07.GroupHeader78{
				MsgId:   "pain013-DS01-20220322",
				CreDtTm: common.ISODateTime(time.Now().Format(time.RFC3339)),
				NbOfTxs: "1",
				InitgPty: common.PartyIdentification135{
					Nm: "Poste Italiane",
					Id: &common.Party38Choice{
						OrgId: &common.OrganisationIdentification29{
							Othr: []common.GenericOrganisationIdentification1{
								{
									Id:   "01114601006",
									Issr: "POSTE ITALIANE",
								},
							},
						},
					},
				},
			},
			PmtInf: []pain_013_001_07.PaymentInstruction31{
				{
					PmtInfId: "Fatt2022-000001-2022-03-22-11:50:45",
					PmtMtd:   "TRF",
					PmtTpInf: &common.PaymentTypeInformation26{
						SvcLvl: []common.ServiceLevel8Choice{
							{
								Cd: "SRTP",
							},
						},
						LclInstrm: &common.LocalInstrument2Choice{
							Prtry: "NOTPROVIDED",
						},
						CtgyPurp: &common.CategoryPurpose1Choice{
							Cd: "OTHR",
						},
					},
					ReqdExctnDt: common.DateAndDateTime2Choice{
						Dt: "2022-03-31",
					},
					XpryDt: &common.DateAndDateTime2Choice{
						Dt: "2022-03-31",
					},
					Dbtr: common.PartyIdentification135{
						Nm: "LoremIpsumSPA",
						PstlAdr: &common.PostalAddress24{
							AdrLine: []common.Max70Text{
								"ViaLoremIpsum 30 Roma",
							},
						},
						Id: &common.Party38Choice{
							PrvtId: &common.PersonIdentification13{
								Othr: []common.GenericPersonIdentification1{
									{
										Id: "123456789",
										SchmeNm: &common.PersonIdentificationSchemeName1Choice{
											Cd: "POID",
										},
									},
								},
							},
						},
					},
					DbtrAcct: &common.CashAccount38{
						Id: common.AccountIdentification4Choice{
							IBAN: "IT60X0760111101000000123456",
						},
					},
					DbtrAgt: common.BranchAndFinancialInstitutionIdentification6{
						FinInstnId: common.FinancialInstitutionIdentification18{
							BICFI: "BPPIITRRXXX",
						},
					},
					CdtTrfTx: []pain_013_001_07.CreditTransferTransaction35{
						{
							PmtId: pain_013_001_07.PaymentIdentification6{
								EndToEndId: "fatt2022-000001",
							},
							PmtTpInf: &common.PaymentTypeInformation26{
								SvcLvl: []common.ServiceLevel8Choice{
									{
										Cd: "SRTP",
									},
								},
								LclInstrm: &common.LocalInstrument2Choice{
									Prtry: "NOTPROVIDED",
								},
								CtgyPurp: &common.CategoryPurpose1Choice{
									Cd: "OTHR",
								},
							},
							PmtCond: &common.PaymentCondition1{
								AmtModAllwd:   false,
								EarlyPmtAllwd: false,
								GrntedPmtReqd: false,
							},
							Amt: common.AmountType4Choice{
								InstdAmt: &common.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   "EUR",
									Value: xsdt.Decimal(fmt.Sprintf("%f", 535.25)),
								},
							},
							ChrgBr: "SLEV",
							CdtrAgt: common.BranchAndFinancialInstitutionIdentification6{
								FinInstnId: common.FinancialInstitutionIdentification18{
									BICFI: "BPPIITRRXXX",
								},
							},
							Cdtr: common.PartyIdentification135{
								Nm: "AFC Poste Italiane",
								PstlAdr: &common.PostalAddress24{
									AdrLine: []common.Max70Text{
										"Via Del Creditore 75 Roma",
									},
								},
								Id: &common.Party38Choice{
									OrgId: &common.OrganisationIdentification29{
										Othr: []common.GenericOrganisationIdentification1{
											{
												Id: "0468651441",
												SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
													Cd: "BOID",
												},
											},
										},
									},
								},
							},
							CdtrAcct: &common.CashAccount38{
								Id: common.AccountIdentification4Choice{
									IBAN: "IT60X0760111101000004545561",
								},
							},
							RmtInf: &common.RemittanceInformation16{
								Ustrd: []common.Max140Text{
									"AT41/fatt2022-000001/AT05/pagamento fattura",
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

	err = ioutil.WriteFile("test-document.xml", b, fs.ModePerm)
	require.NoError(t, err)

	err = os.Remove("test-document.xml")
	require.NoError(t, err)
}
