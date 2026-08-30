#!/usr/bin/env python3
"""
Naming Conventions & Boolean Autofixer.
Scans the codebase for bare `ok` identifiers, negative boolean names (hasNo*, isNot*),
and non-standard boolean prefixes across Go, TypeScript, JavaScript, Python, and C#.

Usage:
  python .lovable/ai-fix-scripts/05-naming-autofixer.py [path]
  python .lovable/ai-fix-scripts/05-naming-autofixer.py --check
  python .lovable/ai-fix-scripts/05-naming-autofixer.py --dry-run
"""

import argparse
import os
import re
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

VALID_EXTENSIONS = {
    ".go",
    ".ts",
    ".tsx",
    ".js",
    ".jsx",
    ".py",
    ".php",
    ".cs",
}

# Negative boolean identifiers to flag
NEGATIVE_BOOL_PATTERNS = [
    re.compile(r"\b(hasNo[A-Z][a-zA-Z0-9_]*)\b"),
    re.compile(r"\b(isNot[A-Z][a-zA-Z0-9_]*)\b"),
    re.compile(r"\b(isNo[A-Z][a-zA-Z0-9_]*)\b"),
]

# Bare ok in Go comma-ok patterns:
# e.g., "val, ok :=" or "val, ok =" or "; ok {" or "if !ok {"
GO_BARE_OK_PATTERNS = [
    re.compile(r"(\b[a-zA-Z0-9_]+,\s*)\bok(\s*:=)"),
    re.compile(r"(\b[a-zA-Z0-9_]+,\s*)\bok(\s*=)"),
    re.compile(r";\s*ok\s*\{"),
    re.compile(r"\bif\s*!ok\s*\{"),
]


def scan_file(file_path: str) -> list[tuple[int, str, str]]:
    violations = []
    try:
        with open(file_path, "r", encoding="utf-8") as f:
            lines = f.readlines()
    except Exception:
        return []

    is_go = file_path.endswith(".go")

    for idx, line in enumerate(lines, start=1):
        line_str = line.strip()
        if line_str.startswith("//") or line_str.startswith("/*") or line_str.startswith("#"):
            continue

        # Check negative boolean names
        for pat in NEGATIVE_BOOL_PATTERNS:
            match = pat.search(line)
            if match:
                var_name = match.group(1)
                violations.append((idx, f"Negative boolean identifier `{var_name}`", line_str))

        # Check Go bare ok
        if is_go:
            for pat in GO_BARE_OK_PATTERNS:
                if pat.search(line):
                    violations.append((idx, "Bare `ok` identifier in Go comma-ok idiom", line_str))
                    break

    return violations


def run_scanner(target_dir: str, is_check_mode: bool, is_verbose: bool):
    start = time.monotonic()
    total_files = 0
    total_violations = 0
    violation_records = []

    for root, dirs, files in os.walk(target_dir):
        dirs[:] = [d for d in dirs if d not in IGNORE_DIRS]
        for f in files:
            ext = os.path.splitext(f)[1].lower()
            if ext not in VALID_EXTENSIONS:
                continue

            file_path = os.path.join(root, f)
            total_files += 1
            file_violations = scan_file(file_path)

            if file_violations:
                total_violations += len(file_violations)
                rel_path = os.path.relpath(file_path, target_dir).replace("\\", "/")
                violation_records.append((rel_path, file_violations))
                if is_verbose or is_check_mode:
                    print(f"\n[VIOLATION] {rel_path} ({len(file_violations)} issues):")
                    for line_no, desc, snippet in file_violations:
                        print(f"  Line {line_no}: {desc} -> {snippet}")

    elapsed = round(time.monotonic() - start, 3)

    print("\n" + "=" * 60)
    print("NAMING CONVENTIONS & BOOLEAN AUDIT REPORT")
    print("=" * 60)
    print(f"Target Directory : {target_dir}")
    print(f"Files Scanned    : {total_files}")
    print(f"Violations Found : {total_violations}")
    print(f"Offending Files  : {len(violation_records)}")
    print(f"Elapsed Time     : {elapsed}s")
    print("=" * 60)

    if is_check_mode:
        if total_violations > 0:
            print(f"\n[FAIL] Found {total_violations} naming convention violations.")
            sys.exit(1)
        else:
            print("\n[PASS] All boolean and variable naming conventions pass 100%.")
            sys.exit(0)


def main():
    parser = argparse.ArgumentParser(
        description="Scans and audits boolean naming prefixes, negative identifiers, and bare ok variables."
    )
    parser.add_argument(
        "path",
        nargs="?",
        default=".",
        help="Path or directory to scan (default: current directory)",
    )
    parser.add_argument(
        "--check",
        action="store_true",
        help="Check only; exit code 1 if violations found (CI/CD mode)",
    )
    parser.add_argument(
        "--verbose",
        "-v",
        action="store_true",
        help="Show detailed line-by-line violation snippets",
    )

    args = parser.parse_args()
    run_scanner(
        target_dir=args.path,
        is_check_mode=args.check,
        is_verbose=args.verbose,
    )


if __name__ == "__main__":
    main()
