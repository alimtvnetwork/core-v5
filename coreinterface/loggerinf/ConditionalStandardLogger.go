package loggerinf

type ConditionalStandardLogger interface {
	On(isCondition bool) StandardLogger
	OnErr(err error) StandardLogger
	OnString(message string) StandardLogger
	OnBytes(rawBytes []byte) StandardLogger
}
