package main

import (
	"testing"
)

func TestDecodeInt(t *testing.T) {
	s := "i52e"
	expected := 52
	result, err := decodeInt(s)
	if err != nil {
		t.Errorf("expected %d but got %d", expected, result)
	}

	t.Logf("result: %v", result)
}
