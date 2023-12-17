package aukast

import "fmt"

type AstCollection struct {
	Parent     *AstElem
	childNodes []AstElem
}

func (it *AstCollection) IsEmpty() bool {
	return it == nil ||
		it.Parent == nil ||
		it.Parent.astReader.IsEmpty() ||
		len(it.childNodes) == 0
}

func (it *AstCollection) IsInvalid() bool {
	return it.IsEmpty()
}

func (it *AstCollection) IsValid() bool {
	return !it.IsEmpty()
}

func (it *AstCollection) FullCode() string {
	if it.IsEmpty() {
		return ""
	}

	return it.Parent.FullCode()
}

func (it *AstCollection) RawChildNodes() []AstElem {
	if it.IsEmpty() {
		return []AstElem{}
	}

	return it.childNodes
}

func (it AstCollection) String() string {
	if it.IsEmpty() {
		return ""
	}

	return fmt.Sprintf(
		"Parent:%s" +
			"Childs:%d" +
			"Parent:%s" +
			"Parent:%s" +
			"Parent:%s" +
			"Parent:%s" +
			"Parent:%s" +
			"Parent:%s" +,
	)
}
