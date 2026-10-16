# Subtask 02 — prompts+skills sync mechanism (W1)

Owner: W1. Owns the new sync script + its doc. Do NOT touch `01-prompts/`.

## What to do

1. Read the existing sync scripts for convention:
   `03-ai-scripts/38-sync-prompts-skills-scripts.py`,
   `03-ai-scripts/42-sync-v3-prompts-and-skills.py`,
   `03-ai-scripts/43-upgrade-v4-prompts-and-skills.py`, plus
   `scripts/prompt-sync-config.json`.
2. Create `03-ai-scripts/50-sync-prompts-skills-to-targets.py` (next free
   number; verify with `ls 03-ai-scripts/` first) + a short doc beside it
   (follow how 38/42/43 document themselves, and the `03-ai-scripts/readme.md`
   index if it lists scripts).
3. Requirements:
   - Syncs PROMPTS (`01-prompts/`) + SKILLS (`.agents/skills/`,
     `.cursor/skills/`) from coding-guidelines-v24 into target repos.
   - NEVER syncs `02-spec/` — explicit fail-safe exclusion (default-deny, not
     default-allow); document this prominently.
   - Target registry: explicit list (config file and/or CLI args) — e.g.
     story-telling-v1.
   - Re-runnable and idempotent; `--dry-run` mode; clear per-target report.
   - Doc answers: what syncs, what never syncs, how to add a target, how to run.
4. Test-run against `~/workspace/custom-skills` (no git there — safe). Do NOT
   sync into `~/workspace/repos/story-telling-v1` (sibling coordinator working
   there; the parent runs that sync later).

## Rules

- TOTAL BAN on rg/grep/git grep/Select-String → gitmap search tools.
- Python via `gitmap py`. Relative paths only. Never delete.
- The user's V6 master prompt files are not to be rewritten — this subtask
  doesn't touch prompts at all.

## Acceptance

- [ ] Script + doc exist and follow repo conventions.
- [ ] `02-spec/` exclusion is explicit and tested (dry-run shows it skipped).
- [ ] Test run vs `~/workspace/custom-skills` succeeds with a per-target report.
