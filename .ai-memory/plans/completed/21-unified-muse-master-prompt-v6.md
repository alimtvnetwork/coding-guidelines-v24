# 21 Unified Muse Master Onboarding & Execution Prompt V6 — COMPLETED

**Version:** 1.1.0
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
   - Authored `02-spec/21-app/16-unified-muse-master-prompt-v6/03-onboarding-protocol-spec.md` with the 9-step onboarding ritual, incremental reporting rules, platform commands, SSH pick flow, 8 default repos + 5 side chats, and 9-point final report.
   - Authored `02-spec/21-app/16-unified-muse-master-prompt-v6/04-gitmap-llm-training-spec.md` specifying `gitmap llm train <ai-name>`, 6-phase lifecycle, and flags.
   - Authored `02-spec/21-app/16-unified-muse-master-prompt-v6/readme.md` directory index.
   - Registered in `02-spec/21-app/readme.md`.

2. **GitMap Skills LLM Training & Command Reference Enhancement:**
   - Updated `.agents/skills/gitmap/SKILL.md` and `.cursor/skills/gitmap/SKILL.md` (100% parity).
   - Documented `gitmap llm train <ai-name>` (e.g. `gitmap llm train muse --days 7`).
   - Documented the full 6-phase self-looping AI lifecycle: Discovery, Refactoring, Verification, Semantic Commit, Telemetry, Heal & Fix.
   - Documented flags (`--days`, `--loop`, `--self-loop`, `--heal-apply <stash|wip|discard>`, `--skill-path`, `--text-only`, `--urls`, `-j/--json`).
   - Updated Invariant 1: Mandatory `gitmap llm train muse --days 7` for memory IO sync.

3. **Unified Master Onboarding & Autonomous Execution Prompt (v6.3.0):**
   - Authored in `01-prompts/27-muse-prompts/01-muse-master-prompt.md`.
   - Synchronized mirror in `01-prompts/27-muse-prompts/05-muse-master-prompt-v2.md`.
   - Integrated V6 parameter block (`N=300`, `A=2`, `H=2`, `C=30`), Top-Instruction Priority Mandate, slash commands (`[/goal]`, `[/learn]`, `[/plan]`), and actionable agent checklist (`## 🎯 Actionable Agent Checklist`).
   - Integrated the complete 9-Step Onboarding Ritual with "No Silent Advance" rule (report after every step with action + verification output):
     - Step 1: Install/update GitMap (Linux/macOS curl, Windows PowerShell irm, update gitmap, source build Go 1.26+ fallback, gh existence check, REPORT).
     - Step 2: Login before cloning (gh auth status, pick ONE: a. Browser login or b. SSH keygen and show in chat, hard gate stopping until login succeeds, REPORT).
     - Step 3: Clone default repos table (8 core repos: `gitmap-v28`, `Antigravity-Manager`, `cat-my-v12`, `alim-seo-writing`, `story-telling-v1`, `coding-guidelines-v24`, `repo-cache`, `repo-secrets`), create 5 side chats seeded from `chat-seeds/<name>.md`, define 3 vault/infra repos without side chats, discover secondary repos via `02-repos.md` with explicit user approval gate, REPORT.
     - Step 4: Pull all handpicked repos (`cd ~/workspace/repos && gitmap pa`), REPORT.
     - Step 5: Read master prompt and custom-skills, obey execution contract for steps 6-9, no fenced code blocks around formatters, REPORT restating contract.
     - Step 6: Memory sync (`repo-cache/12-muse-ai/README.md`, `gitmap llm train muse --days 7`, DECISION TRACE + SINCE LAST READ, read sequenced files), REPORT.
     - Step 7: Connect email (`go-email-reader` from `repo-secrets/11-go-email-reader/`, password-only prompt rule, standing SMTP pre-approval, default sender & notifier), REPORT.
     - Step 8: Scan everything into gitmap (`gitmap scan ~/workspace/repos` after cloning), REPORT scan summary.
     - Step 9: Final report block with exact 9 items including OPEN items from digest.
   - Maintained full V6 autonomous execution engine: Turn 1 confirmed task breakdown FIRST, RUNNING + ETA declaration, 5-minute status pings (preventing `LISTING-WITHOUT-STARTING`), SQLite task DB preflight, pull first, mandatory subagents (`A=2, H=2`, disjoint file boxes, Worker Git Ban).
   - Integrated Pre-Commit Secrets Gate and GitMap hyphen atomic push format (`gitmap cpf "<module> - <summary>"`).
   - Integrated CI telemetry (`gitmap pe -t` / `gitmap pe`).
   - Integrated Muse Memory IO re-sync protocol with standing SMTP pre-approval.

4. **Skills & Catalogs Synchronization:**
   - Updated `.agents/skills/muse-master-prompt/skill.md` with 9-step onboarding ritual.
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
