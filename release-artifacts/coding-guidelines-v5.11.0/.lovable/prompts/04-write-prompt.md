# Write Memory (v2.0.0)

**Trigger:** "write memory" / "end memory" / "update memory"
**Updated:** 2026-04-23

---

## Purpose

End-of-session persistence. Captures everything learned, done, and pending so the next AI session resumes with zero context loss.

## Phases

1. **Audit** — list completed tasks, pending tasks, learnings, failures.
2. **Memory** — update `.lovable/memory/<topic>/*.md` and `index.md`.
3. **Plans & Suggestions** — update `.lovable/plan.md` and `.lovable/suggestions.md` (single-file, with `## Completed` / `## Implemented` sections).
4. **Issues** — `.lovable/pending-issues/NN-name.md`, `.lovable/solved-issues/NN-name.md`, append hard prohibitions to `.lovable/strictly-avoid.md`.
5. **CI/CD issues** — `.lovable/cicd-issues/NN-name.md` + `.lovable/cicd-index.md` summary.
6. **Validate** — every memory file indexed, every solved issue has `## Solution`, no orphans.

## Status markers

| Status | Marker |
|--------|--------|
| Done | ✅ |
| In Progress | 🔄 |
| Pending | ⏳ |
| Blocked / Avoid | 🚫 |

## Naming

- All md files: `lowercase-hyphen.md` with numeric prefix `NN-`.
- Plans/suggestions live in **one** file each — never per-task folders.
- Issues: one file per issue.
- Memory: grouped by topic folder, every file listed in `index.md`.

## Anti-corruption

1. Never delete history — move to `## Completed` / `## Implemented`.
2. Never overwrite blindly — read first.
3. Never leave orphans — every file indexed.
4. Never split unified files (plan, suggestions).
5. Never mix states (pending vs solved).
6. Update `index.md` in the same operation as creating a memory file.

## Final confirmation format

```
✅ Memory update complete.
- Tasks completed: X
- Tasks pending: Y
- New memory files: Z
- Issues resolved: N / opened: M
- Suggestions added: S / implemented: T
Files modified: [list]
Next session can pick up from: [state + next step]
```

---

*Write-memory prompt — v2.0.0 — 2026-04-23*