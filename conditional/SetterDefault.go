package conditional

import "github.com/alimtvnetwork/core/issetter"

func SetterDefault(
	currentSetter issetter.Value,
	defVal issetter.Value,
) issetter.Value {
	if currentSetter.IsUnSetOrUninitialized() {
		return defVal
	}

	return currentSetter
}
