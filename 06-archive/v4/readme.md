# Archived V4 Prompt Tier (`06-archive/v4`)

> [!NOTE]
> Tier: Historical V4 Prompt Architecture
> Status: Archived / Read-Only Reference
> Active Canonical Prompts: [`01-prompts/`](01-prompts/)

---

## Overview

This directory archives historical V4 prompt specifications from the Prompt Architect system.

As part of the architecture evolution to V6 (parameter-driven execution, SQLite task concurrency, GitMap hyphen-separated atomic commits, secrets gate, and ledger resume):
- **Archived V4 Prompts:** Stored here under `06-archive/v4/` for historical reference and audit lineage.
- **Canonical V6 Prompts:** The active, operational prompt suite resides in [`01-prompts/`](01-prompts/). All active workflows must invoke canonical prompts from `01-prompts/`.
- **Sync Exclusion Policy:** `06-archive/` and all its subdirectories (`v1/`, `v2/`, `v3/`, `v4/`, `execute/`) are strictly excluded from downstream synchronization. Downstream repositories only receive active canonical files from `01-prompts/`.
