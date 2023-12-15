package chmodhelper

import "os"

type fileReader struct{}

func (it fileReader) Read(filePath string) (string, error) {
	b, err := os.ReadFile(filePath)

	return string(b), err
}

func (it fileReader) ReadBytes(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
