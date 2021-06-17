package stringslice

// LinesSimpleProcess split text using constants.NewLineUnix
// then returns lines processed by lineProcessor
func LinesSimpleProcess(
	splitsLines []string,
	lineProcessor func(lineIn string) (lineOut string),
) []string {
	length := len(splitsLines)
	slice := Make(length, length)

	for i, lineIn := range splitsLines {
		lineOut := lineProcessor(lineIn)

		slice[i] = lineOut
	}

	return slice
}
