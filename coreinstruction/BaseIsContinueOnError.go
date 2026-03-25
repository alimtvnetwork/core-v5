package coreinstruction

type BaseIsContinueOnError struct {
	IsContinueOnError bool `json:"IsContinueOnError"`
}

func (it *BaseIsContinueOnError) IsExitOnError() bool {
	return it != nil && !it.IsContinueOnError
}
