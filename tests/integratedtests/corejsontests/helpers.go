package corejsontests

// exampleStruct is a simple struct used across Unmarshal tests.
type exampleStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
