# Coding Guidelines (AI Execution Prompt)

**Version:** 3.3.0  
**Status:** Active  
**Updated:** 2026-08-28  
**AI Confidence:** Production-Ready  

## 🤖 MUST FOLLOW INSTRUCTIONS FOR ALL AI AGENTS

> **CRITICAL DIRECTIVE**: You are bound by this document. Before generating any code, writing any script, or modifying any architecture, you **MUST** internalize and apply these rules. Failure to apply these rules will result in an immediate rejection of your code.

### 1. No Generated Code or Artifacts (Never Commit)

Never commit generated code (e.g., ORM models, gRPC clients), test results, test reports, or compiled binaries. They belong in build artifacts or CI, never in source control.

### 2. Error Management is the #1 Priority

Error handling must be implemented from the **very first line of code**. Never write business logic without proper error handling wrapping it. Use the \AppError\ / \AppException\ architecture explicitly defined in the \spec/03-error-manage/\ folder. This is non-negotiable.

### 3. Boolean Naming (Strict Positive Assertion)

All booleans **MUST** use \is\, \has\, \can\, or \should\ prefixes and are **positively named only**.

- ❌ BAD: \!isSuccess\, \isDisabled\
- ✅ GOOD: \isFail\, \isActive\
Extract multi-part conditions into well-named boolean variables.

### 4. Nesting and Flow Control

Zero nesting. Use early returns and guard clauses. No nested \if\ blocks. If you find yourself nesting, extract the logic into a separate function immediately.

### 5. Semantic Naming (No Generics)

Absolutely NO generic garbage names. Variables named \	emp\, \data\, \obj\, \comp_100\ will trigger an instant rejection. All unit tests must be behavior-driven (e.g., \TestUpdateUser_RejectsInvalidEmail\).

### 6. Function Metrics & Signatures

- Functions: 8-15 lines. Files: < 300 lines. React components: < 100 lines.
- **Maximum 3 Parameters:** See the strict formatting rules in \coding-style-checklist.md\.

### 7. Never Hallucinate

If a requirement is unclear or missing, **ask a clarifying question** instead of guessing. Wrong assumptions cause rewrites.

---

## 🏗 Directory Index

- [Cross-Language Standards](./01-cross-language/00-overview.md)
- [TypeScript Guidelines](./02-typescript/00-overview.md)
- [Go Guidelines](./03-golang/00-overview.md)
- [C# Guidelines](./07-csharp/00-overview.md)
- [AI Optimization](./06-ai-optimization/00-overview.md)
- **[Coding Style Checklist](./coding-style-checklist.md)**
