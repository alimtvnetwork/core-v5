package internalenuminf

type HttpMethodTyper interface {
	BasicEnumer
	ByteValuePlusEqualer

	IsGetHttp() bool
	IsPutHttp() bool
	IsPostHttp() bool
	IsDeleteHttp() bool
	IsPatchHttp() bool

	IsAnyHttpMethod(methodNames ...string) bool
	IsAnyHttp() bool
	IsNotHttpMethod() bool
}

type ActionTyper interface {
	BasicEnumer
	ByteValuePlusEqualer

	IsStart() bool
	IsStop() bool
	IsRestart() bool
	IsReload() bool
	IsStopSleepStart() bool
	IsSuspend() bool
	IsPause() bool
	IsResumed() bool

	IsAnyAction() bool
	IsNotAnyAction() bool
}
