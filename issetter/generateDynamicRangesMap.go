package issetter

func generateDynamicRangesMap() map[string]any {
	newMap := make(map[string]any, len(valuesNames))

	for i, name := range valuesNames {
		newMap[name] = i
	}

	return newMap
}
