package rtp

import "github.com/rs/zerolog/log"

var listOfDS12ValidTransactionStatus = []string{
	StatusCancelledAsPerRequest, StatusRejectedCancellationRequest,
}

var listOfDS12ValidNegativeTransactionReasons = []string{
	RejectCancellationReasonAlreadyCancelledRTP,
	StatusReasonAlreadyExpiredRTP,
	StatusReasonAlreadyRefusedRTP,
	StatusReasonAlreadyRejectedRTP,
	RejectCancellationReasonPaymentAlreadyTransmittedExecution,
	RejectReasonRegulatoryReason,
	RejectCancellationReasonUnknownRTP,
}

func RequestForCancellationNextStatus(dataset string, current Status, event Status) (Status, bool) {

	if current.Code != StatusZero {
		return current, false
	}

	rc := true
	var st = current
	switch dataset {

	case DataSetDS12:
		st = event

	default:
		log.Error().Str("dataset", dataset).Str("current", current.Code).Str("event", event.Code).Msg("cannot compute wf-status on unrecognized dataset")
		rc = false
	}

	if rc {
		_ = st.UpdateText(dataset)
	}

	return st, rc
}
