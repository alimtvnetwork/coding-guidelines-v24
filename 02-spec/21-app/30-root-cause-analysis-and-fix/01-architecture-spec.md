# Architectural Specification: Commit Pseudo-Module Stripping & Subsystem Casing Standards

> **Module:** `02-spec/21-app/30-root-cause-analysis-and-fix`  
> **Document:** Architecture Specification  
> **Status:** Draft / Active  

---

## 1. Architectural Boundaries

This specification defines the dual-level architecture for eliminating tautological `commit` pseudo-modules and supporting flexible Title Case / lowercase module conventions:

1. **Instruction & Prompt Boundary (`coding-guidelines`):**
   - Prompt templates in `01-prompts/` and skills in `.agents/skills/`.
   - Global AI rules in `AGENTS.md` and `.agents/rules/`.
   - Explicit negative invariant: AI agents must never pass `commit` or `commits` as the module in `gitmap cpf`, `gitmap cpc`, `gitmap cbf`, `gitmap cfr`, or `gitmap cpr`.
2. **Engine & Parser Boundary (`gitmap`):**
   - Command parser and normalizer in `cli/cmdcommit/commit_formatter.go`.
   - Automated token deduplication stripping `commit`, `commits`, `gitcommit`.
   - Module casing normalizer: Title Case (`Auth`, `Scanner`, `Ui`) or lowercase (`auth`, `scanner`, `ui`) allowed; ALL-CAPS words (`AUTH`, `SCANNER`) normalized to lowercase or Title Case.

---

## 2. Redundant Token Normalization Algorithm

```text
Input: verb, raw string

1. Extract candidate module:
   candidateMod = strings.Trim(rawModule, "[]() ")
   upperMod = strings.ToUpper(candidateMod)

2. Evaluate isCommitTypeToken(upperMod):
   Types: FEATURE, FEAT, RELEASE, REL, VERSION, BUG, FIX, CHORE, REFACTOR, REFACT, SPEC,
          COMMIT, COMMITS, GITCOMMIT

3. If isRedundantModule(prefix, upperMod) || upperMod in ["COMMIT", "COMMITS", "GITCOMMIT"]:
   - Check if nested hyphen exists in remainder:
     a. If nested exists (e.g. "docs - update readme"):
        candidateMod = nestedMod
        Re-evaluate step 2.
     b. If no nested module exists:
        candidateMod = ""
        hasMod = false
        Formatted message = prefix + item

4. Module Casing Rule (when hasMod is true):
   - If candidateMod is all uppercase (len >= 2 and strings.ToUpper(candidateMod) == candidateMod):
     Normalize to strings.ToLower(candidateMod) (or Title Case).
   - If candidateMod is already lowercase or Title Case:
     Preserve natural casing.
```

---

## 3. Compatibility Invariants

- **Backward Compatibility:** All existing commands like `gitmap cpc "docs - update readme"` or `gitmap cpf "scanner - parallel walker"` remain fully valid.
- **Accidental Tautology Healing:** If an agent or human runs `gitmap cpc "commit - fix formatting"`, the engine transforms it to `Chore: fix formatting` without failing.
- **Nested Module Promotion:** If an agent runs `gitmap cpc "commit - docs - update readme"`, the engine promotes `docs` to produce `Chore: docs - update readme`.
