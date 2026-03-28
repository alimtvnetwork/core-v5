# ─────────────────────────────────────────────────────────────────
# Makefile — Thin wrapper around run.sh for Unix conventions
# Usage: make <target> [ARGS="..."]
# ─────────────────────────────────────────────────────────────────
.PHONY: test test-pkg test-cover test-cover-pkg test-int test-fail \
        goconvey run build build-run fmt lint vet tidy pre-commit \
        clean help

ARGS ?=

# ── Testing ──────────────────────────────────────────────────────

test:                ## Run all tests (verbose)
	@bash run.sh t $(ARGS)

test-pkg:            ## Run tests for a specific package: make test-pkg ARGS="regexnewtests"
	@bash run.sh tp $(ARGS)

test-cover:          ## Run tests with coverage (HTML + summary)
	@bash run.sh tc $(ARGS)

test-cover-pkg:      ## Run coverage for a specific package: make test-cover-pkg ARGS="regexnewtests"
	@bash run.sh tcp $(ARGS)

test-int:            ## Run integrated tests only
	@bash run.sh ti $(ARGS)

test-fail:           ## Show last failing tests log
	@bash run.sh tf

goconvey:            ## Launch GoConvey browser test runner
	@bash run.sh gc $(ARGS)

# ── Shortcuts (match run.ps1 shorthand names) ────────────────────

t:  test
tp: test-pkg
tc: test-cover
tcp: test-cover-pkg
ti: test-int
tf: test-fail
gc: goconvey

# ── Build & Run ──────────────────────────────────────────────────

run:                 ## Run the main application
	@bash run.sh r

build:               ## Build the binary
	@bash run.sh b

build-run:           ## Build then run
	@bash run.sh br

r:   run
b:   build
br:  build-run

# ── Code Quality ─────────────────────────────────────────────────

fmt:                 ## Format all Go files
	@bash run.sh f

lint:                ## Run go vet
	@bash run.sh l

vet:                 ## Run go vet
	@bash run.sh v

tidy:                ## Run go mod tidy
	@bash run.sh ty

pre-commit:          ## Check Coverage* files for API mismatches
	@bash run.sh pc $(ARGS)

f:   fmt
l:   lint
v:   vet
ty:  tidy
pc:  pre-commit

# ── Other ────────────────────────────────────────────────────────

clean:               ## Clean build artifacts
	@bash run.sh c

c:   clean

help:                ## Show this help
	@echo ""
	@echo "  Available targets:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*## "}; {printf "    \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "  Pass flags via ARGS: make tc ARGS=\"--open --dry-run\""
	@echo ""

h:   help

.DEFAULT_GOAL := help
