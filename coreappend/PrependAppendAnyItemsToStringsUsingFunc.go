package coreappend

func PrependAppendAnyItemsToStringsUsingFunc(
	isSkipEmptyString bool,
	compilerFunc func(item any) string,
	prependItem, appendItem any,
	anyItems ...any,
) []string {
	slice := make([]string, 0, len(anyItems)+3)

	prependString := compilerFunc(prependItem)

	if isSkipEmptyString && prependString != "" {
		slice = append(
			slice,
			prependString)
	} else if !isSkipEmptyString {
		slice = append(
			slice,
			prependString)
	}

	for _, item := range anyItems {
		if item == nil {
			continue
		}

		currentStr := compilerFunc(item)

		if isSkipEmptyString && currentStr == "" {
			continue
		}

		slice = append(
			slice,
			currentStr)
	}

	appendString := compilerFunc(appendItem)

	if isSkipEmptyString && appendString != "" {
		slice = append(
			slice,
			appendString)
	} else if !isSkipEmptyString {
		slice = append(
			slice,
			appendString)
	}

	return slice
}
