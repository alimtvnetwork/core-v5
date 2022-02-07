package baseactioninf

type SimpleExecutor interface {
	Name() string
	TypeName() string
	Executor
	IsApply() (isSuccess bool)
}
