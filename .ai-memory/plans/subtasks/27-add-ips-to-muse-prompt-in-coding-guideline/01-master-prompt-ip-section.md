# Subtask 27-01: Master Prompt IP Section

Status: COMPLETED
Spec: `02-spec/21-app/27-add-ips-to-muse-prompt-in-coding-guideline/01-architecture-spec.md`

- Add section "Requested Network Endpoints (IPs Muse Has Requested)" to `01-prompts/27-muse-prompts/01-muse-master-prompt.md`, placed after the Master Verification Checklist (Phase 9) and before the closing version footer.
- Content: the 41 `198.19.0.x` port-22 endpoints and the hostname grants exactly as listed in the spec, section 2.
- Bump `> Prompt Version: 6.5.0` (header) and the footer version line to 6.5.1, and extend the footer's summary phrase with "requested-endpoint registry".
- Update `01-prompts/27-muse-prompts/readme.md` catalog row for `01-muse-master-prompt.md`: version 6.5.1 + endpoint registry mention.
- Do NOT touch `02-muse-generic.md`.
