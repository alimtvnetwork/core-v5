# defaultcapacity — Slice/Map Capacity Helpers

Package `defaultcapacity` provides functions for computing predictive initial capacities for slices and maps, reducing allocations in hot paths.

## API

| Function | Description |
|----------|-------------|
| `Predictive(possibleLen, multiplier, additionalCap)` | Computes `ceil(possibleLen × multiplier) + additionalCap` |
| `PredictiveDefault(possibleLen)` | Shorthand with default multiplier and additional capacity |
| `PredictiveDefaultSmall(possibleLen)` | Smaller default for low-growth scenarios |
| `PredictiveFiftyPercentIncrement(possibleLen)` | 1.5× growth strategy |
| `MaxLimit(possibleLen)` | Caps capacity at a configured maximum |
| `OfSearch(length)` | Capacity hint for search-related buffers |
| `OfSplits(length)` | Capacity hint for string-split results |

## Usage

```go
import "github.com/alimtvnetwork/core/defaultcapacity"

// Pre-allocate a slice with predictive capacity
cap := defaultcapacity.PredictiveDefault(len(input))
result := make([]string, 0, cap)
```

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
