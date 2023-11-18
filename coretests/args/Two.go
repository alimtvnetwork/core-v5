package args

type Two struct {
	First  interface{} `json:",omitempty"`
	Second interface{} `json:",omitempty"`
	Expect interface{} `json:",omitempty"`
}
