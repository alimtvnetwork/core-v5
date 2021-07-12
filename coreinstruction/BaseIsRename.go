package coreinstruction

type BaseIsRename struct {
	IsRename bool `json:"IsRename"`
}

func NewRename(isRename bool) BaseIsRename {
	return BaseIsRename{
		IsRename: isRename,
	}
}
