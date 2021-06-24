package testconsts

import "regexp"

var (
	WhitespaceFinderRegex       = regexp.MustCompile("\\s+")
	WhitespaceOrPipeFinderRegex = regexp.MustCompile("\\s+|\\|+")
)
