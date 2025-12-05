# cleanupBot Doc

Small Go utility to find and delete files older than X days.

Usage:
  go run ./cmd/cleanup --dir ./logs --days 30 --dry-run=true

To actually delete:
  please be careful doing this:
  ALLOW_DELETE=true go run ./cmd/cleanup --dir ./logs --days 30 --dry-run=false

You can read about this on my medium page: https://timiajiboye.medium.com/automating-repetitive-cleanup-tasks-with-go-and-github-actions-ci-cd-for-beginners-b5d8de41e978
