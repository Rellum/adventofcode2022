package main

import (
	"bytes"
	"os"
	"testing"
)

func TestRun_example(t *testing.T) {
	f, _ := os.Open("example-input.txt")
	var buf bytes.Buffer

	err := run([]string{}, f, &buf)
	if err != nil {
		t.Error("Received error:", err)
	}

	want := "Part 1 sum: 95437\nPart 2 min size: 24933642\n"
	got := buf.String()
	if got != want {
		t.Error("Incorrect output:", got)
	}
}
