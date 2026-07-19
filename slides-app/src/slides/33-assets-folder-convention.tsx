import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 35: Assets folder convention.
 *
 * Source: spec/17/31 line 44 (rule 12). Assets live at
 * assets/<NN-folder>/<NN-file>.<ext> with two-digit sequence prefixes.
 */

const BEFORE = `src/
  img/
    logo.svg
    logo-final.svg           // ❌ which one is current?
    logo-final-v2.svg        // ❌ version in the filename
    hero-bg-large-2.png      // ❌ no folder, no order
public/
  icons/
    check.svg                // ❌ split between src/ and public/
    Check-Copy.svg           // ❌ mixed case, "Copy" suffix
components/
  Button/
    icon.svg                 // ❌ colocated, not discoverable`;

const AFTER = `assets/
  01-icons/
    01-check.svg
    02-close.svg
    03-arrow-right.svg
  02-logos/
    01-wordmark.svg
    02-monogram.svg
    03-favicon.png
  03-illustrations/
    01-hero-background.png
    02-empty-state.svg
  04-fonts/
    01-inter-variable.woff2

// Two-digit prefix on every folder AND every file.
// Ordering is explicit. No "final", no "v2", no "-copy".`;

export default function AssetsFolderConventionSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 35 · Assets folder convention"
      title="Two-digit prefix on every folder. Two-digit prefix on every file."
      subtitle="Every image, icon, font, and illustration lives under `assets/<NN-folder>/<NN-file>.<ext>`. The sequence number is the ordering contract, so no reviewer ever asks 'which logo is current' and no PR ships a `-final-v2` filename. One tree, one convention, discoverable by `ls`."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="text"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Assets scattered across src/img, public/icons, and component folders. Filenames encode versions."
          afterLabel="✅ One assets/ tree, NN-prefixed folders and files, no version suffixes, ordering is explicit."
        />
        <ActionPanel
          slideId="33-assets-folder-convention"
          symptom="A designer ships a new logo. Two weeks later the marketing page renders `logo-final.svg`, the app header renders `logo-final-v2.svg`, and the favicon still points at `logo.svg`. Nobody can answer 'which one is current' without opening three folders. A rebrand PR then breaks half the surfaces because assets were split between `src/img/`, `public/icons/`, and per-component folders."
          rule="All assets live under a single top-level `assets/` tree. Every folder gets a two-digit sequence prefix (`01-icons/`, `02-logos/`). Every file inside gets its own two-digit prefix (`01-wordmark.svg`, `02-monogram.svg`). No version suffixes in filenames (no `-final`, `-v2`, `-copy`, `-old`). If a file is replaced, edit it in place and let git own the history. Per spec/17/31 line 44 rule 12."
          doThis="On any PR that adds or replaces an asset: place it under the correct `assets/NN-folder/` subtree, pick the next free `NN-` prefix for the filename, and update every import to the new path in the same PR. If you find yourself typing `-final` or `-v2`, stop, overwrite the current file, and let the git diff speak. Reject PRs that add assets outside `assets/` or without the sequence prefix."
        />
      </div>
    </SlideLayout>
  );
}
