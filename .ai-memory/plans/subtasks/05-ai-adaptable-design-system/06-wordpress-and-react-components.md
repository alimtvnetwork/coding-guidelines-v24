# Subtask 06: Page Expansion Rules, WordPress Migration & Modular React Components
Traceability ID: Task-08
Spec Reference: [02-spec/07-design-system/25-ai-adaptable-design-system/06-expansion-and-wordpress/01-page-expansion-and-wordpress.md](../../../../02-spec/07-design-system/25-ai-adaptable-design-system/06-expansion-and-wordpress/01-page-expansion-and-wordpress.md)
Target Files: `02-spec/07-design-system/25-ai-adaptable-design-system/06-expansion-and-wordpress/01-page-expansion-and-wordpress.md`
Action: Document page expansion consistency rules, Gutenberg `theme.json` migration schema, and modular React reference components adhering to small-function rules.
Acceptance Criteria:
- Expansion rules maintaining 70px header, alternating section backgrounds, and 16–20px card geometries.
- WordPress `theme.json` schema mapping Tier 1 color variables to FSE settings.
- Four production-grade React micro-components (`SlideButton`, `FilterChip`, `FloatingAiButton`, `HeroHeadline`), each <= 15 lines per function and <= 100 lines per file.
Targeted Verification: Verified positive boolean states (`isHovered`, `isOpen`), vertical line gaps, and zero explicit true comparisons.
