# cleanup-go

Small Go utility to find and delete files older than X days.

Usage:
  go run ./cmd/cleanup --dir ./logs --days 30 --dry-run=true

To actually delete:
  ALLOW_DELETE=true go run ./cmd/cleanup --dir ./logs --days 30 --dry-run=false
