package fileutil

import "fmt"

type FileWriteModeType uint8

const (
	FileWriteModeDirect   FileWriteModeType = 1
	FileWriteModeAtomic   FileWriteModeType = 2
	FileWriteModeTruncate FileWriteModeType = 3
)

var writeModeNames = map[FileWriteModeType]string{
	FileWriteModeDirect:   "Direct",
	FileWriteModeAtomic:   "Atomic",
	FileWriteModeTruncate: "Truncate",
}

func (m FileWriteModeType) Name() string {
	if name, ok := writeModeNames[m]; ok {
		return name
	}

	return fmt.Sprintf("FileWriteMode(%d)", uint8(m))
}

func (m FileWriteModeType) String() string {
	return m.Name()
}

func (m FileWriteModeType) IsValid() bool {
	_, ok := writeModeNames[m]

	return ok
}
