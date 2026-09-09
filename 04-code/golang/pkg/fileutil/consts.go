package fileutil

import (
	"os"

	"coding-guidelines/common/pkg/appfault"
)

// Default I/O chunk and buffer size constants (64 KB).
const (
	DefaultChunkSize  = 64 * 1024
	DefaultBufferSize = DefaultChunkSize
)

// Windows path prefixes.
const (
	PrefixWinLongPath    = `\\?\`
	PrefixWinUNCLongPath = `\\?\UNC\`
	PrefixUNC            = `\\`
)

// Path separators and character tokens.
const (
	SepSlash      = "/"
	SepBackslash  = "\\"
	CharSlash     = '/'
	CharBackslash = '\\'
	CharColon     = ':'
	CharEqual     = '='
)

// Special directory and extension tokens.
const (
	CurrentDir = "."
	ParentDir  = ".."
	ExtDot     = "."
	CharDot    = '.'
)

// User home directory and tilde expansion prefixes.
const (
	HomeTilde            = "~"
	PrefixTildeSlash     = "~/"
	PrefixTildeBackslash = "~\\"
)

// Environment variable syntax delimiter characters.
const (
	CharDollar     = '$'
	CharPercent    = '%'
	CharLeftBrace  = '{'
	CharRightBrace = '}'
)

// Standard system and cache directory names.
const (
	DirCache    = ".cache"
	DirTmp      = "tmp"
	DirTemp     = "Temp"
	DirAppData  = "AppData"
	DirLocal    = "Local"
	UnixTempDir = "/tmp"
)

// Environment variable names for temporary directory discovery.
const (
	EnvTemp          = "TEMP"
	EnvTmp           = "TMP"
	EnvLocalAppData  = "LOCALAPPDATA"
	EnvUserProfile   = "USERPROFILE"
	EnvTmpDir        = "TMPDIR"
	EnvXdgRuntimeDir = "XDG_RUNTIME_DIR"
)

// Operating system platform identifiers for runtime.GOOS comparison.
const (
	OSWindows = "windows"
	OSDarwin  = "darwin"
)

// Slug sanitization constants and patterns.
const (
	SlugHyphen             = "-"
	PatternSlugNonAlphaNum = `[^a-z0-9-]+`
	PatternSlugDashes      = `-{2,}`
)

// Default filesystem permission modes.
const (
	DefaultDirPerm os.FileMode = 0755
)

type (
	ChunkHandlerFunc func(chunk []byte) error

	ChunkCallbackFunc func(chunk []byte) *appfault.AppError

	BoundFileActionFunc func(w *BoundFileWriter) *appfault.AppError

	WithLockFunc = BoundFileActionFunc

	FileFilterFunc func(path string, info os.FileInfo) bool
)
