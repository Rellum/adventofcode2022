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

	want := "Top crates (part 1): CMZ\nTop crates (part 2): MCD\n"
	got := buf.String()
	if got != want {
		t.Error("Incorrect output:", got)
	}
}
