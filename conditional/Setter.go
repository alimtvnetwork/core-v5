package conditional

import "gitlab.com/evatix-go/core/issetter"

func Setter(
	isTrue bool,
	trueValue, falseValue issetter.Value,
) interface{} {
	if isTrue {
		return trueValue
	}

	return falseValue
}
