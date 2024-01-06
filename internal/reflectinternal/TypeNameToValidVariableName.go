package reflectinternal

import (
	"strings"
	"unicode"

	"gitlab.com/auk-go/core/coredata/stringslice"
)

func TypeNameToValidVariableName(currentTypeName string) (validVarName string) {
	if len(currentTypeName) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.Grow(len(currentTypeName))
	splitsByDot := strings.Split(currentTypeName, ".")
	hasDot := len(splitsByDot) > 1

	if hasDot {
		var leftPartSb strings.Builder
		firstPart := splitsByDot[0]

		for i, c := range firstPart {
			if unicode.IsSpace(c) {
				continue
			}

			if c == '[' && firstPart[i+1] == ']' {
				i++
				leftPartSb.WriteString(typeReplacerMap["[]"])
				continue
			}

			_, has := typeReplacerMap[string(c)]
			if has {
				leftPartSb.WriteString(typeReplacerMap[string(c)])

				continue
			}
		}

		currentTypeName = stringslice.Joins(
			"",
			leftPartSb.String(),
			splitsByDot[len(splitsByDot)-1],
		)
	}

	for i := 0; i < len(currentTypeName); i++ {
		currentChar := currentTypeName[i]

		if unicode.IsSpace(rune(currentChar)) {
			continue
		}

		if currentChar == '[' && currentTypeName[i+1] == ']' {
			i++
			sb.WriteString(typeReplacerMap["[]"])

			continue
		}

		replace, has := typeReplacerMap[string(currentChar)]

		if has {
			sb.WriteString(typeReplacerMap[replace])

			continue
		}

		sb.WriteByte(currentChar)
	}

	return sb.String()
}
