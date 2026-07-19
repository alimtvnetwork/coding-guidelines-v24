import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 WF-005: release ceremony. One command, deterministic order, signed artifacts.
 */

const BEFORE = `# Friday afternoon, ad-hoc release.
git tag v5.42.0
git push --tags
# ...forget to bump package.json, forget CHANGELOG, forget to build the slides zip,
# forget to attach it to the GitHub release, forget to run diagrams:rebaseline,
# forget to run npm run sync, forget to update readme pinned version.
# Monday: "hey the release notes are empty and the download link 404s."`;

const AFTER = `# One command. Deterministic order. Every step logged. Fails fast.
npm run release -- --minor

# Under the hood (scripts/release.mjs) runs in this exact order and stops on the first failure:
# 1. verify clean working tree
# 2. verify current branch is main and up to date
# 3. run full CI locally (tsgo, vitest, playwright smoke, axe, SRA validator)
# 4. bump version in package.json (patch, minor, or major flag)
# 5. run node scripts/sync-version.mjs
# 6. run npm run diagrams:rebaseline (only if diagram sources changed)
# 7. run npm run sync (mirrors, health score, readme stats)
# 8. verify CHANGELOG has an entry for the new version
# 9. build slides-app/dist and zip
# 10. create signed git tag vX.Y.Z
# 11. push commit + tag
# 12. create GitHub Release, upload dist.zip, paste CHANGELOG section as notes,
#     inject "a11y clean" and "SRA green" badges
# 13. print SHA-256 of the uploaded artifact so consumers can verify`;

export default function ReleaseCeremonySlide() {
  return (
    <SlideLayout
      eyebrow="Rule 63 · Workflow · release is one command, deterministic order, signed artifacts"
      title="`npm run release` runs the full ceremony end-to-end. Manual releases are banned."
      subtitle="Ad-hoc `git tag && git push` releases skip half the ceremony half the time. Codify the order, gate every step on the previous one succeeding, log everything, and make the human's job pressing one button."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="bash"
          before={BEFORE}
          after={AFTER}
          beforeLabel="Bad. Every release is a checklist re-invented in Slack. Three of the last five releases shipped with a stale CHANGELOG, two shipped with the readme pinned to the previous version, one shipped with no dist.zip attached and 400+ users hit a 404 on the download link. Root cause: the ceremony lived in a human's head; humans forget on Friday afternoons."
          afterLabel="Good. `npm run release -- --minor` (or `--patch` / `--major`) runs every step in a fixed order. Any step that fails aborts the release before the tag is pushed, so a half-released state is impossible. Every step writes a timestamped log line to `.release-logs/vX.Y.Z.log` and is greppable after the fact. The human's job is one command, one confirmation prompt, and a link to the finished release."
        />
        <ActionPanel
          slideId="62-release-ceremony"
          symptom="Q1 release audit: 5 of 12 releases had at least one missing artifact (dist.zip, changelog entry, readme pin, diagram baseline). Mean time between release and first hotfix was 4.2 hours, with 3 of those hotfixes being 'we forgot to run sync' or 'we forgot to bump the docs version'. Two customer-facing incidents traced to a release where `diagrams:rebaseline` was skipped and the docs site showed stale diagrams for 9 days. Root cause: the release procedure was tribal knowledge, and every author executed a slightly different subset of it."
          rule="Release is one command, `npm run release -- --{patch,minor,major}`, backed by `scripts/release.mjs`. Steps run in the exact order shown and each is a hard gate on the next: (1) clean tree, (2) on `main` and up to date, (3) full CI green locally (tsgo, vitest, playwright smoke, axe-core, SRA validator, mermaid render cache), (4) version bump, (5) `sync-version.mjs`, (6) `diagrams:rebaseline` when diagram sources changed, (7) `npm run sync` for mirrors and stats, (8) CHANGELOG entry present for the new version (fails otherwise), (9) build and zip `slides-app/dist`, (10) signed git tag, (11) push commit + tag, (12) GitHub Release with dist.zip attached and CHANGELOG section pasted as notes plus `a11y clean` and `SRA green` badges, (13) print SHA-256 of the zip. Any step failure aborts before the tag is pushed. Manual tags are banned; the pre-push hook (`.husky/pre-push`) rejects tag pushes not created by the release script (checked via a signature file dropped in step 10)."
          doThis="Enforce mechanically: (1) commit `scripts/release.mjs` as the single entry point, with `--dry-run` for rehearsal; (2) `npm run release` is the only sanctioned invocation, documented in the root readme and in `mem://preferences/release-ceremony`; (3) `.husky/pre-push` rejects any tag push where `.release-logs/vX.Y.Z.log` is absent, so hand-typed `git tag && git push --tags` fails at the client; (4) branch protection requires the release commit to be authored by the release script (check for the trailer `Release-Script: scripts/release.mjs`); (5) release notes template baked into the script pulls the matching CHANGELOG section verbatim, so drift is impossible; (6) SHA-256 of the uploaded zip is printed at the end and posted to the release notes for consumer verification; (7) monthly retro (RCA-002) reviews any release that needed a hotfix within 24 hours and lists which step, if any, should be added or hardened."
        />
      </div>
    </SlideLayout>
  );
}
