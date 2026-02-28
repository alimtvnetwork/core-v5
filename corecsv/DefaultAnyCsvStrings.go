package corecsv

func DefaultAnyCsvStrings(
	references ...any,
) []string {
	return AnyItemsToCsvStrings(
		true,
		false,
		references...)
}
