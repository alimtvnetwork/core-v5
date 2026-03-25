package loggerinf

type DebugLogger interface {
	Debug(args ...any) // Debug logs a message at Debug level.
}
