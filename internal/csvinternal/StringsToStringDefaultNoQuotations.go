package csvinternal

import "github.com/alimtvnetwork/core/constants"

func StringsToStringDefaultNoQuotations(
	references ...string,
) string {
	return StringsToCsvString(
		constants.CommaSpace,
		false,
		false,
		references...)
}
