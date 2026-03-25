package loggerinf

type ErrorLogger interface {
	Error(args ...any) // Error logs a message at Error level.
}
