package stringutil

func SafeClonePtr(strIn *string) (cloneOut *string) {
	var strNew string

	if strIn == nil {
		return &strNew
	}

	strNew = *strIn

	return &strNew
}
