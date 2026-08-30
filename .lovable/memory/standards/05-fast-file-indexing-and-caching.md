---
name: Fast File Indexing & Caching Strategy
description: Explains how AI agents leverage .lovable/ai-fix-scripts/08-fast-file-scanner.py to index repository files into tmp/ cache for rapid retrieval and zero-overhead file discovery.
type: standard
---

# Fast File Indexing & Caching Strategy

**Tool Script:** `.lovable/ai-fix-scripts/08-fast-file-scanner.py`  
**Cache Output:** `tmp/repo-file-cache.json` & `tmp/repo-file-list.txt`  
**Tags:** `#file-scanner`, `#caching`, `#performance`, `#ai-workflow`

## 1. The Multi-Step File Discovery Problem

In large repositories (>2,000 files), repeated ad-hoc file searches (e.g. recursive `os.walk`, deep grep searches, or multiple shell glob queries) consume significant time and token budget across multiple agent steps.

## 2. The Solution: Pre-Computed Fast File Caching

Before embarking on multi-file audits, refactoring passes, or cross-spec updates, AI agents can execute `08-fast-file-scanner.py` once to build a filtered, memory-efficient index in `tmp/`:

```bash
# Scan full repository:
python .lovable/ai-fix-scripts/08-fast-file-scanner.py

# Scan specific language files:
python .lovable/ai-fix-scripts/08-fast-file-scanner.py --lang go,ts,tsx

# Scan specifications only:
python .lovable/ai-fix-scripts/08-fast-file-scanner.py --path spec/ --ext .md
```

## 3. Cached File Contract

The scanner automatically writes structured outputs into `tmp/`:

1. **`tmp/repo-file-cache.json`**: Structured JSON containing metadata, scan duration, extension statistics, and an array of relative file paths.
2. **`tmp/repo-file-list.txt`**: Clean newline-delimited text list for immediate ingestion by subsequent Python/Bash scripts.

## 4. Agent Consumption Pattern

In subsequent turns or script executions, agents do NOT need to re-query the filesystem:

```python
import json

# Read pre-computed file cache instantly (<1ms):
with open("tmp/repo-file-cache.json", "r", encoding="utf-8") as f:
    file_list = json.load(f)["files"]

for filepath in file_list:
    # Process files directly
    pass
```

## 5. Exclusions & Clean Filtering

The scanner automatically filters out:
- Version control: `.git/`
- Package managers & builds: `node_modules/`, `dist/`, `build/`, `bin/`, `obj/`, `.next/`, `.cache/`, `.venv/`
- Internal agent dirs: `.gemini/`, `.agent/`
- Binary assets: images, zip archives, compiled executables, SQLite databases.
