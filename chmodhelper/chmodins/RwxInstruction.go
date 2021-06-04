package chmodins

type RwxInstruction struct {
	RwxOwnerGroupOther
	IsSkipOnNonExist  bool `json:"IsSkipOnNonExist"`
	IsContinueOnError bool `json:"IsContinueOnError"`
	IsRecursive       bool `json:"IsRecursive"`
}
