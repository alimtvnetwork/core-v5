package errcore

import "gitlab.com/auk-go/core/internal/reflectinternal"

func PathMeaningfulError(
	rawErrType RawErrorType,
	err error,
	location string,
) error {
	if err == nil {
		return nil
	}

	funcName := reflectinternal.
		GetFunc.
		NameOnlyByStack(1)

	errMsg := err.Error() +
		", location: [" + location + "]"

	return rawErrType.Error(
		funcName,
		errMsg,
	)
}
