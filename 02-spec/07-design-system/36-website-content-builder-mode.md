# 36 — Website Content Builder Mode

> **/goal** Point at the authoritative website visual builder contract.
> **/learn** Do not implement a second builder from this file. The authoritative specification is `26-visual-builder.md`.

**Version:** 4.3.0
**Updated:** 2026-10-02
**Status:** Pointer

---

The canonical website visual builder specification is [26-visual-builder.md](./26-visual-builder.md).

This file does not define a separate gate, an alternate id scheme, a second store, or a divergent export. Follow [`26-visual-builder.md`](./26-visual-builder.md) for all website builder overlay implementations:
- Gate query: `?builder=1&email={OWNER_EMAIL}`
- Modes: Text, Images, Menu, Layout, Off
- Contenteditable inline text editing with caret management
- Single-click image and icon replacement
- Menu link editing and group reordering
- Review-only client storage and deterministic ZIP export (`content-changes--all-pages--YYYY-MM-DD-HHmm.zip`)
