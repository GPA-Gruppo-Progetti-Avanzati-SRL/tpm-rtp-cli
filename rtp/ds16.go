package rtp

import "github.com/rs/zerolog/log"

var listOfDS16ValidTransactionStatus = []string{
	"",
}

var listOfDS16ValidTransactionReasons = []string{
	StatusReasonAlreadyExpiredRTP,
	StatusReasonAlreadyAcceptedRTP,
	StatusReasonAlreadyRefusedRTP,
	StatusReasonAlreadyRejectedRTP,
	StatusReasonInitialRTPNeverReceived,
	StatusReasonRTPReceivedCanBeProcessed,
}

func MapDS16ReasonsToRtpStatus(rsn string) Status {
	var st Status
	switch rsn {
	case StatusReasonAlreadyExpiredRTP:
		st.Code = StatusRejected
		st.ReasonCode = rsn

	case StatusReasonAlreadyAcceptedRTP:
		st.Code = StatusAccepted

	case StatusReasonAlreadyRefusedRTP:
		st.Code = StatusRejected
		st.ReasonCode = rsn
	case StatusReasonAlreadyRejectedRTP:
		st.Code = StatusRejected
		st.ReasonCode = rsn
	case StatusReasonInitialRTPNeverReceived:
		st.Code = StatusRejected
		st.ReasonCode = rsn
	case StatusReasonRTPReceivedCanBeProcessed:
		st.Code = StatusReasonRTPReceivedCanBeProcessed
	}

	st.UpdateText(DataSetDS16)
	return st
}

func Request4StatusUpdateNextStatus(dataset string, current Status, rsn string) (Status, bool) {

	if current.Code != StatusZero {
		return current, false
	}

	rc := true
	var st Status
	switch dataset {
	case DataSetDS16:

		switch rsn {
		case StatusReasonAlreadyExpiredRTP:
			st.Code = StatusRejected
			st.ReasonCode = rsn
		case StatusReasonAlreadyAcceptedRTP:
			st.Code = StatusAccepted

		case StatusReasonAlreadyRefusedRTP:
			st.Code = StatusRejected
			st.ReasonCode = rsn

		case StatusReasonAlreadyRejectedRTP:
			st.Code = StatusRejected
			st.ReasonCode = rsn

		case StatusReasonInitialRTPNeverReceived:
			st.Code = StatusRejected
			st.ReasonCode = rsn

		case StatusReasonRTPReceivedCanBeProcessed:
			st.Code = StatusReasonRTPReceivedCanBeProcessed
		}

	default:
		log.Error().Str("dataset", dataset).Str("current", current.Code).Str("event", rsn).Msg("cannot compute wf-status on unrecognized dataset")
		rc = false
	}

	if rc {
		_ = st.UpdateText(dataset)
	}

	return st, rc
}
