package coretests

func GetAssertMessage(testCaseMessenger TestCaseMessenger, counter int) string {
	return GetAssert.Quick(
		testCaseMessenger.Value(),
		testCaseMessenger.Actual(),
		testCaseMessenger.Expected(),
		counter,
	)
}
