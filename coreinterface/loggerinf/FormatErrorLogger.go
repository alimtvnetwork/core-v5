package loggerinf

type FormatErrorLogger interface {
	ErrorFmt(
		format string,
		args ...any,
	)
	ErrorFmtIf(
		isLog bool,
		format string,
		args ...any,
	)
	ErrorFmtStackSkip(
		stackSkipIndex int,
		format string,
		args ...any,
	)

	// ErrorFmtUsingError
	//
	// Skip if no error
	ErrorFmtUsingError(
		format string,
		err error,
	)

	// ErrorFmtUsingErrorStackSkip
	//
	// Skip if no error
	ErrorFmtUsingErrorStackSkip(
		stackSkipIndex int,
		format string,
		err error,
	)
	WarnFmtStackSkip(
		stackSkipIndex int,
		format string,
		args ...any,
	)
	InfoFmtStackSkip(
		stackSkipIndex int,
		format string,
		args ...any,
	)
	WarnStackSkip(
		stackSkipIndex int,
		args ...any,
	)

	InfoStackSkip(
		stackSkipIndex int,
		args ...any,
	)
}
