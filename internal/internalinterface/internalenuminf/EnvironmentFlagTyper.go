package internalenuminf

type EnvironmentFlagTyper interface {
	BasicEnumer
	ByteValuePlusEqualer

	ToNumberString() string
	IsValidInvalidChecker

	IsVerbose() bool
	IsDebug() bool
	IsProduction() bool
	IsStackTrace() bool
	IsJson() bool
	IsLog() bool
	IsLoggedIn() bool

	SessionDetailsDynamic() []byte
	KeyValues() map[string]string
	AnyKeyValues() map[string]any
	NameBoolMap() map[string]bool

	FlagValue(name string) string
	FlagAnyValue(name string) any
	FlagAnyValueReflectSet(name string, toPointer any) error

	HasFlag(name string) bool
	HasAnyFlag(names ...string) bool
	HasAllFlags(names ...string) bool

	IsFlagEnabled(name string) bool
	IsFlagDisabled(name string) bool
}
