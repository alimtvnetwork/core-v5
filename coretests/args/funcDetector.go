package args

type funcDetector struct{}

func (it funcDetector) GetFuncWrap(i interface{}) *FuncWrap {
	switch v := i.(type) {
	case Map:
		return v.FuncWrap()
	case *FuncWrap:
		return v
	case OneFunc, TwoFunc, ThreeFunc:
		return v.FuncWrap()
	}
}
