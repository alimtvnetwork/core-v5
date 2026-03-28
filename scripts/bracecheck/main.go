// Package main provides a Go-based brace balance and syntax pre-checker for test files.
// Usage: go run ./scripts/bracecheck/ [files or dirs...]
// If no args, defaults to tests/integratedtests/corestrtests/
package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"tests/integratedtests/corestrtests/"}
	}

	var files []string
	for _, arg := range args {
		info, err := os.Stat(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s: %v\n", arg, err)
			os.Exit(1)
		}
		if info.IsDir() {
			entries, _ := os.ReadDir(arg)
			for _, e := range entries {
				if !e.IsDir() && strings.HasSuffix(e.Name(), ".go") {
					files = append(files, filepath.Join(arg, e.Name()))
				}
			}
		} else {
			files = append(files, arg)
		}
	}

	fset := token.NewFileSet()
	failed := 0

	for _, f := range files {
		_, err := parser.ParseFile(fset, f, nil, parser.AllErrors)
		if err != nil {
			fmt.Fprintf(os.Stderr, "✗ %s:\n  %v\n", f, err)
			failed++
		}
	}

	if failed > 0 {
		fmt.Fprintf(os.Stderr, "\n✗ %d file(s) have syntax errors\n", failed)
		os.Exit(1)
	}
	fmt.Printf("✓ %d file(s) parsed OK\n", len(files))
}
