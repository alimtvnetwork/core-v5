package stringslice

import (
	"gitlab.com/evatix-go/core/constants"
)

// LinesProcess split text using constants.NewLineUnix
// then returns lines processed by lineProcessor
func LinesProcess(
	splitsLines []string,
	lineProcessor func(index int, lineIn string) (lineOut string, isTake, isBreak bool),
) []string {
	slice := Make(constants.Zero, len(splitsLines))

	for i, lineIn := range splitsLines {
		lineOut, isTake, isBreak := lineProcessor(i, lineIn)

		if isTake {
			slice = append(slice, lineOut)
		}

		if isBreak {
			break
		}
	}

	return slice
}
