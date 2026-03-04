package args

// emptyCreator provides factory methods for creating empty/invalid
// instances of common arg types.
type emptyCreator struct{}

// Map returns an empty Map.
func (it emptyCreator) Map() Map {
	return map[string]any{}
}

// FuncWrap returns an invalid FuncWrapAny.
func (it emptyCreator) FuncWrap() *FuncWrapAny {
	return &FuncWrapAny{
		isInvalid: true,
	}
}

// FuncMap returns an empty FuncMap.
func (it emptyCreator) FuncMap() FuncMap {
	return map[string]FuncWrapAny{}
}
