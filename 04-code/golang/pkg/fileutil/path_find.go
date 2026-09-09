package fileutil

import (
	"os"
	"path/filepath"
	"strings"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

func isPatternMatch(pattern string, target string) bool {
	matched, err := filepath.Match(pattern, target)
	if err != nil {
		return false
	}

	return matched
}

func matchesPattern(pattern, name, relPath string) bool {
	hasSlash := strings.Contains(pattern, SepSlash) || strings.Contains(pattern, SepBackslash)
	if hasSlash {
		return isPatternMatch(pattern, ToSlash(relPath))
	}

	return isPatternMatch(pattern, name)
}

func findMatchingPaths(root string, pattern string) ResultSlice[string] {
	if len(root) == 0 {
		return result.FailSlice[string](appfault.NewFile(errtype.Validation, "", "empty root path"))
	}

	var matches []string
	fault := executeWalk(root, func(path string, isDir bool) *appfault.AppError {
		rel, _ := filepath.Rel(root, path)
		if matchesPattern(pattern, filepath.Base(path), rel) {
			matches = append(matches, path)
		}

		return nil
	})
	if fault != nil {
		return result.FailSlice[string](fault)
	}

	return result.OkSlice(matches)
}

func appendMatchedFile(matches *[]*FileInfo, root, path, pattern string) {
	rel, _ := filepath.Rel(root, path)
	if matchesPattern(pattern, filepath.Base(path), rel) {
		*matches = append(*matches, NewFileInfo(path))
	}
}

func findMatchingFiles(root string, pattern string) ResultSlice[*FileInfo] {
	if len(root) == 0 {
		return result.FailSlice[*FileInfo](appfault.NewFile(errtype.Validation, "", "empty root path"))
	}

	var matches []*FileInfo
	fault := executeWalk(root, func(path string, isDir bool) *appfault.AppError {
		if !isDir {
			appendMatchedFile(&matches, root, path, pattern)
		}

		return nil
	})
	if fault != nil {
		return result.FailSlice[*FileInfo](fault)
	}

	return result.OkSlice(matches)
}

func appendMatchedFolder(matches *[]*FolderInfo, root, path, pattern string) {
	rel, _ := filepath.Rel(root, path)
	if matchesPattern(pattern, filepath.Base(path), rel) {
		*matches = append(*matches, NewFolderInfo(path))
	}
}

func findMatchingFolders(root string, pattern string) ResultSlice[*FolderInfo] {
	if len(root) == 0 {
		return result.FailSlice[*FolderInfo](appfault.NewFile(errtype.Validation, "", "empty root path"))
	}

	var matches []*FolderInfo
	fault := executeWalk(root, func(path string, isDir bool) *appfault.AppError {
		if isDir {
			appendMatchedFolder(&matches, root, path, pattern)
		}

		return nil
	})
	if fault != nil {
		return result.FailSlice[*FolderInfo](fault)
	}

	return result.OkSlice(matches)
}

func appendFilterMatch(matches *[]string, path string, filterFn FileFilterFunc) {
	info, err := os.Stat(path)
	if err == nil && filterFn(path, info) {
		*matches = append(*matches, path)
	}
}

func filterMatchingPaths(root string, filterFn FileFilterFunc) ResultSlice[string] {
	if len(root) == 0 || filterFn == nil {
		return result.FailSlice[string](appfault.NewFile(errtype.Validation, "", "empty root or nil filter"))
	}

	var matches []string
	fault := executeWalk(root, func(path string, isDir bool) *appfault.AppError {
		appendFilterMatch(&matches, path, filterFn)

		return nil
	})
	if fault != nil {
		return result.FailSlice[string](fault)
	}

	return result.OkSlice(matches)
}
