package issetter

func generateDynamicRangesMap() map[string]interface{} {
	newMap := make(map[string]interface{}, len(values))

	for i, name := range values {
		newMap[name] = i
	}

	return newMap
}
