package corepayload

import "gitlab.com/auk-go/core/coredata/corejson"

// PayloadWrapperJson.go — JSON-related methods extracted from PayloadWrapper.go
// These methods handle serialization, deserialization, and JSON representation.

// PayloadsJsonResult returns the payloads as a corejson.Result pointer.
func (it PayloadWrapper) PayloadsJsonResult() *corejson.Result {
	if it.IsEmpty() || len(it.Payloads) == 0 {
		return corejson.Empty.ResultPtr()
	}

	return corejson.NewResult.UsingTypeBytesPtr(
		attributesTypeName,
		it.Payloads,
	)
}

// PayloadsPrettyString returns the payloads formatted as pretty JSON.
func (it PayloadWrapper) PayloadsPrettyString() string {
	if it.IsEmpty() || len(it.Payloads) == 0 {
		return ""
	}

	return corejson.BytesToPrettyString(it.Payloads)
}
