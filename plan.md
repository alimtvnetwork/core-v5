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

## Phase 3: README Upgrades 🔲

**Goal**: Rewrite root README with quick start, examples, and spec links.

**Inputs**: Current README.md, spec docs
**Outputs**: Updated `/README.md`

**Acceptance Criteria**:
- [ ] README has end-to-end onboarding
- [ ] README includes modern examples
- [ ] README points to spec docs
- [ ] Prerequisites updated to modern Go

**Open Questions**:
- Should the README maintain backward-compatible examples for Go 1.17?

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

## Prioritized Backlog

| Priority | Task | Spec Reference |
|----------|------|---------------|
| 🔴 High | Upgrade Go to 1.22+ | `11-go-modernization.md` |
| 🔴 High | Replace `interface{}` → `any` project-wide | `11-go-modernization.md` |
| 🟡 Medium | Rewrite root README | `spec/13-app-issues/docs/01-readme-outdated.md` |
| 🟡 Medium | Fix `convertinteranl` typo | `spec/13-app-issues/golang/01-convertinteranl-typo.md` |
| 🟡 Medium | Fix `refeflectcore` typo | `spec/13-app-issues/golang/02-refeflectcore-typo.md` |
| 🟡 Medium | Add generics to `conditional/` | `11-go-modernization.md` |
| 🟡 Medium | Add generics to `core.go` | `11-go-modernization.md` |
| 🟡 Medium | Deprecate codegen | `10-codegen-deprecation-plan.md` |
| 🟢 Low | Add missing unit tests | `spec/13-app-issues/testing/01-missing-unit-tests.md` |
| 🟢 Low | Create per-package READMEs | `spec/13-app-issues/docs/02-missing-package-docs.md` |
| 🟢 Low | Remove codegen completely | `10-codegen-deprecation-plan.md` |

## Next Task Selection

Pick your next task from the backlog above. Recommended order:

1. **Start with Phase 3** — Rewrite root README (improves onboarding immediately).
2. **Then Go upgrade** — Unlocks all generics work.
3. **Then generics** — Reduces codebase size significantly.
4. **Then codegen removal** — Simplifies maintenance.
