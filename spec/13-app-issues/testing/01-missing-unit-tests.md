# Missing Unit Tests for Many Packages

## Issue Summary

While `tests/integratedtests/` has test directories for some packages, many packages lack dedicated test coverage (e.g., `simplewrap/`, `coreutils/`, `mutexbykey/`, `namevalue/`, etc.).

## Root Cause Analysis

Test coverage was focused on core packages during development.

## Fix Description

Create test directories and test cases for all packages following the existing AAA pattern with `coretestcases.CaseV1`.

## TODO and Follow-Ups

- [ ] Audit which packages lack tests
- [ ] Prioritize by risk/usage
- [ ] Create tests following established patterns

## Done Checklist

- [ ] All critical packages have test coverage
