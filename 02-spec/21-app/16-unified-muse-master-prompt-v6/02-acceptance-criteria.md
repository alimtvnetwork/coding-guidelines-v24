# Acceptance Criteria: Unified Muse Master Onboarding & Execution Prompt (V6)

> Category: `21-app`
> Status: Active
> Reference: `02-spec/21-app/16-unified-muse-master-prompt-v6/01-overview-and-architecture.md`

---

## Acceptance Criteria Register

### AC-MUSE-V6-001: Consolidated Single Master Prompt
- **Given** The user reviews `01-prompts/27-muse-prompts/`.
- **When** Checking `01-muse-master-prompt.md` and related category files.
- **Then** All capabilities from `01-muse-master-prompt.md`, `02-muse-execute-in-a-step.md`, `03-muse-master-with-repos.md`, `04-muse-master-with-repos-v2.md`, and `05-muse-master-prompt-v2.md` are unified into one single master prompt file following Execute N-Steps V6 standards (`N=300`, `A=2`, `H=2`, `C=30`).

### AC-MUSE-V6-002: Authentication with SSH Primacy & Fast-Path Bypass
- **Given** A fresh or existing Muse session boots with the master prompt.
- **When** Phase 1 authentication check runs.
- **Then** The agent checks whether GitHub CLI and GitMap authentication are already active; if active, it immediately moves to next steps without waiting; if inactive, it prefers SSH authentication (`eval "$(ssh-agent -s)"`, `gh auth login -p ssh -w`) with browser fallback.

### AC-MUSE-V6-003: GitMap Installation, Self-Update & Fleet Scan
- **Given** Phase 2 executes.
- **When** Checking GitMap availability.
- **Then** It verifies or installs GitMap via `eval "$(curl -fsSL https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/install-quick.sh)"`, updates GitMap first, and executes `gitmap scan` across the workspace so GitMap discovers and stores all repository metadata in SQLite.

### AC-MUSE-V6-004: Coding Guidelines Repo Starting URL & Fleet Cloning
- **Given** Phase 3 executes.
- **When** Reviewing repository URLs and fleet discovery.
- **Then** The prompt specifies the canonical Coding Guidelines repository URL `https://github.com/alimtvnetwork/coding-guidelines-v24.git`, notes the stale old repository, clones missing repositories from the 14+ fleet registry into `~/workspace/repos/`, pulls all repositories with `gitmap pa`, and checks for conflict-free working trees.

### AC-MUSE-V6-005: Repo-Cache Integration & Dedicated Side Chats
- **Given** Repositories are cloned and updated.
- **When** Managing repository state and side conversations.
- **Then** `repo-cache/NN-<slug>` monotonic sequencing is respected, and dedicated side chats are established for each repository named `<repo> repo` alongside an "all summary" chat.

### AC-MUSE-V6-006: V6 Actionable Checklist & Anti-Looping Discipline
- **Given** The prompt is pasted into a Muse AI chat.
- **When** Inspecting the prompt structure.
- **Then** It features a prominent `## 🎯 Actionable Agent Checklist` with `/goal` and `/learn`, verbatim capture, Turn 1 confirmed task breakdown, RUNNING + ETA declaration, 5-minute status pings, mandatory multi-agent dispatch (`A=2, H=2`), Worker Git Ban, pre-commit secrets gate, large-file 15MB 7-Zip split gate, atomic push via GitMap hyphen format (`gitmap cpf "<module> - <summary>"`), CI telemetry (`gitmap pe -t`), and Muse Memory IO re-sync.

### AC-MUSE-V6-007: Skill and Catalog Synchronization
- **Given** The unified prompt is authored.
- **When** Reviewing `.agents/skills/muse-master-prompt/skill.md`, `.cursor/skills/muse-master-prompt/skill.md`, `01-prompts/27-muse-prompts/readme.md`, `01-prompts/readme.md`, `.ai-memory/prompts.md`, and `02-spec/21-app/readme.md`.
- **Then** All skill definitions, catalogs, and indexes reference the unified master prompt and reflect the consolidated architecture.
