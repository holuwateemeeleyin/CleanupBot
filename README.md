# cleanupBot Doc

Small Go utility to find and delete files older than X days.

Usage:
  go run ./cmd/cleanup --dir ./logs --days 30 --dry-run=true

To actually delete:
  please be careful doing this:
  ALLOW_DELETE=true go run ./cmd/cleanup --dir ./logs --days 30 --dry-run=false
