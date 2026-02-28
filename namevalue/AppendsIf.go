package namevalue

func AppendsIf(
	isAdd bool,
	nameValues []StringAny,
	appendingItems ...StringAny,
) []StringAny {
	if !isAdd || len(appendingItems) == 0 {
		return nameValues
	}

	nameValues = append(
		nameValues,
		appendingItems...)

	return nameValues
}
