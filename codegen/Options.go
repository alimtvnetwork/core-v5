package codegen

type Options struct {
	IsGenerateInSameFile  bool
	IsWriteTestCasesFirst bool
	IsIncludeFunction     bool
	IsOverwrite           bool
}

func (it Options) IsOverwriteNoSupported() bool {
	return !it.IsOverwrite
}
