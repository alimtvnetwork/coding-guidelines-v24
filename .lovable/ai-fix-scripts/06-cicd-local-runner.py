#!/usr/bin/env python3
"""
Fast Multi-Threaded Local CI/CD Runner
Executes repository quality gates in parallel using ThreadPoolExecutor and enforces zero-failure tolerance.

All Enums and Constants are imported directly from 02-shared-engine.py.
"""

from concurrent.futures import ThreadPoolExecutor
from importlib import import_module
from pathlib import Path
import subprocess
import sys

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")

sys.path.insert(0, str(Path(__file__).parent))
engine = import_module("02-shared-engine")

ExitCodeType = engine.ExitCodeType
DEFAULT_ENCODING = engine.DEFAULT_ENCODING
LINE_SEPARATOR = engine.LINE_SEPARATOR

JOBS_MATRIX = {
    "Relative Path Check": [sys.executable, "linter-scripts/check-relative-paths.py"],
    "Prompts Loaded Check": [sys.executable, "linter-scripts/check-prompts-loaded.py"],
    "Readme Install Section Check": [sys.executable, "linter-scripts/check-readme-install-section.py"],
    "Forbidden Strings Check": [sys.executable, "linter-scripts/check-forbidden-strings.py"],
    "Newline Styling Check": [sys.executable, "linter-scripts/check-newline-styling.py"],
    "Fast File Scanner Cache": [sys.executable, ".lovable/ai-fix-scripts/11-fast-file-scanner.py", "--check"],
    "File Size Guard": [sys.executable, ".lovable/ai-fix-scripts/13-file-size-guard.py"],
    "Version Sync Check": [sys.executable, ".lovable/ai-fix-scripts/14-version-sync-checker.py"],
    "Bundle Installer Generation": ["node", "scripts/generate-bundle-installers.mjs"],
    "Spec Tree Sync": ["node", "scripts/sync-spec-tree.mjs"],
    "Codegen Determinism Check": [sys.executable, "linters-cicd/codegen/scripts/verify_codegen_determinism.py"],
    "Spec Verification Coverage": ["node", "scripts/spec-verification/generate-coverage-report.mjs", "--strict", "--out", "reports/spec-verification/coverage.md"],
    "Validate Version JSON": ["node", "scripts/validate-version-json.mjs"],
    "Doc Links Check": ["node", "scripts/docs/check-doc-links.mjs", "readme.md", "docs/installer-fix-repo-flags.md"],
    "Check File Sizes Baseline": [sys.executable, "linter-scripts/check-file-sizes.py", "--check"],
    "Newline Styling MJS Check": ["node", "linter-scripts/check-newline-styling.mjs"],
    "Spec Folder References Check": [sys.executable, "linter-scripts/check-spec-folder-refs.py"],
    "Linters CI/CD Test Suite": [sys.executable, "linters-cicd/tests/run.py"],
}

def execute_ci_job(job_name: str, command: list[str]) -> tuple[str, bool, str]:
    """Executes a single validation check asynchronously."""
    try:
        res = subprocess.run(command, capture_output=True, text=True, encoding=DEFAULT_ENCODING, errors="replace")
        is_success = (res.returncode == 0)
        if is_success:
            return (job_name, True, res.stdout)
        return (job_name, False, res.stdout + LINE_SEPARATOR + res.stderr)
    except Exception as e:
        return (job_name, False, str(e))

def run_pipeline() -> int:
    """Dispatches all jobs concurrently and prints clean summary report."""
    print("🚀 Running Local CI/CD Pipeline via ThreadPoolExecutor...")
    print(f"📋 Enqueued Jobs: {', '.join(JOBS_MATRIX.keys())}{LINE_SEPARATOR}")

    results = []
    with ThreadPoolExecutor(max_workers=4) as executor:
        futures = [executor.submit(execute_ci_job, name, cmd) for name, cmd in JOBS_MATRIX.items()]
        for f in futures:
            results.append(f.result())

    has_failures = False
    print(f"{LINE_SEPARATOR}================ FINAL SUMMARY ================")
    for name, is_success, log in results:
        if is_success:
            print(f"✅ {name}: PASSED")
        else:
            print(f"❌ {name}: FAILED")
            has_failures = True
            print(f"--- {name} LOG ---")
            print(log.strip())
            print(f"--------------------{LINE_SEPARATOR}")

    if has_failures:
        print(f"{LINE_SEPARATOR}❌ Pipeline failed.")
        return ExitCodeType.VIOLATIONS_FOUND.value

    print("🎉 All jobs passed successfully!")
    return ExitCodeType.SUCCESS.value

def main():
    sys.exit(run_pipeline())

if __name__ == "__main__":
    main()
