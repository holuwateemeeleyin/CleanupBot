package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	dir := flag.String("dir", ".", "Directory to scan")
	days := flag.Int("days", 30, "Remove files older than this many days")
	flag.Parse()

	cutoff := time.Now().Add(-time.Duration(*days) * 24 * time.Hour)

	fmt.Printf("Scanning folder: %s\n", *dir)
	fmt.Printf("Cutoff date: %s\n", cutoff.Format(time.RFC3339))

	err := filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// If file is older than the cutoff, delete it
		if info.ModTime().Before(cutoff) {
			fmt.Printf("Deleting: %s\n", path)
			return os.Remove(path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cleanup complete")
	}
}
