package fileutil

import (
	"os"
	"path/filepath"

	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/enum/openfiletype"
	"coding-guidelines/common/pkg/errtype"
)

type FilePathOps struct {
	workDir string
	relPath string
	absPath string
}

func resolveAbs(path string) string {
	if len(path) == 0 {
		return ""
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return filepath.Clean(path)
	}

	return filepath.Clean(abs)
}

func newAbsFilePathOps(cleanPath string) *FilePathOps {
	return &FilePathOps{
		workDir: filepath.Dir(cleanPath),
		relPath: filepath.Base(cleanPath),
		absPath: cleanPath,
	}
}

func newRelFilePathOps(cleanPath string) *FilePathOps {
	return &FilePathOps{
		workDir: CurrentDir,
		relPath: cleanPath,
		absPath: resolveAbs(cleanPath),
	}
}

func NewFilePathOps(path string) *FilePathOps {
	if len(path) == 0 {
		return newRelFilePathOps(CurrentDir)
	}

	cleanPath := filepath.Clean(path)
	if filepath.IsAbs(cleanPath) {
		return newAbsFilePathOps(cleanPath)
	}

	return newRelFilePathOps(cleanPath)
}

func resolveAtAbs(w, r string) string {
	if filepath.IsAbs(r) {
		return r
	}

	return resolveAbs(filepath.Join(w, r))
}

func NewFilePathOpsAt(workDir, relPath string) *FilePathOps {
	cleanRel := filepath.Clean(relPath)
	cleanWork := filepath.Clean(workDir)
	var absPath string

	if filepath.IsAbs(cleanRel) {
		absPath = cleanRel
	} else {
		absPath = filepath.Join(cleanWork, cleanRel)
	}

	return &FilePathOps{
		workDir: cleanWork,
		relPath: cleanRel,
		absPath: absPath,
	}
}

func (f *FilePathOps) Clone() *FilePathOps {
	if f == nil {
		return nil
	}

	return &FilePathOps{
		workDir: f.workDir,
		relPath: f.relPath,
		absPath: f.absPath,
	}
}

func (f *FilePathOps) WithWorkDir(dir string) *FilePathOps {
	if f == nil {
		return NewFilePathOpsAt(dir, CurrentDir)
	}

	return NewFilePathOpsAt(dir, f.relPath)
}

func (f *FilePathOps) WithRelPath(rel string) *FilePathOps {
	if f == nil {
		return NewFilePathOps(rel)
	}

	return NewFilePathOpsAt(f.workDir, rel)
}

func (f *FilePathOps) Join(elem ...string) *FilePathOps {
	if f == nil {
		return NewFilePathOps(filepath.Join(elem...))
	}

	newRel := filepath.Join(append([]string{f.relPath}, elem...)...)
	newAbs := filepath.Join(append([]string{f.absPath}, elem...)...)

	return &FilePathOps{
		workDir: f.workDir,
		relPath: newRel,
		absPath: newAbs,
	}
}

func (f *FilePathOps) WorkDir() string {
	if f == nil {
		return ""
	}

	return f.workDir
}

func (f *FilePathOps) RelPath() string {
	if f == nil {
		return ""
	}

	return f.relPath
}

func (f *FilePathOps) AbsPath() string {
	if f == nil {
		return ""
	}

	return f.absPath
}

func (f *FilePathOps) String() string {
	if f == nil {
		return ""
	}

	return f.absPath
}

func (f *FilePathOps) Dir() string {
	if f == nil {
		return ""
	}

	return filepath.Dir(f.absPath)
}

func (f *FilePathOps) Base() string {
	if f == nil {
		return ""
	}

	return filepath.Base(f.absPath)
}

func (f *FilePathOps) Ext() string {
	if f == nil {
		return ""
	}

	return filepath.Ext(f.absPath)
}

func (f *FilePathOps) IsExists() bool {
	if f == nil {
		return false
	}

	if len(f.absPath) == 0 {
		return false
	}

	_, err := os.Stat(f.absPath)

	return err == nil
}

func (f *FilePathOps) Exists() bool {
	return f.IsExists()
}

func (f *FilePathOps) Stat() FileInfoResult {
	if f == nil {
		return FileInfoFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	return Stat(f.absPath)
}

func ensureParentDirOf(path string) BoolResult {
	dir := filepath.Dir(path)
	if !isValidParentDir(dir) {
		return BoolSuccess(true)
	}

	return EnsureDir(dir, filepermtype.Executable)
}

func (f *FilePathOps) EnsureParentDir() BoolResult {
	if f == nil {
		return BoolFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	if len(f.absPath) == 0 {
		return BoolFailureMsg(errtype.Validation, "", "empty path")
	}

	return ensureParentDirOf(f.absPath)
}

func (f *FilePathOps) CreateIfNotExist(perm FilePermType) FileResult {
	if f == nil {
		return FileFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	dirRes := f.EnsureParentDir()
	if dirRes.IsFailed() {
		return FileFailureFault(dirRes.Fault())
	}

	return OpenFile(f.absPath, openfiletype.ReadWriteOrCreateOnly, perm)
}

func (f *FilePathOps) createAndClose(perm FilePermType) BoolResult {
	res := f.CreateIfNotExist(perm)
	if res.IsFailed() {
		return BoolFailureFault(res.Fault())
	}

	_ = res.Data().Close()

	return BoolSuccess(true)
}

func (f *FilePathOps) EnsureFile(perm FilePermType) BoolResult {
	if f == nil {
		return BoolFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	if f.IsExists() {
		return BoolSuccess(true)
	}

	return f.createAndClose(perm)
}

func (f *FilePathOps) ReadBytes() BytesResult {
	if f == nil {
		return BytesFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	return ReadBytes(f.absPath)
}

func (f *FilePathOps) ReadString() StringResult {
	if f == nil {
		return StringFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	return ReadString(f.absPath)
}

func (f *FilePathOps) ReadLines() LinesResult {
	if f == nil {
		return LinesFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	return ReadLines(f.absPath)
}

func (f *FilePathOps) WriteBytes(data []byte, perm FilePermType) BoolResult {
	if f == nil {
		return BoolFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	return WriteBytes(f.absPath, data, perm)
}

func (f *FilePathOps) WriteString(content string, perm FilePermType) BoolResult {
	if f == nil {
		return BoolFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	return WriteString(f.absPath, content, perm)
}

func (f *FilePathOps) WriteLines(lines []string, perm FilePermType) BoolResult {
	if f == nil {
		return BoolFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	return WriteLines(f.absPath, lines, perm)
}

func (f *FilePathOps) WriteAtomic(data []byte, perm FilePermType) BoolResult {
	if f == nil {
		return BoolFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	return WriteAtomic(f.absPath, data, perm)
}

func (f *FilePathOps) AppendBytes(data []byte, perm FilePermType) BoolResult {
	if f == nil {
		return BoolFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	return AppendBytes(f.absPath, data, perm)
}

func (f *FilePathOps) AppendString(content string, perm FilePermType) BoolResult {
	if f == nil {
		return BoolFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	return AppendString(f.absPath, content, perm)
}

func (f *FilePathOps) AppendLines(lines []string, perm FilePermType) BoolResult {
	if f == nil {
		return BoolFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	return AppendLines(f.absPath, lines, perm)
}

func (f *FilePathOps) Open(openMode FileOpenModeType, perm FilePermType) FileResult {
	if f == nil {
		return FileFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	return OpenFile(f.absPath, openMode, perm)
}

func (f *FilePathOps) Create(perm FilePermType) FileResult {
	if f == nil {
		return FileFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	return CreateFile(f.absPath, perm)
}

func (f *FilePathOps) Delete() BoolResult {
	if f == nil {
		return BoolFailureMsg(errtype.Validation, "", "nil FilePathOps")
	}

	return DeleteFile(f.absPath)
}

func (fileNamespace) Target(path string) *FilePathOps {
	return NewFilePathOps(path)
}

func (fileNamespace) At(workDir, relPath string) *FilePathOps {
	return NewFilePathOpsAt(workDir, relPath)
}

func (fileNewCreator) Target(path string) *FilePathOps {
	return NewFilePathOps(path)
}

func (fileNewCreator) At(workDir, relPath string) *FilePathOps {
	return NewFilePathOpsAt(workDir, relPath)
}

func (fileNewCreator) FilePathOps(path string) *FilePathOps {
	return NewFilePathOps(path)
}

func (fileNewCreator) FilePathOpsAt(workDir, relPath string) *FilePathOps {
	return NewFilePathOpsAt(workDir, relPath)
}
