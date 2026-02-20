package main

import (
	"bytes"
	"testing"
)

// TestIconDataNotEmpty verifies that the icon data is loaded and not empty
func TestIconDataNotEmpty(t *testing.T) {
	t.Helper()

	iconData := getIconData()
	if len(iconData) == 0 {
		t.Fatal("expected icon data to be non-empty")
	}
}

// TestIconIsValidPNG verifies that the icon data is a valid PNG format
func TestIconIsValidPNG(t *testing.T) {
	t.Helper()

	iconData := getIconData()

	// PNG signature: 89 50 4E 47 0D 0A 1A 0A
	pngSignature := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}

	if len(iconData) < len(pngSignature) {
		t.Fatalf("icon data is too short to be a valid PNG (got %d bytes)", len(iconData))
	}

	if !bytes.Equal(iconData[:len(pngSignature)], pngSignature) {
		t.Fatal("expected icon data to start with valid PNG signature")
	}
}
