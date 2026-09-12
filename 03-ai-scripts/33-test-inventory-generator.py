#!/usr/bin/env python3
"""
33-test-inventory-generator.py
==============================
Generates and maintains the centralized test inventory manifest at `.lovable/test-inventory.json`
and provides safe, atomic file change recording into `.lovable/temp/recent-file-changes.json`
with file locking to ensure concurrency safety across multi-agent turns.

Usage:
  # Scan and generate / update test inventory:
  python 03-ai-scripts/33-test-inventory-generator.py

  # Record modified files safely under lock:
  python 03-ai-scripts/33-test-inventory-generator.py --record "04-code/golang/pkg/appfault/appfault.go"

  # Query tests associated with recent changes:
  python 03-ai-scripts/33-test-inventory-generator.py --query-recent
"""

import argparse
import contextlib
import datetime
import hashlib
import json
import os
from pathlib import Path
import re
import sys
import time
from typing import Any

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")
if hasattr(sys.stderr, "reconfigure"):
    sys.stderr.reconfigure(encoding="utf-8")

REPO_ROOT = Path(__file__).resolve().parent.parent
LOVABLE_DIR = REPO_ROOT / ".lovable"
TEMP_DIR = LOVABLE_DIR / "temp"
TEST_INVENTORY_PATH = LOVABLE_DIR / "test-inventory.json"
RECENT_CHANGES_PATH = TEMP_DIR / "recent-file-changes.json"
LOCK_FILE_PATH = TEMP_DIR / "recent-file-changes.lock"

FUNC_START_RE = re.compile(r"^func\s+(?:\([^)]+\)\s+)?([A-Za-z0-9_]+)\s*\(")
TEST_START_RE = re.compile(r"^func\s+(Test[A-Za-z0-9_]*)\s*\(")
PY_TEST_START_RE = re.compile(r"^def\s+(test_[A-Za-z0-9_]*)\s*\(")
TS_TEST_START_RE = re.compile(r"""(?:it|test)\s*\(\s*["'`]([^"'`]+)["'`]""")


def compute_file_hash(filepath: Path) -> str:
    """Computes a 16-character SHA-256 hash for a file."""
    if not filepath.is_file():
        return ""
    try:
        data = filepath.read_bytes()
        return hashlib.sha256(data).hexdigest()[:16]
    except OSError:
        return ""


def normalize_repo_rel(path_input: str | Path) -> str:
    """Normalizes any path to a forward-slash relative path from the repository root."""
    raw_str = str(path_input).strip()
    try:
        p = Path(raw_str)
        if p.is_absolute():
            rel = p.resolve().relative_to(REPO_ROOT.resolve())
            raw_str = str(rel)
    except Exception:
        pass
    norm = raw_str.replace("\\", "/").strip()
    if norm.startswith("./"):
        norm = norm[2:]
    return norm


@contextlib.contextmanager
def file_lock(lock_path: Path, timeout_sec: float = 10.0):
    """Acquires a cross-platform cooperative lock file with timeout."""
    lock_path.parent.mkdir(parents=True, exist_ok=True)
    start_time = time.time()
    pid = os.getpid()
    acquired = False

    while not acquired:
        try:
            fd = os.open(str(lock_path), os.O_CREAT | os.O_EXCL | os.O_RDWR)
            os.write(fd, f"{pid}\n".encode("utf-8"))
            os.close(fd)
            acquired = True
        except FileExistsError:
            if time.time() - start_time > timeout_sec:
                try:
                    lock_path.unlink()
                except OSError:
                    pass
            time.sleep(0.05)

    try:
        yield
    finally:
        try:
            lock_path.unlink()
        except OSError:
            pass


def atomic_write_json(filepath: Path, data: dict[str, Any]) -> None:
    """Writes JSON data atomically using a temp file replacement."""
    filepath.parent.mkdir(parents=True, exist_ok=True)
    tmp_file = filepath.with_suffix(f".tmp.{os.getpid()}")
    content = json.dumps(data, indent=2, ensure_ascii=False) + "\n"
    tmp_file.write_text(content, encoding="utf-8")
    tmp_file.replace(filepath)


def extract_go_function_hashes(filepath: Path) -> dict[str, str]:
    """Extracts Go function names and hashes from a source file."""
    funcs: dict[str, str] = {}
    if not filepath.is_file():
        return funcs
    try:
        lines = filepath.read_text(encoding="utf-8", errors="ignore").splitlines(keepends=True)
    except OSError:
        return funcs

    curr_name: str | None = None
    curr_lines: list[str] = []
    brace_depth = 0
    in_func = False

    for line in lines:
        if not in_func:
            m = FUNC_START_RE.match(line)
            if m and "{" in line:
                curr_name = m.group(1)
                curr_lines = [line]
                brace_depth = line.count("{") - line.count("}")
                in_func = (brace_depth > 0)
                if not in_func:
                    funcs[curr_name] = hashlib.sha256(line.encode("utf-8")).hexdigest()[:16]
        else:
            curr_lines.append(line)
            brace_depth += line.count("{") - line.count("}")
            if brace_depth <= 0:
                in_func = False
                if curr_name:
                    body = "".join(curr_lines)
                    funcs[curr_name] = hashlib.sha256(body.encode("utf-8")).hexdigest()[:16]

    return funcs


def extract_go_test_functions(filepath: Path) -> dict[str, str]:
    """Extracts Test* function names and hashes from a Go test file."""
    tests: dict[str, str] = {}
    if not filepath.is_file():
        return tests
    try:
        lines = filepath.read_text(encoding="utf-8", errors="ignore").splitlines(keepends=True)
    except OSError:
        return tests

    curr_name: str | None = None
    curr_lines: list[str] = []
    brace_depth = 0
    in_func = False

    for line in lines:
        if not in_func:
            m = TEST_START_RE.match(line)
            if m and "{" in line:
                curr_name = m.group(1)
                curr_lines = [line]
                brace_depth = line.count("{") - line.count("}")
                in_func = (brace_depth > 0)
                if not in_func:
                    tests[curr_name] = hashlib.sha256(line.encode("utf-8")).hexdigest()[:16]
        else:
            curr_lines.append(line)
            brace_depth += line.count("{") - line.count("}")
            if brace_depth <= 0:
                in_func = False
                if curr_name:
                    body = "".join(curr_lines)
                    tests[curr_name] = hashlib.sha256(body.encode("utf-8")).hexdigest()[:16]

    return tests


def extract_python_tests(filepath: Path) -> dict[str, str]:
    """Extracts test functions from a Python test file."""
    tests: dict[str, str] = {}
    if not filepath.is_file():
        return tests
    try:
        content = filepath.read_text(encoding="utf-8", errors="ignore")
    except OSError:
        return tests

    for line in content.splitlines():
        m = PY_TEST_START_RE.match(line)
        if m:
            test_name = m.group(1)
            tests[test_name] = hashlib.sha256(line.encode("utf-8")).hexdigest()[:16]
    return tests


def extract_ts_tests(filepath: Path) -> dict[str, str]:
    """Extracts test descriptions from a TypeScript/JavaScript test file."""
    tests: dict[str, str] = {}
    if not filepath.is_file():
        return tests
    try:
        content = filepath.read_text(encoding="utf-8", errors="ignore")
    except OSError:
        return tests

    for line in content.splitlines():
        m = TS_TEST_START_RE.search(line)
        if m:
            test_name = m.group(1)
            tests[test_name] = hashlib.sha256(line.encode("utf-8")).hexdigest()[:16]
    return tests


def scan_go_tests(repo_root: Path) -> tuple[dict[str, Any], int, int]:
    """Discovers all Go unit tests and maps them to target files and functions."""
    source_funcs: dict[str, dict[str, str]] = {}
    source_hashes: dict[str, str] = {}
    pkg_to_files: dict[str, list[str]] = {}
    tests_dict: dict[str, Any] = {}

    for root, _, files in os.walk(repo_root):
        rel_dir = normalize_repo_rel(root)
        if ".git" in rel_dir or "node_modules" in rel_dir or ".lovable" in rel_dir:
            continue
        pkg_files: list[str] = []
        for f in files:
            if f.endswith(".go") and not f.endswith("_test.go"):
                p = Path(root) / f
                rel_path = normalize_repo_rel(p)
                source_funcs[rel_path] = extract_go_function_hashes(p)
                source_hashes[rel_path] = compute_file_hash(p)
                pkg_files.append(rel_path)
        if pkg_files:
            pkg_to_files[rel_dir] = pkg_files

    for root, _, files in os.walk(repo_root):
        rel_dir = normalize_repo_rel(root)
        if ".git" in rel_dir or "node_modules" in rel_dir or ".lovable" in rel_dir:
            continue
        for f in files:
            if f.endswith("_test.go"):
                p = Path(root) / f
                rel_test_file = normalize_repo_rel(p)
                go_tests = extract_go_test_functions(p)
                candidates = pkg_to_files.get(rel_dir, [])
                base_stem = f.replace("_test.go", "").replace("_unit", "")
                primary_target = next((c for c in candidates if Path(c).stem.startswith(base_stem)), "")
                if not primary_target and candidates:
                    primary_target = candidates[0]

                for test_func, test_hash in go_tests.items():
                    test_id = f"{rel_dir}.{test_func}"
                    func_suffix = test_func[4:] if test_func.startswith("Test") else test_func
                    matched_func = ""
                    matched_file = primary_target
                    matched_hash = ""

                    for c in candidates:
                        funcs = source_funcs.get(c, {})
                        for fn, fhash in funcs.items():
                            if fn.lower() == func_suffix.lower() or func_suffix.lower().startswith(fn.lower()):
                                matched_func = fn
                                matched_file = c
                                matched_hash = fhash
                                break
                        if matched_func:
                            break

                    if not matched_hash and matched_file:
                        matched_hash = source_hashes.get(matched_file, "")

                    tests_dict[test_id] = {
                        "id": test_id,
                        "package": rel_dir,
                        "test_file": rel_test_file,
                        "test_func": test_func,
                        "test_hash": test_hash,
                        "target_file": matched_file,
                        "target_func": matched_func,
                        "code_hash": matched_hash or test_hash,
                        "duration_sec": 0.0,
                        "last_status": "never_run",
                        "last_run_at": "",
                        "needs_run": True,
                    }

    return tests_dict, len(tests_dict), len(set(t["package"] for t in tests_dict.values()))


def scan_python_and_ts_tests(repo_root: Path) -> dict[str, Any]:
    """Scans and indexes Python and TypeScript test files."""
    tests_dict: dict[str, Any] = {}
    for root, _, files in os.walk(repo_root):
        rel_dir = normalize_repo_rel(root)
        if ".git" in rel_dir or "node_modules" in rel_dir or ".lovable" in rel_dir or "dist" in rel_dir:
            continue
        for f in files:
            p = Path(root) / f
            rel_file = normalize_repo_rel(p)
            if (f.startswith("test_") or f.endswith("_test.py")) and f.endswith(".py"):
                py_tests = extract_python_tests(p)
                for tname, thash in py_tests.items():
                    tid = f"{rel_dir}.{tname}"
                    tests_dict[tid] = {
                        "id": tid,
                        "package": rel_dir,
                        "test_file": rel_file,
                        "test_func": tname,
                        "test_hash": thash,
                        "target_file": rel_file.replace("test_", "").replace("_test.py", ".py"),
                        "target_func": "",
                        "code_hash": thash,
                        "duration_sec": 0.0,
                        "last_status": "never_run",
                        "last_run_at": "",
                        "needs_run": True,
                    }
            elif f.endswith(".test.ts") or f.endswith(".test.tsx") or f.endswith(".spec.ts"):
                ts_tests = extract_ts_tests(p)
                for tname, thash in ts_tests.items():
                    tid = f"{rel_dir}.{tname}"
                    target_file = rel_file.replace(".test.ts", ".ts").replace(".test.tsx", ".tsx").replace(".spec.ts", ".ts")
                    tests_dict[tid] = {
                        "id": tid,
                        "package": rel_dir,
                        "test_file": rel_file,
                        "test_func": tname,
                        "test_hash": thash,
                        "target_file": target_file,
                        "target_func": "",
                        "code_hash": thash,
                        "duration_sec": 0.0,
                        "last_status": "never_run",
                        "last_run_at": "",
                        "needs_run": True,
                    }
    return tests_dict


def build_test_inventory(repo_root: Path) -> dict[str, Any]:
    """Compiles the complete test inventory across all supported languages."""
    go_tests, go_count, _ = scan_go_tests(repo_root)
    py_ts_tests = scan_python_and_ts_tests(repo_root)

    combined_tests = {**go_tests, **py_ts_tests}
    packages = set(t["package"] for t in combined_tests.values())

    inventory = {
        "version": 1,
        "updated_at": datetime.datetime.now(datetime.timezone.utc).strftime("%Y-%m-%dT%H:%M:%SZ"),
        "total_tests": len(combined_tests),
        "summary": {
            "total": len(combined_tests),
            "cached": 0,
            "dirty": len(combined_tests),
            "packages": len(packages),
        },
        "tests": combined_tests,
    }

    atomic_write_json(TEST_INVENTORY_PATH, inventory)
    return inventory


def record_recent_changes(changed_files: list[str]) -> dict[str, Any]:
    """Safely appends distinct modified relative paths to recent-file-changes.json under lock."""
    with file_lock(LOCK_FILE_PATH):
        existing_data: dict[str, Any] = {}
        if RECENT_CHANGES_PATH.is_file():
            try:
                existing_data = json.loads(RECENT_CHANGES_PATH.read_text(encoding="utf-8"))
            except Exception:
                existing_data = {}

        file_set = set(existing_data.get("files", []))
        for f in changed_files:
            rel = normalize_repo_rel(f)
            if rel:
                file_set.add(rel)

        inventory_tests: dict[str, Any] = {}
        if TEST_INVENTORY_PATH.is_file():
            try:
                inv = json.loads(TEST_INVENTORY_PATH.read_text(encoding="utf-8"))
                inventory_tests = inv.get("tests", {})
            except Exception:
                pass

        associated_tests: set[str] = set()
        for fpath in file_set:
            stem = Path(fpath).stem
            fdir = str(Path(fpath).parent).replace("\\", "/")
            for tid, tmeta in inventory_tests.items():
                target = tmeta.get("target_file", "")
                test_file = tmeta.get("test_file", "")
                pkg = tmeta.get("package", "")
                if (
                    target == fpath
                    or test_file == fpath
                    or (target and Path(target).stem == stem)
                    or (pkg and (pkg == fdir or fpath.startswith(pkg + "/")))
                ):
                    associated_tests.add(test_file)

        recent_payload = {
            "version": 1,
            "updated_at": datetime.datetime.now(datetime.timezone.utc).strftime("%Y-%m-%dT%H:%M:%SZ"),
            "files": sorted(list(file_set)),
            "associated_tests": sorted(list(associated_tests)),
        }

        atomic_write_json(RECENT_CHANGES_PATH, recent_payload)
        return recent_payload


def main():
    parser = argparse.ArgumentParser(description="Test inventory generator & atomic change recorder.")
    parser.add_argument("--record", nargs="+", help="Record modified file paths to recent-file-changes.json under lock.")
    parser.add_argument("--query-recent", action="store_true", help="Display recently modified files and associated tests.")
    args = parser.parse_args()

    if args.record:
        result = record_recent_changes(args.record)
        print(f"Recorded {len(args.record)} modified file(s). Total tracked: {len(result.get('files', []))}")
        print(f"Associated test files to run on release: {len(result.get('associated_tests', []))}")
        return

    if args.query_recent:
        if RECENT_CHANGES_PATH.is_file():
            print(RECENT_CHANGES_PATH.read_text(encoding="utf-8"))
        else:
            print("No recent changes recorded.")
        return

    print("Scanning codebase to generate test inventory...")
    inv = build_test_inventory(REPO_ROOT)
    print(f"Generated test inventory at {normalize_repo_rel(TEST_INVENTORY_PATH)}")
    print(f"Total Tests Indexed : {inv['total_tests']}")
    print(f"Total Packages      : {inv['summary']['packages']}")


if __name__ == "__main__":
    main()
