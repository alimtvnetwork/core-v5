package chmodhelper

func VerifyChmodLocationsUsingPartialRwxDirect(
	isContinueOnError,
	isSkipOnInvalid bool,
	partialRwx string,
	locations ...string,
) error {
	return VerifyChmodLocationsUsingPartialRwx(
		isContinueOnError,
		isSkipOnInvalid,
		partialRwx,
		locations)
}
