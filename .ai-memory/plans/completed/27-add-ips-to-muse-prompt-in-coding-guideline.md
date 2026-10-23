# Plan 27: Add IPs to Muse Prompt in Coding Guideline

Status: COMPLETED
Owner: Muse (this session)
Spec: `02-spec/21-app/27-add-ips-to-muse-prompt-in-coding-guideline/01-architecture-spec.md`
Subtasks: `.ai-memory/plans/subtasks/27-add-ips-to-muse-prompt-in-coding-guideline/`

## Subtasks
1. `01-master-prompt-ip-section.md` — add the requested-endpoints section to `01-prompts/27-muse-prompts/01-muse-master-prompt.md`, bump to 6.5.1, sync the `27-muse-prompts` readme catalog.
2. `02-skill-mirror.md` — mirror the section into `.agents/skills/muse-master-prompt/skill.md` and `.cursor/skills/muse-master-prompt/skill.md`.

## Invariants
- Pull first (done, tree clean at claim time); heatmap checked (0 in-flight actions).
- GitMap-only search; relative paths only; atomic commit + push via GitMap; verify with the GitHub API afterwards (not `gitmap pe`, whose local DB can be stale).
