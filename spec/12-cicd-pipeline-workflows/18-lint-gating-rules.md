# Lint Gating Rules & Reusable CI Guard Architecture

> **/goal** Provide strict, automated lint gating across all language toolchains using baseline diffing and non-blocking grandfathered rules.
> **/learn** Read the modular guard specifications under `03-reusable-ci-guards/` before configuring repository lint pipelines.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Apply baseline diff linting to block new violations while grandfathering existing debt.
- [ ] `/learn` Review `03-reusable-ci-guards/05-baseline-diff-lint-gate.md` for algorithm details.
- [ ] `/goal` Ensure all guard scripts emit standard `::error file=...,line=...::` GitHub annotations.
- [ ] `/learn` Execute all guards locally via `python .lovable/ai-fix-scripts/03-cicd-local-runner.py`.

## Modular Guard Index

All concrete guard implementations and patterns are located under [`03-reusable-ci-guards/`](./03-reusable-ci-guards/01-index.md):

* **[02-forbidden-name-guard.md](./03-reusable-ci-guards/02-forbidden-name-guard.md)**: Blocks collision-prone helper names.
* **[03-grandfather-baseline-naming.md](./03-reusable-ci-guards/03-grandfather-baseline-naming.md)**: Enforces naming rules on new identifiers only.
* **[04-cross-file-collision-audit.md](./03-reusable-ci-guards/04-cross-file-collision-audit.md)**: Detects duplicate identifiers across files.
* **[05-baseline-diff-lint-gate.md](./03-reusable-ci-guards/05-baseline-diff-lint-gate.md)**: Fails build only on new lint findings vs baseline.
* **[06-actionable-lint-suggestions.md](./03-reusable-ci-guards/06-actionable-lint-suggestions.md)**: PR comment generator for lint fixes.
* **[07-matrix-test-aggregator.md](./03-reusable-ci-guards/07-matrix-test-aggregator.md)**: Aggregates multi-shard matrix test outputs.
* **[08-shared-cli-wrapper.md](./03-reusable-ci-guards/08-shared-cli-wrapper.md)**: Unified CLI wrapper around all guards.
* **[09-config-schema.md](./03-reusable-ci-guards/09-config-schema.md)**: Unified YAML configuration schema for guards.
* **[10-workflow-templates.md](./03-reusable-ci-guards/10-workflow-templates.md)**: Reusable GitHub Actions workflow starters.
* **[11-changelog-awk-integration.md](./03-reusable-ci-guards/11-changelog-awk-integration.md)**: AWK-free changelog parsing guard.
* **[12-strict-enum-enforcement.md](./03-reusable-ci-guards/12-strict-enum-enforcement.md)**: Type suffix and enum convention checker.
* **[13-query-wrapper-python-ts.md](./03-reusable-ci-guards/13-query-wrapper-python-ts.md)**: Query wrapper CI validation.
