# Plans Index

Master directory of architectural and execution plans.

## Pending Plans

- [02-slides-system-overhaul.md](pending/02-slides-system-overhaul.md): Full slides deck UI and system overhaul.
- [04-guideline-prompt-and-installer-upgrade.md](pending/04-guideline-prompt-and-installer-upgrade.md): Guideline prompt and installer enhancements.
- [09-update-prompts-and-release.md](pending/09-update-prompts-and-release.md): Update prompts and release lifecycle (deferred under WOR policy).
- [11-code-red-refactor-remediation.md](pending/11-code-red-refactor-remediation.md): Remediate Code Red enum, boolean, and query wrapper violations across the codebase.

## Completed Plans

- [01-repository-infrastructure-cicd-and-consolidation.md](completed/01-repository-infrastructure-cicd-and-consolidation.md): Repository hygiene, encoding normalization, lowercase conventions, AI scripts `<details>` documentation, CI/CD quality automation, and plan memory consolidation.
- [02-appfault-result-monad-and-verification-systems.md](completed/02-appfault-result-monad-and-verification-systems.md): Go `*appfault.AppError` standard, `Result[T]` generic containers, dynamic type conversions, deterministic map sorting, and `ReflectSetTo` fast path.
- [03-fileutil-pathinfo-and-enum-architecture.md](completed/03-fileutil-pathinfo-and-enum-architecture.md): Modular file operations, cross-platform temp resolution, .NET-style `PathInfo`/`FolderInfo`/`FileInfo` objects, concurrency locking, and 1:1 modular `BaseEnumer` enums.
- [04-applogger-taxonomy-streaming-and-task-db.md](completed/04-applogger-taxonomy-streaming-and-task-db.md): Structured AppLogger subsystem, split SQLite logging (`logs.db` + `tasks/<task-id>.db`), rotating file sink, generic `LazyOnce`, errcmd streaming, named writers, and typed streamers.
