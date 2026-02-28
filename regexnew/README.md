# regexnew — Lazy Regex Package

## Overview

The `regexnew` package provides a **lazy-loaded, thread-safe** regex compilation system using the [New Creator Pattern](../spec/01-app/21-new-creator-pattern.md). It ensures regex patterns are compiled only once, cached globally, and safe for concurrent use.

## Specs

- [New Creator Pattern](../spec/01-app/21-new-creator-pattern.md)
- [Regex Implementation Details](../spec/01-app/) *(see architecture docs)*

## Quick Start

### Creating a LazyRegex

```go
// From package-level vars (non-locking, for init-time use)
var emailRegex = regexnew.New.Lazy(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// From inside methods (locking, for runtime use)
func validateInput(input string) bool {
    regex := regexnew.New.LazyLock(`^\d{3}-\d{4}$`)
    return regex.IsMatch(input)
}
```

### Using the New Creator

The `regexnew.New` variable provides the entry point:

```go
// LazyRegex — lazy-compiled, cached, thread-safe
regex := regexnew.New.Lazy("pattern")           // for var declarations
regex := regexnew.New.LazyLock("pattern")        // for method-internal use

// Multiple patterns at once (locked)
first, second := regexnew.New.LazyRegex.TwoLock("pattern1", "pattern2")
allMap := regexnew.New.LazyRegex.ManyUsingLock("p1", "p2", "p3")

// Direct compilation (standard regexp)
compiled, err := regexnew.New.Default("pattern")
compiled, err := regexnew.New.DefaultLock("pattern")
compiled, err := regexnew.New.DefaultLockIf(isLock, "pattern")

// Applicability check (compile + check in one call)
regex, err, isApplicable := regexnew.New.DefaultApplicableLock("pattern")
```

## LazyRegex Methods

### Lifecycle

| Method | Description |
|--------|-------------|
| `IsNull()` | Returns true if the LazyRegex pointer is nil |
| `IsDefined()` | Returns true if pattern and compiler are set |
| `IsUndefined()` | Opposite of IsDefined |
| `IsCompiled()` | Returns true if compilation has been attempted |
| `IsApplicable()` | Returns true if compiled successfully (triggers compilation if needed) |
| `IsInvalid()` | Opposite of IsApplicable |
| `HasError()` | Returns true if compilation produced an error |
| `HasAnyIssues()` | Returns true if nil, undefined, or has compilation errors |

### Compilation

| Method | Description |
|--------|-------------|
| `Compile()` | Compiles the pattern (once). Returns `(*regexp.Regexp, error)` |
| `CompileMust()` | Compiles or panics on error. Returns `*regexp.Regexp` |
| `OnRequiredCompiled()` | Ensures compilation happened. Returns error if any |
| `OnRequiredCompiledMust()` | Ensures compilation or panics |
| `MustBeSafe()` | Panics if there's a compilation error |

### Matching

| Method | Description |
|--------|-------------|
| `IsMatch(string)` | Returns true if the string matches the pattern |
| `IsMatchBytes([]byte)` | Returns true if the bytes match the pattern |
| `IsFailedMatch(string)` | Opposite of IsMatch |
| `IsFailedMatchBytes([]byte)` | Opposite of IsMatchBytes |
| `MatchError(string)` | Returns nil on match, error on mismatch |
| `MatchUsingFuncError(string, func)` | Custom match function with error reporting |
| `FirstMatchLine(string)` | Returns the first submatch or empty string |

### Inspection

| Method | Description |
|--------|-------------|
| `Pattern()` | Returns the raw pattern string |
| `String()` | Returns the pattern (implements Stringer) |
| `FullString()` | Returns JSON with pattern, compiled status, applicable status, error |
| `Error()` | Returns compilation error (alias for OnRequiredCompiled) |
| `CompiledError()` | Same as Error |

## New Creator Structure

```
regexnew.New (newCreator)
├── Lazy(pattern)              → *LazyRegex (var-level, no lock)
├── LazyLock(pattern)          → *LazyRegex (method-level, locked)
├── Default(pattern)           → (*regexp.Regexp, error)
├── DefaultLock(pattern)       → (*regexp.Regexp, error)
├── DefaultLockIf(bool, str)   → (*regexp.Regexp, error)
├── DefaultApplicableLock(str) → (*regexp.Regexp, error, bool)
└── LazyRegex (newLazyRegexCreator)
    ├── New(pattern)           → *LazyRegex
    ├── NewLock(pattern)       → *LazyRegex
    ├── NewLockIf(bool, str)   → *LazyRegex
    ├── TwoLock(p1, p2)        → (first, second *LazyRegex)
    ├── ManyUsingLock(ps...)   → map[string]*LazyRegex
    └── AllPatternsMap()       → map[string]*LazyRegex
```

## Thread Safety

- **`Lazy` / `New`**: No mutex; use only at package-level `var` declarations where Go guarantees single-goroutine init.
- **`LazyLock` / `NewLock`**: Uses `sync.Mutex`; safe for concurrent method calls.
- **`TwoLock` / `ManyUsingLock`**: Batch creation under a single lock for efficiency.
- **Pattern caching**: All LazyRegex instances are stored in a global `lazyRegexMap` to ensure each pattern is compiled at most once.

## Examples

### Email Validation

```go
var emailPattern = regexnew.New.Lazy(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func IsValidEmail(email string) bool {
    return emailPattern.IsMatch(email)
}
```

### Dynamic Pattern Matching

```go
func MatchDynamic(pattern, input string) (bool, error) {
    regex := regexnew.New.LazyLock(pattern)
    
    if regex.HasError() {
        return false, regex.Error()
    }
    
    return regex.IsMatch(input), nil
}
```

### Batch Pattern Registration

```go
var patterns = regexnew.New.LazyRegex.ManyUsingLock(
    `^\d+$`,
    `^[a-zA-Z]+$`,
    `^[a-zA-Z0-9]+$`,
)

func MatchAny(input string) bool {
    for _, regex := range patterns {
        if regex.IsMatch(input) {
            return true
        }
    }
    return false
}
```

## Contributors

## Issues for Future Reference
