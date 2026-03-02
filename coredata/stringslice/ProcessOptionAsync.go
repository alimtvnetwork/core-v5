package stringslice

func ProcessOptionAsync(
	isSkipOnNil bool,
	processor func(index int, item any) string,
	items ...any,
) []string {
	if len(items) == 0 {
		return []string{}
	}

	list := ProcessAsync(processor, items...)

	isReturnAll := !isSkipOnNil

	if isReturnAll {
		return list
	}

	newSlice := make([]string, 0, len(list))

	for _, item := range list {
		if item == "" {
			continue
		}

		newSlice = append(newSlice, item)
	}

	return newSlice
}
