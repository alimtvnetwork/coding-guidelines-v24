#!/usr/bin/env python3
"""
Universal Polyglot Codebase & Topology Discovery Engine
Automatically inspects any codebase (Go, Rust, Python, TypeScript, PHP, C#, SQL),
classifies subsystems (Backend, Database, Frontend, CI/CD, Docs), and maintains
a high-speed TTL-cached topology map in tmp/cache/paths/codebase-topology-cache.json.

Usage:
  python .lovable/ai-fix-scripts/18-codebase-topology-discoverer.py [--summary]
  python .lovable/ai-fix-scripts/18-codebase-topology-discoverer.py --query <subsystem-or-language>
  python .lovable/ai-fix-scripts/18-codebase-topology-discoverer.py --refresh [--ttl <seconds>]
"""

import argparse
import datetime
from enum import Enum
import json
import os
from pathlib import Path
import sys
import time
from typing import Any

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")

sys.path.insert(0, str(Path(__file__).parent))
try:
    from importlib import import_module
    engine = import_module("02-shared-engine")
    process_repository_files = engine.process_repository_files
    normalize_rel_path = engine.normalize_rel_path
    is_ignored_directory = engine.is_ignored_directory
    is_ignored_path = engine.is_ignored_path
    is_binary_file = engine.is_binary_file
    atomic_cache_lock = engine.atomic_cache_lock
    DEFAULT_ENCODING = engine.DEFAULT_ENCODING
    LINE_SEPARATOR = engine.LINE_SEPARATOR
    PATH_SEPARATOR = engine.PATH_SEPARATOR
    ExitCodeType = engine.ExitCodeType
except Exception:
    class ExitCodeType(int, Enum):
        SUCCESS = 0
        VIOLATIONS_FOUND = 1
        TOOL_ERROR = 2
    DEFAULT_ENCODING = "utf-8"
    LINE_SEPARATOR = "\n"
    PATH_SEPARATOR = "/"
    normalize_rel_path = lambda p: str(p).replace("\\", "/")
    is_ignored_directory = lambda d: d in {".git", "node_modules", ".venv", "dist", "build", "tmp"}
    is_ignored_path = lambda p: False
    is_binary_file = lambda p: False
    atomic_cache_lock = None

# --- Top-Level Enums Following Standard ---

class SubsystemType(str, Enum):
    """Enumeration for major codebase subsystems."""
    BACKEND = "BACKEND"
    DATABASE = "DATABASE"
    FRONTEND = "FRONTEND"
    CICD = "CICD"
    DOCS = "DOCS"
    CLI = "CLI"
    TESTS = "TESTS"
    UNKNOWN = "UNKNOWN"

class LanguageType(str, Enum):
    """Enumeration for detected programming languages."""
    GO = "GO"
    RUST = "RUST"
    PYTHON = "PYTHON"
    TYPESCRIPT = "TYPESCRIPT"
    JAVASCRIPT = "JAVASCRIPT"
    PHP = "PHP"
    CSHARP = "CSHARP"
    SQL = "SQL"
    SHELL = "SHELL"
    MARKDOWN = "MARKDOWN"
    OTHER = "OTHER"

# --- Constants & Cache Paths ---
TOPOLOGY_CACHE_DIR = Path("tmp/cache/paths")
PRIMARY_TOPOLOGY_CACHE_FILE = TOPOLOGY_CACHE_DIR / "codebase-topology-cache.json"
LEGACY_TOPOLOGY_CACHE_FILE = Path("tmp/codebase-topology-cache.json")
DEFAULT_TTL_SECONDS = 1800  # 30 Minutes

# Subsystem & Language Aliases
QUERY_ALIASES = {
    "db": SubsystemType.DATABASE,
    "database": SubsystemType.DATABASE,
    "sql": SubsystemType.DATABASE,
    "migrations": SubsystemType.DATABASE,
    "schema": SubsystemType.DATABASE,
    "backend": SubsystemType.BACKEND,
    "server": SubsystemType.BACKEND,
    "api": SubsystemType.BACKEND,
    "frontend": SubsystemType.FRONTEND,
    "ui": SubsystemType.FRONTEND,
    "web": SubsystemType.FRONTEND,
    "client": SubsystemType.FRONTEND,
    "app": SubsystemType.FRONTEND,
    "ci": SubsystemType.CICD,
    "cicd": SubsystemType.CICD,
    "workflow": SubsystemType.CICD,
    "actions": SubsystemType.CICD,
    "docs": SubsystemType.DOCS,
    "doc": SubsystemType.DOCS,
    "spec": SubsystemType.DOCS,
    "prompt": SubsystemType.DOCS,
    "prompts": SubsystemType.DOCS,
    "cli": SubsystemType.CLI,
    "commands": SubsystemType.CLI,
    "tests": SubsystemType.TESTS,
    "test": SubsystemType.TESTS,
    "qa": SubsystemType.TESTS,
    "go": LanguageType.GO,
    "golang": LanguageType.GO,
    "rs": LanguageType.RUST,
    "rust": LanguageType.RUST,
    "py": LanguageType.PYTHON,
    "python": LanguageType.PYTHON,
    "ts": LanguageType.TYPESCRIPT,
    "typescript": LanguageType.TYPESCRIPT,
    "js": LanguageType.JAVASCRIPT,
    "javascript": LanguageType.JAVASCRIPT,
    "php": LanguageType.PHP,
    "cs": LanguageType.CSHARP,
    "csharp": LanguageType.CSHARP,
    "sh": LanguageType.SHELL,
    "shell": LanguageType.SHELL,
    "bash": LanguageType.SHELL,
    "ps1": LanguageType.SHELL,
    "powershell": LanguageType.SHELL,
    "md": LanguageType.MARKDOWN,
    "markdown": LanguageType.MARKDOWN,
}

# Known Subsystem Indicators
DATABASE_DIR_HINTS = {"db", "database", "migrations", "migration", "sql", "schemas", "schema", "prisma", "drizzle"}
BACKEND_DIR_HINTS = {"cmd", "internal", "pkg", "api", "routes", "controllers", "handlers", "server", "services", "backend"}
FRONTEND_DIR_HINTS = {"components", "views", "pages", "ui", "web", "frontend", "client", "app", "slides-app"}
CICD_DIR_HINTS = {".github", "workflows", "scripts", "linter-scripts", "linters-cicd", "ci", ".lovable"}
DOCS_DIR_HINTS = {"spec", "docs", "doc", "documentation", "prompts", ".lovable/prompts"}
TESTS_DIR_HINTS = {"tests", "test", "spec", "__tests__", "testing"}

# --- Topology Detection & Analysis Logic ---

def detect_manifests(root_dir: str = ".") -> dict[LanguageType, list[str]]:
    """Detects top-level and submodule package manifests."""
    detected = {lang: [] for lang in LanguageType}
    for root, dirs, files in os.walk(root_dir):
        dirs[:] = [d for d in dirs if not is_ignored_directory(d)]
        for f in files:
            norm_rel = normalize_rel_path(os.path.join(root, f))
            if f in {"go.mod", "go.sum", "go.work"}:
                detected[LanguageType.GO].append(norm_rel)
            elif f in {"Cargo.toml", "Cargo.lock"}:
                detected[LanguageType.RUST].append(norm_rel)
            elif f in {"pyproject.toml", "setup.py", "requirements.txt", "Pipfile", "poetry.lock", "uv.lock"}:
                detected[LanguageType.PYTHON].append(norm_rel)
            elif f == "tsconfig.json":
                detected[LanguageType.TYPESCRIPT].append(norm_rel)
            elif f == "package.json":
                detected[LanguageType.JAVASCRIPT].append(norm_rel)
            elif f in {"composer.json", "composer.lock", "artisan"}:
                detected[LanguageType.PHP].append(norm_rel)
            elif f.endswith(".csproj") or f.endswith(".sln"):
                detected[LanguageType.CSHARP].append(norm_rel)
    return {k: v for k, v in detected.items() if v}

def classify_codebase_subsystems(root_dir: str = ".") -> dict[SubsystemType, dict[str, Any]]:
    """Scans and categorizes directory roots into functional subsystems."""
    subsystems: dict[SubsystemType, dict[str, Any]] = {
        SubsystemType.BACKEND: {"roots": set(), "entrypoints": []},
        SubsystemType.DATABASE: {"roots": set(), "schemaFiles": []},
        SubsystemType.FRONTEND: {"roots": set(), "entrypoints": []},
        SubsystemType.CICD: {"roots": set(), "workflows": []},
        SubsystemType.DOCS: {"roots": set(), "specRoots": []},
        SubsystemType.CLI: {"roots": set(), "entrypoints": []},
        SubsystemType.TESTS: {"roots": set(), "testRunners": []},
    }

    for root, dirs, files in os.walk(root_dir):
        dirs[:] = [d for d in dirs if not is_ignored_directory(d)]
        norm_dir = normalize_rel_path(root)
        dir_name = Path(root).name.lower()
        parts = {p.lower() for p in Path(root).parts}

        # 1. Database & Migrations
        if parts & DATABASE_DIR_HINTS or dir_name in DATABASE_DIR_HINTS:
            subsystems[SubsystemType.DATABASE]["roots"].add(norm_dir)

        # 2. Backend & Server Services
        if parts & BACKEND_DIR_HINTS or dir_name in BACKEND_DIR_HINTS:
            subsystems[SubsystemType.BACKEND]["roots"].add(norm_dir)

        # 3. Frontend UI & Apps
        if parts & FRONTEND_DIR_HINTS or dir_name in FRONTEND_DIR_HINTS:
            subsystems[SubsystemType.FRONTEND]["roots"].add(norm_dir)

        # 4. CI/CD & Workflow Automation
        if parts & CICD_DIR_HINTS or dir_name in CICD_DIR_HINTS:
            subsystems[SubsystemType.CICD]["roots"].add(norm_dir)

        # 5. Docs & Specifications
        if parts & DOCS_DIR_HINTS or dir_name in DOCS_DIR_HINTS:
            subsystems[SubsystemType.DOCS]["roots"].add(norm_dir)

        # 6. Tests & QA
        if parts & TESTS_DIR_HINTS or dir_name in TESTS_DIR_HINTS:
            subsystems[SubsystemType.TESTS]["roots"].add(norm_dir)

        # File-level categorization
        for f in files:
            norm_file = normalize_rel_path(os.path.join(root, f))
            ext = os.path.splitext(f)[1].lower()
            f_lower = f.lower()

            if ext == ".sql" or "schema" in f_lower or "migration" in f_lower:
                subsystems[SubsystemType.DATABASE]["schemaFiles"].append(norm_file)
            if f in {"main.go", "main.py", "main.rs", "server.ts", "server.js", "app.py", "app.go", "index.ts"}:
                subsystems[SubsystemType.BACKEND]["entrypoints"].append(norm_file)
            if "cli" in f_lower or f.startswith("cmd") or "command" in f_lower:
                subsystems[SubsystemType.CLI]["entrypoints"].append(norm_file)
            if "test" in f_lower or f.startswith("test_") or f.endswith("_test.go") or f.endswith(".test.ts"):
                subsystems[SubsystemType.TESTS]["testRunners"].append(norm_file)
            if f.endswith(".yml") or f.endswith(".yaml"):
                if ".github" in norm_file:
                    subsystems[SubsystemType.CICD]["workflows"].append(norm_file)

    serialized = {}
    for st, data in subsystems.items():
        serialized[st.value] = {
            k: sorted(list(v)) if isinstance(v, set) else sorted(v)
            for k, v in data.items()
        }
    return serialized

def build_topology_map(root_dir: str = ".", ttl_seconds: int = DEFAULT_TTL_SECONDS) -> dict[str, Any]:
    """Builds complete topology map with timestamps and TTL expiry information."""
    start_time = time.perf_counter()
    now_utc = datetime.datetime.now(datetime.timezone.utc)
    expires_at = now_utc + datetime.timedelta(seconds=ttl_seconds)

    manifests = detect_manifests(root_dir=root_dir)
    subsystems = classify_codebase_subsystems(root_dir=root_dir)

    lang_file_counts: dict[str, int] = {}
    lang_file_roots: dict[str, set[str]] = {}
    total_files = 0

    for root, dirs, files in os.walk(root_dir):
        dirs[:] = [d for d in dirs if not is_ignored_directory(d)]
        norm_dir = normalize_rel_path(root)
        for f in files:
            total_files += 1
            ext = os.path.splitext(f)[1].lower()

            target_lang = None
            if ext == ".go":
                target_lang = LanguageType.GO.value
            elif ext == ".rs":
                target_lang = LanguageType.RUST.value
            elif ext in {".py", ".pyi"}:
                target_lang = LanguageType.PYTHON.value
            elif ext in {".ts", ".tsx"}:
                target_lang = LanguageType.TYPESCRIPT.value
            elif ext in {".js", ".jsx", ".mjs"}:
                target_lang = LanguageType.JAVASCRIPT.value
            elif ext in {".php", ".phtml"}:
                target_lang = LanguageType.PHP.value
            elif ext == ".cs":
                target_lang = LanguageType.CSHARP.value
            elif ext == ".sql":
                target_lang = LanguageType.SQL.value
            elif ext in {".md", ".markdown"}:
                target_lang = LanguageType.MARKDOWN.value
            elif ext in {".sh", ".bash", ".ps1"}:
                target_lang = LanguageType.SHELL.value

            if target_lang:
                lang_file_counts[target_lang] = lang_file_counts.get(target_lang, 0) + 1
                if target_lang not in lang_file_roots:
                    lang_file_roots[target_lang] = set()
                lang_file_roots[target_lang].add(norm_dir)

    elapsed_ms = (time.perf_counter() - start_time) * 1000

    return {
        "version": "1.0.0",
        "generatedAt": now_utc.isoformat(),
        "expiresAt": expires_at.isoformat(),
        "ttlSeconds": ttl_seconds,
        "scanDurationMs": round(elapsed_ms, 2),
        "totalFiles": total_files,
        "rootPath": normalize_rel_path(root_dir),
        "manifests": {k.value: v for k, v in manifests.items()},
        "languageDistribution": dict(sorted(lang_file_counts.items(), key=lambda x: x[1], reverse=True)),
        "languageRoots": {k: sorted(list(v)) for k, v in lang_file_roots.items()},
        "subsystems": subsystems,
    }

# --- TTL Cache Management ---

def load_cached_topology() -> dict[str, Any] | None:
    """Loads cached topology and checks if TTL is still valid."""
    for cache_p in (PRIMARY_TOPOLOGY_CACHE_FILE, LEGACY_TOPOLOGY_CACHE_FILE):
        if cache_p.exists():
            try:
                with open(cache_p, "r", encoding=DEFAULT_ENCODING) as f:
                    data = json.load(f)
                if isinstance(data, dict):
                    expires_str = data.get("expiresAt")
                    if expires_str:
                        expires_at = datetime.datetime.fromisoformat(expires_str)
                        now_utc = datetime.datetime.now(datetime.timezone.utc)
                        is_valid = (now_utc < expires_at)
                        if is_valid:
                            return data
            except Exception:
                pass
    return None

def save_topology_cache(topology_data: dict[str, Any]) -> None:
    """Saves topology data to primary and legacy cache locations with atomic locking."""
    TOPOLOGY_CACHE_DIR.mkdir(parents=True, exist_ok=True)
    for target in (PRIMARY_TOPOLOGY_CACHE_FILE, LEGACY_TOPOLOGY_CACHE_FILE):
        try:
            temp_file = target.with_suffix(".json.tmp")
            with open(temp_file, "w", encoding=DEFAULT_ENCODING) as f:
                json.dump(topology_data, f, indent=2)
            temp_file.replace(target)
        except Exception:
            pass

def get_or_create_topology(
    root_dir: str = ".",
    is_force_refresh: bool = False,
    ttl_seconds: int = DEFAULT_TTL_SECONDS
) -> dict[str, Any]:
    """Retrieves cached topology if valid, or regenerates and stores a new cache."""
    if not is_force_refresh:
        cached = load_cached_topology()
        has_valid_cache = bool(cached)
        if has_valid_cache:
            return cached

    topology = build_topology_map(root_dir=root_dir, ttl_seconds=ttl_seconds)
    save_topology_cache(topology)
    return topology

# --- Query & Routing Functions ---

def query_topology(query_term: str, topology: dict[str, Any]) -> None:
    """Searches topology for language or subsystem matches and outputs routing paths."""
    q_clean = query_term.strip().lower()
    subsystems = topology.get("subsystems", {})
    manifests = topology.get("manifests", {})
    lang_dist = topology.get("languageDistribution", {})
    lang_roots = topology.get("languageRoots", {})

    print(f"\n⚡ Topology Routing Query for: `{query_term}`\n")
    found = False

    # Check Aliases (e.g. db -> DATABASE, py -> PYTHON)
    matched_target = QUERY_ALIASES.get(q_clean)

    # 1. Match Subsystems
    for subsys_name, info in subsystems.items():
        is_subsys_match = (
            q_clean in subsys_name.lower() or
            (matched_target and isinstance(matched_target, SubsystemType) and matched_target.value == subsys_name)
        )
        if is_subsys_match:
            found = True
            print(f"📦 Subsystem: [{subsys_name}]")
            roots = info.get("roots", [])
            has_roots = bool(roots)
            if has_roots:
                print("   📁 Directory Roots:")
                for r in roots[:15]:
                    print(f"      • {r}")
                if len(roots) > 15:
                    print(f"      ... and {len(roots) - 15} more roots.")
            for k, v in info.items():
                if k != "roots" and isinstance(v, list) and v:
                    print(f"   📄 {k}:")
                    for item in v[:10]:
                        print(f"      • {item}")
                    if len(v) > 10:
                        print(f"      ... and {len(v) - 10} more files.")
            print()

    # 2. Match Languages
    for lang_name, count in lang_dist.items():
        is_lang_match = (
            q_clean in lang_name.lower() or
            (matched_target and isinstance(matched_target, LanguageType) and matched_target.value == lang_name)
        )
        if is_lang_match:
            found = True
            print(f"🔤 Language: [{lang_name}] — {count} tracked files")
            man = manifests.get(lang_name, [])
            has_man = bool(man)
            if has_man:
                print("   📋 Manifests & Configs:")
                for m in man:
                    print(f"      • {m}")
            roots = lang_roots.get(lang_name, [])
            has_roots = bool(roots)
            if has_roots:
                print("   📁 Primary Directories:")
                for r in roots[:12]:
                    print(f"      • {r}")
                if len(roots) > 12:
                    print(f"      ... and {len(roots) - 12} more folders.")
            print()

    if not found:
        print(f"⚠️ No direct subsystem or language matches for '{query_term}'.")
        print("💡 Available Subsystems: BACKEND, DATABASE, FRONTEND, CICD, DOCS, CLI, TESTS")
        print(f"💡 Available Languages: {', '.join(lang_dist.keys())}")

def print_topology_summary(topology: dict[str, Any]) -> None:
    """Prints a clear, high-density terminal summary of the codebase topology."""
    print("================================================================================")
    print(f"🌐 Polyglot Codebase Topology Map (Generated: {topology['generatedAt']})")
    print(f"⏱️ TTL Expiry: {topology['expiresAt']} | Scan Time: {topology['scanDurationMs']}ms")
    print("================================================================================")

    print("\n📊 Polyglot Language Breakdown:")
    for lang, count in topology.get("languageDistribution", {}).items():
        man = topology.get("manifests", {}).get(lang, [])
        man_str = f" (Manifests: {', '.join(man[:2])})" if man else ""
        print(f"   • {lang:<14} : {count:>5} files{man_str}")

    print("\n🏛️ Subsystem Map & Navigation Roots:")
    subsystems = topology.get("subsystems", {})
    for subsys_name, data in subsystems.items():
        roots = data.get("roots", [])
        if roots:
            display_roots = ", ".join(roots[:4])
            if len(roots) > 4:
                display_roots += f" (+{len(roots)-4} more)"
            print(f"   • {subsys_name:<10} : {display_roots}")
        else:
            print(f"   • {subsys_name:<10} : (none detected)")

    print("\n💡 AI Navigation Tip:")
    print("   Query specific folders instantly using:")
    print("   `python .lovable/ai-fix-scripts/18-codebase-topology-discoverer.py --query <go|python|db|backend>`")
    print("================================================================================")

# --- CLI Entry Point ---

def main():
    parser = argparse.ArgumentParser(
        description="Universal Polyglot Codebase & Topology Discovery Engine",
        formatter_class=argparse.RawDescriptionHelpFormatter
    )
    parser.add_argument("--query", "-q", help="Search subsystem or language routing (e.g. go, rust, db, backend)")
    parser.add_argument("--summary", "-s", action="store_true", help="Print topology summary breakdown")
    parser.add_argument("--refresh", "-r", action="store_true", help="Force regenerate topology cache")
    parser.add_argument("--ttl", type=int, default=DEFAULT_TTL_SECONDS, help="TTL cache duration in seconds (default: 1800)")
    parser.add_argument("--json", "-j", action="store_true", help="Output raw JSON topology map")
    parser.add_argument("--path", "-p", default=".", help="Root directory to discover (default: .)")
    args = parser.parse_args()

    topology = get_or_create_topology(
        root_dir=args.path,
        is_force_refresh=args.refresh,
        ttl_seconds=args.ttl
    )

    has_json = args.json
    if has_json:
        print(json.dumps(topology, indent=2))
        sys.exit(0)

    has_query = bool(args.query)
    if has_query:
        query_topology(args.query, topology)
        sys.exit(0)

    print_topology_summary(topology)
    sys.exit(0)

if __name__ == "__main__":
    main()
