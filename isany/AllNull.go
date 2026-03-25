package isany

func AllNull(anyItems ...any) bool {
	if len(anyItems) == 0 {
		return true
	}

	for _, anyItem := range anyItems {
		if !Null(anyItem) {
			return false
		}
	}

	return true
}
