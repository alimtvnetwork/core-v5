package loggerinf

type LoggerStackSkip interface {
	FatalStackSkip(
		stackSkipIndex int,
		args ...any,
	)
	ErrorStackSkip(
		stackSkipIndex int,
		args ...any,
	)

	// ErrorUsingError
	//
	// Skip if no error
	ErrorUsingError(err error)

	// ErrorUsingErrorStackSkip
	//
	// Skip if no error
	ErrorUsingErrorStackSkip(
		stackSkipIndex int,
		err error,
	)

	ErrorIf(isLog bool, args ...any)
	DebugFmtIf(
		isLog bool,
		formatter string,
		args ...any,
	)
	DebugFmtStackSkip(
		stackSkipIndex int,
		format string,
		args ...any,
	)

	DebugIf(isLog bool, args ...any) // Debug logs a message at Debug level.
	DebugStackSkip(
		stackSkipIndex int,
		args ...any,
	)

	DebugIncludingStackTracesIf(
		isLog bool,
		stackSkipIndex int,
		args ...any,
	)
	DebugIncludingStackTraces(
		stackSkipIndex int,
		args ...any,
	)
}
