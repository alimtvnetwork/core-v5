# conditional

## Folder Purpose

Provides ternary-style helper functions for all Go primitive types, replacing the lack of a ternary operator in Go.

## Responsibilities

1. `Bool(condition, trueVal, falseVal)` — and equivalents for `Int`, `String`, `Byte`, `Interface`, etc.
2. Function-based conditionals (`BoolFunc`, `StringFunc`, `ErrorFunc`).
3. Nil-default helpers (`NilDefStr`, `NilDefInt`, etc.).
4. Void/error function execution based on conditions.

## Usage Example

```go
result := conditional.Int(true, 2, 7)   // returns 2
result := conditional.Int(false, 2, 7)  // returns 7
name := conditional.String(len(s) > 0, s, "default")
```

## Related Docs

- [Repo Overview](../00-repo-overview.md)
