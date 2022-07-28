package rtp

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/camt_029_001_09"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/camt_055_001_08"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/pain_013_001_07"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/pain_014_001_07"
	"github.com/rs/zerolog/log"
)

type Status struct {
	Code          string `yaml:"cd" mapstructure:"cd" json:"cd"`
	ReasonCode    string `yaml:"rsn,omitempty" mapstructure:"rsn,omitempty" json:"rsn,omitempty"`
	Text          string `yaml:"text,omitempty" mapstructure:"text,omitempty" json:"text,omitempty"`
	RfCInProgress bool   `yaml:"rfc-in-progress,omitempty" mapstructure:"rfc-in-progress,omitempty" json:"rfc-in-progress,omitempty"`
}

func (st *Status) IsFinal() bool {
	return IsStatusFinal(st.Code)
}

func (st *Status) UpdateText(dataset string) string {
	st.Text = LookupStsCdRsnText(dataset, st.Code, st.ReasonCode)
	return st.Text
}

func (st *Status) ShowInfo() {
	log.Trace().Str("tx-sts", st.Code).Str("rsn", st.ReasonCode).Str("text", st.Text).Msg("status info")
}

const (
	DataSetDS01 = "DS-01"
	DataSetDS02 = "DS-02"
	DataSetDS03 = "DS-03"
	DataSetDS04 = "DS-04"
	DataSetDS05 = "DS-05"
	DataSetDS06 = "DS-06"
	DataSetDS07 = "DS-07"
	DataSetDS08 = "DS-08"
	DataSetDS09 = "DS-09"
	DataSetDS10 = "DS-10"
	DataSetDS11 = "DS-11"
	DataSetDS12 = "DS-12"
	DataSetDS13 = "DS-13"
	DataSetDS14 = "DS-14"
	DataSetDS15 = "DS-15"
	DataSetDS16 = "DS-16"
	DataSetDS17 = "DS-17"

	// IG 2.0 Pag. 64 (DS-04a), 109 (DS-04b+), 159 (DS-05, DS-06), 203 (DS-08, DS-09 positive), 257 (DS-08, DS-09 negative), 495 (DS-17)

	StatusZero            = "RCRE" // Non ISO. Represent a non status.
	RTPRejectInvalidState = "ISTS" // Non Iso20022 Represent an invalid transition.

	StatusRejected            = "RJCT" // applies to: DS04, DS08, DS09 (negative)
	StatusTechnicallyAccepted = "ACTC" // applies to: DS05, DS06
	StatusAccepted            = "ACCP" // applies to: DS08, DS09 (positive)
	StatusAcceptedWithChange  = "ACWC" // applies to: DS08, DS09 (positive)

	StatusReasonAlreadyExpiredRTP         = "AEXR" // applies to: DS17 StatusUpdate pain.014
	StatusReasonAlreadyAcceptedRTP        = "ALAC" // applies to: DS10b, DS13, DS17 StatusUpdate pain.014
	StatusReasonAlreadyRefusedRTP         = "ARFR" // applies to: DS10b, DS13  StatusUpdate pain.014
	StatusReasonAlreadyRejectedRTP        = "ARJR" // applies to: DS10b, DS13   StatusUpdate pain.014
	StatusReasonInitialRTPNeverReceived   = "IRNR" // applies to: DS17 StatusUpdate pain.014
	StatusReasonRTPReceivedCanBeProcessed = "REPR" // applies to: DS17 StatusUpdate pain.014

	RejectReasonInvalidDebtorAccountNumber            = "AC02" // applies to: DS04
	RejectReasonNotAllowedCurrency                    = "AM03" // applies to: DS04, DS08, DS09
	RejectReasonDuplication                           = "AM05" // applies to: DS04, DS08, DS09
	RejectReasonWrongAmount                           = "AM09" // applies to: DS08, DS09
	RejectReasonAttachmentsNotSupported               = "ATNS" // applies to: DS04
	RejectReasonInvalidDebtorIdentificationCode       = "BE16" // applies to: DS04
	RejectReasonExpiryDateTooLong                     = "EDTL" // applies to: DS04
	RejectReasonExpiryDateTimeReached                 = "EDTR" // applies to: DS04
	RejectReasonInvalidFileFormat                     = "FF01" // applies to: DS04
	RejectReasonFraudulentOrigin                      = "FRAD" // applies to: DS04
	RejectReasonIncorrectExpiryDateTime               = "IEDT" // applies to: DS08, DS09
	RejectReasonNotSpecifiedReasonCustomerGenerated   = "MS02" // applies to: DS08, DS09
	RejectReasonNotSpecifiedReasonAgentGenerated      = "MS03" // applies to: DS04
	RejectReasonNonAgreedRTP                          = "NOAR" // applies to: DS08, DS09
	RejectReasonPayerOrPayerRTPSPNotReachable         = "NRCH" // applies to: DS04
	RejectReasonTypeOfPaymentInstrumentNotSupported   = "PINS" // applies to: DS04
	RejectReasonRegulatoryReason                      = "RR04" // applies to: DS04, DS10b, DS13
	RejectReasonRTPNotSupportedForDebtor              = "RTNS" // applies to: DS04
	RejectReasonRTPServiceProviderIdentifierIncorrect = "SPII" // applies to: DS04
	RejectReasonUnknownCreditor                       = "UCRD" // applies to: DS08, DS09

	// IG 2.0 Pag. 348 (DS-10b), 380 (DS-12, DS-13 positive), 411 (DS-12, DS-13 negative), 520 (DS-16, DS-17 status update)

	StatusRejectedCancellationRequest                          = "RJCR" // applies to: DS10b, DS13 (negative)
	StatusCancelledAsPerRequest                                = "CNCL" // applies to: DS12, DS13 (positive)
	RejectCancellationReasonAlreadyCancelledRTP                = "ACLR" // applies to: DS10b, DS13
	RejectCancellationReasonPaymentAlreadyTransmittedExecution = "PATE" // applies to: DS10b, DS13
	RejectCancellationReasonUnknownRTP                         = "URTP" // applies to: DS10b, DS13
	RfCStatusUpdateAlreadyRejected                             = "RCAR" // applies to: DS17
	RfCStatusUpdateNeverReceived                               = "RCNR" // applies to: DS17
	RfCStatusUpdateReceivedAndProcessed                        = "RCPR" // applies to: DS17
)

type StsRsnCodeDescription struct {
	Code        string
	Description string
}

var StsRsnCodeDescriptionRegistry = map[string]StsRsnCodeDescription{
	StatusZero: {
		Code:        StatusZero,
		Description: "Created",
	},
	RTPRejectInvalidState: {
		Code:        RTPRejectInvalidState,
		Description: "Invalid state",
	},
	StatusRejected: {
		Code:        StatusRejected,
		Description: "Rejected",
	},
	StatusTechnicallyAccepted: {
		Code:        StatusTechnicallyAccepted,
		Description: "Technically Accepted",
	},
	StatusAccepted: {
		Code:        StatusAccepted,
		Description: "Accepted",
	},
	StatusAcceptedWithChange: {
		Code:        StatusAcceptedWithChange,
		Description: "Accepted With Change",
	},
	StatusReasonAlreadyExpiredRTP: {
		Code:        StatusReasonAlreadyExpiredRTP,
		Description: "Already Expired",
	},
	StatusReasonAlreadyAcceptedRTP: {
		Code:        StatusReasonAlreadyAcceptedRTP,
		Description: "Already Accepted",
	},
	StatusReasonAlreadyRefusedRTP: {
		Code:        StatusReasonAlreadyRefusedRTP,
		Description: "Already Refused",
	},
	StatusReasonAlreadyRejectedRTP: {
		Code:        StatusReasonAlreadyRejectedRTP,
		Description: "Already Rejected",
	},
	StatusReasonInitialRTPNeverReceived: {
		Code:        StatusReasonInitialRTPNeverReceived,
		Description: "Initial RTP Never Received",
	},
	StatusReasonRTPReceivedCanBeProcessed: {
		Code:        StatusReasonRTPReceivedCanBeProcessed,
		Description: "RTP Received Can Be Processed",
	},
	RejectReasonInvalidDebtorAccountNumber: {
		Code:        RejectReasonInvalidDebtorAccountNumber,
		Description: "Invalid Debtor Account Number",
	},
	RejectReasonNotAllowedCurrency: {
		Code:        RejectReasonNotAllowedCurrency,
		Description: "Not Allowed Currency",
	},
	RejectReasonDuplication: {
		Code:        RejectReasonDuplication,
		Description: "Duplication",
	},
	RejectReasonWrongAmount: {
		Code:        RejectReasonWrongAmount,
		Description: "Wrong Amount",
	},
	RejectReasonAttachmentsNotSupported: {
		Code:        RejectReasonAttachmentsNotSupported,
		Description: "Attachments Not Supported",
	},
	RejectReasonInvalidDebtorIdentificationCode: {
		Code:        RejectReasonInvalidDebtorIdentificationCode,
		Description: "Invalid Debtor Identification Code",
	},
	RejectReasonExpiryDateTooLong: {
		Code:        RejectReasonExpiryDateTooLong,
		Description: "Expiry Date Too Long",
	},
	RejectReasonExpiryDateTimeReached: {
		Code:        RejectReasonExpiryDateTimeReached,
		Description: "Expiry Date Time Reached",
	},
	RejectReasonInvalidFileFormat: {
		Code:        RejectReasonInvalidFileFormat,
		Description: "Invalid File Format",
	},
	RejectReasonFraudulentOrigin: {
		Code:        RejectReasonFraudulentOrigin,
		Description: "Fraudulent Origin",
	},
	RejectReasonIncorrectExpiryDateTime: {
		Code:        RejectReasonIncorrectExpiryDateTime,
		Description: "Incorrect Expiry DateTime",
	},
	RejectReasonNotSpecifiedReasonCustomerGenerated: {
		Code:        RejectReasonNotSpecifiedReasonCustomerGenerated,
		Description: "Not Specified Reason Customer Generated",
	},
	RejectReasonNotSpecifiedReasonAgentGenerated: {
		Code:        RejectReasonNotSpecifiedReasonAgentGenerated,
		Description: "Reason has not been specified by agent",
	},
	RejectReasonNonAgreedRTP: {
		Code:        RejectReasonNonAgreedRTP,
		Description: "Non Agreed RTP",
	},
	RejectReasonPayerOrPayerRTPSPNotReachable: {
		Code:        RejectReasonPayerOrPayerRTPSPNotReachable,
		Description: "Payer Or Payer RTPSP Not Reachable",
	},
	RejectReasonTypeOfPaymentInstrumentNotSupported: {
		Code:        RejectReasonTypeOfPaymentInstrumentNotSupported,
		Description: "Type Of Payment Instrument Not Supported",
	},
	RejectReasonRegulatoryReason: {
		Code:        RejectReasonRegulatoryReason,
		Description: "Regulatory Reason",
	},
	RejectReasonRTPNotSupportedForDebtor: {
		Code:        RejectReasonRTPNotSupportedForDebtor,
		Description: "RTP Not Supported ForDebtor",
	},
	RejectReasonRTPServiceProviderIdentifierIncorrect: {
		Code:        RejectReasonRTPServiceProviderIdentifierIncorrect,
		Description: "RTP Service Provider Identifier Incorrect",
	},
	RejectReasonUnknownCreditor: {
		Code:        RejectReasonUnknownCreditor,
		Description: "Unknown creditor",
	},
	StatusRejectedCancellationRequest: {
		Code:        StatusRejectedCancellationRequest,
		Description: "Rejected Cancellation Request",
	},
	StatusCancelledAsPerRequest: {
		Code:        StatusCancelledAsPerRequest,
		Description: "Cancelled as per request",
	},
	RejectCancellationReasonAlreadyCancelledRTP: {
		Code:        RejectCancellationReasonAlreadyCancelledRTP,
		Description: "Request-to-pay has already been cancelled",
	},

	RejectCancellationReasonPaymentAlreadyTransmittedExecution: {
		Code:        RejectCancellationReasonPaymentAlreadyTransmittedExecution,
		Description: "Payment related to the request-to-pay has already been transmitted for execution",
	},
	RejectCancellationReasonUnknownRTP: {
		Code:        RejectCancellationReasonUnknownRTP,
		Description: "Request-to-pay is unknown",
	},
	RfCStatusUpdateAlreadyRejected: {
		Code:        RfCStatusUpdateAlreadyRejected,
		Description: "Already Rejected",
	},
	RfCStatusUpdateNeverReceived: {
		Code:        RfCStatusUpdateNeverReceived,
		Description: "Never received",
	},
	RfCStatusUpdateReceivedAndProcessed: {
		Code:        RfCStatusUpdateReceivedAndProcessed,
		Description: "Already Processed",
	},
}

func LookupStsCdRsnText(dataset string, cd string, rsn string) string {

	if rsn != "" {
		if s, ok := StsRsnCodeDescriptionRegistry[rsn]; ok {
			return s.Description
		}
	}

	if s, ok := StsRsnCodeDescriptionRegistry[cd]; ok {
		return s.Description
	}
	return ""
}

func IsTransactionStatusValidForDataset(ds string, txSts string) bool {

	var rc bool
	switch ds {
	case DataSetDS04:
		if util.StringArrayContains(listOfDS04ValidTransactionStatus, txSts) {
			rc = true
		}
	case DataSetDS05:
		if util.StringArrayContains(listOfDS05ValidTransactionStatus, txSts) {
			rc = true
		}
	case DataSetDS08:
		if util.StringArrayContains(listOfDS08ValidTransactionStatus, txSts) {
			rc = true
		}
	case DataSetDS12:
		if util.StringArrayContains(listOfDS12ValidTransactionStatus, txSts) {
			rc = true
		}
	case DataSetDS16:
		if util.StringArrayContains(listOfDS16ValidTransactionStatus, txSts) {
			rc = true
		}
	default:
		log.Error().Str("dataset", ds).Msg("unsupported dataset")
	}

	return rc
}

func IsTransactionStatusReasonValidForDataset(ds string, st Status) bool {

	var rc bool
	switch ds {
	case DataSetDS04:
		if util.StringArrayContains(listOfDS04ValidNegativeTransactionReasons, st.ReasonCode) {
			rc = true
		}
	case DataSetDS05:
		if st.ReasonCode == "" {
			rc = true
		}
	case DataSetDS08:
		if st.Code == StatusRejected {
			if util.StringArrayContains(listOfDS08ValidNegativeTransactionReasons, st.ReasonCode) {
				rc = true
			}
		} else if st.ReasonCode == "" {
			rc = true
		}
	case DataSetDS12:
		if st.Code == StatusRejectedCancellationRequest {
			if util.StringArrayContains(listOfDS12ValidNegativeTransactionReasons, st.ReasonCode) {
				rc = true
			}
		} else if st.ReasonCode == "" {
			rc = true
		}
	case DataSetDS16:
		if util.StringArrayContains(listOfDS16ValidTransactionReasons, st.ReasonCode) {
			rc = true
		}
	default:
		log.Error().Str("dataset", ds).Msg("unsupported dataset")
	}

	return rc
}

var listOfFinalTransactionStatus = []string{
	StatusRejected, StatusRejectedCancellationRequest, StatusAccepted, StatusCancelledAsPerRequest, StatusAcceptedWithChange,
}

func IsStatusFinal(txSts string) bool {
	return util.StringArrayContains(listOfFinalTransactionStatus, txSts)
}

func RequestToPayNextStatus(dataset string, current Status, event Status) (Status, bool) {

	var rc bool
	var st = current
	switch dataset {
	case DataSetDS04:
		if current.Code == StatusZero {
			rc = true
			st = event
		}
	case DataSetDS05:
		if current.Code == StatusZero {
			rc = true
			st = event
			st.RfCInProgress = current.RfCInProgress
		}
	case DataSetDS08:
		if current.Code == StatusZero || current.Code == StatusTechnicallyAccepted {
			rc = true
			st = event
		}
	case DataSetDS12:
		// This condition refers to the status of the DS10 request....
		// Different issue when there is the need to propagate to the original rtp status.
		if !current.IsFinal() {
			rc = true
			if event.Code == StatusCancelledAsPerRequest {
				st = event
			} else {
				if !st.RfCInProgress {
					log.Warn().Str("dataset", dataset).Msg("dataset on an Rtp with no RfCInProgress flag")
				}
				st.RfCInProgress = false
			}
		}
	case DataSetDS16:
		if !current.IsFinal() {
			if event.IsFinal() {
				rc = true
				st = event
			}
		}

	default:
		log.Error().Str("dataset", dataset).Str("current", current.Code).Str("event", event.Code).Msg("cannot compute wf-status on unrecognized dataset")

	}

	_ = st.UpdateText(dataset)
	return st, rc
}

type DataSetIsoMessageName struct {
	DataSet string
	MsgName []string
}

var DataSetIsoMessageNameRegistry = map[string]DataSetIsoMessageName{
	DataSetDS02: {DataSet: DataSetDS02, MsgName: []string{pain_013_001_07.Iso20022MsgName}},
	DataSetDS03: {DataSet: DataSetDS03, MsgName: []string{pain_013_001_07.Iso20022MsgName}},
	DataSetDS04: {DataSet: DataSetDS04, MsgName: []string{pain_014_001_07.Iso20022MsgName}},
	DataSetDS05: {DataSet: DataSetDS05, MsgName: []string{pain_014_001_07.Iso20022MsgName}},
	DataSetDS06: {DataSet: DataSetDS06, MsgName: []string{pain_014_001_07.Iso20022MsgName}},
	DataSetDS07: {DataSet: DataSetDS07, MsgName: []string{pain_014_001_07.Iso20022MsgName}},
	DataSetDS08: {DataSet: DataSetDS08, MsgName: []string{pain_014_001_07.Iso20022MsgName}},
	DataSetDS01: {DataSet: DataSetDS01, MsgName: []string{pain_013_001_07.Iso20022MsgName}},
	DataSetDS09: {DataSet: DataSetDS09, MsgName: []string{pain_014_001_07.Iso20022MsgName}},
	DataSetDS10: {DataSet: DataSetDS10, MsgName: []string{camt_055_001_08.Iso20022MsgName}},
	DataSetDS11: {DataSet: DataSetDS11, MsgName: []string{camt_055_001_08.Iso20022MsgName}},
	DataSetDS12: {DataSet: DataSetDS12, MsgName: []string{camt_029_001_09.Iso20022MsgName}},
	DataSetDS13: {DataSet: DataSetDS13, MsgName: []string{camt_029_001_09.Iso20022MsgName}},
	DataSetDS14: {DataSet: DataSetDS14, MsgName: []string{camt_055_001_08.Iso20022MsgName}},
	DataSetDS15: {DataSet: DataSetDS15, MsgName: []string{camt_055_001_08.Iso20022MsgName}},
	DataSetDS16: {DataSet: DataSetDS16, MsgName: []string{pain_014_001_07.Iso20022MsgName}},
	DataSetDS17: {DataSet: DataSetDS17, MsgName: []string{pain_014_001_07.Iso20022MsgName, camt_029_001_09.Iso20022MsgName}},
}
