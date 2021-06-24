package coretests

import "strings"

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

func IsStringMessageWithoutWhitespaceSortedEqual(
	isPrintOnFail bool,
	hasWhitespace bool,
	when interface{},
	actual, expected string,
	counter int,
) bool {
	if hasWhitespace {
		return isStringMessageWithoutWhitespaceSortedEqual(
			isPrintOnFail,
			when,
			actual,
			expected,
			counter)
	}

	trimActual := strings.TrimSpace(actual)
	trimExpected := strings.TrimSpace(actual)
	isEqual := trimActual == trimExpected

	if !isEqual && isPrintOnFail {
		PrintFailedExpectedSortedMessage(
			!isEqual,
			when,
			actual,
			trimActual,
			expected,
			trimExpected,
			counter)
	}

	return isEqual
}

func isStringMessageWithoutWhitespaceSortedEqual(
	isPrintOnFail bool,
	when interface{},
	actual, expected string,
	counter int,
) bool {
	actualSortedDefault := GetMessageToSorted(
		false,
		strings.TrimSpace(actual),
		commonJoiner)

	expectedSortedDefault := GetMessageToSorted(
		false,
		strings.TrimSpace(expected),
		commonJoiner)

	isEqual := actualSortedDefault == expectedSortedDefault
	isFailed := !isEqual

	if isPrintOnFail && isFailed {
		PrintFailedExpectedSortedMessage(
			isFailed,
			when,
			actual,
			actualSortedDefault,
			expected,
			expectedSortedDefault,
			counter)
	}

	return isEqual
}
