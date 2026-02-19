package main

import (
	"os"
	"strings"
	"testing"
)

func TestAppInitialization(t *testing.T) {
	t.Helper()
	t.Setenv("FYNE_DRIVER", "test")

	defer func() {
		if recovered := recover(); recovered != nil {
			t.Fatalf("app initialization panicked: %v", recovered)
		}
	}()

	application := initializeApp()
	if application == nil {
		t.Fatal("expected initialized app to be non-nil")
	}
}

func TestModuleDependencies(t *testing.T) {
	t.Helper()

	goModContent, err := os.ReadFile("go.mod")
	if err != nil {
		t.Fatalf("failed to read go.mod: %v", err)
	}

	if !strings.Contains(string(goModContent), "fyne.io/fyne/v2") {
		t.Fatal("expected fyne.io/fyne/v2 to be present in go.mod dependencies")
	}
}
