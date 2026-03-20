# Integrated Testing Guidelines for Go

> **Portable guideline** — drop this folder into any Go project that uses the `coretests` framework.
> It covers folder structure, test case types, input holders, result assertions, branch coverage strategy, and real-world examples.

## Table of Contents

| File | Purpose |
|------|---------|
| [01-folder-structure.md](01-folder-structure.md) | Directory layout, naming conventions, file separation rules |
| [02-test-case-types.md](02-test-case-types.md) | `CaseV1`, `CaseNilSafe`, `GenericGherkins` — when to use each |
| [03-args-reference.md](03-args-reference.md) | `args.Map`, `args.One`–`args.Six`, `args.Dynamic`, `args.Holder`, `args.LeftRight` |
| [04-results-reference.md](04-results-reference.md) | `results.Result`, `results.ResultAny`, `results.ExpectAnyError`, `InvokeWithPanicRecovery` |
| [05-assertion-patterns.md](05-assertion-patterns.md) | `ShouldBeEqual`, `ShouldBeEqualMap`, `ShouldBeSafe`, diff-based assertions |
| [06-branch-coverage.md](06-branch-coverage.md) | Positive/negative paths, if/else, nil guards, boundary conditions |
| [07-good-vs-bad.md](07-good-vs-bad.md) | Concrete examples of good tests vs. bad tests |
| [08-creating-custom-cases.md](08-creating-custom-cases.md) | How to create your own case type using `BaseTestCase` |

## Core Principles

1. **Separation of data and logic** — test cases (`_testcases.go`) never contain assertions; test runners (`_test.go`) never contain expected values
2. **AAA comments are mandatory** — every test function has `// Arrange`, `// Act`, `// Assert` comments
3. **No raw `t.Error` / `t.Errorf`** — always use framework assertions (`ShouldBeEqual`, `ShouldBeEqualMap`, etc.)
4. **Native types in expectations** — use `bool`, `int`, `string` in `args.Map`, not `"true"`, `"5"`
5. **One test function per logical scenario** — no branching logic inside test bodies
6. **No coverage tests for `internal/` packages** — never write coverage-motivated tests for any package under the `internal/` folder. Business-critical tests are allowed, but coverage push work must target only public packages.
