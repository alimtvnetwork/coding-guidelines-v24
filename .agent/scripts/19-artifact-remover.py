#!/usr/bin/env python3
"""
Fast Repository Artifact Remover & Git Cleanup Guard
Safely finds, previews, and deletes unneeded test artifacts, binary blobs, pycache,
and temporary files from both the filesystem and Git index (`git rm`).

Includes interactive confirmation prompts, `--dry-run` preview, and `--force` flag.
All Enums, Constants, and Functions are imported directly from 02-shared-engine.py.

Usage:
  python .lovable/ai-fix-scripts/19-artifact-remover.py <path-or-pattern> [--dry-run]
  python .lovable/ai-fix-scripts/19-artifact-remover.py <path-or-pattern> --force
  python .lovable/ai-fix-scripts/19-artifact-remover.py --clean-pycache [--force]
  python .lovable/ai-fix-scripts/19-artifact-remover.py --clean-binaries [--force]
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

# Centralized Enums & Constants
ExitCodeType = engine.ExitCodeType
CURRENT_DIR = engine.CURRENT_DIR
EMPTY_STRING = engine.EMPTY_STRING
LINE_SEPARATOR = engine.LINE_SEPARATOR
PATH_SEPARATOR = engine.PATH_SEPARATOR
DEFAULT_ENCODING = engine.DEFAULT_ENCODING
BINARY_EXTENSIONS = engine.BINARY_EXTENSIONS
EXCLUDE_DIRS = engine.EXCLUDE_DIRS
is_ignored_directory = engine.is_ignored_directory
is_binary_file = engine.is_binary_file
is_allowed_large_file = engine.is_allowed_large_file
normalize_rel_path = engine.normalize_rel_path
normalize_extensions = engine.normalize_extensions

# Artifact Preset Target Folders & Extensions
PYCACHE_PATTERNS = {"__pycache__", ".pytest_cache", ".coverage", ".mypy_cache", ".ruff_cache"}
TEMP_ARTIFACT_PATTERNS = {".tmp", ".log", ".swp", ".bak", ".pyc", ".pyo"}

def is_git_tracked(file_path: Path) -> bool:
    """Checks if a file is currently tracked by Git."""
    try:
        res = subprocess.run(
            ["git", "ls-files", "--error-unmatch", str(file_path)],
            capture_output=True,
            text=True,
            encoding=DEFAULT_ENCODING,
            errors="replace"
        )
        return (res.returncode == 0)
    except Exception:
        return False

def remove_from_git_and_disk(file_path: Path) -> bool:
    """Removes a file from Git index and filesystem using `git rm -f` or `os.remove`."""
    norm_p = normalize_rel_path(file_path)
    is_tracked = is_git_tracked(file_path)

    if is_tracked:
        try:
            res = subprocess.run(
                ["git", "rm", "-f", str(file_path)],
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
    target: str = EMPTY_STRING,
    is_clean_pycache: bool = False,
    is_clean_binaries: bool = False,
    root_dir: str = CURRENT_DIR
) -> list[Path]:
    """Collects candidate files and directories to be removed."""
    candidates: list[Path] = []
    clean_target = target.strip()

    # 1. Direct file or directory path
    if clean_target and clean_target not in {".", "./"}:
        p_target = Path(clean_target)
        if p_target.exists():
            candidates.append(p_target)
            if p_target.is_dir():
                for root, dirs, files in os.walk(p_target, topdown=False):
                    for f in files:
                        candidates.append(Path(root) / f)
            return sorted(list(set(candidates)))

        # Try glob matching if direct path does not exist
        root_path = Path(root_dir)
        for match in root_path.glob(f"**/{clean_target}"):
            if not any(part in match.parts for part in {".git", ".gitmap", "node_modules"}):
                candidates.append(match)

    # 2. Pycache sweep
    if is_clean_pycache:
        for root, dirs, files in os.walk(root_dir, topdown=False):
            dirs[:] = [d for d in dirs if d not in {".git", ".gitmap", "node_modules"}]
            dir_name = Path(root).name
            if dir_name in PYCACHE_PATTERNS:
                candidates.append(Path(root))
            for f in files:
                ext = os.path.splitext(f)[1].lower()
                if ext in {".pyc", ".pyo"}:
                    candidates.append(Path(root) / f)

    # 3. Binaries sweep (excluding approved assets)
    if is_clean_binaries:
        for root, dirs, files in os.walk(root_dir, topdown=False):
            dirs[:] = [d for d in dirs if not is_ignored_directory(d)]
            for f in files:
                p_file = Path(root) / f
                if is_allowed_large_file(p_file):
                    continue
                if is_binary_file(p_file):
                    candidates.append(p_file)

    return sorted(list(set(candidates)))

def confirm_action(prompt_text: str) -> bool:
    """Prompts user for interactive confirmation."""
    try:
        response = input(f"{prompt_text} (y/N): ").strip().lower()
        return response in {"y", "yes"}
    except (EOFError, KeyboardInterrupt):
        return False

def run_artifact_remover(
    target: str = EMPTY_STRING,
    is_clean_pycache: bool = False,
    is_clean_binaries: bool = False,
    is_force_mode: bool = False,
    is_dry_run_mode: bool = False,
    root_dir: str = CURRENT_DIR
) -> int:
    """Executes artifact discovery, preview, confirmation, and safe deletion."""
    start_time = time.perf_counter()

    artifacts = collect_matching_artifacts(
        target=target,
        is_clean_pycache=is_clean_pycache,
        is_clean_binaries=is_clean_binaries,
        root_dir=root_dir
    )

    has_artifacts = len(artifacts) > 0
    if not has_artifacts:
        target_name = target if target else "presets"
        print(f"✅ No matching artifacts found for target: '{target_name}'")
        return ExitCodeType.SUCCESS.value

    # Display Preview
    print("================================================================================")
    print(f"🗑️ Artifact Removal Target: Found **{len(artifacts)}** candidate item(s)")
    print("================================================================================")
    for idx, item in enumerate(artifacts[:25], start=1):
        tracked_badge = " [Git Tracked]" if is_git_tracked(item) else " [Untracked]"
        print(f"   {idx:>3}. {normalize_rel_path(item)}{tracked_badge}")
    if len(artifacts) > 25:
        print(f"   ... and {len(artifacts) - 25} more items.")

    if is_dry_run_mode:
        print(f"{LINE_SEPARATOR}ℹ️ Dry-run mode enabled. No items were deleted.")
        return ExitCodeType.SUCCESS.value

    # Require Confirmation unless --force is provided
    if not is_force_mode:
        print(f"{LINE_SEPARATOR}⚠️ CAUTION: Deleting these files removes them from the filesystem and Git index.")
        is_confirmed = confirm_action(f"Are you sure you want to permanently remove {len(artifacts)} item(s)?")
        if not is_confirmed:
            print("❌ Operation cancelled by user. No files were modified.")
            return ExitCodeType.SUCCESS.value

    # Execute Safe Removal
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
    parser.add_argument("target", nargs="?", default=EMPTY_STRING, help="Path, directory, or pattern to remove (e.g. tests/visual/run.png, *.pyc)")
    parser.add_argument("--path", "-p", dest="opt_path", help="Alternative target path")
    parser.add_argument("--clean-pycache", action="store_true", help="Remove all __pycache__, .pytest_cache, and .pyc files")
    parser.add_argument("--clean-binaries", action="store_true", help="Remove untracked binary blobs and images")
    parser.add_argument("--force", "-f", action="store_true", help="Bypass interactive confirmation prompt")
    parser.add_argument("--dry-run", "-d", action="store_true", help="Preview items without deleting")
    parser.add_argument("--dir", default=CURRENT_DIR, help="Root directory to search (default: .)")
    args = parser.parse_args()

    target_str = args.opt_path or args.target or EMPTY_STRING
    has_target = bool(target_str or args.clean_pycache or args.clean_binaries)
    if not has_target:
        parser.print_help()
        sys.exit(0)

    sys.exit(run_artifact_remover(
        target=target_str,
        is_clean_pycache=args.clean_pycache,
        is_clean_binaries=args.clean_binaries,
        is_force_mode=args.force,
        is_dry_run_mode=args.dry_run,
        root_dir=args.dir
    ))

if __name__ == "__main__":
    main()
