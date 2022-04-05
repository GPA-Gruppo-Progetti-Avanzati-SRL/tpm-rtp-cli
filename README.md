## tpm-rtp-cli

### What's that and why doing it
This project main objective is the generation of artifacts to manipulate IS20022 messages using Golang code. 
On the net there is stuff to generate go structs from schemas but all this stuff lacks some bells and whistles to easily use this structures. This generation is not general
if we mean a general generation of structs from schema files (it lacks some support of builtin types that are not used in IS20022 schemas for example) but tries to offer an approach useful for the task at hand.
For example, it doesn't generate a schme ato a single file in a single package but tries to find common types across different messages and split them in different packages.
As such, you will have a package for each message and a common package where simple and complex types shared among messages are put together to simplify the creation of a message from a different one used as a source.
More over it provides some convenience methods for setting and getting values out of the structures with some protection for arrays and nil pointers. 
I hope you will find this stuff edible to get life easier down the road considering that:

* arrays are a matter of the schema where plurality gets in the way,
* pointers are a fact of life when generating empty structs is not an option
* those messages can have a number of nodes ranginbg from 3000 up to 8500 (considering leaves and intermediate)
* real cases in general set a very few number of nodes: might be in the range of 30 or 40 leaves: more intermediate nodes than leaves as to speak. As such these trees are pretty sparse.
* the code to map some data to a message can be very much specific and the code can result very much hard-coded as to speak.

### Layout
A brief description of the layout of the project and whate you can find where.

* [iso-20022](iso-20022). In here you can find the go generated structures of some iso20022 messages. This stuff is self-consistent in that it doesn't depend on external packages beside the standard ones.
* [cmds/tpm-rtp-cli](cmds/tpm-rtp-cli). It contains the driver of the command line utilities. At the moment the generation command is under way. It doesn't do much.
* [cmds/generate](cmds/generate). It is the generation stuff that is used to produce the iso-20022 folder content. The output has been produced by a test case [cmds/generate/golang_test.go](cmds/generate/golang_test.go).

### iso-20022 folder Layout
The folder contains artefacts produced by the generator and other stuff. More precisely.

* [iso-20022/schemas](iso-20022/schemas). Contains the xsd used for the generation.
* [iso-20022/messages](iso-20022/messages). In here a number of folders.
    - One folder for each message named after the message
    - [iso-20022/messages/common](iso-20022/messages/common) a common folder that contains stuff related to types that are shared among messages.
    - [iso-20022/messages/xsdt](iso-20022/messages/xsdt) a folder for definitions of go types related to schema builtin types.
* [iso-20022/examples](iso-20022/examples). A few test cases hand coded to give the idea of the structuring in some real case.

#### iso-20022 message folder
The folders are named after the message in a go-frienly way (i.e. `pain.013.001.07` gets to `pain_013_001_07`). 
Each folder contains a number of files. The size and the complexity of the messages and the size of the generated stuff has forced to distribute code on various files.

* document.go, document-paths.go, document-complex-set-ops.go, document-simple-set-ops.go: they contain the document structure definition and the document level commands. They are considered below in a dedicated paragraph.
  (note: in the repo not every file for every message have been committed due to size and amangeability).
* complex-types.go: contains the declaration of structures not shared and specific to the message (at tleast in the context of the messages considered)
* complex-types-ops.go: contains the methods of the structures.
* <message-id>_test.go (where message-id stands for `pain_013_001_07` etc..): this file contains a literal sample of the document with sample values, it is meant to provide an overview of the message and to 
extract snippet that might be used somewhere else.

##### Typical layout of a structure
The snippet below provide an example of the type of stuff that is in there (Methods and structure might change over time but pretty much this is the idea). In the snippet below it is worth of note:

* Optional properties of complex types are pointers to structs.
* Mandatory structs are normal structs.
* Plurals are handled as array of items (not pointers to).

```
// CreditTransferTransaction35 type definition
type CreditTransferTransaction35 struct {
	PmtId           PaymentIdentification6                               `xml:"PmtId"`
	PmtTpInf        *common.PaymentTypeInformation26                     `xml:"PmtTpInf,omitempty"`
	PmtCond         *common.PaymentCondition1                            `xml:"PmtCond,omitempty"`
	Amt             common.AmountType4Choice                             `xml:"Amt"`
	ChrgBr          common.ChargeBearerType1Code                         `xml:"ChrgBr"`
	ChqInstr        *Cheque11                                            `xml:"ChqInstr,omitempty"`
	UltmtDbtr       *common.PartyIdentification135                       `xml:"UltmtDbtr,omitempty"`
	IntrmyAgt1      *common.BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty"`
	IntrmyAgt2      *common.BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty"`
	IntrmyAgt3      *common.BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty"`
	CdtrAgt         common.BranchAndFinancialInstitutionIdentification6  `xml:"CdtrAgt"`
	Cdtr            common.PartyIdentification135                        `xml:"Cdtr"`
	CdtrAcct        *common.CashAccount38                                `xml:"CdtrAcct,omitempty"`
	UltmtCdtr       *common.PartyIdentification135                       `xml:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt []InstructionForCreditorAgent1                       `xml:"InstrForCdtrAgt,omitempty"`
	Purp            *common.Purpose2Choice                               `xml:"Purp,omitempty"`
	RgltryRptg      []RegulatoryReporting3                               `xml:"RgltryRptg,omitempty"`
	Tax             *TaxInformation8                                     `xml:"Tax,omitempty"`
	RltdRmtInf      []RemittanceLocation7                                `xml:"RltdRmtInf,omitempty"`
	RmtInf          *common.RemittanceInformation16                      `xml:"RmtInf,omitempty"`
	NclsdFile       []common.Document12                                  `xml:"NclsdFile,omitempty"`
	SplmtryData     []common.SupplementaryData1                          `xml:"SplmtryData,omitempty"`
}

// IsValid checks if CreditTransferTransaction35 is valid
func (s CreditTransferTransaction35) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.PmtId.IsValid(false)
	valid = valid && (s.PmtTpInf == nil || (s.PmtTpInf != nil && s.PmtTpInf.IsValid(true)))

	valid = valid && (s.PmtCond == nil || (s.PmtCond != nil && s.PmtCond.IsValid(true)))

	valid = valid && s.Amt.IsValid(false)
	valid = valid && s.ChrgBr.IsValid(false)
	valid = valid && (s.ChqInstr == nil || (s.ChqInstr != nil && s.ChqInstr.IsValid(true)))

	valid = valid && (s.UltmtDbtr == nil || (s.UltmtDbtr != nil && s.UltmtDbtr.IsValid(true)))

	valid = valid && (s.IntrmyAgt1 == nil || (s.IntrmyAgt1 != nil && s.IntrmyAgt1.IsValid(true)))

	valid = valid && (s.IntrmyAgt2 == nil || (s.IntrmyAgt2 != nil && s.IntrmyAgt2.IsValid(true)))

	valid = valid && (s.IntrmyAgt3 == nil || (s.IntrmyAgt3 != nil && s.IntrmyAgt3.IsValid(true)))

	valid = valid && s.CdtrAgt.IsValid(false)
	valid = valid && s.Cdtr.IsValid(false)
	valid = valid && (s.CdtrAcct == nil || (s.CdtrAcct != nil && s.CdtrAcct.IsValid(true)))

	valid = valid && (s.UltmtCdtr == nil || (s.UltmtCdtr != nil && s.UltmtCdtr.IsValid(true)))

	for j := 0; j < len(s.InstrForCdtrAgt); j++ {
		valid = valid && s.InstrForCdtrAgt[j].IsValid(true)
	}

	valid = valid && (s.Purp == nil || (s.Purp != nil && s.Purp.IsValid(true)))

	for j := 0; j < len(s.RgltryRptg); j++ {
		valid = valid && s.RgltryRptg[j].IsValid(true)
	}

	valid = valid && (s.Tax == nil || (s.Tax != nil && s.Tax.IsValid(true)))

	for j := 0; j < len(s.RltdRmtInf); j++ {
		valid = valid && s.RltdRmtInf[j].IsValid(true)
	}

	valid = valid && (s.RmtInf == nil || (s.RmtInf != nil && s.RmtInf.IsValid(true)))

	for j := 0; j < len(s.NclsdFile); j++ {
		valid = valid && s.NclsdFile[j].IsValid(true)
	}

	for j := 0; j < len(s.SplmtryData); j++ {
		valid = valid && s.SplmtryData[j].IsValid(true)
	}

	return valid
}
```

##### document files
In here a brief explanation of what is in the document* files that are found in the message folder.

* document.go: it is the structure of the document and a few methods. Previously put in the `complex-types.go` has been extracted since the Document is the root of the thing and is easier to find.
* document-paths.go: it is a set of constants with each one constant to each different path from the root to each node (not just the leaves).
* document-complex-set-ops.go, document-simple-set-ops.go: they contain a set function for setting each individual node of the document. These functions are methods of the Document struct and are named
after the path of each node from the root of the document. They are two files because of size: the complex contains function where the parameter is of struct or array type whereas the simple contain
functions that address only the leaves of the tree.

These are few consts from the pain.013.001.07 stuff. What are them used for? Nothing at the moment but of course there is a plan.

```
	Path_Doc                                    = "Doc"
	Path_Doc_CdtrPmtActvtnReq                   = "Doc.CdtrPmtActvtnReq"
	Path_Doc_CdtrPmtActvtnReq_GrpHdr            = "Doc.CdtrPmtActvtnReq.GrpHdr"
	Path_Doc_CdtrPmtActvtnReq_GrpHdr_MsgId      = "Doc.CdtrPmtActvtnReq.GrpHdr.MsgId"
	Path_Doc_CdtrPmtActvtnReq_GrpHdr_CreDtTm    = "Doc.CdtrPmtActvtnReq.GrpHdr.CreDtTm"
	Path_Doc_CdtrPmtActvtnReq_GrpHdr_NbOfTxs    = "Doc.CdtrPmtActvtnReq.GrpHdr.NbOfTxs"
	Path_Doc_CdtrPmtActvtnReq_GrpHdr_CtrlSum    = "Doc.CdtrPmtActvtnReq.GrpHdr.CtrlSum"
	Path_Doc_CdtrPmtActvtnReq_GrpHdr_InitgPty   = "Doc.CdtrPmtActvtnReq.GrpHdr.InitgPty"
```

Examples of functions that can be found in the document-complex-set-ops.go and document-simple-set-ops.go files. As you can note, the code takes care of handling
pointers and arrays. We might want to set that specific node down the hierarchy without worrying of checks of `nil` arrays and `nil` pointers.
If we didn't have pointers and arrays the stuff was a lot simpler.

```
func (d *Document) With_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm(_schmeNm *common.OrganisationIdentificationSchemeName1Choice) {
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id = &common.Party38Choice{}
	}
	if d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId == nil {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId = &common.OrganisationIdentification29{}
	}
	if len(d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr) == 0 {
		d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr = append(d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr, common.GenericOrganisationIdentification1{})
	}
	d.CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr[0].SchmeNm = _schmeNm
}
```

#### iso-20022 common folder
The folder contains declaration of simple types and complex types shared between the messages. Files have been split in declarations and methods.

* complex-types.go: contains the declaration of shared complex types.
* complex-types-ops.go: contains the methods of shared complex types.
* simple-types.go: contains the declaration of shared simple types.
* complex-types-ops.go: contains the methods of shared simple types.
* restriction-util.go: a few functions referenced by other files and with bridge scope for some external custom utility projects. The idea is to limit external dependencies 
to be easily managed.

Below a snippet of the generation of a simple string type with enum restriction (Methods and structure might change over time but pretty much this is the idea). 

```
// ChargeBearerType1Code May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code xsdt.String

/*
 * ChargeBearerType1Code Ops
 */

const (
	ChargeBearerType1CodeDEBT = "DEBT"
	ChargeBearerType1CodeCRED = "CRED"
	ChargeBearerType1CodeSHAR = "SHAR"
	ChargeBearerType1CodeSLEV = "SLEV"
)

var ChargeBearerType1CodeEnumRestriction = []string{
     ChargeBearerType1CodeDEBT, ChargeBearerType1CodeCRED, ChargeBearerType1CodeSHAR, ChargeBearerType1CodeSLEV}

// IsValid checks if ChargeBearerType1Code of type String is valid
func (t ChargeBearerType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), ChargeBearerType1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ChargeBearerType1Code) String() string {
	return string(t)
}

// ToChargeBearerType1Code method for easy conversion with application of restrictions
func ToChargeBearerType1Code(s string) (ChargeBearerType1Code, error) {
	if !isEnumRestrictionValid(s, ChargeBearerType1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type ChargeBearerType1Code", s)
	}

	return ChargeBearerType1Code(s), nil
}

// MustToChargeBearerType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToChargeBearerType1Code(s string) ChargeBearerType1Code {
	v, err := ToChargeBearerType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

// ChargeBearerType1CodeExample method for generation of valid sample data
func ChargeBearerType1CodeExample() ChargeBearerType1Code {
	return ChargeBearerType1Code(generateSampleDataWithEnumRestriction(ChargeBearerType1CodeEnumRestriction))

}
```

#### iso-20022 examples
In the [iso-20022/examples](iso-20022/examples) are contained a few examples to show the use of structures in literals.
The use of the `MustTo` type of function is to allow the literal to be layed-out without dealing with error (The `MustTo`)
version panics in case of problems.

```
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
								Value: xsdt.ToDecimal(535.35),
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
```