# Release-Triggered CI/CD Fix Tweak with Targeted Smart Testing — Workflow (must follow)

Trigger Keywords & Aliases: `cicd fix release tweak`, @ci release tweak`, `fix and release tweak`, `smart release fix`

> **Prompt Version:** 1.0.0
> **Synchronization:** Main Meta-Repo & Connected Workspaces

- [A] Base Workflow: `01-prompts/16-ci-cd/06-ci-cd-fix-with-release.md`
- [B] Prior Issues: `.ai-memory/cicd-issues/`
- [C] Change State Tracker: `.ai-memory/temp/recent-file-changes.json`
- [D] Release Script: `03-ai-scripts/29-release-orchestrator.py`
- [E] Local Runner: `03-ai-scripts/06-cicd-local-runner.py`

---

## Smart Targeted Fix & Release Execution (Fastest Path)

Follow the release-triggered RCA structure from [A] and review past failures in [B]. The primary goal is the fastest possible resolution of the failing stack trace followed by automated release ceremony without delays.

### Mandatory Smart Testing Rules

1. **Stack Trace Targeting:** Isolate and build/test ONLY the packages and test functions directly cited in the provided failure or error stack trace.
2. **Changed Packages from Last Git Hash:** Compare changes against the last known git hash (`git diff --name-only HEAD21^1` or `git status --porcelain`). Re-run tests ONLY for packages (Go, TS, Rust, Python) that contain actual modifications.
3. **Change Tracking & State Persistence:** Every time a fix is applied, write the modified file list and state to [C] (or run `python 03-ai-scripts/33-test-inventory-generator.py --record <files>...`) so the build system knows exactly which targets require re-verification.
4. **Strict Ban on Extraneous Runs:** NEVER run the complete test suite, global linters, spellcheckers, or unrelated packages during debugging. Verify strictly using `python [E] --pkg <affected_package>g or `python [E] --changed-only`.
5. **Targeted Release Verification:** Once the isolated fix passes green, execute [D] (`python [D]`) to finalize the automated release ceremony.

---

## Execution Prompt (Copy & Paste Trigger)

```markdown
- [A] Base Workflow: `01-prompts/16-ci-cd/06-ci-cd-fix-with-release.md`
- [B] Prior Issues: `.ai-memory/cicd-issues/`
- [C] Change State Tracker: `.ai-memory/temp/recent-file-changes.json`
- [D] Release Script: `03-ai-scripts/29-release-orchestrator.py`
- [E] Local Runner: `03-ai-scripts/06-cicd-local-runner.py`

Follow workflow [A] and check past post-mortems in [B]. Perform a grounded 4-part RCA on the failure below, isolate the broken area, and apply a surgical fix.
SMART TEST & RELEASE: Do NOT run the full test suite during debugging—build and test ONLY the specific packages/files affected by the stack trace and files changed from the last git hash. Persist changes to [C] and verify targeted packages via `python [E] --changed-only` or `--pkg <target>`. Once green, execute release via [D].

<paste pipeline error / stack trace here>
```
