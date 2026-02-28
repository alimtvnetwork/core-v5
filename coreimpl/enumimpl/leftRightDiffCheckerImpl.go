package enumimpl

import "fmt"

type leftRightDiffCheckerImpl struct{}

func (it leftRightDiffCheckerImpl) GetSingleDiffResult(isLeft bool, l, r any) any {
	diff := DiffLeftRight{
		Left:  l,
		Right: r,
	}

	return diff.DiffString()
}

func (it leftRightDiffCheckerImpl) GetResultOnKeyMissingInRightExistInLeft(lKey string, lVal any) any {
	return fmt.Sprintf("%+v (type:%T) - left - key is missing!", lVal, lVal)
}

func (it leftRightDiffCheckerImpl) IsEqual(isRegardless bool, l, r any) bool {
	return DefaultDiffCheckerImpl.IsEqual(isRegardless, l, r)
}

func (it leftRightDiffCheckerImpl) AsChecker() DifferChecker {
	return &it
}
