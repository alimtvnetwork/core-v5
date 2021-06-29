package stringutil

func ClonePtr(strIn *string) (cloneOut *string) {
	if strIn == nil {
		return nil
	}

	var strNew string

	strNew = *strIn

	return &strNew
}
