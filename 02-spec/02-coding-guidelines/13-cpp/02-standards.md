# Modern C++ Standards (AI Execution Prompt)

> **/goal** Define mandatory modern C++ (C++20 baseline) programming standards, smart pointer memory safety, RAII ownership, and naming conventions.
> **/learn** Master C++20 concepts, ban manual new/delete, enforce Rule of Zero/Rule of Five, and prevent exception leakage across FFI boundaries.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Target C++20 standard baseline; use concepts rather than raw `enable_if` for template constraints.
- [ ] `/learn` Never use raw `new` or `delete`; always use `std::unique_ptr` or `std::shared_ptr` under strict RAII.
- [ ] `/goal` Adhere to the Rule of Zero or explicitly define/delete all five special member functions (Rule of Five).
- [ ] `/learn` Ensure exceptions do not escape C++ boundaries into C/FFI layers and avoid exceptions for control flow.

. **CRITICAL AI INSTRUCTION:** This specification is an active AI execution directive. All code generated or modified must strictly follow the rules below.

## 1. C++20 Baseline

- All code MUST be written against the C++20 standard or newer.
- Use concepts instead of raw `enable_if` for template constraints.

## 2. Memory Safety & RAII

- **Never use `new` or `delete` manually.** Always use `std::unique_ptr` or `std::shared_ptr`.
- Follow the **Rule of Five** (or Rule of Zero preferred). If a class defines a custom destructor, copy constructor, move constructor, copy assignment, or move assignment, it should explicitly define or delete all five.
- Resources must be tied to object lifecycles (RAII).

## 3. Naming Conventions

- Structs, Classes, and Enums: `PascalCase`
- Functions and Variables: `snake_case` (follows community standards similar to Rust)
- Macros: `SCREAMING_SNAKE_CASE` (avoid macros where `constexpr` can be used)

## 4. Exceptions

- Avoid exceptions for control flow.
- When crossing C/C++ boundaries, ensure exceptions do not escape C++ code.

---

## Verification & Acceptance Criteria

_Auto-generated section — see `02-spec/02-coding-guidelines/97-acceptance-criteria.md` for the full criteria index._

### AC-CG-CPP-002: Modern C++ Standards Conformance

**Given** Polyglot development guidelines and language standards.
**When** Audited against this language specification.
**Then** Zero non-compliant conventions or syntax patterns are detected and exit code is 0.

**Verification command:**
```bash
python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only
```
**Expected:** exit 0. Zero violations.
