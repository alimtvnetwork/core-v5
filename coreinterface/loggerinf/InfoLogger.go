package loggerinf

type InfoLogger interface {
	Info(args ...any) // Info logs a message at Info level.
}
