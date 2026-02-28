package namevalue

func PrependsIf(
	isAdd bool,
	nameValues []StringAny,
	prependingItems ...StringAny,
) []StringAny {
	if !isAdd || len(prependingItems) == 0 {
		return nameValues
	}

	nameValues = append(prependingItems, nameValues...)

	return nameValues
}
