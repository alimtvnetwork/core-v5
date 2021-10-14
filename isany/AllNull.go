package isany

func AllNull(anyItems ...interface{}) bool {
	if len(anyItems) == 0 {
		return false
	}

	for _, anyItem := range anyItems {
		if !Null(anyItem) {
			return false
		}
	}

	return true
}
