package coreversiontests

const (
	defaultCreationFmt          = "%d : %s (compact: %s, display: %s)"
	defaultInvalidV1CreationFmt = "%d : invalid - %s (raw: %s)"
	defaultInvalidV2CreationFmt = "%d : invalid - %s "

	comparisonFmt = "%d : Left [%s, raw(%s)] %s [%s, raw(%s)] Right | Expect: %s - %t"
)
