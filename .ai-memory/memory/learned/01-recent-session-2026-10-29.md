# Recent session — 2026-10-29

This note records what was created and shipped on 2026-10-29. Paths below are relative git paths.

## go-email-reader

- Specifications for warm-up sequences 08 through 13 are in `02-spec/21-app/08-ui-and-warm-up-algorithm-improvements/` through `02-spec/21-app/13-ui-and-warm-up-algorithm-improvements/`.
- Plans are in `.ai-memory/plans/pending/08-ui-and-warm-up-algorithm-improvements.md` through `.ai-memory/plans/pending/13-ui-and-warm-up-algorithm-improvements.md`.
- Sequence 08 is the connector and gate map. Sequence 09 is brakes. Sequence 10 is the send lease and cool-down. Sequence 11 is the shadow compare. Sequence 12 is the pace envelope. Sequence 13 is the shared quiet window.
- The live tick now consults `internal/warmup/gatefeed.go`. A peer can tighten the remaining-send cap and cannot raise it. An absent or stale peer does not stop local warm-up. Local quiet hours hold. A quiet skip does not send. A dry run does not spend the cap. The gate order is health, cadence, connector, theme.
- The reader UI gate map reads `GET /api/warmup/gates`.
- Commit `f3864ee` is on `main` and is pushed. The remote uses SSH.
- Sequences 08 through 13 stay pending as catalogs. Only the gate engine is implemented.
- `user.email` is `local@gitmap`. The name is `gitmap automation`.

## gitmap-v28

- Cursor-like IDE setup landed in `cli/cmdide/ide_setup.go`, `cli/cmdide/setup/`, and `cli/cmdui/ui_server_ide_setup.go`.
- The spec is `02-spec/21-app/352-cursor-like-ide-setup-from-gitmap/`.
- Pushed commits include `3229df9b4`, `864012968`, and `04690b952`.
- `cli/cov` is a coverage profile and was not committed.

## repo-cache

- Commit `f474298` stores heatmap databases, including `.heatmap/06-cursor/go-email-reader/heatmap.db`.
- SQLite `-shm` and `-wal` sidecars were not committed.

## Push denied

These local commits are still unpushed because GitHub denied `alimtvnetwork`:

- `alim.karim.profile` (19 commits)
- `kubernetes-training-v1` (19 commits)
- `common-linux-installer` (1 commit)

## Rules learned

- Do not put a personal access token in a git remote URL. Use SSH.
- Do not commit coverage profiles or SQLite sidecar files.

## Last 30 commits in this repo

```text
d397809b cfr(prompts): condense letterly and cursor prompts under 5000 characters and standardize relative skill links
98fe7283 Refactor: Prompt - enforce file; relative path format in V7 canonical prompt specifications
845d8966 Feature: Prompt - update skill mentions and letterly 14 to @[skill](file;.agents/skills/skill) syntax
f9f470ac Feature: Skills - update skill invocation mentions to double at prefix and fix spec21 filter in sync script
94ec4626 Feature: Prompt - update execute v7 invocation to double at prefix and letterly formatters
fa5e484a Feature: Prompt - consolidate dual mode time restriction in V7 skill and letterly 14
89959ace Feature: Prompt - add gitmap time restriction skills, V7 execution prompt, and letterly 14 prompt
1b5a73a6 Feature: Plans - close finished pending plans 03 05 30 34 and 35
eac42aa7 Feature: Authoring - add blind-agent spec ticket and skip missing consistency report
fcd4df35 Feature: Design System - append design spec checks 301 through 400
39d7a516 Feature: Design System - append design spec checks 201 through 300
11403b29 Feature: Design System - append design spec checks 101 through 200
d7b2f2b8 Feature: Design System - add blind-ai slide and tree spec with 100 checks
e6b2ed55 Feature: Skills - update gitmap skill with aggressive Python DB ban and db-find priority
99d93b64 Chore: Sync - dd-MMM-yy 12h date format and 10-day branch pruning
e26155a3 Chore: Sync - add consolidated skills to sync purge list
36979f11 Release: Guidelines - v6.88.0 GitMap clone bare slug fallback, show-log non-git resilience, and skills reduction matrix
719d22be Release: Skills - v6.87.0 PowerShell completion fix, follow-cg-execute 39-step orchestrator, and skills reduction
c45f471b Release: v6.86.0 - minor bump, follow coding guidelines skill, and gitmap skill sanitization
e410829e Feature: Skills - archive redundant skills, deduplicate commit reference, and add heatmap shorthand matrix
94f647d3 Feature: Skills - purge cluster and starbar bloat from gitmap skill and archive cluster fleet
1e4c4899 Feature: Skills - update gitmap skill flexible path cat, proc kill, and casing parity
bccb00e8 Feature: SKILL - document gitmap swap-binary command and ban manual binary copy
f807f122 Feature: Skills - add follow-coding-guidelines and follow-cg-execute, sanitize gitmap skill casing, archive obsolete prompts
df3cb412 Release: v6.85.0 - subsystem module conventions and commit pseudo-module ban
244d6e60 Chore: guidelines - standardize lowercase module in commit rules and skills
39002171 CHORE(GUIDELINE): coding guideline and folder structure improvement
b24a82e6 Feature: PROMPTS - add 04 muse universal master prompt following 01 architecture with dynamic git onboarding
8bba13aa Feature: PROMPTS - add generic standalone master prompt for muse without gitmap
7343a8bd Feature: PROMPTS - condense muse long update and sleepy letterly formatters under 5800 chars
```
