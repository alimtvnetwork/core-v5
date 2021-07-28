package reqtype

import "gitlab.com/evatix-go/core/msgtype"

func RangesOnlySupportedFor(
	message string,
	requests ...Request,
) error {
	if len(requests) == 0 {
		return nil
	}

	referencesMessage := RangesStringDefaultJoiner(
		requests...)

	return msgtype.RangesOnlySupported.Error(
		message,
		referencesMessage)
}
