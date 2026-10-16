# SPEC — coding-guideline-skills-sync-and-prompt-updates

- slug: `coding-guideline-skills-sync-and-prompt-updates`
- repo: `coding-guidelines-v24` (github.com/alimtvnetwork/coding-guidelines-v24)
- date: 2026-10-17
- status: in progress
- source: user prompt "Learning and Storytelling Audio Creation" (coding-guidelines track)
- orchestrator: coordinator subagent, v6 playbook (A=2 per phase, disjoint file boxes)

## 1. Goal

Bring the newest gitmap skills into coding-guidelines-v24 (the canonical home of
all prompts), add a re-runnable prompts+skills sync mechanism for repos that do
not need spec, and ship new prompt versions that pull before working and never
wait for approval — then release the repo with green CI.

## 2. Workstreams

### WS1 — gitmap skills re-read + inventory (W1)

Inventory new/updated gitmap skills across `01-prompts/26-gitmap/`,
`.agents/skills/`, `.cursor/skills/`, and skill docs. Diff against the ingested
set (`gitmap`, `gitmap-agent-orchestrator` + 6 subsystems: scanner-and-cloner,
ssh-cluster-fleet, split-db-engine, pipeline-and-diagnostics,
developer-hygiene-and-agy, macro-automation-engine). Record the inventory
(what changed, what's new, what to adopt) in the plan.

### WS2 — sync mechanism (W1)

A re-runnable sync living INSIDE coding-guidelines-v24 (script under
`03-ai-scripts/` + short doc, following the convention of the existing
`38/42/43` sync scripts): syncs PROMPTS + SKILLS from this repo into target
repos that don't need spec (e.g. story-telling-v1). Must document: what syncs,
what NEVER syncs (`02-spec/` — explicit fail-safe exclusion), how to add a
target. Test-run against `~/workspace/custom-skills` (no git there, safe).

### WS3 — letterly prompt, new version (W2)

New file `01-prompts/22-letterly/11-<slug>-letterly.md` (next sequence number,
keep the `-letterly.md` suffix; update `readme.md` index). Its FIRST
instruction is unmistakable: Git pull BEFORE starting any work — first try
`gitmap pull`; if that fails, fall back to regular `git pull` in the current
repository/folder. Numbered step 1, plain language.

### WS4 — muse prompts, new versions (W2)

New versions only (never edit in place) in `01-prompts/27-muse-prompts/`:
(a) repo-aware master (03 lineage): git-pull-first rule for ANY repository it
works with (`gitmap pull` → fallback `git pull`), in the operating contract AND
the step protocol; (b) registry additions — `story-telling-v1` (Bengali kids
storytelling, ছোট গল্পের দেশ), `macro-ahk`, `scripts-fixture`, and annotate
`cat-my-v12` as "Cat Malaysia"; (c) explicit checklist item: DO NOT wait for
confirmation/approval — break down the task, start working immediately, the
reviewer reviews after completion. Also read `01-muse-master-prompt.md` and
decide whether it needs the same treatment (new version if yes, recorded
decision if no). Update the category `readme.md`.

## 3. Release

Commit via `gitmap cpf` (hyphen format, no colons in args). Then release per
`.ai-memory/release/release-method.md`:
`gitmap py .ai-memory/release/bump_versions.py --type minor --create-release`
(minor: new prompts + new sync mechanism). Verify CI GREEN via `gh run list` /
checks; diagnose red, fix only own breakage, report the rest.

## 4. Constraints

- TOTAL BAN on rg/grep/git grep/Select-String → `gitmap aum search`,
  `gitmap find`, `gitmap cat` (export `TERM=xterm-256color`).
- Relative git paths only, everywhere (commits, notes, release).
- Never delete anything. Never rewrite the user's V6 master prompt in place —
  new versions only.
- Do NOT sync into story-telling-v1 in this run (sibling coordinator is
  committing there in parallel; the parent runs that sync).
- gitmap here is v6.557.0 (freshly rebuilt); use it, don't reinstall.
- No silent errors: every DONE cites file paths / SHAs / command output.

## 5. Out of scope

- story-telling-v1 sync execution (parent's job after the sibling finishes).
- Changes to the gitmap CLI itself (that lives in gitmap-v28).

## 6. Acceptance criteria

- [ ] Skills inventory recorded in the plan with new/changed items named.
- [ ] Sync script + doc exist, re-runnable and idempotent; `02-spec/` explicitly
      excluded; test run against `~/workspace/custom-skills` succeeds.
- [ ] New letterly prompt version exists with pull-first as step 1; readme
      index updated.
- [ ] New muse prompt versions exist with pull-first rule, expanded registry,
      and no-wait checklist; readme catalog updated.
- [ ] Committed via `gitmap cpf`; released (minor bump); CI green verified via
      `gh`; push verified via `git ls-remote`.
