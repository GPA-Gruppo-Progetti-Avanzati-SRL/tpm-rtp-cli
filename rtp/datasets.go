package rtp

const (
	DataSetDS01 = "DS01"
	DataSetDS02 = "DS02"
	DataSetDS03 = "DS03"
	DataSetDS04 = "DS04"
	DataSetDS05 = "DS05"
	DataSetDS06 = "DS06"
	DataSetDS07 = "DS07"
	DataSetDS08 = "DS08"
	DataSetDS09 = "DS09"
	DataSetDS10 = "DS10"
	DataSetDS11 = "DS11"
	DataSetDS12 = "DS12"
	DataSetDS13 = "DS13"
	DataSetDS14 = "DS14"
	DataSetDS15 = "DS15"
	DataSetDS16 = "DS16"
	DataSetDS17 = "DS17"

	// IG 2.0 Pag. 64 (DS-04a), 109 (DS-04b+), 159 (DS-05, DS-06), 203 (DS-08, DS-09 positive), 257 (DS-08, DS-09 negative), 495 (DS-17)

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

	return "UKN:" + cd
}
