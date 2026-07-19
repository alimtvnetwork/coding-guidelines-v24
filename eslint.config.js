import js from "@eslint/js";
import globals from "globals";
import reactHooks from "eslint-plugin-react-hooks";
import reactRefresh from "eslint-plugin-react-refresh";
import tseslint from "typescript-eslint";
import codingGuidelines from "./eslint-plugins/coding-guidelines/index.js";

export default tseslint.config(
  { ignores: ["dist", "linters-cicd/tests/fixtures/**"] },
  {
    extends: [js.configs.recommended, ...tseslint.configs.recommended],
    files: ["**/*.{ts,tsx}"],
    languageOptions: {
      ecmaVersion: 2020,
      globals: globals.browser,
    },
    plugins: {
      "react-hooks": reactHooks,
      "react-refresh": reactRefresh,
      "coding-guidelines": codingGuidelines,
    },
    rules: {
      ...reactHooks.configs.recommended.rules,
      "react-refresh/only-export-components": ["warn", { allowConstantExport: true }],
      "@typescript-eslint/no-unused-vars": "off",

      // ═══════════════════════════════════════════════════════════════
      // 🔴 CANONICAL SIZE TIER (single source of truth)
      // Canonical doc: spec/02-coding-guidelines/00-canonical-size-tier.md
      // ┌────────────────────────────┬──────────────┬─────────┐
      // │ Metric                     │ Limit        │ Level   │
      // ├────────────────────────────┼──────────────┼─────────┤
      // │ Function body (preferred)  │ ≤ 8 lines    │ warn    │
      // │ Function body (hard cap)   │ ≤ 15 lines   │ error   │
      // │ File length                │ ≤ 300 lines  │ error   │
      // │ React component file       │ ≤ 100 lines  │ error   │  (*.tsx override)
      // │ Struct / class             │ ≤ 120 lines  │ error   │  (linter-cicd)
      // └────────────────────────────┴──────────────┴─────────┘
      // Line counts skip blanks and comments. Waiver syntax:
      //   // lint-allow: function-length reason="..." max=N
      // ═══════════════════════════════════════════════════════════════

      // File length ≤ 300 (canonical tier).
      "max-lines": ["error", { max: 300, skipBlankLines: true, skipComments: true }],

      // `any` is prohibited (TS Standards §2.1).
      "@typescript-eslint/no-explicit-any": "error",

      // ═══════════════════════════════════════════════════════════════
      // 🔴 CODE RED RULES — Automatic PR rejection
      // ═══════════════════════════════════════════════════════════════

      // Zero nested if — flatten with early returns or named booleans
      "coding-guidelines/no-nested-if": "error",

      // Boolean variables must use is/has/can/should/was/will prefix
      "coding-guidelines/boolean-naming": "error",

      // No raw string literals in comparisons — use enum/typed constants
      "coding-guidelines/no-magic-strings": "warn",

      // Function body hard cap: 15 lines (canonical tier, error level)
      "coding-guidelines/max-function-lines": ["error", { max: 15 }],

      // Function body preferred: 8 lines (canonical tier, warn level).
      // Kept as warn so the 15-line hard cap is the sole build-failing gate;
      // eliminates the prior 8=error vs 15=error contradiction.
      "coding-guidelines/prefer-function-lines": ["warn", { prefer: 8 }],


      // Promise.all for independent async calls — no sequential await
      "coding-guidelines/promise-all-independent": "error",

      // ═══════════════════════════════════════════════════════════════
      // ⚠️ STYLE RULES — Warnings
      // ═══════════════════════════════════════════════════════════════

      // Blank line before return/throw when preceded by statements (R4)
      "coding-guidelines/blank-line-before-return": "warn",

      // No else after return/throw/continue/break (R7)
      "coding-guidelines/no-else-after-return": "error",
    },
  },
  {
    // React component files: tighter 100-line file cap (canonical tier).
    files: ["**/*.tsx"],
    rules: {
      "max-lines": ["error", { max: 100, skipBlankLines: true, skipComments: true }],
    },
  },
);

