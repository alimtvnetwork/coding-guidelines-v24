# 35 — Slide Builder Mode, Interactive Canvas & Inspector Specification

> **/goal** Master and enforce the dual-store architecture, 7 visual layers, interactive selection overlays, bounding box overrides, and audio cue debouncing of the Slide Presentation Builder Engine.
> **/learn** Master the separation of `useDeckStore` (persisted) and `useEditStore` (ephemeral), the 7 canvas stacking layers, builder hotkeys (`B`/`E`, `Tab`, `Cmd+Z`, `1`–`4`), audio debouncing windows, and headless Chromium print-ready PDF exports.

**Version:** 4.0.0
**Status:** Active
**AI Confidence:** High
**Ambiguity:** None

---

## 1. System Overview & Architectural Role

The **Slide Builder Engine** empowers authors, presenters, and AI agents to visually configure presentations in real-time directly on the scaled `1920×1080` virtual canvas:
- Move, drag, and resize text blocks, bullet cards, and media plates.
- Switch layout models and theme palettes on the fly.
- Reassign pill preset colors and 9-cell alignment coordinates.
- Maintain a non-destructive undo/redo history stack without polluting presentation playback timers.

---

## 2. Decoupled Dual-Store Architecture

To eliminate unnecessary re-renders and guarantee that editing artifacts never leak into audience presentation view, state is decoupled into two discrete stores:

```
┌─────────────────────────────────┐     ┌─────────────────────────────────┐
│   useDeckStore (Persisted)      │     │    useEditStore (Ephemeral)     │
│ ─────────────────────────────── │     │ ─────────────────────────────── │
│ • deck: DeckData                │     │ • isEditMode: boolean           │
│ • activeSlideIndex: number      │     │ • selectedElementId: string     │
│ • upsertSlide(slide)            │     │ • activePanel: PanelType        │
│ • updateSlideTheme(themeId)     │     │ • undoStack: HistoryAction[]    │
│ • reorderSlides(from, to)       │     │ • redoStack: HistoryAction[]    │
│ • LocalStorage Sync: "deck-v1"  │     │ • activeBoxHover: string | null │
└─────────────────────────────────┘     └─────────────────────────────────┘
                 ▲                                       ▲
                 └───────────────────┬───────────────────┘
                                     │
                             applyEdit() Gateway
                                     │
                        ┌────────────────────────┐
                        │ Single Mutator Gateway │
                        └────────────────────────┘
```

1. **`useDeckStore` (Persisted):** Holds the authoritative presentation JSON tree. Persists changes to `localStorage` under a versioned key (`deck-v1`). Synchronizes cross-window edits via browser `storage` events.
2. **`useEditStore` (Ephemeral):** Manages interactive UI states, selected element IDs, drag coordinates, undo/redo stacks, and toolbars. Cleared on page refresh.

---

## 3. Visual Canvas Layer Stack (7 Discrete Layers)

Every element on the `1920×1080` canvas is assigned to one of 7 isolated stacking layers (`z-index` and CSS `isolation`):

```
┌─────────────────────────────────────────────────────────────────────────────┐
│ Layer 6: Floating Inspector & HUD (Controls, Theme Bar, Builder Sidebar)   │
├─────────────────────────────────────────────────────────────────────────────┤
│ Layer 5: Builder Selection Overlays (Bounding boxes, blue focus outlines)  │
├─────────────────────────────────────────────────────────────────────────────┤
│ Layer 4: Ink Annotation Layer (Live presenter drawing & highlighter paths) │
├─────────────────────────────────────────────────────────────────────────────┤
│ Layer 3: Live DOM Typography & Cards (Headings, bullet lists, pricing)     │
├─────────────────────────────────────────────────────────────────────────────┤
│ Layer 2: Media Plates (Photographic hero silhouettes, feathered masks)     │
├─────────────────────────────────────────────────────────────────────────────┤
│ Layer 1: Brand Watermarks & Ribbons (Concentric arcs, organic SVG waves)   │
├─────────────────────────────────────────────────────────────────────────────┤
│ Layer 0: Canvas Base Background (Pure white #FFFFFF, 10-step gradient stop)│
└─────────────────────────────────────────────────────────────────────────────┘
```

| Layer Index | Name | CSS Stacking | Architectural Scope |
|:---:|:---|:---|:---|
| **Layer 0** | Canvas Base | `z-index: 0` | Pure background fill or base theme gradient ($S_0$–$S_2$). |
| **Layer 1** | Watermarks | `z-index: 10` | Non-interactive SVG organic waves, concentric circles. |
| **Layer 2** | Media Plates | `z-index: 20` | Hero photos, team avatars with feathered gradient masks. |
| **Layer 3** | DOM Typography | `z-index: 30` | Live HTML text, bullet clusters, pricing cards, pills. |
| **Layer 4** | Ink Annotations| `z-index: 40` | Real-time freehand SVG pen/marker canvas paths. |
| **Layer 5** | Selection Box | `z-index: 50` | 2px solid cyan/blue bounding box, corner resize handles. |
| **Layer 6** | Inspector HUD | `z-index: 60` | Fixed floating toolbars, slide reordering list, modal menus. |

---

## 4. Key Actions & Hotkey Navigation Matrix

| Hotkey | Action Payload | Context Guard |
|:---|:---|:---|
| **`B`** or **`E`** | Toggle Builder Mode ON / OFF | Suppressed when typing in text fields |
| **`Tab`** | Focus next editable element on slide | Active when builder mode is ON |
| **`Shift + Tab`** | Focus previous editable element | Active when builder mode is ON |
| **`Cmd/Ctrl + Z`** | Undo last canvas modification | Pops from `undoStack`, pushes to `redoStack` |
| **`Cmd/Ctrl + Shift + Z`** | Redo canvas modification | Pops from `redoStack`, pushes to `undoStack` |
| **`Escape`** | Deselect active element / Close modal | Blurs active focus ring |
| **`1`** | Quick-switch to Royal Violet theme | Applies to current slide or deck |
| **`2`** | Quick-switch to Corporate Gold theme | Applies to current slide or deck |
| **`3`** | Quick-switch to Enterprise Blue theme | Applies to current slide or deck |
| **`4`** | Quick-switch to Clinical Emerald theme | Applies to current slide or deck |
| **`F`** | Toggle Fullscreen Mode | Invokes `requestFullscreen()` |

---

## 5. Element Identity & Custom Bounding Boxes

Each editable canvas element is registered with a unique key:
```typescript
export interface EditBox {
  x: number;       // Left offset in 1920px canvas space
  y: number;       // Top offset in 1080px canvas space
  width?: number;  // Explicit pixel width override
  height?: number; // Explicit pixel height override
}

export interface SlideData {
  id: string;
  type: SlideType;
  title: string;
  boxes?: Record<string, EditBox>; // e.g., { "headline": { x: 140, y: 240, width: 1200 } }
}
```

When builder mode is active:
- Element renders a `SelectionOverlay` with a `2px solid #38bdf8` outline.
- Four corner drag handles (`size: 8×8px`, background `#FFFFFF`, border `#0284c7`).
- Live coordinates display in a floating micro-tooltip (`x: 140px, y: 240px`).

---

## 6. Acoustic & Audio Cue Engine

Tactile acoustic cues trigger dynamically during live presentations:

| Audio Event | Asset File | Debounce Window | Default Volume |
|:---|:---|:---:|:---:|
| **Slide Transition Swoosh** | `/sounds/fade_swoosh_v4.mp3` | `120ms` | `0.90 × Master` |
| **Sub-Step Advance Click** | `/sounds/click.mp3` | `80ms` | `0.70 × Master` |
| **Typewriter Character Tap** | `/sounds/tap.mp3` | `45ms` | `0.35 × Master` |

---

## 7. Ultra-High Resolution Headless PDF Export

To generate pixel-perfect, print-ready PDF handouts:
1. Load deck in headless Chromium (`puppeteer` / `playwright`).
2. Lock viewport dimensions to exact `1920 × 1080`.
3. Disable all animations via URL query parameter `?export=pdf&reducedMotion=true`.
4. Iterate slides 1 through $N$, executing page capture:
   ```javascript
   await page.pdf({
     path: 'deck-handout-print.pdf',
     width: '1920px',
     height: '1080px',
     printBackground: true,
     margin: { top: 0, right: 0, bottom: 0, left: 0 }
   });
   ```

---

## 8. Anti-Hallucination & Quality Verification Checklist

- [ ] State architecture strictly maintains dual-store separation (`useDeckStore` vs `useEditStore`).
- [ ] Stacking context adheres strictly to the 7 defined visual layers.
- [ ] Hotkeys `B` and `E` toggle builder mode; `Escape` clears selection.
- [ ] Audio cue triggers adhere to debounce windows (swoosh: 120ms, click: 80ms).
- [ ] PDF export forces 1920×1080 dimensions with background graphics enabled and animations bypassed.
