# GitMap CI/CD Self-Healing & Letterly Integration Protocol

> **Reference Specification:** `02-spec/21-app/04-cicd-fix-gitmap/01-architecture-spec.md`  
> **Associated Prompt:** `01-prompts/16-ci-cd/11-ci-cd-fix-gitmap-release.md`  
> **Associated Skills:** `.agents/skills/ci-cd-fix-gitmap-release/skill.md`, `.cursor/skills/ci-cd-fix-gitmap-release/skill.md`

## 1. Core Objectives & Principles

When diagnosing and healing failing CI/CD pipelines across any connected repository:

1. **Continuous Telemetry Loop:** Never rely on blind guesses or manual web browser checks. Run `gitmap pe -t` to watch workflow execution timelines with dynamic ETA, or `gitmap pe` to dump the latest failure logs directly to stdout.
2. **Grounded 4-Part Root Cause Analysis (RCA):**
   - **Part 1 (Symptom):** Exact failing step, job name, workflow YAML, and error stack trace.
   - **Part 2 (Root Cause):** Fundamental defect (syntax error, missing argument, type mismatch, breaking change).
   - **Part 3 (Surgical Fix):** Minimal targeted code modification strictly addressing the root cause.
   - **Part 4 (Verification):** Targeted local linter command verifying the fix before committing.
3. **Strict Zero-Storage & No-Disabling Ban:**
   - Under NO circumstances may an AI agent comment out, disable, or delete a failing CI/CD test, linter, or workflow step to force green status.
   - Always fix the source code so that the check passes legitimately.
4. **Mandatory Minor Version Bump Ceremony:**
   - Every resolved CI/CD cycle must bump the minor version using the repository's standard script:
     `python 03-ai-scripts/37-bump-version.py -t minor -s "<summary of fixes>"`
   - Ensure manifests (`package.json`, `version.json`, `readme.md`, `changelog.md`) are updated.
   - Commit atomically via `gitmap cpf "<module> - fix CI/CD and release minor version"`.
   - Tag release (`git tag -a vX.Y.Z -m "..."`) and push (`git push origin vX.Y.Z`).

## 2. Letterly Voice-to-Prompt Transformation

Letterly converts spoken voice notes into executable prompts for autonomous AI agents:

1. **Desktop Format:** Structured with `# High Priority Instruction`, `[/goal]`, `[/learn]`, `# Actionable Items Must Follow Non-Negotiable`, and mandatory skill suffix:
   `[ci-cd-fix-gitmap-release](file;.agents/skills/ci-cd-fix-gitmap-release)`.
2. **Mobile Format:** Strictly ONE single line / continuous paragraph with zero line breaks or code fences, prefixing with `[/goal] [/learn] Run gitmap pe -t to diagnose...` and suffixing with:
   `- must follow the skill [ci-cd-fix-gitmap-release](file;.agents/skills/ci-cd-fix-gitmap-release)`.

## 3. Tool & Skill Synchronization Invariant

Whenever a CI/CD or execution skill is authored or updated:
- The skill file MUST exist in both `.agents/skills/<skill-name>/skill.md` (for Google Antigravity) and `.cursor/skills/<skill-name>/skill.md` (for Cursor IDE).
- Re-index prompts using `python linter-scripts/check-prompts-loaded.py --fix`.
- Ensure zero absolute paths using `python linter-scripts/check-relative-paths.py`.
