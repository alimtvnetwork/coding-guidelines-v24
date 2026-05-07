---
id: PH14-14
phase: 14
chapter: 14
title: Rbac And Status Seed
specRef: spec/19-main-worker-service/14-rbac-and-status-seed.md
status: Todo
language: LanguageTBD
repo: RepoTBD
mutationGate: ">=80%"
---

# PH14-14 — Implement Rbac And Status Seed

## Spec Reference
- Source of truth: [`spec/19-main-worker-service/14-rbac-and-status-seed.md`](../../spec/19-main-worker-service/14-rbac-and-status-seed.md)
- This file is **issue-tracking only**. No implementation code lives in `spec/19/`
  (constraint: `mem://constraints/spec19-no-implementation`).

## Scope Summary
Materialize chapter 14 (Rbac And Status Seed) per the spec. Implementation lives outside
`spec/19/` once the target language and repo are decided.

## Acceptance Criteria
- [ ] Implementation passes spec/19 chapter 14 acceptance items.
- [ ] Mutation score >=80% (Patch D gate, closed v2.0.0 in spec/06).
- [ ] Error codes conform to `spec/19/13-error-codes.md` ranges.
- [ ] PascalCase identifiers (Rust exempt per naming memory).

## Dependencies
TBD — populate when adjacent chapters are picked up.

## Status
Todo
