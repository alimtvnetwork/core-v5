package internalinterface

type BaseRawErrorTyper interface {
	String() string
	NameWithNameEqualer
	Combine(
		otherMsg string, reference any,
	) string
	TypesAttach(
		otherMsg string,
		reflectionTypes ...any,
	) string
	TypesAttachErr(
		otherMsg string,
		reflectionTypes ...any,
	) error
	SrcDestination(
		otherMsg string,
		srcName string, srcValue any,
		destinationName string, destinationValue any,
	) string
	SrcDestinationErr(
		otherMsg string,
		srcName string, srcValue any,
		destinationName string, destinationValue any,
	) error
	Error(otherMsg string, reference any) error
	MsgCsvRef(
		otherMsg string,
		csvReferenceItems ...any,
	) string
	MsgCsvRefError(
		otherMsg string,
		csvReferenceItems ...any,
	) error
	ErrorRefOnly(reference any) error
	Expecting(expecting, actual any) error
	NoRef(otherMsg string) string
	ErrorNoRefs(otherMsg string) error
	HandleUsingPanic(otherMsg string, reference any)
}
