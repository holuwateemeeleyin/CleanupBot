package cleaner

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFindOldFilesAndDelete(t *testing.T) {
	tmp := t.TempDir()

	// create two files: one old, one new
	oldFile := filepath.Join(tmp, "old.log")
	newFile := filepath.Join(tmp, "new.log")

	if err := os.WriteFile(oldFile, []byte("old"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(newFile, []byte("new"), 0644); err != nil {
		t.Fatal(err)
	}

	// set modification times: old = 10 days ago, new = now
	now := time.Now()
	if err := os.Chtimes(oldFile, now.Add(-10*24*time.Hour), now.Add(-10*24*time.Hour)); err != nil {
		t.Fatal(err)
	}

	cutoff := time.Now().Add(-5 * 24 * time.Hour) // files older than 5 days
	files, err := FindOldFiles(tmp, cutoff)
	if err != nil {
		t.Fatalf("FindOldFiles failed: %v", err)
	}

	if len(files) != 1 {
		t.Fatalf("expected 1 old file, got %d", len(files))
	}
	if files[0].Path != oldFile {
		t.Fatalf("expected %s, got %s", oldFile, files[0].Path)
	}

	// test delete
	deleted, err := DeleteFiles(files)
	if err != nil {
		t.Fatalf("DeleteFiles error: %v", err)
	}
	if deleted != 1 {
		t.Fatalf("expected 1 deletion, got %d", deleted)
	}

	// ensure file removed
	if _, err := os.Stat(oldFile); !os.IsNotExist(err) {
		t.Fatalf("expected old file to be removed, stat error: %v", err)
	}
}
