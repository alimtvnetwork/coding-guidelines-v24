# Write and Enhance a Design Spec That a Blind AI Can Follow

> **Prompt Version:** 1.0.0
> **Trigger keywords:** `design-spec`, `enhance-design-spec`, `ui-spec`, `slide-spec`, `blind-spec`

**/goal** Turn source designs (code, specs, decks, screenshots) into public design-system specs so exact that an AI with no other context can build the same website, blog, menu, button, slide deck, or image without guessing a single value.

**/learn** A spec is finished only when every component in the source has its values written down, every gap is marked, and the completion gate at the end passes. A file that only points at another file does not count as covered.

---

## 1. Inputs

Fill these before you start. If one is missing, ask once, then stop.

| Input | Meaning | Default |
|---|---|---|
| `SOURCE` | Folders or files to learn from. Read-only | none |
| `TARGET` | Spec folder to write | `02-spec/07-design-system/` |
| `MODE` | `create` (new topic) or `enhance` (fix and complete existing files) | `enhance` |
| `DOMAINS` | Any of: website, blog, menu, button, color, type, spacing, motion, sections, slide deck, slide builder, site builder, image | all that the source contains |
| `BANNED_NAMES` | Product, client, company, person, domain, and private repo names that must never appear in output | none, ask |

---

## 2. Hard rules (any violation is an auto-reject)

1. **No invented values.** Every hex, HSL, px, ms, easing, count, key, limit, and default must come from `SOURCE` or from an existing file in `TARGET`. If the source does not state it, write `Not specified in source. Do not invent.` Never write "about", "around", "roughly", or "~" in place of a number unless the source itself says it.
2. **Keep a private ledger.** While working, record every value and the source file it came from in `.ai-memory/temp/design-spec-ledger.md`. The public spec never cites the private source path.
3. **Anonymize everything.** No entry from `BANNED_NAMES` in any file, heading, example, comment, storage key, logo filename, or sample copy. Use generic ids and placeholders such as `{SITE_DOMAIN}`, `{OWNER_EMAIL}`, `{BRAND}`. Typeface names (Ubuntu, Poppins, Inter) and open-source library names (Lucide, Radix) are allowed.
4. **Enhance, do not rewrite.** Keep every existing correct value. Change a value only when the source contradicts it, and record the old and new value in the ledger.
5. **Conflicts are scoped, not averaged.** When two sources disagree (for example `28×8` versus `32×8` for an active dot), write both, say which surface each belongs to, and name one canonical value for new work.
6. **One palette per surface.** Marketing pages, slide decks, and images each get their own palette. Never borrow a slide accent for a page or a page accent for a slide.
7. **Closed sets.** Every list a builder picks from (section ids, layout ids, button variants, sizes, theme ids, pill colors, transitions, keys, builder modes) is written as a closed list followed by: anything else does not exist.
8. **Values beside the rule.** Do not write "use the brand color". Write the token, the hex, and the HSL in the same row.
9. **Relative paths only.** No absolute paths, no `file:///`.
10. **Files stay small.** At most 300 lines per file. Split by topic. Lowercase `NN-kebab-name.md` with the next free number in `TARGET`.
11. **Plain sentences.** Short sentences, tables for values, no slogans.
12. **No false "done".** You may not say the work is complete until every box in section 8 is checked. If something is still missing, say what it is.

---

## 3. Workflow

### Phase 1: Inventory (read only, no writing)

1. List every file in `SOURCE`. Skip `node_modules`, build output, `.git`, lockfiles, and generated skills.
2. Read token files, style sheets, component files, type definitions, and existing specs.
3. Build an inventory table of every item a builder could need:

| Kind | Examples to look for |
|---|---|
| Tokens | colors, gradients, shadows, radii, spacing, breakpoints, durations, easings, z-index |
| Type | families, weights, size scale, line-height, tracking |
| Components | buttons and every variant, header, nav link, mega panel, promo card, drawer, footer, cards, badges, pills, tooltips, toasts |
| Sections | hero kinds, grids, rails, tabs, pricing, FAQ, forms, closers |
| Motion | each keyframe, entrance, hover, stagger, reduced-motion behavior |
| Slides | canvas, themes, palette, layouts and their fields, transitions, controller, dots, keys, step reveal, builder, export |
| Builders | gate, modes, limits, storage, history, export |
| Images | canvas sizes, export sizes, safe zones, text rules, file naming |

### Phase 2: Gap audit

For each inventory row, find the matching file and section in `TARGET`, then mark one status:

| Status | Meaning |
|---|---|
| `exact` | Every value is written and matches the source |
| `wrong` | A value differs from the source |
| `pointer-only` | The target only links elsewhere, and the linked file lacks the values |
| `missing` | Nothing in `TARGET` |

Write this coverage table in the ledger. Everything that is not `exact` becomes a task.

### Phase 3: Write

Fix `wrong` rows first, then fill `pointer-only` and `missing` rows. Use the template in section 6. Put one topic in each file.

### Phase 4: Wire up

1. Add every new file to the `TARGET` readme reading checklist and file inventory table.
2. Add every new file to the consistency report inventory and cross-reference table.
3. Link sibling files in both directions (for example, slide layouts link to slide color options, and the reverse).
4. Update the prompts that consume the specs (`01-prompts/10-ui-and-design/07-follow-ui-ux-design-system.md`, `08-create-slide-deck.md`) so their reading lists include the new files.

### Phase 5: Verify

Run the checks in section 7, then tick section 8.

---

## 4. What "complete" means for each domain

A domain is complete only when every bullet has real values.

**Color:** every token name, hex, HSL (plus RGB and OKLCH when the family already uses them), its role, and where it is banned. Surfaces, gradients, focus ring, hover and active tints, and the light or dark rule.

**Type:** families and weights, the full size scale with line-height and tracking per role, which color each role uses, and how fonts load.

**Spacing and layout:** container width, gutters per breakpoint, section rhythm, radii, shadows, safe area, and z-index layers.

**Buttons:** every variant crossed with every size: height, padding, font size and weight, radius, icon size, gap. States: rest, hover, focus-visible, active, disabled. Motion and its duration. Accessible name.

**Menu and header:** height, sticky rule, scroll threshold and what changes after it, nav link type, underline, chevron, panel geometry, column templates, entrance timings and stagger, close delay and safe region, promo card faces and flip, CTA pair, mobile drawer, keyboard and `Escape` behavior.

**Motion:** duration and easing tokens, each keyframe with from and to values, stagger, reduced-motion fallback, and a per-page budget.

**Sections and pages:** closed section id list, anatomy of each, band tone, copy caps, and page shells for marketing, blog index, and article.

**Slides:** canvas and scale formula, theme id table, palette, type scale, shell layers, safe area, closed layout union with every field and default, transitions with ms and easing, controller buttons, dots, tooltip, keys, step reveal, builder, and export.

**Site builder:** access gate, modes, size and type limits, sanitizer, storage, history, and export format.

**Images (logo, thumbnail, banner):** canvas and export pixel sizes, safe and collision zones, text size and word limits, palette, and file naming. Reuse the existing image prompts in `01-prompts/10-ui-and-design/` (`01` to `06`) instead of duplicating them.

---

## 5. Component entry fields

Every component section has these rows. If the source is silent, write `Not specified in source.`

| Field | Example of the level of detail required |
|---|---|
| Name and job | `MegaLink`: one link inside a mega panel group |
| Anatomy | label, optional description, left rule, trailing arrow |
| Geometry | `px-3 py-2`, radius `10px`, arrow `14px` |
| Tokens | hover wash `color-mix(in oklab, var(--primary) 7%, transparent)` |
| States | rest, hover, focus-visible, active, disabled |
| Motion | rule `scaleY(0)` to `scaleY(1)` over `420ms`, `cubic-bezier(0.16, 1, 0.3, 1)` |
| Reduced motion | opacity only, `120ms` |
| Accessibility | role, accessible name, keyboard |
| Closed options | the variants or sizes allowed |
| Do not | the mistakes an AI is likely to make |

---

## 6. Spec file template

```markdown
# NN — Title

> **/goal** One sentence: what this file lets an AI build.
> **/learn** One sentence: what lives here and what lives in a sibling file.

**Version:** 1.0.0
**Status:** Active

---

## 0. Anti-hallucination

If a value is not in this file or a file it names, do not invent it.

## 1. Scope

When to use this file and when not to.

## 2. Tokens

| Token | Hex | HSL | Role |

## 3. Components

One section per component, using the entry fields.

## 4. Closed sets

The allowed ids, then: anything else does not exist.

## 5. Refusal list

What an AI must not do.
```

---

## 7. Verification commands

Run from the repository root. Each check must return nothing, except the line count.

```bash
# Banned names (replace with BANNED_NAMES, joined by |)
rg -i "name-one|name-two" 02-spec/07-design-system 01-prompts/10-ui-and-design

# Absolute paths and file URIs
rg -n "[A-Za-z]:\\\\|file:///|/home/|/Users/" 02-spec/07-design-system

# Vague numbers
rg -n "\b(about|around|roughly|approx\.?)\s+\d|~\s?\d" 02-spec/07-design-system

# Line counts, each must be 300 or fewer
wc -l 02-spec/07-design-system/*.md
```

PowerShell line count:

```powershell
Get-ChildItem 02-spec/07-design-system/*.md | ForEach-Object { "{0} {1}" -f (Get-Content $_).Count, $_.Name }
```

Then open the readme and the consistency report and confirm every new file appears in both.

---

## 8. Completion gate

Do not report done until every box is checked.

- [ ] Every inventory row is `exact`, or marked `Not specified in source` with a reason.
- [ ] No `pointer-only` row remains.
- [ ] Every number, color, and duration in the new text is in the ledger with its source.
- [ ] Conflicting values are written with their scope and one canonical choice.
- [ ] Each closed set ends with: anything else does not exist.
- [ ] Each component has all entry fields from section 5.
- [ ] Marketing, slide, and image palettes are separate.
- [ ] The banned-name check returns nothing.
- [ ] The absolute-path check returns nothing.
- [ ] The vague-number check returns nothing, or each hit is quoted from the source.
- [ ] Every file is 300 lines or fewer and named `NN-lowercase-kebab.md`.
- [ ] Every new file is in the readme and the consistency report.
- [ ] The consuming prompts list the new files.
- [ ] A test read passes: pick one component and one slide layout, and confirm you could build each from the spec alone without opening `SOURCE`.

---

## 9. Final report

End with exactly this:

```markdown
### Design spec result

- Mode: create | enhance
- Inventory items: N
- Exact before / after: N / N
- Wrong fixed: N
- Pointer-only filled: N
- Missing added: N
- Not specified in source: N (list them)

### Files

- relative/path/one.md (created | updated)

### Gate

All boxes in section 8 checked: yes | no (list the open boxes)
```

If the gate is `no`, the work is not done. Say so.
