# Architecture Specification: Letterly Prompt Cleanup & Fleet Sync

- **Module / Task Slug:** `letterly-prompt-cleanup-and-fleet-sync`
- **Specification ID:** `02-spec/21-app/letterly-prompt-cleanup-and-fleet-sync/01-architecture-spec.md`
- **Plan File:** `.ai-memory/plans/letterly-prompt-cleanup-and-fleet-sync.md`
- **Status:** APPROVED & ACTIVE
- **Target Release:** `v6.83.0`

---

## 1. Executive Summary & Purpose

Letterly is an input-to-prompt conversion engine designed to take raw, dictated, or spoken user thoughts and format them into canonical, high-discipline execution prompts. 

Execution engine details (such as the 400-step budget, phase partitioning, agent concurrency `A = 3, H = 2`, and Muse persona specifics) belong exclusively in the execution skills ([`execute-parent-task-with-n-steps-v6`](.agents/skills/execute-parent-task-with-n-steps-v6)) and the Muse master prompt family ([`01-prompts/27-muse-prompts/`](01-prompts/27-muse-prompts/)). 

This specification directs:
1. **Generic Letterly Decoupling:** Remove all step budget mandates (e.g. 400-step budget, phase allocations) and any Muse persona specifics from [`01-prompts/22-letterly/03-execute-n-steps-letterly.md`](01-prompts/22-letterly/03-execute-n-steps-letterly.md) and related Letterly formatters.
2. **Text Compaction & Filler Reduction:** Restructure instructions, eliminate filler words ("the", repetitive explanations, verbose rules), and reduce line counts so the prompt formatter is crisp, compact, and deterministic.
3. **Skill Reference Integrity:** Retain clean delegation to `[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)` (and Cursor mirror to `[execute-parent-task-with-n-steps-v6](file;.cursor/skills/execute-parent-task-with-n-steps-v6)`).
4. **Fleet Synchronization:** Mirror changes across `01-prompts/22-letterly/` and `01-prompts/23-cursor-prompts/`, verify against linters, release `v6.83.0`, and propagate to all 43 connected repositories via [`03-ai-scripts/38-sync-prompts-skills-scripts.py`](03-ai-scripts/38-sync-prompts-skills-scripts.py).

---

## 2. Letterly Formatter Canonical Structure

The cleaned, compact Letterly formatter follows this lean contract:

```markdown
# <task title>: high priority instruction, non-negotiable task

${Input Text Verbatim}

## task: <task title>
## slug: <task-slug>

# Actionable Items Must Follow Non-Negotiable

1. Git pull BEFORE starting any work — first try `gitmap pull <repo path>` (or run `gitmap pull` in repository folder); if that fails, fall back to regular `git pull` and resolve any merge conflicts so the repository is completely up to date and clean before starting.
2. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/) first
3. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String
4. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well
5. [Derived discrete actionable items from input text...]

## Must follow and spawn agent using

[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)

## Additional Instructions

- [/plan](slashCommand;plan) first before doing the work to reduce the credits.
- [/learn](slashCommand;learn) from [gitmap](file;.agents/skills/gitmap) skill to leverage GitMap high-speed search, toolchain discovery, and caching.
- Track in-flight file touches and prevent write collisions using GitMap RC Heatmap (`gitmap rc heatmap status` / `gitmap rc heatmap tree` / `gitmap rc heatmap live` / `gitmap rc heatmap sync`).
- Only add relative paths, never add absolute paths during work; this must be respected on the release page and in release notes as well.
```

---

## 3. Boundaries & Invariants

- **No Step Budgeting in Letterly:** Do not mandate 400 steps or phase splits in Letterly prompts. The spawned skill handles execution limits.
- **No Muse Infiltration:** Muse directives belong strictly in `01-prompts/27-muse-prompts/`.
- **Strict Relative Paths:** All citations must remain strictly relative Git paths.
- **Hyphen GitMap Commits:** All GitMap commits use hyphens without colons in arguments.
- **Non-Negotiable Sync Boundaries:** Spec 21 exclusion, bump script protection, memory isolation, and secrets isolation apply during fleet sync.
