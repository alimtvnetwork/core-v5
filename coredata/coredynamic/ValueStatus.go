package coredynamic

type ValueStatus struct {
	IsValid bool
	Message string
	Index   int
	Value   interface{}
}
