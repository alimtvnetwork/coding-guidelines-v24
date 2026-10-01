# 36 — Website Content Builder Mode & In-Page Editor Specification

> **/goal** Master and enforce the architecture, activation mechanics, 3 operating modes, element identity tagging, inline text editing, media replacement, and Markdown+ZIP export format for client-side website editing.
> **/learn** Master the 3 modes (Browse, Edit, Preview), stable element keys (`data-edit-id` and `data-source-file`), contenteditable caret preservation, gradient accent markup (`[[accent]]`), SEO image replacement, localStorage quota safety, and the deterministic Markdown changelog + ZIP bundle export.

**Version:** 4.0.0
**Status:** Active
**AI Confidence:** High
**Ambiguity:** None

---

## 1. System Vision & Architecture

The **Website Content Builder Mode** is a zero-server, client-side visual editing layer for marketing websites, blogs, and documentation portals:
- Stakeholders or content authors open the live website via a secure query parameter (`?builder=true` or secret key combination).
- Edit text copy, headlines, bullet lists, menu labels, and images directly on the rendered page with instant visual fidelity.
- Generates a deterministic **Markdown Change Register** and a **ZIP archive** containing all modified text and uploaded image assets.
- Developers or autonomous AI coding agents ingest the exported ZIP and apply all changes back to source code in one deterministic pass.
- **Zero Server Footprint:** No backend database, no user accounts, no server API, and zero bloat for ordinary public visitors.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│ 1. Activation (`?builder=true`) ──► 2. In-Page Visual Editing (3 Modes)     │
│    - Browse: Normal navigation         - Edit: Blue outlines, contenteditable│
│    - Preview: Clean presentation view  - Toolbar: [[accent]] formatting     │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
                                       ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│ 3. Client Storage & History Store   ──► 4. One-Click Export Bundle          │
│    - LocalStorage + IndexedDB Blobs     - `changes.md` structured log       │
│    - Per-element undo/redo stacks       - `assets/` directory inside ZIP    │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 2. The Three Operating Modes

| Mode | Purpose | User Interaction | Visual Chrome |
|:---|:---|:---|:---|
| **`Browse`** | Normal Navigation | All links and buttons navigate normally | Floating builder trigger pill in bottom-right corner. |
| **`Edit`** | In-Page Modification | Editable elements intercept clicks, activate `contenteditable` | Blue focus outlines (`2px solid #2563eb`), floating format toolbars, image swap overlays. |
| **`Preview`** | Pre-Flight Verification | Links disabled; shows edited copy in exact published styles | Clean interface with floating "Export Changes" panel. |

---

## 3. Stable Element Identity & Source Mapping

Every editable element in the codebase is tagged with two non-colliding HTML attributes:

```html
<h1
  data-edit-id="homepage-hero-headline"
  data-source-file="src/content/pages/home.ts"
  class="text-h1 font-display"
>
  Empower Your Enterprise With Autonomous Intelligence
</h1>
```

1. **`data-edit-id`:** Hierarchical string formatted as `[page]-[section]-[element]` (e.g., `home-hero-title`, `about-leadership-bio-01`).
2. **`data-source-file`:** Strictly relative path from the repository root to the content file defining that copy.

---

## 4. Text Editing, Caret Preservation & Accent Toolbar

### 4.1 Contenteditable & Caret Position Management
When an editable text node is clicked in `Edit` mode:
1. Element toggles `contentEditable = "true"`.
2. A mutation listener records changes without forcing React re-mounts that disrupt native caret positioning.
3. On blur, text is sanitized: strips rogue HTML tags, collapses non-breaking spaces, and stores plain text or tokenized markup.

### 4.2 Gradient Accent Formatting (`[[accent]]`)
Marketing headlines often include a gradient-styled accent word. The builder provides a floating micro-toolbar with an **Accent Highlight** button:
- Selecting a phrase and clicking "Accent" wraps the selection in `[[accent]]` markers:
  ```text
  Transforming Legacy Systems into [[accent]]Streamlined Agility[[/accent]]
  ```
- The renderer dynamically transforms `[[accent]]...[[/accent]]` into `<span class="gradient-text">...</span>` on the fly.

---

## 5. Image & Media Replacement Engine

Clicking an editable image plate (`<img data-edit-id="...">` or `<div data-edit-id="..." style="background-image: ...">`) opens the **Media Replacement Modal**:
1. **Upload or Drop:** Drag-and-drop a replacement `.png`, `.webp`, `.jpg`, or `.svg`.
2. **Alt Text Editing:** Inline field to update image accessibility and SEO description.
3. **SEO Filename Normalization:** Automatically slugifies the filename based on the section title (e.g., `enterprise-platform-hero-visual.webp`).
4. **Storage Architecture:** Images are converted to `Blob` objects and cached in IndexedDB (bypassing the 5MB `localStorage` ceiling).

---

## 6. Menu & Navigation Item Editing

Primary navigation labels and URLs are editable directly:
- In `Edit` mode, hovering the header reveals an "Edit Navigation" pill.
- Clicking opens a reorderable list of menu items.
- Authors can update nav labels, edit link targets, or update dropdown descriptions without touching code.

---

## 7. Storage, History & Quota Management

Edits are recorded in an authoritative client-side transaction store:

```typescript
export interface ElementChange {
  editId: string;
  sourceFile: string;
  originalValue: string;
  newValue: string;
  timestamp: number;
  type: "text" | "image" | "link";
  imageBlob?: Blob;
  seoFileName?: string;
  altText?: string;
}

export interface BuilderStore {
  mode: "browse" | "edit" | "preview";
  changes: Record<string, ElementChange>;
  undoStack: ElementChange[][];
  redoStack: ElementChange[][];
}
```

- **Undo / Redo:** Full session history accessible via floating controls or `Cmd/Ctrl+Z`.
- **Reset Baseline:** Single click resets all modifications back to original source values.

---

## 8. Export Format Specification (Markdown + ZIP Bundle)

Clicking **"Export Change Package"** generates a downloadable ZIP archive containing:
1. `changes.md` — The structured change manifest.
2. `assets/` — Subfolder containing all replacement image binaries.

### 8.1 Structure of `changes.md`

```markdown
# Website Content Change Register

**Export Date:** 2026-10-02 03:15:00 UTC  
**Total Changes:** 3 text elements, 1 image replacement

---

## File: src/content/pages/home.ts

### Element: `homepage-hero-headline`
- **Original:** "Old Manual Software Architecture"
- **Updated:** "Transforming Legacy Systems into [[accent]]Streamlined Agility[[/accent]]"

### Element: `homepage-hero-lead`
- **Original:** "We build traditional web apps."
- **Updated:** "Autonomous enterprise software delivery with zero downtime and guaranteed precision."

---

## File: src/content/pages/about.ts

### Element: `about-hero-visual` (Image)
- **Original:** `/assets/images/legacy-team.png`
- **Updated:** `assets/enterprise-leadership-team.webp`
- **Alt Text:** "Executive leadership team collaborating on modern platform architecture"
```

---

## 9. Autonomous Developer / AI Agent Intake Protocol

When an AI agent is tasked with applying a builder export:
1. Read `changes.md` in the unzipped package.
2. For each file referenced in `data-source-file`:
   - Locate the target string using the `Original` value.
   - Replace it with the `Updated` value.
   - Convert `[[accent]]word[[/accent]]` to `<GradientText>word</GradientText>` or the target framework's token syntax.
3. For images:
   - Copy assets from `assets/` to the repository's public image directory.
   - Update image imports or source strings to reference the new path.
4. Run tests or linter to verify zero regressions.

---

## 10. Anti-Hallucination & Quality Verification Checklist

- [ ] Every editable element has both `data-edit-id` and `data-source-file`.
- [ ] No server communication, database connection, or external login API is introduced.
- [ ] Caret position is preserved during in-place contenteditable typing.
- [ ] Gradient accent tags use `[[accent]]` and `[[/accent]]`.
- [ ] Image assets over 1MB are stored via IndexedDB blobs to prevent localStorage exhaustion.
- [ ] Export package strictly outputs a valid ZIP containing `changes.md` and `assets/`.
