# Archived Agent Skills (`.agents/skills-archived/`)

This directory houses deprecated, legacy, stub, and redundant skills that have been retired from active agent discovery in `.agents/skills/`.

Archiving these skills prevents prompt context bloat, eliminates conflicting skill definitions, and ensures all autonomous agents route directly to modern, consolidated canonical skills.

> [!NOTE]
> Skills located in this directory are preserved strictly for historical reference and audit purposes. They must not be synchronized across downstream repositories or linked from active prompt templates.

---

## Archived Skills Catalog

| # | Archived Skill | Category | Reason for Archival | Canonical Replacement |
| :---: | :--- | :--- | :--- | :--- |
| **01** | `audit-app-spec-old` | Deprecated -old | Legacy specification audit implementation superseded by modern v2 audit engine. | [audit-app-spec](file;.agents/skills/audit-app-spec) |
| **02** | `fix-spec-from-audit-old` | Deprecated -old | Superseded by modern 1:1 remediation checklist engine in `fix-spec-from-audit`. | [fix-spec-from-audit](file;.agents/skills/fix-spec-from-audit) |
| **03** | `read-memory-old` | Deprecated -old | Superseded by enhanced memory retrieval with git commit ingestion in `read-memory-enhanced`. | [read-memory-enhanced](file;.agents/skills/read-memory-enhanced) |
| **04** | `plan-maximum-enforcement` | Legacy Planning | Fixed 5-step legacy planning stub; superseded by modular V2 spec planning engine. | [plan-spec-steps-v2](file;.agents/skills/plan-spec-steps-v2) |
| **05** | `plan-maximum-enforcement-v5` | Legacy Planning | Arbitrary `[N=?]` step planning stub; replaced by structured tree-swarm planning. | [plan-spec-steps-v2](file;.agents/skills/plan-spec-steps-v2) |
| **06** | `plan-spec-steps` | Legacy Planning | Original v1 spec steps planner; upgraded to comprehensive v2 application spec engine. | [plan-spec-steps-v2](file;.agents/skills/plan-spec-steps-v2) |
| **07** | `plan-steps` | Legacy Planning | Unstructured step-by-step task decomposition stub; replaced by bounded subtask engine. | [plan-spec-steps-v2](file;.agents/skills/plan-spec-steps-v2) |
| **08** | `plan-steps-by-groups` | Legacy Planning | Grouped task decomposition stub; replaced by bounded subtask engine. | [plan-spec-steps-v2](file;.agents/skills/plan-spec-steps-v2) |
| **09** | `boolean-improvements` | Duplicate v1 Wrapper | Legacy v1 boolean improvement wrapper superseded by v2 and comprehensive `cg-*` suite. | [boolean-improvements-v2](file;.agents/skills/boolean-improvements-v2) / [cg-boolean-and-naming](file;.agents/skills/cg-boolean-and-naming) |
| **10** | `proofread` | Duplicate v1 Wrapper | First-generation proofreading stub superseded by on-the-fly v2 proofreading engine. | [proofread-v2](file;.agents/skills/proofread-v2) |
| **11** | `autonomous-qa-and-testing` | Duplicate v1 Wrapper | Initial test execution wrapper superseded by comprehensive v4 test suite verification. | [autonomous-qa-and-testing-v4](file;.agents/skills/autonomous-qa-and-testing-v4) |
| **12** | `temp-end-to-end-test` | Duplicate v1 Wrapper | Singular naming stub superseded by canonical plural form `temp-end-to-end-tests`. | [temp-end-to-end-tests](file;.agents/skills/temp-end-to-end-tests) |
| **13** | `nuclear-package-modularization` | Duplicate v1 Wrapper | Early package modularization stub superseded by full unit test optimization skill. | [nuclear-package-modularization-and-unit-test-optimization](file;.agents/skills/nuclear-package-modularization-and-unit-test-optimization) |
| **14** | `branch-immutability-and-clean-construction` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-branch-immutability-and-clean-construction](file;.agents/skills/cg-branch-immutability-and-clean-construction) |
| **15** | `clean-repo-build-and-caches` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-clean-repo-build-and-caches](file;.agents/skills/cg-clean-repo-build-and-caches) |
| **16** | `clean-work-artifacts-and-os-caches` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-clean-repo-build-and-caches](file;.agents/skills/cg-clean-repo-build-and-caches) |
| **17** | `code-dryness-and-library-extraction` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-code-dryness-and-library-extraction](file;.agents/skills/cg-code-dryness-and-library-extraction) |
| **18** | `file-size-and-function-reduction` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-file-size-and-function-reduction](file;.agents/skills/cg-file-size-and-function-reduction) |
| **19** | `golang-pointer-reduction-and-value-semantics` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-golang-pointer-reduction-and-value-semantics](file;.agents/skills/cg-golang-pointer-reduction-and-value-semantics) |
| **20** | `extract-generic-types-to-types-go` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-extract-types](file;.agents/skills/cg-extract-types) |
| **21** | `string-operations-and-efficiency` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-string-efficiency](file;.agents/skills/cg-string-efficiency) |
| **22** | `string-normalization-and-equalfoldany` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-string-normalization-and-equalfoldany](file;.agents/skills/cg-string-normalization-and-equalfoldany) |
| **23** | `variadic-and-spread-parameters` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-variadic-and-spread-parameters](file;.agents/skills/cg-variadic-and-spread-parameters) |
| **24** | `isolate-destructive-os-and-heavy-unit-tests` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-isolate-os-tests](file;.agents/skills/cg-isolate-os-tests) |
| **25** | `booleans-and-complex-conditions` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-boolean-and-naming](file;.agents/skills/cg-boolean-and-naming) |
| **26** | `naming-conventions-and-boolean-prefixes` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-boolean-and-naming](file;.agents/skills/cg-boolean-and-naming) |
| **27** | `constants-and-enums` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-enum-standards](file;.agents/skills/cg-enum-standards) |
| **28** | `data-and-schema` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-database-conventions](file;.agents/skills/cg-database-conventions) |
| **29** | `error-management` | 20-Line Stub Duplicate | Single-rule prompt wrapper consolidated into repository-wide `cg-*` standard suite. | [cg-error-management](file;.agents/skills/cg-error-management) |
| **30** | `ci-cd-fix-with-n-steps` | Redundant CI/CD Variant | Legacy iterative runner stub; consolidated into unified CI/CD remediation engine. | [ci-cd-fix](file;.agents/skills/ci-cd-fix) |
| **31** | `ci-cd-fix-with-release` | Redundant CI/CD Variant | Redundant permutation; consolidated into `ci-cd-fix-gitmap-release`. | [ci-cd-fix-gitmap-release](file;.agents/skills/ci-cd-fix-gitmap-release) |
| **32** | `ci-cd-fix-with-release-n-steps` | Redundant CI/CD Variant | Redundant permutation; consolidated into `ci-cd-fix-gitmap-release`. | [ci-cd-fix-gitmap-release](file;.agents/skills/ci-cd-fix-gitmap-release) |
| **33** | `ci-cd-fix-with-release-tweak` | Redundant CI/CD Variant | Redundant permutation; consolidated into `ci-cd-fix-gitmap-release`. | [ci-cd-fix-gitmap-release](file;.agents/skills/ci-cd-fix-gitmap-release) |
| **34** | `ci-cd-fix-tweak` | Redundant CI/CD Variant | Redundant permutation; consolidated into `ci-cd-fix`. | [ci-cd-fix](file;.agents/skills/ci-cd-fix) |
| **35** | `cicd-pipeline-create` | Redundant CI/CD Variant | Legacy single-file prompt wrapper; superseded by modern Python pipeline architect. | [ci-cd-create](file;.agents/skills/ci-cd-create) |
| **36** | `fix-ci-cd-and-run-scripts` | Redundant CI/CD Variant | Broad legacy combined runner stub; split into specialized CI/CD fix and runner skills. | [ci-cd-fix](file;.agents/skills/ci-cd-fix) / [run](file;.agents/skills/run) |
| **37** | `raw-insults` | Tone Stub | Pure conversational tone modifier containing zero executable actions or tooling. | (Removed — Tone stubs eliminated) |
| **38** | `consolidated-insults-v2` | Tone Stub | Pure conversational tone modifier containing zero executable actions or tooling. | (Removed — Tone stubs eliminated) |
| **39** | `insult-code-fix` | Tone Stub | Pure conversational tone modifier containing zero executable actions or tooling. | (Removed — Tone stubs eliminated) |
| **40** | `read-memory-latest` | Deprecated Wrapper | Superseded by enhanced memory retrieval with git commit ingestion in `read-memory-enhanced`. | [read-memory-enhanced](file;.agents/skills/read-memory-enhanced) |
| **41** | `unified-ai-prompt-v4` | Legacy Orchestrator | Deprecated unified orchestrator prompt stub superseded by V6 multi-agent swarm architecture. | [execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6) |

---

## Archival Invariants

1. **Strict Downstream Sync Exclusion:** The `.agents/skills-archived/` folder is strictly excluded from `03-ai-scripts/38-sync-prompts-skills-scripts.py` and `gitmap sync`. Archived skills are never mirrored to connected downstream repositories.
2. **Strict Relative Git Paths:** All documentation, catalogs, and markdown links must use repository-root relative paths. Absolute filesystem paths and absolute file scheme URIs are strictly banned.
3. **Strict Lowercase File Naming:** All directory names, file names, and references must strictly adhere to lowercase hyphenated naming conventions.
