package loggerinf

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/coreinterface/enuminf"
	"gitlab.com/auk-go/core/coreinterface/errcoreinf"
)

type StandardLogger interface {
	StandardLoggerChecker
	ConditionalStandardLogger

	FullLogger() FullLogger
	EnvOptioner() enuminf.EnvironmentOptioner

	TaskNamedLogger(
		taskName string,
	) StandardLogger

	TaskWithPayloadLogger(
		taskName string,
		payloadAny any, // can be bytes, payloadWrapper, can be any
	) StandardLogger

	GetLoggerByTaskName(taskName string) StandardLogger
	GetLoggerByTaskNamer(taskNamer enuminf.Namer) StandardLogger

	Success(args ...any) StandardLogger
	Info(args ...any) StandardLogger
	Trace(args ...any) StandardLogger
	Debug(args ...any) StandardLogger
	Warn(args ...any) StandardLogger
	Error(args ...any) StandardLogger
	Fatal(args ...any) StandardLogger
	Panic(args ...any) StandardLogger

	SuccessFmt(format string, args ...any) StandardLogger
	InfoFmt(format string, args ...any) StandardLogger
	TraceFmt(format string, args ...any) StandardLogger
	DebugFmt(format string, args ...any) StandardLogger
	WarnFmt(format string, args ...any) StandardLogger
	ErrorFmt(format string, args ...any) StandardLogger
	FatalFmt(format string, args ...any) StandardLogger
	PanicFmt(format string, args ...any) StandardLogger

	SuccessExtend() SingleLogger
	InfoExtend() SingleLogger
	TraceExtend() SingleLogger
	DebugExtend() SingleLogger
	WarnExtend() SingleLogger
	FatalExtend() SingleLogger
	PanicExtend() SingleLogger

	ErrorDirect(err error) StandardLogger
	OnErrStackTrace(err error) StandardLogger
	ErrInterface(errInf errcoreinf.BaseErrorOrCollectionWrapper) StandardLogger
	ErrInterfaceStackTraces(errInfWithStackTraces errcoreinf.BaseErrorOrCollectionWrapper) StandardLogger

	ReflectSetter

	InfoOrError(isError bool) SingleLogger
	Log(loggerType enuminf.LoggerTyper) StandardLogger

	ErrorJsoner(jsoner corejson.Jsoner) StandardLogger
	DebugJsoner(jsoner corejson.Jsoner) StandardLogger
	ErrorJsonerTitle(title string, jsoner corejson.Jsoner) StandardLogger
	DebugJsonerTitle(title string, jsoner corejson.Jsoner) StandardLogger
}
