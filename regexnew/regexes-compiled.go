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
)
