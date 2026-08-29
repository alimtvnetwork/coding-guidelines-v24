# AI Fix Scripts

This directory contains reusable utility scripts designed to be executed by AI agents (or humans) to enforce repository standards, fix formatting issues, and automate codebase maintenance.

## Available Scripts

### 
ewline_fixer.py
**Purpose:** Scans codebase text and source files to ensure they end with exactly one UNIX-style newline (\n). 
**Supported Extensions:** .md, .txt, .go, .ts, .js, .mjs, .cjs, .jsx, .cs, .vb, .rs, .json, .yml, .yaml, .sh, .ps1
**Usage:** python .lovable/ai-fix-scripts/newline_fixer.py

### 03-cicd-local-runner.py
**Purpose:** Local CI/CD pipeline orchestrator. Runs all local quality gates and linters in parallel using ThreadPoolExecutor.
**Usage:** python .lovable/ai-fix-scripts/03-cicd-local-runner.py

<details><summary>Why this script exists</summary>
Provides a fast, local way to simulate CI quality gates using native Python concurrency, ensuring jobs don't block each other and that all tests must pass before finalizing work.
</details>

### 04-relative-path-fixer.py
**Purpose:** Scans the codebase and automatically resolves absolute paths and file:/// URIs into clean, strictly relative paths.
**Usage:** python .lovable/ai-fix-scripts/04-relative-path-fixer.py

<details><summary>Why this script exists</summary>
Absolute paths break portability across different OS environments (Windows vs Unix) and CI pipelines. This script ensures 100% relative path references across markdown, plans, code, and documentation.
</details>

