# Architecture Specification: Relative Paths Mandate & Legacy Execute Skills Purge

## 1. Executive Summary

This specification governs two architectural directives requested for the Prompt Architect meta-repository:
1. **Strict Relative Git Paths Mandate:** Enforcing across all Letterly formatters (`01-prompts/22-letterly/`), Cursor mirrors (`01-prompts/23-cursor-prompts/`), companion skills (`.agents/skills/letterly-*/`, `.cursor/skills/letterly-*/`), and release management specifications (`02-spec/16-generic-release/`, `01-prompts/17-release-management/`) that all paths, markdown links, code citations, changelog entries, and release documentation MUST strictly use relative Git paths from repository root. A TOTAL BAN is placed on absolute filesystem paths (e.g. `C:\...`, `/home/...`, `/root/...`) and `file:///` URIs.
2. **Purge of Legacy Execute Parent Task Skills:** Removing all legacy previous versions of the execute parent task skill (`execute-parent-task`, `execute-parent-task-with-n-steps`, `parent-task-in-below-steps`, `parent-task-n-step-loop`) from both `.agents/skills/` and `.cursor/skills/`, leaving `execute-parent-task-with-n-steps-v6` as the sole canonical execution engine.

---

## 2. User Request (Verbatim)

```text
In all these latterly prompts, can you please add one more line to only add the relative paths, never add the absolute path during your work. Make sure you include that, and also this should be respected in the release page as well. Mention that. Okay. I hope you can respect this and make a final improvements on all the prompts and also seeing the skills and make sure that the skill, like the execute parent task in n steps. So this one, all the previous versions, I want you to remove from the skills. Remember that and confirm that this is really applied properly.
```

---

## 3. Root Cause & Architectural Boundaries

### 3.1 Why Relative Paths Are Non-Negotiable
- **Cross-Platform Portability:** Absolute filesystem paths tie repository assets to a single developer machine (e.g. `D:\work\coding-guidelines`, `C:\Users\Administrator\...`). When opened on Linux runners, macOS laptops, or peer developer environments, absolute paths break immediately.
- **Git Hygiene:** Git tracks repository-relative paths. Inserting absolute paths into plans, specifications, comments, or changelogs pollutes version history and creates false CI/CD drift reports.
- **Release Page Integrity:** Release notes in `changelog.md` and release announcements published via GitHub Releases (`02-spec/16-generic-release/`, `01-prompts/17-release-management/`) are read publicly. Leaking local filesystem paths exposes workstation usernames, paths, and internal drive structures, violating repository security standards.
- **IDE Link Resolution:** Both Antigravity and Cursor IDE engines parse relative paths (`02-spec/...`, `.agents/skills/...`) cleanly, whereas raw `file:///` URIs or host-dependent absolute paths break IDE symbol resolution and navigation.

### 3.2 Why Legacy Execute Skills Must Be Purged
- **Skill Redundancy & Confusion:** The existence of multiple execute skills (`execute-parent-task`, `execute-parent-task-with-n-steps`, `parent-task-in-below-steps`, `parent-task-n-step-loop`) alongside `execute-parent-task-with-n-steps-v6` leads autonomous agents to inadvertently select legacy workflows lacking parameterization, SQLite logging, or GitMap search primacy.
- **Canonical V6 Supremacy:** Canonical V6 (`execute-parent-task-with-n-steps-v6`) incorporates SQLite task management (`46-agent-sqlite-task-manager.py`), mandatory subagent dispatch (`A = 2, H = 2`), GitMap search primacy (`gitmap aum search`, `gitmap find`), and secrets gates. Older skills are completely obsolete and retired.
- **Clean Registry Surface:** Purging older skill directories from `.agents/skills/` and `.cursor/skills/` ensures that any AI agent querying available skills sees only the active, canonical V6 skill.

---

## 4. Architectural Invariants

1. **Strict Relative Git Paths:** All generated prompts, actionable items, release notes, changelogs, and specs MUST use relative paths starting from the repository root (e.g. `02-spec/21-app/<slug>/`, `.ai-memory/plans/<slug>.md`, `changelog.md`).
2. **TOTAL BAN on Absolute Paths & `file:///` URIs:** Under no circumstances may an agent emit or record absolute paths (`C:\...`, `/home/...`) or `file:///` URIs in repository files.
3. **Single Canonical Execute Skill:** The only execute parent task skill permitted in `.agents/skills/` and `.cursor/skills/` is `execute-parent-task-with-n-steps-v6`. All predecessor versions are permanently removed.
4. **Positive Booleans Only:** Invariant boolean checks must use positive prefixes (`is` and `has`) with zero explicit `== true` evaluations.
5. **GitMap Primacy:** All discovery and search operations must use GitMap commands (`gitmap aum search`, `gitmap find`, `gitmap cat`).
