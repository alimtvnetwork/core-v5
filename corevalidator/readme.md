# corevalidator — Text, Line & Slice Validators

Package `corevalidator` provides composable validators for text matching, line-level comparison, and slice-level verification. Used extensively by the test framework (`coretests`) and assertion pipelines to validate multi-line output against expected patterns with configurable comparison methods, trimming, whitespace normalization, and sorting.

## Architecture

```
corevalidator/
├── Condition.go                   # Condition struct (trim, unique, sort flags)
├── Parameter.go                   # Verification parameter bag
├── LineNumber.go                  # Line number matching type (-1 = skip)
│
├── TextValidator.go               # Single text match with configurable comparison
├── TextValidators.go              # Collection of TextValidator items
│
├── LineValidator.go               # Text match + line number verification
├── LinesValidators.go             # Collection of LineValidator items
├── BaseLinesValidators.go         # Embeddable base for line collections
│
├── SliceValidator.go              # Compare actual vs expected string slices (~630 lines)
├── SliceValidators.go             # Collection of SliceValidator items
├── SimpleSliceValidator.go        # Simplified slice validator wrapper
├── HeaderSliceValidator.go        # Slice validator with header metadata
├── HeaderSliceValidators.go       # Collection of HeaderSliceValidator items
│
├── RangeSegmentsValidator.go      # Range-based segment validation
├── RangesSegment.go               # Range segment type definition
│
├── BaseValidatorCoreCondition.go  # Lazy-init condition base for embedding
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
    IsNonEmptyWhitespace bool // Split by whitespace, ignore empty segments
    IsSortStringsBySpace bool // Sort words alphabetically
}
```

**Pre-built conditions:**

| Variable | Trim | Unique | NonEmpty | Sort | Use Case |
|----------|------|--------|----------|------|----------|
| `DefaultDisabledCoreCondition` | ✗ | ✗ | ✗ | ✗ | Exact raw comparison |
| `DefaultTrimCoreCondition` | ✓ | ✗ | ✗ | ✗ | Ignore leading/trailing whitespace |
| `DefaultSortTrimCoreCondition` | ✓ | ✗ | ✓ | ✓ | Whitespace-normalized sorted comparison |
| `DefaultUniqueWordsCoreCondition` | ✓ | ✓ | ✓ | ✓ | Deduplicated, sorted word comparison |

### TextValidator

Validates a single text against a search term using a configurable comparison method from `stringcompareas`:

```go
type TextValidator struct {
    Search   string
    SearchAs stringcompareas.Variant // Equal, Contains, StartsWith, EndsWith, Regex, etc.
    Condition
}
```

**Supported comparison methods** (`stringcompareas.Variant`):

| Positive | Negative | Description |
|----------|----------|-------------|
| `Equal` | `NotEqual` | Exact string match |
| `StartsWith` | `NotStartsWith` | Prefix match |
| `EndsWith` | `NotEndsWith` | Suffix match |
| `Contains` / `Anywhere` | `NotContains` | Substring match |
| `AnyChars` | `NotAnyChars` | Any character present |
| `Regex` | `NotMatchRegex` | Regular expression match (cached) |

| Method | Description |
|--------|-------------|
| `IsMatch(content, isCaseSensitive)` | Returns true if content matches search |
| `IsMatchMany(skipEmpty, caseSensitive, ...contents)` | Match against multiple strings |
| `VerifyDetailError(params, content)` | Returns detailed error on mismatch |
| `VerifySimpleError(index, params, content)` | Returns simple error on mismatch |
| `VerifyMany(continueOnErr, params, ...contents)` | Verify multiple contents |
| `VerifyFirstError(params, ...contents)` | Stop on first error |
| `AllVerifyError(params, ...contents)` | Collect all errors |
| `SearchTextFinalized()` | Get preprocessed search term (cached) |
| `GetCompiledTermBasedOnConditions(input, caseSensitive)` | Apply conditions to any string |
| `MethodName()` | Comparison method name |
| `String()` / `ToString(singleLine)` | Human-readable representation |

### LineValidator

Extends `TextValidator` with line number matching. Line number `-1` means "skip line number check":

```go
type LineValidator struct {
    LineNumber
    TextValidator
}
```

| Method | Description |
|--------|-------------|
| `IsMatch(lineNumber, content, caseSensitive)` | Match with line number check |
| `IsMatchMany(skipEmpty, caseSensitive, ...contentsWithLine)` | Multi-item match |
| `VerifyError(params, lineNumber, content)` | Error with line number context |
| `VerifyMany(continueOnErr, params, ...contentsWithLine)` | Routes to First/All |
| `VerifyFirstError(params, ...contentsWithLine)` | Stop on first error |
| `AllVerifyError(params, ...contentsWithLine)` | Collect all errors |

### SliceValidator

Compares actual vs expected string slices line-by-line. Lazily creates `TextValidators` from `ExpectedLines` on first use:

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
| `AllVerifyErrorExceptLast(params)` | Verify up to second-last item |
| `AllVerifyErrorUptoLength(isFirstOnly, params, length)` | Verify up to N items |
| `AllVerifyErrorQuick(index, header, ...actual)` | Quick verify with defaults |
| `AllVerifyErrorTestCase(index, header, caseSensitive)` | Test-oriented verify (prints on error) |
| `AssertAllQuick(t, index, header, ...actual)` | GoConvey assertion |
| `NewSliceValidatorUsingErr(err, expected, ...)` | Construct from error output |
| `NewSliceValidatorUsingAny(any, expected, ...)` | Construct from any value |
| `SetActual(lines)` / `SetActualVsExpected(actual, expected)` | Mutate data |
| `ComparingValidators()` | Get/create cached TextValidators |
| `IsUsedAlready()` | Check if validators have been created |
| `UserInputsMergeWithError(params, err)` | Merge error with user input context |
| `Dispose()` | Release all internal references |

### Parameter

Configuration bag for verification calls:

```go
type Parameter struct {
    CaseIndex                  int    // Test case index for error reporting
    Header                     string // Test case title for error reporting
    IsCaseSensitive            bool   // Enable case-sensitive comparison
    IsSkipCompareOnActualEmpty bool   // Skip if actual is empty/nil
    IsAttachUserInputs         bool   // Append actual/expected to error messages
}
```

### Collections

| Collection | Item Type | Key Methods |
|------------|-----------|-------------|
| `TextValidators` | `TextValidator` | `Add`, `Adds`, `AddSimple`, `AddSimpleAllTrue`, `IsMatch`, `IsMatchMany`, `VerifyFirstError`, `AllVerifyError`, `VerifyErrorMany`, `Dispose` |
| `LinesValidators` | `LineValidator` | `Add`, `Adds`, `AddPtr`, `IsMatch`, `IsMatchText`, `VerifyFirstDefaultLineNumberError`, `AllVerifyError` |
| `SliceValidators` | `SliceValidator` | `IsMatch`, `IsValid`, `VerifyAll`, `VerifyFirst`, `VerifyUpto`, `SetActualOnAll` |
| `HeaderSliceValidators` | `HeaderSliceValidator` | `IsMatch`, `VerifyAll`, `VerifyFirst`, `VerifyUpto`, `AssertVerifyAll`, `VerifyAllErrorUsingActual` |

### Supporting Types

| Type | Description |
|------|-------------|
| `LineNumber` | Wraps `int`; `-1` means "no line number check". Methods: `HasLineNumber()`, `IsMatch(n)`, `VerifyError(n)` |
| `BaseValidatorCoreCondition` | Embeddable lazy-init `*Condition` pointer with `ValidatorCoreConditionDefault()` |
| `BaseLinesValidators` | Embeddable `[]LineValidator` with helper methods and `ToLinesValidators()` |
| `SimpleSliceValidator` | Wraps `SimpleSlice` expected + conditions, produces `SliceValidator` on demand |
| `RangeSegmentsValidator` | Validates specific line ranges within a larger output using `RangesSegment` definitions |
| `RangesSegment` | Combines `RangeInt` + `ExpectedLines` + `CompareAs` + `Condition` |

## Usage Examples

### Simple Text Validation

```go
validator := corevalidator.TextValidator{
    Search:    "expected output",
    SearchAs:  stringcompareas.Contains,
    Condition: corevalidator.DefaultTrimCoreCondition,
}

isMatch := validator.IsMatch("  some expected output here  ", true)
// true — trims content, then checks "contains"
```

### Case-Insensitive Match

```go
v := corevalidator.TextValidator{
    Search:    "Hello World",
    SearchAs:  stringcompareas.Equal,
    Condition: corevalidator.DefaultDisabledCoreCondition,
}

v.IsMatch("hello world", false) // true — case-insensitive
v.IsMatch("hello world", true)  // false — case-sensitive
```

### Regex Matching

```go
v := corevalidator.TextValidator{
    Search:   `^func\s+\w+\(`,
    SearchAs: stringcompareas.Regex,
}

v.IsMatch("func main() {", true) // true
v.IsMatch("var x = 1", true)     // false
```

### Unique Word Comparison

```go
v := corevalidator.TextValidator{
    Search:    "  banana  apple  apple  cherry  ",
    SearchAs:  stringcompareas.Equal,
    Condition: corevalidator.DefaultUniqueWordsCoreCondition,
}

// Preprocesses both sides: deduplicate → sort → "apple banana cherry"
v.IsMatch("cherry banana apple", false) // true
```

### Slice Comparison in Tests

```go
sv := &corevalidator.SliceValidator{
    ActualLines:   []string{"line1", "line2", "line3"},
    ExpectedLines: []string{"line1", "line2", "line3"},
    CompareAs:     stringcompareas.Equal,
    Condition:     corevalidator.DefaultTrimCoreCondition,
}

params := &corevalidator.Parameter{
    CaseIndex:          0,
    Header:             "output comparison",
    IsCaseSensitive:    true,
    IsAttachUserInputs: true,
}

err := sv.AllVerifyError(params)
// nil — all lines match
```

### Quick Slice Assertion

```go
sv := corevalidator.SliceValidator{
    CompareAs:     stringcompareas.Equal,
    Condition:     corevalidator.DefaultDisabledCoreCondition,
    ExpectedLines: []string{"a", "b", "c"},
}

sv.AssertAllQuick(t, 0, "my test case", "a", "b", "c")
```

### Line-Level Validation

```go
lv := &corevalidator.LineValidator{
    LineNumber: corevalidator.LineNumber{LineNumber: 3},
    TextValidator: corevalidator.TextValidator{
        Search:    "func main",
        SearchAs:  stringcompareas.Contains,
        Condition: corevalidator.DefaultTrimCoreCondition,
    },
}

err := lv.VerifyError(params, 3, "  func main() {  ")
// nil — line 3 matches, "func main" found in trimmed content
```

### Multi-Validator Pipeline

```go
validators := corevalidator.NewTextValidators(3)
validators.AddSimple("package main", stringcompareas.Equal)
validators.AddSimple("import", stringcompareas.Contains)
validators.AddSimple("func", stringcompareas.StartsWith)

// All validators must pass for a single content string
isMatch := validators.IsMatch("package main", true)
// false — "import" not found, "func" doesn't start
```

### Range Segment Validation

```go
rsv := &corevalidator.RangeSegmentsValidator{
    Title: "validate header and footer",
    VerifierSegments: []corevalidator.RangesSegment{
        {
            RangeInt:      corerange.RangeInt{Start: 0, End: 2},
            ExpectedLines: []string{"header1", "header2"},
            CompareAs:     stringcompareas.Equal,
            Condition:     corevalidator.DefaultTrimCoreCondition,
        },
    },
}

err := rsv.VerifySimple(actualLines, params, false)
```

### Error Construction from Test Output

```go
// Build validator from error output for comparison
sv := corevalidator.NewSliceValidatorUsingErr(
    functionUnderTest(),          // error result
    "expected line 1\nexpected line 2",
    true,  // trim
    false, // no whitespace split
    false, // no sort
    stringcompareas.Contains,
)

err := sv.AllVerifyError(params)
```

## Test Coverage

Comprehensive tests in `tests/integratedtests/corevalidatortests/` (~208 tests):

| File | Coverage |
|------|----------|
| `Parameter_test.go` | `IsIgnoreCase`, default values |
| `Condition_test.go` | `IsSplitByWhitespace` flag combinations, preset conditions |
| `LineNumber_test.go` | `HasLineNumber`, `IsMatch` with -1 skip, `VerifyError` |
| `TextValidator_test.go` | `IsMatch` (Equal/Contains/NotEqual/Trim/UniqueWords), `IsMatchMany`, `VerifyDetailError`, `VerifyMany`, caching |
| `TextValidators_test.go` | Collection ops, `IsMatch`, `VerifyFirstError`, `AllVerifyError`, `Dispose`, nil receiver |
| `LineValidator_test.go` | Combined line+text matching, `VerifyError`, `VerifyMany` |
| `LinesValidators_test.go` | `IsMatch` with contents, `VerifyFirstDefaultLineNumberError`, `AllVerifyError`, boundary scenarios |
| `SliceValidatorUnit_test.go` | `IsValid`, `Verify*`, `Dispose`, case sensitivity, `NewSliceValidatorUsingErr`, helpers |
| `SliceValidatorExtra_test.go` | `AllVerifyErrorExceptLast`, `AllVerifyErrorQuick`, caching, `UserInputsMergeWithError`, `ToString` |
| `SliceValidators_test.go` | Collection `IsMatch`, `VerifyAll`, `VerifyFirst`, `VerifyUpto`, nil receiver |
| `SimpleSliceValidator_test.go` | `VerifyAll`, `VerifyFirst`, `VerifyUpto` |
| `BaseLinesValidators_test.go` | `BaseLinesValidators`, `LinesValidators`, `BaseValidatorCoreCondition` |
| `TestValidators_test.go` | Table-driven integration tests |

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [Testing Guidelines](/spec/01-app/16-testing-guidelines.md)
- [errcore readme](/errcore/readme.md)
- [coretests readme](/coretests/readme.md)
- [stringcompareas Enum](/enums/stringcompareas/)
- [Improvement Plan](/spec/13-app-issues/corevalidator/01-improvement-plan.md)
