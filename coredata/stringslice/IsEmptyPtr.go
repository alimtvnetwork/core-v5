package stringslice

func IsEmptyPtr(slice *[]string) bool {
	return slice == nil || len(*slice) == 0
}
