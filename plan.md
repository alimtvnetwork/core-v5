# Phase-by-Phase Plan

## Status Overview

| Phase | Status | Description |
|-------|--------|-------------|
| Phase 1 | ✅ Done | Repository scan and folder map draft |
| Phase 2 | ✅ Done | Per-folder spec docs |
| Phase 3 | ✅ Done | README upgrades (root README rewrite) |
| Phase 4 | ✅ Done | Special module docs and conventions |
| Phase 5 | ✅ Done | Codegen deprecation and Go modernization plan |
| Phase 6 | ✅ Done | Code review report and improvement backlog |
| Phase 7 | ✅ Done | Package-level README docs |
| Phase 8 | ✅ Done | Code review fixes (return types, typos, README accuracy) |

---

## Phase 1: Repository Scan and Folder Map Draft ✅

**Goal**: Understand the full repo structure and create foundational spec docs.

**Outputs**:
- `/spec/01-app/00-repo-overview.md` ✅
- `/spec/01-app/01-folder-map.md` ✅

**Acceptance Criteria**:
- [x] Folder map exists and covers all top-level folders
- [x] Repo overview exists and links to folder map
- [x] Missing doc areas are listed

---

## Phase 2: Per-Folder Spec Docs ✅

**Goal**: Create spec docs for each major folder.

**Outputs**:
- `/spec/01-app/folders/01-chmodhelper.md` ✅
- `/spec/01-app/folders/02-cmd.md` ✅
- `/spec/01-app/folders/03-codegen.md` ✅
- `/spec/01-app/folders/04-coreinterface.md` ✅
- `/spec/01-app/folders/05-coredata.md` ✅
- `/spec/01-app/folders/06-errcore.md` ✅
- `/spec/01-app/folders/07-coretests.md` ✅
- `/spec/01-app/folders/08-conditional.md` ✅
- `/spec/01-app/folders/09-internal.md` ✅
- `/spec/01-app/folders/10-remaining-packages.md` ✅

**Acceptance Criteria**:
- [x] Each major folder has a spec doc
- [x] Each doc links back to repo overview

---

## Phase 3: README Upgrades ✅

**Goal**: Rewrite root README with quick start, examples, and spec links.

**Acceptance Criteria**:
- [x] README has end-to-end onboarding
- [x] README includes modern examples
- [x] README points to spec docs
- [x] Prerequisites updated to modern Go

---

## Phase 4: Special Module Docs ✅

**Goal**: Document key modules and conventions.

**Outputs**:
- `/spec/01-app/modules/01-chmod-helper.md` ✅
- `/spec/01-app/12-cmd-entrypoints.md` ✅
- `/spec/01-app/13-testing-patterns.md` ✅
- `/spec/01-app/14-core-interface-conventions.md` ✅
- Updated `/cmd/README.md` ✅

**Acceptance Criteria**:
- [x] Module docs exist and are consistent
- [x] Conventions are explicit and reusable

---

## Phase 5: Codegen Deprecation and Go Modernization ✅

**Outputs**:
- `/spec/01-app/10-codegen-deprecation-plan.md` ✅
- `/spec/01-app/11-go-modernization.md` ✅

**Acceptance Criteria**:
- [x] Codegen removal plan has clear exit criteria
- [x] Go modernization plan is actionable

---

## Phase 6: Code Review and Issues ✅

**Outputs**:
- `/spec/01-app/15-code-review-report.md` ✅
- `/spec/13-app-issues/golang/01-convertinteranl-typo.md` ✅
- `/spec/13-app-issues/golang/02-refeflectcore-typo.md` ✅
- `/spec/13-app-issues/golang/03-go-version-outdated.md` ✅
- `/spec/13-app-issues/golang/04-type-duplication-no-generics.md` ✅
- `/spec/13-app-issues/docs/01-readme-outdated.md` ✅
- `/spec/13-app-issues/docs/02-missing-package-docs.md` ✅
- `/spec/13-app-issues/codegen/01-codegen-deprecation.md` ✅
- `/spec/13-app-issues/testing/01-missing-unit-tests.md` ✅
- `/spec/13-app-issues/cmd/01-cmd-readme-minimal.md` ✅

---

## Phase 7: Package-Level README Docs ✅

**Goal**: Create comprehensive README.md files for core packages with usage examples.

**Outputs**:
- `coredata/corejson/README.md` ✅
- `coredata/coreapi/README.md` ✅
- `corefuncs/README.md` ✅
- Verified `coredata/coredynamic/README.md` ✅ (pre-existing)
- Verified `coredata/corepayload/README.md` ✅ (pre-existing)

**Acceptance Criteria**:
- [x] Each README has accurate folder tree
- [x] Each README has usage examples
- [x] All method names and return types verified against source

---

## Phase 8: Code Review Fixes ✅

**Goal**: Fix issues found during README review and code modernization.

**Outputs**:
- Migrated `*[]string` → `[]string` return types in Hashset, Collection, CharHashsetMap ✅
- Fixed README examples: `Serialize.ToString` signature, `MapResults` API, wrapper constructors ✅
- Added missing files to folder trees in corejson and coreapi READMEs ✅
- Added deprecation notices to `*Ptr()` methods with direct-return alternatives ✅

**Acceptance Criteria**:
- [x] All README examples compile-correct
- [x] Folder trees match actual directory contents
- [x] Return type modernization complete for string collections

---

## Prioritized Backlog

| Priority | Task | Spec Reference |
|----------|------|---------------|
| ✅ Done | `interface{}` → `any` in `coreinterface/` (verified 0 remaining) | `20-improvement-plan.md` |
| ✅ Done | Split `Attributes.go` — already split (Getters/Setters/Json) | `20-improvement-plan.md` Phase 5 |
| ✅ Done | Split `Dynamic.go` — already split (Getters/Reflect/Json) | `20-improvement-plan.md` Phase 5 |
| ✅ Done | Split `Info.go` (646→4 files: Info/Getters/Json/Map) | `20-improvement-plan.md` Phase 5 |
| 🟡 Medium | Split `BaseTestCase.go` (435 lines) | `20-improvement-plan.md` Phase 5 |
| 🟡 Medium | Add generics to `conditional/` | `11-go-modernization.md` |
| 🟡 Medium | Deprecate codegen | `10-codegen-deprecation-plan.md` |
| 🟡 Medium | Create READMEs for coregeneric, corestr, coreonce, corerange, stringslice | `13-app-issues/docs/02-missing-package-docs.md` |
| 🟡 Medium | Value receiver migration (ongoing) | `20-improvement-plan.md` Phase 6 |
| 🟢 Low | Add missing unit tests (P1-P3 packages) | `20-improvement-plan.md` Phase 4 |
| 🟢 Low | Remove codegen completely | `10-codegen-deprecation-plan.md` |

## Next Task Selection

Recommended order:

1. **`coreinterface/` `any` migration** — Largest remaining modernization surface (569 matches).
2. **File splitting** — Reduces cognitive load on large files.
3. **Remaining package READMEs** — Completes documentation coverage.
4. **Codegen deprecation** — Simplifies maintenance.
