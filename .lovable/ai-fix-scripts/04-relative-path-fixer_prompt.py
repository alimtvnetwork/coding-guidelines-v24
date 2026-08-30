#!/usr/bin/env python3
"""
Relative Path Fixer & Validator.
Scans the codebase for absolute filesystem paths (e.g. D:\\..., C:\\..., /home/...)
and file:/// URIs in repository files, plans, specs, citations, and markdown docs,
and automatically converts them to strictly relative Git repository paths.

Usage:
  python .lovable/ai-fix-scripts/04-relative-path-fixer.py [path]
  python .lovable/ai-fix-scripts/04-relative-path-fixer.py --check
  python .lovable/ai-fix-scripts/04-relative-path-fixer.py --dry-run
"""

import argparse
import os
import re
import subprocess
import sys
import time

IGNORE_DIRS = {
    ".git",
    "node_modules",
    ".venv",
    "venv",
    "vendor",
    "build",
    "dist",
    "bin",
    ".gemini",
    "__pycache__",
}

IGNORE_FILES = {
    "04-relative-path-fixer.py",
}

VALID_EXTENSIONS = {
    ".md",
    ".go",
    ".ts",
    ".tsx",
    ".js",
    ".jsx",
    ".py",
    ".php",
    ".cs",
    ".json",
    ".yaml",
    ".yml",
    ".toml",
    ".sh",
    ".ps1",
    ".txt",
}

BAD_EXAMPLE_MARKERS = [
    "❌",
    "BAD",
    "INVALID",
    "FORBIDDEN",
    "ANTI-PATTERN",
    "INCORRECT",
]


def get_git_root() -> str:
    """Discover the root directory of the current git repository."""
    try:
        result = subprocess.run(
            ["git", "rev-parse", "--show-toplevel"],
            capture_output=True,
            text=True,
            check=True,
        )
        return os.path.abspath(result.stdout.strip())
    except Exception:
        return os.path.abspath(os.getcwd())


def normalize_separators(path_str: str) -> str:
    return path_str.replace("\\", "/")


def build_replacement_patterns(repo_root: str):
    norm_root = normalize_separators(repo_root)
    drive_match = re.match(r"^([a-zA-Z]):", norm_root)
    drive_letter = drive_match.group(1) if drive_match else ""

    patterns = []

    # 1. file:///d:/repo/path or file:///D:/repo/path
    if drive_letter:
        prefix_lower = f"file:///{drive_letter.lower()}:" + norm_root[2:]
        prefix_upper = f"file:///{drive_letter.upper()}:" + norm_root[2:]
        patterns.append((prefix_lower.rstrip("/") + "/", ""))
        patterns.append((prefix_upper.rstrip("/") + "/", ""))

    # 2. file:///path (Unix style)
    patterns.append((f"file://{norm_root}".rstrip("/") + "/", ""))
    patterns.append((f"file:///{norm_root}".rstrip("/") + "/", ""))

    # 3. Windows drive path: D:/repo/path or d:/repo/path or D:\repo\path
    patterns.append((norm_root.rstrip("/") + "/", ""))
    patterns.append((norm_root.lower().rstrip("/") + "/", ""))
    patterns.append((repo_root.rstrip("\\") + "\\", ""))
    patterns.append((repo_root.lower().rstrip("\\") + "\\", ""))

    return patterns


def is_bad_example_context(line: str, in_bad_codeblock: bool) -> bool:
    """Detect if a line is an intentional documentation demonstration of a bad/invalid example."""
    if in_bad_codeblock:
        return True
    return any(marker in line for marker in BAD_EXAMPLE_MARKERS)


def fix_line(line: str, repo_root: str, patterns, is_markdown: bool) -> tuple[str, int]:
    modified_line = line
    total_matches = 0
    norm_root = normalize_separators(repo_root)

    # 1. Replace explicit repo prefixes
    for prefix, repl in patterns:
        if prefix in modified_line:
            count = modified_line.count(prefix)
            total_matches += count
            modified_line = modified_line.replace(prefix, repl)

    # 2. In markdown files, catch file:/// URLs inside markdown links [Label](file:///...)
    if is_markdown:
        def md_link_sub(match):
            nonlocal total_matches
            label = match.group(1)
            raw_url = match.group(2)
            norm_url = normalize_separators(raw_url)

            if norm_url.startswith("file:///"):
                stripped = norm_url[8:]
                if re.match(r"^[a-zA-Z]:/", stripped):
                    if stripped.lower().startswith(norm_root.lower() + "/"):
                        rel = stripped[len(norm_root) + 1 :]
                        total_matches += 1
                        return f"[{label}]({rel})"
                    rel = stripped[3:]
                    total_matches += 1
                    return f"[{label}]({rel})"
                elif stripped.startswith("/"):
                    rel = stripped.lstrip("/")
                    total_matches += 1
                    return f"[{label}]({rel})"
            return match.group(0)

        modified_line = re.sub(
            r"\[([^\]]+)\]\((file:///[^)]+)\)", md_link_sub, modified_line
        )

        # Clean backslashes in markdown links: [Label](dir\file.md) -> [Label](dir/file.md)
        def clean_md_backslashes(match):
            nonlocal total_matches
            label = match.group(1)
            target = match.group(2)
            if "\\" in target and not target.startswith("http") and ("/" in target or "\\" in target or "." in target):
                cleaned = normalize_separators(target)
                if cleaned != target:
                    total_matches += 1
                return f"[{label}]({cleaned})"
            return match.group(0)

        modified_line = re.sub(
            r"\[([^\]]+)\]\(([a-zA-Z0-9_\-\.\\]+\\[a-zA-Z0-9_\-\.\\]+)\)", clean_md_backslashes, modified_line
        )

    return modified_line, total_matches


def fix_content(content: str, repo_root: str, patterns, is_markdown: bool = True) -> tuple[str, int]:
    """Scan line-by-line, preserving explicit bad-example documentation snippets."""
    lines = content.splitlines(keepends=True)
    new_lines = []
    total_matches = 0
    in_bad_codeblock = False

    for line in lines:
        stripped = line.strip()
        # Track bad code block headers (e.g. #### ❌ INVALID)
        if any(marker in stripped for marker in BAD_EXAMPLE_MARKERS) and (
            stripped.startswith("#") or stripped.startswith("```") or stripped.startswith("-")
        ):
            if "INVALID" in stripped or "BAD" in stripped or "❌" in stripped:
                in_bad_codeblock = True
            else:
                in_bad_codeblock = False
        elif (
            "VALID" in stripped
            or "GOOD" in stripped
            or "✅" in stripped
            or "REQUIRED" in stripped
        ):
            in_bad_codeblock = False

        if is_bad_example_context(line, in_bad_codeblock):
            new_lines.append(line)
            continue

        mod_line, count = fix_line(line, repo_root, patterns, is_markdown)
        total_matches += count
        new_lines.append(mod_line)

    return "".join(new_lines), total_matches


def process_file(
    file_path: str, repo_root: str, patterns, is_dry_run: bool, is_check_mode: bool
) -> tuple[int, bool]:
    try:
        with open(file_path, "r", encoding="utf-8") as f:
            content = f.read()
    except Exception:
        return 0, False

    is_md = file_path.endswith(".md")
    new_content, match_count = fix_content(content, repo_root, patterns, is_markdown=is_md)

    if match_count > 0 and not is_check_mode and not is_dry_run:
        with open(file_path, "w", encoding="utf-8") as f:
            f.write(new_content)

    return match_count, (match_count > 0)


def scan_and_fix(target_dir: str, is_dry_run: bool, is_check_mode: bool, is_verbose: bool):
    start_time = time.monotonic()
    repo_root = get_git_root()
    patterns = build_replacement_patterns(repo_root)

    total_scanned = 0
    total_violations = 0
    modified_files = []

    for root, dirs, files in os.walk(target_dir):
        dirs[:] = [d for d in dirs if d not in IGNORE_DIRS]
        for f in files:
            if f in IGNORE_FILES:
                continue
            ext = os.path.splitext(f)[1].lower()
            if ext not in VALID_EXTENSIONS:
                continue

            file_path = os.path.join(root, f)
            rel_path = normalize_separators(os.path.relpath(file_path, repo_root))
            total_scanned += 1

            matches, is_modified = process_file(
                file_path, repo_root, patterns, is_dry_run, is_check_mode
            )

            if matches > 0:
                total_violations += matches
                modified_files.append((rel_path, matches))
                if is_verbose or is_check_mode or is_dry_run:
                    action_tag = "VIOLATION" if is_check_mode else ("WOULD FIX" if is_dry_run else "FIXED")
                    print(f"[{action_tag}] {rel_path} ({matches} instances)")

    elapsed = round(time.monotonic() - start_time, 3)

    print("\n" + "=" * 60)
    print("RELATIVE PATH FIXER REPORT")
    print("=" * 60)
    print(f"Target Directory : {target_dir}")
    print(f"Git Root         : {repo_root}")
    print(f"Files Scanned    : {total_scanned}")
    print(f"Violations Found : {total_violations}")
    print(f"Files Modified   : {len(modified_files)}")
    print(f"Elapsed Time     : {elapsed}s")
    print("=" * 60)

    if is_check_mode:
        if total_violations > 0:
            print(f"\n[FAIL] Found {total_violations} absolute path / file URI violations across {len(modified_files)} files.")
            sys.exit(1)
        else:
            print("\n[PASS] All paths are strictly relative to git root.")
            sys.exit(0)


def main():
    parser = argparse.ArgumentParser(
        description="Scans and converts absolute paths / file:/// URIs to strictly relative Git root paths."
    )
    parser.add_argument(
        "path",
        nargs="?",
        default=".",
        help="Path or directory to scan and fix (default: current directory)",
    )
    parser.add_argument(
        "--check",
        action="store_true",
        help="Check only; exit code 1 if violations found (CI/CD mode)",
    )
    parser.add_argument(
        "--dry-run",
        action="store_true",
        help="Report violations without writing changes to disk",
    )
    parser.add_argument(
        "--verbose",
        "-v",
        action="store_true",
        help="Show detailed output per modified file",
    )

    args = parser.parse_args()
    scan_and_fix(
        target_dir=args.path,
        is_dry_run=args.dry_run,
        is_check_mode=args.check,
        is_verbose=args.verbose,
    )


if __name__ == "__main__":
    main()
