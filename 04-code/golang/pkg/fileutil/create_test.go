package fileutil

import (
	"os"
	"path/filepath"
	"testing"
)

func verifyCreateDir(t *testing.T, dirPath string) {
	resDir := CreateDir(dirPath, FilePermExecutable)
	if resDir.HasError() {
		t.Fatalf("Expected CreateDir to succeed, got %v", resDir.Fault().Error())
	}

	stat, err := os.Stat(dirPath)
	if err != nil || !stat.IsDir() {
		t.Fatalf("Directory was not created")
	}
}

func verifyCreateFile(t *testing.T, filePath string) {
	resFile := CreateFile(filePath, FilePermStandard)
	if resFile.HasError() {
		t.Fatalf("Expected CreateFile to succeed, got %v", resFile.Fault().Error())
	}

	_ = resFile.Data().Close()

	stat, err := os.Stat(filePath)
	if err != nil || stat.IsDir() {
		t.Fatalf("File was not created")
	}
}

func TestCreateDirAndFile(t *testing.T) {
	tmp := t.TempDir()
	dirPath := filepath.Join(tmp, "sub")
	verifyCreateDir(t, dirPath)

	filePath := filepath.Join(dirPath, "test.txt")
	verifyCreateFile(t, filePath)
}
