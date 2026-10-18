# Component Specification: MIME & Media Type Enum Architecture

## 1. Overview & Deliverables Inventory
This specification defines the exact component updates, prompt enhancements, skill revisions, and inspection matrices required to enforce the MIME and Media Type Enum standard across the prompt architecture and fleet repositories.

### Primary Deliverables
1. **Cross-Language Coding Guideline:**
   - Path: `02-spec/02-coding-guidelines/01-cross-language/36-mime-and-media-type-enum-standards.md`
   - Scope: Universal multi-language patterns (TS, Go, Python, Rust, PHP), lookup dispatch tables, Blob/Clipboard/Header patterns.
2. **Coding Guideline Execute Prompts:**
   - Paths:
     - `01-prompts/15-cg-execute/06-constants-and-enums.md`
     - `01-prompts/v4/15-cg-execute/06-constants-and-enums.md`
   - Updates: Add explicit MIME/Media type enum requirement, Blob/ClipboardItem inspection checklists, and HTTP header refactoring rules.
3. **Companion Skills:**
   - Paths:
     - `.agents/skills/cg-enum-standards/skill.md`
     - `.cursor/skills/cg-enum-standards/skill.md`
     - `.agents/skills/multi-language-enums-and-traits/skill.md`
     - `.cursor/skills/multi-language-enums-and-traits/skill.md`
   - Updates: Include MIME type enum detection patterns, Blob/Clipboard inspection guidelines, and file extension mapping rules.
4. **Fleet Synchronization:**
   - Script: `03-ai-scripts/38-sync-prompts-skills-scripts.py`
   - Scope: Propagate updated prompts and skills across all connected repositories in `d:\work`.

---

## 2. Component Design & Modifications

### A. CG Execute Prompt: `06-constants-and-enums.md`
The prompt instructs AI agents during coding guideline audits and refactorings. The following explicit section MUST be embedded:
- **MIME & Media Type Enum Rule:**
  - Scan for magic string literals representing MIME/Media types (`'text/html'`, `'text/plain'`, `'application/json'`, etc.) in `new Blob()`, `new ClipboardItem()`, headers, and response setters.
  - Require creating or importing a centralized `MimeType` enum / typed constant dictionary.
  - Refactor all call sites to use `MimeType.<Member>`.
  - Ban loose file extension to MIME `if-else` branching; mandate map dispatch tables (`EXTENSION_MIME_MAP`).

### B. Antigravity & Cursor Skills: `cg-enum-standards`
Add dedicated pattern guidance for MIME and Media Types:
- Detection regex patterns for magic strings.
- Implementation guide for TypeScript `as const` object enums and Go/Python string enums.
- Clipboard and Blob API refactoring examples.

### C. Antigravity & Cursor Skills: `multi-language-enums-and-traits`
Add cross-language reference tables for `MimeType` across Go, TypeScript, Python, Rust, and PHP.

---

## 3. Linter & Static Analysis Signatures

### Detection Rules for Magic Strings
- Regex for Blob constructors with raw MIME strings:
  `new\s+Blob\s*\(\s*\[.*?\],\s*\{\s*type:\s*['"`](text|application|image|audio|video|multipart)\/[a-zA-Z0-9.+_-]+['"`]\s*\}\s*\)`
- Regex for ClipboardItem constructors with raw MIME strings:
  `new\s+ClipboardItem\s*\(\s*\{\s*['"`](text|application|image)\/[a-zA-Z0-9.+_-]+['"`]\s*:`
- Regex for HTTP Header `Content-Type` with raw MIME strings:
  `(?:setHeader|Set|headers\.append)\s*\(\s*['"`](?:Content-Type|Accept)['"`]\s*,\s*['"`](text|application|image)\/[a-zA-Z0-9.+_-]+['"`]\s*\)`
