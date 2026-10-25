# Confusion branch notes — muse prompt changes, 25-Oct-25

Branch: `confusion/muse-25-Oct-25` (named by DATE only, no time — time is irrelevant here).
Purpose: collect all current muse-prompt changes in one place so another AI can fix and merge this branch.
Base: latest `main` at branch time, so every change below is already on this branch.

## Change 1 — commit discipline (commit `6bbdd74f`)

What: master prompt Phase 7 + generic prompt commit section.
Why added: the user corrected the plan-325 working style — dozens of micro-commits made review and debugging hard, and per-commit CI watches multiplied wait time.
Reasoning: comment all changes (brief WHY) so intent survives; one compact commit per reviewable unit so review and bisection look at one commit; verify CI/CD once at the end instead of after every commit.

## Change 2 — fast-draft mindset MUST checklist + scope rule (commit `6aba40cd`)

What: master prompt Phase 7, generic prompt commit section, muse-master-prompt skills (`.agents` + `.cursor`).
Why added: the user's mindset — git is a quick draft, not a ceremony.
Reasoning: with many agents working (the user cited 10), verifying everything on the fly in memory wastes time and everyone's resources. Commit all modified files together and push immediately; verify at the end of the tunnel. Scope rule added in the same change: prompt/skill updates apply to the coding-guidelines prompts and skills ONLY — never sync across other repositories, because the user has other tools for that.

## Change 3 — doubt marking + execute-v6 N-task rule (commit `f4ca2953`)

What: master + generic checklists, muse-master-prompt skills, execute-v6 prompt + its skills.
Why added: the user wants doubtful code to be findable and verifiable by a future AI, and wants the v6 CI rule stated as end-of-Nth-task batching.
Reasoning: a strict DOUBT comment at the exact spot plus a `gitmap heatmap log` entry (file + line number + what needs verification) makes re-verification tool-driven instead of memory-driven. Execute v6 already banned CI/CD unless asked (R1); the added sentence only makes the batching explicit — N tasks verify build, test cases, and CI/CD once at the end of the Nth task.

## Change 4 — date & time standard (commit `a0f3ef3c`)

What: master prompt, generic prompt, execute-v6 prompt, and all three skill mirrors.
Why added: the user set a single date/time standard for all prompts.
Reasoning: DD-MMM-YY dates are unambiguous across locales; 24-hour time removes AM/PM ambiguity; saving everything as UTC keeps stored data canonical with no timezone drift; GMT+8 is shown only when the user wants to see a date/time, so display converts on request instead of corrupting storage; branch names use the date only and time is omitted wherever it is irrelevant, keeping names and records clean.
Note on a wording conflict in the request: the user first mentioned 12-hour format, then corrected to 24-hour ("time should be always in 24 hours"). The prompts record 24-hour as the final rule — flag this if 12-hour was intended anywhere.

## For the AI fixing and merging this branch

- All prompt changes live in coding-guidelines ONLY, nowhere else — do not sync them to other repositories.
- Do not rewrite shipped history; this branch is for fix-and-merge, not history repair.
- If any rule above conflicts with a newer user instruction, the newer instruction wins — record the decision in this file.
