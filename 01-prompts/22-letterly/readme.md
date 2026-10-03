# Letterly Prompt Formatters (`22-letterly`)

This collection contains deterministic prompt transformation templates designed to ingest raw voice-dictated user transcripts from Letterly (mobile and desktop) and format them into rigorous, non-negotiable agent prompts.

## Directory Index

| # | Prompt File | Target Mode | Description | Companion Skill |
| :---: | :--- | :--- | :--- | :--- |
| **01** | [`01-mobile-letterly.md`](01-mobile-letterly.md) | Mobile Single-Line | Compact one-liner without line breaks starting with `# High Priority Instruction:` | `letterly-mobile` |
| **02** | [`02-desktop-letterly.md`](02-desktop-letterly.md) | Desktop Structured | Formats voice dictation into high-priority instructions, action items, and V6 suffix | `letterly-desktop` |
| **03** | [`03-execute-n-steps-letterly.md`](03-execute-n-steps-letterly.md) | Execute N-Steps | Verbatim input capture with discrete action items and V6 execution engine | `letterly-execute-n-steps` |
| **04** | [`04-plan-letterly.md`](04-plan-letterly.md) | Planning Spec | Formats raw input into architectural planning directives with no-build rules | `letterly-plan` |
| **05** | [`05-release-letterly.md`](05-release-letterly.md) | Minor Release (Letter U) | Formats input into minor release update ceremony directives | `letterly-release` |
| **06** | [`06-cicd-fix-release-letterly.md`](06-cicd-fix-release-letterly.md) | CI/CD Fix & Release | Formats input into `gitmap pe -t` telemetry diagnosis, 4-part RCA, and minor release | `letterly-cicd-fix-release` |
| **07** | [`07-mobile-cicd-fix-letterly.md`](07-mobile-cicd-fix-letterly.md) | Mobile CI/CD Fix | One-liner mobile prompt diagnosing via `gitmap pe -t` and triggering release | `letterly-mobile-cicd-fix` |

### Cursor Prompt Suite (`cursor/`)

For Cursor environments requiring `.cursor/skills/` resolution, see [`cursor/readme.md`](cursor/readme.md) which houses the dedicated Cursor prompt suite (`01-mobile-letterly-cursor.md` through `07-mobile-cicd-fix-letterly-cursor.md`).

## Core Invariants

1. **Zero Conversational Filler:** Never output "Certainly! Here is your output:".
2. **Lossless Verbatim Capture:** Never drop specific technical flags, file paths, or commands from input.
3. **Mandatory Suffix Invocations:** Every formatted prompt must conclude with its designated skill link.
4. **Standard Letterly Suffix:** All prompt files in this directory strictly carry the `-letterly.md` file suffix.
