#!/usr/bin/env python3
"""
Fast Repository Artifact Remover & Git Cleanup Guard
Safely discovers, previews, and permanently deletes unneeded test artifacts,
binary blobs, __pycache__, and temporary files from both the filesystem and Git index (`git rm`).

All Enums, Artifact Presets, Git Constants, and Utility Functions are imported
directly from 02-shared-engine.py as the single source of truth.

Usage:
  # 1. Preview specific file or glob pattern
  python .lovable/ai-fix-scripts/19-artifact-remover.py tests/visual/baselines/plan69/run.png --dry-run

  # 2. Remove specific file with interactive confirmation
  python .lovable/ai-fix-scripts/19-artifact-remover.py tests/visual/baselines/plan69/run.png

  # 3. Clean all pycache and test cache artifacts
  python .lovable/ai-fix-scripts/19-artifact-remover.py --clean-pycache [--force]

  # 4. Clean temporary files (.tmp, .log, .swp, .bak, .DS_Store)
  python .lovable/ai-fix-scripts/19-artifact-remover.py --clean-temp [--force]

  # 5. Clean unapproved binary files and blobs
  python .lovable/ai-fix-scripts/19-artifact-remover.py --clean-binaries [--force]

  # 6. Add custom paths, extensions, or glob patterns dynamically
  python .lovable/ai-fix-scripts/19-artifact-remover.py --add-path build/,dist/ --add-ext .tmp,.log --force
"""

import argparse
from importlib import import_module
import os
from pathlib import Path
import shutil
import subprocess
import sys
import time

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")

sys.path.insert(0, str(Path(__file__).parent))
engine = import_module("02-shared-engine")

# Centralized Enums
ArtifactCategoryType = engine.ArtifactCategoryType
ExitCodeType = engine.ExitCodeType

# Centralized String Literals & Separators
CURRENT_DIR = engine.CURRENT_DIR
EMPTY_STRING = engine.EMPTY_STRING
LINE_SEPARATOR = engine.LINE_SEPARATOR
PATH_SEPARATOR = engine.PATH_SEPARATOR
DEFAULT_ENCODING = engine.DEFAULT_ENCODING

# Centralized Git Command Constants
GIT_EXECUTABLE = engine.GIT_EXECUTABLE
GIT_CMD_LS_FILES = engine.GIT_CMD_LS_FILES
GIT_CMD_RM = engine.GIT_CMD_RM
GIT_FLAG_FORCE = engine.GIT_FLAG_FORCE
GIT_FLAG_ERROR_UNMATCH = engine.GIT_FLAG_ERROR_UNMATCH

# Centralized Status Badges
STATUS_GIT_TRACKED = engine.STATUS_GIT_TRACKED
STATUS_UNTRACKED = engine.STATUS_UNTRACKED

# Centralized Preset Artifact Collections
PYCACHE_DIR_NAMES = engine.PYCACHE_DIR_NAMES
PYCACHE_FILE_EXTENSIONS = engine.PYCACHE_FILE_EXTENSIONS
TEMP_ARTIFACT_EXTENSIONS = engine.TEMP_ARTIFACT_EXTENSIONS
TEMP_ARTIFACT_FILENAMES = engine.TEMP_ARTIFACT_FILENAMES
BINARY_EXTENSIONS = engine.BINARY_EXTENSIONS
EXCLUDE_DIRS = engine.EXCLUDE_DIRS

# Centralized Helper Functions
is_ignored_directory = engine.is_ignored_directory
is_binary_file = engine.is_binary_file
is_allowed_large_file = engine.is_allowed_large_file
normalize_rel_path = engine.normalize_rel_path
normalize_extensions = engine.normalize_extensions

def is_git_tracked(file_path: Path) -> bool:
    """
    Checks whether a given file is actively tracked in the Git index.
    Executes `git ls-files --error-unmatch <path>` with exit code inspection.
    """
    try:
        res = subprocess.run(
            [GIT_EXECUTABLE, GIT_CMD_LS_FILES, GIT_FLAG_ERROR_UNMATCH, str(file_path)],
            capture_output=True,
            text=True,
            encoding=DEFAULT_ENCODING,
            errors="replace"
        )
        return (res.returncode == 0)
    except Exception:
        return False

def remove_from_git_and_disk(file_path: Path) -> bool:
    """
    Safely removes a file or directory from both the Git index and the local filesystem.
    - For Git-tracked items: executes `git rm -f <path>` to stage deletion immediately.
    - For untracked items: deletes directly via `unlink` or `rmtree`.
    """
    is_tracked = is_git_tracked(file_path)

    if is_tracked:
        try:
            res = subprocess.run(
                [GIT_EXECUTABLE, GIT_CMD_RM, GIT_FLAG_FORCE, str(file_path)],
                capture_output=True,
                text=True,
                encoding=DEFAULT_ENCODING,
                errors="replace"
            )
            is_success = (res.returncode == 0)
            if is_success:
                return True
        except Exception:
            pass

    try:
        if file_path.is_file() or file_path.is_symlink():
            file_path.unlink(missing_ok=True)
            return True
        if file_path.is_dir():
            shutil.rmtree(file_path, ignore_errors=True)
            return True
    except Exception:
        pass

    return False

def collect_matching_artifacts(
    targets: list[str] | None = None,
    add_paths: set[str] | None = None,
    add_exts: set[str] | None = None,
    add_patterns: list[str] | None = None,
    is_clean_pycache: bool = False,
    is_clean_binaries: bool = False,
    is_clean_temp: bool = False,
    root_dir: str = CURRENT_DIR
) -> list[Path]:
    """
    Collects and deduplicates candidate files and directories to be removed.
    Supports preset categories (pycache, binaries, temp) as well as extensible CLI paths and extensions.
    """
    candidates: list[Path] = []
    root_path = Path(root_dir)

    # 1. Direct explicit file/folder targets
    if targets:
        for t in targets:
            clean_t = t.strip()
            if not clean_t or clean_t in {CURRENT_DIR, f"{CURRENT_DIR}{PATH_SEPARATOR}"}:
                continue
            p_target = Path(clean_t)
            if p_target.exists():
                candidates.append(p_target)
                if p_target.is_dir():
                    for r, dirs, files in os.walk(p_target, topdown=False):
                        for f in files:
                            candidates.append(Path(r) / f)
            else:
                for match in root_path.glob(f"**/{clean_t}"):
                    if not any(part in match.parts for part in {".git", ".gitmap", "node_modules"}):
                        candidates.append(match)

    # 2. Extensible Custom Paths (--add-path)
    if add_paths:
        for p_str in add_paths:
            p_obj = Path(p_str)
            if p_obj.exists():
                candidates.append(p_obj)
                if p_obj.is_dir():
                    for r, dirs, files in os.walk(p_obj, topdown=False):
                        for f in files:
                            candidates.append(Path(r) / f)

    # 3. Pycache & Bytecode Preset (--clean-pycache)
    if is_clean_pycache:
        for r, dirs, files in os.walk(root_dir, topdown=False):
            dirs[:] = [d for d in dirs if d not in {".git", ".gitmap", "node_modules"}]
            dir_name = Path(r).name
            if dir_name in PYCACHE_DIR_NAMES:
                candidates.append(Path(r))
            for f in files:
                ext = os.path.splitext(f)[1].lower()
                if ext in PYCACHE_FILE_EXTENSIONS:
                    candidates.append(Path(r) / f)

    # 4. Temporary Artifacts Preset (--clean-temp)
    if is_clean_temp:
        for r, dirs, files in os.walk(root_dir, topdown=False):
            dirs[:] = [d for d in dirs if not is_ignored_directory(d)]
            for f in files:
                ext = os.path.splitext(f)[1].lower()
                if ext in TEMP_ARTIFACT_EXTENSIONS or f in TEMP_ARTIFACT_FILENAMES:
                    candidates.append(Path(r) / f)

    # 5. Unapproved Binary Blobs Preset (--clean-binaries)
    if is_clean_binaries:
        for r, dirs, files in os.walk(root_dir, topdown=False):
            dirs[:] = [d for d in dirs if not is_ignored_directory(d)]
            for f in files:
                p_file = Path(r) / f
                if is_allowed_large_file(p_file):
                    continue
                if is_binary_file(p_file):
                    candidates.append(p_file)

    # 6. Extensible Custom Extension Filter (--add-ext)
    if add_exts:
        for r, dirs, files in os.walk(root_dir, topdown=False):
            dirs[:] = [d for d in dirs if not is_ignored_directory(d)]
            for f in files:
                ext = os.path.splitext(f)[1].lower()
                if ext in add_exts:
                    candidates.append(Path(r) / f)

    # 7. Extensible Glob Patterns (--add-pattern)
    if add_patterns:
        for pat in add_patterns:
            for match in root_path.glob(f"**/{pat}"):
                if not any(part in match.parts for part in {".git", ".gitmap", "node_modules"}):
                    candidates.append(match)

    return sorted(list(set(candidates)))

def confirm_action(prompt_text: str) -> bool:
    """Prompts user for interactive confirmation."""
    try:
        response = input(f"{prompt_text} (y/N): ").strip().lower()
        return response in {"y", "yes"}
    except (EOFError, KeyboardInterrupt):
        return False

def run_artifact_remover(
    targets: list[str] | None = None,
    add_paths: set[str] | None = None,
    add_exts: set[str] | None = None,
    add_patterns: list[str] | None = None,
    is_clean_pycache: bool = False,
    is_clean_binaries: bool = False,
    is_clean_temp: bool = False,
    is_force_mode: bool = False,
    is_dry_run_mode: bool = False,
    root_dir: str = CURRENT_DIR
) -> int:
    """
    Master orchestrator:
    1. Collects all matching items.
    2. Displays full target preview with Git Tracked / Untracked badges.
    3. Prompts user for interactive confirmation (unless --force).
    4. Executes atomic Git & disk removal.
    """
    start_time = time.perf_counter()

    artifacts = collect_matching_artifacts(
        targets=targets,
        add_paths=add_paths,
        add_exts=add_exts,
        add_patterns=add_patterns,
        is_clean_pycache=is_clean_pycache,
        is_clean_binaries=is_clean_binaries,
        is_clean_temp=is_clean_temp,
        root_dir=root_dir
    )

    has_artifacts = len(artifacts) > 0
    if not has_artifacts:
        print("✅ No matching artifacts found for the specified criteria.")
        return ExitCodeType.SUCCESS.value

    # Display Preview
    print("================================================================================")
    print(f"🗑️ Artifact Removal Target: Found **{len(artifacts)}** candidate item(s)")
    print("================================================================================")
    for idx, item in enumerate(artifacts[:25], start=1):
        badge = STATUS_GIT_TRACKED if is_git_tracked(item) else STATUS_UNTRACKED
        print(f"   {idx:>3}. {normalize_rel_path(item)}{badge}")
    if len(artifacts) > 25:
        print(f"   ... and {len(artifacts) - 25} more items.")

    if is_dry_run_mode:
        print(f"{LINE_SEPARATOR}ℹ️ Dry-run mode enabled. No items were deleted.")
        return ExitCodeType.SUCCESS.value

    # Require Interactive Confirmation unless --force is specified
    if not is_force_mode:
        print(f"{LINE_SEPARATOR}⚠️ CAUTION: Deleting these files removes them from the filesystem and Git index.")
        is_confirmed = confirm_action(f"Are you sure you want to permanently remove {len(artifacts)} item(s)?")
        if not is_confirmed:
            print("❌ Operation cancelled by user. No files were modified.")
            return ExitCodeType.SUCCESS.value

    # Execute Removal
    removed_count = 0
    for item in artifacts:
        is_removed = remove_from_git_and_disk(item)
        if is_removed:
            removed_count += 1
            print(f"  ✓ Removed: {normalize_rel_path(item)}")

    elapsed_ms = (time.perf_counter() - start_time) * 1000
    print(f"{LINE_SEPARATOR}✅ Successfully removed {removed_count} of {len(artifacts)} item(s) in {elapsed_ms:.2f}ms.")
    return ExitCodeType.SUCCESS.value

def main():
    parser = argparse.ArgumentParser(
        description="Fast Repository Artifact Remover & Git Cleanup Guard",
        formatter_class=argparse.RawDescriptionHelpFormatter
    )
    parser.add_argument("targets", nargs="*", default=[], help="Target files, directories, or patterns to remove (e.g. tests/visual/run.png)")
    parser.add_argument("--add-path", help="Comma-separated additional folders/paths to remove")
    parser.add_argument("--add-ext", help="Comma-separated additional file extensions to remove (e.g. .tmp,.log)")
    parser.add_argument("--add-pattern", help="Comma-separated additional glob patterns to remove")
    parser.add_argument("--clean-pycache", action="store_true", help="Remove __pycache__, .pytest_cache, and .pyc/.pyo files")
    parser.add_argument("--clean-temp", action="store_true", help="Remove temporary files (.tmp, .log, .swp, .bak, .DS_Store)")
    parser.add_argument("--clean-binaries", action="store_true", help="Remove unapproved binary blobs and image artifacts")
    parser.add_argument("--clean-all", action="store_true", help="Enable pycache, temp files, and binary cleanup presets")
    parser.add_argument("--force", "-f", action="store_true", help="Bypass interactive confirmation prompt")
    parser.add_argument("--dry-run", "-d", action="store_true", help="Preview matching items without deleting")
    parser.add_argument("--dir", default=CURRENT_DIR, help="Root directory to search (default: .)")
    args = parser.parse_args()

    # Parse extensible CLI inputs
    custom_paths = {p.strip() for p in args.add_path.split(",") if p.strip()} if args.add_path else set()
    custom_exts = normalize_extensions(args.add_ext)
    custom_patterns = [p.strip() for p in args.add_pattern.split(",") if p.strip()] if args.add_pattern else []

    is_clean_pycache = args.clean_pycache or args.clean_all
    is_clean_temp = args.clean_temp or args.clean_all
    is_clean_binaries = args.clean_binaries or args.clean_all

    has_work = bool(
        args.targets or custom_paths or custom_exts or custom_patterns or
        is_clean_pycache or is_clean_temp or is_clean_binaries
    )
    if not has_work:
        parser.print_help()
        sys.exit(0)

    sys.exit(run_artifact_remover(
        targets=args.targets,
        add_paths=custom_paths,
        add_exts=custom_exts,
        add_patterns=custom_patterns,
        is_clean_pycache=is_clean_pycache,
        is_clean_binaries=is_clean_binaries,
        is_clean_temp=is_clean_temp,
        is_force_mode=args.force,
        is_dry_run_mode=args.dry_run,
        root_dir=args.dir
    ))

if __name__ == "__main__":
    main()
