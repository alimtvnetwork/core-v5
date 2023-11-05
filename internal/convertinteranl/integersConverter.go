package convertinteranl

type integersConverter struct{}

func (it integersConverter) ToMapBool(
	items ...int,
) (hashMap map[int]bool) {
	if len(items) == 0 {
		return map[int]bool{}
	}

	hashMap = make(map[int]bool, len(items))

	for _, item := range items {
		hashMap[item] = true
	}

	return hashMap
}

func (it integersConverter) Int8ToMapBool(
	items ...int8,
) (hashMap map[int8]bool) {
	if len(items) == 0 {
		return map[int8]bool{}
	}

	hashMap = make(map[int8]bool, len(items))

	for _, item := range items {
		hashMap[item] = true
	}

	return hashMap
}
