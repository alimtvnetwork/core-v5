# Architecture Decision: Module Splitting

## Status: 📋 DECISION — Keep Single Module (Monorepo)

## Date: 2026-03-21

## Context

The `github.com/alimtvnetwork/core` repository is a single Go module containing ~62 top-level packages, ~15 internal packages, and ~65 test packages. As the codebase has grown (generics adoption, new collection types, expanded test suites), the question arises: should this monorepo be split into multiple Go modules?

This document evaluates the trade-offs and recommends an approach.

---

## Current Architecture

### Module Stats

| Metric | Value |
|--------|-------|
| Top-level packages | ~62 |
| Internal packages | ~15 |
| Test packages (integratedtests) | ~65 |
| External dependencies | 3 (goconvey, assertions, x/tools) |
| Go version | 1.25.0 |

### Dependency Clusters (by import frequency)

| Cluster | Packages | Internal Imports |
|---------|----------|-----------------|
| **Core Foundation** | `constants`, `defaultcapacity`, `defaulterr`, `conditional` | 0–1 deps each |
| **Type Utilities** | `isany`, `issetter`, `iserror`, `converters`, `typesconv` | Foundation only |
| **String/Format** | `simplewrap`, `corecsv`, `coreutils/stringutil`, `regexnew` | Foundation + Type Utils |
| **Interfaces** | `coreinterface/*` (6 sub-packages) | Near-zero internal deps |
| **Error System** | `errcore` | Foundation + `internal/*` + `namevalue` |
| **Data Structures** | `coredata/*` (9 sub-packages) | Heavy: errcore, converters, internal/*, interfaces |
| **Comparison/Sort** | `anycmp`, `corecmp`, `corecomparator`, `coresort` | Foundation + constants |
| **Test Framework** | `coretests/*` | Heavy: errcore, coredata, validators, internal/* |
| **File System** | `chmodhelper/*` | Heavy: errcore, coredata, internal/*, interfaces |
| **Reflection** | `reflectcore/*`, `internal/reflectinternal` | Foundation |
| **Leaf Packages** | `dtformats`, `extensionsconst`, `filemode`, `pagingutil`, `regconsts`, `testconsts`, `coremath`, `ostype` | Zero internal imports |

### Dependency Heat Map (top importees in non-test code)

```
constants          — 427 imports (ubiquitous)
errcore            — 212 imports (error construction)
coredata/corestr   — 210 imports (string collections)
coredata/corejson  — 189 imports (JSON utilities)
internal/*         — ~300 imports combined
corecomparator     —  71 imports
isany              —  43 imports
issetter           —  45 imports
```

---

## Options Evaluated

### Option A: Keep Single Module (Monorepo) ✅ RECOMMENDED

Keep `github.com/alimtvnetwork/core` as one `go.mod`.

**Pros:**
- Zero versioning coordination overhead
- Atomic refactors across packages (rename, signature changes)
- Single CI pipeline, single coverage gate
- `internal/` packages remain naturally scoped
- No diamond dependency risk
- Current external dependency count is minimal (3 deps)
- All existing tooling (`run.ps1 PC`, `TC`) works unchanged

**Cons:**
- Consumers pull entire module even if they only need `constants`
- Single version number for unrelated changes

### Option B: Split into 3–4 Focused Modules

```
core-foundation/  → constants, conditional, isany, issetter, defaulterr, defaultcapacity, coremath, ...
core-data/        → coredata/*, converters, simplewrap, corecsv, regexnew, ...
core-errors/      → errcore, coreinterface/errcoreinf, codestack
core-testing/     → coretests/*, corevalidator, testconsts
```

**Pros:**
- Consumers can import only what they need
- Independent versioning per module
- Smaller dependency trees for downstream projects

**Cons:**
- **High coupling makes clean cuts difficult**: `errcore` imports `internal/reflectinternal` and `constants`; `coredata` imports `errcore`, `converters`, `internal/*`, and `coreinterface/*`. Every split boundary creates cross-module dependencies.
- **`internal/` packages cannot be shared across modules**: The 15 `internal/*` packages would need to be either duplicated, promoted to public API, or consolidated into one module — all undesirable.
- **Circular dependency risk**: `coredata` → `errcore` → `internal/reflectinternal` → used by `coredata/coredynamic`. Splitting requires careful layering.
- **Version coordination**: Any interface change in `core-foundation` forces synchronized releases across all downstream modules.
- **CI complexity**: Multiple pipelines, cross-module integration tests, release orchestration.
- **Breaking change for all consumers**: Every `import "github.com/alimtvnetwork/core/..."` path changes.

### Option C: Extract Only Leaf Packages

Extract zero-dependency packages as standalone modules:
```
github.com/alimtvnetwork/core-constants
github.com/alimtvnetwork/core-dtformats
github.com/alimtvnetwork/core-filemode
```

**Pros:**
- Low risk — no coupling to break
- Consumers of just constants get a tiny dependency

**Cons:**
- Marginal benefit — these packages are tiny (<10 files each)
- Maintenance overhead of separate repos for trivial code
- Fragments discoverability

---

## Decision

**Keep single module (Option A).**

### Rationale

1. **Coupling is too high for clean splitting.** The dependency analysis shows that `constants`, `errcore`, `internal/*`, and `coreinterface/*` are imported by nearly every package. There is no natural seam where a module boundary would not create cross-module imports.

2. **`internal/` is the primary blocker.** Go's `internal/` visibility rule means these 15 packages can only be used within the same module. Splitting would require promoting them to public API (increasing API surface and maintenance burden) or duplicating code.

3. **The monorepo is not large enough to justify splitting.** With only 3 external dependencies and ~62 packages, the module is well within the size range where a single `go.mod` is the standard Go practice (cf. `golang.org/x/tools`, `k8s.io/apimachinery`).

4. **Consumer impact is minimal.** All known consumers are within the `alimtvnetwork` organization, so the "pull entire module" cost is negligible (Go module cache + proxy handle this efficiently).

5. **Leaf package extraction (Option C) has poor ROI.** The zero-dependency packages are tiny and rarely imported independently.

### When to Revisit

Re-evaluate module splitting if any of these conditions change:

| Trigger | Action |
|---------|--------|
| External consumers outside the org need only a subset | Consider Option C for high-demand leaf packages |
| `internal/` packages need to be shared with a separate repo | Promote specific internals to a `pkg/` public package |
| Module exceeds ~150 packages | Re-evaluate Option B with updated dependency analysis |
| Go tooling adds "workspace-aware module splitting" | Re-evaluate with new tooling |

---

## Alternatives Considered but Rejected

| Alternative | Why Rejected |
|-------------|--------------|
| Go workspaces (`go.work`) | Workspaces are for development, not distribution. Consumers still import individual modules. |
| Build tags for optional packages | Go build tags don't reduce module download size and add complexity. |
| Interface-only module extraction | `coreinterface/` has near-zero deps but is too small (~20 files) to justify a separate module. |

---

## Related Docs

- [Folder Map](./01-folder-map.md)
- [Code Review Report](./15-code-review-report.md)
- [Coding Guidelines](./17-coding-guidelines.md)
- [Generic Interfaces Decision](./20-generic-interfaces-decision.md)
