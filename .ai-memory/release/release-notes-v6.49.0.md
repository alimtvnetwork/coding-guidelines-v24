## Quick Install v6.49.0

### Windows (PowerShell)

```powershell
Invoke-WebRequest -Uri https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v24/v6.49.0/install.ps1 -OutFile install.ps1; .\install.ps1 -TargetDir ".ai-memory/prompts" -Version "v6.49.0"
```

### Unix / Linux / macOS (Bash)

```bash
curl -sL https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v24/v6.49.0/install.sh | bash -s -- ".ai-memory/prompts" "v6.49.0"
```

---

## What's Changed in v6.49.0

### Added
- `01-prompts/14-execute/11-execute-parent-task-with-n-steps-v4.md`: V4 of the execute-parent-task prompt, rewritten for Antigravity 2.0 (rules R1 to R15 stated once, capability preflight, resumable ledger, corrected `invoke_subagent` payload, explicit-path staging, evidence-gated checks).
- Thin pointer skills for V4 in `.agents/skills/execute-parent-task-with-n-steps-v4/` and `.cursor/skills/execute-parent-task-with-n-steps-v4/`.
- Plan 15 (V3 audit and V4 design), plan 16 (this release and the execute-folder hardening), and the open-conventions note `.ai-memory/ambiguous-questions/01-new-ambiguity/02-execute-v4-open-conventions.md`.

### Fixed
- Renumbered the duplicate `01-prompts/15-cg-execute/34-clean-work-artifacts-and-os-caches.md` to `35-clean-work-artifacts-and-os-caches.md`, and updated the prompt index, the folder readme, and its Cursor skill pointer.
- Regenerated the 14 bundle installers (`*-install.sh`, `*-install.ps1`), whose banners and usage examples still pinned v6.46.0.

### Issues
- `.ai-memory/release/issues/01-6.49.0-cg-execute-duplicate-sequence.md`: the pre-release gate failed on a duplicate `34-` prompt number.
- `.ai-memory/release/issues/02-6.49.0-stale-installer-version-pins.md`: the bump script does not update the bundle installers.
