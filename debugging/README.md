# Debugging exercise

The original repository supplied to this workspace did not contain a `debugging/` directory. To keep the assignment runnable, this directory contains a small equivalent JSONL exercise and its fixed implementation. The four defects documented in `../BUGS.md` are the defects addressed in this reconstructed fixture.

Run from this directory:

```text
go run -race . events.jsonl
go run . events.jsonl
```

Both commands must print the contents of `expected_output.txt` exactly.