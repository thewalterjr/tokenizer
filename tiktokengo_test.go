package tiktokengo_test

import (
	"testing"

	tiktokengo "github.com/thewalterjr/tokenizer"
)

func TestCountTokens(t *testing.T) {
	testCases := []struct {
		prompt string
		count  int
	}{
		{prompt: "", count: 0},
		{prompt: "Hello world", count: 2},
		{prompt: "Walter é lindo", count: 5},
	}

	for _, tc := range testCases {
		got := tiktokengo.CountTokens(tc.prompt)
		if got != tc.count {
			t.Errorf("CountTokens(%q) = %v; want %v", tc.prompt, got, tc.count)
		}
	}
}
