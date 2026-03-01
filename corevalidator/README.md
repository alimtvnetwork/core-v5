# corevalidator — Text, Line & Slice Validators

Package `corevalidator` provides composable validators for text matching, line-level comparison, and slice-level verification. Used extensively by the test framework (`coretests`) and assertion pipelines to validate multi-line output against expected patterns with configurable comparison methods, trimming, whitespace normalization, and sorting.

## Architecture

```
corevalidator/
├── TextValidator.go               # Single text match with configurable comparison
├── TextValidators.go              # Collection of TextValidator items
├── LineValidator.go               # Text match with line number verification
├── LinesValidators.go             # Collection of LineValidator items
├── SliceValidator.go              # Compare actual vs expected string slices (630+ lines)
├── SliceValidators.go             # Collection of SliceValidator items
├── SimpleSliceValidator.go        # Simplified slice validator
├── HeaderSliceValidator.go        # Slice validator with header metadata
├── HeaderSliceValidators.go       # Collection of HeaderSliceValidator items
├── BaseLinesValidators.go         # Base validator for line collections
├── BaseValidatorCoreCondition.go  # Shared condition base
├── Condition.go                   # Condition struct (trim, unique, sort flags)
├── LineNumber.go                  # Line number matching type
├── Parameter.go                   # Verification parameter bag
├── RangeSegmentsValidator.go      # Range segment validation
├── RangesSegment.go               # Range segment type
├── consts.go                      # Internal message format constants
└── vars.go                        # Pre-built condition & validator instances
```

## Core Types

### Condition

Controls text preprocessing before comparison:

```go
type Condition struct {
    IsTrimCompare        bool // Trim whitespace before comparing
    IsUniqueWordOnly     bool // Deduplicate words
    IsNonEmptyWhitespace bool // Split by whitespace, ignore empty
    IsSortStringsBySpace bool // Sort words alphabetically
}
```

**Pre-built conditions:**

| Variable | Description |
|----------|-------------|
| `DefaultDisabledCoreCondition` | All flags `false` |
| `DefaultTrimCoreCondition` | Only `IsTrimCompare = true` |
| `DefaultSortTrimCoreCondition` | Trim + non-empty whitespace + sort |
| `DefaultUniqueWordsCoreCondition` | All flags `true` |

### TextValidator

Validates a single text against a search term using a configurable comparison method:

```go
type TextValidator struct {
    Search   string
    SearchAs stringcompareas.Variant // Equal, Contains, StartsWith, etc.
    Condition
}
```

| Method | Description |
|--------|-------------|
| `IsMatch(content, isCaseSensitive)` | Returns true if content matches search |
| `IsMatchMany(skipEmpty, caseSensitive, ...contents)` | Match against multiple strings |
| `VerifyDetailError(params, content)` | Returns detailed error on mismatch |
| `VerifySimpleError(index, params, content)` | Returns simple error on mismatch |
| `VerifyMany(continueOnErr, params, ...contents)` | Verify multiple contents |
| `VerifyFirstError(params, ...contents)` | Stop on first error |
| `AllVerifyError(params, ...contents)` | Collect all errors |
| `SearchTextFinalized()` | Get preprocessed search term |
| `MethodName()` | Comparison method name |
| `String()` / `ToString(singleLine)` | Human-readable representation |

### LineValidator

Extends `TextValidator` with line number matching:

```go
type LineValidator struct {
    LineNumber
    TextValidator
}
```

| Method | Description |
|--------|-------------|
| `IsMatch(lineNumber, content, caseSensitive)` | Match with line number check |
| `VerifyError(params, lineNumber, content)` | Error with line number context |
| `VerifyMany(continueOnErr, params, ...contentsWithLine)` | Multi-line verify |

### SliceValidator

Compares actual vs expected string slices line-by-line:

```go
type SliceValidator struct {
    Condition
    CompareAs                  stringcompareas.Variant
    ActualLines, ExpectedLines []string
}
```

| Method | Description |
|--------|-------------|
| `IsValid(caseSensitive)` | Full slice comparison |
| `IsValidOtherLines(caseSensitive, lines)` | Compare against other lines |
| `VerifyFirstError(params)` | Stop on first mismatch |
| `AllVerifyError(params)` | Collect all mismatches |
| `AllVerifyErrorQuick(index, header, ...actual)` | Quick verify with defaults |
| `AllVerifyErrorTestCase(index, header, caseSensitive)` | Test-oriented verify |
| `AssertAllQuick(t, index, header, ...actual)` | GoConvey assertion |
| `SetActual(lines)` | Set actual lines |
| `SetActualVsExpected(actual, expected)` | Set both |
| `ActualLinesString()` / `ExpectingLinesString()` | Quoted display |

### Parameter

Configuration bag for verification calls:

```go
type Parameter struct {
    CaseIndex                  int
    Header                     string
    IsCaseSensitive            bool
    IsSkipCompareOnActualEmpty bool
    IsAttachUserInputs         bool
}
```

## Constructors

```go
// From error output
sv := corevalidator.NewSliceValidatorUsingErr(
    err,
    expectedContent,
    true,  // trim
    true,  // non-empty whitespace
    false, // no sort
    stringcompareas.Contains,
)

// From any value
sv := corevalidator.NewSliceValidatorUsingAny(
    anyValue,
    expectedContent,
    true, true, false,
    stringcompareas.Equal,
)
```

## Usage

### Simple Text Validation

```go
validator := corevalidator.TextValidator{
    Search:    "expected output",
    SearchAs:  stringcompareas.Contains,
    Condition: corevalidator.DefaultTrimCoreCondition,
}

isMatch := validator.IsMatch("some expected output here", true)
// true
```

### Slice Comparison in Tests

```go
sv := &corevalidator.SliceValidator{
    ActualLines:   actualOutput,
    ExpectedLines: expectedLines,
    CompareAs:     stringcompareas.Equal,
    Condition:     corevalidator.DefaultTrimCoreCondition,
}

err := sv.AllVerifyErrorQuick(caseIndex, "test header", actualOutput...)
if err != nil {
    t.Fatal(err)
}
```

### Line-Level Validation

```go
lv := &corevalidator.LineValidator{
    LineNumber:    corevalidator.LineNumber{LineNumber: 3},
    TextValidator: corevalidator.TextValidator{
        Search:   "func main",
        SearchAs: stringcompareas.Contains,
        Condition: corevalidator.DefaultTrimCoreCondition,
    },
}

err := lv.VerifyError(params, 3, "func main() {")
```

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [Testing Guidelines](/spec/01-app/16-testing-guidelines.md)
- [errcore README](/errcore/README.md)
- [coretests README](/coretests/README.md)
