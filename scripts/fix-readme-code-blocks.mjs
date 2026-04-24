#!/usr/bin/env node
/**
 * fix-readme-code-blocks.mjs
 *
 * Reformats fenced code blocks in readme.md to comply with the blank-line
 * rules from:
 *   spec/02-coding-guidelines/01-cross-language/04-code-style/
 *     03-blank-lines-and-spacing.md
 *
 * Rules implemented:
 *   R4  Blank line before `return` / `throw` when preceded by other statements
 *   R5  Blank line after `}` when followed by more code
 *   R10 Blank line before control structures (if/for/foreach/while/switch/try)
 *       when preceded by a non-brace statement
 *   R12 No leading blank line at the start of a function/method body
 *   R13 Collapse 2+ consecutive blank lines down to 1
 *
 * Targets only TS / JS / Go / PHP fenced blocks. Skips bash/json/text/markdown
 * fences and any block whose first comment line marks it as an intentional
 * violation (`// ❌` or `// FORBIDDEN`).
 *
 * Idempotent: running twice produces zero diff.
 *
 * Usage:
 *   node scripts/fix-readme-code-blocks.mjs            # rewrite in place
 *   node scripts/fix-readme-code-blocks.mjs --check    # exit 1 if changes needed
 *   node scripts/fix-readme-code-blocks.mjs --file X   # target a different file
 */

import { readFileSync, writeFileSync } from "node:fs";
import { resolve } from "node:path";

const ARGS = process.argv.slice(2);
const CHECK_ONLY = ARGS.includes("--check");
const FILE_FLAG_INDEX = ARGS.indexOf("--file");
const TARGET_FILE = resolve(
  FILE_FLAG_INDEX >= 0 ? ARGS[FILE_FLAG_INDEX + 1] : "readme.md"
);

const ELIGIBLE_LANGS = new Set([
  "ts",
  "tsx",
  "typescript",
  "js",
  "jsx",
  "javascript",
  "go",
  "golang",
  "php",
]);

const CONTROL_KEYWORDS = ["if", "for", "foreach", "while", "switch", "try"];

const isBlank = (line) => line.trim() === "";

const stripLineComment = (line) => {
  // Strip // and # comments but keep string contents intact (best-effort).
  // This is good enough for whole-line classification.
  const trimmed = line.trim();

  if (trimmed.startsWith("//") || trimmed.startsWith("#")) {
    return "";
  }

  return trimmed;
};

const startsControlStructure = (line) => {
  const code = stripLineComment(line);

  for (const keyword of CONTROL_KEYWORDS) {
    if (code.startsWith(`${keyword} `) || code.startsWith(`${keyword}(`)) {
      return true;
    }
  }

  return false;
};

const startsReturnOrThrow = (line) => {
  const code = stripLineComment(line);

  return (
    code.startsWith("return ") ||
    code === "return;" ||
    code === "return" ||
    code.startsWith("throw ")
  );
};

const isCloseBraceOnly = (line) => {
  const code = stripLineComment(line);

  return code === "}" || code === "};" || code === "})" || code === "});";
};

const startsContinuation = (line) => {
  // Lines that legitimately follow `}` without a blank line: else/catch/finally,
  // another close brace, chained method calls (rare in our examples).
  const code = stripLineComment(line);

  return (
    code.startsWith("}") ||
    code.startsWith("else") ||
    code.startsWith("catch") ||
    code.startsWith("finally") ||
    code.startsWith(".") ||
    code.startsWith(")")
  );
};

const isOpenBlockLine = (line) => {
  // Heuristic: line ends in `{` (function/method/block opener).
  const code = stripLineComment(line);

  return code.endsWith("{");
};

const isStatementLine = (line) => {
  // Any non-blank, non-pure-brace, non-comment line counts as a statement
  // for the purposes of R4/R10 "preceded by other statements".
  const trimmed = line.trim();

  if (trimmed === "") {
    return false;
  }

  if (trimmed.startsWith("//") || trimmed.startsWith("#") || trimmed.startsWith("*")) {
    return false;
  }

  if (trimmed === "{" || trimmed === "}" || trimmed === "};") {
    return false;
  }

  return true;
};

const blockIsIntentionalViolation = (lines) => {
  // Look at the first ~3 non-blank lines for an explicit "❌" / "FORBIDDEN"
  // marker. If found, leave the block untouched — it is a teaching example.
  let scanned = 0;

  for (const line of lines) {
    if (scanned >= 3) {
      break;
    }

    if (isBlank(line)) {
      continue;
    }

    scanned += 1;

    if (line.includes("❌") || /\bFORBIDDEN\b/.test(line)) {
      return true;
    }
  }

  return false;
};

const reformatBlock = (rawLines) => {
  if (blockIsIntentionalViolation(rawLines)) {
    return rawLines;
  }

  // Step 1: collapse 2+ blank lines down to 1 (R13).
  const collapsed = [];
  let prevBlank = false;

  for (const line of rawLines) {
    const blank = isBlank(line);

    if (blank && prevBlank) {
      continue;
    }

    collapsed.push(line);
    prevBlank = blank;
  }

  // Step 2: strip leading blank lines after each `{` opener (R12).
  const noLeadingBlanks = [];

  for (let i = 0; i < collapsed.length; i += 1) {
    const line = collapsed[i];
    noLeadingBlanks.push(line);
    const opensBlock = isOpenBlockLine(line);
    const next = collapsed[i + 1];

    if (opensBlock && next !== undefined && isBlank(next)) {
      // Skip the blank that follows the opener.
      i += 1;
    }
  }

  // Step 3: walk and insert blank lines per R4, R5, R10.
  const out = [];

  for (let i = 0; i < noLeadingBlanks.length; i += 1) {
    const line = noLeadingBlanks[i];
    const prev = out[out.length - 1] ?? "";
    const prevTrimmed = prev.trim();
    const prevIsBlank = prevTrimmed === "";
    const prevIsOpener = isOpenBlockLine(prev);
    const prevIsCloser = isCloseBraceOnly(prev);
    const prevIsStatement = isStatementLine(prev);

    // R5: blank line after `}` when followed by more code (not another close,
    // not else/catch/finally, not a chained call).
    if (prevIsCloser && !isBlank(line) && !startsContinuation(line)) {
      out.push("");
    }

    // R10: blank line before control structure when preceded by a statement.
    else if (
      startsControlStructure(line) &&
      prevIsStatement &&
      !prevIsBlank &&
      !prevIsOpener
    ) {
      out.push("");
    }

    // R4: blank line before return/throw when preceded by other statements.
    else if (
      startsReturnOrThrow(line) &&
      prevIsStatement &&
      !prevIsBlank &&
      !prevIsOpener
    ) {
      out.push("");
    }

    out.push(line);
  }

  // Final pass: re-collapse any double-blanks introduced (defensive).
  const final = [];
  prevBlank = false;

  for (const line of out) {
    const blank = isBlank(line);

    if (blank && prevBlank) {
      continue;
    }

    final.push(line);
    prevBlank = blank;
  }

  return final;
};

const FENCE_RE = /^(\s*)(```+)([A-Za-z0-9_-]*)\s*$/;

const processMarkdown = (source) => {
  const lines = source.split("\n");
  const out = [];
  let i = 0;

  while (i < lines.length) {
    const line = lines[i];
    const fenceMatch = line.match(FENCE_RE);

    if (fenceMatch === null) {
      out.push(line);
      i += 1;
      continue;
    }

    const [, indent, fence, lang] = fenceMatch;
    const langLower = lang.toLowerCase();
    out.push(line);
    i += 1;

    // Collect block body until matching fence.
    const body = [];
    let closed = false;

    while (i < lines.length) {
      const inner = lines[i];
      const innerFence = inner.match(FENCE_RE);
      const isClosing =
        innerFence !== null &&
        innerFence[1] === indent &&
        innerFence[2].length === fence.length &&
        innerFence[3] === "";

      if (isClosing) {
        closed = true;
        break;
      }

      body.push(inner);
      i += 1;
    }

    const reformatted = ELIGIBLE_LANGS.has(langLower)
      ? reformatBlock(body)
      : body;

    for (const bodyLine of reformatted) {
      out.push(bodyLine);
    }

    if (closed) {
      out.push(lines[i]);
      i += 1;
    }
  }

  return out.join("\n");
};

const main = () => {
  const original = readFileSync(TARGET_FILE, "utf8");
  const updated = processMarkdown(original);
  const changed = updated !== original;

  if (CHECK_ONLY) {
    if (changed) {
      console.error(
        `[fix-readme-code-blocks] ${TARGET_FILE} would be modified — run without --check to apply.`
      );
      process.exit(1);
    }

    console.log(`[fix-readme-code-blocks] ${TARGET_FILE} already compliant.`);
    return;
  }

  if (changed === false) {
    console.log(`[fix-readme-code-blocks] ${TARGET_FILE} already compliant.`);
    return;
  }

  writeFileSync(TARGET_FILE, updated, "utf8");
  console.log(`[fix-readme-code-blocks] ${TARGET_FILE} reformatted.`);
};

main();
