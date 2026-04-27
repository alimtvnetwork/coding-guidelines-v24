#!/usr/bin/env python3
"""
check-function-lengths.py
=========================

🔴 CODE RED enforcement: every function in scripts/**/*.{sh,ps1} MUST be
8–15 lines (body lines, brace/closer-exclusive). Hard ceiling enforced
here is 15; the lower bound is advisory and not failed on.

Discovery:
  - Walks `scripts/` (override with --root).
  - Includes top-level installer/runner files: fix-repo.{sh,ps1},
    visibility-change.{sh,ps1}, run.{sh,ps1}.

Detection:
  - Bash:   `name() {` ... matching `}` at column 0 (or a single `}` line).
  - PowerShell: `function Verb-Noun {` ... matching `}` via brace counting.

Exit codes:
  0  all functions ≤ MAX_LINES
  1  one or more violations (printed with file:line and length)
  2  usage / discovery error
"""
from __future__ import annotations
import argparse
import re
import sys
from pathlib import Path

MAX_LINES = 15
TOP_LEVEL_FILES = (
    "fix-repo.sh", "fix-repo.ps1",
    "visibility-change.sh", "visibility-change.ps1",
    "run.sh", "run.ps1",
)

BASH_FN_RE = re.compile(r"^\s*([A-Za-z_][A-Za-z0-9_]*)\s*\(\)\s*\{?\s*$")
PS_FN_RE = re.compile(r"^\s*function\s+([A-Za-z_][A-Za-z0-9_-]*)\s*\{?\s*", re.IGNORECASE)


def is_bash_file(path: Path) -> bool:
    return path.suffix == ".sh"


def is_ps_file(path: Path) -> bool:
    return path.suffix == ".ps1"


def measure_braced_block(lines: list[str], start_idx: int) -> int:
    """Return body line count by brace counting from `{` on/after start_idx."""
    depth = 0
    seen_open = False
    body = 0
    for i in range(start_idx, len(lines)):
        line = lines[i]
        opens = line.count("{")
        closes = line.count("}")
        if not seen_open and opens == 0:
            continue
        seen_open = True
        depth += opens - closes
        if i > start_idx:
            body += 1
        if depth <= 0 and seen_open:
            return max(body - 1, 0)
    return body


def scan_file(path: Path) -> list[tuple[str, int, int]]:
    """Return list of (function_name, line_number, body_length) violations + sizes."""
    try:
        lines = path.read_text(encoding="utf-8", errors="replace").splitlines()
    except OSError as exc:
        print(f"::error::cannot read {path}: {exc}", file=sys.stderr)
        return []
    fn_re = BASH_FN_RE if is_bash_file(path) else PS_FN_RE
    findings: list[tuple[str, int, int]] = []
    for idx, line in enumerate(lines):
        match = fn_re.match(line)
        if not match:
            continue
        name = match.group(1)
        length = measure_braced_block(lines, idx)
        findings.append((name, idx + 1, length))
    return findings


def discover_targets(root: Path, repo_root: Path) -> list[Path]:
    targets: list[Path] = []
    if root.exists():
        for ext in ("*.sh", "*.ps1"):
            targets.extend(root.rglob(ext))
    for name in TOP_LEVEL_FILES:
        candidate = repo_root / name
        if candidate.exists():
            targets.append(candidate)
    return sorted(set(targets))


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--root", default="scripts", help="directory to scan")
    parser.add_argument("--repo-root", default=".", help="repo root for top-level files")
    parser.add_argument("--max", type=int, default=MAX_LINES)
    parser.add_argument("--verbose", action="store_true")
    args = parser.parse_args()

    repo_root = Path(args.repo_root).resolve()
    root = (repo_root / args.root).resolve()
    targets = discover_targets(root, repo_root)
    if not targets:
        print(f"::error::no .sh/.ps1 files found under {root}", file=sys.stderr)
        return 2

    violations = 0
    checked = 0
    for path in targets:
        for name, lineno, length in scan_file(path):
            checked += 1
            if length > args.max:
                rel = path.relative_to(repo_root)
                print(f"{rel}:{lineno}: function '{name}' has {length} body lines (max {args.max})")
                violations += 1
            elif args.verbose:
                rel = path.relative_to(repo_root)
                print(f"OK  {rel}:{lineno} {name} ({length})")

    print(f"\nChecked {checked} functions across {len(targets)} files; "
          f"{violations} violation(s).", file=sys.stderr)
    return 1 if violations else 0


if __name__ == "__main__":
    sys.exit(main())
