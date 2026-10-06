# 42 — Slide Quiz Preview Chrome & Default Shadow Pair

> **/goal** One copy-paste contract for default **text-shadow** and **box-shadow** tokens, quiz option cards, botanical-light colors, presenter shortcuts, and center-stage layout.
> **/learn** Product context: `02-spec/21-app/13-slide-quiz-preview-and-presenter-chrome/`. HUD geometry: `31-slide-controller-buttons.md`. Deck keys subset: `29-slide-navigation-and-builder.md` section 3.

**Version:** 1.0.0
**Status:** Active
**AI Confidence:** High

---

## 0. Anti-hallucination

Do not invent shadow rgba, shortcut keys, or green hex. Copy section 2 verbatim into CSS. If a key is not listed in section 5, do not bind it.

---

## 1. Semantic tokens (text + elevation)

These names are **global defaults** for slides, quiz runners, image plates, and pricing cards unless a file explicitly overrides them.

| Token | Purpose |
|:---|:---|
| `--text-shadow-rest` | Body and option labels at rest |
| `--text-shadow-hover` | Option labels on card hover |
| `--elevation-rest` | Cards, images, plates at rest |
| `--elevation-hover` | Cards, images, plates on hover |
| `--elevation-selected` | Selected option glow (optional) |

### 1.1 Light surfaces (`:root`, `.light`, `[data-theme="botanical-light"]`)

```css
:root,
.light,
[data-theme="botanical-light"] {
  --text-shadow-rest: 0 1px 2px rgba(0, 0, 0, 0.06);
  --text-shadow-hover: rgba(0, 0, 0, 0.3) 1px 0.7px 0px;
  --elevation-rest: 0 1px 3px 0 rgba(0, 0, 0, 0.04), 0 1px 2px -1px rgba(0, 0, 0, 0.03);
  --elevation-hover: 0 10px 30px -4px hsl(var(--primary) / 0.24), 0 2px 8px -1px rgba(0, 0, 0, 0.35);
  --elevation-selected: 0 0 16px hsl(var(--primary) / 0.27);
}
```

### 1.2 Dark / high-contrast decks (`.dark`, `[data-theme="bright-gold-tech"]` stage text on dark chrome)

```css
.dark,
[data-theme="vscode-dark"],
[data-theme="dracula"],
[data-theme="noir-gold"] {
  --text-shadow-rest: 0 1px 4px rgba(0, 0, 0, 0.45), 0 2px 8px rgba(0, 0, 0, 0.25);
  --text-shadow-hover: #000000 1px 0.7px 0px;
  --elevation-rest: 0 2px 8px rgba(0, 0, 0, 0.35);
  --elevation-hover: 0 12px 36px -6px hsl(var(--primary) / 0.32), 0 4px 12px -2px rgba(0, 0, 0, 0.55);
  --elevation-selected: 0 0 20px hsl(var(--primary) / 0.35);
}
```

Apply utility:

```css
.text-shadow-default {
  text-shadow: var(--text-shadow-rest);
  transition: text-shadow 200ms ease, color 200ms ease;
}
.elevation-default {
  box-shadow: var(--elevation-rest);
  transition: transform 220ms cubic-bezier(0.16, 1, 0.3, 1),
    box-shadow 220ms ease,
    border-color 200ms ease,
    background-color 200ms ease;
}
.elevation-default:hover,
.elevation-default:focus-visible {
  box-shadow: var(--elevation-hover);
}
img.elevation-default,
.media-plate.elevation-default {
  border-radius: var(--radius, 0.75rem);
}
img.elevation-default:hover {
  transform: translate3d(0, -2px, 0);
}
```

---

## 2. Presentation option card (normative)

```css
.presentation-option-card {
  opacity: 0.82;
  box-shadow: var(--elevation-rest);
  transition: transform 220ms cubic-bezier(0.16, 1, 0.3, 1),
    box-shadow 220ms ease,
    border-color 200ms ease,
    background-color 200ms ease,
    background 200ms ease,
    opacity 200ms ease;
  will-change: transform, opacity, box-shadow;
}
.presentation-option-card:hover {
  opacity: 1 !important;
  transform: translate3d(6px, 0, 0) !important;
  border-color: hsl(var(--primary) / 0.75) !important;
  background: linear-gradient(
    90deg,
    hsl(var(--primary) / 0.14) 0%,
    hsl(var(--card) / 0.92) 100%
  ) !important;
  box-shadow: var(--elevation-hover) !important;
}
.presentation-option-card:hover .option-text,
.presentation-option-card:hover .option-text-shadow {
  color: hsl(var(--foreground)) !important;
  font-weight: 600 !important;
  text-shadow: var(--text-shadow-hover) !important;
}
.presentation-option-card:hover .option-badge {
  border-color: hsl(var(--primary) / 0.75) !important;
  background-color: hsl(var(--primary) / 0.22) !important;
  color: hsl(var(--primary)) !important;
  transform: scale(1.05);
}
.option-text,
.option-text-shadow {
  text-shadow: var(--text-shadow-rest);
  transition: text-shadow 200ms ease, color 200ms ease;
}
@media (prefers-reduced-motion: reduce) {
  .presentation-option-card {
    animation-duration: 0.01ms !important;
    transition-duration: 0.01ms !important;
    transform: none !important;
  }
}
```

Selected state (inline styles in TSX are allowed):

```typescript
const selectedGlow = "var(--elevation-selected)";
const selectedStyle = {
  opacity: 1,
  backgroundColor: "hsl(var(--card-active-bg))",
  border: "1px solid hsl(var(--primary) / 0.85)",
  boxShadow: selectedGlow,
};
```

Define `--card-active-bg` on the theme root (botanical-light section 3).

---

## 3. Theme `botanical-light` (improved green)

Register as a light quiz skin. HSL triples only—implementers map to Tailwind `@theme` or `:root`.

```css
[data-theme="botanical-light"] {
  --primary: 142 65% 38%;
  --primary-foreground: 0 0% 100%;
  --background: 140 18% 97%;
  --foreground: 160 22% 10%;
  --card: 0 0% 100%;
  --card-foreground: 160 22% 10%;
  --secondary: 142 45% 93%;
  --secondary-foreground: 142 55% 28%;
  --muted: 150 14% 96%;
  --muted-foreground: 160 9% 42%;
  --border: 150 14% 90%;
  --ring: 142 65% 38%;
  --card-active-bg: 142 48% 94%;
  --radius: 0.75rem;
  color-scheme: light;
  background-color: hsl(var(--background));
  background-image:
    radial-gradient(ellipse 70% 50% at 50% -10%, hsl(142 65% 38% / 0.06), transparent 70%),
    radial-gradient(circle 500px at 100% 100%, hsl(160 35% 40% / 0.03), transparent 60%);
  background-attachment: fixed;
}
```

Hex twins for design reviews only (do not hardcode in components):

| Role | Hex |
|:---|:---|
| Ground | `#F3F7F5` |
| Primary | `#15803D` |
| Active fill | `#E8F5EC` |
| Body text | `#13201B` |
| Muted text | `#5F7369` |

---

## 4. Stagger entrance (optional)

```css
@keyframes slide-up-fade {
  from {
    opacity: 0;
    transform: translate3d(0, 12px, 0);
  }
  to {
    opacity: 1;
    transform: translate3d(0, 0, 0);
  }
}
.slide-up-anim {
  animation: slide-up-fade 420ms cubic-bezier(0.16, 1, 0.3, 1) both;
}
.stagger-1 { animation-delay: 0.06s; }
.stagger-2 { animation-delay: 0.12s; }
.stagger-3 { animation-delay: 0.18s; }
.stagger-4 { animation-delay: 0.24s; }
.stagger-5 { animation-delay: 0.3s; }
.stagger-6 { animation-delay: 0.36s; }
```

---

## 5. Presenter keyboard shortcuts (normative groups)

Implement as `ShortcutGroup[]`. Merge with `29-slide-navigation-and-builder.md` section 3 without deleting rows here.

| Group | Keys | Action |
|:---|:---|:---|
| **Deck navigation** | `→`, `Space`, `Enter` | Next slide / step |
| | `←`, `Backspace` | Previous slide / step |
| | `F` | Toggle fullscreen |
| | `G` | Slide grid overview |
| | `J` | Top slide jumper |
| | `T` | Theme palette |
| | `Esc` | Close overlay / exit fullscreen |
| | `/` | Open keyboard map |
| **Deck builder** | `E` | Toggle slide builder (`35-slide-builder-canvas-inspector.md`) |
| | `S` | Settings panel (`29` section 3) |
| **Quick jump** | `2`–`9` | Start typing slide number (`1` reserved) |
| | `Enter` | Jump to typed number |
| | `Backspace` | Delete last digit |
| | `Esc` | Cancel pending jump |
| **Sidebar** | `Ctrl+1` / `⌘+1` | Toggle slide outline |
| **Camera — power** | `I` | Hard toggle camera |
| | `M` | Soft minimize / restore stream |
| | `P` | Camera fullscreen |
| | `[` | Exit camera fullscreen |
| | `]` | Cinematic 3-state cycle |
| | `1` | Camera stage-fill (bare `1` only) |
| **Camera — sizing** | `+` / `−` | Step PIP size |
| | `O` | Circle ↔ rectangle frame |
| | `H` | Vignette halo |
| **Camera passthrough** | `→`, `↓`, `Enter`, `Space` | Next slide while camera fullscreen |
| | `←`, `PageUp`, `PageDown` | Previous slide |

**Form focus guard:** ignore all single-key shortcuts when `event.target` is `INPUT`, `TEXTAREA`, or `contentEditable`.

**HUD:** Cam button toggles the same pipeline as **`I`**. A keyboard icon or **?** chip calls the same dialog as **`/`**.

---

## 6. Center-stage layout

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
.slide-center-stage .headline {
  max-width: 920px;
  font-weight: 700;
  font-size: clamp(40px, 4.2vw, 72px);
  line-height: 1.08;
  text-shadow: var(--text-shadow-rest);
}
.slide-center-stage .kicker {
  font-weight: 600;
  font-size: clamp(14px, 1.1vw, 18px);
  letter-spacing: 0.12em;
  text-transform: uppercase;
  text-shadow: var(--text-shadow-rest);
}
.slide-center-stage .subcopy {
  max-width: 720px;
  font-size: clamp(18px, 1.6vw, 28px);
  line-height: 1.45;
  color: hsl(var(--muted-foreground));
}
```

---

## 7. Cross-references

| Topic | File |
|:---|:---|
| Five transition names, 0.45s timing | `29-slide-navigation-and-builder.md` |
| HUD pill, Cam, Build, dots | `31-slide-controller-buttons.md` |
| Builder stores and layers | `35-slide-builder-canvas-inspector.md` |
| Eight deck themes | `40-theme-switch.md` |
| Bright gold corporate slides | `05-bright-gold-tech/` |
| General motion curves | `21-css3-animations-and-interactions.md` section 2 |

---

## 8. Checklist for blind agents

- [ ] Paste section 1 and 2 CSS into the app global stylesheet or scoped quiz bundle.
- [ ] Map legacy `green-choice` product id to `botanical-light` tokens or alias both attributes.
- [ ] Wire `SHORTCUTS` array to a dialog; bind `/` and HUD button.
- [ ] Use `.slide-center-stage` for title slides and quiz intro screens.
- [ ] Run AC-SQZ-001 through AC-SQZ-008 in `02-spec/21-app/13-slide-quiz-preview-and-presenter-chrome/04-acceptance-criteria.md`.
