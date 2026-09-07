package fileutil

import (
	"os"

	"coding-guidelines/common/pkg/enum/openfiletype"
	"coding-guidelines/common/pkg/result"
)

type (
	FileOpenModeType = openfiletype.Variant

	FileResult     = result.Wrap[*os.File]
	BytesResult    = result.Wrap[[]byte]
	StringResult   = result.Wrap[string]
	LinesResult    = result.Wrap[[]string]
	BoolResult     = result.Wrap[bool]
	FileInfoResult = result.Wrap[os.FileInfo]
	Int64Result    = result.Wrap[int64]
)
