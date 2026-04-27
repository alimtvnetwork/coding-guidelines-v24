/**
 * JS mirror of linters-cicd/checks/_lib/effective_lines.py
 *
 * MUST stay byte-for-byte semantically identical to the Python module.
 * A fixture-driven parity test in
 * linters-cicd/tests/test_effective_lines_parity.py feeds the same
 * fixtures through both implementations and asserts the same counts.
 *
 * If you change one, change the other in the same commit.
 *
 * Canonical prose spec:
 *   linters-cicd/checks/function-length-prefer8/README.md  §2
 */

const SYNTAX = {
  go:         { lineTokens: ["//"],        blockOpen: "/*", blockClose: "*/", docstringTokens: [] },
  typescript: { lineTokens: ["//"],        blockOpen: "/*", blockClose: "*/", docstringTokens: [] },
  javascript: { lineTokens: ["//"],        blockOpen: "/*", blockClose: "*/", docstringTokens: [] },
  rust:       { lineTokens: ["//", "///"], blockOpen: "/*", blockClose: "*/", docstringTokens: [] },
  php:        { lineTokens: ["//", "#"],   blockOpen: "/*", blockClose: "*/", docstringTokens: [] },
  python:     { lineTokens: ["#"],         blockOpen: null, blockClose: null, docstringTokens: ['"""', "'''"] },
};

function isLineComment(stripped, tokens) {
  return tokens.some((tok) => stripped.startsWith(tok));
}

function countEffective(bodyLines, language) {
  const syntax = SYNTAX[language];
  if (!syntax) throw new Error(`Unknown language for effective-lines counter: ${language}`);
  let count = 0;
  let inBlock = false;
  for (const raw of bodyLines) {
    const stripped = raw.trim();
    if (!stripped) continue;
    if (inBlock) {
      if (syntax.blockClose && stripped.includes(syntax.blockClose)) inBlock = false;
      continue;
    }
    if (syntax.blockOpen && stripped.startsWith(syntax.blockOpen)) {
      const rest = stripped.slice(syntax.blockOpen.length);
      if (syntax.blockClose && !rest.includes(syntax.blockClose)) inBlock = true;
      continue;
    }
    if (isLineComment(stripped, syntax.lineTokens)) continue;
    if (isLineComment(stripped, syntax.docstringTokens)) continue;
    count++;
  }
  return count;
}

export { countEffective, SYNTAX };