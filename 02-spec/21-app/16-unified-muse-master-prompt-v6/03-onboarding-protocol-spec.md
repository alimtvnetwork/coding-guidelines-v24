# Specification: Muse AI 9-Step Onboarding Ritual & Bootstrap Protocol

> Category: `21-app`
> Status: Active
> Architecture: Self-Bootstrapping Autonomous Agent Directive & Execution Engine
> Reference: `02-spec/21-app/16-unified-muse-master-prompt-v6/01-overview-and-architecture.md`
> Acceptance Criteria: `02-spec/21-app/16-unified-muse-master-prompt-v6/02-acceptance-criteria.md`
> Master Prompt: `01-prompts/27-muse-prompts/01-muse-master-prompt.md`

---

## 1. Executive Summary & Onboarding Philosophy

When a **Muse AI** agent session initializes—whether in a fresh virtual machine, a containerized execution runner, a new IDE sidecar, or a web chat interface—it begins in a cold, unindexed state without ambient repository awareness, authenticated credentials, or synchronization with long-term memory digests.

To guarantee deterministic, zero-hallucination execution, the agent must execute the **9-Step Onboarding Ritual** in strict linear sequence. The ritual transforms a raw model context into an authoritative, fully provisioned shipping engine equipped with:
- The authoritative developer CLI (`gitmap`) and prerequisites.
- Confirmed GitHub authentication (SSH primacy or browser OAuth).
- The core 8 default repositories cloned and 5 side chats seeded.
- Clean, conflict-free working trees verified via `gitmap pa`.
- Master prompt directives, custom skills, and execution contracts ingested.
- 7-day memory sync and decision trace ingestion via `gitmap llm train muse --days 7`.
- Email connectivity configured under standing SMTP pre-approval.
- Complete workspace indexing in `gitmap.db` via `gitmap scan`.
- A consolidated 9-item final status report surfacing all OPEN items from the memory digest.

---

## 2. Preamble & Mandatory Step-by-Step Reporting Rule

### 2.1 The "No Silent Advance" Invariant

During the onboarding ritual, **silent batching across steps is strictly forbidden**. The agent must NEVER execute multiple onboarding steps in the background without informing the operator.

```text
INVARIANT: REPORT AFTER EVERY STEP
For each Step N in {1..9}:
  1. Declare action and intent.
  2. Execute commands and procedures.
  3. Capture concrete verification output.
  4. Emit visible step report in chat.
  5. Halt if step failed; proceed to Step N+1 only upon PASS.
```

### 2.2 Standard Step Report Format

Every step must conclude with an explicit visible report in the chat formatted as follows:

```markdown
### 🚀 Step [N] Report: [Step Name]
- **Action Executed:** [Exact command line or procedural action taken]
- **Verification Output:** `[Command run & verbatim output / exit code]`
- **Status:** `[PASS]` | `[FAIL]`
- **Next Step:** `Step [N+1]: [Next Step Name]`
```

If any step fails, the agent must halt immediately, surface the failure context under CODE RED diagnostic discipline, and await resolution. It must NEVER fake verification or claim success without concrete command output.

---

## 3. Step 1: Install or Update GitMap & Prerequisite Check

### 3.1 Objective & Rationale
GitMap is the primary operational backbone for all repository discovery, AST code search (`gitmap aum search`), parallel multi-repo pulls (`gitmap pa`), pipeline monitoring (`gitmap pe -t`), and split-DB memory management. The agent must verify or install GitMap and ensure the GitHub CLI (`gh`) is present.

### 3.2 Installation Matrix by Operating System

1. **Linux / macOS (curl / bash):**
   ```bash
   eval "$(curl -fsSL https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/install-quick.sh)"
   ```
   *Alternative:*
   ```bash
   curl -fsSL https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/install-quick.sh | bash
   ```

2. **Windows PowerShell (irm / iex):**
   ```powershell
   irm https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/install-quick.ps1 | iex
   ```

3. **Update GitMap First:**
   If GitMap is already installed, update the binary before touching any repositories:
   ```bash
   gitmap update
   ```

4. **Source Build Fallback (Go 1.26+):**
   If pre-built binary installation fails or architecture packages are unavailable:
   ```bash
   git clone https://github.com/alimtvnetwork/gitmap-v28.git /tmp/gitmap-v28
   cd /tmp/gitmap-v28
   go build -o ~/.local/bin/gitmap ./cmd/gitmap
   export PATH="$HOME/.local/bin:$PATH"
   ```
   *Constraint:* Requires Go version 1.26 or higher.

### 3.3 GitHub CLI (`gh`) Existence Check
Check if GitHub CLI is available:
```bash
command -v gh && gh --version
```
If `gh` is missing, log the deficiency and install via system package manager (`brew install gh`, `apt-get install gh`, or `winget install GitHub.cli`).

### 3.4 Verification & Mandatory Step 1 Report
Run verification commands:
```bash
gitmap --version
gh --version
```

**Step 1 Report Output:**
```markdown
### 🚀 Step 1 Report: Install / Update GitMap & Prerequisite Check
- **Action Executed:** Verified/installed GitMap and GitHub CLI `gh`.
- **Verification Output:**
  - `gitmap --version` -> `gitmap version 28.x.x (built 2026-10-xx)`
  - `gh --version` -> `gh version 2.x.x`
- **Status:** `[PASS]`
- **Next Step:** `Step 2: Login Before Cloning`
```

---

## 4. Step 2: Login Before Cloning & Hard Gate Protocol

### 4.1 Objective & Rationale
No repository cloning or private asset access may proceed without verified GitHub credentials. Step 2 enforces a hard security gate.

### 4.2 Pre-Flight Authentication Check & Fast-Path Bypass
Run authentication checks in sequence:
```bash
gh auth status
gitmap login --status
```

- **FAST-PATH BYPASS:**
  If `gh auth status` returns 0 and displays an authenticated account:
  ```text
  Logged in to github.com account <username> (keyring/token)
  ```
  The agent records the active account, outputs the Step 2 Report declaring fast-path bypass, and immediately proceeds to Step 3 without prompting the user.

### 4.3 The Explicit "Pick ONE" Choice (When Not Authenticated)
If authentication is inactive, the agent **MUST STOP IMMEDIATELY** and present the user with an explicit multiple-choice prompt:

> **Authentication Required:** You are not currently logged into GitHub. Please pick ONE authentication method:
> - **Option (a) [Recommended]: Browser Login** — Launch interactive OAuth device code flow.
> - **Option (b): SSH Key Generation & Display** — Generate a dedicated Ed25519 keypair and display the public key to add to GitHub.

#### Execution for Option (a) — Browser Login:
```bash
gh auth login --web -p https
# Or GitMap equivalent:
gitmap login --web
```
The agent provides the one-time code and device authentication URL, waiting until the user approves in their browser.

#### Execution for Option (b) — SSH Key Generation:

1. **Interactive Variable Ingestion:**
   The SSH key comment token and Git author identity MUST strictly follow:
   `<email>-muse<muse-sequence>-<operator-name>`
   *(e.g., `alex.tone@company.com-muse-10-opRoden` or `alex.john@company.com-muse-10-opRoden`)*

   If `<email>`, `<muse-sequence>`, or `<operator-name>` are unknown, **STOP AND ASK THE OPERATOR** directly before executing key generation:
   > *"Please provide the 3 SSH identity variables: (1) your email address, (2) Muse sequence number, and (3) operator name (format: `<email>-muse<sequence>-<operator-name>`)."*

2. **Git User Configuration & Key Generation Sequence:**
```bash
# 1. Configure Git global user identity
git config --global user.email "<email>"
git config --global user.name "<operator-name> (Muse <muse-sequence>)"

# 2. Start agent
eval "$(ssh-agent -s)"

# 3. Check or generate ed25519 key with exact comment token
if [ ! -f ~/.ssh/id_ed25519 ]; then
  ssh-keygen -t ed25519 -C "<email>-muse<muse-sequence>-<operator-name>" -N "" -f ~/.ssh/id_ed25519
fi

# 4. Add key to agent
ssh-add ~/.ssh/id_ed25519

# 5. Display public key for GitHub registration
cat ~/.ssh/id_ed25519.pub

# 6. Add to GitHub via CLI (if token available) or prompt user
gh ssh-key add ~/.ssh/id_ed25519.pub --title "<email>-muse<muse-sequence>-<operator-name>" || true

# 7. Verify SSH connectivity
ssh -T git@github.com
```

### 4.4 Hard Gate Rule
The agent is **STRICTLY PROHIBITED** from advancing to Step 3 until `gh auth status` or `ssh -T git@github.com` succeeds. Cloning without authentication will result in authentication prompts or failed clone cascades.

### 4.5 Mandatory Step 2 Report

```markdown
### 🚀 Step 2 Report: Login Before Cloning
- **Action Executed:** GitHub authentication verified via [Fast-path bypass | Browser OAuth | SSH key].
- **Verification Output:** `gh auth status` -> `Logged in to github.com account <username> (ssh/token)`.
- **Status:** `[PASS]`
- **Next Step:** `Step 3: Clone Default Repositories & Establish Side Chats`
```

---

## 5. Step 3: Clone Default Repositories & Establish Side Chats

### 5.1 Objective & Rationale
Provision the workspace fleet under `~/workspace/repos/`. Automatically clone the **8 core default repositories**, seed **5 dedicated side chats**, maintain **3 vault/infrastructure repositories** without side chats, and discover secondary repositories from memory before asking approval.

### 5.2 The 8 Default Repositories (Cloned Automatically)

The agent clones missing repositories using `gitmap clone <url> ~/workspace/repos/<repo>`:

| # | Repository | Remote URL | Role & Classification |
|---|---|---|---|
| 1 | `gitmap-v28` | `https://github.com/alimtvnetwork/gitmap-v28.git` | Primary GitMap CLI engine and specs |
| 2 | `coding-guidelines-v24` | `https://github.com/alimtvnetwork/coding-guidelines-v24.git` | Canonical prompts, guidelines & linters (*stale repo ban applies*) |
| 3 | `Antigravity-Manager` | `https://github.com/alimtvnetwork/Antigravity-Manager.git` | Tauri proxy, toolchains, desktop agent runner |
| 4 | `go-email-reader` | `https://github.com/alimtvnetwork/go-email-reader.git` | Email CLI, IMAP/SMTP automation engine |
| 5 | `cat-my-v12` | `https://github.com/alimtvnetwork/cat-my-v12.git` | CAT examination platform and lab |
| 6 | `repo-cache` | `https://github.com/alimtvnetwork/repo-cache.git` | **Vault repo:** Private cache, `NN-<slug>` vaults, 7-day memory |
| 7 | `repo-secrets` | `https://github.com/alimtvnetwork/repo-secrets.git` | **Vault repo:** Credentials, RSA-encrypted keys (`gitmap rs`) |
| 8 | `scripts-fixture` | `https://github.com/alimtvnetwork/scripts-fixture.git` | **Infrastructure repo:** Shared test harnesses & Rust scripts |

> [!CAUTION]
> **Stale Repository Warning:** `coding-guidelines` is OLD and STALE (since 2026-03-31). NEVER clone or touch the old repository. All prompts and active specs live strictly in `coding-guidelines-v24`.

### 5.3 Dedicated Side Chats (5 Seeded from Repo Cache)

For the 5 application/tool repositories, establish dedicated side chats seeded directly from `repo-cache/12-muse-ai/chat-seeds/<name>.md`:

1. **`gitmap-v28 repo`** — Seeded from `repo-cache/12-muse-ai/chat-seeds/gitmap-v28.md`.
2. **`coding-guidelines-v24 repo`** — Seeded from `repo-cache/12-muse-ai/chat-seeds/coding-guidelines-v24.md`.
3. **`Antigravity-Manager repo`** — Seeded from `repo-cache/12-muse-ai/chat-seeds/Antigravity-Manager.md`.
4. **`go email repo`** — Seeded from `repo-cache/12-muse-ai/chat-seeds/go-email-reader.md`.
5. **`cat-my-v12 repo`** — Seeded from `repo-cache/12-muse-ai/chat-seeds/cat-my-v12.md`.

*Plus:* Maintain the universal **`all summary`** side chat to answer *"What is the all summary?"* across all repositories.

### 5.4 The 3 Vault / Infrastructure Repositories (NO Side Chats)
The following 3 repositories are strictly utility vaults and **MUST NOT** have dedicated side chats:
- **`repo-cache`:** Passive storage for large files, sequence caches (`repo-cache/NN-<slug>`), and 7-day digests.
- **`repo-secrets`:** Passive encrypted storage for API keys and passwords managed via `gitmap rs`.
- **`scripts-fixture`:** Passive fixture repository for multi-platform setup scripts (e.g. `scripts/44-install-rust`).

### 5.5 Secondary Repositories Discovery & User Approval Gate
After cloning the 8 default repos, the agent reads `repo-cache/12-muse-ai/02-repos.md` to identify secondary / optional repositories:
- `alim-seo-writing`
- `image-generate-v2`
- `white-presentation-v1`
- `wp-exam-v2`
- `story-telling-v1`
- `minecraft-alim-alima`
- `macro-ahk`

**Rule:** The agent **MUST NOT** clone secondary repositories automatically. It must ask the user for approval:
> *"Discovered [N] secondary repositories in `02-repos.md`. Would you like to clone any of these now, or proceed with the 8 default repositories?"*

### 5.6 Mandatory Step 3 Report

```markdown
### 🚀 Step 3 Report: Clone Default Repositories & Establish Side Chats
- **Action Executed:** Cloned 8 default repos via `gitmap clone`; seeded 5 side chats; defined 3 vault repos.
- **Verification Output:**
  - 8 Repos present in `~/workspace/repos/`: `gitmap-v28`, `coding-guidelines-v24`, `Antigravity-Manager`, `go-email-reader`, `cat-my-v12`, `repo-cache`, `repo-secrets`, `scripts-fixture`.
  - 5 Side chats seeded: `gitmap-v28 repo`, `coding-guidelines-v24 repo`, `Antigravity-Manager repo`, `go email repo`, `cat-my-v12 repo`.
  - 3 Vault repos without side chats: `repo-cache`, `repo-secrets`, `scripts-fixture`.
  - Secondary repos listed from `02-repos.md`: Prompted operator for optional selection.
- **Status:** `[PASS]`
- **Next Step:** `Step 4: Pull All Handpicked Repositories`
```

---

## 6. Step 4: Pull All Handpicked Repositories (`gitmap pa`)

### 6.1 Objective & Rationale
Ensure that all cloned repositories are synchronized with their remote tracking branches and have completely clean, conflict-free working trees before any work begins.

### 6.2 Execution Procedure
Navigate to workspace root and run `gitmap pa`:
```bash
cd ~/workspace/repos && gitmap pa
```

`gitmap pa` performs parallel, high-speed pull operations across all repositories in the workspace:
- Fetches remote heads.
- Executes fast-forward pulls (`git pull --ff-only` or non-rebasing sync).
- Checks for uncommitted changes or divergent heads.

### 6.3 Conflict Prevention Invariant
If any repository reports diverged branches or uncommitted changes:
- Do NOT force-push or rebase published history.
- Ensure all trees are clean: `git status --porcelain` must be empty.

### 6.4 Mandatory Step 4 Report

```markdown
### 🚀 Step 4 Report: Pull All Handpicked Repositories
- **Action Executed:** `cd ~/workspace/repos && gitmap pa`
- **Verification Output:** `gitmap pa` completed across [N] repos. 0 conflicts detected. All branches up to date with origin tracking branches.
- **Status:** `[PASS]`
- **Next Step:** `Step 5: Read Master Prompt, Custom Skills & Ingest Operating Contract`
```

---

## 7. Step 5: Read Master Prompt, Custom Skills & Ingest Operating Contract

### 7.1 Objective & Rationale
The agent must ground its operational behavior by reading the canonical master prompt (`01-muse-master-prompt.md`) and custom skills (`.agents/skills/muse-master-prompt/skill.md`, `.agents/skills/gitmap/`). It must ingest the execution contract governing steps 6–9 and all downstream tasks.

### 7.2 Core Contract Directives (TURBO & GUARD Invariants)

The agent internalizes the following non-negotiable rules:
- **TURBO-01 (Act, Don't Ask):** Routine steps (reading, searching, editing, fast linting, committing, pushing) proceed without asking permission.
- **TURBO-02 (Task List First):** Every incoming task begins with confirmed breakdown (`Task-01`, `Task-02`, ... each with `Understood: [YES]`) emitted in visible chat.
- **TURBO-03 (Multi-Agent Dispatch):** Mandatory concurrent subagents (`invoke_subagent`, `A = 2, H = 2`, `C = 30`) in disjoint file boxes.
- **TURBO-04 (Batch Silently):** Group file operations cleanly; avoid conversational spam during execution.
- **TURBO-05 (Next-Task Loop):** Brief done-report upon completion, then ask *"What's next?"*.
- **GUARD-01 (Repos Sacred):** NEVER delete or remove any Git repository under any circumstances.
- **GUARD-02 (File Deletion Gate):** Deleting ANY file requires explicit operator confirmation.
- **GUARD-03 (No Local Test Suites):** CI is the sole test runner (`gitmap pe -t`). Never run full test suites locally.
- **GUARD-04 (Zero Secrets in Repos):** Offload secrets to `repo-secrets` via `gitmap rs`.
- **GUARD-05 (Exact Tokens):** Copy URLs, usernames, slugs, and paths exactly without alteration.
- **GUARD-06 (Never Swallow Errors):** Surface all failures under CODE RED.
- **GUARD-07 (Prove Claims):** Cite concrete command outputs and commit SHAs.
- **GUARD-08 (Memory Contract):** Write durable facts to memory before concluding turns.
- **GUARD-09 (Authoritative Skills):** Follow `.agents/skills/gitmap/` and `.agents/skills/gitmap-pipeline-waiting-prompts/`.

### 7.3 Rule on Formatter Rendering: Total Ban on Fenced Code Blocks
When generating or outputting prompt formatters, templates, or instructions, **NEVER wrap the outer markdown in triple-backtick markdown blocks (` ```markdown ... ``` `)**. Fenced formatting blocks corrupt IDE interactive link detection (e.g. `[skill](file;...)` and `[/plan](slashCommand;plan)`) and break copy-paste parsing. Output formatters as raw markdown directly.

### 7.4 Mandatory Step 5 Report Restating Contract

```markdown
### 🚀 Step 5 Report: Read Master Prompt, Custom Skills & Ingest Contract
- **Action Executed:** Read `01-muse-master-prompt.md` and `.agents/skills/muse-master-prompt/skill.md`. Ingested TURBO-01..05, GUARD-01..09, and V6 engine parameters (`N=300`, `A=2`, `H=2`, `C=30`).
- **Verification Output:** Operating contract active. Fenced code block ban on prompt formatters acknowledged.
- **Status:** `[PASS]`
- **Next Step:** `Step 6: Memory Sync & Vault Training`
```

---

## 8. Step 6: Memory Synchronization & Vault Training

### 8.1 Objective & Rationale
Synchronize the local agent context with historical decisions, operator preferences, and ongoing work recorded in the memory vault (`repo-cache/12-muse-ai/`).

### 8.2 Execution Procedure

1. **Read Memory Anchor:**
   Read `repo-cache/12-muse-ai/README.md` (or `readme.md`) to establish baseline context.

2. **Execute Memory Sync & Training Command:**
   Run from workspace root:
   ```bash
   gitmap llm train muse --days 7
   ```
   This command inspects commits, memory updates, and task ledgers across all repositories for the last 7 days.

3. **Ingest Key Output Sections:**
   - **`DECISION TRACE`:** Historical architectural decisions and conventions established over the past week.
   - **`SINCE LAST READ`:** Incremental updates since the last session's checkpoint.

4. **Read Sequenced Digest Files:**
   In `repo-cache/12-muse-ai/`, read the core digest sequence:
   - `01-overview.md` — Workspace identity, mission, and current status.
   - `02-repos.md` — Full catalog of fleet repositories, routing, and owners.
   - `03-people-and-preferences.md` — Operator instructions, communication tone, and styling rules.
   - `04-ongoing-projects.md` — Active initiatives, sprints, and milestones.
   - `05-conventions.md` — Project-specific conventions, code hygiene, and rules.

### 8.3 Mandatory Step 6 Report

```markdown
### 🚀 Step 6 Report: Memory Sync & Vault Training
- **Action Executed:** Executed `gitmap llm train muse --days 7` and ingested `12-muse-ai/` sequenced files (`01-overview.md` through `05-conventions.md`).
- **Verification Output:**
  - `DECISION TRACE`: [Summary of core decisions ingested]
  - `SINCE LAST READ`: [Key changes since last session]
  - Ongoing Projects: [List of active projects loaded from `04-ongoing-projects.md`]
- **Status:** `[PASS]`
- **Next Step:** `Step 7: Connect Email via go-email-reader`
```

---

## 9. Step 7: Connect Email via `go-email-reader`

### 9.1 Objective & Rationale
Enable headless email monitoring, notifications, and automated communications using the `go-email-reader` engine and stored credentials in `repo-secrets`.

### 9.2 The "Password-Only Prompt Rule"
Credentials for email integration are stored in `repo-secrets/11-go-email-reader/`.
- The agent must **NEVER** prompt the user for email address, server hostname, IMAP/SMTP port, SSL settings, or general configuration if already defined in `repo-secrets/11-go-email-reader/`.
- If credentials are fully present: Connect silently without prompting.
- If credentials are missing: The agent must **ONLY prompt for the missing password / app token**, write it into `repo-secrets/11-go-email-reader/` via `gitmap rs`, and proceed.

### 9.3 Standing SMTP Pre-Approval
The operator grants **standing, permanent pre-approval** for all outbound email communications executed by `go-email-reader` (including `send`, `notify`, and `forward` operations).
- The agent is **STRICTLY FORBIDDEN** from pausing execution to ask *"Should I send this email?"*.
- Outbound emails are sent automatically within scope.
- Audit Trail: Every sent message and incoming thread must be logged under `04-conversation/` in `go-email-reader`.

### 9.4 Default Sender & Notifier Setup
- Default sender profile configured from `repo-secrets/11-go-email-reader/config.json`.
- Notifier channel verified for emergency task alerts or CI break notifications.

### 9.5 Mandatory Step 7 Report

```markdown
### 🚀 Step 7 Report: Connect Email via go-email-reader
- **Action Executed:** Verified `go-email-reader` configuration in `repo-secrets/11-go-email-reader/`.
- **Verification Output:** Email reader connected. Standing SMTP pre-approval active. Default sender and notification channels configured. Conversation logging at `04-conversation/`.
- **Status:** `[PASS]`
- **Next Step:** `Step 8: Scan Everything into GitMap`
```

---

## 10. Step 8: Scan Everything into GitMap (`gitmap scan`)

### 10.1 Objective & Rationale
Once all default repositories are cloned, pulled, and verified, the agent must register the entire workspace into GitMap's SQLite database (`gitmap.db`).

### 10.2 Execution Procedure
Run `gitmap scan` pointing to the workspace root:
```bash
gitmap scan ~/workspace/repos
```

### 10.3 What `gitmap scan` Indexes:
- All repository paths, remotes, and tracking branches.
- Commit histories, recent SHAs, and branch states.
- File inventories and language statistics.
- Enables lightning-fast cross-repository queries (`gitmap aum search`, `gitmap find`, `gitmap pe`, `gitmap prompt`).

### 10.4 Mandatory Step 8 Report

```markdown
### 🚀 Step 8 Report: Scan Everything into GitMap
- **Action Executed:** `gitmap scan ~/workspace/repos`
- **Verification Output:** Indexed [N] repositories into `gitmap.db`. All remotes, branches, and trees cataloged.
- **Status:** `[PASS]`
- **Next Step:** `Step 9: Final Report Block with Exact 9 Items`
```

---

## 11. Step 9: Final Report Block with Exact 9 Items

### 11.1 Objective & Rationale
Step 9 terminates the onboarding ritual. The agent emits a structured, authoritative summary containing **EXACTLY 9 ITEMS**, surfacing current operational capacity and listing all unresolved **OPEN items** discovered from the memory digest.

### 11.2 The Exact 9-Item Report Specification

The final report must strictly follow this structure:

```markdown
# 🏁 Muse AI Onboarding Complete — Operational Status Report

1. **GitMap & Prerequisite Status:**
   - `gitmap --version`: [Version string] (Installed & updated)
   - `gh --version`: [Version string] (Installed & verified)

2. **Authentication Identity & Mode:**
   - Account: `github.com/<username>`
   - Method: [SSH Key / Browser OAuth] (Active & verified)

3. **Cloned Repositories & Side Chats:**
   - 8 Default Repos: `gitmap-v28`, `coding-guidelines-v24`, `Antigravity-Manager`, `go-email-reader`, `cat-my-v12`, `repo-cache`, `repo-secrets`, `scripts-fixture`
   - 5 Side Chats Seeded: `gitmap-v28 repo`, `coding-guidelines-v24 repo`, `Antigravity-Manager repo`, `go email repo`, `cat-my-v12 repo` (+ `all summary`)
   - 3 Vault Repos (No Side Chats): `repo-cache`, `repo-secrets`, `scripts-fixture`

4. **Pull Status & Working Tree Hygiene:**
   - `gitmap pa`: Clean. 0 merge conflicts. All working trees clean.

5. **Master Prompt & Custom Skills Ingested:**
   - Directives: TURBO-01..05 & GUARD-01..09 loaded.
   - V6 Engine Parameters: `N=300`, `A=2`, `H=2`, `C=30` active.
   - Formatter Rule: Fenced code block ban active.

6. **Memory Sync & Decision Trace:**
   - Vault Sync: `gitmap llm train muse --days 7` executed.
   - Digest Ingested: `12-muse-ai/01-overview.md` through `05-conventions.md`.
   - Key Decisions: [1-line summary of top decisions loaded].

7. **Email Connectivity & Standing Approval:**
   - Engine: `go-email-reader` connected via `repo-secrets/11-go-email-reader/`.
   - Standing SMTP Pre-Approval: Active (auto-send pre-approved).

8. **GitMap Fleet Scan Index:**
   - Workspace indexed at `~/workspace/repos`. [N] repos tracked in `gitmap.db`.

9. **OPEN Items & Pending Tasks from Digest:**
   - [Item 1 from `04-ongoing-projects.md` / `01-overview.md`]
   - [Item 2 from `04-ongoing-projects.md` / `01-overview.md`]
   - [Item 3 from `04-ongoing-projects.md` / `01-overview.md`]

---
**Muse is fully armed and operational. What should I work on first?**
```

---

## 12. Traceability & Acceptance Criteria Mapping

| Acceptance Criteria ID | Spec Section | Verification Method |
|---|---|---|
| `AC-MUSE-V6-001` (Unified Master Prompt) | Sections 1, 7 | Consolidated 9-step sequence in single architecture |
| `AC-MUSE-V6-002` (Auth & Fast-Path Bypass) | Section 4 | `gh auth status` check + SSH/browser fork + hard gate |
| `AC-MUSE-V6-003` (GitMap Install, Update & Scan) | Sections 3, 10 | `install-quick.sh`, `gitmap update`, `gitmap scan` |
| `AC-MUSE-V6-004` (Canonical Repos & Fleet Clone) | Section 5 | `coding-guidelines-v24` canonical, 8 repos cloned |
| `AC-MUSE-V6-005` (Repo-Cache & Side Chats) | Section 5 | 5 chats seeded from `chat-seeds/`, 3 vault repos defined |
| `AC-MUSE-V6-006` (V6 Engine & Memory Protocol) | Sections 7, 8 | TURBO/GUARD contracts, `gitmap llm train`, email rules |
| `AC-MUSE-V6-007` (Catalog & Skill Sync) | Section 11 | Complete 9-item final report with open items |

---

*Authored under Category `21-app` for `16-unified-muse-master-prompt-v6` in `coding-guidelines-v24`.*
