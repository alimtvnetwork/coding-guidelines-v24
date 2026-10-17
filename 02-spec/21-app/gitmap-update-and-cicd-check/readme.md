# gitmap-update-and-cicd-check — parent spec

- slug: `gitmap-update-and-cicd-check`
- date: 2026-10-18
- home repo: `coding-guidelines-v24` (routing assumption: multi-repo ops task; spec/plan conventions and one release live here)
- priority: high, non-negotiable
- executor model: `execute-parent-task-with-n-steps-v6` (A=2 coordinators, H=2, disjoint file boxes)

## User Request (Verbatim)

# GitMap Update and CICD Check: high priority instruction, non-negotiable task

I want you to first do the gitmap update and run the gitmap scan on the root repositories folder. That means workspace/repos or repo, whatever the name is, where all the repos are there. Run that. And then you should run a gitmap space PA so that you have all the updates. Now, the next job is for you to check all these repositories, including the Gat MY `Antigravity Manager (`Antigravity Manager (AGM)`)`, which is the Antigravity Manager, gitmap, and Go email. All of those, I want the CICD to be green. Until the CICD green, you can check using the gitmap space PE space hyphen T. So it would show you as long as the CICD running that how long you have to wait. So once you have waited, and then you can check and fix the git issues and make a minor bump and release for Antigravity Gat MY coding guidelines, gitmap repo, and also Go email pro repo. Go email sender, I think, GER. And also please confirm that we have an icon for GER.

## task: GitMap Update and CICD Check
## slug: gitmap-update-and-cicd-check

# Actionable Items Must Follow Non-Negotiable

1. Git pull BEFORE starting any work — first try `gitmap pull <repo path>` (or run `gitmap pull` in repository folder); if that fails, fall back to regular `git pull` and resolve any merge conflicts so the repository is completely up to date and clean before starting.
2. Write spec under 02-spec/21-app/gitmap-update-and-cicd-check/ and enqueue plan task in .ai-memory/plans/gitmap-update-and-cicd-check.md (subtasks in .ai-memory/plans/subtasks/gitmap-update-and-cicd-check/) first
3. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String
4. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well
5. Perform gitmap update and run gitmap scan on the root repositories folder (workspace/repos or repo).
6. Run gitmap space PA to ensure all updates are captured.
7. Check all repositories including Gat MY `Antigravity Manager (`Antigravity Manager (AGM)`)`, gitmap, and Go email to ensure CICD is green.
8. Use gitmap space PE -T to monitor CICD status and wait until it is green.
9. Fix any git issues, make a minor bump, and release for Antigravity Gat MY coding guidelines, gitmap repo, and Go email pro repo.
10. Confirm that there is an icon for GER.

## Must follow and spawn agent using

[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)

## Additional Instructions

- [/plan](slashCommand;plan) first before doing the work to reduce the credits.
- [/learn](slashCommand;learn) from [gitmap](file;.agents/skills/gitmap) skill to leverage GitMap high-speed search, toolchain discovery, and caching.
- Only add the relative paths, never add the absolute path during your work; this should be respected on the release page and in release notes as well.

## Confirmed Task Breakdown

- Task-00: gitmap CLI reinstall/upgrade (binary missing on VM 2026-10-18), `gitmap scan ~/workspace/repos`, `gitmap pa`. Owner: Coordinator-01.
- Task-01: gitmap-v28 — pull, CICD green via `gitmap pe -t`, fix, minor bump + release. Owner: Coordinator-01.
- Task-02: Antigravity-Manager — pull, CICD green, fix, minor bump + release. Owner: Coordinator-02.
- Task-03: coding-guidelines-v24 — pull, CICD green, fix, minor bump + release. Owner: Coordinator-02.
- Task-04: go-email-reader (GER) — clone via git-over-SSH (not present locally; `gitmap clone` needs API token), pull, CICD green, fix, minor bump + release, confirm icon asset. Owner: Coordinator-02.
- Task-05: CICD-green sweep, fix-only (no bump/release): Coordinator-01 owns alim-seo-writing, story-telling-v1; Coordinator-02 owns cat-my-v12, repo-cache, repo-secrets.
- Task-06: lead consolidation — verify evidence, update `02-spec/21-app/readme.md` + `.ai-memory/plans/readme.md` indexes, final atomic commit, report. Owner: lead.

## Non-negotiable constraints

1. Pull before work (`gitmap pull`, fallback `git pull`); clean tree, conflicts resolved.
2. This spec + `.ai-memory/plans/gitmap-update-and-cicd-check.md` + subtask files first.
3. Codebase search ONLY via gitmap (`aum search`, `find`, `cat`, `ps`, `py`, `llm train`); TOTAL BAN on rg/ripgrep/grep/git grep/Select-String.
4. Relative git paths only — also on release pages and release notes.
5. R1: zero local builds/test suites (`go build`, `npm test`, `go test ./...`, etc. banned); CI verifies. `gitmap pe -t` waits are allowed.
6. Commits: `gitmap cpf/cpb/cpr "<module> - <summary>"`, hyphen-separated, NO colons; one atomic commit + push per task; repo-local identity `MD ALIM UL KARIM <devorg.bd@gmail.com>`.
7. R3: every DONE cites evidence (file path, diffstat, or exit code).
8. Minor bumps + releases are explicitly user-authorized (overrides zero-release rule) for: Antigravity-Manager, coding-guidelines-v24, gitmap-v28, go-email-reader only.
9. SSH: github.com via gitmap-managed key; note ssh resolves home as /root — keys/config mirrored at /root/.ssh per 2026-10-17 onboarding.
10. Out-of-scope findings go under `Follow-ups:` in the plan; never fix in this run.

## Evidence ledger

- Coordinator-01: `.ai-memory/plans/subtasks/gitmap-update-and-cicd-check/c1-gitmap-track.md`
- Coordinator-02: `.ai-memory/plans/subtasks/gitmap-update-and-cicd-check/c2-apps-track.md`
