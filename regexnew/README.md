# regexnew — Lazy-Compiled Regex

Package `regexnew` provides a **lazy-loaded, thread-safe** regex compilation system using the [New Creator Pattern](../spec/01-app/21-new-creator-pattern.md). Patterns are compiled only once, cached globally, and safe for concurrent use. All methods handle nil `*LazyRegex` receivers gracefully without panicking.

## Architecture

```
regexnew/
├── vars.go                              # Package-level singletons: New, regexMaps, lazyRegexOnceMap
├── newCreator.go                        # New.* — Lazy, LazyLock, Default, DefaultLock entry points
├── newLazyRegexCreator.go               # New.LazyRegex.* — TwoLock, ManyUsingLock batch creators
├── LazyRegex.go                         # LazyRegex struct: lazy-compiled, cached, nil-safe
├── lazyRegexMap.go                      # Global pattern → *LazyRegex cache with mutex
├── funcs.go                             # Standalone helpers
├── Create.go                            # Create(pattern) — direct compilation
├── CreateLock.go                        # CreateLock(pattern) — locked compilation
├── CreateLockIf.go                      # CreateLockIf(bool, pattern) — conditional lock
├── CreateApplicableLock.go              # CreateApplicableLock — compile + applicability check
├── CreateMust.go                        # CreateMust — panics on error
├── CreateMustLockIf.go                  # CreateMustLockIf — conditional lock, panics on error
├── NewMustLock.go                       # NewMustLock — locked LazyRegex, panics on error
├── IsMatchLock.go                       # IsMatchLock — one-shot locked match
├── IsMatchFailed.go                     # IsMatchFailed — negated match
├── MatchError.go                        # MatchError — match with error return
├── MatchErrorLock.go                    # MatchErrorLock — locked match with error
├── MatchUsingFuncErrorLock.go           # MatchUsingFuncErrorLock — custom match function
├── MatchUsingCustomizeErrorFuncLock.go  # MatchUsingCustomizeErrorFuncLock — custom error func
├── regExMatchValidationError.go         # Match validation error type
├── regexes-compiled.go                  # Pre-compiled common patterns
├── prettyJson.go                        # JSON pretty-print for LazyRegex
└── README.md
```

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

## Thread Safety

| Creator | Lock | When to Use |
|---------|------|-------------|
| `Lazy` / `New` | None | Package-level `var` declarations only (Go guarantees single-goroutine init) |
| `LazyLock` / `NewLock` | `sync.Mutex` | Inside methods — safe for concurrent calls |
| `TwoLock` / `ManyUsingLock` | Single `sync.Mutex` | Batch creation under one lock for efficiency |

All `LazyRegex` instances are stored in a global `lazyRegexMap` to ensure each pattern is compiled at most once.

## Usage

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

## Related Docs

- [New Creator Pattern](/spec/01-app/21-new-creator-pattern.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [Folder Spec](/spec/01-app/folders/10-remaining-packages.md)
