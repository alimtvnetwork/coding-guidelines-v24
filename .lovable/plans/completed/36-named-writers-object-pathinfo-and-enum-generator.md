# Plan 36: Logger Named Writers, Object-Oriented PathInfo, and Enum Scaffolder Tooling

## 1. Executive Summary

This plan addresses four direct user requirements:
1. **Logger Named Writers & Introspection:**
   - Add `Name() string` to the `LogSinker` interface and implement on all sinks (`ConsoleSink`, `FileSink`, `RotatingFileSink`, `ApiSink`, `SQLiteSink`, `ZapAdapter`, `CompositeSink`, `StreamerSink`).
   - Add introspection methods `Writers() []LogSink`, `WriterNames() []string`, and `Streamers() []any` to `Logger` and `appLogger`.
2. **Object-Oriented `PathInfo` Architecture:**
   - Transform `PathInfo` from a thin shell into a comprehensive domain object initialized from any path (`fileutil.NewPathInfo(path)`, `fileutil.Path.Inspect(path)`, `fileutil.Path.Info(path)`).
   - Equip `PathInfo`, `FolderInfo`, and `FileInfo` with normalization (`Normalize()`, `ToSlash()`, `Clean()`), navigation (`Up()`, `UpN()`, `Cd()`, `Sub()`, `Parent()`), finding and filtering (`Find(pattern)`, `FindFiles()`, `FindFolders()`, `Filter(fn)`), and recursive looping/walking.
3. **Enum Generator CLI & Root Greeting in `readme.md`:**
   - Enhance `03-ai-scripts/30-enum-generator.py` with `--out` / `-o` parameter and automated directory creation.
   - Showcase the one-line enum generation command in the root `readme.md` greeting section.
4. **Codebase Architecture Review:**
   - Provide an in-depth, candid assessment of the codebase's strengths, enum design, and architectural enhancement opportunities.

---

## 2. Task-Specific Rules

1. **Strictly Relative Git Paths:** All documentation, plans, links, and code paths MUST be strictly relative to the repository root. Zero absolute paths or `file:///` URIs.
2. **Interface Naming Mandate:** Every Go interface MUST end with the `er` suffix (e.g., `LogSinker`, `WritersProvider`, `WriterNamesProvider`, `StreamersProvider`).
3. **Hard Cap $\le 15$ Lines:** Every function or method in Go and Python must stay strictly $\le 15$ lines.
4. **Implicit Booleans:** Implicit boolean evaluations only; fields/variables prefixed with `is` or `has`.

---

## 3. Subtask Decomposition

- [x] `01-task-logger-named-writers-and-introspection.md`: Add `Name()` to `LogSinker` and all sinks; implement `Writers()`, `WriterNames()`, `Streamers()` on `Logger`.
- [x] `02-task-object-oriented-pathinfo-and-search-filters.md`: Refactor `PathInfo`, `FolderInfo`, and `FileInfo` with normalization, navigation, search filters, and walking.
- [x] `03-task-enum-generator-cli-and-root-readme.md`: Add `--out` flag to `30-enum-generator.py` and document one-line command in `readme.md`.
- [x] `04-task-codebase-architecture-analysis-and-verification.md`: Verify all tests and CI/CD quality gates pass 100% green.
