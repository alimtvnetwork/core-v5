package stringslice

type IndexValuesDetail struct {
	Values         []string
	MissingIndexes []int
	IsAnyMissing   bool
	IsValid        bool // represent all missing or not found any
}

func InvalidIndexValuesDetail() *IndexValuesDetail {
	return &IndexValuesDetail{
		Values:         []string{},
		MissingIndexes: []int{},
		IsAnyMissing:   true,
		IsValid:        false,
	}
}
