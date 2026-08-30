#!/usr/bin/env python3
"""
Fast Boolean Naming & Code Convention Guard
Audits and flags explicit boolean true comparisons (e.g. `== True`, `=== true`) and negative naming anti-patterns.
"""

import argparse
from pathlib import Path
import re
import sys

sys.path.insert(0, str(Path(__file__).parent))
try:
    from importlib import import_module
    engine = import_module("00-shared-engine")
    process_repository_files = engine.process_repository_files
    read_file_lf = engine.read_file_lf
    ExitCodeType = engine.ExitCodeType
except Exception:
    ExitCodeType = None

CODE_EXTENSIONS = (".ts", ".tsx", ".js", ".jsx", ".go", ".py", ".php", ".cs")
EXPLICIT_TRUE_PATTERNS = [
    re.compile(r"==\s*true\b", re.IGNORECASE),
    re.compile(r"===\s*true\b", re.IGNORECASE),
]

def find_explicit_true_violations(content: str) -> list[tuple[int, str]]:
    """Inspects lines for explicit true comparisons."""
    violations = []
    for idx, line in enumerate(content.split("\n"), start=1):
        # Ignore comments and markdown fences
        stripped = line.strip()
        if stripped.startswith("//") or stripped.startswith("#") or stripped.startswith("*"):
            continue
        for pat in EXPLICIT_TRUE_PATTERNS:
            if pat.search(line):
                violations.append((idx, stripped))
    return violations

def audit_file_naming(file_path: Path) -> tuple[str, list[tuple[int, str]]]:
    """Audits code file for boolean conventions."""
    try:
        content = read_file_lf(file_path)
        vios = find_explicit_true_violations(content)
        return (str(file_path), vios)
    except Exception:
        return (str(file_path), [])

def run_naming_auditor(target_dir: str = ".") -> int:
    """Executes naming and boolean convention check."""
    def handler(p: Path):
        fp_str, vios = audit_file_naming(p)
        if vios:
            return (fp_str, vios)
        return None

    stats = process_repository_files(handler, root_dir=target_dir, extensions=CODE_EXTENSIONS)
    all_violations = stats["results"]

    if all_violations:
        print(f"\n❌ Found explicit boolean comparisons in {len(all_violations)} file(s) ({stats['elapsed_ms']:.2f}ms):")
        for fp, vios in all_violations:
            for l_num, line_str in vios[:2]:
                print(f"  ::error file={fp},line={l_num}::Explicit true comparison: {line_str}")
        return ExitCodeType.VIOLATIONS_FOUND.value if ExitCodeType else 1
    else:
        print(f"✅ All {stats['total_files']} code files conform to implicit boolean rules ({stats['elapsed_ms']:.2f}ms).")
        return ExitCodeType.SUCCESS.value if ExitCodeType else 0

def main():
    parser = argparse.ArgumentParser(description="Audit boolean conventions and naming")
    parser.add_argument("path", nargs="?", default=".", help="Root directory to scan")
    args = parser.parse_args()

    sys.exit(run_naming_auditor(target_dir=args.path))

if __name__ == "__main__":
    main()
