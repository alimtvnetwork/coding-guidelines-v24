package fileutil

import (
	"path/filepath"
	"sync"
	"testing"
)

func runConcurrentWriters(path string, count int, wg *sync.WaitGroup) {
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ExportTextLocked(path, "locked content", FilePermStandard)
		}()
	}
}

func runConcurrentReaders(t *testing.T, path string, count int, wg *sync.WaitGroup) {
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := ReadTextLocked(path)
			if res.IsSuccess() {
				if res.Data() != "locked content" {
					t.Errorf("Unexpected content: %s", res.Data())
				}
			}
		}()
	}
}

func TestLockerMechanisms(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "locked.txt")
	var wg sync.WaitGroup

	runConcurrentWriters(path, 50, &wg)
	runConcurrentReaders(t, path, 50, &wg)
	wg.Wait()
}

func TestLocker_RefCountAndEviction(t *testing.T) {
	path := "sample_lock_path.txt"
	_ = GetFileLock(path)
	if GetLockRefCount(path) != 1 {
		t.Fatalf("expected refCount 1, got %d", GetLockRefCount(path))
	}

	_ = GetFileLock(path)
	if GetLockRefCount(path) != 2 {
		t.Fatalf("expected refCount 2, got %d", GetLockRefCount(path))
	}

	verifyLockerEviction(t, path)
}

func verifyLockerEviction(t *testing.T, path string) {
	ReleaseFileLock(path)
	if GetLockRefCount(path) != 1 {
		t.Fatalf("expected refCount 1, got %d", GetLockRefCount(path))
	}

	ReleaseFileLock(path)
	if GetLockRefCount(path) != 0 {
		t.Fatalf("expected refCount 0, got %d", GetLockRefCount(path))
	}
}

func TestLocker_ExtraReleaseZeroPanics(t *testing.T) {
	ReleaseFileLock("non_existent_path_evicted.txt")
	ReleaseFileLock("non_existent_path_evicted.txt")
	if ActiveLockCount() < 0 {
		t.Fatalf("unexpected negative lock count")
	}
}
