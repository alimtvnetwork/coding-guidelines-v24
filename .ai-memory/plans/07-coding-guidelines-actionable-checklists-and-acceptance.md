# Plan: Coding Guidelines Actionable Checklists & Acceptance Criteria Standard (AI Execution Prompt)

> **/goal** Transform all coding guidelines across the repository into authoritative, active AI execution prompts featuring standardized agent checklists and testable Gherkin acceptance criteria.
> **/learn** Enforce the 4-part architectural anatomy, multi-agent wave partitioning (A = 2, H = 2), disjoint file boxes, worker git ban, and automated verification gates.

## 🎯 Actionable CI/CD & Agent Checklist

- [x] `/goal` Review the coding guidelines in `02-spec/02-coding-guidelines/readme.md` and verify architecture specifications.
- [x] `/learn` Harmonize application specifications in `02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/`.
- [x] `/goal` Resolve `AC-CG-STYLE-008` criteria collision by assigning `AC-CG-STYLE-009` to newline styling examples.
- [x] `/learn` Harmonize polyglot registries (Python, C++, Go, PHP, Rust, AI, C#, Security) with self-conformance criteria.
- [x] `/goal` Update Master Registry in `02-spec/02-coding-guidelines/97-acceptance-criteria.md` with application specs and table parity.
- [x] `/learn` Execute full verification suite (`05-guideline-autofixer.py`, `check-relative-paths.py`, `check-forbidden-strings.py`).

. **CRITICAL AI INSTRUCTION:** This completed plan records the architecture, execution proof, and verification evidence for coding guidelines actionable checklists and acceptance criteria standardization.

**Version:** 1.1.0  
**Updated:** 2026-10-03  
**Status:** Completed  
**AI Confidence:** Production-Ready  
**Ambiguity:** None  

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

## 2. Multi-Agent Execution Waves (A = 2, H = 2)

### Wave 1: Disjoint Execution Subtasks
- **Worker 1:** Subtask `01-app-spec-and-style-collision.md`
  - Targets: `02-spec/21-app/07-.../` (3 files) + `01-cross-language/` (2 files)
  - Outcome: Retagged `21-newline-styling-examples.md` to `AC-CG-STYLE-009`, normalized headings in `07-...`, and added discrete bash verification commands for all criteria in `02-component-spec.md`.
- **Worker 2:** Subtask `02-polyglot-and-master-registry.md`
  - Targets: `12-python/` (2 files), `13-cpp/` (2 files), 7 domain registries (7 files), Master Registry `97-acceptance-criteria.md` (1 file)
  - Outcome: Harmonized Python and C++ index criteria, added missing `*-REG-001` criteria across 7 domain registries, updated Master Registry tables with standardized 4-column layout, and indexed application spec `AC-15`.

---

## 3. Verification Commands & Evidence

1. **Coding Guideline Conformance Scanner:**
   ```bash
   python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only
   python 03-ai-scripts/05-guideline-autofixer.py 02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance --check-only
   ```
   - **Result:** Exit code 0 across all 199 files. Zero newline or boolean violations.

2. **Relative Path Linter:**
   ```bash
   python linter-scripts/check-relative-paths.py
   ```
   - **Result:** Exit code 0. Zero absolute paths or `file:///` URIs.

3. **Forbidden Strings Linter:**
   ```bash
   python linter-scripts/check-forbidden-strings.py
   ```
   - **Result:** Exit code 0. Zero forbidden strings detected.
