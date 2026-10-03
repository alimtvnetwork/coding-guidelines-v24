# Architecture Specification: Letterly Prompts Modernization, CI/CD Fix GitMap Release & Run Script Architecture

> **Specification Version:** 1.0.0  
> **Status:** Approved / Active  
> **Target Package / Directory:** `01-prompts/22-letterly/`, `01-prompts/16-ci-cd/`, `01-prompts/14-execute/`, `.agents/skills/`, `.cursor/skills/`

---

## 1. Executive Summary & Core Objectives

This architecture specification standardizes three critical operational domains:
1. **Letterly Voice-Prompt Engineering (`01-prompts/22-letterly/`):** Adopts the superior `execute-n-steps` structural blueprint across all desktop and specialized prompts (Plan, Release, CI/CD Fix), adds the explicit `-letterly.md` file suffix, refines the introduction, implements the strict one-line compact mobile base prompt with `# High Priority Instruction:`, and purges older non-letterly skills.
2. **CI/CD Fix GitMap Release (`01-prompts/16-ci-cd/11-ci-cd-fix-gitmap-release.md`):** Renames and modernizes `ci-cd-fix-gitmap` to `ci-cd-fix-gitmap-release`, explicitly codifying the minor version bump script ceremony (`python 03-ai-scripts/37-bump-version.py -t minor -s "<summary>"`), release branch option, atomic GitMap commit (`gitmap cpf`), and synchronization with companion release skills.
3. **Execution & Run Orchestration (`01-prompts/14-execute/14-run.md` & `01-prompts/16-ci-cd/04-create-run-ps1-file.md`):** Introduces a dedicated `run.md` prompt in the execute folder to execute `run.ps1` (or `run.sh`), refactors `04-cicd-run-ps1.md` into `04-create-run-ps1-file.md` following V6 execute N-steps architecture, and codifies the self-healing fallback to `local-install.ps1` / `local-install.sh` via GitMap or standalone installers.

---

## 2. Letterly Prompts Architecture (`22-letterly`)

### 2.1. File Suffix & Organization
All prompt files in `01-prompts/22-letterly/` MUST carry the explicit `-letterly.md` suffix:
- `01-mobile-letterly.md` (renamed from `01-mobile.md`)
- `02-desktop-letterly.md` (renamed from `02-desktop.md`)
- `03-execute-n-steps-letterly.md` (renamed from `03-execute-n-steps.md`)
- `04-plan-letterly.md` (renamed from `04-plan.md`)
- `05-release-letterly.md` (renamed from `05-release.md`)
- `06-cicd-fix-release-letterly.md` (renamed from `06-cicd-fix-release.md`)
- `07-mobile-cicd-fix-letterly.md` (renamed from `07-mobile-cicd-fix.md`)
- `readme.md` (updated directory index)

### 2.2. The Superior Blueprint (`execute-n-steps`)
The desktop format across all prompts follows the deterministic 5-block structure:
```markdown
# High Priority Instruction

${Input Text Verbatim}

# Actionable Items Must Follow Non-Negotiable

1. Write spec and plan first (or Write plan and spec first in 02-spec/21-app/)
2. [Domain-specific step 1]
3. [Domain-specific step 2]
...

Must follow and spawn agent using

[<canonical-skill-link>](file;.agents/skills/<skill-name>)

## Additional Instructions

learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.
```

### 2.3. Mobile Base Prompt Invariant
The mobile prompt (`01-mobile-letterly.md`) MUST be strictly compacted into a single line without any newlines or line gaps:
- No `[/goal]` or `[/learn]` at the beginning.
- Starts with `# High Priority Instruction: `.
- Followed by cleaned verbatim text.
- Suffixes with ` - must follow the skill [execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)`.

---

## 3. CI/CD Fix GitMap Release Architecture

### 3.1. Purpose & Continuous Loop
- Prompt file: `01-prompts/16-ci-cd/11-ci-cd-fix-gitmap-release.md`.
- Skills: `.agents/skills/ci-cd-fix-gitmap-release/skill.md` and `.cursor/skills/ci-cd-fix-gitmap-release/skill.md`.
- Continuously inspects `gitmap pe -t` (dynamic ETA timeline) and `gitmap pe` (failure logs).
- For every failure, documents 4-part RCA under `.ai-memory/cicd-issues/`.
- Applies surgical code fixes in disjoint files.
- Executes minor version bump: `python 03-ai-scripts/37-bump-version.py -t minor -s "<summary of fixes>"`.
- Verifies manifests (`package.json`, `version.json`, `readme.md`, `changelog.md`).
- Commits atomically via `gitmap cpf "<module> - fix CI/CD and release minor version"`.
- Tags release version (`git tag -a vX.Y.Z -m "..."`) and pushes (`git push origin vX.Y.Z`).
- Loops until remote CI/CD is green (`✔ clean`).

---

## 4. Run Script & Local-Install Architecture

### 4.1. Run Prompt (`01-prompts/14-execute/14-run.md`)
- Invocation: `/run <target>` or `/run-ps1`.
- Executes `.\run.ps1` on Windows pwsh, or `./run.sh` on Unix/macOS.
- Passes configuration flags (`-CI`, `-Verbose`, target task).

### 4.2. Run Script Creation Prompt (`01-prompts/16-ci-cd/04-create-run-ps1-file.md`)
- Upgraded to V6 parameters (`N = 300, A = 2, H = 2, C = 30`).
- Specifies construction of dynamic `run.ps1` / `run.sh` driven by `run.config.json`.
- Integrates self-healing installation fallback:
  ```powershell
  # Dependency pre-flight check
  if (-not (Test-DependenciesInstalled)) {
      Write-Host "Dependencies missing. Attempting installation via GitMap / local-install..."
      if (Get-Command gitmap -ErrorAction SilentlyContinue) {
          gitmap aum install
      } else {
          .\local-install.ps1
      }
      # Re-verify and rerun
      .\run.ps1 @PSBoundParameters
      exit $LASTEXITCODE
  }
  ```

### 4.3. Local-Install Script Contract (`local-install.ps1` & `local-install.sh`)
- Purpose: Standalone environment and dependency bootstrap.
- Strategy:
  1. Detects OS and package manager (`winget`, `choco`, `brew`, `apt`).
  2. Uses GitMap environment bootstrap if available.
  3. Otherwise installs required runtimes (Node, Python, Go, Rust, PHP) and project dependencies.

---

## 5. Acceptance Criteria

1. All 7 Letterly prompts use the `-letterly.md` naming convention.
2. Older non-letterly skills (`desktop`, `mobile`, `execute-n-steps`) are completely purged from `.agents/skills/` and `.cursor/skills/`.
3. `01-mobile-letterly.md` enforces single-paragraph one-liner output starting with `# High Priority Instruction: `.
4. Desktop, Plan, Release, and CI/CD Fix formatters adhere strictly to the `execute-n-steps` format with `1. Write spec and plan first`.
5. `11-ci-cd-fix-gitmap-release.md` and companion skills are authored, registered, and synchronized.
6. `14-run.md` and `04-create-run-ps1-file.md` are established in their respective prompt directories and companion skills created.
7. Linters pass with 0 errors, 0 absolute paths, and full cursor synchronization.
