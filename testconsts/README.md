# testconsts — Test-Only Constants

Package `testconsts` provides shared constant values and sample data used exclusively in test files across the project. **Not for production use.**

## Constants

| Constant | Value | Type |
|----------|-------|------|
| `JohnDoe` | `"John Doe"` | `string` |
| `JaneDoe` | `"Jane Doe"` | `string` |
| `John` | `"John"` | `string` |
| `Jane` | `"Jane"` | `string` |
| `Doe` | `"Doe"` | `string` |
| `NumberOne` | `1` | `int` |
| `NumberTwo` | `2` | `int` |
| `NumberThree` | `3` | `int` |
| `NumberFive` | `5` | `int` |

## Variables

| Variable | Type | Description |
|----------|------|-------------|
| `LowerCaseStringsArray` | `[]string` | `["hello", "world", "one", "two"]` |

## Usage

```go
import "github.com/alimtvnetwork/core/testconsts"

func Test_Something(t *testing.T) {
    input := testconsts.JohnDoe
    // ...
}
```

## Related Docs

- [Testing Patterns](/spec/01-app/13-testing-patterns.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
