# Python File Manipulator CLI Specification — Tooling Spec (must follow)

> **Prompt Version:** 2.2.0
> **Synchronization:** Main Meta-Repo & Connected Workspaces

/goal Autonomously generate a robust, dependency-free Python CLI tool to handle mass file renaming, sequencing, and encoding normalization across any specified folder.

## Overview

You are an expert Python Developer AI. Your task is to write a standalone, reusable Python script that handles mass file renaming (lowercasing), sequence fixing, and encoding normalization across any target directory. This script will act as an autonomous tool for other AIs and developers to organize files without needing a compiled binary.

**Target Path:** `.lovable/ai-fix-scripts/01-file-manipulator.py`

## Non-Negotiable Rules for the Python Script

1. **Zero Dependencies**: The script MUST use only Python standard libraries (`os`, `sys`, `argparse`, `shutil`, `subprocess`, `pathlib`, `enum`, `json`, `time`, `re`).
2. **Root Constants & Enums**: Top-level configuration constants and Enums ending with `Type` suffix (`ScanModeType`, `ExitCodeType`) MUST be declared at the root of the file.
3. **Small Decomposed Functions**: Code must be decomposed into small, testable functions under 25 lines each following the Single Responsibility Principle.
4. **DRY Shared Engine**: Import common filesystem scanning, caching, and line-ending utilities from `00-shared-engine.py`.
5. **Two-Phase Incremental Caching**: Leverage `tmp/cache/repo-file-cache.json` to start with cached entries first and stream newly discovered or modified files concurrently.
6. **Pre-compiled Regex Engine**: Compile all regex patterns (`re.compile`) at module initialization or function entry to avoid repetitive in-loop compilation and ensure sub-15ms execution.
7. **Multi-Folder Scoping**: Support target directory specification (`<path>` or `--path` / `--dir`) so the tool can run on any directory, repository root, or nested subfolder.
8. **Customizable File Extensions**: Support `--ext` for filtering specific file extensions (e.g. `.md,.ts,.py`), normalizing extensions to lowercase with leading dots.
9. **Nested Ignore Pruning**: Prune `.git`, `.gitmap`, `node_modules`, `dist`, `build`, `.venv`, `.gemini`, `tmp`, `.system_generated`, and `release-artifacts` at all subtree depths including nested subprojects.
10. **Windows Long Paths**: Normalize paths and safely handle Windows `MAX_PATH` limitations.
11. **Git Awareness**: Attempt `git mv` via subprocess first, gracefully falling back to standard `shutil.move` / `os.rename`.
12. **Update Index**: Document usage in `.lovable/ai-fix-scripts/01-index.md`.

---

## Core Feature 1: Lowercase Renamer

**Command Pattern**:
```bash
python .lovable/ai-fix-scripts/01-file-manipulator.py lowercase <target_directory> [flags]
```

**Requirements**:
1. Recursively convert all files matching a target pattern to lowercase.
2. **Extension Enforcement**: It MUST strictly ensure that all markdown and code file extensions are lowercased (e.g., converting `.MD` to `.md`).
3. **Default Ignores**: By default, the script MUST silently ignore `node_modules`, `.git`, `.gitmap`, and other excluded folders. Do not traverse them.
4. **Extendable Ignores**: Provide an `--except` flag accepting a comma-separated list of additional files, folders, or wildcard patterns to ignore (e.g., `--except "vendor/*, build/*"`).

**Example Output in `--help`**:
- `python .lovable/ai-fix-scripts/01-file-manipulator.py lowercase ./src`
- `python .lovable/ai-fix-scripts/01-file-manipulator.py lowercase ./spec --except "images/*"`

---

## Core Feature 2: Fix File Sequencing (`fix-seq-files`)

**Command Pattern**:
```bash
python .lovable/ai-fix-scripts/01-file-manipulator.py fix-seq-files <target_directory> [flags]
```

**Requirements**:
1. Scan the specified directory for sequenced files (e.g., `01-draft.md`, `02-notes.md`).
2. **Ordering Flags**:
   - `--order-by-time`: Re-sequence files sequentially based on their filesystem modification time.
   - `--order-by-az`: Re-sequence files alphabetically based on the string following the sequence number.
3. **Tie-Breaker / Preservation**:
   - `--keep-old-order`: Preserve existing numeric ordering as much as possible. Only assign new sequence numbers to unnumbered files or resolve direct conflicts using time/alphabetization.
4. **Fixated / Pinned Sequences**:
   - `--pin "<mapping>"`: Allow users to explicitly lock specific files to a sequence number (e.g., `--pin "readme=00,draft=01"`). The script must increment other files around these locked sequences.

**Example Output in `--help`**:
- `python .lovable/ai-fix-scripts/01-file-manipulator.py fix-seq-files ./docs --order-by-time`
- `python .lovable/ai-fix-scripts/01-file-manipulator.py fix-seq-files ./spec/02-coding-guidelines --order-by-az --keep-old-order`
- `python .lovable/ai-fix-scripts/01-file-manipulator.py fix-seq-files ./spec --pin "readme=00,intro=01"`

---

## Core Feature 3: Fix Encoding & Line Endings (`fix-encoding`)

**Command Pattern**:
```bash
python .lovable/ai-fix-scripts/01-file-manipulator.py fix-encoding <target_directory> [flags]
```

**Requirements**:
1. Scan the specified directory and aggressively fix encoding issues for all text files (customizable via `--ext`, defaulting to standard text extensions).
2. **BOM Stripping**: Detect and strip any UTF-8 Byte Order Marks (BOM) or UTF-16 encodings, standardizing everything strictly to UTF-8 without BOM.
3. **Line Ending Normalization**: Automatically convert all Windows CRLF (`\r\n`) line endings to Unix LF (`\n`) to prevent git warnings and cross-platform issues.

**Example Output in `--help`**:
- `python .lovable/ai-fix-scripts/01-file-manipulator.py fix-encoding ./src`
- `python .lovable/ai-fix-scripts/01-file-manipulator.py fix-encoding ./spec --ext .md`

---

## Execution Checklist for the AI

Before completing this task, you MUST verify:

- [ ] I saved the script precisely to `.lovable/ai-fix-scripts/01-file-manipulator.py`.
- [ ] I used `argparse` to handle subcommands (`lowercase`, `fix-seq-files`, and `fix-encoding`) and provided detailed help text with examples.
- [ ] `node_modules`, `.git`, and `.gitmap` are hardcoded into the default ignore list for all commands.
- [ ] Renames use `git mv` where applicable to preserve history.
- [ ] I implemented the pinning (`--pin`) logic for sequences.
- [ ] I handled Windows long paths properly via path normalization.
- [ ] I successfully implemented the encoding fix (BOM stripping and CRLF to LF normalization).
- [ ] I enforced strict lowercase `.md` extensions.
- [ ] I updated `.lovable/ai-fix-scripts/01-index.md` with instructions on how to use this script.
- [ ] I pre-compiled all regex patterns at module level for optimal speed.
- [ ] I did NOT leave any TODO placeholders in the generated Python code.

## No Automatic Releases (Strict Policy)

You MUST NOT bump versions, update changelogs, or cut a release at the end of this task. Commits must remain standard development commits. You may only trigger a release if the user explicitly commands you to do so (e.g., "cut a release" or "bump the version").

## STRICT AVOIDANCE: Never Disable CI/CD

> [!CAUTION]
> **NEVER disable any CI/CD checks, GitHub Actions, or validation workflows.**
> Strictly avoid commenting out, bypassing, or deleting CI/CD steps to force a pipeline to pass. Your job is to fix the underlying code so that the CI/CD pipeline passes legitimately. Disabling CI/CD is an auto-reject failure.
