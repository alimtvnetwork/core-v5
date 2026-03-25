package stringslice

func AppendLineNew(slice []string, line string) []string {
	mergedNew := MergeNew(slice, line)

	return mergedNew
}
