package corecsv

import "github.com/alimtvnetwork/core/constants"

func DefaultCsv(
	references ...string,
) string {
	return StringsToCsvString(
		constants.CommaSpace,
		true,
		false,
		references...)
}
