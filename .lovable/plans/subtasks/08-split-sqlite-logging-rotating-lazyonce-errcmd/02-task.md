# Subtask 02: Text-Based Rotating File Logger with Configurable Size, Retention & Archiving

## Parent Plan
`.lovable/plans/completed/08-split-sqlite-logging-rotating-lazyonce-errcmd.md`

## Target Files
- `04-code/golang/pkg/applogger/rotation_config.go`
- `04-code/golang/pkg/applogger/rotating_file_sink.go`

## Instructions
1. Implement `RotationConfig` struct with fields:
   - `FilePath string`
   - `MaxSizeBytes int64` (default 2 MB: `2 * 1024 * 1024`)
   - `MaxBackups int` (default 20)
   - `IsArchiveEnabled bool`
   - `ArchiveDir string`
   - `IsCompress bool`
2. Implement `RotatingFileSink` conforming to `LogSink` interface (`WriteEntry`, `Sync`, `Close`).
3. Check size on write; when current file size + entry size > `MaxSizeBytes`, rotate active file.
4. Manage backup index and purge or archive backups exceeding `MaxBackups`.
5. Implement thread-safe operations (`sync.Mutex`).
6. Enforce <= 15 lines per function and `*appfault.AppError` return type.
