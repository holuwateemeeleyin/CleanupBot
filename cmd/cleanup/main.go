package main

import (
	"cleanUpWithGo/internal/cleaner"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	dir := flag.String("dir", ".", "Directory to scan")
	days := flag.Int("days", 30, "Remove files older than this many days")
	dryRun := flag.Bool("dry-run", true, "If true, do not perform deletions (default true)")
	allowDeleteEnv := flag.String("allow-delete-env", "ALLOW_DELETE", "Env var that must be set to 'true' to allow deletions when dry-run=false")
	flag.Parse()

	// compute cutoff
	cutoff := time.Now().Add(-time.Duration(*days) * 24 * time.Hour)

	fmt.Printf("Scanning directory: %s\n", *dir)
	fmt.Printf("Removing files older than %d days (before %s)\n", *days, cutoff.Format(time.RFC3339))

	files, err := cleaner.FindOldFiles(*dir, cutoff)
	if err != nil {
		log.Fatalf("error scanning files: %v", err)
	}

	if len(files) == 0 {
		fmt.Println("No files to remove.")
		return
	}

	for _, f := range files {
		fmt.Printf("Candidate: %s (modified: %s)\n", f.Path, f.Info.ModTime().Format(time.RFC3339))
	}

	if *dryRun {
		fmt.Printf("Dry run enabled — no files will be deleted. %d candidates found.\n", len(files))
		return
	}

	// dryRun is false, but require explicit env var to actually delete
	if os.Getenv(*allowDeleteEnv) != "true" {
		log.Fatalf("Deletion blocked: env %s is not set to 'true'. Set it in CI secrets if you want deletions.", *allowDeleteEnv)
	}

	deleted, delErr := cleaner.DeleteFiles(files)
	if delErr != nil {
		log.Printf("delete completed with errors, deleted %d files; last error: %v\n", deleted, delErr)
	} else {
		fmt.Printf("Deleted %d files successfully.\n", deleted)
	}
}
