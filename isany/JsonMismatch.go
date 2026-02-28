package isany

// JsonMismatch
//
//	Inverse of JsonEqual
func JsonMismatch(
	left, right any,
) bool {
	return !JsonEqual(left, right)
}
