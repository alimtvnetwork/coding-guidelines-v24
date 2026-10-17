# Subtask 02: Muse Master Prompt & Onboarding Spec Updates

## Objective
Update Muse master prompt and onboarding specifications with SSH identity format `<email>-muse<muse-sequence>-<operator-name>`, account email auto-discovery, native `gitmap ssh create` command execution, key file delivery protocol (banning chat text pasting), and explicit key type designation.

## Targets
- `01-prompts/27-muse-prompts/01-muse-master-prompt.md`
- `.agents/skills/muse-master-prompt/skill.md`
- `.cursor/skills/muse-master-prompt/skill.md`
- `02-spec/21-app/16-unified-muse-master-prompt-v6/03-onboarding-protocol-spec.md`

## Acceptance Criteria
- Exact identity format `<email>-muse<muse-sequence>-<operator-name>`.
- Auto-extract email via `git config user.email` or `gh api user --jq .email`. Prompt operator for missing variables.
- Command `gitmap ssh create "<identity>" -y`.
- Absolute ban on raw chat key pasting: deliver as file (`~/.ssh/<identity>.pub`, `.ai-memory/ssh/<identity>.pub`) or copy via `gitmap ssh copy`.
- Explicitly state key type (Authentication Key vs Signing Key).
