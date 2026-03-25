package isanytests

func convertFuncType(i any) (resultFunc isBoolCheckerFunc) {
	if f, ok := i.(func(x any) bool); ok {
		return f
	}

	return nil
}
