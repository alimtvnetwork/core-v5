package args

// funcDetector provides utility methods for detecting and extracting
// FuncWrapAny instances from various input types.
type funcDetector struct{}

// GetFuncWrap extracts a FuncWrapAny from various input types.
// Supports Map, *FuncWrapAny, FuncWrapGetter, ArgsMapper,
// and falls back to creating a new FuncWrapAny via reflection.
func (it funcDetector) GetFuncWrap(i any) *FuncWrapAny {
	switch v := i.(type) {
	case Map:
		return v.FuncWrap()
	case *FuncWrapAny:
		return v
	case FuncWrapGetter:
		return v.FuncWrap()
	case ArgsMapper:
		return v.FuncWrap()
	default:
		return NewFuncWrap.Default(i)
	}
}
