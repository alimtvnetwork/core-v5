package csvinternal

import "github.com/alimtvnetwork/core/constants"

func AnyItemsToStringDefault(
	references ...any,
) string {
	return AnyItemsToCsvString(
		constants.CommaSpace,
		true,
		false,
		references...)
}
