# Specification: Single Muse Master Prompt & Fleet-Wide Deprecated Prompt Purge

> Category: `21-app`
> Status: Active
> Architecture: Prompt Consolidation, Legacy Prompt Deletion & Automated Fleet Purge
> Canonical Prompt: `01-prompts/27-muse-prompts/01-muse-master-prompt.md` (v6.4.0)
> Sync Automation: `03-ai-scripts/38-sync-prompts-skills-scripts.py`

---

## Overview

This specification documents the consolidation of category `27-muse-prompts` into a single authoritative master prompt (`01-muse-master-prompt.md`) and the permanent eradication of redundant and deprecated prompt files (`02-muse-execute-in-a-step.md`, `03-muse-master-with-repos.md`, `04-muse-master-with-repos-v2.md`, and `05-muse-master-prompt-v2.md`) across `coding-guidelines` and all 43 connected repositories.

## Documents

| Document | Purpose |
| :--- | :--- |
| [`01-overview-and-architecture.md`](01-overview-and-architecture.md) | Architectural rationale, file scope, and sync automation mechanics |
| [`02-acceptance-criteria.md`](02-acceptance-criteria.md) | Gherkin acceptance criteria for single master prompt integrity and fleet deployment |
