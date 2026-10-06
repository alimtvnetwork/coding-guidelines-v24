# 04 — Acceptance Criteria

---

## AC-SQZ-001 — Default shadow pair exists

**Given** an agent reads file 42  
**When** it implements any interactive card, image plate, or option row  
**Then** rest uses `--elevation-rest` and hover uses `--elevation-hover`, and label text uses `--text-shadow-rest` / `--text-shadow-hover` without ad hoc rgba literals.

---

## AC-SQZ-002 — Option card hover physics

**Given** a `.presentation-option-card` button  
**When** the pointer hovers and motion is allowed  
**Then** transform is exactly `translate3d(6px, 0, 0)`, duration **220ms**, easing **`cubic-bezier(0.16, 1, 0.3, 1)`**, and opacity becomes **1**.

---

## AC-SQZ-003 — Botanical-light contrast

**Given** theme attribute `data-theme="botanical-light"`  
**When** primary text sits on `--background`  
**Then** contrast ratio is at least **4.5:1** for body copy using file 42 HSL table (no raw `#16A34A` on `#F4F8F8` without verification).

---

## AC-SQZ-004 — Shortcut overlay

**Given** presenter mode without focus in an text field  
**When** user presses **`/`**  
**Then** shortcut dialog opens and lists every group defined in file 42 section 5.

---

## AC-SQZ-005 — Camera and slide jump conflict

**Given** camera stage-fill is enabled  
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
