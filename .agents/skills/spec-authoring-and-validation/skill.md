---
name: spec-authoring-and-validation
description: Author, structure, sequence, and validate repository specifications adhering to 02-spec/01-spec-authoring-guide/.
---

# Specification Authoring & Validation Guide

This skill governs the creation, organization, and automated validation of architectural specifications in `02-spec/`.

## Structure & File Naming Conventions

1. **Folder Naming:**
   - Folders follow the hyphenated two-digit sequence pattern: `02-spec/<NN>-<slug>/` (e.g. `02-spec/02-coding-guidelines/`, `02-spec/21-app/`).

2. **Mandatory Files per Spec Folder:**
   - `01-index.md`: Primary entry point explaining scope, version, goal, and learn checklists.
   - Numbered markdown files: Detailed topic-specific policies.
   - `97-acceptance-criteria.md`: Verification commands and criteria.
   - `98-changelog.md`: Evolution history of the specification.
   - `99-consistency-report.md`: Audit log verifying alignment with global rules.

3. **Strict Path Rules:**
   - All internal links must use relative paths starting from the repository root or relative markdown paths.
   - Never write absolute filesystem paths or `file:///` URIs.

4. **Validation Checklist:**
   - Run spec cross-link validation:
     ```bash
     python linter-scripts/check-spec-cross-links.py --root 02-spec --repo-root .
     ```
   - Ensure header spacing and markdown gap linters pass.
