package coretests

type TestCaseMessenger interface {
	GetFuncName() string
	Value() any
	Expected() any
	Actual() any
}
