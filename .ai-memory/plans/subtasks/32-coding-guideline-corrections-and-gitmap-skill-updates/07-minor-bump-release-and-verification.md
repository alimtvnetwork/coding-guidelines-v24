# Subtask 07: Minor Bump Release & Verification

- **Parent Plan:** [.ai-memory/plans/pending/32-coding-guideline-corrections-and-gitmap-skill-updates.md](.ai-memory/plans/pending/32-coding-guideline-corrections-and-gitmap-skill-updates.md)
- **Target Repos:** `coding-guidelines` and `gitmap`
- **Status:** Completed

---

## Action Items

1. Rebuild `gitmap.exe` in `d:/work/gitmap` and execute full unit test suite.
2. Deploy new binary to `C:\Users\Administrator\AppData\Local\gitmap-cli\gitmap.exe`.
3. In `coding-guidelines`, perform minor version bump from `v6.86.0` to `v6.87.0`:
   - Update `version.json`, `package.json`, `prompt-version.template.json`, `readme.md`, `changelog.md`.
   - Update `public/health-score.json` and `src/data/specTree.json`.
   - Write `.ai-memory/release/release-notes-v6.87.0.md`.
4. Stage and commit with dedicated release verb `gitmap cpr "v6.87.0 - PowerShell completion fix, follow-cg-execute 39-step orchestrator, and skills reduction"`.
5. Tag `v6.87.0` and push upstream to GitHub across both repositories ('No Push = Not Done').
