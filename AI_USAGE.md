# AI usage

This project was developed with GitHub Copilot in VS Code.

- Tools used: workspace file inspection, repository search, patch-based file editing, and PowerShell command execution.
- AI helped scaffold the Go standard-library service, reason about idempotent batch writes and mutex protection, write focused tests, and draft the assignment documentation.
- A SQLite-backed design was considered but not used. The reason was to keep the empty-workspace submission dependency-free and easy to run; the scale trade-off and migration path are documented in `NOTES.md`.
- AI helped explain why duplicate detection and counter mutation must share one critical section for concurrent requests.
- The generated implementation could not be compiled or debugged in this environment because the `go` executable is not installed. No test, race-detector, seed execution, or debugging output is claimed as completed; the exact commands remain in `README.md`.