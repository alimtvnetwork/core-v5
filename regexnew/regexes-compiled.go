package regexnew

import (
	"gitlab.com/evatix-go/core/regconsts"
)

var (
	WhitespaceFinderRegex             = NewMust(regconsts.AllWhitespaces)
	HashCommentWithSpaceOptionalRegex = NewMust(regconsts.HashCommentWithSpaceOptional)
	WhitespaceOrPipeFinderRegex       = NewMust(regconsts.AllWhitespacesOrPipe)
	DollarIdentifierRegex             = NewMust(regconsts.EachWordsWithDollarSymbolDefinition)
	PercentIdentifierRegex            = NewMust(regconsts.EachWordsWithinPercentSymbolDefinition)
	PrettyNameRegex                   = NewMust(regconsts.PrettyName)
	ExactIdFieldMatchingRegex         = NewMust(regconsts.ExactIdFieldMatching)
	ExactVersionIdFieldMatchingRegex  = NewMust(regconsts.ExactVersionIdFieldMatching)
	UbuntuNameCheckerRegex            = NewMust(regconsts.UbuntuNameChecker)
	CentOsNameCheckerRegex            = NewMust(regconsts.CentOsNameChecker)
	RedHatNameCheckerRegex            = NewMust(regconsts.RedHatNameChecker)
	FirstNumberAnyWhereCheckerRegex   = NewMust(regconsts.FirstNumberAnyWhere)
	WindowsVersionNumberCheckerRegex  = FirstNumberAnyWhereCheckerRegex
)
