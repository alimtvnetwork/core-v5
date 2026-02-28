# Repository Overview

## Purpose

`gitlab.com/auk-go/core` is a shared Go utility framework — the foundational package of the **auk-go** ecosystem. It provides reusable primitives, data structures, interfaces, converters, validators, file-system helpers, testing utilities, and code-generation tools that keep downstream packages DRY and consistent.

## Who It Is For

- **Go developers** building applications within the auk-go ecosystem.
- **AI agents** that need to read, extend, or generate code following this framework's conventions.
- **Contributors** adding new utilities, data types, or platform-specific helpers.

## High-Level Architecture

```
core.go                  ← root package: slice/map factory helpers
├── chmodhelper/          ← file-system chmod parsing, verification, application
├── cmd/                  ← CLI entry-points (main, server, client, sample)
├── codegen/              ← unit-test code generation (DEPRECATED — see 10-codegen-deprecation-plan.md)
├── codestack/            ← call-stack capture & trace formatting
├── conditional/          ← generic ternary helpers (Bool, Int, String, …)
├── constants/            ← global constants, line separators per OS, capacity defaults
├── converters/           ← type converters (strings ↔ bytes, maps, pointers)
├── coreappend/           ← append utilities
├── corecmp / corecmparator / anycmp ← comparison helpers
├── corecsv/              ← CSV formatting & serialization
├── coredata/             ← rich data structures (corestr, corejson, coredynamic, coreonce, …)
├── corefuncs/            ← function-level utilities
├── coreimpl/             ← concrete implementations (e.g. enumimpl)
├── coreindexes/          ← index-to-name mapping
├── coreinstruction/      ← instruction abstractions
├── coreinterface/        ← canonical interface contracts (getters, checkers, serializers, …)
├── coremath/             ← min/max for all numeric types
├── coresort/             ← sorting utilities (Quick, QuickDsc)
├── coretaskinfo/         ← task metadata
├── coretests/            ← test helpers, assertions, Gherkins-style case wrappers
├── coreunique/           ← unique-value generators
├── coreutils/            ← string utilities, template replacers
├── corevalidator/        ← line / slice / text / range validators
├── coreversion/          ← semantic versioning data type
├── defaultcapacity/      ← default capacity constants
├── defaulterr/           ← default error factories
├── dtformats/            ← date-time format strings
├── enums/                ← enum helpers (stringcompareas, versionindexes)
├── errcore/              ← rich error construction, stack traces, merge, formatting
├── extensionsconst/      ← file-extension constants
├── filemode/             ← file-mode constants
├── internal/             ← private helpers (reflect, convert, csv, json, path, …)
├── isany/                ← nil/zero/type checks on interface{}
├── iserror/              ← error-defined checks
├── issetter/             ← 4-valued boolean (Uninitialized/True/False/Wildcard)
├── keymk/                ← key-maker utilities
├── mutexbykey/           ← per-key mutex locking
├── namevalue/            ← name-value pair utilities
├── osconsts/             ← OS-specific constants
├── ostype/               ← OS type detection
├── pagingutil/           ← paging calculations
├── refeflectcore/        ← reflection core helpers
├── regconsts/            ← regex constant strings
├── regexnew/             ← lazy-compiled regex with lock support
├── reqtype/              ← request-type helpers
├── scripts/              ← deployment & Docker scripts
├── simplewrap/           ← string wrapping (quotes, brackets, curly)
├── testconsts/           ← test-only constants
├── tests/                ← integration & wrapper tests
└── typesconv/            ← additional type conversions
```

## How to Build, Test, and Run

### Prerequisites

| Tool | Version |
|------|---------|
| Go   | **1.22+** |
| Git  | ≥ 2.29 |

### Common Commands

```bash
# Run default CLI
make                        # or: make run-main

# Run tests (integration)
make run-tests              # runs tests/ folder

# Build binary
make build                  # outputs to build/cli

# Run specific cmd targets
make run-server
make run-client
make run-sample
```

### Installation as Dependency

```bash
go get gitlab.com/auk-go/core
```

## How Specs Relate to README and cmd/README

| Document | Purpose |
|----------|---------|
| `/README.md` | Public-facing onboarding, examples, and links |
| `/cmd/README.md` | CLI entry-point documentation |
| `/spec/01-app/` | Deep architecture specs for AI and contributor training |
| `/spec/13-app-issues/` | Tracked improvement issues by category |

## Related Docs

- [Folder Map](./01-folder-map.md)
- [Codegen Deprecation Plan](./10-codegen-deprecation-plan.md)
- [Go Modernization Plan](./11-go-modernization.md)
- [CMD Entrypoints](./12-cmd-entrypoints.md)
- [Testing Patterns](./13-testing-patterns.md)
- [Core Interface Conventions](./14-core-interface-conventions.md)
- [Code Review Report](./15-code-review-report.md)
