# Subtask 02: Header, Menu Navigation & Icon Micro-Transitions
Traceability ID: Task-04
Spec Reference: [02-spec/07-design-system/25-ai-adaptable-design-system/02-menu/01-navigation-and-menu.md](../../../../02-spec/07-design-system/25-ai-adaptable-design-system/02-menu/01-navigation-and-menu.md)
Target Files: `02-spec/07-design-system/25-ai-adaptable-design-system/02-menu/01-navigation-and-menu.md`
Action: Specify sticky ~70px navigation header, expanding CSS3 underline link accents, and 3-phase header icon micro-interactions.
Acceptance Criteria:
- Fixed 70px height, translucent backdrop (`rgba(255, 255, 255, 0.85)` + `backdrop-filter: blur(12px)`), hairline border.
- Animated menu underline via `transform: scaleX(...)` with `cubic-bezier(0.16, 1, 0.3, 1)` easing.
- Header icon hover contract with 1.08x scale lift, `--color-primary` color shift, and ambient background wash.
Targeted Verification: Verified layout geometry, hardware-accelerated transforms, and outline/solid button pairing.
