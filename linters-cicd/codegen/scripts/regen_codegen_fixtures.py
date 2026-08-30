#!/usr/bin/env python3
"""
Regenerate Codegen Expected Fixtures
Regenerates expected artifacts from committed sources for inverted field codegen.
"""

from enum import Enum
from pathlib import Path
import subprocess
import sys

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

def regenerate_target(lang: str, ext: str) -> bool:
    """Regenerates a single expected fixture file for a given language."""
    input_file = SOURCES_DIR / f"User.{ext}"
    output_file = EXPECTED_DIR / f"User.generated.{ext}"
    if not input_file.is_file():
        print(f"::error::Missing source fixture: {input_file}", file=sys.stderr)
        return False

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

    print(f"  ✓ Regenerated {output_file.relative_to(CODEGEN_DIR.parent)}")
    return True

def run_regeneration() -> int:
    """Executes fixture regeneration across all target languages."""
    EXPECTED_DIR.mkdir(parents=True, exist_ok=True)
    print(f"Regenerating codegen fixtures from {SOURCES_DIR.name}...")
    
    is_all_success = True
    for lang, ext in TARGET_CONFIGS:
        if not regenerate_target(lang, ext):
            is_all_success = False

    if not is_all_success:
        return ExitCodeType.FAILURE

    print("Done. All codegen fixtures successfully regenerated.")
    return ExitCodeType.SUCCESS

if __name__ == "__main__":
    sys.exit(run_regeneration())
