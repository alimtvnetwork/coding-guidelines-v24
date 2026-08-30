# Plan: CI/CD Pipeline Consolidation & Owner Review

**Status:** RESOLVED & CONSOLIDATED

> **/goal** Consolidate pipeline workflow specifications into canonical standards under `spec/12-cicd-pipeline-workflows/` upon owner review.
> **/learn** Review consolidated specifications in `spec/12-cicd-pipeline-workflows/` and modular quality guards under `03-reusable-ci-guards/`.

## Owner Decisions Applied (2026-08-30)

1. **Known Issues & Fixes & RCA Skew**:
   - Merged `22-known-issues-and-fixes.md` and `28-rca-release-skew.md` directly into [`17-release-pipeline-issues-rca.md`](../../../spec/12-cicd-pipeline-workflows/17-release-pipeline-issues-rca.md) as Issue #13. Redundant files removed.
2. **Lint Gating Rules**:
   - Established [`18-lint-gating-rules.md`](../../../spec/12-cicd-pipeline-workflows/18-lint-gating-rules.md) as the unified architectural overview pointing to all 12 concrete guard specifications in [`03-reusable-ci-guards/`](../../../spec/12-cicd-pipeline-workflows/03-reusable-ci-guards/01-index.md).
3. **AI Release Synchronization**:
   - Merged `24-ai-release-synchronization.md` directly into [`05-release-pipeline.md`](../../../spec/12-cicd-pipeline-workflows/05-release-pipeline.md) with full step-by-step checklist.
4. **Changelog & Enum Enforcement**:
   - Moved into `03-reusable-ci-guards/` as [`11-changelog-awk-integration.md`](../../../spec/12-cicd-pipeline-workflows/03-reusable-ci-guards/11-changelog-awk-integration.md), [`12-strict-enum-enforcement.md`](../../../spec/12-cicd-pipeline-workflows/03-reusable-ci-guards/12-strict-enum-enforcement.md), and [`13-query-wrapper-python-ts.md`](../../../spec/12-cicd-pipeline-workflows/03-reusable-ci-guards/13-query-wrapper-python-ts.md).
5. **Gitmap Extended Workflows**:
   - Preserved under `04-gitmap-pipeline/` as reference templates.

## Summary of Changes
- Total root spec files in `spec/12-cicd-pipeline-workflows/`: **22 sequential files** (`01-index.md` to `22-e2e-testing-pattern.md` + `99-consistency-report.md`).
- All `❓ Pending Owner Review` markers resolved and cleared.
