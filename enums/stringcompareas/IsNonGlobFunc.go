package stringcompareas

// IsNonGlobFunc is the inversion of isGlobFunc.
var IsNonGlobFunc = func(
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
