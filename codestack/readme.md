# codestack

Runtime call-stack capture, trace formatting, and stack trace collection. Provides structured access to Go's `runtime.Caller` information with automatic filtering of standard library noise.

## Architecture

```
codestack/
├── Trace.go                 # Single stack frame: package, method, file, line
├── TraceCollection.go       # Ordered collection of Trace items with filtering, paging, serialization
├── FileWithLine.go          # Lightweight file-path + line-number pair
├── FileWithLiner.go         # Interface: FullFilePath() + LineNumber()
├── newCreator.go            # New.Create(skipIndex) — builds a single Trace from runtime.Caller
├── newStacksCreator.go      # New.StackTrace.* — builds TraceCollection with skip/count control
├── newTraceCollection.go    # New.traces.* — TraceCollection constructors (Cap, Using, Empty)
├── currentNameOf.go         # NameOf.* — extract package/method names from runtime function info
├── stacksTo.go              # StacksTo.* — convenience converters (String, Bytes, JsonString)
├── fileGetter.go            # File.* — current file path/name from runtime
├── dirGetter.go             # Dir.* — current directory, repo root from runtime
├── skippablePrefixes.go     # Go standard library package map + isSkippablePackage()
├── funcs.go                 # FilterFunc, Formatter type aliases
├── consts.go                # Format strings, skip/take constants
└── vars.go                  # Package-level singletons: NameOf, New, StacksTo, File, Dir
```

## Core Types

### Trace

A single stack frame captured via `runtime.Caller`. Fields:

| Field | Type | Description |
|-------|------|-------------|
| `SkipIndex` | `int` | The skip index used to capture this frame |
| `PackageName` | `string` | Extracted package name (e.g. `myapp/handlers`) |
| `MethodName` | `string` | Method or function name (e.g. `HandleRequest`) |
| `PackageMethodName` | `string` | Full qualified name (e.g. `myapp/handlers.HandleRequest`) |
| `FilePath` | `string` | Absolute file path |
| `Line` | `int` | Line number |
| `IsOkay` | `bool` | Whether `runtime.Caller` succeeded |
| `IsSkippable` | `bool` | True if the frame originates from a Go standard library package |

Key methods: `Message()`, `ShortString()`, `FileWithLine()`, `FileName()`, `Clone()`, `Json()`, `Dispose()`.

### TraceCollection

An ordered slice of `Trace` items with rich collection operations:

| Category | Methods |
|----------|---------|
| Access | `First`, `Last`, `FirstOrDefault`, `LastOrDefault` |
| Slicing | `Skip`, `Take`, `Limit`, `SafeLimitCollection` |
| Filtering | `Filter(FilterFunc)`, `FilterWithLimit`, `Find`, `FindFirst` |
| Paging | `GetPagesSize`, `GetPagedCollection`, `GetSinglePageCollection` |
| Output | `CodeStacksString`, `StackTracesBytes`, `Strings`, `JoinUsingFmt` |
| Mutation | `Add`, `Adds`, `InsertAt`, `ConcatNew` |
| Serialization | `Json`, `JsonPtr`, `JsonString` |
| Lifecycle | `Clone`, `Dispose` |

### FileWithLine

Lightweight struct holding `FilePath` + `Line`. Implements the `FileWithLiner` interface.

### FileWithLiner

```go
type FileWithLiner interface {
    FullFilePath() string
    LineNumber() int
}
```

Both `Trace` and `FileWithLine` satisfy this interface via `AsFileLiner()`.

## IsSkippable — Standard Library Filtering

Traces originating from Go's standard library are automatically flagged as `IsSkippable = true` during creation. These frames are **filtered out at collection-creation time** in `AddsUsingSkip` and `AddsUsingSkipUsingFilter`, so they never appear in `TraceCollection.Items`.

### How It Works

1. **Detection** — When `New.Create(skipIndex)` builds a `Trace`, it calls `isSkippablePackage(packageName)` against the extracted package name.
2. **Map lookup** — Defined in `skippablePrefixes.go` as a `map[string]bool` for O(1) lookup. Matches exact package names:

   ```
   net, net/http, runtime, reflect, fmt, strings, strconv, os, io, sync,
   encoding, crypto, math, testing, log, bytes, bufio, context,
   database/sql, path, sort, time, regexp, errors, syscall, unicode
   ```

3. **Filtering** — During stack collection (`AddsUsingSkip`, `AddsUsingSkipUsingFilter`), any trace where `IsSkippable == true` is skipped via `continue` and never appended to `Items`.

### Why

Stack traces from Go's core packages (e.g. `runtime.goexit`, `net/http.(*conn).serve`, `reflect.Value.Call`) add noise and are rarely useful for debugging application-level issues. Filtering them at creation time produces cleaner, more actionable stack traces.

## Entry Points

All access goes through package-level singletons defined in `vars.go`:

| Singleton | Type | Purpose |
|-----------|------|---------|
| `New` | `newCreator` | Create single traces or trace collections |
| `NameOf` | `currentNameOf` | Extract package/method names from the call stack |
| `StacksTo` | `stacksTo` | Convert stack traces to string/bytes/JSON |
| `File` | `fileGetter` | Get current file path or name |
| `Dir` | `dirGetter` | Get current directory or repo root |

## Usage Examples

### Capture a Single Trace

```go
trace := codestack.New.Default()
fmt.Println(trace.ShortString())
// Output: myapp/handlers.HandleRequest (42) -> /app/handlers/request.go:42
```

### Capture a Stack Trace Collection

```go
stacks := codestack.New.StackTrace.SkipOne()
fmt.Println(stacks.CodeStacksString())
// Only application frames — standard library frames are automatically excluded
```

### Get Current Method Name

```go
methodName := codestack.NameOf.Method()
// Returns: "HandleRequest"
```

### Convert to JSON

```go
jsonStr := codestack.StacksTo.JsonStringDefault()
```

### Get Current File Path and Line

```go
filePath, lineNumber := codestack.File.PathLineSepDefault()
```

## Skip Index Constants

The package provides named constants for common skip values:

- `SkipNone` (0) through `Skip20` (20) — how many stack frames to skip
- `Take2` through `Take200` — how many frames to capture
- `DefaultStackCount` = 12

## Function Types

```go
type FilterFunc func(trace *Trace) (isTake, isBreak bool)
type Formatter  func(trace *Trace) (output string)
```

`FilterFunc` is used by `Filter`, `FilterWithLimit`, and `AddsUsingSkipUsingFilter` for custom trace selection. `Formatter` is used by `JoinUsingFmt` and `StringUsingFmt` for custom output formatting.

## Related Docs

- [Remaining Packages](../spec/01-app/folders/10-remaining-packages.md) — package listing with `codestack` summary
