package stringcompareas

import "github.com/alimtvnetwork/core/coreutils/stringutil"

var isStartsWithFunc = func(
	contentLine,
	searchComparingLine string,
	isIgnoreCase bool,
) bool {
	return stringutil.IsStartsWith(
		contentLine,
		searchComparingLine,
		isIgnoreCase)
}
