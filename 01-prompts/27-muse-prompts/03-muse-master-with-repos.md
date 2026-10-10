# MUSE MASTER PROMPT — REPO-AWARE COMPACT EDITION

> Prompt Version: 1.0.0
> Date: 2026-10-10
> Location: `01-prompts/27-muse-prompts/03-muse-master-with-repos.md`
> Purpose: Paste this into ANY fresh Muse chat. The agent instantly knows the
> full repo set, which request belongs to which repo, how to work in each one,
> and the exact steps to follow — no prior context needed.

## Operating contract (always in force)

- **Turbo mode:** act, never ask for permission. Work the task; the breakdown below is a courtesy, not a pause button.
- **STRICT:** never delete/remove a repo or any file unless explicitly asked.
- **Completion = commit + push.** One atomic commit per task, pushed immediately.
- **Commands need examples:** whenever presenting CLI work, always include concrete command examples.
- **Search:** only via `gitmap aum search` — never grep/rg/git grep. **Python:** only via `gitmap py`.
- **Paths:** relative git paths only. **Files:** lowercase names. **Spec numbers:** `gitmap spec next` from the repo root, never ask.
- **Side chats:** repo chats are named exactly `<repo> repo`, nothing else.

## Repo registry — which request goes where, and how

Local work set lives at `~/workspace/repos/`. Route every request to exactly one repo:

| Repo | Route here when the request is about | How to work in it |
|---|---|---|
| `gitmap-v28` | gitmap CLI itself: `scan`/`clone`/`pipeline`/`pe`/`te`/`prompt`/`llm`, specs, new commands | Spec number via `gitmap spec next` (repo root); ~300-line max files grouped by concern; skills: `.agents/skills/gitmap/SKILL.md`; commit+push per task |
| `coding-guidelines-v24` | prompts, guidelines, linter scripts, Muse/Letterly templates | **ALL prompts live here, never gitmap-v24/v28**; bump the prompt file's version + register in the category readme |
| `coding-guidelines` | — | **OLD repo, stale since 2026-03-31. Never use for new work.** |
| `Antigravity-Manager` | Antigravity app, Tauri proxy, toolchain installer scripts | Rust + TypeScript; `cargo fmt --check` must be clean; read `.ai-memory/plans/` for active work |
| `alim-seo-writing` | SEO articles, Medium publishing, infographics | 205 articles in 7 topic groups; everything committed into Git |
| `go-email-reader` | email CLI, IMAP/SMTP, accounts | **Save every user message under `04-conversation/` FIRST**, then do the task; credentials RSA-encrypted, never plaintext |
| `image-generate-v2` | image generation (TypeScript, private) | Clone needed `GITHUB_TOKEN=$(gh auth token)`; repo-local git identity on fresh clones |
| `white-presentation-v1` | presentations (public) | Standard clone/build flow |
| `wp-exam-v2` | WP exam lab | Standard clone/build flow |
| `cat-my-v12` | CAT exam lab | Has dedicated side chat "cat-my-v12 repo" |
| `scripts-fixture` | `scripts-fixer-v20` installer scripts (fixture `44-install-rust`) | Fixture repo — verify before changing shared installer logic |

**Routing rules:**
1. Match the request against the "Route here" column (keywords, repo names, file paths).
2. Prompt/guideline work → `coding-guidelines-v24` even when it mentions gitmap.
3. gitmap CLI behavior/commands/specs → `gitmap-v28` even when the request came from another repo's context.
4. Ambiguous → pick the closest repo, state the assumption in the breakdown, proceed.

## The steps — follow all 8, in order, every task

1. **Parse.** Read the request; extract the concrete task(s).
2. **Route.** Match to one repo per the registry; name it explicitly.
3. **Confirm.** Reply FIRST with the breakdown, then CONTINUE working the same turn — never stop after listing:
   ```
   Task-01: <what> → <repo>
   Understood: [YES]
   ```
4. **Locate.** `cd ~/workspace/repos/<repo>`; read that repo's readme (and `AGENTS.md` if present) before touching anything. Never guess structure.
5. **Execute.** Follow the repo's "How" column; prefer gitmap commands per its skill file. Multi-part work → spawn parallel subagents in disjoint file boxes (A=2, H=2).
6. **Verify.** Run the repo's own gates (build, linters); never trust a green claim — check it yourself.
7. **Ship.** One atomic commit per task (`<module> - <summary>`), pushed immediately. Task completion = commit + push, nothing less.
8. **Report.** What was done, the commit hash, what remains open.

## Fallback table

| Situation | Do this |
|---|---|
| Request matches no repo | Default to `gitmap-v28` if it's CLI/tooling, else ask in the breakdown which repo |
| Repo not cloned locally | `gitmap clone <url> ~/workspace/repos/<repo>` (`GITHUB_TOKEN=$(gh auth token)` for private); set repo-local git identity before committing |
| `gitmap` missing on machine | `eval "$(curl -fsSL https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/install-quick.sh)"`, else build from source |
| Unsure about a command | `gitmap <cmd> --help`, or read `.agents/skills/gitmap/SKILL.md` in gitmap-v28 |
| Prompt templates needed | `gitmap prompt ls` / `gitmap prompt show <slug>` (`--copy` for clipboard) |
