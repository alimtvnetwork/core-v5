package stringcompareas

// isNotGlobFunc is the inversion of isGlobFunc.
var isNotGlobFunc = func(
	contentLine,
	globPattern string,
	isIgnoreCase bool,
) bool {
	return !isGlobFunc(
		contentLine,
		globPattern,
		isIgnoreCase,
	)
}
