package loggerinf

type WarningLogger interface {
	Warn(args ...any) // Warn logs a message at Warning level.
}
