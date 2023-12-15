package codegen

import (
	"go/parser"
	"path"
)

type newAstReaderCreator struct{}

// Create
//
// src is usually a string Golang code.
func (it newAstReaderCreator) Create(filePath string, src any) *AstReader {
	return &AstReader{
		filePath: path.Clean(filePath),
		src:      src,
		mode:     parser.AllErrors,
	}
}

func (it newAstReaderCreator) All(
	filePath string,
	src any,
	mode parser.Mode,
) *AstReader {
	return &AstReader{
		filePath: path.Clean(filePath),
		src:      src,
		mode:     mode,
	}
}

// Src
//
// src is usually a string Golang code.
func (it newAstReaderCreator) Src(src any) *AstReader {
	return &AstReader{
		src:  src,
		mode: parser.AllErrors,
	}
}

func (it newAstReaderCreator) FilePath(filePath string) *AstReader {
	return &AstReader{
		filePath: path.Clean(filePath),
		mode:     parser.AllErrors,
	}
}
