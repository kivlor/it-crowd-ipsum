package main

import (
	"testing"
)

func TestGenerateLispum(t *testing.T) {
	var subject []string

	subject = GenerateLipsum(5)

	if len(subject) != 5 {
		t.Error("expected a len of 5, got ", len(subject))
	}

	subject = GenerateLipsum(1)

	if len(subject) != 1 {
		t.Error("expected a len of 1, got ", len(subject))
	}
}
