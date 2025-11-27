package cleaner

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

// FileInfo holds info about a candidate file
type FileInfo struct {
	Path string
	Info fs.FileInfo
}

// FindOldFiles walks root and returns files (not directories) older than cutoff.
func FindOldFiles(root string, cutoff time.Time) ([]FileInfo, error) {
	if root == "" {
		return nil, errors.New("root path is empty")
	}

	var files []FileInfo
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// If one path errors, continue but record error to caller
			return err
		}
		if d.IsDir() {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		if info.ModTime().Before(cutoff) {
			files = append(files, FileInfo{Path: path, Info: info})
		}
		return nil
	})
	if err != nil {
		return files, err
	}
	return files, nil
}

// DeleteFiles deletes the given files. Returns number deleted and any error encountered.
func DeleteFiles(list []FileInfo) (int, error) {
	deleted := 0
	var lastErr error
	for _, f := range list {
		err := os.Remove(f.Path)
		if err != nil {
			lastErr = err
			// continue attempting to delete others
			continue
		}
		deleted++
	}
	return deleted, lastErr
}
