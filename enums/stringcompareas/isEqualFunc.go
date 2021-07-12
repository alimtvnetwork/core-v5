package stringcompareas

import "strings"

var isEqualFunc = func(
	contentLine,
	searchComparingLine string,
	isCaseSensitive bool,
) bool {
	if isCaseSensitive {
		return contentLine ==
			searchComparingLine
	}

	return strings.EqualFold(
		searchComparingLine,
		contentLine)
}
