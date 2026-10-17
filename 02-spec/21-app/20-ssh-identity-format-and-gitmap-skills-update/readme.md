# Specification: SSH Identity Format & GitMap Skills Fleet Synchronization

> Category: `21-app`
> Status: Active
> Architecture: SSH Key Identity Formatting, Git Configuration, GitMap LLM Training & Fleet Synchronization
> Canonical Prompt: `01-prompts/27-muse-prompts/01-muse-master-prompt.md` (v6.4.0)
> Companion Skills: `.agents/skills/muse-master-prompt/`, `.agents/skills/gitmap/`

---

## Overview

This specification establishes the standardized format for generating SSH keypairs, setting Git configuration, interactively requesting unknown variables from operators, and expanding GitMap skills across `coding-guidelines` and the 43 connected repositories.

## Documents

| Document | Purpose |
| :--- | :--- |
| [`01-overview-and-architecture.md`](01-overview-and-architecture.md) | Architectural rationale, SSH format `<email>-muse<sequence>-<operator>`, Git commands, and GitMap LLM training |
| [`02-acceptance-criteria.md`](02-acceptance-criteria.md) | Gherkin acceptance criteria for SSH generation, variable prompts, and fleet deployment |
