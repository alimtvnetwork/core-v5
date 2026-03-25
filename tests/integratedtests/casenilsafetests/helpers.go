package casenilsafetests

// sampleStruct is a test-only struct used to validate CaseNilSafe
// across all edge cases: pointer receivers, value receivers, void
// methods, multi-return, and methods with complex args.
type sampleStruct struct {
	Name  string
	Value int
}

// IsValid — pointer receiver, nil-safe, returns bool.
func (it *sampleStruct) IsValid() bool {
	if it == nil {
		return false
	}

	return it.Name != ""
}

// Length — pointer receiver, nil-safe, returns int.
func (it *sampleStruct) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Name)
}

// String — pointer receiver, nil-safe, returns string.
func (it *sampleStruct) String() string {
	if it == nil {
		return ""
	}

	return it.Name
}

// Reset — pointer receiver, nil-safe, void method.
func (it *sampleStruct) Reset() {
	if it == nil {
		return
	}

	it.Name = ""
	it.Value = 0
}

// Parse — pointer receiver, nil-safe, returns (int, error).
func (it *sampleStruct) Parse(input string) (int, error) {
	if it == nil {
		return 0, nil
	}

	return len(input), nil
}

// Lookup — pointer receiver, nil-safe, returns (string, bool).
func (it *sampleStruct) Lookup(key string) (string, bool) {
	if it == nil {
		return "", false
	}

	if key == it.Name {
		return it.Name, true
	}

	return "", false
}

// UnsafeMethod — pointer receiver, NOT nil-safe (panics on nil).
func (it *sampleStruct) UnsafeMethod() string {
	return it.Name // will panic if it == nil
}

// ValueString — value receiver, always panics on nil pointer.
func (it sampleStruct) ValueString() string {
	return it.Name
}
