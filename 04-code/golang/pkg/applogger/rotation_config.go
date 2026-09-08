package applogger

import "path/filepath"

const (
	// DefaultMaxSizeBytes sets the default rotation threshold to 2 MB.
	DefaultMaxSizeBytes int64 = 2 * 1024 * 1024
	// DefaultMaxBackups defines the default number of historical logs retained.
	DefaultMaxBackups int = 20
)

// RotationConfig controls file size limits, backup retention, and archiving.
type RotationConfig struct {
	FilePath         string
	MaxSizeBytes     int64
	MaxBackups       int
	IsArchiveEnabled bool
	ArchiveDir       string
	IsCompress       bool
}

// DefaultRotationConfig constructs standard settings with 2MB limits and 20 backups.
func DefaultRotationConfig(filePath string) RotationConfig {
	archiveDir := filepath.Join(filepath.Dir(filePath), "archives")

	return RotationConfig{
		FilePath:         filePath,
		MaxSizeBytes:     DefaultMaxSizeBytes,
		MaxBackups:       DefaultMaxBackups,
		IsArchiveEnabled: true,
		ArchiveDir:       archiveDir,
		IsCompress:       false,
	}
}

// Normalize sets fallback defaults for unspecified thresholds.
func (c *RotationConfig) Normalize() {
	if c.MaxSizeBytes <= 0 {
		c.MaxSizeBytes = DefaultMaxSizeBytes
	}

	if c.MaxBackups <= 0 {
		c.MaxBackups = DefaultMaxBackups
	}

	if c.ArchiveDir == "" && c.IsArchiveEnabled {
		c.ArchiveDir = filepath.Join(filepath.Dir(c.FilePath), "archives")
	}
}
