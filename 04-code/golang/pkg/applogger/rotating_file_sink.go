package applogger

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/fileutil"
)

// RotatingFileSink manages an active log file with size-based rotation.
type RotatingFileSink struct {
	lock        sync.Mutex
	cfg         RotationConfig
	file        *os.File
	currentSize int64
}

// NewRotatingFileSink initializes a rotating file sink from configuration.
func NewRotatingFileSink(cfg RotationConfig) (*RotatingFileSink, error) {
	cfg.Normalize()
	if err := fileutil.EnsureDir(filepath.Dir(cfg.FilePath), filepermtype.Standard).Fault(); err != nil {
		return nil, err
	}

	sink := &RotatingFileSink{cfg: cfg}
	if err := sink.openActiveFile(); err != nil {
		return nil, err
	}

	return sink, nil
}

// openActiveFile opens or creates the active log file and seeds currentSize.
func (s *RotatingFileSink) openActiveFile() error {
	f, err := os.OpenFile(s.cfg.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}

	stat, err := f.Stat()
	if err != nil {
		_ = f.Close()

		return err
	}

	s.file = f
	s.currentSize = stat.Size()

	return nil
}

// WriteEntry persists the log entry, triggering rotation if size threshold exceeded.
func (s *RotatingFileSink) WriteEntry(e LogEntry) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	data, err := json.Marshal(e)
	if err != nil {
		return err
	}

	line := append(data, '\n')
	if s.currentSize+int64(len(line)) > s.cfg.MaxSizeBytes {
		if rotErr := s.rotate(); rotErr != nil {
			return rotErr
		}
	}

	n, writeErr := s.file.Write(line)
	s.currentSize += int64(n)

	return writeErr
}

// rotate closes the current file, renames/archives it, prunes old backups, and opens fresh file.
func (s *RotatingFileSink) rotate() error {
	if s.file != nil {
		_ = s.file.Close()
		s.file = nil
	}

	backupPath := s.generateBackupPath()
	if err := os.Rename(s.cfg.FilePath, backupPath); err != nil {
		return err
	}

	if s.cfg.IsArchiveEnabled {
		_ = s.archiveFile(backupPath)
	}

	s.pruneBackups()

	return s.openActiveFile()
}

// generateBackupPath produces a timestamped rotated filename.
func (s *RotatingFileSink) generateBackupPath() string {
	ts := time.Now().UTC().Format("20060102-150405.000")
	dir := filepath.Dir(s.cfg.FilePath)
	base := filepath.Base(s.cfg.FilePath)
	ext := filepath.Ext(base)
	stem := strings.TrimSuffix(base, ext)

	return filepath.Join(dir, fmt.Sprintf("%s-%s%s", stem, ts, ext))
}

// archiveFile moves and optionally compresses the backup into the archive folder.
func (s *RotatingFileSink) archiveFile(srcPath string) error {
	_ = fileutil.EnsureDir(s.cfg.ArchiveDir, filepermtype.Standard)
	destName := filepath.Base(srcPath)
	destPath := filepath.Join(s.cfg.ArchiveDir, destName)

	if s.cfg.IsCompress {
		return compressLogFile(srcPath, destPath+".gz")
	}

	return os.Rename(srcPath, destPath)
}

// compressLogFile compresses src into a gzip dest and deletes src.
func compressLogFile(src, dest string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}

	defer in.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer out.Close()

	gz := gzip.NewWriter(out)
	if _, copyErr := io.Copy(gz, in); copyErr != nil {
		return copyErr
	}

	_ = gz.Close()

	return os.Remove(src)
}

// pruneBackups enforces retention limits by removing oldest files exceeding MaxBackups.
func (s *RotatingFileSink) pruneBackups() {
	targetDir := filepath.Dir(s.cfg.FilePath)
	if s.cfg.IsArchiveEnabled && s.cfg.ArchiveDir != "" {
		targetDir = s.cfg.ArchiveDir
	}

	entries, err := os.ReadDir(targetDir)
	if err != nil || len(entries) <= s.cfg.MaxBackups {
		return
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	excess := len(entries) - s.cfg.MaxBackups
	for i := 0; i < excess; i++ {
		_ = os.Remove(filepath.Join(targetDir, entries[i].Name()))
	}
}

// Sync flushes pending data to storage.
func (s *RotatingFileSink) Sync() error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.file != nil {
		return s.file.Sync()
	}

	return nil
}

// Close safely shuts down the active log file.
func (s *RotatingFileSink) Close() error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.file != nil {
		err := s.file.Close()
		s.file = nil

		return err
	}

	return nil
}
