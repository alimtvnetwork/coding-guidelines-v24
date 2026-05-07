---
id: PH14-13
phase: 14
chapter: 13
title: Error Codes
specRef: spec/19-main-worker-service/13-error-codes.md
status: Todo
language: LanguageTBD
repo: RepoTBD
mutationGate: ">=80%"
---

# PH14-13 — Implement Error Codes

## Spec Reference
- Source of truth: [`spec/19-main-worker-service/13-error-codes.md`](../../spec/19-main-worker-service/13-error-codes.md)
- This file is **issue-tracking only**. No implementation code lives in `spec/19/`
  (constraint: `mem://constraints/spec19-no-implementation`).

## Scope Summary
Materialize chapter 13 (Error Codes) per the spec. Implementation lives outside
`spec/19/` once the target language and repo are decided.

## Acceptance Criteria
- [ ] Implementation passes spec/19 chapter 13 acceptance items.
- [ ] Mutation score >=80% (Patch D gate, closed v2.0.0 in spec/06).
- [ ] Error codes conform to `spec/19/13-error-codes.md` ranges.
- [ ] PascalCase identifiers (Rust exempt per naming memory).

## Dependencies
TBD — populate when adjacent chapters are picked up.

## Status
Todo
