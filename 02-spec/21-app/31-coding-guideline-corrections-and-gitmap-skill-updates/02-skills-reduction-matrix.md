# Skills Reduction & Archival Matrix

> **Specification Sequence:** 31  
> **Slug:** `coding-guideline-corrections-and-gitmap-skill-updates`  
> **Target:** Skills Inventory & Consolidation  
> **Date:** 2026-10-27  

---

## 1. Inventory & Proposed Actions

### Group A: Letterly Skills (Archived / Removed from Active Skills)
These are raw dictation formatters intended for transcription ingestion rather than active coding assistants:
- `letterly-cicd-fix-release`
- `letterly-desktop`
- `letterly-execute-n-steps`
- `letterly-execute-with-release`
- `letterly-execute-with-verification`
- `letterly-mobile`
- `letterly-mobile-cicd-fix`
- `letterly-muse-long-update`
- `letterly-muse-sleepy`
- `letterly-plan`
- `letterly-release`
- `letterly-run`

### Group B: Muse Prompt Skills (Archived)
- `muse-master-prompt`

### Group C: Fragmented Python Helper Skills (Consolidate into `coding-guidelines`)
- `python-array-constants-and-dynamic-enums`
- `python-boolean-and-conditional-enhancement`
- `python-constants-and-magic-number-elimination`
- `python-dry-architecture-and-caching`
- `python-dry-caching`

### Group D: Prompt Library Setup (Archived)
- `prompt-library`
- `prompt-library-setup`

### Group E: Release Management Variants (Consolidate)
Currently there are 7 overlapping release skills:
- `release`
- `release-now`
- `release-now-without-test`
- `release-orchestrator`
- `release-management`
- `minor-bump`
- `major-bump`
- `patch-bump`

### Group F: CG Execute Fragmentation -> Proposed `follow-cg-execute`
Currently there are dozens of individual `cg-*` skills that mirror individual prompts:
- `cg-boolean-and-naming`
- `cg-branch-immutability-and-clean-construction`
- `cg-clean-repo-build-and-caches`
- `cg-code-dryness-and-library-extraction`
- `cg-database-conventions`
- `cg-enum-standards`
- `cg-error-management`
- `cg-extract-types`
- `cg-file-size-and-function-reduction`
- `cg-golang-pointer-reduction-and-value-semantics`
- `cg-golang-pointers-and-mutation-reduction`
- `cg-isolate-os-tests`
- `cg-map-dispatch-and-handler-lookup`
- `cg-string-efficiency`
- `cg-string-normalization-and-equalfoldany`
- `cg-variadic-and-spread-parameters`

**Proposed Solution:** Create a single unified `follow-cg-execute` skill that orchestrates the 39 CG prompt directives sequentially with "next" fast-forwarding micro-batches.
