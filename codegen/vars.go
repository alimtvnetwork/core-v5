package codegen

import "gitlab.com/auk-go/core/coreindexes"

var (
	NewCodeOutput  = newCodeOutputCreator{}
	indexByNameMap = map[int]string{
		coreindexes.First:  "First",
		coreindexes.Second: "Second",
		coreindexes.Third:  "Third",
		coreindexes.Fourth: "Fourth",
		coreindexes.Fifth:  "Fifth",
		coreindexes.Sixth:  "Sixth",
	}
)
