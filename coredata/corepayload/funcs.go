package corepayload

type (
	FilterFunc     func(payloadWrapper *PayloadWrapper) (isTake, isBreak bool)
	SkipFilterFunc func(payloadWrapper *PayloadWrapper) (isSkip, isBreak bool)
	Formatter      func(payloadWrapper *PayloadWrapper) (out string)
)
