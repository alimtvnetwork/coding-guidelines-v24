# Learned Protocol: Letterly Blueprint, CI/CD Fix GitMap Release & Run Script Architecture

> **Reference Specification:** `02-spec/21-app/05-letterly-prompts-and-run-scripts/01-architecture-spec.md`  
> **Associated Prompts:** `01-prompts/22-letterly/*-letterly.md`, `01-prompts/16-ci-cd/11-ci-cd-fix-gitmap-release.md`, `01-prompts/14-execute/14-run.md`, `01-prompts/16-ci-cd/04-create-run-ps1-file.md`

## 1. Letterly Prompt Architecture & The `execute-n-steps` Blueprint

1. **The Superior Structural Blueprint:**
   All desktop-oriented Letterly prompts (Desktop Base, Plan, Release, CI/CD Fix Release) MUST strictly adopt the `execute-n-steps` 5-block structure:
   - `# High Priority Instruction`
   - `${Input Text Verbatim}` (placed directly beneath the header)
   - `# Actionable Items Must Follow Non-Negotiable`
     - Action item 1 is ALWAYS: `1. Write spec and plan first` (or `Write the plan and architectural spec first` for planning)
     - Subsequent items capture discrete technical directives extracted from verbatim input
   - Mandatory Skill Invocation Suffix: `Must follow and spawn agent using [<skill>](file;.agents/skills/<skill>)`
   - Additional Instructions: `learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.`

2. **Standard File Suffix Mandate:**
   All prompt transformation files under `01-prompts/22-letterly/` MUST use the explicit `-letterly.md` file suffix (`01-mobile-letterly.md`, `02-desktop-letterly.md`, etc.).

3. **Mobile Base Prompt Invariant:**
   The mobile prompt format is strictly ONE continuous single-line paragraph with ZERO newlines, line breaks, or code fences:
   - Starts with `# High Priority Instruction: `.
   - Followed by cleaned verbatim text.
   - Suffixes with ` - must follow the skill [execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)`.
   - Strictly omits `[/goal]` or `[/learn]` prefixes.

4. **Skill Hygiene & Older Skills Purge:**
   - Keep only `letterly-*` skills in `.agents/skills/` and `.cursor/skills/`.
   - Purge older non-prefixed skills (`desktop`, `mobile`, `execute-n-steps`).

---

## 2. CI/CD Fix GitMap Release Protocol (`ci-cd-fix-gitmap-release`)

1. **Continuous Telemetry Loop:**
   - Run `gitmap pe -t` to watch workflow execution timelines with dynamic ETA.
   - Run `gitmap pe` to dump failure logs.
2. **Grounded 4-Part RCA:**
   - Document Symptom, Root Cause, Surgical Fix, and Verification under `.ai-memory/cicd-issues/`.
3. **Mandatory Minor Version Bump:**
   - Run `python 03-ai-scripts/37-bump-version.py -t minor -s "<summary>"`.
   - Update manifests (`package.json`, `version.json`, `readme.md`, `changelog.md`).
   - If requested, create release branch `release/vX.Y.Z`.
   - Commit via `gitmap cpf "<module> - fix CI/CD and release minor version"`.
   - Tag release `git tag -a vX.Y.Z -m "..."` and push to remote.
   - Re-check `gitmap pe -t` until completely green.

---

## 3. Run & Install Script Architecture

1. **Manifest-Driven Execution (`run.config.json`):**
   - Single source of truth defining ports, service directories, and lifecycle commands.
2. **Self-Healing Fallback in `run.ps1` / `run.sh`:**
   - When executing `run.ps1`, perform a preflight dependency check.
   - If dependencies or compilers are missing:
     1. Check if `gitmap` is available and run `gitmap aum install`.
     2. Otherwise, fall back to executing `.\local-install.ps1` (or `./local-install.sh`).
     3. Once installation completes, automatically rerun `run.ps1`.
3. **Dedicated Run Prompt (`14-run.md`):**
   - Provides a direct `/run <target>` command in the execute library.
