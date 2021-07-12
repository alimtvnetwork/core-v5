package regexnew

import "regexp"

// NewLockIf calls New with mutex lock and unlock if true.
func NewLockIf(isLock bool, regularExpressionSyntax string) (*regexp.Regexp, error) {
	if isLock {
		regexMutex.Lock()

		defer regexMutex.Unlock()
	}

	return New(regularExpressionSyntax)
}
