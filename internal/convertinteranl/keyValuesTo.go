package convertinteranl

type keyValuesTo struct{}

// ToMap keys nil will return empty map[string]string
func (it mapConverter) ToMap(keys, values []string) map[string]string {
	if keys == nil || values == nil {
		return map[string]string{}
	}

	newArray := make(map[string]string, len(keys))

	for i, key := range keys {
		newArray[key] = values[i]
	}

	return newArray
}

// ToMapPtr keys nil will return empty map[string]string
func (it mapConverter) ToMapPtr(keys, values *[]string) *map[string]string {
	if keys == nil || *keys == nil {
		var emptyResult map[string]string

		return &emptyResult
	}

	newArray := make(map[string]string, len(*keys))

	for i, key := range *keys {
		newArray[key] = (*values)[i]
	}

	return &newArray
}
