# Instruction (must follow): Enforce Markdown Linting Rules (MD022 & MD032)

/goal Enforce strict Markdown formatting rules across this codebase, specifically targeting heading and list spacing, to ensure the documentation renders perfectly on GitHub and other platforms. Avoid hallucinations by verifying rules against actual files and ensuring no steps are missed.

## 1. Configure the Linter
If the repository does not have a `.markdownlint.json` file, create it in the root directory. If it does exist, update it to explicitly enable rules `MD022` and `MD032`:
```json
{
  "default": false,
  "MD022": true,
  "MD032": true
}
```

## 2. The Spacing Rules (What these rules mean)
- **MD022 (Heading Spacing):** Every Markdown heading (`#`, `##`, `###`, etc.) MUST be surrounded by a completely blank line directly above and below it.
- **MD032 (List Spacing):** Every list or checklist (`- `, `- [ ]`, `1. `) MUST be surrounded by blank lines. If there is a paragraph, title, or header immediately preceding a list, you MUST insert a blank line before the first list item.

## 3. Execution
1. Install or run the markdownlinter CLI (e.g., `npx markdownlint-cli "**/*.md" --fix`).
2. If the linter cannot auto-fix everything, you must manually sweep the `.md` files and ensure there are no lists glued directly to the bottom of a paragraph or header.
3. Commit the configuration and the formatting fixes to the repository.

## Anti-Hallucination & Compliance Checklist
Before finalizing this turn, you MUST manually verify the following to prevent hallucination and ignored instructions:
- [ ] **File Presence:** Did I verify that the files I am referencing actually exist on disk? (No guessing).
- [ ] **README Version Pinning:** If a release occurred, did I verify that the root `readme.md` has the latest version strictly pinned?
- [ ] **Guideline Compliance:** Did I adhere strictly to the coding guidelines, including Markdown rules (MD022/MD032) and the parameter-per-line formatting rule?
