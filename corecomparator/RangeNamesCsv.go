package corecomparator

import "github.com/alimtvnetwork/core/internal/csvinternal"

func RangeNamesCsv() string {
	return csvinternal.RangeNamesWithValuesIndexesCsvString(
		CompareNames[:]...)
}
