package coretests

type SimpleTestCaseWrapper interface {
	CaseTitle() string
	Input() any
	Expected() any
	Actual() any
	SetActual(actual any)
}
