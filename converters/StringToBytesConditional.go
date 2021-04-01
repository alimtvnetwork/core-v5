package converters

import "strings"

func StringToBytesConditional(
	stringInput,
	separator string,
	processor func(in string) (out byte, isTake, isBreak bool),
) *[]byte {
	if stringInput == "" {
		return &[]byte{}
	}

	splits := strings.Split(stringInput, separator)
	results := make([]byte, 0, len(splits))

	for _, v := range splits {
		out, isTake, isBreak := processor(v)

		if isTake {
			results = append(results, out)
		}

		if isBreak {
			break
		}
	}

	return &results
}
