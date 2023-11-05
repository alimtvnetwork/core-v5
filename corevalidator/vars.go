package corevalidator

var (
	DefaultDisabledCoreCondition = ValidatorCoreCondition{
		IsTrimCompare:        false,
		IsUniqueWordOnly:     false,
		IsNonEmptyWhitespace: false,
		IsSortStringsBySpace: false,
	}

	DefaultTrimCoreCondition = ValidatorCoreCondition{
		IsTrimCompare: true,
	}
)
