import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 28: Line-gap and whitespace discipline.
 *
 * Source: spec/17/31 §"Line-Gap and Whitespace Style" lines 53-62.
 *   1. Blank line before every return / throw (unless only statement).
 *   2. Blank line after closing `}` (unless next is `}` / else / case / catch).
 *   3. Never two blank lines in a row.
 *   4. No blank line right after `{` or right before `}`.
 *   5. One blank line between top-level declarations.
 *   6. Imports grouped stdlib / third-party / first-party absolute / relative,
 *      one blank line between groups, never mixed.
 */

const BEFORE = `import { z } from "zod";
import { readFile } from "node:fs/promises";
import { Button } from "@/components/ui/button";
import { formatDate } from "./utils";
import { logger } from "@/lib/logger";

function loadUser(id: string) {
  if (!id) {
    throw new Error("id required");
  }

  const raw = readFile(path);

  return parse(raw);
}

function parse(raw: string) {
  return schema.parse(raw);
}`;

const AFTER = `import { readFile } from "node:fs/promises";

import { z } from "zod";

import { Button } from "@/components/ui/button";
import { logger } from "@/lib/logger";

import { formatDate } from "./utils";

function loadUser(id: string) {
  if (!id) {
    throw new Error("id required");
  }

  const raw = readFile(path);

  return parse(raw);
}

function parse(raw: string) {
  return schema.parse(raw);
}`;

export default function LineGapDisciplineSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 28 · Line-gap discipline"
      title="Blank lines are punctuation. Use them on purpose."
      subtitle="Imports grouped stdlib, third-party, first-party absolute, then relative. One blank line before every return or throw. Blank line after a closing brace. Never two blank lines in a row, never one right after an opening brace."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 18, marginTop: 4 }}>
        <CodeDiff
          language="typescript"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ mixed import groups, double blanks, tight return"
          afterLabel="✅ four grouped imports, blank before return, blank between decls"
        />
        <ActionPanel
          slideId="26-line-gap-discipline"
          symptom="A reviewer opens the diff and sees imports jumbled across origins, a blank line right after `function loadUser(id) {`, no blank line before `return parse(raw)`, and two adjacent function declarations glued together. Reading rhythm collapses; real logic changes are hidden by whitespace churn."
          rule="Group imports as stdlib, third-party, first-party absolute, first-party relative, with one blank line between groups. One blank line before every `return` or `throw` unless it is the only statement in the block. Blank line after a closing `}` unless the next line is `}`, `else`, `case`, or `catch`. No blank line right after `{` or right before `}`. Never two blank lines in a row. Per spec/17/31 lines 55 to 62."
          doThis="When you touch a file, fix the whitespace in the same diff: reorder and regroup imports, insert the return/throw blank, delete accidental double blanks, remove the blank after opening braces. If you feel the urge to add section-separator blanks inside one function, split the function instead."
        />
      </div>
    </SlideLayout>
  );
}
