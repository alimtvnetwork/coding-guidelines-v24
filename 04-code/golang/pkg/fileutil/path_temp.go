package fileutil

import (
	"os"
	"path/filepath"
	"runtime"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

type envLookupFunc func(string) string

type pathExistsFunc func(string) bool

func appendEnvVal(list []string, val string) []string {
	if len(val) == 0 {
		return list
	}

	return append(list, val)
}

func appendEnvSubpath(list []string, base string, sub ...string) []string {
	if len(base) == 0 {
		return list
	}

	elems := append([]string{base}, sub...)

	return append(list, filepath.Join(elems...))
}

func appendCacheTmp(list []string, home string, dirExists pathExistsFunc) []string {
	if len(home) == 0 {
		return list
	}

	cacheTmp := filepath.Join(home, ".cache", "tmp")
	if dirExists(cacheTmp) {
		return append(list, cacheTmp)
	}

	return list
}

func windowsTempCandidates(getenv envLookupFunc) []string {
	var list []string
	list = appendEnvVal(list, getenv("TEMP"))
	list = appendEnvVal(list, getenv("TMP"))
	list = appendEnvSubpath(list, getenv("LOCALAPPDATA"), "Temp")
	list = appendEnvSubpath(list, getenv("USERPROFILE"), "AppData", "Local", "Temp")

	return append(list, os.TempDir())
}

func darwinTempCandidates(getenv envLookupFunc) []string {
	var list []string
	list = appendEnvVal(list, getenv("TMPDIR"))
	list = append(list, os.TempDir())

	return append(list, "/tmp")
}

func linuxTempCandidates(getenv envLookupFunc, home string, dirExists pathExistsFunc) []string {
	var list []string
	list = appendEnvVal(list, getenv("TMPDIR"))
	list = appendEnvVal(list, getenv("XDG_RUNTIME_DIR"))
	list = appendCacheTmp(list, home, dirExists)

	return append(list, "/tmp")
}

func isDirectoryExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func hostTempCandidates() []string {
	if runtime.GOOS == "windows" {
		return windowsTempCandidates(os.Getenv)
	}

	if runtime.GOOS == "darwin" {
		return darwinTempCandidates(os.Getenv)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		home = ""
	}

	return linuxTempCandidates(os.Getenv, home, isDirectoryExists)
}

func ensureCandidateDir(cand string) error {
	info, err := os.Stat(cand)
	if err == nil {
		if info.IsDir() {
			return nil
		}

		return os.ErrInvalid
	}

	return os.MkdirAll(cand, 0755)
}

func selectTempDir(candidates []string) (string, error) {
	var lastErr error

	for _, cand := range candidates {
		if len(cand) == 0 {
			continue
		}

		err := ensureCandidateDir(cand)
		if err == nil {
			return cand, nil
		}

		lastErr = err
	}

	return "", lastErr
}

func UserTempDir() StringResult {
	dir, err := selectTempDir(hostTempCandidates())
	if err != nil {
		return StringFailure(errtype.IO, err, dir, "failed to resolve temp directory")
	}

	return StringSuccess(dir)
}

func UserTempPath(subpath ...string) StringResult {
	res := UserTempDir()
	if res.IsFailed() {
		return res
	}

	elems := append([]string{res.Data()}, subpath...)

	return StringSuccess(filepath.Join(elems...))
}

func resolveTargetDir(dir string) (string, *appfault.AppError) {
	if len(dir) > 0 {
		return dir, nil
	}

	res := UserTempDir()
	if res.IsFailed() {
		return "", res.Fault()
	}

	return res.Data(), nil
}

func applyPermToFile(f *os.File, perm FilePermType) {
	if perm != 0 {
		_ = f.Chmod(perm.Mode())
	}
}

func CreateTempFile(dir string, pattern string, perm FilePermType) FileResult {
	targetDir, fault := resolveTargetDir(dir)
	if fault != nil {
		return result.WrapFailure[*os.File](fault)
	}

	f, err := os.CreateTemp(targetDir, pattern)
	if err != nil {
		return FileFailure(errtype.IO, err, targetDir, "failed to create temp file")
	}

	applyPermToFile(f, perm)

	return FileSuccess(f)
}

func applyPermToDir(dir string, perm FilePermType) {
	if perm != 0 {
		_ = os.Chmod(dir, perm.Mode())
	}
}

func CreateTempDir(dir string, pattern string, perm FilePermType) StringResult {
	targetDir, fault := resolveTargetDir(dir)
	if fault != nil {
		return result.WrapFailure[string](fault)
	}

	d, err := os.MkdirTemp(targetDir, pattern)
	if err != nil {
		return StringFailure(errtype.IO, err, targetDir, "failed to create temp dir")
	}

	applyPermToDir(d, perm)

	return StringSuccess(d)
}

func TempFile(pattern string) FileResult {
	return CreateTempFile("", pattern, filepermtype.Standard)
}

func TempDir(pattern string) StringResult {
	return CreateTempDir("", pattern, filepermtype.Standard)
}

type pathTempNamespace struct{}

func (pathTempNamespace) UserTempDir() StringResult {
	return UserTempDir()
}

func (pathTempNamespace) UserTempPath(subpath ...string) StringResult {
	return UserTempPath(subpath...)
}

func (pathTempNamespace) TempFile(pattern string) FileResult {
	return TempFile(pattern)
}

func (pathTempNamespace) TempDir(pattern string) StringResult {
	return TempDir(pattern)
}

func (pathTempNamespace) CreateTempFile(dir string, pattern string, perm FilePermType) FileResult {
	return CreateTempFile(dir, pattern, perm)
}

func (pathTempNamespace) CreateTempDir(dir string, pattern string, perm FilePermType) StringResult {
	return CreateTempDir(dir, pattern, perm)
}
