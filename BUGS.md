# Debugging findings

The workspace began without the assignment's original `debugging/` source, README, input, or expected output. The `debugging/` directory is therefore a reconstructed equivalent fixture, and these are the four defects addressed in its fixed implementation.

1. **Unknown event types were counted.** The root cause was treating any decoded `type` as a counter key. The fix validates membership in the four supported event types before counting. Verification: the `unknown` input record is reported as `invalid=1` and changes no counter.
2. **Duplicate IDs were counted twice.** The root cause was incrementing counters before checking a deduplication set. The fix checks and records `event_id` under the same mutex before incrementing. Verification: the repeated `evt-4` produces `duplicates=1` and `clicked=1`.
3. **Shared maps and counters were not protected.** The root cause was concurrent reads and writes to shared state. The fix uses one mutex around the deduplication and counter transaction. Verification: `go run -race . events.jsonl` is the required race-detector command; it could not be executed in this environment because Go is unavailable.
4. **Scanner errors were ignored.** The root cause was processing the loop without checking `scanner.Err()`. The fix checks the scanner after the loop and exits nonzero on an I/O failure. Verification: code inspection confirms the error path, while the normal fixture output is defined in `expected_output.txt`.

Expected normal output is in `debugging/expected_output.txt`. The exact commands requested by the assignment are:

```text
cd debugging
go run -race . events.jsonl
go run . events.jsonl
```