package coretests

func IsStringErrorWithoutWhitespaceSortedEqual(
	isPrintOnFail bool,
	hasWhitespace bool,
	when interface{},
	actual error, expected string,
	counter int,
) bool {
	var actualErrorMessage string

	if actual != nil {
		actualErrorMessage = actual.Error()
	}

	if expected == "" && actualErrorMessage == "" {
		return true
	}

	return IsStringMessageWithoutWhitespaceSortedEqual(
		isPrintOnFail,
		hasWhitespace,
		when,
		actualErrorMessage,
		expected,
		counter)
}
