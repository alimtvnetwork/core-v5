package codestack

// skippablePackagePrefixes contains Go standard library and core framework
// package prefixes. Stack traces originating from these packages are flagged
// as skippable since they add noise and are rarely useful for debugging
// application-level issues.
var skippablePackagePrefixes = []string{
	"net",
	"net/http",
	"runtime",
	"reflect",
	"fmt",
	"strings",
	"strconv",
	"os",
	"io",
	"sync",
	"encoding",
	"crypto",
	"math",
	"testing",
	"log",
	"bytes",
	"bufio",
	"context",
	"database/sql",
	"path",
	"sort",
	"time",
	"regexp",
	"errors",
	"syscall",
	"unicode",
}

// isSkippablePackage checks whether the given package name matches any of the
// Go standard library prefixes, indicating the trace is from core framework
// code rather than application code.
func isSkippablePackage(packageName string) bool {
	for _, prefix := range skippablePackagePrefixes {
		if packageName == prefix {
			return true
		}
	}

	return false
}
