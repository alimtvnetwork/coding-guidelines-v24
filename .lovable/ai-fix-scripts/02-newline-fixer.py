#!/usr/bin/env python3
"""
Fast Newline & Trailing Whitespace Fixer
Enforces clean UNIX LF line endings, trims trailing spaces, and ensures a single trailing newline.
Multi-folder capable, customizable extensions, and thread-safe lazy regex engine.
"""

import argparse
from pathlib import Path
import sys

sys.path.insert(0, str(Path(__file__).parent))
try:
    from importlib import import_module
    engine = import_module("00-shared-engine")
    process_repository_files = engine.process_repository_files
    read_file_lf = engine.read_file_lf
    write_file_lf = engine.write_file_lf
    normalize_extensions = engine.normalize_extensions
    normalize_rel_path = engine.normalize_rel_path
    ExitCodeType = engine.ExitCodeType
    RegexPatternType = engine.RegexPatternType
    get_compiled_regex = engine.get_compiled_regex
    DEFAULT_TEXT_EXTENSIONS = engine.DEFAULT_TEXT_EXTENSIONS
except Exception:
    ExitCodeType = None
    RegexPatternType = None
    get_compiled_regex = None
    DEFAULT_TEXT_EXTENSIONS = (".md", ".py", ".ts")

def clean_file_content(content: str) -> str:
    """Strips trailing whitespace per line and guarantees a single final newline."""
    re_crlf = get_compiled_regex(RegexPatternType.Crlf)
    normalized = re_crlf.sub("\n", content)
    lines = [line.rstrip() for line in normalized.split("\n")]
    while lines:
        if lines[-1]:
            break
        lines.pop()
    return "\n".join(lines) + "\n"

def process_file_newlines(file_path: Path, is_fix_mode: bool = False) -> tuple[str, bool]:
    """Checks and optionally fixes newlines and trailing whitespace in a single file."""
    norm_p = normalize_rel_path(file_path)
    try:
        raw = read_file_lf(file_path)
        if not raw:
            return (norm_p, False)
        cleaned = clean_file_content(raw)
        if raw != cleaned:
            if is_fix_mode:
                write_file_lf(file_path, cleaned)
            return (norm_p, True)
    except Exception:
        pass
    return (norm_p, False)

def run_newline_auditor(
    target_dir: str = ".",
    is_fix_mode: bool = False,
    extensions: set[str] | tuple | None = None
) -> int:
    """Executes two-phase repository scan to audit/fix newlines across any target directory."""
    exts = normalize_extensions(extensions) or DEFAULT_TEXT_EXTENSIONS

    def handler(p: Path):
        path_str, has_issue = process_file_newlines(p, is_fix_mode=is_fix_mode)
        return path_str if has_issue else None

    stats = process_repository_files(handler, root_dir=target_dir, extensions=exts)
    violations = stats["results"]

    if violations:
        action_word = "Fixed" if is_fix_mode else "Found issues in"
        print(f"\n⚠️ {action_word} {len(violations)} file(s) ({stats['elapsed_ms']:.2f}ms):")
        for v in violations[:10]:
            print(f"  ::notice file={v}::{v}")
        if not is_fix_mode:
            return ExitCodeType.ViolationsFound.value if ExitCodeType else 1
    else:
        print(f"✅ All {stats['total_files']} files in '{target_dir}' have clean newlines ({stats['elapsed_ms']:.2f}ms).")

    return ExitCodeType.Success.value if ExitCodeType else 0

def main():
    parser = argparse.ArgumentParser(description="Fix trailing whitespace and newlines across folders")
    parser.add_argument("path", nargs="?", default=".", help="Root directory or subfolder to scan")
    parser.add_argument("--path", "-p", dest="opt_path", help="Alternative flag to specify target directory")
    parser.add_argument("--fix", action="store_true", help="Auto-fix whitespace issues in-place")
    parser.add_argument("--ext", help="Comma-separated extensions to scan (e.g. .md,.ts,.py)")
    args = parser.parse_args()

    target_path = args.opt_path or args.path or "."
    sys.exit(run_newline_auditor(target_dir=target_path, is_fix_mode=args.fix, extensions=args.ext))

if __name__ == "__main__":
    main()
