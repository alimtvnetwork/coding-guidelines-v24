# Plans Index

Master directory of architectural and execution plans.

## Pending Plans

- [02-slides-system-overhaul.md](pending/02-slides-system-overhaul.md): Full slides deck UI and system overhaul.
- [04-guideline-prompt-and-installer-upgrade.md](pending/04-guideline-prompt-and-installer-upgrade.md): Guideline prompt and installer enhancements.
- [09-update-prompts-and-release.md](pending/09-update-prompts-and-release.md): Update prompts and release lifecycle (deferred under WOR policy).
- [11-code-red-refactor-remediation.md](pending/11-code-red-refactor-remediation.md): Remediate Code Red enum, boolean, and query wrapper violations across the codebase.

## Completed Plans

- [26-prune-enum-const-aliases-and-enumer-types.md](completed/26-prune-enum-const-aliases-and-enumer-types.md): Eliminate Redundant Enum Const Aliases & Prune Over-Engineered Enumer Types.
- [25-enum-packages-isolation.md](completed/25-enum-packages-isolation.md): Dedicated Enum Packages Isolation (`fileoptype`, `filewritemodetype`, `severitytype`, `prioritytype`).
- [24-filepermtype-enum-package.md](completed/24-filepermtype-enum-package.md): FilePermType Enum Package Isolation (`pkg/enum/filepermtype`).
- [23-errtype-enum-folder-isolation.md](completed/23-errtype-enum-folder-isolation.md): Errtype Enum Folder Isolation (`logleveltype` & `processstatetype`).
- [22-bytetype-package-and-baseenumer-consolidation.md](completed/22-bytetype-package-and-baseenumer-consolidation.md): Canonical ByteType Package & BaseEnumer Generic Helpers Consolidation.
- [21-repo-wide-enum-isolation.md](completed/21-repo-wide-enum-isolation.md): Repo-Wide Enum Isolation — Dedicated Files and Packages.
- [20-fileutil-filename-matching-and-bound-path-ops.md](completed/20-fileutil-filename-matching-and-bound-path-ops.md): Fileutil Struct-Filename Alignment, Bound FilePathOps, and AI Skill Playbook.
- [19-fileutil-struct-grouping-and-modular-decomposition.md](completed/19-fileutil-struct-grouping-and-modular-decomposition.md): Fileutil Struct Grouping & Creator Modular Decomposition.
- [18-modular-filepath-util-and-cross-platform-temp.md](completed/18-modular-filepath-util-and-cross-platform-temp.md): Modular Filepath Utilities, Cross-Platform User Temp Path, Environment Variable Expansion, and Namespace Grouping.
- [17-structured-fileutil-context-and-concrete-results.md](completed/17-structured-fileutil-context-and-concrete-results.md): Structured fileutil context injection, path error constructors, and concrete result type aliases.
- [16-coredata-wrap-and-baseenumer-expansion.md](completed/16-coredata-wrap-and-baseenumer-expansion.md): Modular BaseEnum family (byte/utf8, utf16, utf32/rune, string, number/int) and coredata collection combinators & dynamic struct formatting.
- [15-simple-verifier-consolidation.md](completed/15-simple-verifier-consolidation.md): Unified SimpleVerifier and SimpleVerifyChecker contracts, AsSimpleVerifier/AsSimpleVerifyChecker methods, and compile-time static type assertions.
- [14-generic-typecast-and-result-checkers.md](completed/14-generic-typecast-and-result-checkers.md): Bulletproof typecast conversions, reflection performance optimization, Result and AppError generic casting, and Checker interface contracts.

- [01-apperror-new-constructors.md](completed/01-apperror-new-constructors.md): Go AppError namespace constructors.
- [03-apperror-human-logger-methods.md](completed/03-apperror-human-logger-methods.md): Human-readable logger methods on AppError.
- [05-rename-overviews-and-installer-json.md](completed/05-rename-overviews-and-installer-json.md): Rename overviews and installer JSON configurations.
- [06-fix-encoding.md](completed/06-fix-encoding.md): Repository encoding and BOM normalization.
- [07-trailing-newlines-and-ai-scripts.md](completed/07-trailing-newlines-and-ai-scripts.md): Trailing newlines and AI fix scripts catalog.
- [08-lowercase-changelog.md](completed/08-lowercase-changelog.md): Lowercase changelog migration.
- [10-rca-and-boolean-fix.md](completed/10-rca-and-boolean-fix.md): Root cause analysis and boolean condition remediation.
- [12-prompt-architect-version-tracking.md](completed/12-prompt-architect-version-tracking.md): Version tracking safeguards for prompt architect.
- [13-cicd-pipeline-consolidation-and-owner-review.md](completed/13-cicd-pipeline-consolidation-and-owner-review.md): CI/CD pipeline workflow consolidation and open review questions.
