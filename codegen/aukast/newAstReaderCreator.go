package aukast

import (
	"go/parser"
	"go/token"
	"path"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/internal/pathinternal"
)

type newAstReaderCreator struct{}

// Create
//
// src is usually a string Golang code.
func (it newAstReaderCreator) Create(filePath string, src any) (*AstReader, error) {
	currentMode := parser.AllErrors

	return it.All(filePath, src, currentMode)
}

func (it newAstReaderCreator) All(
	filePath string,
	src any,
	mode parser.Mode,
) (*AstReader, error) {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(
		fileSet,
		filePath,
		src,
		mode,
	)

	if err != nil {
		return &AstReader{
			filePath: filePath,
			src:      src,
			mode:     mode,
		}, err
	}

	fullCode, fileErr := chmodhelper.
		SimpleFileWriter.
		FileReader.
		Read(filePath)

	return &AstReader{
		filePath: path.Clean(filePath),
		src:      src,
		astFile:  node,
		fullCode: fullCode,
		fileSet:  fileSet,
		mode:     mode,
	}, fileErr
}

// Src
//
// src is usually a string Golang code.
func (it newAstReaderCreator) Src(src any, mode parser.Mode) (*AstReader, error) {
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(
		fileSet,
		"",
		src,
		mode,
	)

	if err != nil {
		return &AstReader{
			src:  src,
			mode: mode,
		}, err
	}

	fullCode, astErr := astUtil.AstFileToCode(fileSet, astFile)

	return &AstReader{
		src:      src,
		astFile:  astFile,
		fullCode: fullCode,
		fileSet:  fileSet,
		mode:     mode,
	}, astErr
}

func (it newAstReaderCreator) FilePath(filePath string) (*AstReader, error) {
	return it.All(
		pathinternal.Clean(filePath),
		nil,
		parser.AllErrors,
	)
}
