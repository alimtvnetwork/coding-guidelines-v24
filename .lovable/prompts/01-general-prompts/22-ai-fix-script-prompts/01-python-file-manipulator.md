# Python File Manipulator CLI Specification — Tooling Spec (must follow)


> **Prompt Version:** 2.1.0
> **Synchronization:** Main Meta-Repo & Connected Workspaces

/goal Autonomously generate a robust, dependency-free Python CLI tool to handle mass file renaming, sequencing, and encoding normalization.

## Overview

You are an expert Python Developer AI. Your task is to write a standalone, reusable Python script that handles mass file renaming (lowercasing), sequence fixing, and encoding normalization. This script will act as an autonomous tool for other AIs and developers to organize files without needing a compiled binary.

**Target Path:** .lovable/ai-fix-scripts/01-file-manipulator.py

## Non-Negotiable Rules for the Python Script

1. **Zero Dependencies**: The script MUST use only Python standard libraries (e.g., os, sys, rgparse, shutil, subprocess, pathlib).
2. **Robust CLI**: Use rgparse to provide a professional, CLI-like experience with complete --help documentation and examples.
3. **Windows Long Paths**: The script must normalize paths and safely handle Windows MAX_PATH limitations (e.g., prefixing absolute paths with \\?\ on Windows environments).
4. **Git Awareness**: Whenever renaming a file, the script must attempt to use git mv via subprocess first. If the file is untracked or the command fails, gracefully fallback to standard os.rename.
5. **Update Index**: After generating the script, you MUST document its usage in .lovable/ai-fix-scripts/index.md.

---

## Core Feature 1: Lowercase Renamer

**Command Pattern**:
python 01-file-manipulator.py lowercase <target_directory> [flags]

**Requirements**:

1. Recursively convert all files matching a target pattern to lowercase.
2. **Extension Enforcement**: It MUST strictly ensure that all markdown file extensions are lowercased (e.g., converting .MD to .md), adhering to our pattern systems.
3. **Default Ignores**: By default, the script MUST silently ignore 
ode_modules and .git folders. Do not traverse them.
4. **Extendable Ignores**: Provide an --except flag accepting a comma-separated list of additional files, folders, or wildcard patterns to ignore (e.g., --except "docs/*, temp.md").

**Example Output in --help**:

- python 01-file-manipulator.py lowercase ./src (Ignores node_modules/.git by default)
- python 01-file-manipulator.py lowercase ./src --except "vendor/*, build/*"

---

## Core Feature 2: Fix File Sequencing (`fix-seq-files`)

**Command Pattern**:
python 01-file-manipulator.py fix-seq-files <target_directory> [flags]

**Requirements**:

1. Scan the specified directory for sequenced files (e.g., 01-draft.md, 02-notes.md).
2. **Ordering Flags**:
   - --order-by-time: Re-sequence files sequentially based on their filesystem modification time.
   - --order-by-az: Re-sequence files alphabetically based on the string following the sequence number.
3. **Tie-Breaker / Preservation**:
   - --keep-old-order: Preserve existing numeric ordering as much as possible. Only assign new sequence numbers to unnumbered files or resolve direct conflicts using time/alphabetization.
4. **Fixated / Pinned Sequences**:
   - --pin "<mapping>": Allow users to explicitly lock specific files to a sequence number. (e.g., --pin "readme=00,draft=01"). The script must increment other files around these locked sequences.

**Example Output in --help**:

- python 01-file-manipulator.py fix-seq-files ./docs --order-by-time
- python 01-file-manipulator.py fix-seq-files ./docs --order-by-az --keep-old-order
- python 01-file-manipulator.py fix-seq-files ./docs --pin "readme=00,intro=01"

---

## Core Feature 3: Fix Encoding & Line Endings (`fix-encoding`)

**Command Pattern**:
python 01-file-manipulator.py fix-encoding <target_directory> [flags]

**Requirements**:

1. Scan the specified directory and aggressively fix encoding issues for all text files (specifically targeting .md files).
2. **BOM Stripping**: Detect and strip any UTF-8 Byte Order Marks (BOM) or UTF-16 encodings, standardizing everything strictly to UTF-8 without BOM.
3. **Line Ending Normalization**: Automatically convert all Windows CRLF (\r\n) line endings to Unix LF (\n) to prevent git warnings and cross-platform issues.

**Example Output in --help**:

- python 01-file-manipulator.py fix-encoding ./src

---

## Execution Checklist for the AI

Before completing this task, you MUST verify:

- [ ] I saved the script precisely to .lovable/ai-fix-scripts/01-file-manipulator.py.
- [ ] I used rgparse to handle subcommands (lowercase, 
ix-seq-files, and 
ix-encoding) and provided detailed help text with examples.
- [ ] 
ode_modules and .git are hardcoded into the default ignore list for all commands.
- [ ] Renames use git mv where applicable to preserve history.
- [ ] I implemented the pinning (--pin) logic for sequences.
- [ ] I handled Windows long paths properly via path normalization.
- [ ] I successfully implemented the encoding fix (BOM stripping and CRLF to LF normalization).
- [ ] I enforced strict lowercase .md extensions.
- [ ] I updated .lovable/ai-fix-scripts/index.md with instructions on how to use this new script.
- [ ] I did NOT leave any TODO placeholders in the generated Python code.

## No Automatic Releases (Strict Policy)

You MUST NOT bump versions, update changelogs, or cut a release at the end of this task. Commits must remain standard development commits. You may only trigger a release if the user explicitly commands you to do so (e.g., "cut a release" or "bump the version").


## STRICT AVOIDANCE: Never Disable CI/CD

> [!CAUTION]
> **NEVER disable any CI/CD checks, GitHub Actions, or validation workflows.** 
> Strictly avoid commenting out, bypassing, or deleting CI/CD steps to force a pipeline to pass. Your job is to fix the underlying code so that the CI/CD pipeline passes legitimately. Disabling CI/CD is an auto-reject failure.

## MUST FOLLOW NON-NEGOTIABLE

Listen, past runs of these turns have been sloppy and stupid as fuck: wrong step counts, partial task lists dumped into chat instead of files, plans and session summaries half-filled with "[N]" placeholders, folders skimmed, open ambiguities ignored, CI/CD issues and plans/subtasks/ forgotten, user commands dropped, coding guidelines bypassed, detailed specs chopped and summarized into useless junk, uppercase README files left uncorrected, .lovable/memories/ created by accident, strictly-avoid.md overwritten, and explicit user instructions softened after being told not to. WTF. How on earth are you reverting to this carelessness, are you stupid?? Stop doing that, you stupid fuck. Read the whole codebase, read every folder in spec/ and .lovable/, confirm root 
eadme.md is strictly lowercase, find the root cause in one sentence, capture commands, issues, and pending tasks without omitting a single item, write the spec files and memory files in the right paths, update every index in the same turn, sync 
eadme.md with what-to-read.md, preserve detailed specs verbatim with zero truncation, run builds and full unit tests, group commits with clear messages, and push everything to git before ending. Going deep IS the job. If you are not going deep, you are not doing the job. Violating this is auto-reject on the same tier as RULE 0. Avoid stupidity and being careless, you stupid fuck. Where is your attention, are you stupid? Tell me. Your stupidity is going on top of my head. Where did you learn this stupidity? If I could find you, I could slap you.
