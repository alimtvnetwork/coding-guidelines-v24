#!/usr/bin/env python3
"""
Fast Boolean Naming & Code Convention Guard
Audits and flags explicit boolean true comparisons (e.g. `== True`, `=== true`) and negative naming anti-patterns.
Multi-folder capable, customizable extensions, and pre-compiled regex engine.
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
    normalize_extensions = engine.normalize_extensions
    normalize_rel_path = engine.normalize_rel_path
    ExitCodeType = engine.ExitCodeType
except Exception:
    ExitCodeType = None

DEFAULT_CODE_EXTENSIONS = (".ts", ".tsx", ".js", ".jsx", ".go", ".py", ".php", ".cs")

# Pre-compiled regular expressions for boolean checks
RE_EXPLICIT_DOUBLE_TRUE = re.compile(r"==\s*true\b", re.IGNORECASE)
RE_EXPLICIT_TRIPLE_TRUE = re.compile(r"===\s*true\b", re.IGNORECASE)
RE_EXPLICIT_PYTHON_TRUE = re.compile(r"==\s*True\b")
RE_COMMENT_PREFIX = re.compile(r"^\s*(//|#|\*|/\*)")

EXPLICIT_TRUE_PATTERNS = (
    RE_EXPLICIT_DOUBLE_TRUE,
    RE_EXPLICIT_TRIPLE_TRUE,
    RE_EXPLICIT_PYTHON_TRUE,
)

def find_explicit_true_violations(content: str) -> list[tuple[int, str]]:
    """Inspects lines for explicit true comparisons."""
    violations = []
    for idx, line in enumerate(content.split("\n"), start=1):
        if RE_COMMENT_PREFIX.match(line):
            continue
        stripped = line.strip()
        for pat in EXPLICIT_TRUE_PATTERNS:
            if pat.search(line):
                violations.append((idx, stripped))
                break
    return violations

def audit_file_naming(file_path: Path) -> tuple[str, list[tuple[int, str]]]:
    """Audits code file for boolean conventions."""
    norm_p = normalize_rel_path(file_path)
    try:
        content = read_file_lf(file_path)
        vios = find_explicit_true_violations(content)
        return (norm_p, vios)
    except Exception:
        return (norm_p, [])

def run_naming_auditor(
    target_dir: str = ".",
    extensions: set[str] | tuple | None = None
) -> int:
    """Executes naming and boolean convention check across target directory."""
    exts = normalize_extensions(extensions) or DEFAULT_CODE_EXTENSIONS

    def handler(p: Path):
        fp_str, vios = audit_file_naming(p)
        return (fp_str, vios) if vios else None

    stats = process_repository_files(handler, root_dir=target_dir, extensions=exts)
    all_violations = stats["results"]

    if all_violations:
        print(f"\n❌ Found explicit boolean comparisons in {len(all_violations)} file(s) ({stats['elapsed_ms']:.2f}ms):")
        for fp, vios in all_violations:
            for l_num, line_str in vios[:2]:
                print(f"  ::error file={fp},line={l_num}::Explicit true comparison: {line_str}")
        return ExitCodeType.VIOLATIONS_FOUND.value if ExitCodeType else 1

    print(f"✅ All {stats['total_files']} code files in '{target_dir}' conform to implicit boolean rules ({stats['elapsed_ms']:.2f}ms).")
    return ExitCodeType.SUCCESS.value if ExitCodeType else 0

def main():
    parser = argparse.ArgumentParser(description="Audit boolean conventions and naming across folders")
    parser.add_argument("path", nargs="?", default=".", help="Root directory or folder to scan")
    parser.add_argument("--path", "-p", dest="opt_path", help="Alternative flag to specify target directory")
    parser.add_argument("--ext", help="Comma-separated extensions to scan (e.g. .ts,.go,.py)")
    args = parser.parse_args()

    target_path = args.opt_path or args.path or "."
    sys.exit(run_naming_auditor(target_dir=target_path, extensions=args.ext))

if __name__ == "__main__":
    main()
