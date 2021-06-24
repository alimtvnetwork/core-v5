package chmodhelper

func VerifyChmodLocationsUsingPartialRwx(
	isContinueOnError,
	isSkipOnInvalid bool,
	partialRwx string,
	locations []string,
) error {
	varWrapper, err := NewRwxVariableWrapper(partialRwx)

	if err != nil {
		return err
	}

	status := varWrapper.RwxMatchingStatus(
		isContinueOnError,
		isSkipOnInvalid,
		locations)

	return status.CreateErrFinalError()
}
