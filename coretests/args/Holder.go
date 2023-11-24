package args

// Holder
//
// it is used to hold more dynamic parameters for
// the Act function in the unit or integration test
//
// If parameters are not enough use the Hashmap
type Holder struct {
	First   interface{} `json:",omitempty"`
	Second  interface{} `json:",omitempty"`
	Third   interface{} `json:",omitempty"`
	Fourth  interface{} `json:",omitempty"`
	Fifth   interface{} `json:",omitempty"`
	Sixth   interface{} `json:",omitempty"`
	Expect  interface{} `json:",omitempty"`
	Hashmap Map         `json:",omitempty"`
}
