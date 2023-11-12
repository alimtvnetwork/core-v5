package corefuncs

// GetFuncFullName
//
// Get the function name, passing non function may result panic
func GetFuncFullName(i interface{}) string {
	f := GetFunc(i)

	if f == nil {
		return ""
	}

	return f.Name()
}

// GetFuncFullName
//
// Get the function name, passing non function may result panic
func GetFuncFullName(i interface{}) string {
	f := GetFunc(i)

	if f == nil {
		return ""
	}

	return f.Name()
}
