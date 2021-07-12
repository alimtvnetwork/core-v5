package corevalidator

type ValidatorCoreCondition struct {
	IsTrimCompare        bool
	IsUniqueWordOnly     bool
	IsNonEmptyWhitespace bool
	IsSortStringsBySpace bool
}

func (it *ValidatorCoreCondition) IsSplitByWhitespace() bool {
	return it.IsUniqueWordOnly ||
		it.IsNonEmptyWhitespace ||
		it.IsSortStringsBySpace
}
