// Package main provides an auto-fixer for common Go test syntax errors.
// It iteratively parses files, detects known error patterns, applies fixes, and re-checks.
// Usage: go run ./scripts/autofix/ [--dry-run] [files or dirs...]
// If no args, defaults to tests/integratedtests/corestrtests/
package main

import (
	"fmt"
	"go/parser"
	"go/scanner"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var dryRun bool

func main() {
	var args []string
	for _, a := range os.Args[1:] {
		if a == "--dry-run" {
			dryRun = true
		} else {
			args = append(args, a)
		}
	}
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

	totalFixed := 0
	totalFiles := 0

	for _, f := range files {
		n := fixFile(f)
		if n > 0 {
			totalFiles++
			totalFixed += n
			if dryRun {
				fmt.Printf("  → %s: %d fix(es) would be applied\n", f, n)
			} else {
				fmt.Printf("  ✓ %s: %d fix(es) applied\n", f, n)
			}
		}
	}

	if totalFixed == 0 {
		fmt.Println("✓ No fixable issues found.")
	} else if dryRun {
		fmt.Printf("\n→ Would apply %d fix(es) across %d file(s). (dry-run, no files modified)\n", totalFixed, totalFiles)
	} else {
		fmt.Printf("\n✓ Applied %d fix(es) across %d file(s).\n", totalFixed, totalFiles)
		fmt.Println("  Run bracecheck again to verify: go run ./scripts/bracecheck/")
	}
}

// fixFile attempts up to maxPasses of parse-fix cycles on a single file.
func fixFile(path string) int {
	const maxPasses = 10
	totalFixes := 0

	for pass := 0; pass < maxPasses; pass++ {
		src, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "  ✗ %s: read error: %v\n", path, err)
			return totalFixes
		}

		fset := token.NewFileSet()
		_, parseErr := parser.ParseFile(fset, path, src, parser.AllErrors)
		if parseErr == nil {
			return totalFixes // clean
		}

		errList, ok := parseErr.(scanner.ErrorList)
		if !ok {
			return totalFixes
		}

		lines := strings.Split(string(src), "\n")
		fixes := 0

		// Process errors in reverse line order so line numbers stay valid
		applied := make(map[int]bool) // track lines already modified this pass
		for i := len(errList) - 1; i >= 0; i-- {
			e := errList[i]
			lineIdx := e.Pos.Line - 1
			if lineIdx < 0 || lineIdx >= len(lines) || applied[lineIdx] {
				continue
			}

			fixed := false
			switch {
			case strings.Contains(e.Msg, "missing ',' before newline in argument list"):
				fixed = fixMissingComma(lines, lineIdx)

			case strings.Contains(e.Msg, "expected statement, found ')'"):
				fixed = fixUnexpectedCloseParen(lines, lineIdx)

			case strings.Contains(e.Msg, "expected declaration, found ')'"):
				fixed = fixUnexpectedCloseParenTopLevel(lines, lineIdx)

			case strings.Contains(e.Msg, "expected '}', found 'EOF'"):
				fixed = fixMissingCloseBrace(lines)

			case strings.Contains(e.Msg, "expected 1 expression"):
				fixed = fixExpectedOneExpression(lines, lineIdx)
			}

			if fixed {
				fixes++
				applied[lineIdx] = true
			}
		}

		if fixes == 0 {
			return totalFixes // no more auto-fixable errors
		}

		totalFixes += fixes

		if dryRun {
			// Don't write changes in dry-run mode; stop after first pass
			return totalFixes
		}

		result := strings.Join(lines, "\n")
		if err := os.WriteFile(path, []byte(result), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "  ✗ %s: write error: %v\n", path, err)
			return totalFixes
		}
	}
	return totalFixes
}

// fixMissingComma adds a trailing comma to the line before the error.
// The error points to the line AFTER the one missing the comma.
func fixMissingComma(lines []string, errLine int) bool {
	// The missing comma is on the previous non-empty line
	target := -1
	for i := errLine - 1; i >= 0; i-- {
		trimmed := strings.TrimSpace(lines[i])
		if trimmed == "" || strings.HasPrefix(trimmed, "//") {
			continue
		}
		target = i
		break
	}
	if target < 0 {
		return false
	}

	line := lines[target]
	trimmed := strings.TrimRight(line, " \t\r")

	// Don't add comma if line already ends with comma, opening brace/paren, or is a comment
	if trimmed == "" {
		return false
	}
	lastChar := trimmed[len(trimmed)-1]
	if lastChar == ',' || lastChar == '{' || lastChar == '(' || lastChar == '[' {
		return false
	}
	// Don't add comma after comment lines
	if strings.HasPrefix(strings.TrimSpace(trimmed), "//") {
		return false
	}

	lines[target] = trimmed + ","
	return true
}

// rxOnlyCloseParen matches lines that are only whitespace + ")" or "),"
var rxOnlyCloseParen = regexp.MustCompile(`^\s*\)\s*,?\s*$`)

// fixUnexpectedCloseParen handles "expected statement, found ')'"
// Common cause: extra ')' inside a closure, or a }) that should be just }
func fixUnexpectedCloseParen(lines []string, errLine int) bool {
	if errLine < 0 || errLine >= len(lines) {
		return false
	}
	trimmed := strings.TrimSpace(lines[errLine])

	// Case 1: Line is just ")" — likely extra paren, remove it
	if trimmed == ")" {
		// Check if previous non-empty line ends with "}" — this is a "})" split across lines
		prev := findPrevNonEmpty(lines, errLine)
		if prev >= 0 && strings.TrimSpace(lines[prev]) == "}" {
			// Merge: change prev to "})" and remove current line
			indent := leadingWhitespace(lines[prev])
			lines[prev] = indent + "})"
			lines = removeLineInPlace(lines, errLine)
			return true
		}
		// Otherwise just remove the stray ")"
		lines[errLine] = ""
		return true
	}

	// Case 2: Line has "})}" or similar — normalize safeTest closure
	if trimmed == "})}" {
		indent := leadingWhitespace(lines[errLine])
		lines[errLine] = indent + "})"
		return true
	}

	return false
}

// fixUnexpectedCloseParenTopLevel handles "expected declaration, found ')'"
// Usually a stray ')' at top level from a mismatched safeTest closure.
func fixUnexpectedCloseParenTopLevel(lines []string, errLine int) bool {
	if errLine < 0 || errLine >= len(lines) {
		return false
	}
	trimmed := strings.TrimSpace(lines[errLine])

	// If the line is just ")" at top level, remove it
	if trimmed == ")" {
		lines[errLine] = ""
		return true
	}

	return false
}

// fixMissingCloseBrace appends a closing "}" if the file ends without one.
func fixMissingCloseBrace(lines []string) bool {
	// Find last non-empty line
	for i := len(lines) - 1; i >= 0; i-- {
		trimmed := strings.TrimSpace(lines[i])
		if trimmed == "" {
			continue
		}
		if trimmed != "}" && trimmed != "})" {
			lines = append(lines, "}")
			return true
		}
		return false
	}
	return false
}

// --- helpers ---

func findPrevNonEmpty(lines []string, from int) int {
	for i := from - 1; i >= 0; i-- {
		if strings.TrimSpace(lines[i]) != "" {
			return i
		}
	}
	return -1
}

func leadingWhitespace(s string) string {
	for i, c := range s {
		if c != ' ' && c != '\t' {
			return s[:i]
		}
	}
	return s
}

func removeLineInPlace(lines []string, idx int) []string {
	// We can't change the slice header from the caller, so blank the line instead
	// This is simpler and avoids line-number drift in the same pass
	lines[idx] = ""
	return lines
}
