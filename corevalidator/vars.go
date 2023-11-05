package corevalidator

import "gitlab.com/auk-go/core/enums/stringcompareas"

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

	DefaultSortTrimCoreCondition = ValidatorCoreCondition{
		IsTrimCompare:        true,
		IsNonEmptyWhitespace: true,
		IsSortStringsBySpace: true,
	}

	DefaultUniqueWordsCoreCondition = ValidatorCoreCondition{
		IsTrimCompare:        true,
		IsUniqueWordOnly:     true,
		IsNonEmptyWhitespace: true,
		IsSortStringsBySpace: true,
	}

	EmptyValidator = TextValidator{
		Search:                 "",
		SearchAs:               stringcompareas.Equal,
		ValidatorCoreCondition: DefaultTrimCoreCondition,
	}
)
