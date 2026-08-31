# Plan: Trailing Newlines & AI Fix Scripts

**Status:** PENDING

## Problem Statement
Various files in the repository (such as recent markdown plans) were created without a trailing newline (\n) at the end of the file. POSIX standard and most linters require a trailing newline. Additionally, there is a need to build a reusable library of AI helper scripts in .lovable/ai-fix-scripts/ that can automatically rectify codebase-wide issues like this.

## Acceptance Criteria
1. **AI Fix Scripts Directory:** Establish .lovable/ai-fix-scripts/ with an index.md documenting available scripts.
2. **Newline Fixer Script:** Create a robust script
ewline_fixer.py (or Node equivalent) inside the fix scripts directory. It must scan source code, text, and markdown files (.md, .txt, .go, .ts, .js, .mjs, .cjs, .jsx, .cs, .vb, .rs, etc.) and ensure they end with exactly one \n. It must respect the UTF-8 without BOM requirement.
3. **Configuration:** Update .editorconfig (or Prettier/Linter configs) to enforce insert_final_newline = true.
4. **Guideline Update:** Update spec/02-coding-guidelines/03-03-coding-style-checklist.md to explicitly require trailing newlines.
5. **Execution:** Run the fixer script across the repository.

## Subtasks
- 01-create-fix-scripts-repo.md: Scaffold .lovable/ai-fix-scripts/ and index.md.
- 02-write-newline-fixer.md: Author
ewline_fixer.py and execute it.
- 03-update-tool-configs.md: Configure Prettier/EditorConfig.
- 04-update-guidelines.md: Update 03-coding-style-checklist.md.
