#!/usr/bin/env node
/**
 * SS-02 structure check: every rule slide must expose the
 * Symptom -> Rule -> Action pattern via either <ActionPanel> or
 * <PrincipleCard> (both are the sanctioned carriers of the pattern).
 *
 * Exempt slides (title, section overviews, closing) live in EXEMPT below.
 *
 * Exits 1 on any violation, with a per-file report of what is missing.
 *
 * Run locally:  node scripts/validate-slides-sra.mjs
 * CI:           .github/workflows/slides-smoke.yml (step "Validate SRA")
 */
import { readFileSync, readdirSync } from "node:fs";
import { resolve, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const HERE = dirname(fileURLToPath(import.meta.url));
const SLIDES_DIR = resolve(HERE, "..", "slides-app", "src", "slides");

// Slides that legitimately do NOT carry a rule (title card, closing).
// Keep this list minimal; every rule/content slide MUST follow SRA.
const EXEMPT = new Set(["00-title.tsx", "01-table-of-contents.tsx", "12-closing.tsx"]);

// Required prop names on each carrier component.
const CARRIERS = [
  { tag: "ActionPanel",   props: ["symptom", "rule", "doThis"] },
  { tag: "PrincipleCard", props: ["symptom", "rule", "action"] },
];

function hasProp(source, tag, propName) {
  // Match `<Tag ... propName=` allowing multi-line prop blocks. We only
  // need presence, not value shape, because the SRA copy itself is
  // reviewed by humans; the check guards structure, not prose.
  const re = new RegExp(
    `<${tag}\\b[^>]*?\\b${propName}\\s*=`,
    "s",
  );
  return re.test(source);
}

function auditFile(fileName) {
  const source = readFileSync(resolve(SLIDES_DIR, fileName), "utf8");
  const violations = [];
  for (const carrier of CARRIERS) {
    const usesCarrier = new RegExp(`<${carrier.tag}\\b`).test(source);
    if (!usesCarrier) continue;
    const missing = carrier.props.filter(
      (p) => !hasProp(source, carrier.tag, p),
    );
    if (missing.length > 0) {
      violations.push(
        `<${carrier.tag}> is missing required prop(s): ${missing.join(", ")}`,
      );
    }
  }
  const usesAnyCarrier = CARRIERS.some((c) =>
    new RegExp(`<${c.tag}\\b`).test(source),
  );
  if (!usesAnyCarrier) {
    violations.push(
      "no <ActionPanel> or <PrincipleCard> found. Every SS-02 slide must " +
        "carry Symptom -> Rule -> Action via one of these components.",
    );
  }
  return violations;
}

function main() {
  const files = readdirSync(SLIDES_DIR)
    .filter((f) => f.endsWith(".tsx"))
    .filter((f) => !EXEMPT.has(f))
    .sort();

  const report = new Map();
  for (const file of files) {
    const violations = auditFile(file);
    if (violations.length > 0) report.set(file, violations);
  }

  if (report.size === 0) {
    console.log(
      `[slides-sra] OK: ${files.length} slide(s) follow the Symptom -> Rule -> Action pattern.`,
    );
    process.exit(0);
  }

  console.error("[slides-sra] FAIL: SS-02 structural violations found.\n");
  for (const [file, violations] of report) {
    console.error(`  ${file}`);
    for (const v of violations) console.error(`    - ${v}`);
  }
  console.error(
    `\n  Fix by wrapping the slide body in <ActionPanel symptom=... rule=... doThis=...> ` +
      `or one or more <PrincipleCard symptom=... rule=... action=...>.`,
  );
  console.error(
    `  If a slide is legitimately structural (title, closing, section header), ` +
      `add its filename to EXEMPT in scripts/validate-slides-sra.mjs with a reason.`,
  );
  process.exit(1);
}

main();
