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

	want := "Sum (part 1): 157\nSum (part 2): 70\n"
	got := buf.String()
	if got != want {
		t.Error("Incorrect output:", got)
	}
}
