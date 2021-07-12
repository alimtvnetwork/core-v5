package stringcompareas

import "gitlab.com/evatix-go/core/regexnew"

// isRegexFunc no use of isCaseSensitive
//
// isCaseSensitive is kept for consistency and calling ability
var isRegexFunc = func(
	contentLine,
	regexStringSearching string,
	isCaseSensitive bool,
) bool {
	return regexnew.NewMust(regexStringSearching).
		MatchString(contentLine)
}
