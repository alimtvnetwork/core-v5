package corevalidator

type BaseLinesValidators struct {
	LinesValidators []LineValidator `json:"LinesValidators,omitempty"`
}

func (it *BaseLinesValidators) LinesValidatorsLength() int {
	if it == nil {
		return 0
	}

	return len(it.LinesValidators)
}

func (it *BaseLinesValidators) IsEmptyLinesValidators() bool {
	return it.LinesValidatorsLength() == 0
}

func (it *BaseLinesValidators) HasLinesValidators() bool {
	return it.LinesValidatorsLength() > 0
}

func (it *BaseLinesValidators) ToLinesValidators() *LinesValidators {
	length := it.LinesValidatorsLength()
	linesValidators := NewLinesValidators(length)

	if length == 0 {
		return linesValidators
	}

	return linesValidators.Adds(it.LinesValidators...)
}
