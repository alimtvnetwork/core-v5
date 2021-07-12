package devenv

type Info struct {
	IsDev, IsProd, IsTest, IsVerbose, IsLog, IsWarning, IsDebug bool
	EnvName                                                     string
}
