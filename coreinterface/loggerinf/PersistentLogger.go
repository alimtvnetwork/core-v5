package loggerinf

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/internal/internalinterface"
)

type PersistentLogger interface {
	LogPathExtender() internalinterface.PathExtender
	IsRotating() bool
	IsDbLogger() bool
	IsFileLogger() bool

	Config() interface{}
	ConfigReflectSetTo(toPointer interface{}) error
	internalinterface.IdStringerWithNamer

	Info(message string) PersistentLogger
	Trace(message string) PersistentLogger
	Debug(message string) PersistentLogger
	Warn(message string) PersistentLogger
	Error(message string) PersistentLogger
	Fatal(message string) PersistentLogger
	Panic(message string) PersistentLogger

	InfoAttr(message, attr string) PersistentLogger
	TraceAttr(message, attr string) PersistentLogger
	DebugAttr(message, attr string) PersistentLogger
	WarnAttr(message, attr string) PersistentLogger
	ErrorAttr(message, attr string) PersistentLogger
	FatalAttr(message, attr string) PersistentLogger
	PanicAttr(message, attr string) PersistentLogger

	InfoStackSkip(stackSkipIndex int, message string) PersistentLogger
	TraceStackSkip(stackSkipIndex int, message string) PersistentLogger
	DebugStackSkip(stackSkipIndex int, message string) PersistentLogger
	WarnStackSkip(stackSkipIndex int, message string) PersistentLogger
	ErrorStackSkip(stackSkipIndex int, message string) PersistentLogger
	FatalStackSkip(stackSkipIndex int, message string) PersistentLogger
	PanicStackSkip(stackSkipIndex int, message string) PersistentLogger

	InfoJson(jsonResult *corejson.Result) PersistentLogger
	ErrorJson(jsonResult *corejson.Result) PersistentLogger
	DebugJson(jsonResult *corejson.Result) PersistentLogger

	InfoBytes(rawBytes []byte) PersistentLogger
	ErrorBytes(rawBytes []byte) PersistentLogger
	DebugBytes(rawBytes []byte) PersistentLogger

	InfoTitleBytes(title string, rawBytes []byte) PersistentLogger
	ErrorTitleBytes(title string, rawBytes []byte) PersistentLogger
	DebugTitleBytes(title string, rawBytes []byte) PersistentLogger

	Log(message string) PersistentLogger
	LogRaw(logType LogTypeChecker, message, attr string) PersistentLogger
	LogRawStackSkip(stackSkipIndex int, logType LogTypeChecker, message, attr string) PersistentLogger
	Jsoner(logType LogTypeChecker, message string, jsonResult *corejson.Result) PersistentLogger
	JsonerStackSkip(
		stackSkipIndex int, logType LogTypeChecker, message string, jsonResult *corejson.Result,
	) PersistentLogger

	LogStackSkip(stackSkipIndex int, message string) PersistentLogger
	LogFmtStackSkip(stackSkipIndex int, format string, message string) PersistentLogger
	LogAttr(message, attr string) PersistentLogger
	LogAttrStackSkip(stackSkipIndex int, message, attr string) PersistentLogger

	AnErr(err error) PersistentLogger
	ErrorMessage(message string) PersistentLogger
	ErrorMessageAttr(message, attr string) PersistentLogger
	ErrorMessageStackSkip(stackSkipIndex int, message string) PersistentLogger
	ErrorMessageAttrStackSkip(stackSkipIndex int, message, attr string) PersistentLogger
	ErrorMessageFmtStackSkip(stackSkipIndex int, format string, message string) PersistentLogger

	DebugMessage(message string) PersistentLogger
	DebugMessageAttr(message, attr string) PersistentLogger
	Err(err error) PersistentLogger
	// ErrorStackTraces
	//
	// Includes stack-traces
	ErrorStackTraces(err error) PersistentLogger
	// DebugStackTraces
	//
	// Includes stack-traces
	DebugStackTraces(message string) PersistentLogger
	// DebugAttrStackTraces
	//
	// Includes stack-traces
	DebugAttrStackTraces(message, attr string) PersistentLogger
	StackTraces() PersistentLogger
	StackTracesSkip(stackSkipIndex int) PersistentLogger
	TitleStackTraces(title string) PersistentLogger
	TitleStackTracesSkip(stackSkipIndex int, title string) PersistentLogger
	HasError() bool
	CompileErrorWithTraces() error

	ConditionalStandardLogger
}
