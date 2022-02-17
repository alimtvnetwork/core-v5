package internalenuminf

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
