package isany

func DefinedItems(
	anyItems ...any,
) (isAllDefined bool, definedItems []any) {
	if len(anyItems) == 0 {
		return false, nil
	}

	isAllDefined = true
	definedItems = make([]any, 0, len(anyItems))

	for _, anyItem := range anyItems {
		if Null(anyItem) {
			isAllDefined = false
		} else {
			// defined
			definedItems = append(definedItems, anyItem)
		}
	}

	return isAllDefined, definedItems
}
