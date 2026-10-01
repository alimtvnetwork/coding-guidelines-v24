# Follow the UI and UX design system

> **Prompt Version:** 1.0.0

Before you draw a website, a blog, a menu, or a CSS animation, read these files in order and follow them. Relative paths are from the repository root.

1. `02-spec/07-design-system/04-white-blue-theme/readme.md`
2. `02-spec/07-design-system/04-white-blue-theme/01-colors-typography-and-tokens.md`
3. `02-spec/07-design-system/04-white-blue-theme/02-header-mega-menu-and-footer.md`
4. `02-spec/07-design-system/04-white-blue-theme/03-buttons-motion-and-interactions.md`
5. `02-spec/07-design-system/04-white-blue-theme/04-cards-heroes-and-section-library.md`
6. `02-spec/07-design-system/25-page-assembly.md`

When the task is the on-page editor, also read `02-spec/07-design-system/26-visual-builder.md`. Do not use that file for slides.

## Menu (do not retune)

From `02-header-mega-menu-and-footer.md`:

- Header height `72px`, glass `bg-background/90`, `backdrop-blur-xl`.
- Scroll threshold `12px`. Border and shadow ease over `420ms` with `cubic-bezier(0.16, 1, 0.3, 1)`.
- `SlideSwapLabel` stagger `0.04` on nav labels. The primitive default `0.018` is not the nav value.
- Underline grows `scale-x-0` to `scale-x-100` over `--dur-fast`.
- Mega menu safe region `pad = 14`. Close delay `110ms` to `220ms`. Chevron rotates `180deg`.
- Promo card flips `rotateY(180deg)` over `820ms` at `perspective: 1400px`.

## Pages

Compose only the shell and section ids in `25-page-assembly.md`. The H1 is at most 9 words (`20-ai-training-and-checklist-guide.md`). Do not add another word cap.

Marketing colors stay navy `#0D2975`, cobalt `#2563EB`, violet `#822EE8`. Do not put slide amber (`#ffae00`, `#F5A623`, `#FFD83A`) on a page.

## Refusal

If a number is not in the files above, stop and leave it out. Do not name a client, a vendor, or a private repository.
