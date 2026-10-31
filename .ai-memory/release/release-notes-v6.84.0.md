## [v6.84.0] - 2026-10-22

### Added
- **SQL Repository Pattern & Enum Actions Rules:** Codified four normative rules (spec `02-spec/21-app/03-execute-n-steps-letterly-prompt-formatter/04-sql-repository-pattern-and-enum-actions-spec.md`): (A) all SQL access goes through the repository pattern in every programming language — no SQL strings in services or handlers; (B) every action is dispatched from an enum so actions stay switchable and exhaustively checkable; (C) an enum authoring standard with a util-methods checkup (`String` / `Parse` / `IsValid` / `All`); (D) typed-error tie-ins between enums, repositories, and the error-management guidelines.
- **Fixed Items 5–8 in the Execute-N-Steps Formatters:** The four rules now ship as fixed items 5–8 in `01-prompts/22-letterly/03-execute-n-steps-letterly.md` and its Cursor mirror `01-prompts/23-cursor-prompts/03-execute-n-steps-letterly-cursor.md`, so every generated execution prompt carries them.
- **New Skill `sql-repository-pattern-and-enum-actions`:** Added to `.agents/skills/` and `.cursor/skills/`, covering the repository-pattern layout, enum authoring, exhaustive dispatch, and the util-methods checkup; hash-verified during propagation to all six fleet repositories.

### Changed
- **Enum / Database / Error Skills Fine-Tuned:** Seven related skills and five guideline/error prompts were aligned with the new rules (enum definitions, repository boundaries, typed-error propagation), with no conflicting legacy wording retained.

---

### Quick Install One-Liners (v6.84.0)

**PowerShell:**
```powershell
irm https://github.com/alimtvnetwork/coding-guidelines-v24/releases/download/v6.84.0/release-install.ps1 | iex
```

**Bash:**
```bash
curl -fsSL https://github.com/alimtvnetwork/coding-guidelines-v24/releases/download/v6.84.0/release-install.sh | bash
```

### Release Assets Included
- `coding-guidelines-v6.84.0.zip`
- `coding-guidelines-v6.84.0.tar.gz`
- `release-install.ps1`
- `release-install.sh`
- `checksums.txt`
