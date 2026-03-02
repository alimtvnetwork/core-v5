package coredata

type PointerIntegers []*int

func (p PointerIntegers) Len() int {
	if p == nil {
		return 0
	}

	return len(p)
}

func (p PointerIntegers) Less(i, j int) bool {
	if p[i] == nil {
		return true
	}
	if p[j] == nil {
		return false
	}

	return *p[i] < *p[j]
}
func (p PointerIntegers) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
