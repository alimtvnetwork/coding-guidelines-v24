# Architecture Specification: Icon Generation Modernization, Flat Commit Integration & 400-Step Lifecycle

## User Request (Verbatim)

```text
You are Spec Author 01 for task coding-guideline-and-skill-optimization. You have no prior chat context; this brief is your complete specification.

### Boundaries & Crash Prevention:
- You are an authoring subagent (`TypeName: 'self'`).
- Edit and create ONLY your assigned disjoint files:
  1. `02-spec/21-app/coding-guideline-and-skill-optimization/01-architecture-spec.md`
  2. `.ai-memory/plans/subtasks/coding-guideline-and-skill-optimization/01-icon-and-flat-commit-spec.md`
- TOTAL BAN ON RAW SEARCH TOOLS: NEVER run `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem -Recurse`, or `findstr`.
- TOTAL BAN ON GIT COMMANDS: NEVER run `git add`, `git commit`, `git push`, `git status`, `git diff`, `git checkout`.
- Strict Relative Git Paths: All paths cited in files MUST be strictly relative to the git repository root. TOTAL BAN on absolute filesystem paths (C:\..., /home/...) and `file:///` URIs.
- Strict Lowercase Filenames: All files, directory paths, and references must be strictly lowercase.
- Boolean Standards: Use positive booleans (`is`, `has`), never compare explicit `== true`.

### Required Spec Content for 02-spec/21-app/coding-guideline-and-skill-optimization/01-architecture-spec.md:
1. Executive Summary & Problem Statement: Modernizing icon prompts and skills, introducing flat commit documentation across CLI and UI help guides, establishing a 400-step budget.
2. Icon Generation Architecture:
   - 100% Transparent Canvas (`alpha = 0`) rule.
   - Dual-Mode Inverted Artwork Clarity: Dark Mode means inverted high-luminance strokes (luminance > 0.70, e.g. #FFFFFF, #FCD34D, #38BDF8) on a 100% transparent canvas—NEVER a dark box container (<rect>)! Light Mode uses rich saturated hues (luminance < 0.40).
   - SVG Precision & Quiet Zone: 8-10% padding to prevent clipping at small favicon sizes (16px), 2-decimal rounded coordinates, semantic `<defs><linearGradient>`, accessibility (`<title>`, `<desc>`, `role="img"`, `aria-labelledby`), and `fill="currentColor"` for monochrome glyphs.
   - Design Token Alignment: Mapping color attributes to design system tokens (`--pres-accent`, `--pres-text`, `--primary`, `--foreground`).
3. Flat Commit Segment Specification:
   - Exact CLI Syntax: `gitmap commit "<message>"` and alias `gitmap cm "<message>"`.
   - Mechanics: 1-step auto-stage (`git add -A`) and local commit without semantic prefixes ('Feature: ', 'Bug: ') and without remote push.
   - Contrast with Semantic Commits (`gitmap cpf`, `gitmap cpb`, `gitmap cpr`): Semantic commits auto-push and format with prefixes. Flat commits are local.
   - Completion Invariant Reminder: While flat commit is useful for local iteration, final task completion mandates the 'No Push = Not Done' invariant via `gitmap cpf` or explicit push.
   - Documentation additions to CLI help guides, UI docs, and AGENTS.md §10.
4. 400-Step Budget Specification:
   - Phase 1 (Steps 1..100): Planning, discovery subagents, specs, and Brain roadmap.
   - Phase 2 (Steps 101..250): Intermediate tasks, archival, worktree sanitization, and prerequisites.
   - Phase 3 (Steps 251..400): Implementation, worker waves, targeted checks, fleet synchronization, and atomic push.
```

---

## 1. Executive Summary & Problem Statement

### 1.1 Context & Purpose
The Prompt Architect ecosystem relies on standardized coding guidelines, prompt engineering templates, and automated tooling to coordinate autonomous AI agents across repositories. Over multiple production cycles, three key friction points have been identified across our prompt libraries, documentation, and agent workflows:

1. **Suboptimal Icon & Vector Asset Generation:** Existing icon prompts and skills frequently suffer from visual clipping at small favicon sizes (16px/32px), inconsistent coordinate precision, lack of design token binding, and a critical anti-pattern where AI agents wrap dark mode vector icons in artificial dark rectangular backgrounds (`<rect fill="#1e293b">`) instead of rendering inverted high-luminance strokes directly on a transparent canvas.
2. **Missing Flat Commit Documentation:** Developers and automated agents frequently perform fast, iterative local checkpoints where semantic prefixes (`Feature: `, `Bug: `) and instantaneous remote git pushes (`git push origin <branch>`) are undesirable. While GitMap provides `gitmap commit` (alias `gitmap cm`) for 1-step local staging and committing, this capability lacks structured documentation in CLI help guides, UI documentation, and repository conventions.
3. **Unstructured Step Budgets for Long-Horizon Tasks:** Large-scale refactors and guideline migrations require predictable execution horizons. Without an explicit phase budget, multi-agent runs risk either premature termination or unbounded task thrashing. A formalized 400-step budget establishes deterministic guardrails across discovery, sanitization, and parallel execution.

### 1.2 Core Objectives
- Establish an authoritative vector standard for SVG/PNG icon generation featuring mandatory canvas transparency, dual-mode inverted artwork, quiet-zone geometry, and design-token alignment.
- Codify the specification, mechanics, and operational boundaries of GitMap flat commits (`gitmap commit` / `gitmap cm`) alongside semantic auto-pushing commits (`gitmap cpf`, `gitmap cpb`, `gitmap cpr`), reinforcing the global 'No Push = Not Done' completion invariant.
- Architect the 400-Step Execution Lifecycle across three discrete, bounded phases to standardize agent task velocity, tree-swarm coordination, and atomic release hygiene.

---

## 2. Icon Generation Architecture

### 2.1 100% Transparent Canvas (`alpha = 0`) Rule
Vector assets (SVG) and raster exports (PNG) created for logos, app glyphs, favicons, and HUD indicators must never include an opaque background layer unless explicitly mandated by the user as a physical card component.

- **Mandatory Root Transparency:** The root `<svg>` element must never have a background fill style, and must never contain a root-level `<rect width="100%" height="100%" ...>` or bounding backdrop.
- **Raster Export Alpha:** All rendered PNG graphics must be exported with a 32-bit RGBA channel where background pixels have `alpha = 0`.
- **Composability Invariant:** Vector assets must seamlessly overlay diverse host backgrounds—including dark HUDs, glassmorphic presentation slides, white documentation pages, and system docks—without bounding-box halos or mismatched corner radiuses.

```xml
<!-- ❌ WRONG: Background container rect enclosing the icon -->
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48">
  <rect width="48" height="48" rx="8" fill="#0f172a"/>
  <path d="..." stroke="#38bdf8"/>
</svg>

<!-- ✅ CORRECT: Pure transparent canvas, artwork floats freely -->
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" fill="none">
  <path d="..." stroke="#38bdf8"/>
</svg>
```

### 2.2 Dual-Mode Inverted Artwork Clarity
A major source of AI hallucination is conflating **Dark Mode Icons** with **Dark Card Backgrounds**. In this architecture:

- **Dark Mode Requirement (Inverted High-Luminance Strokes):**
  - Dark mode assets represent the artwork itself inverted for readability against dark host environments.
  - Strokes and primary fills must utilize high-luminance tones (perceptual luminance $L > 0.70$), such as `#ffffff` ($L = 1.0$), `#fcd34d` ($L \approx 0.82$), `#38bdf8` ($L \approx 0.78$), or `#a7f3d0` ($L \approx 0.88$).
  - **TOTAL BAN ON DARK BOX CONTAINERS:** Dark mode icons must NEVER include a dark rectangular container (`<rect>`, `<circle>`, or `<path>` acting as a dark card backdrop). The background remains 100% transparent.
- **Light Mode Requirement (Rich Saturated Hues):**
  - Light mode assets utilize deep, rich, high-contrast saturated hues (perceptual luminance $L < 0.40$), such as deep indigos (`#1e1b4b`), slate navy (`#0f172a`), deep cobalt (`#1d4ed8`), or dark emerald (`#064e3b`).
  - The canvas remains 100% transparent.

| Mode | Canvas State | Stroke / Fill Luminance Range | Primary Palette Samples | Banned Elements |
|:---|:---|:---|:---|:---|
| **Dark Mode** | 100% Transparent (`alpha = 0`) | High Luminance ($L > 0.70$) | `#ffffff`, `#fcd34d`, `#38bdf8`, `#cbd5e1` | `<rect>` dark backdrops, dark cards |
| **Light Mode** | 100% Transparent (`alpha = 0`) | Low Luminance ($L < 0.40$) | `#0f172a`, `#1e1b4b`, `#1e3a8a`, `#14532d` | Artificial border boxes, matte backgrounds |

### 2.3 SVG Precision & Quiet Zone Geometry
To prevent visual clipping, rendering artifacts, and blurry subpixel anti-aliasing across displays:

- **Quiet Zone (8–10% Inner Padding):**
  - All visual vector paths, strokes, and markers must remain within the inner 80–84% safe area of the `viewBox`.
  - For a standard `0 0 24 24` viewBox, maintain a minimum 2.0px quiet zone margin on all four boundaries (content bounds: `x: 2..22`, `y: 2..22`).
  - For a `0 0 48 48` viewBox, maintain a minimum 4.0px quiet zone margin (content bounds: `x: 4..44`, `y: 4..44`).
  - Prevents stroke width clipping when rendered as 16px/32px desktop browser favicons, mobile homescreen shortcuts, or IDE activity bar icons.
- **Coordinate Precision (2 Decimal Places):**
  - All path data coordinates (`d="..."`), transformation matrices, and dimension attributes must be rounded to a maximum of 2 decimal places (e.g., `d="M 12.50 4.25 L 18.75 14.10 ..."`).
  - Eliminates bloated floating-point noise (e.g. `12.500000000004px`), minimizes DOM payload size, and optimizes browser SVG geometry parsing.
- **Semantic Gradients & Defs:**
  - Complex fills must be grouped inside `<defs>` with unique, descriptive, lowercase IDs (e.g. `id="gradient-accent-flow"`).
  - Gradients must specify explicit coordinate space (`gradientUnits="userSpaceOnUse"` or `objectBoundingBox`) and linear stops with color and opacity.
- **Accessibility & Compliance:**
  - Root element must declare `role="img"`.
  - Provide semantic descriptions via `<title id="icon-title">` and `<desc id="icon-desc">` referenced via `aria-labelledby="icon-title icon-desc"`.
  - Monochrome glyphs must use `fill="currentColor"` or `stroke="currentColor"` to inherit parent CSS styling dynamically.

```xml
<svg
  xmlns="http://www.w3.org/2000/svg"
  viewBox="0 0 48 48"
  fill="none"
  role="img"
  aria-labelledby="app-icon-title app-icon-desc"
  class="icon-glyph"
>
  <title id="app-icon-title">Workflow Execution Engine</title>
  <desc id="app-icon-desc">Dual node connector symbol with high-luminance accent gradient</desc>
  <defs>
    <linearGradient id="gradient-accent-flow" x1="4.00" y1="4.00" x2="44.00" y2="44.00" gradientUnits="userSpaceOnUse">
      <stop offset="0.00" stop-color="#38bdf8"/>
      <stop offset="1.00" stop-color="#fcd34d"/>
    </linearGradient>
  </defs>
  <!-- Content strictly bounded inside 4.00px quiet zone (4.00 to 44.00) -->
  <path
    d="M 12.00 8.00 L 36.00 8.00 C 40.42 8.00 44.00 11.58 44.00 16.00 L 44.00 32.00 C 44.00 36.42 40.42 40.00 36.00 40.00 L 12.00 40.00 C 7.58 40.00 4.00 36.42 4.00 32.00 L 4.00 16.00 C 4.00 11.58 7.58 8.00 12.00 8.00 Z"
    stroke="url(#gradient-accent-flow)"
    stroke-width="2.50"
    stroke-linecap="round"
    stroke-linejoin="round"
  />
  <circle cx="24.00" cy="24.00" r="6.00" fill="#ffffff"/>
</svg>
```

### 2.4 Design Token Alignment
Icons must directly reference or map to our established design token system to ensure strict aesthetic coherence across Lovable web applications, presentations (`slides-app`), and desktop interfaces.

- **Presentation Engine Tokens:**
  - Primary Accent: `var(--pres-accent, #38bdf8)`
  - Primary Content / Typography: `var(--pres-text, #f8fafc)`
  - Subdued Accents: `var(--pres-muted, #94a3b8)`
- **Tailwind / Shadcn Tokens:**
  - Primary Stroke / Fill: `var(--primary, #0f172a)`
  - Foreground Contrast: `var(--foreground, #020817)`
  - Muted Foreground: `var(--muted-foreground, #64748b)`
- **Dynamic CSS Variable Binding:**
  When rendering inline SVGs in React or HTML components, color attributes should leverage CSS variable fallbacks (e.g. `stroke="var(--pres-accent, #38bdf8)"`) or leverage `currentColor` alongside CSS utility classes (`text-primary`, `text-accent`).

---

## 3. Flat Commit Segment Specification

### 3.1 Syntax & CLI Command Surface
GitMap introduces and formalizes the **flat commit** command set for rapid, single-step staging and local committing:

- **Primary Command:** `gitmap commit "<message>"`
- **Canonical Alias:** `gitmap cm "<message>"`

```bash
# Standard syntax
gitmap commit "update icon generator quiet zone specifications"

# High-velocity alias
gitmap cm "refactor boolean evaluation in token parser"
```

### 3.2 Internal Mechanics & Behavioral Contract
The flat commit command executes a lean, deterministic two-step atomic sequence:
1. **Auto-Stage (`git add -A`):** Gathers all modified, newly created, and deleted tracked/untracked files within the active git working tree.
2. **Local Commit (`git commit -m "<message>"`):** Records the staged changes locally using the verbatim `<message>` passed by the user.

Key behavioral traits:
- **No Semantic Message Mutation:** Unlike semantic commit commands, `gitmap commit` DOES NOT prepend prefixes such as `Feature: `, `Bug: `, or `Refactor: `. The exact message string provided is written to the git commit object.
- **No Remote Git Push:** `gitmap commit` operates purely on the local repository. It NEVER triggers `git push`, never interacts with remote tracking branches, and never communicates over SSH or HTTPS remotes during execution.
- **Speed & Isolation:** Designed for rapid micro-batch checkpoints, scratch saves, and pre-refactor local commits without consuming network bandwidth or triggering remote CI/CD webhooks prematurely.

### 3.3 Contrast: Flat Commits vs. Semantic Auto-Pushing Commits

| Feature Attribute | Flat Commit (`gitmap cm` / `gitmap commit`) | Semantic Feature (`gitmap cpf`) | Semantic Bugfix (`gitmap cpb`) | Semantic Refactor (`gitmap cpr`) |
|:---|:---|:---|:---|:---|
| **Staging** | Automatic (`git add -A`) | Automatic (`git add -A`) | Automatic (`git add -A`) | Automatic (`git add -A`) |
| **Commit Prefix** | None (Verbatim message) | Prepends `Feature: ` | Prepends `Bug: ` | Prepends `Refactor: ` |
| **Message Arg Constraint** | Free text, no colon restriction | Hyphen format: `<module> - <summary>` | Hyphen format: `<module> - <summary>` | Hyphen format: `<module> - <summary>` |
| **Remote Git Push** | **No** (Local commit only) | **Yes** (Pushes to remote branch) | **Yes** (Pushes to remote branch) | **Yes** (Pushes to remote branch) |
| **CI/CD Pipeline Trigger** | No (Local only) | Yes (Remote push triggers pipeline) | Yes (Remote push triggers pipeline) | Yes (Remote push triggers pipeline) |
| **Intended Use Case** | Local WIP checkpoints, scratch steps | Completed feature subtasks | Completed bugfix subtasks | Completed architecture refactors |

### 3.4 Completion Invariant Reminder ('No Push = Not Done')
While flat commits (`gitmap commit`, `gitmap cm`) provide a fast, reliable mechanism for intermediate progress snapshots within multi-step tasks, they **DO NOT satisfy the repository completion invariant**.

Under `AGENTS.md` §13:
> [!IMPORTANT]
> **Mandatory Completion Invariant ('No Push = Not Done'):**
> If code changes are not committed and pushed upstream to GitHub (main/master/tracking branch), the task is strictly considered **INCOMPLETE and NOT DONE**.
> 
> Therefore, whenever flat commits are used during local micro-steps, the agent MUST conclude the overall task with a semantic auto-pushing commit (`gitmap cpf "<module> - <summary>"`) or execute an explicit `gitmap push` before marking the task or plan complete.

### 3.5 Documentation Updates & Locations
The flat commit specification must be propagated across:
1. `AGENTS.md` §10 under GitMap Commit Formatting.
2. CLI help documentation in `02-spec/commands/` and GitMap usage manuals.
3. Prompt templates in `01-prompts/` and developer companion skills (`.agents/skills/gitmap/`).

---

## 4. 400-Step Budget Specification

Long-running agentic tasks require a structured operational lifecycle to prevent execution degradation, context drift, and unfinished worktrees. The 400-Step Budget defines a standard budget distribution across three sequential phases.

```
+---------------------------------------------------------------------------------------------------+
|                                     400-STEP EXECUTION BUDGET                                     |
+------------------------------------+------------------------------------+-------------------------+
| Phase 1: Planning & Discovery      | Phase 2: Sanitization & Prep       | Phase 3: Implementation |
| Steps 1 .. 100                     | Steps 101 .. 250                   | Steps 251 .. 400        |
| - Requirements ingestion           | - Worktree sanitization            | - Wave 1 workers        |
| - GitMap symbol discovery          | - Stale cache purging              | - Wave 2 workers        |
| - Canonical spec authoring         | - Architecture prerequisite checks | - Targeted checks       |
| - Subtask plan decomposition       | - Test fixture preparation         | - Multi-repo sync       |
| - Brain roadmap generation         | - Baseline validation              | - Atomic release & push |
+------------------------------------+------------------------------------+-------------------------+
```

### 4.1 Phase 1: Planning, Discovery & Specification (Steps 1 .. 100)
- **Step 1 .. 20 (Requirement Ingestion & Pre-flight):** Parse verbatim user instructions, enforce git pull (`gitmap pull`), verify repository identity, load memory context (`.ai-memory/`), and output the confirmed task breakdown.
- **Step 21 .. 50 (High-Speed Codebase Discovery):** Deploy read-only exploration or discovery subagents using GitMap search commands (`gitmap aum search`, `gitmap find`, `gitmap cat`). Zero raw shell search tools allowed. Identify all affected files, symbol dependencies, and schema contracts.
- **Step 51 .. 80 (Canonical Spec Authoring):** Author domain specifications in `02-spec/21-app/<slug>/` containing verbatim user requests, data contracts, and visual standards.
- **Step 81 .. 100 (Plan Decomposition & Brain Roadmap):** Author parent plans in `.ai-memory/plans/` and bounded subtasks in `.ai-memory/plans/subtasks/<slug>/`. Initialize Brain roadmap artifacts and prepare execution ledgers.

### 4.2 Phase 2: Sanitization, Archival & Prerequisites (Steps 101 .. 250)
- **Step 101 .. 140 (Worktree Sanitization & Cache Purging):** Verify working tree state, clean orphan temporary files, purge build caches (Go build cache, npm/vite caches, python pycache) using `cg-clean-repo-build-and-caches`.
- **Step 141 .. 180 (Legacy Archival & Consolidation):** Identify deprecated plans or superseded specifications. Consolidate completed subtasks into archival indexes via `ai-memory-consolidate-reduce`.
- **Step 181 .. 220 (Prerequisite Infrastructure & Test Baseline):** Verify linter readiness, configure test runners, prepare mock datasets, and confirm dependency resolution without modifying application domain logic.
- **Step 221 .. 250 (Checkpoint Commit & Gate Review):** Record an intermediate local flat commit (`gitmap cm "phase 2 complete: sanitization and baseline verification"`) to create an immutable checkpoint before code modifications begin.

### 4.3 Phase 3: Implementation, Worker Swarms & Completion (Steps 251 .. 400)
- **Step 251 .. 310 (Worker Wave 1 - Core Refactoring):** Spawn concurrent, disjoint subagents bounded to 5–8 files each. Execute architectural migrations, icon asset modernizations, and coding guideline fixes.
- **Step 311 .. 350 (Worker Wave 2 - Secondary Polish & Integration):** Resolve secondary cross-references, update documentation, align CLI help flags, and harmonize UI design tokens.
- **Step 351 .. 375 (Targeted Verification & Quality Gates):** Execute smart incremental checks (`smart-test-runner-and-inventory`), check spec link consistency, and verify zero absolute path leaks.
- **Step 376 .. 390 (Fleet Synchronization):** If cross-repository canonical prompts or skills are affected, trigger `prompts-and-skills-sync` across connected repositories.
- **Step 391 .. 400 (Atomic Commit, Push & Release Ceremony):** Execute the mandatory semantic push (`gitmap cpf "<module> - <summary>"`) to satisfy the 'No Push = Not Done' invariant. Record session outcomes in `.ai-memory/memory/` and finalize user report.

---

## 5. Architectural Quality Gates & Invariants

1. **Strict Lowercase File Naming:** All created and modified specifications, subtasks, icons, and scripts must strictly use lowercase characters and hyphens (e.g., `01-architecture-spec.md`, `icon-workflow.svg`).
2. **Strict Relative Git Paths:** No absolute filesystem paths (`C:\...`, `/home/...`) or `file:///` URIs may appear in specifications, plans, code comments, or documentation.
3. **Boolean Hygiene:** All boolean fields and conditional statements must use positive prefixes (`is`, `has`) and implicit truth evaluations (`if isReady`). Explicit comparisons (`== true`) and mixed polarity conditions (`isA && !isB`) are forbidden.
4. **Transparent Icon Purity:** Automated asset generation must reject any SVG or PNG containing bounding background boxes or opaque dark rectangles.
5. **Upstream Completion Verification:** No task lifecycle may exit without verifying git remote push status on the tracking branch.
