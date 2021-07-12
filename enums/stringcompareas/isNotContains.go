package stringcompareas

import "strings"

// isNotContainsFunc tided with NotContains
var isNotContainsFunc = func(
	contentLine,
	searchComparingLine string,
	isCaseSensitive bool,
) bool {
	if isCaseSensitive {
		return !strings.Contains(
			contentLine,
			searchComparingLine,
		)
	}

	return !strings.Contains(
		strings.ToLower(contentLine),
		strings.ToLower(searchComparingLine),
	)
}
