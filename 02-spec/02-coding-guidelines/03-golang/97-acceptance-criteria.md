# Golang Standards — Acceptance Criteria Registry (AI Execution Prompt)

> **/goal** Provide a consolidated, traceable registry of testable acceptance criteria across all Go coding guideline specifications in `03-golang/`.
> **/learn** Enforce the canonical criteria taxonomy (`AC-CG-GO-[NUM]`), positive boolean rules (`is`/`has` only), `*appfault.AppError` return types, and verify compliance using targeted linters.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify each criterion maps 1:1 to an authoritative specification file in `03-golang/`.
- [ ] `/learn` Ensure all criteria follow the structured `Given / When / Then` verification contract.
- [ ] `/goal` Validate that all verification commands execute cleanly with `Expected: exit 0`.
- [ ] `/learn` Verify 100% relative paths and zero absolute filesystem paths.

. **CRITICAL AI INSTRUCTION:** This specification is an active AI execution directive. All code generated or modified must strictly follow the rules below.

**Version:** 4.0.0
**Last Updated:** 2026-10-02
**Status:** Active
**AI Confidence:** Production-Ready

---

## 1. Golang Criteria Inventory (`AC-CG-GO-`)

| ID | Title | Authoritative Specification | Verification Command |
|:---|:---|:---|:---|
| `AC-CG-GO-000` | Golang Coding Standards Index & Specification Conformance | [`readme.md`](readme.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/03-golang --check-only` |
| `AC-CG-GO-001` | Go Positive Boolean Naming, Negation Elimination, and Guard Clauses | [`02-boolean-standards.md`](02-boolean-standards.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/03-golang/02-boolean-standards.md --check-only` |
| `AC-CG-GO-003` | HttpMethod Enum Standard and Magic String Elimination | [`03-httpmethod-enum.md`](03-httpmethod-enum.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/03-golang/03-httpmethod-enum.md --check-only` |
| `AC-CG-GO-005` | Resource Defer Rules and Loop Safety Standards | [`05-defer-rules.md`](05-defer-rules.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/03-golang/05-defer-rules.md --check-only` |
| `AC-CG-GO-006` | Go String and Slice Preallocation & Memory Efficiency Standards | [`06-string-slice-internals.md`](06-string-slice-internals.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/03-golang/06-string-slice-internals.md --check-only` |
| `AC-CG-GO-007` | Code Severity Taxonomy, Error Classification, and Fault Logging | [`07-code-severity-taxonomy.md`](07-code-severity-taxonomy.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/03-golang/07-code-severity-taxonomy.md --check-only` |
| `AC-CG-GO-008` | Unified PathUtil & FileUtil Cross-Platform Specification | [`08-pathutil-fileutil-spec.md`](08-pathutil-fileutil-spec.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/03-golang/08-pathutil-fileutil-spec.md --check-only` |
| `AC-CG-GO-009` | Wrapped Result Monad and Single Return Parameter Standards | [`09-wrapped-boolean-results.md`](09-wrapped-boolean-results.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/03-golang/09-wrapped-boolean-results.md --check-only` |

---

## Cross-References

- [Golang Coding Standards Overview](./readme.md)
- [Master Coding Guidelines Acceptance Criteria](../97-acceptance-criteria.md)
- [Cross-Language Standards](../01-cross-language/readme.md)
- [Cross-Language Acceptance Criteria](../01-cross-language/97-acceptance-criteria.md)
