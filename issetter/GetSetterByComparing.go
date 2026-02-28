package issetter

// GetSetterByComparing
//
// returns true value if any of ranges value matches
func GetSetterByComparing(
	trueVal, falseVal Value,
	expectedVal any,
	trueRanges ...any,
) Value {
	for _, s := range trueRanges {
		if s == expectedVal {
			return trueVal
		}
	}

	return falseVal
}
