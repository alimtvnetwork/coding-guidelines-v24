# 05 — Blind-Agent Fault Register & Remediation

> **/goal** List every known spec gap that caused sub-100% confidence, and point to the fixed section.
> **/learn** After remediation, blind agents implement only from file **42 v2** + this register. Run `node scripts/verify-botanical-light-contrast.mjs` before merge.

**Version:** 1.0.0
**Status:** Active

---

## Fault table

| ID | Fault | Symptom for blind AI | Remediation |
|:---|:---|:---|:---|
| F-01 | `botanical-light` looked like a 9th deck theme | Agent adds id to `40-theme-switch.md` table | **42 §9** + **40 §7** — embed-only; never increments count of 8 |
| F-02 | Primary `142 65% 38%` on background **3.19:1** (fails AA) | Agent uses `--primary` for body links on ground | **42 §3** — primary **`142 70% 30%`**; rule: body copy uses `--foreground` only |
| F-03 | No legacy CSS aliases | Half the repo keeps `--option-text-shadow-*` | **42 §1.4** — mandatory forward aliases |
| F-04 | Shortcuts were markdown only | Agent omits keys or invents labels | **42 §5.1–5.3** — paste `PRESENTER_SHORTCUTS_CORE` + handlers |
| F-05 | HUD shortcut control unspecified | Agent adds random icon/placement | **31 §3.2.1** — `Keyboard` icon, id `ctrl-shortcuts` |
| F-06 | Dark elevation looked inferred | Agent changes rgba freely | **42 §1.3** — dark text-shadow from exam L99–100; plate shadow from `modern-quiz-card` L132–136 |
| F-07 | AC-SQZ-003 wording wrong | Tested primary on background | **04-acceptance** — AC-SQZ-003 tests `--foreground` on `--background` |
| F-08 | No executable verification | Agent claims pass without math | **scripts/verify-botanical-light-contrast.mjs** |
| F-09 | `green-choice` vs `botanical-light` split | Two palettes drift | **42 §3** — duplicate selectors, same custom properties |
| F-10 | Builder hotkey `B` vs `E` in 31 vs 29 | Double bind or wrong mode | **42 §5.4** — `E` toggles slide builder; `B` only if file 35 documents it on same store |
| F-11 | Slide-specific shortcut groups | Agent drops or hallucinates profiles | **42 §5.2** — optional `SLIDE_SPECIFIC_SHORTCUTS`; never required for AC pass |
| F-12 | Nested theme scope unclear | Shadows missing inside embed | **42 §9.2** — inner root must re-declare §1 tokens |
| F-13 | Key **`1`** means camera, builder theme, or profile | Triple bind | **42 §5.4** dispatch order; **AC-SQZ-011** |

---

## Score after remediation

| Scope | Score | License file |
|:---|:---:|:---|
| File 42 sections 1–6, 9 + §5 TS + verification script | **100** | Copy verbatim; run script |
| Full product (all slide layouts + custom decks) | **74** | Still needs layout files 34/28 |

The **100** score applies only to the closed contract in the right column. Anything outside that list remains out of license.
