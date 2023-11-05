package corevalidator

type ValidatorCoreCondition struct {
	IsTrimCompare        bool
	IsUniqueWordOnly     bool
	IsNonEmptyWhitespace bool // Split by whitespace and then compare
	IsSortStringsBySpace bool
}

func (it *ValidatorCoreCondition) IsSplitByWhitespace() bool {
	return it.IsUniqueWordOnly ||
		it.IsNonEmptyWhitespace ||
		it.IsSortStringsBySpace
}
