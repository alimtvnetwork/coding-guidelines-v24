# Coding Guideline and Design System Spec Improvements

> **/goal** Train any AI, including a weak one, to produce professional slides and a compact tree-plus-detail panel in one shot.
> **/learn** Follow this folder in order. Do not mix token files. Do not revive archived slide skills.

**Slug:** `coding-guideline-and-design-system-spec-improvements`
**Status:** Active
**Updated:** 28-Oct-26
**Sequence:** 35

## Read order

| Order | File | Job |
| :--- | :--- | :--- |
| 1 | [01-architecture-and-token-contract.md](01-architecture-and-token-contract.md) | Which file wins for slides, websites, and trees |
| 2 | [02-one-hundred-design-improvements.md](02-one-hundred-design-improvements.md) | 100 checks. The next run appends 101 onward |
| 3 | [03-tree-view-and-detail-panel.md](03-tree-view-and-detail-panel.md) | Generic compact tree and right panel |
| 4 | [04-blind-ai-execution-checklist.md](04-blind-ai-execution-checklist.md) | Non-negotiable shot checklist |
| 5 | [05-benchmark-previous-versus-new.md](05-benchmark-previous-versus-new.md) | What the old specs let a weak model get wrong |

## Authority

- Slide stage default: `bright-gold-tech` in `02-spec/07-design-system/40-theme-switch.md` and `02-spec/07-design-system/05-bright-gold-tech/`.
- Repeatable prompt: `01-prompts/30-improvements/08-design-system-spec-improvements.md`.
- Plan: `.ai-memory/plans/pending/35-coding-guideline-and-design-system-spec-improvements.md`.

## Reduced skills (do not restore)

Archived, and not the path for new decks:

- `.agents/skills-archived/create-slide-deck/`
- `.agents/skills-archived/slides-deck-management/`
- `.agents/skills-archived/theming-guidelines/`
- `.agents/skills-archived/update-theming/`

Active authoring skills stay `follow-ui-ux-design-system` and `write-and-enhance-design-spec`.

## Sibling evidence (read-only)

- `presentations-repos/slides-spec/` and `presentations-repos/flat-slide-show/` carry the same design-system tree.
- `wp-exam/src/styles/theme.less` is the source for sideways choice-row motion.
- `antigravity-manager/src/components/instances/PromptTreeView/` is the source for the compact tree.
- `letsmarknow` and `letsmarknow-ui` had no UI source on 28-Oct-26. Do not invent a Let's Mark skin.
