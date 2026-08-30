# Folder Structure

This document serves as the absolute source of truth for the `.lovable/` folder ecosystem and repository architecture. All AI agents MUST strictly adhere to this architecture when reading, saving, or generating files.

## The XX-<slug> Sequence ID System (Strict Prefixing)

To maintain deterministic ordering, EVERY file and folder inside `.lovable/plans/`, `.lovable/prompts/`, `.lovable/ai-fix-scripts/`, and subdirectories MUST be prefixed with a 2-digit zero-padded number (`XX-`).

*   **Format:** `XX-<slug-name>.md`, `XX-<slug-name>/`, or `XX-<script-name>.py`
*   **Examples:** `01-setup-database.md`, `02-create-auth-service.md`, `15-deploy-frontend.md`
*   **Rules:**
    *   `XX` always starts at `01` (or `00` for root indexes / foundational files) and increments sequentially.
    *   If you create a new file, calculate the highest existing `XX` in that folder and add 1 (e.g., if `04-foo.md` exists, create `05-bar.md`).
    *   NEVER use un-prefixed names like `setup-database.md`.
    *   NEVER use letters as prefixes (e.g., `A-setup.md`).

## Core Directories

### `.lovable/plans/`

The command center for execution.

*   `01-index.md`: The master registry. MUST be updated simultaneously whenever a file is moved or created in the plans directory.
*   `/pending/`: Holds high-level parent task specs (e.g., `01-auth-spec.md`).
*   `/subtasks/XX-<slug>/`: If a parent task is complex, it is broken down into a dedicated subfolder here, containing sequential micro-tasks (e.g., `01-create-table.md`, `02-write-query.md`).
*   `/completed/`: Where pending tasks are moved once successfully executed.

### `.lovable/memory/`

The long-term cognitive storage for the AI.

*   `01-index.md`: The table of contents for memory.
*   Stores architectural decisions, project-specific mappings, and state tracking.

### `.lovable/prompts/`

The prompt registry for AI workflows.

*   `/01-prompts-category/`: The primary source of truth for all structured prompt categories (organized by topic/workflow).
*   `*.md` (flat prompts): Generated/compiled mirror from `/01-prompts-category/` via `scripts/update-prompts.ps1` and `scripts/prompt-sync-config.json`.
*   All prompt modifications MUST occur in `/01-prompts-category/` and then be compiled to the flat structure.

### `.lovable/release/`

The continuous delivery and versioning hub.

*   `release-method.md`: Documents exactly *which* files contain version strings and *how* to bump them for this specific project.
*   `bump_versions.py`: The automated script that reads `release-method.md` and safely mutates versions without unbounded global searches.
*   `/issues/`: Holds bug reports (e.g., `01-v1.2.3-build-failed.md`) from failed releases.

### `.lovable/ai-fix-scripts/`

A persistent toolkit of reusable AI helper scripts.

*   `01-index.md`: Master index explaining each script, its purpose, execution syntax, and search tags.
*   All scripts must follow numeric prefixing (`01-`, `02-`, `03-`, etc.).
*   Must be committed to the repository so future AI sessions can reuse them.

### `.lovable/assets/`

The media vault.

*   Store all user-uploaded UI mockups, diagrams, and reference images here.
*   Never dump images in the root directory.

### `.lovable/question-and-ambiguity/`

The resolution tracker.

*   `readme.md`: Index of open and resolved project questions and ambiguities.
*   `task-counter.md`: Running tally of tasks and iterations.

### `.lovable/suggestions/`

Optimization and feature recommendations.

*   `01-index.md`: Index of suggestions and enhancements.
