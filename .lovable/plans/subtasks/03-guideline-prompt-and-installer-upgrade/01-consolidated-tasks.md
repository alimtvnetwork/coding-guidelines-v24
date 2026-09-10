# Subtask: Guideline Prompt & Installer Upgrade (Consolidated)

> **Parent Plan:** `.lovable/plans/pending/04-guideline-prompt-and-installer-upgrade.md`  
> **Status:** PENDING  

## 1. Scope & Execution Ledger

| Step | Scope | Description | Status |
|:---:|---|---|:---:|
| 1 | Overview Prompt Conversion | Format `02-spec/02-coding-guidelines/00-overview.md` with explicit AI prompts | PENDING |
| 2 | Installer Upgrade Logic | Update `generate-bundle-installers.mjs` to handle file removals via inner manifest | PENDING |
| 3 | 50 Guideline Improvements | Generate structured improvement rules for polyglot codebases | PENDING |
| 4 | Release Preparation | Verify installer outputs JSON summaries and prepare release notes | PENDING |

## 2. Core Implementation Requirements

- Installers MUST emit JSON summaries after running so CLI tools can parse results.
- Never modify the root `Version` during partial manifest updates.
- All files must use strictly lowercase naming and Unix LF line endings.
