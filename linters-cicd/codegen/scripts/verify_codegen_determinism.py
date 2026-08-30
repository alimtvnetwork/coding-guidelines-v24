#!/usr/bin/env python3
"""
Verify Codegen Determinism
Re-runs inverted-field codegen on source fixtures and asserts byte-for-byte determinism against committed expected fixtures.
"""

import difflib
from enum import Enum
from pathlib import Path
import subprocess
import sys
import tempfile

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8", errors="replace")
if hasattr(sys.stderr, "reconfigure"):
    sys.stderr.reconfigure(encoding="utf-8", errors="replace")

# --- Top-Level Enums & Constants ---
class ExitCodeType(int, Enum):
    SUCCESS = 0
    FAILURE = 1

SCRIPT_DIR = Path(__file__).resolve().parent
CODEGEN_DIR = SCRIPT_DIR.parent
SOURCES_DIR = CODEGEN_DIR / "fixtures" / "sources"
EXPECTED_DIR = CODEGEN_DIR / "fixtures" / "expected"
TOOL_PATH = CODEGEN_DIR / "inverted_fields.py"

TARGET_CONFIGS: tuple[tuple[str, str], ...] = (
    ("go", "go"),
    ("php", "php"),
    ("typescript", "ts"),
)

def execute_codegen(lang: str, input_file: Path, output_file: Path) -> bool:
    """Executes the inverted_fields.py generator for a single target."""
    cmd = [
        sys.executable,
        str(TOOL_PATH),
        "--input", str(input_file),
        "--lang", lang,
        "--output", str(output_file),
    ]
    res = subprocess.run(cmd, capture_output=True, text=True, encoding="utf-8", errors="replace")
    if res.returncode != ExitCodeType.SUCCESS:
        print(f"::error::Codegen failed for {lang}:\n{res.stderr}", file=sys.stderr)
        return False
    return True

def compare_outputs(expected_file: Path, actual_file: Path, lang: str) -> bool:
    """Compares actual generated output against expected fixture and prints diff on drift."""
    expected_text = expected_file.read_text(encoding="utf-8").replace("\r\n", "\n")
    actual_text = actual_file.read_text(encoding="utf-8").replace("\r\n", "\n")

    if expected_text == actual_text:
        line_count = len(expected_text.splitlines())
        print(f"  OK {lang} ({line_count} lines)")
        return True

    print(f"::error::Codegen drift detected for lang={lang}", file=sys.stderr)
    print(f"         Expected: {expected_file}", file=sys.stderr)
    print(f"         Actual:   {actual_file}", file=sys.stderr)
    print(f"         Fix:      python linters-cicd/codegen/scripts/regen_codegen_fixtures.py && commit\n", file=sys.stderr)

    diff = difflib.unified_diff(
        expected_text.splitlines(keepends=True),
        actual_text.splitlines(keepends=True),
        fromfile=str(expected_file),
        tofile=str(actual_file),
    )
    sys.stderr.writelines(diff)
    return False

def verify_target(lang: str, ext: str, temp_dir: Path) -> bool:
    """Runs codegen and verifies determinism for a single language target."""
    input_file = SOURCES_DIR / f"User.{ext}"
    expected_file = EXPECTED_DIR / f"User.generated.{ext}"
    actual_file = temp_dir / f"User.generated.{ext}"

    if not expected_file.is_file():
        print(f"::error::Expected fixtures missing: {expected_file}", file=sys.stderr)
        return False

    is_codegen_ok = execute_codegen(lang, input_file, actual_file)
    if not is_codegen_ok:
        return False

    return compare_outputs(expected_file, actual_file, lang)

def run_verification() -> int:
    """Verifies all target languages for determinism against committed fixtures."""
    if not EXPECTED_DIR.is_dir():
        print(f"::error::Expected fixtures directory missing: {EXPECTED_DIR}", file=sys.stderr)
        print("         Run: python linters-cicd/codegen/scripts/regen_codegen_fixtures.py", file=sys.stderr)
        return ExitCodeType.FAILURE

    print("Verifying codegen determinism against committed fixtures...")
    is_all_clean = True

    with tempfile.TemporaryDirectory(prefix="codegen-verify-") as temp_dir_str:
        temp_dir = Path(temp_dir_str)
        for lang, ext in TARGET_CONFIGS:
            if not verify_target(lang, ext, temp_dir):
                is_all_clean = False

    if not is_all_clean:
        return ExitCodeType.FAILURE

    print("All codegen outputs match expected/ — determinism verified.")
    return ExitCodeType.SUCCESS

if __name__ == "__main__":
    sys.exit(run_verification())
