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

	want := "Max: 24000\nSum of top 3: 45000\n"
	got := buf.String()
	if got != want {
		t.Error("Incorrect output:", got)
	}
}
