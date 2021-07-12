package regexnew

import "regexp"

func NewMustLockIf(
	isLock bool,
	regularExpressionSyntax string,
) *regexp.Regexp {
	if isLock {
		regexMutex.Lock()

		defer regexMutex.Unlock()
	}

	return NewMust(regularExpressionSyntax)
}
