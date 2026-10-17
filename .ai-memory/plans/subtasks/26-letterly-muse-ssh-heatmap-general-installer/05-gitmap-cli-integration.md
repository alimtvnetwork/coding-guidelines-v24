# Subtask 05: GitMap CLI Integration for General & Coding Guidelines Installation

## Objective
Extend `gitmap cg` (and `gitmap cd` alias handling) in `gitmap` to support `gitmap cg install general` / `gitmap cg general` / `gitmap cd general`, synchronizing directly from local `coding-guidelines` or invoking the PowerShell installer to bootstrap general repos automatically.

## Targets
- `gitmap: cli/cmdcg/cg.go`
- `gitmap: cli/cmdcg/cg_worker.go`
- `gitmap: cli/cmdcd/`

## Acceptance Criteria
- `gitmap cg install general`, `gitmap cg general`, `gitmap cg install --general`, and `gitmap cd general` execute general repo installation.
- Automatically synchronizes general assets directly from local repository if present.
