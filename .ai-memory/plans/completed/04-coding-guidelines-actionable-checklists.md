# Completed Plan: Coding Guidelines Actionable Checklists & Acceptance Criteria Transformation

> **/goal** Autonomously transform 32 cross-language coding guidelines into actionable prompt-driven specifications with agent checklists and testable acceptance criteria using execute-parent-task-with-n-steps-v6.
> **/learn** Enforce disjoint file boundaries, bounded 5–8 file micro-batches, parallel subagent dispatch (A = 2, H = 2), and zero-lock execution under execute-parent-task-with-n-steps-v6.

**Version:** 1.0.0  
**Completed:** 2026-10-02  
**Status:** COMPLETED  
**Parent Task ID:** 1 (SQLite DB: `.ai-memory/temp-agents/04-coding-guidelines-actionable-checklists/agent-task.db`)  
**Canonical Spec:** [`02-spec/21-app/04-coding-guidelines-actionable-checklists-and-acceptance-criteria/01-architecture-spec.md`](../../02-spec/21-app/04-coding-guidelines-actionable-checklists-and-acceptance-criteria/01-architecture-spec.md)  

---

## 1. User Request (Verbatim)

```text
# High Priority Instruction

Can you please go into the coding guideline and see the coding guideline, how it is written? It's not written as an action or prompt on every file. Try to have from the README.md file inside the coding guideline folder, like a spec and then coding guideline. Try to have every one of the file try to mention as actionable items that the AI agent must follow. Okay. As a checklist. And what is acceptance criteria at the end that needs to be mentioned on each one of the specs for coding guideline inside, like the style guideline, Boolean guideline, naming convention, things like that. Okay. Do you understand? Can you please follow through? Do you have any question and confusion?

# Actionable Items Must Follow Non-Negotiable

1. Review the coding guidelines in the README.md file within the coding guideline folder.
2. Create a specification and coding guideline document.
3. Ensure each file includes actionable items for AI agents as a checklist.
4. Define acceptance criteria for each spec, including style guidelines, Boolean guidelines, and naming conventions.

Must follow and spawn agent using 

@[.agents/skills/execute-parent-task-with-n-steps-v6]

## Additional Instructions

learn /learn if you have to learn something and /plan stuff before working please./plan
```

---

## 2. Completed Subtasks & Evidence Matrix

| Wave | Task-ID | Worker | Owned Files | Status | Evidence |
|:---|:---|:---|:---|:---|:---|
| **Phase 1** | Task-01 | Lead | `02-spec/21-app/04-coding-guidelines-actionable-checklists-and-acceptance-criteria/*` | DONE | Created architecture spec, rollout spec, registered in 21-app readme |
| **Wave 1** | Task-02A | Worker 01 | `01-cross-language/02-boolean-principles/*` (6 files) | DONE | PASS exit 0, 6 files upgraded with AI checklists and AC-CG-BOOL-001..006 |
| **Wave 1** | Task-02B | Worker 02 | `01-cross-language/12-no-negatives.md`, `24-boolean-flag-methods.md` | DONE | PASS exit 0, 2 files upgraded with AI checklists and AC-CG-BOOL-012/024 |
| **Wave 2** | Task-03A | Worker 01 | `01-cross-language/04-code-style/02-05` (4 files) + readme | DONE | PASS exit 0, 5 files upgraded with AI checklists and AC-CG-STYLE-001..005 |
| **Wave 2** | Task-03B | Worker 02 | `01-cross-language/04-code-style/06-08`, `21-newline...` | DONE | PASS exit 0, 4 files upgraded with AI checklists and AC-CG-STYLE-005..008 |
| **Wave 3** | Task-04A | Worker 01 | `01-cross-language/10`, `11`, `22` | DONE | PASS exit 0, 3 files upgraded with AI checklists and AC-CG-NAME-010..022 |
| **Wave 3** | Task-04B | Worker 02 | `01-cross-language/07`, `28` | DONE | PASS exit 0, 2 files upgraded with AI checklists and AC-CG-NAME-007/028 |
| **Wave 4** | Task-05A | Worker 01 | `01-cross-language/13`, `18`, `32` | DONE | PASS exit 0, 3 files upgraded with AI checklists and AC-CG-TYPE-013..032 |
| **Wave 4** | Task-05B | Worker 02 | `01-cross-language/33`, `34` | DONE | PASS exit 0, 2 files upgraded with AI checklists and AC-CG-TYPE-033/034 |
| **Wave 5** | Task-06A | Worker 01 | `01-cross-language/03`, `06`, `08` | DONE | PASS exit 0, 3 files upgraded with AI checklists and AC-CG-ARCH-006..008 |
| **Wave 5** | Task-06B | Worker 02 | `01-cross-language/16`, `19`, `20`, `23` | DONE | PASS exit 0, 4 files upgraded with AI checklists and AC-CG-ARCH-016..023 |
| **Wave 6** | Task-07 | Lead | `01-cross-language/97-acceptance-criteria.md`, `97-acceptance-criteria.md`, readmes | DONE | PASS exit 0, consolidated Acceptance Criteria registry created |
| **Wave 7A** | Task-08A | Worker 01 | `03-golang/readme.md`, `02-boolean-standards.md`, `03-httpmethod-enum.md`, `05-defer-rules.md` | DONE | PASS exit 0, 4 files upgraded with AI checklists and AC-CG-GO-000..005 |
| **Wave 7B** | Task-08B | Worker 02 | `03-golang/06-string-slice-internals.md`, `07-code-severity-taxonomy.md`, `08-pathutil-fileutil-spec.md`, `09-wrapped-boolean-results.md` | DONE | PASS exit 0, 4 files upgraded with AI checklists and AC-CG-GO-006..009 |
| **Wave 7C** | Task-08C | Lead | `03-golang/97-acceptance-criteria.md`, `02-coding-guidelines/97-acceptance-criteria.md` | DONE | PASS exit 0, Go Acceptance Criteria registry and root criteria synchronized |
| **Wave 8A** | Task-11A | Worker 01 | `02-canonical-size-tier.md`, `03-coding-style-checklist.md`, `04-consolidated-review-guide-condensed.md`, `05-consolidated-review-guide.md` | DONE | PASS exit 0, 4 root style files upgraded with AI checklists and AC-CG-STYLE-002..005 |
| **Wave 8B** | Task-11B | Worker 02 | `03-golang/01-enum-specification/*` (6 files) | DONE | PASS exit 0, 6 Go enum files upgraded with AI checklists and AC-CG-GO-ENUM-000..006 |
| **Wave 9A** | Task-12A | Worker 01 | `03-golang/04-golang-standards-reference/*` (7 files) | DONE | PASS exit 0, 7 Go ref files upgraded with AI checklists and AC-CG-GO-REF-000..007 |
| **Wave 9B** | Task-12B | Worker 02 | `02-typescript/readme.md`, `02..07`, `11` (8 files) | DONE | PASS exit 0, 8 TS enum specs upgraded with AI checklists and AC-CG-TS-000..011 |
| **Wave 10A** | Task-13A | Worker 01 | `02-typescript/08..10`, `12..15` (7 files) | DONE | PASS exit 0, 7 TS advanced specs upgraded with AI checklists and AC-CG-TS-008..015 |
| **Wave 10B** | Task-13B | Worker 02 | `02-typescript/97-acceptance-criteria.md` | DONE | PASS exit 0, full TS AC registry created and synchronized |
| **Wave 11A** | Task-14A | Worker 01 | `04-php/readme.md`, `02..05` (5 files) | DONE | PASS exit 0, 5 PHP core specs upgraded with AI checklists and AC-CG-PHP-000..005 |
| **Wave 11B** | Task-14B | Worker 02 | `04-php/06..09`, `07-php-standards-reference/*`, `97-acceptance-criteria.md` | DONE | PASS exit 0, 11 PHP reference specs and AC registry upgraded |
| **Wave 12A** | Task-15A | Worker 01 | `05-rust/*` (8 files) | DONE | PASS exit 0, 8 Rust specs upgraded with AI checklists and AC-CG-RUST-000..007 |
| **Wave 12B** | Task-15B | Worker 02 | `07-csharp/*` (6 files) | DONE | PASS exit 0, 6 C# specs upgraded with AI checklists and AC-CG-CS-000..005 |
| **Wave 13A** | Task-16A | Worker 01 | `06-ai-optimization/02..05` (4 files) | DONE | PASS exit 0, 4 AI optimization specs upgraded with AI checklists and AC-CG-AI-002..005 |
| **Wave 13B** | Task-16B | Worker 02 | `06-ai-optimization/06..09`, `readme.md`, `97-acceptance-criteria.md` | DONE | PASS exit 0, 6 AI optimization specs and AC registry upgraded |
| **Wave 14A** | Task-17A | Worker 01 | `06-cicd-integration/02..08`, `readme.md`, `97-acceptance-criteria.md` | DONE | PASS exit 0, 9 CI/CD integration specs and AC registry upgraded |
| **Wave 14B** | Task-17B | Worker 02 | `06-cicd-integration/08-fix-repo-and-installers/*` (5 files) | DONE | PASS exit 0, 5 Fix Repo and installer specs upgraded |
| **Wave 15A** | Task-18A | Worker 01 | `08-file-folder-naming/*` (6 files) | DONE | PASS exit 0, 6 File & folder naming specs upgraded with AI checklists and AC-CG-FILE-000..006 |
| **Wave 15B** | Task-18B | Worker 02 | `09-powershell-integration/`, `10-research/`, `12-python/`, `13-cpp/` (6 files) | DONE | PASS exit 0, 6 Polyglot specs upgraded with AI checklists and AC |
| **Wave 16A** | Task-19A | Worker 01 | `11-security/*` (8 files) | DONE | PASS exit 0, 8 Security specs upgraded with AI checklists and AC-CG-SEC-000..004, AC-CG-SEC-AXIOS-000..002 |
| **Wave 16B** | Task-19B | Worker 02 | `21-app/`, `22-app-issues/`, `23-app-db/`, `24-app-ui-design-system/`, `01-cross-language/` stubs (10 files) | DONE | PASS exit 0, 10 App Readmes and cross-language stubs upgraded with AI checklists and AC |
| **Wave 17A** | Task-20A | Worker 01 | `11-security/97-acceptance-criteria.md`, `06-cicd-integration/98-faq.md`, `99-troubleshooting.md`, `08-fix-repo-and-installers/98-faq.md` | DONE | PASS exit 0, Security AC registry created and CI/CD FAQ/Troubleshooting upgraded |
| **Wave 17B** | Task-20B | Worker 02 | `01-cross-language/readme.md`, `99-consistency-report.md` | DONE | PASS exit 0, Cross-language readme AC block and root consistency report upgraded |
| **Wave 18** | Task-21 | Lead | `02-coding-guidelines/97-acceptance-criteria.md`, `01-cross-language/16-static-analysis/97-acceptance-criteria.md` | DONE | PASS exit 0, Master AC registry updated with all language and module registries |

---

## 3. Verification & Compliance Evidence

1. **Relative Paths:** `python linter-scripts/check-relative-paths.py` -> exit 0 (0 absolute paths or `file:///` URIs across 3504 tracked files).
2. **Forbidden Strings & Secrets:** `python linter-scripts/check-forbidden-strings.py` -> exit 0 (zero secrets, zero forbidden patterns).
3. **Guideline Formatting across all 223 files:** `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only` -> exit 0 (All 223 files conform).
4. **SQLite Concurrency & Action Logging:** `python 03-ai-scripts/46-agent-sqlite-task-manager.py status --db .ai-memory/temp-agents/04-coding-guidelines-actionable-checklists/agent-task.db` -> exit 0 (All 38 subtasks marked DONE).



