# Architecture Specification: Coding Guideline Corrections & GitMap Skill Updates

> **Specification Sequence:** 31  
> **Slug:** `coding-guideline-corrections-and-gitmap-skill-updates`  
> **Target:** Architecture & Contracts  
> **Date:** 2026-10-27  

---

## 1. Skill Architecture: `follow-coding-guidelines`

### 1.1 Metadata & Location
- **Path:** `.agents/skills/follow-coding-guidelines/skill.md`
- **Mirror:** `.cursor/skills/follow-coding-guidelines/skill.md`
- **Naming Rule:** Strictly lowercase filename (`skill.md`).
- **Core Intent:** Provide a single comprehensive reference and checklist for AI agents to audit, write, and refactor code adhering to all repository coding guidelines.

### 1.2 Structure & Components
```yaml
---
name: follow-coding-guidelines
description: "Autonomously enforce, audit, and follow repository coding guidelines, zero-nesting rules, leaf error sovereignty, parameter structs, positive booleans, and GitMap commit discipline."
---
```
Sections:
1. **Preamble & Learn Loop (`/learn`):** Explaining how AI models continuously learn from violations and apply guardrails.
2. **The 12 Grounded Rules Checklist:**
   - Rule 1: No Generated Code or Artifacts Committed.
   - Rule 2: Function Size Caps (8 lines preferred, 15 lines max; files max 300 lines, React 100).
   - Rule 3: Zero Nested If Statements (guard clauses, early returns).
   - Rule 4: Positive Boolean Standard (`is`, `has` only; implicit evaluation, no mixed polarity).
   - Rule 5: Error Management & Leaf Error Sovereignty (`*appfault.AppError`, no orphaned errors).
   - Rule 6: Narrow Types & Result Containers (`Result[T]`, no raw `interface{}` or `any`).
   - Rule 7: Typed Constants & Enums (Type suffix, PascalCase enum constants).
   - Rule 8: Discrete Types Location (`types.go`, dedicated definition files).
   - Rule 9: DRY Principle & Extract Utility Rules.
   - Rule 10: Immutability First & Parameter Structs (>2-3 arguments -> `*Params` struct).
   - Rule 11: GitMap Command Discipline & Subsystem Module Standard (Title Case or lowercase subsystem; TOTAL BAN on `commit` as module).
   - Rule 12: Zero-Storage CI/CD Mandate & Relative Path Hygiene (`02-spec/...`, `.ai-memory/...`).

---

## 2. GitMap Skill Overhaul: `.agents/skills/gitmap/skill.md`

### 2.1 Casing Normalization
- Replace ALL-CAPS module tags in lines 715–740:
  - `SKILLS - ...` -> `Skills - ...` (or `skills - ...`)
  - `CLONER - ...` -> `Cloner - ...` (or `cloner - ...`)
  - `DOCS - ...` -> `Docs - ...` (or `docs - ...`)
  - `RELEASE - bump minor version...` -> `v6.619.0 - bump minor version and update changelog`
- Reconcile all commit examples throughout `skill.md` to use Title Case or lowercase subsystem modules.

### 2.2 Bloat Removal & Deduping
- **Remove:**
  - `gitmap disk`, `gitmap os storage`, `gitmap os dev-clean`.
  - `gitmap vscode repair`.
  - `gitmap starbar client ls` / `client list`.
  - `gitmap cluster worker`, `gitmap cluster status`.
- **Retain & Elevate:**
  - `gitmap cargo` suite (`cargo build`, `cargo test`, `cargo run`, `cargo install`).
  - Native process termination: introduce `gitmap proc kill <pid> [--force]`.

---

## 3. GitMap CLI Code Architecture (`d:\work\gitmap`)

### 3.1 Non-Git Root Graceful Handling in Show Log
- **Location:** `d:\work\gitmap\cli\cmdlogviewer\server.go` and `cli/cmdgit/cmd_log.go`
- **Issue:** `git rev-parse --show-toplevel` returns exit 128 when run in non-git directory (`D:\work`), causing `findRepoRoot` to emit a full Go stack trace crash.
- **Remediation:** Check if directory is inside a Git repository before invoking `RunShowLog`. If not in a Git repository, output a clean, formatted error card:
  ```text
  ERROR: Directory "D:\work" is not a Git repository.
  INFO:  Navigate into a repository folder or specify a target repo: gitmap show-log <repo>
  ```
  and exit cleanly with `constants.ExitUsageError` without stack trace dumping.

### 3.2 Short Repo Name Resolution in `gitmap clone`
- **Location:** `d:\work\gitmap\cli\cmdclone\clone_resolver.go`
- **Current Behavior:** Fails if passed bare repo name (e.g. `gitmap clone coding-guidelines`).
- **Remediation:** If input does not contain `/` or `:` and does not start with `http://` or `git@`:
  - Check configured user/org (default `alimtvnetwork` or ambient git config).
  - Attempt resolving short slug or appending default prefix: `git@github.com:alimtvnetwork/<repo>.git`.
  - If bare name matches a known repo in GitMap's SQLite catalog (`repodb`), clone using the registered remote URL.

### 3.3 Show Log UI Modernization
- **Window Favicon / Icon:** Embed GitMap SVG/PNG logo into the HTML `<head><link rel="icon" ...></head>` and title bar.
- **Compact UI Controls:** Group Fullscreen, Refresh, and Settings into compact icon buttons with hover tooltips.
- **Searchable Branch Combo:** Add quick fuzzy search filter within branch dropdown.
- **Diff Comparison:** Add double-click file diff view opening Beyond Compare (`BCompare.exe`) if detected, falling back to VS Code diff (`code -d`).
- **UI Settings Modal:** Allow selecting preferred external diff tool (Beyond Compare / VS Code / Sublime Merge / GitMap Built-in).
