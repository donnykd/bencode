package main

import (
	"bytes"
	"testing"
)

func TestDecodeInt(t *testing.T) {
	s := "i52e"
	expected := 52
	result, err := decodeInt(s)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("expected %d but got %d", expected, result)
	}

	t.Logf("result: %v", result)
}

func TestDecodeByteString(t *testing.T) {
	s := "12:hello world!"
	expected := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd', '!'}
	result, err := decodeByteString(s)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !bytes.Equal(result, expected) {
		t.Errorf("expected %s but got %s", expected, result)
	}

	t.Logf("result: %v", result)
}
