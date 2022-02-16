package corecomparator

import "gitlab.com/evatix-go/core/internal/csvinternal"

func RangeNamesCsv() string {
	return csvinternal.RangeNamesWithValuesIndexesCsvString(
		CompareNames[:]...)
}
