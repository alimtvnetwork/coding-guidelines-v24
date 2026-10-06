# 04 — Acceptance Criteria

---

## AC-SQZ-001 — Default shadow pair exists

**Given** an agent reads file 42 §1  
**When** it implements any interactive card, image plate, or option row  
**Then** rest uses `--elevation-rest` and hover uses `--elevation-hover`, and label text uses `--text-shadow-rest` / `--text-shadow-hover` without ad hoc rgba literals.

---

## AC-SQZ-002 — Option card hover physics

**Given** a `.presentation-option-card` button  
**When** the pointer hovers and motion is allowed  
**Then** transform is exactly `translate3d(6px, 0, 0)`, duration **220ms**, easing **`cubic-bezier(0.16, 1, 0.3, 1)`**, and opacity becomes **1**.

---

## AC-SQZ-003 — Botanical-light body contrast

**Given** `data-theme="botanical-light"`  
**When** body copy uses `color: hsl(var(--foreground))` on `hsl(var(--background))`  
**Then** `node scripts/verify-botanical-light-contrast.mjs` exits **0** (foreground pair ≥ **4.5:1**).

---

## AC-SQZ-004 — Shortcut overlay

**Given** presenter mode without focus in a text field  
**When** user presses **`/`** or clicks `#ctrl-shortcuts`  
**Then** shortcut dialog opens and lists every group in `PRESENTER_SHORTCUTS_CORE` (file 42 §5.3).

---

## AC-SQZ-005 — Camera and slide jump conflict

**Given** camera stage-fill is enabled and slide builder is **off** and no slide-specific **`1`** bind is active  
**When** user presses bare **`1`**  
**Then** camera stage-fill toggles and slide index does not change.

---

## AC-SQZ-006 — Center stage

**Given** a title-only slide using `.slide-center-stage`  
**When** rendered at **1920×1080**  
**Then** headline block is vertically and horizontally centered with **`padding: 96px 120px`** and max text width **`920px`**.

---

## AC-SQZ-007 — Reduced motion

**Given** `prefers-reduced-motion: reduce`  
**When** hovering an option card  
**Then** transform stays **`none`** and transitions are **`0.01ms`**.

---

## AC-SQZ-008 — Builder isolation

**Given** audience presentation route  
**When** page loads  
**Then** `useEditStore` is not initialized and **`E`** does not mount builder chrome.

---

## AC-SQZ-009 — Legacy aliases

**Given** global CSS after migration  
**When** `--option-text-shadow-rest` is read  
**Then** computed value equals `--text-shadow-rest` (file 42 §1.4).

---

## AC-SQZ-010 — Embed theme not in deck switcher

**Given** the deck theme control from file 40  
**When** rendered  
**Then** it lists exactly **8** ids and does **not** include `botanical-light` or `green-choice`.

---

## AC-SQZ-011 — Builder wins on key 1

**Given** slide builder **`isEditMode`** is true  
**When** user presses **`1`**  
**Then** builder theme quick-switch runs (file 35 §7) and camera stage-fill does **not** run.

---

## AC-SQZ-012 — Core shortcut export integrity

**Given** `shortcuts.ts` copied from file 42 §5.3  
**When** counted  
**Then** `PRESENTER_SHORTCUTS_CORE` has **7** groups and **31** total `ShortcutItem` rows (count items in file 42 §5.3).
