package fileutil

import (
	"os"

	"coding-guidelines/common/pkg/enum/fileoptype"
	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/enum/filewritemodetype"
	"coding-guidelines/common/pkg/enum/openfiletype"
	"coding-guidelines/common/pkg/result"
	"coding-guidelines/common/pkg/streamwriter"
)

type (
	FileOpenModeType = openfiletype.Variant

	FileResult          = result.Wrap[*os.File]
	BytesResult         = result.Wrap[[]byte]
	StringResult        = result.Wrap[string]
	LinesResult         = result.Wrap[[]string]
	BoolResult          = result.Wrap[bool]
	FileInfoResult      = result.Wrap[os.FileInfo]
	Int64Result         = result.Wrap[int64]
	FilePermResult      = result.Wrap[filepermtype.Variant]
	FileOpResult        = result.Wrap[fileoptype.Variant]
	FileWriteModeResult = result.Wrap[filewritemodetype.Variant]
	FileOpenModeResult  = result.Wrap[openfiletype.Variant]
	FileWriterResult    = result.Wrap[*streamwriter.PluggableWriter[any]]
)
