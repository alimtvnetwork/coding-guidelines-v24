# Muse Prompts (`27-muse-prompts`) — Index & Catalog

> [!IMPORTANT]
> Category: `27-muse-prompts`
> Architecture: Self-Bootstrapping Autonomous Onboarding Directive
> Scope: Muse AI Session Bootstrapping, Operating Contract, and Multi-Agent Execution

---

## Overview

The `27-muse-prompts` prompt category houses the master onboarding prompt for
Muse AI sessions. Pasting the master prompt into a fresh Muse chat boots a
fully-onboarded autonomous agent: repository intake, terminal + GitHub auth
bootstrap, GitMap clone/build/learn, a memory operating-contract checklist
(turbo mode, strict no-delete guards), coding-guideline and design-system
intake — and then the standing **Task Confirmation & Multi-Agent Execution
Protocol**: every user message that is work gets a confirmed task breakdown
first (`Understood: [YES]` per task), then multiple concurrent agents complete it.

The canonical execution engine it delegates to is
[`02-execute-parent-task-with-n-steps.md`](../14-execute/02-execute-parent-task-with-n-steps.md)
(Canonical V6: `N = 300`, `A = 2`, `H = 2`, `C = 30`).

Its installable skill is `muse-master-prompt`
(`.agents/skills/muse-master-prompt/skill.md`, mirrored in `.cursor/skills/`).

## Muse Memory IO (Memory Re-sync Protocol)

The master prompts [`01-muse-master-prompt.md`](01-muse-master-prompt.md)
(6.1.1) and [`05-muse-master-prompt-v2.md`](05-muse-master-prompt-v2.md)
(6.2.1) carry a "Muse Memory IO" section (§11). The private vault repo
`alimtvnetwork/repo-secrets` holds a 7-day rolling condensed memory digest
under its `Muse/` folder (sequenced `README.md`, `01-overview.md` →
`05-conventions.md`, `chat-seeds/*.md`, Muse/last-read.md checkpoint).

When the user says **`Muse memory IO read`** (in ANY chat), the session runs
the say-READ protocol: clone (or `git pull --ff-only`) repo-secrets into the
work-folder root, diff vs the Muse/last-read.md checkpoint, read the
new/updated sequenced files, summarize changes, recreate the last-5-days chats
from `Muse/chat-seeds/*.md`, and update the checkpoint.

Two standing rules travel with the protocol:

- **PRIVATE ITEMS RULE:** private repo lists, account names, and sensitive
  references resolve FROM the repo cache (`Muse/` in repo-secrets) — never
  hardcoded in prompts. Pointers only, never secret values.
- **EMAIL RULE:** when a memory-IO read shows go-email-reader setup is needed,
  ask the user ONLY for the email password — account, address, and server
  config come from the encrypted store; never re-ask for non-secret info.

---

## Directory Index & Catalog

| Prompt File | Version | Scope | Key Capabilities |
| :--- | :--- | :--- | :--- |
| [`01-muse-master-prompt.md`](01-muse-master-prompt.md) | 6.1.0 | Master Onboarding | Self-loop Phases 0–5, GitMap reinstall guard, auth bootstrap, GitMap intake via skills, memory contract checklist, CI-only test verification (`gitmap pe -t`), task-confirm + multi-agent protocol (A=2/H=2), understanding check-ins, Top-Instruction Priority Mandate. |
| [`02-muse-execute-in-a-step.md`](02-muse-execute-in-a-step.md) | 1.1.0 | Execute-in-a-Step | Breakdown-first listing, RUNNING + ETA declaration ("Are you running or not?"), 5-minute status pings, commit+push completion, CI-only verification, GitMap skill authority, understanding check-ins. |
| [`03-muse-master-with-repos.md`](03-muse-master-with-repos.md) | 1.1.0 | Repo-Aware Master | Compact paste-and-go prompt: embedded 10-repo registry (which request → which repo + how), routing rules, 8-step execution protocol (parse → route → confirm → locate → execute → verify → ship → report), GitMap reinstall + CI-only verification, fallback table. |
| [`04-muse-master-with-repos-v2.md`](04-muse-master-with-repos-v2.md) | 1.3.0 | Repo-Aware Master v2 | 13-repo registry, pull-first + no-wait checklist, large-file 15 MB split rule, repo-cache NN-sequence convention, explicit commit instructions (gitmap cpf + ls-remote verify). |
| [`05-muse-master-prompt-v2.md`](05-muse-master-prompt-v2.md) | 6.2.2 | Master Prompt v2 | V6 agent directive (N=300/A=2/H=2), self-loop onboarding Phases 0–5, turbo ask/don't-ask matrix, large-file 15 MB split rule + repo-cache sequence in the commit gate. |

---

## Directory Tree

```text
01-prompts/27-muse-prompts/
├── 01-muse-master-prompt.md   # the master prompt (paste into a fresh Muse session)
├── 02-muse-execute-in-a-step.md   # one-shot execution prompt (breakdown → RUNNING+ETA → 5-min pings → commit+push)
├── 03-muse-master-with-repos.md   # repo-aware compact master (embedded repo registry + routing)
└── readme.md                   # this index
```
