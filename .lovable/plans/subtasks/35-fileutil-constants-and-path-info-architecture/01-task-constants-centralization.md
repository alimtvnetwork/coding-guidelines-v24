# Subtask 35.1: Fileutil Constants Centralization & Refactoring

## Status: Complete

## Goal
Centralize all path prefixes, separators, environment variable keys, platform identifiers, and directory permissions into `04-code/golang/pkg/fileutil/consts.go`, and replace hardcoded literals across `path_normalize.go`, `path_env.go`, `path_info.go`, `path_temp.go`, `file_path_ops.go`, and `fileutil.go`.

## Target Files
- `04-code/golang/pkg/fileutil/consts.go`
- `04-code/golang/pkg/fileutil/path_normalize.go`
- `04-code/golang/pkg/fileutil/path_env.go`
- `04-code/golang/pkg/fileutil/path_info.go`
- `04-code/golang/pkg/fileutil/path_temp.go`
- `04-code/golang/pkg/fileutil/file_path_ops.go`
- `04-code/golang/pkg/fileutil/fileutil.go`

## Acceptance Criteria
1. `consts.go` defines typed and untyped constants:
   - Windows long path prefixes: `PrefixWinLongPath = "\\\\?\\"`, `PrefixWinUNCLongPath = "\\\\?\\UNC\\"`, `PrefixUNC = "\\\\"`
   - Separators: `SepSlash = "/"`, `SepBackslash = "\\"`, `CharSlash = '/'`, `CharBackslash = '\\'`, `CharColon = ':'`, `CharEqual = '='`
   - Special dirs: `CurrentDir = "."`, `ParentDir = ".."`, `ExtDot = "."`, `CharDot = '.'`
   - Tilde: `HomeTilde = "~"`, `PrefixTildeSlash = "~/"`, `PrefixTildeBackslash = "~\\\\"`
   - Env syntax: `CharDollar = '$'`, `CharPercent = '%'`, `CharLeftBrace = '{'`, `CharRightBrace = '}'`
   - Temp dirs & keys: `DirCache = ".cache"`, `DirTmp = "tmp"`, `DirTemp = "Temp"`, `DirAppData = "AppData"`, `DirLocal = "Local"`, `UnixTempDir = "/tmp"`
   - Env keys: `EnvTemp = "TEMP"`, `EnvTmp = "TMP"`, `EnvLocalAppData = "LOCALAPPDATA"`, `EnvUserProfile = "USERPROFILE"`, `EnvTmpDir = "TMPDIR"`, `EnvXdgRuntimeDir = "XDG_RUNTIME_DIR"`
   - Platforms: `OSWindows = "windows"`, `OSDarwin = "darwin"`
   - Slug: `SlugHyphen = "-"`, `PatternSlugNonAlphaNum = "[^a-z0-9-]+"`, `PatternSlugDashes = "-{2,}"`
   - Permissions: `DefaultDirPerm os.FileMode = 0755`
2. In `path_normalize.go`:
   - Replace literal `\\?\` in `cleanLongPath` and other functions with `PrefixWinLongPath`.
   - Replace `\\?\UNC\` with `PrefixWinUNCLongPath`, `\\` with `PrefixUNC`, `/` and `\` with `SepSlash` and `SepBackslash`.
   - Replace magic string `.` in `Clean` with `CurrentDir`.
3. In `path_env.go`:
   - Replace `~/`, `~\`, `~` with `PrefixTildeSlash`, `PrefixTildeBackslash`, `HomeTilde`.
   - Replace `$`, `%`, `{`, `}` with `CharDollar`, `CharPercent`, `CharLeftBrace`, `CharRightBrace`.
4. In `path_info.go`:
   - Replace regex literals with `PatternSlugNonAlphaNum` and `PatternSlugDashes`.
   - Replace `.` with `ExtDot`, `-` with `SlugHyphen`, `..` with `ParentDir`.
5. In `path_temp.go`:
   - Replace env names and directory names with constants.
   - Replace `0755` with `DefaultDirPerm`.
6. In `file_path_ops.go` and `fileutil.go`:
   - Replace `.` with `CurrentDir` and `0755` with `DefaultDirPerm`.
7. Every refactored function remains $\le 15$ lines.
8. `cd 04-code/golang && go test ./pkg/fileutil/...` passes cleanly.
