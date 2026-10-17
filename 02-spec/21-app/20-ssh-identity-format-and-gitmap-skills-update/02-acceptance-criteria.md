# Acceptance Criteria: SSH Identity Format & GitMap Skills Fleet Synchronization

> Version: 1.0.0
> Date: 2026-10-17
> Status: Normative

---

## Acceptance Criteria Matrix

### Scenario 1: SSH Key Identity Formatting
- **Given** A Muse agent onboarding without active GitHub authentication
- **When** Option (b) (SSH Key Generation) is chosen
- **Then** The key comment token MUST match `<email>-muse<muse-sequence>-<operator-name>`
- **And** Git configuration is set using `git config --global user.email "<email>"` and `git config --global user.name "<operator-name> (Muse <muse-sequence>)"`

### Scenario 2: Interactive Ingestion of Missing Variables
- **Given** Any of `<email>`, `<muse-sequence>`, or `<operator-name>` are uninitialized
- **When** Preparing to generate the SSH key
- **Then** The agent MUST prompt the operator to supply the missing variable values before executing `ssh-keygen`

### Scenario 3: GitMap Skills Currency
- **Given** The `gitmap` skills in `.agents/skills/gitmap/skill.md` and `.cursor/skills/gitmap/skill.md`
- **When** Inspected for LLM and memory commands
- **Then** `gitmap llm train`, `gitmap llm train muse --days 7`, `gitmap agent`, and related subcommands are fully documented with operational invariants

### Scenario 4: Absence of Deprecated Muse Prompts Across Fleet
- **Given** The 43 connected repositories
- **When** Fleet synchronization completes
- **Then** Zero instances of `02-muse-execute-in-a-step.md`, `03-muse-master-with-repos.md`, `04-muse-master-with-repos-v2.md`, or `05-muse-master-prompt-v2.md` exist across all 43 repositories
- **And** `01-muse-master-prompt.md` is the sole prompt in `01-prompts/27-muse-prompts/`
