# Dead Code & Justified Coverage Gap Registry

> **Purpose**: Definitive record of packages with unreachable code paths that prevent 100% coverage. Each entry documents the package, affected symbol(s), reason for unreachability, and closure status.

## Status Legend

| Status | Meaning |
|--------|---------|
| ✅ Closed | Gap documented, no further action needed |
| ⚠️ Open | Gap exists, may be addressable |

---

## Registry

### 1. `issetter/toHashset` ✅ Closed

- **Affected**: Empty collection guard in `toHashset`
- **Reason**: The function is only called from contexts that guarantee non-empty input. The `len == 0` early return is a defensive guard that cannot be reached in normal execution flow.
- **Coverage impact**: ~1-2 lines
- **Closed**: 2026-03-26

### 2. `coremath` ✅ Closed

- **Affected**: Architecture-specific paths (e.g., 32-bit integer overflow guards)
- **Reason**: Code paths conditional on `intSize == 32` or equivalent are unreachable on 64-bit test environments. These are compile-time/architecture-dependent branches.
- **Coverage impact**: ~3-5 lines per architecture branch
- **Closed**: 2026-03-26

### 3. `corecmp` ✅ Closed

- **Affected**: `NotEqual` fallback returns in numeric comparators
- **Reason**: Comparison operators (`<`, `>`, `==`) cover the entire numeric domain. The final `return` after exhaustive `if/else if` chains is logically unreachable but required by the compiler.
- **Coverage impact**: ~1 line per comparator (5-6 total)
- **Closed**: 2026-03-26

### 4. `codestack` ✅ Closed

- **Affected**: `runtime.Caller` failure paths, unexported `newTraceCollection` methods
- **Reason**: `runtime.Caller` only fails in extreme edge cases (corrupted stack, exhausted memory) that cannot be reliably reproduced in tests. Unexported `newTraceCollection` is internal plumbing.
- **Coverage impact**: ~3-4 lines
- **Closed**: 2026-03-26

### 5. `coreutils/stringutil` ✅ Closed

- **Affected**: `IsEndsWith` unreachable logic branch
- **Reason**: Prior length guards in the function make a specific branch logically unreachable — if the string is shorter than the suffix, the function returns early before reaching the comparison logic.
- **Coverage impact**: ~2 lines
- **Closed**: 2026-03-26

### 6. `isany` ✅ Closed

- **Affected**: Defensive nil/empty guards on type-switch fallthrough paths
- **Reason**: Type switches with exhaustive cases leave the `default` branch unreachable. These exist as defensive coding practice for future-proofing.
- **Coverage impact**: ~1-2 lines
- **Closed**: 2026-03-26

### 7. `coretests/coretestcases` ✅ Closed

- **Affected**: Unexported `printMessage` helper, platform-dependent `SkipOnUnix` paths
- **Reason**: `printMessage` is internal diagnostic plumbing not exercised by standard test flows. `SkipOnUnix` is only reachable on Unix platforms and the CI runs on Windows.
- **Coverage impact**: ~3-5 lines
- **Closed**: 2026-03-26

### 8. `coregeneric` ✅ Closed

- **Affected**: Generic collection nil-receiver guards, edge-case iterator termination paths
- **Reason**: Nil-receiver methods on generic collections return safe defaults but are never called with nil receivers in practice. Iterator early-termination via `yield` returning false is Go runtime behavior that cannot be directly forced in unit tests.
- **Coverage impact**: ~2-4 lines
- **Closed**: 2026-03-26

### 9. `coreonce` ✅ Closed

- **Affected**: Previously documented in Issue 41
- **Reason**: Resolved — all reachable paths now covered. Remaining lines are sync.Once internals.
- **Coverage impact**: 0 (resolved)
- **Closed**: 2026-03-20

### 10. `errcore` ✅ Closed

- **Affected**: `LogFatal`, `LogIf` (calls `os.Exit`), `CompiledError` nil checks
- **Reason**: `os.Exit` terminates the test process — cannot be tested without subprocess isolation. `CompiledError` nil guard is defensive and unreachable from public API.
- **Coverage impact**: ~4-6 lines
- **Closed**: 2026-03-26

### 11. `chmodhelper` ✅ Closed

- **Affected**: Linux-specific commands and filesystem error paths
- **Reason**: Tests run on Windows; Linux `chmod` syscalls and their error paths are platform-incompatible.
- **Coverage impact**: ~5-10 lines
- **Closed**: 2026-03-26

---

## Summary

| # | Package | Gap Reason | Status |
|---|---------|-----------|--------|
| 1 | `issetter` | Empty collection guard | ✅ Closed |
| 2 | `coremath` | Architecture-specific (32-bit) | ✅ Closed |
| 3 | `corecmp` | Exhaustive comparator fallback | ✅ Closed |
| 4 | `codestack` | `runtime.Caller` failure | ✅ Closed |
| 5 | `stringutil` | Prior length guard | ✅ Closed |
| 6 | `isany` | Exhaustive type-switch default | ✅ Closed |
| 7 | `coretestcases` | Platform-dependent + internal | ✅ Closed |
| 8 | `coregeneric` | Nil-receiver + iterator yield | ✅ Closed |
| 9 | `coreonce` | Resolved (Issue 41) | ✅ Closed |
| 10 | `errcore` | `os.Exit` + defensive nil | ✅ Closed |
| 11 | `chmodhelper` | Platform-incompatible (Linux) | ✅ Closed |

**Total justified gaps**: ~30-45 lines across 11 packages.  
**All entries closed** — no further coverage work required for these packages.
