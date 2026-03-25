# dtformats — Date-Time Format Layouts

Package `dtformats` provides a strongly-typed `Layout` type wrapping Go's standard `time` layout strings, with classification helpers.

## Type

```go
type Layout string
```

## Constants

All standard Go time layouts are available as typed constants:

| Constant | Layout String |
|----------|--------------|
| `ANSIC` | `Mon Jan _2 15:04:05 2006` |
| `UnixDate` | `Mon Jan _2 15:04:05 MST 2006` |
| `RFC3339` | `2006-01-02T15:04:05Z07:00` |
| `RFC3339Nano` | `2006-01-02T15:04:05.999999999Z07:00` |
| `Kitchen` | `3:04PM` |
| ... | All standard `time` package formats |

## Methods

| Method | Description |
|--------|-------------|
| `Value() string` | Returns the raw layout string |
| `Is(format string) bool` | Compares against a format string |
| `IsTimeOnly() bool` | True for `Kitchen` |
| `IsTimeFocused() bool` | True for `Kitchen`, `Stamp*` formats |
| `IsDateTime() bool` | True for full date-time formats |
| `HasTimeZone() bool` | True for formats with timezone info |

## Usage

```go
import "github.com/alimtvnetwork/core/dtformats"

layout := dtformats.RFC3339
fmt.Println(layout.IsDateTime())   // true
fmt.Println(layout.HasTimeZone())  // true

t := time.Now().Format(layout.Value())
```

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
