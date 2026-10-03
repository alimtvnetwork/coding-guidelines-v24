# Subtask 01: Application Spec Harmonization & Style Collision Resolution (AI Execution Prompt)

> **/goal** Harmonize headings and runnable commands in `02-spec/21-app/07-.../`, retag `21-newline-styling-examples.md` to `AC-CG-STYLE-009`, and update `01-cross-language/97-acceptance-criteria.md`.
> **/learn** Enforce strictly relative paths, zero absolute paths or file:/// URIs, LF line endings, and run guideline autofixer verification.

## 🎯 Actionable CI/CD & Agent Checklist

- [x] `/goal` Update `02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/readme.md` criterion ID to `AC-CG-SPEC-000` with verification command.
- [x] `/learn` Normalize `## 6. Verification & Acceptance Criteria` to unnumbered in `01-architecture-spec.md` and retag criterion to `AC-CG-SPEC-001`.
- [x] `/goal` Normalize `## 4. Verification & Acceptance Criteria` in `02-component-spec.md` and add discrete bash verification commands for all 4 criteria.
- [x] `/learn` Retag criterion in `01-cross-language/21-newline-styling-examples.md` to `AC-CG-STYLE-009`.
- [x] `/goal` Add row and Gherkin entry for `AC-CG-STYLE-009` in `01-cross-language/97-acceptance-criteria.md`.
- [x] `/learn` Run `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance --check-only` and `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/01-cross-language --check-only`.

. **CRITICAL AI INSTRUCTION:** Do NOT run git commands. Report findings and edits back to the orchestrator.

**Assigned Worker:** Worker 1
**Status:** Completed
