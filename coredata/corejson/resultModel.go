package corejson

type resultModel struct {
	Bytes    string `json:"Payload,omitempty"`
	Error    error  `json:"Error,omitempty"`
	TypeName string `json:"TypeName,omitempty"`
}
