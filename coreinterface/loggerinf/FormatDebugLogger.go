package loggerinf

type FormatDebugLogger interface {
	DebugFmt(formatter string, args ...any) // Debug logs a message at Debug level.
}
