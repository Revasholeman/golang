package main

import (
	"fmt"
	"testing"
)

func TestSpamMasker(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"test: 1", "http://hehe.HAHA", "http://*********"},
		{"test: 2", "123 http://hehe.HAHA", "123 http://*********"},
		{"test: 3", "http://hehe.HAHA 123", "http://********* 123"},
		{"test: 4", "123 http://hehe.HAHA 123", "123 http://********* 123"},
		{"test: 5", "1 http://hehe123.HAHA 1", "1 http://************ 1"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", tt.input), func(t *testing.T) {
			if got := spamMasker(tt.input); got != tt.expected {
				t.Errorf("spamMasker(%s) = %s; want %s", tt.input, got, tt.expected)
			}
		})
	}
}
