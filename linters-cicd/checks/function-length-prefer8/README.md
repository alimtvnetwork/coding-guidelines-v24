# CODE-RED-005 — Function Length (Strict 8)

> **Status:** `error` (build-failing) · **Spec source of truth:** this file
> **Owners:** `linters-cicd/checks/function-length-prefer8/` + `eslint-plugins/coding-guidelines/index.js#preferFunctionLines`
> **Last reviewed:** 2026-04-27

This is the **single, canonical** specification for how the strict 8-line
function-length rule is measured and enforced. All other surfaces
(`spec/.../06-rules-mapping.md`, `.lovable/coding-guidelines/coding-guidelines.md`,
the root `readme.md`, the rule docstrings) link **here**. If you change
the threshold or counting rules, update this file first; the others must
follow.

---

## 1. Threshold (the only number that matters)

```
STRICT_LINES = 8     # owned by CODE-RED-005 — binding cap
LEGACY_HARD  = 15    # owned by CODE-RED-004 — redundant safety net
```

A function is **in violation** when its *effective body line count*
(see §2) is **strictly greater than 8**.

| Effective lines | CODE-RED-005 | CODE-RED-004 | Build outcome |
|----------------:|--------------|--------------|---------------|
| 0 – 8           | silent       | silent       | ✅ pass       |
| 9 – 15          | **error**    | silent       | ❌ fail       |
| 16+             | **error**    | error (redundant) | ❌ fail  |

There is no warning band. There is no soft tier. `9` fails the build
just as hard as `9000`.

---

## 2. What counts as an "effective body line"

Every scanner counts lines **inside the function body** — i.e. **between**
(not including) the opening `{` / `def:` and the matching close. The
function signature line, the closing brace/dedent, and lines outside the
body are **never** counted.

A line in the body **is counted** iff, after `.strip()`, it is **not** one
of the following:

| Skip rule              | TS / JS / Go / PHP / Rust | Python |
|------------------------|:-------------------------:|:------:|
| Empty / whitespace-only| ✅ skip                   | ✅ skip |
| Single-line comment    | `//` line                 | `#` line |
| Block comment open     | line starting with `/*`   | n/a    |
| Inside open `/* … */`  | every line until `*/`     | n/a    |
| JSDoc continuation     | line starting with `*`, or exactly `/*` / `*/` (ESLint counter only) | n/a |
| Pure docstring line    | n/a                       | line starting with `"""` or `'''` |

**Everything else counts as one line — including:**

- A statement spread across multiple physical lines (each physical
  line counts; the rule is line-based, not statement-based, by design —
  this discourages "hide-the-statement" reformatting to dodge the cap).
- Lines that are *only* a `}` or `)` or `];` if they sit on their own.
- `return` statements, `await` lines, decorators inside the body.
- Lines inside nested functions / closures (the outer function carries
  the cost of bodies it lexically contains).

**Caveats and known counter divergence** (intentional, do not "fix"
without bumping this doc):

- The ESLint counter (`countEffectiveBodyLines` in
  `eslint-plugins/coding-guidelines/index.js`) skips lines starting
  with `*` to keep JSDoc continuations free. The Go/PHP Python
  counters track block-comment state instead. Both converge on the
  same answer for ordinary code; they only differ for hand-rolled
  banner comments inside a body, which is a non-goal.
- The Python scanner skips a line that *starts with* `"""` or `'''`,
  not full multi-line docstring tracking. Multi-line docstrings whose
  middle lines are plain prose **will count**. This is acceptable
  because docstrings belong on the function, not buried mid-body.
- Rust handles `//`, `///`, and `/* … */` the same way as Go.

---

## 3. What "function" means per language

The rule fires on every construct that has a body:

| Language   | Construct(s)                                       | Detector |
|------------|----------------------------------------------------|----------|
| TypeScript / JavaScript | `FunctionDeclaration`, `FunctionExpression`, `ArrowFunctionExpression` (with `BlockStatement` body), class methods | ESLint AST |
| Go         | `func Name(...) {`, `func (r Recv) Name(...) {`    | regex + brace balance |
| PHP        | `function name(...) {`, modifier-prefixed methods  | regex + brace balance |
| Python     | `def name(...):`, `async def name(...):`           | regex + indentation |
| Rust       | `fn name(...) [-> T] {` (incl. `pub`, `async`, `unsafe`, `const`) | regex + brace balance |

Arrow functions with **expression bodies** (no `{ }`) are not measured —
they are by definition one expression and cannot exceed the cap.

---

## 4. Scope: what the rule does NOT do

- **No length-of-file aggregation.** That is CODE-RED-006
  (`checks/file-length/`).
- **No cyclomatic complexity.** Lines are a proxy, not the metric of
  record. A 6-line function with 4 nested ternaries still passes
  CODE-RED-005 (and is then caught by CODE-RED-001 / no-nested-if).
- **No automatic refactor.** The rule reports; the human decides
  whether to extract, inline, or waiver.

---

## 5. Waivers

There is **no waiver mechanism** for CODE-RED-005 today. A future
`// CODE-RED-005-WAIVER: <reason>` pattern is reserved but not
implemented; introducing it requires updating §1 of this file and the
mapping table in `spec/02-coding-guidelines/06-cicd-integration/06-rules-mapping.md`.

If you need to ship code that exceeds 8 lines today and refactor is
not viable, the supported escape hatches are:

1. Extract a helper (preferred — usually trivial).
2. Disable the rule on the specific line via the underlying linter's
   own disable comment (e.g.
   `// eslint-disable-next-line coding-guidelines/prefer-function-lines`).
   Each such disable is a debt marker and should be tracked.

The CI scanners (Python-driven, in `linters-cicd/checks/`) do **not**
honour ESLint disable comments. They will still emit a SARIF `error`
for the same function. This is intentional — the strict cap is global.

---

## 6. Verifying a change to this rule

If you modify any counter, threshold, or detector, you **must** re-run:

```bash
python3 -m pytest linters-cicd/tests/test_prefer8_fires_on_fixture.py -v
python3 -m pytest linters-cicd/tests/test_function_length_*.py -v
npm run lint -- --no-warn-ignored
```

The fixture at
`linters-cicd/tests/fixtures/code-red-005/too-long.ts` is calibrated to
**11 effective lines**. Tests assert it triggers `CODE-RED-005` at
`error` level and that `CODE-RED-004` stays silent. Do not "fix" the
fixture to 8 lines — it is the canary.

---

## 7. Cross-references (must agree with this file)

- `spec/02-coding-guidelines/06-cicd-integration/06-rules-mapping.md` —
  registry-level mapping table.
- `.lovable/coding-guidelines/coding-guidelines.md` — Rule #1
  developer-facing summary.
- `linters-cicd/checks/function-length-prefer8/_shared.py` — `RULE`
  metadata + `STRICT_LINES` constant (the **executable** source of
  truth; this README is the **prose** source of truth — if they
  disagree, treat it as a P1 bug and fix here first).
- `eslint-plugins/coding-guidelines/index.js` — `preferFunctionLines`
  rule + `countEffectiveBodyLines` helper.
- Root `readme.md` — links here from the "Coding standards" section.
