package conditional

import "github.com/alimtvnetwork/core/constants"

func StringDefault(
	isTrue bool,
	trueValue string,
) string {
	if isTrue {
		return trueValue
	}

	return constants.EmptyString
}
