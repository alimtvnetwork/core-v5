package corevalidator

type ValidatorCoreCondition struct {
	IsTrimCompare        bool
	IsUniqueWordOnly     bool
	IsNonEmptyWhitespace bool // Split by whitespace and then compare, don't keep whitespace into comparison
	IsSortStringsBySpace bool
}

func (it *ValidatorCoreCondition) IsSplitByWhitespace() bool {
	return it.IsUniqueWordOnly ||
		it.IsNonEmptyWhitespace ||
		it.IsSortStringsBySpace
}
