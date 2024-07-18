package main

import (
	"fmt"
	"testing"
)

func TestSpamMasker(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http://hehe.HAHA", "http://*********"},
		{"123 http://hehe.HAHA", "123 http://*********"},
		{"http://hehe.HAHA 123", "http://********* 123"},
		{"123 http://hehe.HAHA 123", "123 http://********* 123"},
		{"1 http://hehe123.HAHA 1", "1 http://************ 1"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", tt.input), func(t *testing.T) {
			if got := spamMasker(tt.input); got != tt.expected {
				t.Errorf("spamMasker(%s) = %s; want %s", tt.input, got, tt.expected)
			}
		})
	}
}
