package stringcompareas

import "github.com/alimtvnetwork/core/coreutils/stringutil"

var isNotStartsWithFunc = func(
	contentLine,
	searchComparingLine string,
	isIgnoreCase bool,
) bool {
	return !stringutil.IsStartsWith(
		contentLine,
		searchComparingLine,
		isIgnoreCase)
}
