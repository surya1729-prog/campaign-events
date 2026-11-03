package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type event struct {
	EventID string `json:"event_id"`
	Type    string `json:"type"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: go run . events.jsonl")
		os.Exit(2)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	var mu sync.Mutex
	seen := make(map[string]bool)
	counts := map[string]int{"sent": 0, "delivered": 0, "opened": 0, "clicked": 0}
	processed, accepted, duplicates, invalid := 0, 0, 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processed++
		var current event
		if json.Unmarshal(scanner.Bytes(), &current) != nil || current.EventID == "" {
			invalid++
			continue
		}
		if _, ok := counts[current.Type]; !ok {
			invalid++
			continue
		}
		mu.Lock()
		if seen[current.EventID] {
			duplicates++
			mu.Unlock()
			continue
		}
		seen[current.EventID] = true
		counts[current.Type]++
		accepted++
		mu.Unlock()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("processed=%d\naccepted=%d\nduplicates=%d\ninvalid=%d\nsent=%d\ndelivered=%d\nopened=%d\nclicked=%d\n", processed, accepted, duplicates, invalid, counts["sent"], counts["delivered"], counts["opened"], counts["clicked"])
}