package enumimpl

type (
	LooperFunc                func(index int, name string, anyVal any) (isBreak bool)
	LooperIntegerFunc         func(index int, name string, anyVal int) (isBreak bool)
	IsEqualCheckerFunc        func(isRegardless bool, l, r any) bool
	GetSingleDiffResultFunc   func(isLeft bool, l, r any) any
	GetOnKeyMissingResultFunc func(lKey string, lVal any) any
)
