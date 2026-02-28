package loggerinf

type FormatInfoLogger interface {
	InfoFmt(format string, args ...any) // Info logs a message at Info level.
}
