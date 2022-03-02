package loggerinf

type SpecificValuer interface {
	IsAnyValueDefined() bool
	BytesVal() []byte
	StringVal() string
	BooleanVal() bool
	IntegerVal() int
	ByteVal() byte
}
