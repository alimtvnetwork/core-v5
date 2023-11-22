package coretests

func GetAssertMessage(testCaseMessenger TestCaseMessenger, counter int) string {
	return GetAssert.QuickGherkins(
		testCaseMessenger.Value(),
		testCaseMessenger.Actual(),
		testCaseMessenger.Expected(),
		counter,
	)
}
