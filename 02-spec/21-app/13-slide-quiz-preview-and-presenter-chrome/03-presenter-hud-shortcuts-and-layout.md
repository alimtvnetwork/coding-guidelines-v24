# 03 — Presenter HUD, Shortcuts & Center Layout

---

## 1. HUD and camera control

Geometry and pill layout are **`31-slide-controller-buttons.md`**. Required affordances:

| Control | Label in HUD | Opens / toggles |
|:---|:---|:---|
| Camera | 📹 **Cam** | Acquires webcam per PIP spec in file 31; hard toggle **`I`** |
| Shortcuts | **?** or keyboard icon (if present) | Same data as **`/`** opener |
| Builder | ✎ **Build** | `E` — slide builder per `29` / `35` |
| Fullscreen | ⤢ **Full** | **`F`** |

Soft minimize camera (**`M`**) keeps tracks alive; **`P`** enters camera fullscreen. Do not bind bare **`1`** to slide jump when camera stage-fill is enabled—reserve **`1`** for camera stage-fill (see file 42 shortcut table).

Mount HUD via portal on `document.body`, `z-index: 60`, never inside scaled stage transform.

---

## 2. Keyboard shortcut overlay

Data model (TypeScript, Gemini-safe):

```typescript
export interface ShortcutItem {
  keys: string[];
  label: string;
}

export interface ShortcutGroup {
  group: string;
  items: ShortcutItem[];
}
```

UI rules copied from global deck implementation:

- Dialog max width **`max-w-3xl`**, max height **`85vh`**, scroll inside.
- Group title: **`10px`**, uppercase, **`tracking-[0.18em]`**, accent color.
- Each row: label left, `<kbd>` chips right (`px-1.5 py-0.5`, `text-[10px]`, mono, muted background, border).
- Global listener: **`/`** key without modifiers; ignore when focus is `INPUT`, `TEXTAREA`, or `contentEditable`.
- Callout banner at top explains **`1`** vs **`2`–`9`** slide jump and **`Ctrl/⌘+1`** outline.

Full group list is **normative in file 42 section 5**. `29-slide-navigation-and-builder.md` section 3 remains the minimal deck subset; file 42 is superset for presenter products.

---

## 3. Center-stage content (`1920×1080`)

Use a single flex wrapper on the virtual canvas:

```css
.slide-center-stage {
  box-sizing: border-box;
  width: 1920px;
  height: 1080px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 96px 120px;
  text-align: center;
  gap: 24px;
}
```

Typography:

| Role | Size | Weight | Shadow |
|:---|:---|:---|:---|
| Kicker | `clamp(14px, 1.1vw, 18px)` | 600 | `--text-shadow-rest` |
| Headline | `clamp(40px, 4.2vw, 72px)` | 700 | `--text-shadow-rest` |
| Subcopy | `clamp(18px, 1.6vw, 28px)` | 400 | none |
| Max line width | `min(920px, 100%)` | — | — |

For left-aligned quiz question + option stack, switch to `align-items: stretch` and `text-align: left` on the question block only; keep progress/timer in HUD.

---

## 4. Responsiveness

| Breakpoint | Rule |
|:---|:---|
| Stage | Always design at **1920×1080**; scale with `transform: scale()` from `24-slide-presentation-system.md` |
| Option cards | `p-4`, `rounded-2xl`, gap `12px`; badge `w-9 h-9` below `640px`, `w-10 h-10` at `sm+` |
| Touch | Minimum hit target **44×44px** including badge and trailing icon |
| Safe area | HUD offsets use `env(safe-area-inset-*)` per file 31 |

Quiz sidebar + hero layouts follow exam spec `02-spec/21-app/` sibling products; when embedding in slides, collapse sidebar below **`1024px`** viewport width.

---

## 5. Builder mode interaction

- **`E`** toggles `useEditStore.isEditMode` (file 35). Audience `/present` route MUST NOT load edit store.
- Builder hotkeys **`B`**, **`Tab`**, **`1`–`7`** layer focus, **`Cmd/Ctrl+Z`** undo—see file 35; they MUST NOT fire while shortcut dialog is open.
- Persist deck mutations only through `useDeckStore` gateway.

---

## 6. Bright-gold deck alignment

Corporate gold deck colors live in **`05-bright-gold-tech/`** and **`40-theme-switch.md`** (`bright-gold-tech`). Quiz **`botanical-light`** is a **light runner skin**, not a replacement for bright-gold slide themes. A deck may use bright-gold slides and botanical-light quiz embed via CSS scope:

```html
<div data-theme="bright-gold-tech" class="slide-root">
  <div data-theme="botanical-light" class="quiz-embed">...</div>
</div>
```

Nested theme scopes MUST set both HSL and elevation tokens on the inner root.
