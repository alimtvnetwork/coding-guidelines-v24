# Letterly Pull-First, Muse SSH Key Delivery, Heat Map Tracking & Dual Installers (`21-letterly-muse-ssh-heatmap-general-installer`)

This specification defines standards and implementations for:
1. Re-sequencing all Letterly prompts so that Item 1 is always Git pull via `gitmap pull` (fallback `git pull`) with conflict resolution.
2. Mobile prompt one-liner invariant.
3. Muse SSH key generation via `gitmap ssh create` with standardized `<email>-muse<muse-sequence>-<operator-name>` identity, account email discovery, and key file download delivery protocol (fixing SSH formatting errors).
4. In-flight heat map action logging in `execute-parent-task-with-n-steps-v6` and `gitmap heatmap`.
5. Dual installers: full coding guidelines installer vs lightweight general repository installer (`general-install.ps1`, `general-install-config.json`).
6. GitMap CLI integration for general repository installation via `gitmap cg install general` and `gitmap cd general`.

## Specifications

- [`01-overview-and-architecture.md`](01-overview-and-architecture.md) — Architectural overview, contracts, and folder structures.
- [`02-acceptance-criteria.md`](02-acceptance-criteria.md) — 23 measurable acceptance criteria across all 6 capability areas.
