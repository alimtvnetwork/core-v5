package internalinterface

type PathExtender interface {
	FullPath() string
	FileName() string
	Extension() string
	Root() string
	Relative() string
	ParentDir() string
	IsFile() bool
}
