package codestack

// skippablePackages contains Go standard library and core framework
// package names. Stack traces originating from these packages are flagged
// as skippable since they add noise and are rarely useful for debugging
// application-level issues. Uses a map for O(1) lookup.
var skippablePackages = map[string]bool{
	"net":          true,
	"net/http":     true,
	"runtime":      true,
	"reflect":      true,
	"fmt":          true,
	"strings":      true,
	"strconv":      true,
	"os":           true,
	"io":           true,
	"sync":         true,
	"encoding":     true,
	"crypto":       true,
	"math":         true,
	"testing":      true,
	"log":          true,
	"bytes":        true,
	"bufio":        true,
	"context":      true,
	"database/sql": true,
	"path":         true,
	"sort":         true,
	"time":         true,
	"regexp":       true,
	"errors":       true,
	"syscall":      true,
	"unicode":      true,
}

// isSkippablePackage checks whether the given package name matches any of the
// Go standard library packages, indicating the trace is from core framework
// code rather than application code.
func isSkippablePackage(packageName string) bool {
	return skippablePackages[packageName]
}
