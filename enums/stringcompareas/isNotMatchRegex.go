package stringcompareas

import "gitlab.com/evatix-go/core/regexnew"

// NotMatchRegex no use of isCaseSensitive
//
// Tided with NotMatchRegex, invert of isRegexFunc
//
// isCaseSensitive is kept for consistency and calling ability
var isNotMatchRegex = func(
	contentLine,
	regexStringSearching string,
	isCaseSensitive bool,
) bool {
	return !regexnew.NewMust(regexStringSearching).
		MatchString(contentLine)
}
