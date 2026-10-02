# Write and Enhance a Design Spec That a Blind AI Can Follow

> **Prompt Version:** 2.0.0
> **Trigger keywords:** `design-spec`, `enhance-design-spec`, `ui-spec`, `slide-spec`, `blind-spec`, `spec-creator`, `spec-enhancer`

**/goal** Autonomously ingest source designs (codebases, specifications, presentation decks, UI components, screenshots) and author or enhance public design-system specifications so comprehensive, mathematically exact, and rigorous that any "blind AI" or human engineer with zero external context can follow them blindly to design websites, presentations, blogs, animations, menus, buttons, builder modes, or images without guessing a single value or hallucinating.

**/learn** A specification is finished ONLY when every component, token, physics parameter, layout slot, and interaction timing in the source has its exact values written down in structured tables, every gap is audited, and the completion gate at the end passes 100%. A file that merely points at another file or leaves vague placeholders does not count as complete.

---

## 1. Required Execution Inputs

Before beginning analysis or writing, establish these parameters:

| Input Parameter | Definition & Constraints | Default Value |
|:---|:---|:---|
| `SOURCE` | Read-only folder paths or files containing reference implementations (code, specs, decks). | User-provided or repository inspection |
| `TARGET` | Spec directory where public design system files are created or enhanced. | `02-spec/07-design-system/` |
| `MODE` | `create` (authoring new specifications) or `enhance` (fixing and expanding existing specs). | `enhance` |
| `DOMAINS` | Any of: `website`, `blog`, `menu`, `button`, `color`, `type`, `spacing`, `motion`, `sections`, `slide-deck`, `slide-builder`, `site-builder`, `image`. | All domains present in source |
| `BANNED_NAMES` | Client, product, company, private repository, and individual person names that MUST NEVER appear in public specifications. | Mandatory anonymization |

---

## 2. Hard Architectural Rules (Zero-Tolerance Violations)

1. **Total Ban on Invented Values:** Every hex, HSL, RGB, OKLCH, pixel, millisecond, cubic-bezier easing, character stagger rate, keycode, and default limit MUST originate from `SOURCE` or grounded mathematics. If the source is silent, explicitly state: `Not specified in source. Do not invent.` Never write "about", "around", "roughly", or "~" in place of a concrete number.
2. **Strict Anonymization Mandate:** Never leak names from `BANNED_NAMES` (no client company names, internal repositories, or personal names) into public specifications, comments, sample text, or image paths. Always use generic enterprise designations (e.g., `{BRAND}`, `{COMPANY}`, `Executive Leadership`, `Enterprise Platform`).
3. **Pure DOM Text Rendering Mandate (Zero Baked-In Text):** In websites, presentations, and UI designs, all headlines, body copy, bullets, metrics, credentials, and captions MUST be rendered as live, selectable DOM HTML elements (`<h1>`, `<p>`, `<span>`) styled with CSS typography tokens. Raster images (`.png`, `.jpg`, `.webp`) are strictly reserved for photographic visual plates, author avatars, and partner logos.
4. **Strict Relative Git Paths Only:** All file paths, markdown links, subtask paths, and citations MUST be strictly relative paths starting from the repository root (e.g., `02-spec/07-design-system/11-button-system.md`). TOTAL BAN on absolute filesystem paths (`C:\...`, `/home/...`) and `file:///` URIs.
5. **Lowercase File Naming Convention:** All generated specification files MUST use strictly lowercase kebab-case naming (e.g., `34-slide-layout-catalog.md`). No uppercase characters, spaces, or underscores in filenames.
6. **Strict Boolean Standards:** Positive booleans MUST ALWAYS be evaluated implicitly (`if isReady`). NEVER evaluate explicitly against true (`if isReady == true` is banned). NEVER combine positive and negative checks in the same condition.
7. **Scoped Domain Palettes:** Marketing websites, presentation decks, and branding graphics each maintain distinct, dedicated color palettes. Never contaminate a marketing page with slide presentation amber or vice-versa.
8. **Closed Sets Rule:** Every selectable options list (button variants, sizes, slide layout types, theme IDs, pill preset colors, transition kinds, hotkeys, builder modes) MUST be specified as a closed set followed by: `Anything else does not exist.`
9. **Values Beside the Rule:** Never simply write "use the primary color". Document the semantic token, HEX value, RGB coordinates, and HSL values in the same row.
10. **Strict File Size Cap:** Specification files MUST remain bounded (at most 300 lines per file). Complex topics must be broken into discrete, single-responsibility files.

---

## 3. Systematic 5-Phase Workflow

### Phase 1: Comprehensive Source Inventory (Read-Only)
1. Recursively scan `SOURCE` directories, skipping `node_modules`, build artifacts, `.git`, and lockfiles.
2. Read token files, stylesheets, component primitives, hooks, route templates, and existing specs.
3. Construct an exhaustive inventory table recording:
   - **Tokens:** Colors, 10-step gradients, shadows, radii, spacing, gutters, breakpoints, durations, easings.
   - **Typography:** Font families, weights, fluid clamp scales, line heights, tracking, tabular numeral rules.
   - **Components:** Button variants, sizes, sticky glass header, nav links, mega panel, 3D flip card, mobile drawer, footer.
   - **Sections:** Hero layouts, feature grids, drag rails, process stacks, pricing tables, FAQ accordions.
   - **Motion:** Keyframe definitions, reveal thresholds, stagger intervals, reduced-motion fallbacks.
   - **Slides:** 16:9 canvas (`1920×1080`), scale math, 10-step ramps ($S_0$–$S_9$), 10 master layouts, controller HUD, dots, hotkeys, step reveals.
   - **Builders:** Slide builder (`useDeckStore` and `useEditStore`) and the website builder in `26-visual-builder.md` only. `36-website-content-builder-mode.md` is a pointer.
   - **Content extracts:** A folder of page-copy changes is content only. Do not take colors, type, or motion from it.
   - **Images:** Canvas dimensions, safe zones, text caps, aspect ratios.

### Phase 2: Gap & Fidelity Audit
Compare every inventory item against existing specifications in `TARGET`:
- `exact`: Value is explicitly documented and matches source ground truth.
- `wrong`: Value conflicts with source implementation.
- `pointer-only`: Specification merely links to another file without providing the concrete numbers.
- `missing`: Component, token, or behavior is absent from specifications.

### Phase 3: Surgical Specification Authoring & Enhancement
1. Remediate all `wrong` values first, aligning them with ground truth.
2. Author dedicated specification files for all `missing` and `pointer-only` areas.
3. Follow the standardized Spec File Template (Section 5) for every file.

### Phase 4: Bi-Directional Cross-Referencing & Prompt Sync
1. Add every new specification file to the reading checklist in `02-spec/07-design-system/readme.md`.
2. Add every new file to the module inventory table in `readme.md`.
3. Update consuming prompts (`01-prompts/10-ui-and-design/07-follow-ui-ux-design-system.md` and `08-create-slide-deck.md`) so their reading sequences reference all new files.

### Phase 5: Verification & Completion Gate
Execute the automated lint checks and review against the Completion Gate (Section 7).

---

## 4. Domain Field Checklist

This section names the fields to fill. It does not supply the values. The source wins. If the source is silent, write `Not specified in source. Do not invent.` If the slide source and the website source disagree, write two palettes and name the domain of each. Do not merge them. Do not copy a number from this prompt into a spec.

### 4.1 Color
- Token name, and HEX, RGB, and HSL only when the source states them.
- A gradient ramp only when the source states the stops.
- Pill names only when the source lists them.

### 4.2 Type
- Family, weight, size, line height, and tracking, each only when the source states it.
- A fluid `clamp()` scale only when the source states the formula.

### 4.3 Buttons
- The variant names the source implements.
- The sizes the source implements.
- Motion only when the source states duration, easing, or strength.

### 4.4 Navigation
- Header height, scroll threshold, panel motion, and drawer behavior, each only when the source states it.

### 4.5 Slides
- Canvas size from the slide store.
- One row per component that exists, with the type string that component's renderer uses.
- A catalog type with no component is marked missing. It is not filled with a sample layout.

### 4.6 Slide builder
- Store names and which one is persisted, from the slide app.
- Hotkeys only when a shortcut file binds them.

### 4.7 Website builder
- The only contract is `02-spec/07-design-system/26-visual-builder.md`.
- Fields: gate query, id rule, sanitizer allow-list, image limit, save delays, store version, export names.
- Do not add `data-edit-id`, `?builder=true`, or `[[accent]]` unless that file states them.

---

## 5. Standardized Spec File Template

Every newly created or enhanced specification file MUST follow this structure:

```markdown
# NN — [Topic Title]

> **/goal** One sentence defining what an AI or human engineer can build using this specification.
> **/learn** One sentence summarizing key tokens, formulas, and references to sibling specifications.

**Version:** 4.3.0
**Status:** Active
**AI Confidence:** High only for numbers this file states. Unstated numbers are not invented.
**Ambiguity:** State it. Do not write None while another file disagrees.

---

## 1. System Overview & Scope
[When to use this specification, architectural pillars, and anti-hallucination boundaries]

---

## 2. Token Registry & Exact Mathematics
[Tables of tokens with Hex, HSL, RGB, OKLCH, dimensions, and mathematical formulas]

---

## 3. Component Taxonomy & Architectural Blueprint
[Component anatomy, geometric dimensions, states, motion curves, and TypeScript interfaces]

---

## 4. Closed Sets & Valid Options
[The complete enumerated list of allowed IDs, variants, or modes, followed by:]
*Anything else does not exist.*

---

## 5. Reference Implementation
[Tested, framework-agnostic code snippet or component blueprint]

---

## 6. Anti-Hallucination & Quality Verification Checklist
[Exhaustive checkboxes guaranteeing that every required number, property, and rule is satisfied]
```

---

## 6. Automated Verification Commands

Run these checks from the repository root:

```powershell
# 1. Verify no absolute paths or file:/// URIs
Get-ChildItem -Path "02-spec/07-design-system", "01-prompts/10-ui-and-design" -Filter "*.md" | Select-String -Pattern "([A-Za-z]:\\|file:///|/home/|/Users/)"

# 2. Verify no vague numbers (about, roughly, approx)
Get-ChildItem -Path "02-spec/07-design-system", "01-prompts/10-ui-and-design" -Filter "*.md" | Select-String -Pattern "\b(about|around|roughly|approx\.?)\s+\d|~\s?\d"

# 3. Check line counts (must be <= 300 lines per file)
Get-ChildItem -Path "02-spec/07-design-system/*.md" | ForEach-Object { "{0,4} lines: {1}" -f (Get-Content $_.FullName).Count, $_.Name }
```

---

## 7. Completion Gate

An AI agent MUST NOT report completion until every gate is verified:

- [ ] Every inventory row from the source has been mapped to an exact specification with concrete numbers.
- [ ] NO `pointer-only` stubs remain; all referenced values are fully declared.
- [ ] Zero baked-in text: Pure DOM typography mandate is enforced across websites and slide decks.
- [ ] Navigation, buttons, and slides contain only numbers that appear in the source. A missing field says `Not specified in source. Do not invent.`
- [ ] Slide types match renderer `case` labels. A type with no component is not given sample pixels.
- [ ] The website builder matches `26-visual-builder.md`. It does not add `data-edit-id` or `?builder=true`.
- [ ] Page-copy extracts contributed no color, type, or motion value.
- [ ] No banned or private company names exist in public specification files.
- [ ] All file paths are strictly relative paths from the git root.
- [ ] All new files are indexed in `02-spec/07-design-system/readme.md`.
- [ ] Consuming prompts (`07-follow-ui-ux-design-system.md` and `08-create-slide-deck.md`) list all new specification files in their reading sequences.
