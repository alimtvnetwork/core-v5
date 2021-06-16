package stringslice

func FirstLastDefaultStatus(slice []string) *FirstLastStatus {
	length := len(slice)

	if length == 0 {
		return InvalidFirstLastStatus()
	}

	if length == 1 {
		return InvalidFirstLastStatusForLast(slice[0])
	}

	// length >= 2
	return &FirstLastStatus{
		First:    slice[0],
		Last:     slice[length-1],
		IsValid:  true,
		HasFirst: true,
		HasLast:  true,
	}
}
