# simplewrap — String Wrapping Utilities

## Overview

Package `simplewrap` provides functions for wrapping strings with delimiters — double quotes, single quotes, curly braces, square brackets, parentheses, and arbitrary start/end characters. Includes conditional wrapping, title-prefixed formats, and JSON name formatting.

## Architecture

```
simplewrap/
├── WrapWith.go                          # With(start, source, end), WithPtr
├── WrapWithStartEnd.go                  # WrapWithStartEnd
├── ConditionalWrapWith.go               # ConditionalWrapWith — only adds missing delimiters
├── WithDoubleQuote.go                   # WithDoubleQuote(string)
├── WithDoubleQuoteAny.go               # WithDoubleQuoteAny(any)
├── WithSingleQuote.go                   # WithSingleQuote(string)
├── WithParenthesis.go                   # WithParenthesis(string)
├── WithParenthesisQuotation.go          # WithParenthesisQuotation
├── WithCurly.go                         # WithCurly(string)
├── WithCurlyQuotation.go               # WithCurlyQuotation
├── WithBrackets.go                      # WithBrackets(string)
├── WithBracketsQuotation.go            # WithBracketsQuotation
├── CurlyWrap.go                         # CurlyWrap(any)
├── CurlyWrapIf.go                       # CurlyWrapIf(condition, any)
├── CurlyWrapOption.go                   # CurlyWrapOption
├── SquareWrap.go                        # SquareWrap(any)
├── SquareWrapIf.go                      # SquareWrapIf(condition, any)
├── ParenthesisWrap.go                   # ParenthesisWrap(any)
├── ParenthesisWrapIf.go                 # ParenthesisWrapIf(condition, any)
├── DoubleQuoteWrapElements.go           # DoubleQuoteWrapElements([]string)
├── DoubleQuoteWrapElementsWithIndexes.go # DoubleQuoteWrapElementsWithIndexes
├── TitleCurlyWrap.go                    # TitleCurlyWrap(title, value)
├── TitleCurlyMeta.go                    # TitleCurlyMeta
├── TitleSquare.go                       # TitleSquare(title, value)
├── TitleSquareMeta.go                   # TitleSquareMeta
├── TitleSquareMetaUsingFmt.go           # TitleSquareMetaUsingFmt
├── TitleSquareCsvMeta.go                # TitleSquareCsvMeta
├── TitleQuotationMeta.go                # TitleQuotationMeta
├── MsgCsvItems.go                       # MsgCsvItems
├── MsgWrapMsg.go                        # MsgWrapMsg
├── MsgWrapNumber.go                     # MsgWrapNumber
├── ToJsonName.go                        # ToJsonName — JSON field name formatting
├── toString.go                          # Internal any-to-string conversion
├── wrapDoubleQuoteOnNonExist.go         # Wrap only if not already quoted
└── readme.md
```

## Key Functions

### Basic Wrapping

| Function | Output | Description |
|----------|--------|-------------|
| `With(start, source, end)` | `start + source + end` | Generic string concatenation |
| `WithPtr(*start, *source, *end)` | `*string` | Nil-safe pointer variant |
| `WithDoubleQuote(s)` | `"s"` | Double-quote wrap |
| `WithSingleQuote(s)` | `'s'` | Single-quote wrap |
| `WithParenthesis(s)` | `(s)` | Parenthesis wrap |
| `WithCurly(s)` | `{s}` | Curly brace wrap |
| `WithBrackets(s)` | `[s]` | Square bracket wrap |

### Any-Type Wrapping

| Function | Description |
|----------|-------------|
| `CurlyWrap(any)` | `{value}` — converts any to string first |
| `SquareWrap(any)` | `[value]` |
| `ParenthesisWrap(any)` | `(value)` |
| `WithDoubleQuoteAny(any)` | `"value"` |

### Conditional Wrapping

| Function | Description |
|----------|-------------|
| `ConditionalWrapWith(startChar, input, endChar)` | Only adds missing delimiters |
| `CurlyWrapIf(bool, any)` | Wrap only if condition true |
| `SquareWrapIf(bool, any)` | Wrap only if condition true |
| `ParenthesisWrapIf(bool, any)` | Wrap only if condition true |

### Title-Prefixed Formats

| Function | Description |
|----------|-------------|
| `TitleCurlyWrap(title, value)` | `"title {value}"` format |
| `TitleSquare(title, value)` | `"title [value]"` format |
| `TitleCurlyMeta(...)` | Title with curly metadata |
| `TitleSquareMeta(...)` | Title with square metadata |
| `TitleSquareCsvMeta(...)` | Title with CSV square metadata |
| `TitleQuotationMeta(...)` | Title with quoted metadata |

### Batch Operations

| Function | Description |
|----------|-------------|
| `DoubleQuoteWrapElements([]string)` | Wrap each element in double quotes |
| `DoubleQuoteWrapElementsWithIndexes(...)` | Wrap with index annotations |

### Message Formatting

| Function | Description |
|----------|-------------|
| `MsgCsvItems(...)` | Message with CSV items |
| `MsgWrapMsg(...)` | Nested message wrapping |
| `MsgWrapNumber(...)` | Message with numeric value |
| `ToJsonName(...)` | JSON field name formatting |

## Usage

```go
import "github.com/alimtvnetwork/core/simplewrap"

// Basic wrapping
simplewrap.WithDoubleQuote("hello")     // "hello"
simplewrap.WithCurly("key")            // {key}
simplewrap.WithBrackets("0")           // [0]

// Conditional — only adds if missing
simplewrap.ConditionalWrapWith('{', "already}", '}') // {already}

// Title format
simplewrap.TitleCurlyWrap("Error", "nil pointer") // Error {nil pointer}

// Any-type
simplewrap.CurlyWrap(42)               // {42}
```

## How to Extend Safely

- **New delimiter types**: Add as `With{Delimiter}.go` files following the existing pattern.
- **New conditional variants**: Add as `{Delimiter}WrapIf.go` files.
- **New title formats**: Add as `Title{Delimiter}*.go` files.
- **Do not** modify `toString.go` — it delegates to `convertinternal.AnyTo.SmartString`.

## Related Docs

- [Coding Guidelines](../spec/01-app/17-coding-guidelines.md)
