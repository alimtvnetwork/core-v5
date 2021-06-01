package coretests

const (
	// Contains name-value using %+v, %v for only value.
	//
	// Expectations : %+v
	// Actual: %+v
	logFormat = "\n ====================================Actual vs Expectation ====================================\n" +
		"\tExpectations : %+v\n" +
		"\tActual: %+v"
	printValuesFormat = "\nHeader:%s\n" +
		"\tType:%T\n" +
		"\tValue:%s\n"
)
