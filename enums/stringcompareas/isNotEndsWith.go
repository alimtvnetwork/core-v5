package stringcompareas

import "github.com/alimtvnetwork/core/coreutils/stringutil"

// isNotEndsWithFunc tided with NotEndsWith
var isNotEndsWithFunc = func(
	contentLine,
	searchComparingLine string,
	isIgnoreCase bool,
) bool {
	return !stringutil.IsEndsWith(
		contentLine,
		searchComparingLine,
		isIgnoreCase)
}
