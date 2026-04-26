# Spec-Author Developer Experience — Deep Dive

**Author:** Lovable AI · **Date:** 2026-04-26
**Companion to:** `docs/guidelines-audit.md`

---

## Goal

Reduce the time from "I have a spec idea" to "CI is green and the doc is live in the viewer" from ~30 min to under 5 min, without lowering the quality bar set by `linter-scripts/validate-guidelines.py`.

---

## Current author flow (measured pain points)

| Step | Today | Friction |
|------|-------|----------|
| 1. Find the right module / number | Read `01-spec-authoring.md` (475 lines) → `19-gap-analysis.md` → `spec-index.md` | 3 hops, easy to pick a duplicate prefix |
| 2. Write the metadata header | Type from memory or copy from a sibling | Drift in version / date / scoring fields |
| 3. Write content | Free-form markdown | Catches `not`/`no` boolean names, missing alt text, broken links only at commit time |
| 4. Validate | `python linter-scripts/validate-guidelines.py` locally | 8–15 s feedback loop, no in-editor squiggles |
| 5. Sync derived artifacts | 4 commands in a specific order | If skipped → `Drift detected in version.json` in CI |
| 6. Open PR | CI re-runs validator + cross-link checker + spec-tree drift check | Same checks as step 4 — should be a no-op but often catches step-5 mistakes |

---

## Recommended improvements (ranked by leverage)

### 1. `spec/_template.md` (shipped in this PR)
Eliminates step 2 friction. Authors `cp` and edit. Already in repo.

### 2. `00-strictly-avoid-quickref.md` (shipped in this PR)
30-second read; eliminates the most common review-time rejections (boolean negatives, magic strings, swallowed errors, version drift).

### 3. VS Code task — one-button validate
Add `.vscode/tasks.json` entry that runs the validator + cross-link checker and surfaces output in the Problems panel:

```jsonc
{
  "label": "Spec: Validate",
  "type": "shell",
  "command": "python linter-scripts/validate-guidelines.py && python linter-scripts/check-spec-cross-links.py --root spec --repo-root .",
  "group": "test",
  "presentation": { "reveal": "silent", "panel": "dedicated" },
  "problemMatcher": []
}
```

Pair with a custom `problemMatcher` once the validator emits `file:line: message` lines, so violations show inline.

### 4. Pre-commit hook expansion
`scripts/hooks/pre-commit` already exists. Extend it to:

1. Detect changes under `spec/` → run validator
2. Detect any change → run sync scripts in order
3. Detect `version.json` / `specTree.json` diff post-sync → re-stage

This collapses step 5 of the current flow into a no-op: the author can't commit without the synced state.

### 5. Live preview wired into the docs viewer
The viewer already reads `src/data/specTree.json`. Add a dev-mode watch (or `bun run sync:watch`) that re-runs `sync-spec-tree.mjs` on `spec/**/*.md` save. Authors see their doc render in the live preview within ~1 s.

### 6. Frontmatter linter (specific to spec metadata)
Today the validator checks header presence but not content. Add micro-rules:

- `Version` is semver
- `Updated` is ISO `YYYY-MM-DD` and ≥ git's last-modified date
- `AI Confidence` ∈ {Production-Ready, High, Medium, Low, Draft}
- `Ambiguity` ∈ {None, Low, Medium, High, Critical}
- Exactly one H1
- Scoring table present when file is `00-overview.md`

Each rule is a 5-line Python check. Output format `path:line:col: SPEC-XXX message` so editors and CI both render it identically.

### 7. Auto-bump helper
`scripts/bump-spec.mjs <path> [patch|minor|major]` that:

- Bumps `Version` in the file's frontmatter
- Updates `Updated` to today
- Re-runs the three sync scripts
- Stages the diff

Removes the most common "I forgot to bump" PR comment.

### 8. Scaffold command
`scripts/new-spec.mjs <module-name> <file-name>`:

- Picks the next free numeric prefix in the chosen module
- Copies `spec/_template.md`
- Pre-fills title, version `1.0.0`, today's date
- Adds an entry stub to `99-consistency-report.md`
- Echoes the file path so the author can `code $(...)` straight into it

### 9. CODE-RED-024 in-editor lint for TS / TSX
Authors of *code* (not specs) keep tripping on bare `true` / `false` positional args. Add an ESLint rule (`no-restricted-syntax` with a tight selector) that flags `CallExpression > Literal[value=true|false]` outside test files. The fix is "import a named flag from `@/constants/boolFlags`."

### 10. Cross-link auto-fix
`check-spec-cross-links.py` reports broken links. A `--fix` flag that uses fuzzy matching against existing files would resolve 80 % of typos in one keystroke.

---

## Suggested rollout

| Wave | Items | Cost | Payoff |
|------|-------|------|--------|
| **Now (this PR)** | 1, 2 | shipped | Authors can write a valid spec without reading the long-form guide |
| **Next sprint** | 3, 4 | ~1 day | One-button validate + commit hook covers most CI failures locally |
| **Following sprint** | 5, 6 | ~2 days | Real-time preview + frontmatter linter |
| **Backlog** | 7, 8, 9, 10 | varies | Polish — every item independently shippable |

---

## What I'm explicitly not recommending

- **Web-based spec editor.** The Monaco-based editor in `MonacoMarkdownEditor.tsx` is sufficient for previewing; writing happens in the author's IDE. A custom web editor adds maintenance burden without a clear payoff.
- **AI-generated specs from prose prompts.** Spec quality depends on author intent. AI-assisted scaffolding (item 8) is fine; AI-authored specs are not.
- **Loosening the validator.** Every rule it enforces was added because it failed silently before. Keep the gate strict; lower the friction *around* the gate instead.

---

*See `docs/guidelines-audit.md` for the audit and the 3 quick wins shipped alongside this doc.*