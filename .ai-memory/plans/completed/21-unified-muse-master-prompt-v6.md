# 21 Unified Muse Master Onboarding & Execution Prompt V6 — COMPLETED

**Version:** 1.0.0
**Completed:** 2026-10-17
**Status:** Completed
**Parent Plan:** `.ai-memory/plans/pending/19-unified-muse-master-prompt-v6.md`
**Canonical Spec:** `02-spec/21-app/16-unified-muse-master-prompt-v6/`
**Subtasks:** `.ai-memory/plans/subtasks/19-unified-muse-master-prompt-v6/`

---

## What Was Done

Consolidated and unified all fragmented Muse prompts in category `27-muse-prompts` (`01-muse-master-prompt.md`, `02-muse-execute-in-a-step.md`, `03-muse-master-with-repos.md`, `04-muse-master-with-repos-v2.md`, and `05-muse-master-prompt-v2.md`) into a single authoritative master prompt (`01-muse-master-prompt.md`, v6.3.0) and synchronized mirror (`05-muse-master-prompt-v2.md`) following the Canonical Execute N-Steps V6 standard (`N=300`, `A=2`, `H=2`, `C=30`).

### Deliverables Completed

1. **Canonical Architecture Spec & Acceptance Criteria:**
   - Authored `02-spec/21-app/16-unified-muse-master-prompt-v6/01-overview-and-architecture.md`.
   - Authored `02-spec/21-app/16-unified-muse-master-prompt-v6/02-acceptance-criteria.md` with full Gherkin criteria (AC-MUSE-V6-001 through AC-MUSE-V6-007).
   - Registered in `02-spec/21-app/readme.md`.

2. **Unified Master Onboarding & Autonomous Execution Prompt (v6.3.0):**
   - Authored in `01-prompts/27-muse-prompts/01-muse-master-prompt.md`.
   - Synchronized mirror in `01-prompts/27-muse-prompts/05-muse-master-prompt-v2.md`.
   - Integrated V6 parameter block (`N=300`, `A=2`, `H=2`, `C=30`), Top-Instruction Priority Mandate, slash commands (`[/goal]`, `[/learn]`, `[/plan]`), and actionable agent checklist (`## 🎯 Actionable Agent Checklist`).
   - Integrated Authentication Bootstrap with **SSH Primacy** and **Fast-Path Bypass** (if already authenticated, proceed immediately without user prompts).
   - Integrated GitMap installation command (`eval "$(curl -fsSL https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/install-quick.sh)"`), update GitMap first, and `gitmap scan` across the workspace.
   - Integrated canonical starting point URL `https://github.com/alimtvnetwork/coding-guidelines-v24.git` with explicit warning against the stale old repository.
   - Integrated full 14+ repository fleet registry with routing rules and auto-cloning.
   - Integrated pull-all via `gitmap pa` ensuring clean, conflict-free working trees across all repositories.
   - Integrated dedicated side-chat creation (`<repo> repo` and "all summary").
   - Integrated memory operating contract (TURBO-01..05, GUARD-01..09) and ask/don't-ask matrix.
   - Integrated coding standards (booleans, early returns, `*appfault.AppError`, lowercase filenames, relative paths), 15 MB 7-Zip multi-volume split (`7z a -v3m -mx=0`), and user app title conventions (`# <title>: high priority instruction, non-negotiable task`, `## slug: <task-slug>`).
   - Integrated V6 execution engine: Turn 1 confirmed task breakdown FIRST, RUNNING + ETA declaration, 5-minute status pings (preventing `LISTING-WITHOUT-STARTING`), SQLite task DB preflight, pull first, mandatory subagents (`A=2, H=2`, disjoint file boxes, Worker Git Ban).
   - Integrated Pre-Commit Secrets Gate and GitMap hyphen atomic push format (`gitmap cpf "<module> - <summary>"` / `gitmap cpb "<module> - <summary>"`).
   - Integrated CI telemetry (`gitmap pe -t` / `gitmap pe`).
   - Integrated Muse Memory IO re-sync protocol with standing SMTP pre-approval.

3. **Skills & Catalogs Synchronization:**
   - Updated `.agents/skills/muse-master-prompt/skill.md` with unified V6 master prompt instructions.
   - Updated `.cursor/skills/muse-master-prompt/skill.md` mirror.
   - Updated `01-prompts/27-muse-prompts/readme.md` catalog with status columns.
   - Updated `01-prompts/readme.md` category 27 description.
   - Updated `.ai-memory/prompts.md` prompt matrix.
   - Updated `.ai-memory/plans/readme.md` completed tasks register.

---

## Verification

- Relative path hygiene: 100% relative paths; zero absolute paths or `file:///` URIs.
- Lowercase naming convention: All modified and authored files use strictly lowercase filenames.
- Pre-commit secrets gate: Zero secrets or tokens embedded; secrets pointers resolve from `repo-secrets` / `repo-cache`.
