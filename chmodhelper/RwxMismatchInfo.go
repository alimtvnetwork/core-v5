package chmodhelper

type RwxMismatchInfo struct {
	FilePath          string
	Expecting, Actual string
}
