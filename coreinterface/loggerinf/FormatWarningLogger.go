package loggerinf

type FormatWarningLogger interface {
	WarnFmt(format string, args ...any) // Warn logs a message at Warning level.
}
