# iserror — Error State Checks

Package `iserror` provides predicate functions for checking error states, supporting both single errors and variadic error lists.

## API

### Single Error

| Function | Description |
|----------|-------------|
| `Defined(err) bool` | `err != nil` |
| `Empty(err) bool` | `err == nil` |
| `NotEmpty(err) bool` | Alias for `Defined` |
| `Equal(a, b) bool` | Compares two errors |
| `NotEqual(a, b) bool` | Negation of `Equal` |
| `EqualString(err, s) bool` | Compares `err.Error()` to string |
| `NotEqualString(err, s) bool` | Negation of `EqualString` |
| `ExitError(err) (int, bool)` | Extracts exit code from `exec.ExitError` |

### Variadic (Multiple Errors)

| Function | Description |
|----------|-------------|
| `AllDefined(errs...) bool` | All non-nil (empty → false) |
| `AllEmpty(errs...) bool` | All nil (empty → true) |
| `AnyDefined(errs...) bool` | At least one non-nil (empty → false) |
| `AnyEmpty(errs...) bool` | At least one nil (empty → true) |

## Usage

```go
import "github.com/alimtvnetwork/core/iserror"

if iserror.Defined(err) {
    log.Fatal(err)
}

if iserror.AllEmpty(err1, err2, err3) {
    fmt.Println("all operations succeeded")
}
```

## Related Docs

- [isany README](/isany/README.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
