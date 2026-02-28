package isany

func DefinedAnyOf(anyItems ...any) bool {
	if len(anyItems) == 0 {
		return false
	}

	for _, anyItem := range anyItems {
		if !Null(anyItem) {
			return true
		}
	}

	return false
}
