package coretests

import "gitlab.com/evatix-go/core/constants"

const (
	// Contains name-value using %+v, %v for only value.
	//
	// Expectations : %+v
	// Actual: %+v
	logFormat = "\n ====================================" +
		"Actual vs IsMatchesExpectation " +
		"====================================\n" +
		"\tExpectations : %+v\n" +
		"\tActual: %+v"
	printValuesFormat = "\nHeader:%s\n" +
		"\tType:%T\n" +
		"\tValue:%s\n"

	quickActualExpectedMessageFormat = "----------------------\n" +
		"%d )\tWhen:%#v\n\t\t" +
		"  Actual:`%#v` ,\n\t\t" +
		"Expected:`%#v`"

	printValuesForActualAndSortedFormat = "----------------------" +
		"\n%d )\t" +
		"    When:%#v\n\t\t" +
		"  Actual:`%#v` ,\n\t\t" +
		"Expected:`%#v`\n\t\t" +
		"  Actual-Processed:`%#v` ,\n\t\t" +
		"Expected-Processed:`%#v`"

	commonJoiner = constants.Pipe
)
