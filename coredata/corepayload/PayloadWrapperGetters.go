package corepayload

// PayloadWrapperGetters.go — Read-only accessor methods extracted from PayloadWrapper.go
// This file contains all getter and query methods.

// PayloadsString returns the Payloads field as a string.
func (it *PayloadWrapper) PayloadsString() string {
	if it.IsEmpty() || len(it.Payloads) == 0 {
		return ""
	}

	return string(it.Payloads)
}
