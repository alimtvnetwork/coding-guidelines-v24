# Subtask Plan: GitMap Skills Command Matrix & Examples Update

## Subtask ID: `04-subtask`
- **Parent Plan:** `.ai-memory/plans/gitmap-commands-agent-spawning-heatmap-sync.md`
- **Spec Reference:** `02-spec/21-app/gitmap-commands-agent-spawning-heatmap-sync/02-component-and-examples-spec.md`
- **Primary Owner:** Worker 02 (Phase 2 Execution Wave)
- **Status:** `COMPLETED`

---

## 1. Intent & Technical Deliverables

Upgrade `.agents/skills/gitmap/skill.md` and `.cursor/skills/gitmap/skill.md` to:
1. **Permanently Ban All Shell Traversal & Unbuffered File Commands:**
   - Explicitly expand the Non-Negotiable Command Replacement Matrix to prohibit ALL variants of `Get-ChildItem` (including non-recursive `Get-ChildItem -Path ...`, `Get-ChildItem`), shell `dir`, shell `ls`, and unbuffered file inspection tools (`cat`, `type`, `Get-Content`).
2. **Provide Prominent Copy-Paste Command Examples:**
   - Add a high-visibility, copy-paste example section illustrating exact syntax for directory inventory, file globbing, streaming, live regex search, and cognitive heatmap logging.
3. **Elevate Cognitive Heatmap Telemetry:**
   - Document lifecycle commands for `gitmap heatmap log`, `gitmap heatmap status`, `gitmap heatmap tree`, and `gitmap agent heatmap`.

---

## 2. Disjoint File Boundaries

### Owned Files (Exclusive Modification Rights for Worker 02):
- `.agents/skills/gitmap/skill.md`
- `.cursor/skills/gitmap/skill.md`

### Read-Only Context Files:
- `02-spec/21-app/gitmap-commands-agent-spawning-heatmap-sync/02-component-and-examples-spec.md`
- `AGENTS.md`
- `.ai-memory/plans/gitmap-commands-agent-spawning-heatmap-sync.md`

---

## 3. Negative Boundaries & Invariants

- **TOTAL BAN ON GIT COMMANDS:** Worker 02 must NEVER run `git add`, `git commit`, `git push`, `git checkout`, or any git commands. Only the lead orchestrator performs atomic commits at the completion of Phase 3.
- **TOTAL BAN ON RAW SEARCH TOOLS:** Worker 02 must never run `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem`, `dir`, or `ls`. Use `gitmap aum search` and `gitmap find` exclusively.
- **Strict Relative Git Paths:** Zero absolute paths (`C:\...`, `/home/...`) or `file:///` URIs.
- **Strict Lowercase Filenames:** All filenames, references, and directories must be lowercase.
- **100% Skill Parity:** Changes made to `.agents/skills/gitmap/skill.md` must be mirrored identically to `.cursor/skills/gitmap/skill.md`.

---

## 4. Step-by-Step Implementation Guide for Worker 02

### Step 1: Update the Non-Negotiable Command Replacement Matrix
In Section "## Non-Negotiable Command Replacement Matrix (DO NOT USE X; USE GITMAP Y)":
1. Update the directory inventory row to prohibit all forms of `Get-ChildItem`:
   ```markdown
   | ❌ `Get-ChildItem` (all variants), `Get-ChildItem -Path "..."`, `dir`, `ls` | ✅ `gitmap lf [dir]` or `gitmap find "*.*" [dir]` | Sub-millisecond directory inventory (<10ms) without slow PowerShell .NET object pipelines, Windows process hangs, or terminal freezing. |
   | ❌ `Get-ChildItem -Recurse`, `find . -name "..."`, `dir /s /b` | ✅ `gitmap find "<pattern>" [-ext <ext>]` | Multi-core indexed glob discovery across 10,000+ files without disk traversal overhead. |
   | ❌ `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `findstr` | ✅ `gitmap aum search "<pattern>" [dir] [--ext <ext>]` or `gitmap search "<pattern>"` | Multi-core streaming live search (<15ms) with binary null-byte probe, 500 KB file cap, and DH2D SQLite hot cache. Prevents buffer bloat and terminal freezing. |
   | ❌ Raw unbuffered `cat`, `type`, `Get-Content` on source files | ✅ `gitmap cat <filepath>` | Direct zero-disk stream of file contents directly into process stdout for resource-constrained CLI sessions. |
   ```

### Step 2: Add High-Visibility Copy-Paste Command Scenarios
Under Section 5 ("High-Performance Automation (AUM) & File Discovery"), add a dedicated subsection:

```markdown
### High-Speed Command Cheat Sheet & Copy-Paste Scenarios

#### Scenario 1: Directory Inventory & File Discovery (Zero Get-ChildItem)
```bash
# List all files in a relative directory:
gitmap lf "02-spec/21-app"

# Find all markdown files matching a glob:
gitmap find "*.md" "02-spec/21-app"

# Find files across any subpath matching a pattern:
gitmap find "*agent*" .ai-memory

# Find files by exact substring:
gitmap find-files-any "heatmap"
```

#### Scenario 2: Live Code & Symbol Search (Zero Select-String / rg)
```bash
# Search for symbol with extension filter:
gitmap aum search "RunFleetPASCommand" cli -e .go

# Case-insensitive regex search:
gitmap aum search -r "(\"pas\"|\"pa\")" cli/cmd -e .go -i

# Search across prompt templates:
gitmap aum search "Subagents" 01-prompts -e .md
```

#### Scenario 3: Stream File Contents to Stdout (Zero cat / Get-Content)
```bash
# Stream specification or config file directly to stdout:
gitmap cat "02-spec/21-app/01-architecture-spec.md"
```
```

### Step 3: Enhance Section 12 (File Contention Heatmap & Agent Activity)
Ensure the following commands are documented with explicit copy-paste syntax:
```bash
# Preflight check before editing files:
gitmap heatmap status
gitmap heatmap tree
gitmap agent heatmap

# Record cognitive reasoning before/after file touch:
gitmap heatmap log -t "<task-slug>" -a "READ" -f "<path>" -r "<reasoning>"
gitmap heatmap log -t "<task-slug>" -a "EDIT" -f "<path>" -r "<reasoning>"

# Direct non-blocking SQL trace query:
gitmap heatmap query "SELECT HeatmapActionLogId, ActionType, FilePath, ReasoningNotes FROM HeatmapActionLog ORDER BY HeatmapActionLogId DESC LIMIT 10" --json
```

### Step 4: Mirror to `.cursor/skills/gitmap/skill.md`
Verify that `.cursor/skills/gitmap/skill.md` is updated with identical content to preserve cross-IDE consistency.

---

## 5. Verification Evidence Contract

1. **Matrix Coverage Check:**
   - Execute: `gitmap aum search "Get-ChildItem (all variants)" .agents/skills/gitmap` -> Must return 1 match.
   - Execute: `gitmap aum search "gitmap lf \[dir\]" .agents/skills/gitmap` -> Must return matches.
2. **Parity Check:**
   - Execute: `gitmap aum search "Get-ChildItem (all variants)" .cursor/skills/gitmap` -> Must return 1 match.
3. **Heatmap Syntax Check:**
   - Execute: `gitmap aum search "gitmap heatmap log -t" .agents/skills/gitmap` -> Must return matches.
