package internalinterface

type BaseErrorTyper interface {
	NameWithNameEqualer
	NameValue() string
	IsValid() bool
	IsInvalid() bool
	IsRawValue(rawValue uint16) bool
	IsNoError() bool
	IsEmptyError() bool
	HasError() bool
	Combine(
		additionalMessage,
		varName string,
		val any,
	) string
	CombineNoRefs(
		additionalMessage string,
	) string
	Error(
		additionalMessage,
		varName string,
		val any,
	) error
	ErrorReferences(
		additionalMessage string,
		references ...any,
	) error
	ErrorNoRefs(
		additionalMessage string,
	) error
	Panic(
		additionalMessage,
		varName string,
		val any,
	)
	PanicNoRefs(
		additionalMessage string,
	)
	// CodeWithTypeName
	//
	// 	errconsts.ErrorCodeHyphenTypeNameFormat  = "(#%d - %s)"
	CodeWithTypeName() string
	TypeName() string
	CodeTypeNameWithCustomMessage(
		customMessage string,
	) string
	ReferencesCsv(
		additionalMessage string,
		references ...any,
	) string
	ReferencesLines(
		additionalMessage string,
		referencesLines ...any,
	) string
	ReferencesLinesError(
		additionalMessage string,
		referencesLines ...any,
	) error
	ReferencesCsvError(
		additionalMessage string,
		references ...any,
	) error
	ShortReferencesCsv(
		references ...any,
	) string
	ShortReferencesCsvError(
		references ...any,
	) error
	RawValue() uint16
	Value() uint16
	ValueInt16() int16
	ValueInt() int
	ValueUInt() uint
}
