import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 WF-004: PR template. Mandatory checklist enforced by CI.
 */

const BEFORE = `## Pull request

Fixes the thing.

- [x] It works on my machine`;

const AFTER = `## Pull request

### Linked issue
Closes #1234

### Root cause (one sentence, before the fix)
The retry loop used the request timestamp instead of the response timestamp, so
retries used a stale clock skew and hit the 5-minute nonce window.

### What changed and why
- Extracted \`computeNonce\` from \`postWithRetry\`; now takes an explicit \`Clock\`.
- Retry loop reads the fresh nonce per attempt.

### Checklist (all boxes required, CI fails otherwise)
- [x] Root cause written above in one sentence, before the fix.
- [x] Tests added at the correct pyramid layer (WF-003). Unit for new logic,
      integration for the retry seam, no new e2e.
- [x] If DB changed: Mermaid ERD committed under \`docs/erd/\` (SCHEMA-006).
- [x] If UI changed: keyboard-only sweep passed AND axe-core clean (A11Y-001).
- [x] CHANGELOG entry added under the current unreleased version.
- [x] No error swallowed. Every new catch logs with context (ERR-001, ERR-002).
- [x] File and function size limits respected (SIZE-001, REACT-007).
- [x] No new placeholder type names (REACT-012).

### Screenshots or logs (before -> after)
Before: 12 percent retry failures per hour.
After:  0 retry failures over 24 h canary.`;

export default function PrTemplateSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 62 · Workflow · PR template is a merge gate, not a suggestion"
      title="Every PR carries linked issue, one-sentence root cause, layered tests, ERD/A11y proofs, and a changelog entry. CI enforces it."
      subtitle="A PR without root cause, without tests at the right layer, or without proof of the a11y sweep is not a PR; it is a hope. The template turns every rule slide (SIZE-001, ERR-001, SCHEMA-006, A11Y-001, REACT-007, REACT-012) into a merge-time check."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="md"
          before={BEFORE}
          after={AFTER}
          beforeLabel="Bad. 'Fixes the thing. It works on my machine.' The reviewer has to reverse-engineer the root cause from the diff, guess the test layer, notice DB changes without an ERD, and hope the author considered accessibility. Every review is a full re-investigation; regressions ship because the reviewer missed what the template should have surfaced; changelog is written days later by someone else who does not remember why."
          afterLabel="Good. The template forces the author to (1) name the linked issue, (2) write the root cause in one sentence before the fix, (3) declare test layer per WF-003, (4) attach an ERD when DB touched, (5) prove keyboard + axe when UI touched, (6) update the changelog in the same commit, (7) confirm no new swallowed errors or placeholder types. The reviewer's job shrinks from investigation to verification. CI fails the PR when any required section is missing."
        />
        <ActionPanel
          slideId="61-pr-template"
          symptom="Six-month audit of merged PRs found: 34 percent had no linked issue, 51 percent had no root cause description, 22 percent shipped DB changes without an ERD, 18 percent shipped UI changes without an axe run, 12 percent shipped test coverage regressions. Two-thirds of Q1 production incidents traced to PRs in these gaps. Reviewers said they 'trusted the author'; authors said they 'thought the reviewer would catch it'. Root cause: no enforced template, so every reviewer applied a different bar, and gaps were nobody's job."
          rule="`.github/pull_request_template.md` is the single source of truth for the required PR shape and is enforced by CI. Required sections (all mandatory, all machine-checkable): (1) Linked issue with `Closes #N`; (2) Root cause in one sentence, above the fix description, present tense; (3) What changed and why, as a bullet list; (4) Merge-gate checklist with 8 items covering WF-003 (test layer), SCHEMA-006 (ERD), A11Y-001 (keyboard + axe), CHANGELOG, ERR-001/002 (no swallowed errors), SIZE-001/REACT-007 (size caps), REACT-012 (no placeholder types); (5) Screenshots or logs showing the before -> after signal (this is the DoD 'verified' bar from the standing prompt). Every checkbox has a linked rule slide, so failing a check points the reviewer at the exact standard. Optional-only when explicitly waived in the PR body with a written reason and an author + reviewer sign-off; waivers are logged to `docs/waivers.md` weekly."
          doThis="Enforce mechanically: (1) commit `.github/pull_request_template.md` with the shape above; (2) GitHub Action `pr-template-check` (`.github/workflows/pr-template.yml`) parses the PR body and fails when any required heading or checkbox is missing, unchecked, or lacks the linked issue; (3) `danger.js` rule that fails when CHANGELOG.md has no entry under the current unreleased version and the diff touches `src/`; (4) `pr-template-check` posts a checklist comment showing which items are missing, linking to the corresponding rule slide (`/#/60-test-pyramid`, `/#/46-erd-required-on-db-prs`, etc.); (5) branch protection requires `pr-template-check` + `danger` + all `test:*` jobs green; (6) waiver flow: author adds `<!-- waiver: SCHEMA-006 reason=doc-only -->`, `danger` accepts it and appends the line to `docs/waivers.md` on merge; (7) monthly review of `docs/waivers.md` in the ops retro (RCA-002)."
        />
      </div>
    </SlideLayout>
  );
}
