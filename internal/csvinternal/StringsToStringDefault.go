package csvinternal

import "github.com/alimtvnetwork/core/constants"

func StringsToStringDefault(
	references ...string,
) string {
	return StringsToCsvString(
		constants.CommaSpace,
		true,
		false,
		references...)
}
