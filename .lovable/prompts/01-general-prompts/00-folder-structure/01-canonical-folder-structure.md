# Canonical Folder Structure & Sizing Rules — Architecture Spec (must follow)


> **Prompt Version:** 2.1.0
> **Synchronization:** Main Meta-Repo & Connected Workspaces

Every prompt in this set MUST write into this exact tree. No invented folders, no alternate paths. If a prompt seems to want a folder not listed here, the prompt is wrong, not the tree.

All dates are UTC. All filenames are lowercase kebab-case with a two-digit zero-padded XX prefix where sequencing applies. XX is monotonic within its folder scope.

`	ext
/                                      # root of the repository
  spec/                                # canonical app specs (source of truth)
    02-coding-guidelines/              # detailed coding architectures
    21-<area>/                         # foundational app spec areas (NEW SPECS GO HERE)
    25-app-spec-audit/                 # output of app-spec audit
      00-overview.md
      NN-audit-<date>-v<version>.md
    17-consolidated-guidelines/

  .lovable/
    what-to-read.md                    # priority reading list for read-memory
    strictly-avoid.md                  # universal anti-patterns
    
    coding-guidelines/
      coding-guidelines.md             # mirror of spec source of truth

    memory/                            # cross-session AI context (FLAT FILES ONLY)
      00-index.md                      # master index of all memory files
      XX-<slug>.md                     # flat memory files (specs, workflow, standards, etc.)

    prompts/                           # ONLY when a prompt body itself changes
      XX-<slug>.md

    plans/
      index.md                         # roll-up of every plan, status, links
      pending/
        XX-<slug>.md                   # active plans (MAPPED & INDEXED IN index.md)
      subtasks/
        XX-<slug>/
          01-<subslug>.md              # granular step-by-step tasks
      completed/
        XX-<slug>.md                   # finished plans (moved, not copied)

    temp-agents/                       # agent lifecycle state management
      XX-agent-state.md                

    issues/                            # app-level bugs and blockers
      XX-<slug>.md

    cicd-issues/                       # CI / CD failures
      XX-<slug>.md

    audits/
      XX-<work-slug>/                  # one folder per recent-work audit run
        01-index.md
        02-<work-slug>.md

    release/
      issues/
        XX-<version>-<slug>.md         # release-time failures and flags

    ambiguous-questions/               # see 14-ambiguity-prompt.md
      01-new-ambiguity/
        XX-<slug>.md
      02-ambiguity-resolved/
        XX-<slug>.md

    assets/                            # attachments referenced by specs / plans / audits
      <category>/
        XX-<slug>.<ext>
`

## Numbering rules (XX and NN)

- Two-digit zero-padded, monotonic per folder scope.
- Next XX = max existing XX in that folder scope + 1. Never reuse.
- Ambiguous questions: XX is monotonic across BOTH 01-new-ambiguity/ and 2-ambiguity-resolved/ combined.
- Audits: XX in .lovable/audits/XX-<work-slug>/ is monotonic across all audit folders combined.
- Release issues: XX monotonic within .lovable/release/issues/ regardless of version.
- Plans: XX monotonic across pending/, subtasks/, and completed/ combined; filename does not change when moving between them.
- Spec Audits: NN in spec/25-app-spec-audit/ increments based on existing files.

## Movement rules

- mv, never cp. Files carry their filename when their lifecycle changes folders (plan pending -> completed, ambiguity open -> resolved).
- Never leave a copy behind in the origin folder.
- Never rename on move.

## What each folder is for

- spec/21-<area>/ - This is where NEW specs go. App specs are sourced here at the ROOT, not inside .lovable/.
- spec/25-app-spec-audit/ - The only place pp-spec-audit writes its output.
- plans/ - Every plan built by planning prompts. pending/ holds the master plan (indexed in plans/index.md), subtasks/ holds granular execution steps.
- memory/ - Retained context so agents do not lose track of decisions. MUST be kept completely FLAT. All files are XX-<slug>.md. The main index is 0-index.md.
- 	emp-agents/ - Explicit state files where agents record their current step and crash logs.
- issues/ and cicd-issues/ - Runtime bugs and pipeline failures.
- udits/ - One folder per recent-work audit run (distinct from app-spec audit).
- 

elease/issues/ - Only issues discovered during release.

- mbiguous-questions/ - Open questions raised by any prompt. 
- ssets/ - Any binary or image. Reference by relative path with a caption.
- prompts/ - Only touched when a prompt body itself changes.

## Hard bans

- No mbiguities/, questions/, lockers/, 	odo/, 
otes/, or any folder name not listed above.
- No subfolders inside .lovable/memory/. Flat files ONLY.
- No per-turn mirror files under .lovable/prompts/.
- No renaming a file when moving it between lifecycle folders.
- No guessing when the folder is unclear. File an ambiguity.
 
