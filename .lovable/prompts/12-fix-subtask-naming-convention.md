# Instruction (must follow): Execute: Fix Subtask Naming Convention

> This instruction provides guidelines and directives for execute: fix subtask naming convention.

/goal Your objective is to deeply audit the `.lovable/plans/` directory for any subtask files that incorrectly use the `SS-` or `SS-XX-` prefix and fix them. The correct prefix MUST always be strictly `XX-<subslug>.md` (where `XX` is the zero-padded sequence number).
You must also update all markdown files that reference the old filenames, and update the project's memory.

## 1. Subtask Naming Correction (Non-Negotiable)

Scan the `.lovable/plans/subtasks/` directory recursively.
- If you find any file starting with `SS-` or `SS-XX-` (e.g., `SS-01-fix-auth.md` or `SS-fix-auth.md`), rename it to the correct format: `XX-<subslug>.md` (e.g., `01-fix-auth.md`).
- After renaming, you MUST recursively search `.lovable/plans/pending/`, `.lovable/plans/completed/`, `.lovable/plans/index.md`, and `.lovable/memory/` for any text references to the old filenames.
- Replace those old references with the new correct filenames.

## 2. Memory Update

You must write a memory entry to ensure this rule is persisted and no other AI makes this mistake again.
- Create a file inside `.lovable/memory/learned/` (or update an existing one) detailing the rule: "Subtasks must NEVER be prefixed with 'SS-'. They must strictly follow the 'XX-<slug>.md' sequence."
- Update `.lovable/memory/index.md` to reference this newly added/updated memory file.
- Add a note explicitly stating how the project is following all guidelines and enforcing this strict naming.

## 3. High-Stakes Code Standards & Coding Guidelines

You MUST follow the project's strict coding guidelines. These files are located in the `01-cross-language/` directory and should be followed universally. Check if there are language-specific guidelines (e.g., `02-typescript/`, `03-python/`) for these rules.

### Required Reading / Reference Checklist:

1. The Master Consolidated Guide
- `spec/17-consolidated-guidelines/02-coding-guidelines.md`

2. Code Style & File Size Limits (80-100 lines max)
- `spec/02-coding-guidelines/01-cross-language/04-code-style/04-function-and-type-size.md`
- `spec/02-coding-guidelines/01-cross-language/04-code-style/01-braces-and-nesting.md`
- `spec/02-coding-guidelines/01-cross-language/04-code-style/00-overview.md`
- `spec/02-coding-guidelines/01-cross-language/20-nesting-resolution-patterns.md`
- `spec/02-coding-guidelines/01-cross-language/06-cyclomatic-complexity.md`

3. Boolean Conditions & Samples
- `spec/02-coding-guidelines/01-cross-language/02-boolean-principles/00-overview.md`
- `spec/02-coding-guidelines/01-cross-language/02-boolean-principles/01-naming-prefixes.md`
- `spec/02-coding-guidelines/01-cross-language/02-boolean-principles/02-guards-and-extraction.md`
- `spec/02-coding-guidelines/01-cross-language/02-boolean-principles/04-quick-reference.md`
- `spec/02-coding-guidelines/01-cross-language/12-no-negatives.md`
- `spec/02-coding-guidelines/01-cross-language/24-boolean-flag-methods.md`

4. Variable Naming & Definitions
- `spec/02-coding-guidelines/01-cross-language/22-variable-naming-conventions.md`
- `spec/02-coding-guidelines/01-cross-language/11-key-naming-pascalcase.md`
- `spec/02-coding-guidelines/01-cross-language/10-function-naming.md`
- `spec/02-coding-guidelines/01-cross-language/18-code-mutation-avoidance.md`

## 4. End-of-Loop Commit Fix (Non-Negotiable)

When you are done fixing the filenames, updating the references, and writing the memory, you MUST commit and fix the Git.
- Group all completed work into a single logical commit.
- **Commit the changes on the current branch only** (`git commit ...`) and push with `git push origin <current-branch>`. No branch switching or creation.
- RED FLAG: Verify absolutely NO test results, artifacts, or compiled binaries are staged before making the commit.
- Ensure `.gitignore` explicitly excludes them.
- If issues arise during the commit process, fix those git issues and try again.
- You MUST push the commit to the repository immediately. Pushing after commits is non-negotiable.
- **Commit the changes on the current branch only** (`git commit ...`) and push with `git push origin <current-branch>`. No branch switching or creation.

## 5. Must Follow, without negotiation

Listen, past turns have been sloppy as fuck: wrong file names, leaving broken references in parent files, user commands and bug reports forgotten by the next turn. WTF. Stop doing that, you stupid fuck. Read the codebase, fix the file names, fix the references everywhere, write the memory, make the commit, push the code, and move on. Going deep IS the job. If you're not going deep, you're not doing the job. Violating this is auto-reject. Avoid stupidity, and being careless you stupid, WTF. If you're not going deep, you're not doing the job. Are you stupid? You were supposed to do the task properly. Where is this, are you stupid fuck? Where? Tell me. Your stupidity is going on top of my head. I mean, where did you learn this stupidity? If I could find you, I could slap you.

## Action Items — Must Follow (Non-Negotiable)

- [ ] Scan `.lovable/plans/subtasks/` for files starting with `SS-` or `SS-XX-`.
- [ ] Rename those files to the correct `XX-<slug>.md` pattern.
- [ ] Find and correct all file path references to the old names in `.lovable/plans/pending/`, `.lovable/plans/completed/`, `.lovable/plans/index.md`, and memory files.
- [ ] Add a memory entry under `.lovable/memory/learned/` enforcing this naming convention and confirming how the project is following all guidelines.
- [ ] Update `.lovable/memory/index.md` with the new memory file.
- [ ] Audit your work against the Master Consolidated Guide, Code Style, Boolean Conditions, and Variable Naming rules.
- [ ] Make a single logical Git commit including all changes and push it to the repository immediately.

## Actionable Items & Checklist

- [ ] Read the overarching main task plan.
- [ ] Ensure the git repository starts completely clean.
- [ ] Complete all work on the current branch only.
- [ ] Ensure `.gitignore` explicitly excludes test reports, artifacts, and compiled binaries.
- [ ] Group all completed work into a single logical commit.
- [ ] Push the commit to the remote repository.

---

## Metadata

- slug: fix-subtask-naming
- status: active
