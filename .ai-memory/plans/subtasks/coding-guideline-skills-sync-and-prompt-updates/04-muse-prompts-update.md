# Subtask 04 — muse prompts new versions (W2)

Owner: W2. Owns `01-prompts/27-muse-prompts/` additions only. Do NOT touch `03-ai-scripts/`.

## What to do

1. Read `01-prompts/27-muse-prompts/03-muse-master-with-repos.md` fully
   (v1.1.0) and the category `readme.md`.
2. Create a NEW version (next number `04-...`, with a `-v2` marker in the
   filename and the version header bumped, e.g. 1.2.0). Never edit `03-...`
   in place. It must contain:
   - (a) **Pull-first rule for ANY repository:** operating-contract bullet AND
     a step in the execution protocol — before any work, run `gitmap pull`;
     on failure fall back to regular `git pull` in the current
     repository/folder. Unmistakable.
   - (b) **Expanded repo registry:** add `story-telling-v1` (Bengali kids
     storytelling, ছোট গল্পের দেশ channel), `macro-ahk` (AHK macro
     automation), `scripts-fixture` (fixture scripts); annotate `cat-my-v12`
     as "Cat Malaysia". Keep the table's Route/How columns style.
   - (c) **No-wait checklist item:** explicit, prominent — DO NOT wait for
     confirmation/approval; break down the task, start working immediately;
     the reviewer reviews after completion.
3. Read `01-muse-master-prompt.md` (v6.1.0) and decide: if it needs the same
   pull-first + no-wait treatment, create `05-muse-master-prompt-v2.md`
   (header 6.2.0); if already covered, record the decision + evidence in the
   parent plan instead of a file.
4. Update `01-prompts/27-muse-prompts/readme.md` catalog (new rows, versions).

## Rules

- TOTAL BAN on rg/grep/git grep/Select-String → gitmap search tools.
- New versions only — never rewrite existing prompt files in place (this
  includes the user's V6 master prompt lineage).
- Relative paths only.

## Acceptance

- [ ] New repo-aware master version exists with (a) pull-first, (b) expanded
      registry, (c) no-wait checklist.
- [ ] Decision on `01-muse-master-prompt.md` recorded (new version or evidence
      it's covered).
- [ ] Category `readme.md` catalog updated.
