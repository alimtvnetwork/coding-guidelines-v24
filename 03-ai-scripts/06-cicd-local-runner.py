#!/usr/bin/env python3
"""
Fast Parallel Multi-Worker Local CI/CD Runner
Executes repository quality gates concurrently via ThreadPoolExecutor worker groups.

Key Features:
- Parallel execution via configurable worker group (ThreadPoolExecutor).
- Smart log display filtering:
  - Default / --failed: Shows clean status for all gates; full logs are ONLY printed for failed gates. If all pass, logs are suppressed.
  - --all: Shows full stdout/stderr output for all gates (both passed and failed).
- Per-gate execution duration and total wall time reporting.
- Clean real-time ticker and final summary table.

All Enums, Constants, and Base Configurations are imported directly from 02-shared-engine.py.
"""

import argparse
from concurrent.futures import ThreadPoolExecutor, as_completed
from dataclasses import dataclass
from importlib import import_module
import os
from pathlib import Path
import subprocess
import sys
import time

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")
if hasattr(sys.stderr, "reconfigure"):
    sys.stderr.reconfigure(encoding="utf-8")

sys.path.insert(0, str(Path(__file__).parent))
engine = import_module("02-shared-engine")

ExitCodeType = engine.ExitCodeType
DEFAULT_ENCODING = engine.DEFAULT_ENCODING
LINE_SEPARATOR = engine.LINE_SEPARATOR
DEFAULT_MAX_WORKERS = engine.DEFAULT_MAX_WORKERS
CI_JOBS_MATRIX = engine.CI_JOBS_MATRIX
format_keys = engine.format_keys


@dataclass
class JobResult:
    """Represents the execution outcome of an individual CI quality gate."""
    name: str
    is_success: bool
    output: str
    duration_sec: float
    return_code: int


def execute_ci_job(job_name: str, command: list[str]) -> JobResult:
    """Executes a single validation check asynchronously and records output and duration."""
    start_time = time.perf_counter()
    try:
        res = subprocess.run(
            command,
            capture_output=True,
            text=True,
            encoding=DEFAULT_ENCODING,
            errors="replace"
        )
        duration = time.perf_counter() - start_time
        is_success = (res.returncode == 0)

        output_parts: list[str] = []
        stdout_clean = res.stdout.strip()
        if stdout_clean:
            output_parts.append(stdout_clean)
        stderr_clean = res.stderr.strip()
        if stderr_clean:
            output_parts.append(stderr_clean)

        combined_output = LINE_SEPARATOR.join(output_parts)
        return JobResult(
            name=job_name,
            is_success=is_success,
            output=combined_output,
            duration_sec=duration,
            return_code=res.returncode
        )
    except Exception as exc:
        duration = time.perf_counter() - start_time
        return JobResult(
            name=job_name,
            is_success=False,
            output=f"Failed to execute process: {exc}",
            duration_sec=duration,
            return_code=-1
        )


def run_pipeline(
    jobs: dict[str, list[str]] | None = None,
    max_workers: int | None = None,
    show_all: bool = False,
    filter_pattern: str | None = None
) -> int:
    """Dispatches all jobs concurrently using parallel worker group and displays targeted reports."""
    all_jobs = jobs or CI_JOBS_MATRIX

    if filter_pattern:
        pattern_lower = filter_pattern.lower()
        target_jobs = {k: v for k, v in all_jobs.items() if pattern_lower in k.lower()}
        if not target_jobs:
            print(f"⚠️ No CI quality gates matched filter: '{filter_pattern}'")
            return ExitCodeType.TOOL_ERROR.value
    else:
        target_jobs = all_jobs

    worker_count = max_workers or min(len(target_jobs), os.cpu_count() or DEFAULT_MAX_WORKERS, 8)
    total_jobs = len(target_jobs)

    mode_description = "SHOW ALL LOGS (--all)" if show_all else "SHOW FAILED LOGS ONLY (--failed / default)"

    print("================================================================")
    print("           PARALLEL LOCAL CI/CD QUALITY GATE RUNNER             ")
    print("================================================================")
    print(f"🚀 Worker Group Concurrency: {worker_count} parallel workers")
    print(f"📋 Total Enqueued Gates    : {total_jobs}")
    print(f"🔍 Display Mode            : {mode_description}")
    print("----------------------------------------------------------------\n")

    start_wall_time = time.perf_counter()
    results: list[JobResult] = []

    with ThreadPoolExecutor(max_workers=worker_count) as executor:
        future_map = {
            executor.submit(execute_ci_job, name, cmd): name
            for name, cmd in target_jobs.items()
        }
        for future in as_completed(future_map):
            res = future.result()
            results.append(res)
            status_icon = "✅" if res.is_success else "❌"
            status_label = "PASS" if res.is_success else "FAIL"
            print(f"[{len(results):2d}/{total_jobs:2d}] {status_icon} [{status_label}] {res.name} ({res.duration_sec:.2f}s)")

    total_wall_duration = time.perf_counter() - start_wall_time

    # Sort results to match original job matrix order
    job_order = list(target_jobs.keys())
    results.sort(key=lambda r: job_order.index(r.name) if r.name in job_order else 999)

    passed_count = sum(1 for r in results if r.is_success)
    failed_count = sum(1 for r in results if not r.is_success)
    has_failures = (failed_count > 0)

    print("\n======================= FINAL SUMMARY =======================")
    for r in results:
        status_icon = "✅" if r.is_success else "❌"
        status_word = "PASSED" if r.is_success else "FAILED"
        print(f"{status_icon} [{status_word}] {r.name:<40} ({r.duration_sec:.2f}s)")

    print("------------------------------------------------------------")
    print(f"Total Duration : {total_wall_duration:.2f}s")
    print(f"Gates Passed   : {passed_count}/{total_jobs}")
    print(f"Gates Failed   : {failed_count}/{total_jobs}")
    print("------------------------------------------------------------")

    # Display Logs
    if show_all:
        print("\n===================== ALL QUALITY GATE LOGS =====================")
        for r in results:
            status_word = "PASSED" if r.is_success else "FAILED"
            print(f"\n--- [{status_word}] {r.name} ({r.duration_sec:.2f}s, exit code: {r.return_code}) ---")
            if r.output:
                print(r.output)
            else:
                print("(no output)")
            print("-----------------------------------------------------------------")
    else:
        if has_failures:
            print("\n================== FAILED QUALITY GATE LOGS ==================")
            for r in results:
                if not r.is_success:
                    print(f"\n❌ FAILED: {r.name} (exit code: {r.return_code}, duration: {r.duration_sec:.2f}s)")
                    print(f"--- {r.name} LOG ---")
                    if r.output:
                        print(r.output)
                    else:
                        print("(no output)")
                    print(f"--- END {r.name} LOG ---")
            print("-----------------------------------------------------------------")

    if has_failures:
        print(f"\n❌ Pipeline failed: {failed_count} quality gate(s) reported violations.")
        return ExitCodeType.VIOLATIONS_FOUND.value

    print("\n🎉 All quality gates passed successfully! Codebase is 100% green.")
    return ExitCodeType.SUCCESS.value


def parse_arguments() -> argparse.Namespace:
    """Parses command-line arguments for CI runner."""
    parser = argparse.ArgumentParser(
        description="Fast Multi-Worker Local CI/CD Runner with selective log display.",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  # Default: run all gates in parallel; only print failed gate logs (if any fail):
  python 03-ai-scripts/06-cicd-local-runner.py

  # Show only failed logs (explicit):
  python 03-ai-scripts/06-cicd-local-runner.py --failed

  # Show full logs for all gates (both passed and failed):
  python 03-ai-scripts/06-cicd-local-runner.py --all

  # Configure parallel worker count:
  python 03-ai-scripts/06-cicd-local-runner.py --workers 8

  # Filter specific gate:
  python 03-ai-scripts/06-cicd-local-runner.py --filter "Go Base"
        """
    )
    parser.add_argument(
        "--all", "-a",
        action="store_true",
        dest="show_all",
        help="Show full logs for all quality gates (both passed and failed)."
    )
    parser.add_argument(
        "--failed", "-f",
        action="store_true",
        dest="show_failed",
        help="Show logs only for failed quality gates (default behavior)."
    )
    parser.add_argument(
        "--workers", "-w",
        type=int,
        default=None,
        help="Number of concurrent worker threads (default: auto-detected CPU count capped at 8)."
    )
    parser.add_argument(
        "--filter", "-k",
        type=str,
        default=None,
        help="Filter jobs matching substring (case-insensitive)."
    )
    return parser.parse_args()


def main():
    args = parse_arguments()
    show_all = args.show_all
    exit_code = run_pipeline(
        max_workers=args.workers,
        show_all=show_all,
        filter_pattern=args.filter
    )
    sys.exit(exit_code)


if __name__ == "__main__":
    main()

