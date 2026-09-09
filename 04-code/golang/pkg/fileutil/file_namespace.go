package fileutil

type fileNamespace struct {
	Open   openOps
	Create createOps
	Write  writeOps
	Append appendOps
	Read   readOps
	Path   pathNamespace
}

var File = &fileNamespace{
	Open:   openOps{},
	Create: createOps{},
	Write:  writeOps{},
	Append: appendOps{},
	Read:   readOps{},
	Path:   Path,
}

type fileNewCreator struct {
	Writer       fileWriterCreator
	Appender     fileAppenderCreator
	BoundWriter  fileBoundWriterCreator
	Path         filePathCreator
	StreamWriter fileStreamWriterCreator
}

var New = &fileNewCreator{
	Writer:       fileWriterCreator{},
	Appender:     fileAppenderCreator{},
	BoundWriter:  fileBoundWriterCreator{},
	Path:         filePathCreator{},
	StreamWriter: fileStreamWriterCreator{},
}
