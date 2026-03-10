# mutexbykey — Per-Key Mutex Locking

## Overview

Package `mutexbykey` provides a global, thread-safe registry of `sync.Mutex` instances keyed by string. Use it when you need fine-grained locking per logical resource (e.g., per user ID, per file path) without creating mutex maps in each consumer.

## Architecture

```
mutexbykey/
├── mutexMap.go     # Get, Delete — global mutex registry
└── readme.md
```

## Functions

| Function | Signature | Description |
|----------|-----------|-------------|
| `Get` | `(key string) *sync.Mutex` | Returns the mutex for `key`, creating one if it doesn't exist |
| `Delete` | `(key string) bool` | Removes the mutex for `key`; returns `true` if it existed |

## Internal Design

- A package-level `globalMutex` protects all reads/writes to the internal map.
- Mutexes are created lazily on first `Get` call per key.
- Initial map capacity is `ArbitraryCapacity10` from `constants`.

## Usage

```go
import "github.com/alimtvnetwork/core/mutexbykey"

// Lock per resource
mu := mutexbykey.Get("user:123")
mu.Lock()
defer mu.Unlock()
// ... critical section for user 123

// Cleanup when resource is no longer needed
mutexbykey.Delete("user:123")
```

## Related Docs

- [Coding Guidelines](../spec/01-app/17-coding-guidelines.md)
