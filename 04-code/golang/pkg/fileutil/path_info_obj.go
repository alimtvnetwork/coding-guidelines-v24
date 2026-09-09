package fileutil

import (
	"os"
	"path/filepath"
)

type PathInfo struct {
	path string
}

func NewPathInfo(path string) *PathInfo {
	if len(path) == 0 {
		return &PathInfo{path: CurrentDir}
	}

	return &PathInfo{path: filepath.Clean(path)}
}

func (p *PathInfo) Path() string {
	if p == nil {
		return ""
	}

	return p.path
}

func (p *PathInfo) Name() string {
	if p == nil || len(p.path) == 0 {
		return ""
	}

	return filepath.Base(p.path)
}

func (p *PathInfo) IsExists() bool {
	if p == nil || len(p.path) == 0 {
		return false
	}

	_, err := os.Stat(p.path)

	return err == nil
}

func (p *PathInfo) Exists() bool {
	return p.IsExists()
}

func (p *PathInfo) IsDir() bool {
	if p == nil || len(p.path) == 0 {
		return false
	}

	info, err := os.Stat(p.path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func (p *PathInfo) IsFile() bool {
	if p == nil || len(p.path) == 0 {
		return false
	}

	info, err := os.Stat(p.path)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

func (p *PathInfo) AsFolder() *FolderInfo {
	if p == nil {
		return NewFolderInfo(CurrentDir)
	}

	return NewFolderInfo(p.path)
}

func (p *PathInfo) AsFile() *FileInfo {
	if p == nil {
		return NewFileInfo(CurrentDir)
	}

	return NewFileInfo(p.path)
}

func (p *PathInfo) Parent() *FolderInfo {
	if p == nil {
		return nil
	}

	return p.AsFolder().Parent()
}
