package isany

// ReflectNotNull
//
// Returns true for not nil.
//
// Reference : https://stackoverflow.com/a/43896204
func ReflectNotNull(item any) bool {
	return !ReflectNull(item)
}
