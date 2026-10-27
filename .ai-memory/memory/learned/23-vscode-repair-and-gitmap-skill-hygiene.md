# VS Code Repair & GitMap Skill Hygiene Invariant

## Context & Decision
VS Code repair commands and bloat commands (`gitmap disk`, OS storage dev-clean, starbar client list, cluster status, cluster workers) have been permanently purged from GitMap autonomous skills (`gitmap` skill) to keep the AI skill compact, high-performance, and focused on core repository operations.

## Invariant Rules
1. **VS Code Repair Exemption:** GitMap does not require or advertise VS Code repair routines in public agent skills. Any manual VS Code cache/extension repair is considered an internal developer environment fix, not a repository workflow command.
2. **Cluster & Starbar Commands Purged:** Commands related to servers-clients topology (`sc list`, `sc ping`), cluster execution workers, and cluster status are removed from the AI skill set.
3. **Rust Toolchain Preservation:** Rust and Cargo toolchain commands (`gitmap cargo status`, `gitmap install cargo --check`, `gitmap install cargo`, `gitmap install --list`) are preserved and supported.
4. **Memory Isolation Invariant:** This memory entry resides strictly within `.ai-memory/memory/learned/` and MUST NEVER be synchronized or shared across repositories.
