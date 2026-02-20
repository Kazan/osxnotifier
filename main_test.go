package main

import (
	"os"
	"strings"
	"testing"
)

func TestOnReadySetup(t *testing.T) {
	t.Helper()

	defer func() {
		if recovered := recover(); recovered != nil {
			t.Fatalf("onReady panicked: %v", recovered)
		}
	}()

	onReady()
}

func TestOnExitCleanup(t *testing.T) {
	t.Helper()

	defer func() {
		if recovered := recover(); recovered != nil {
			t.Fatalf("onExit panicked: %v", recovered)
		}
	}()

	onExit()
}

func TestModuleDependenciesUpdated(t *testing.T) {
	t.Helper()

	goModContent, err := os.ReadFile("go.mod")
	if err != nil {
		t.Fatalf("failed to read go.mod: %v", err)
	}

	content := string(goModContent)

	// Check that fyne.io/systray is a direct dependency (not indirect)
	if !strings.Contains(content, "fyne.io/systray") {
		t.Fatal("expected fyne.io/systray to be present in go.mod dependencies")
	}

	// Verify it's a direct dependency by checking it appears before the indirect section
	lines := strings.Split(content, "\n")
	foundSystray := false
	inIndirectSection := false

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if strings.Contains(trimmedLine, "// indirect") {
			inIndirectSection = true
		}

		if strings.Contains(trimmedLine, "fyne.io/systray") {
			foundSystray = true
			if inIndirectSection || strings.Contains(trimmedLine, "// indirect") {
				t.Fatal("expected fyne.io/systray to be a direct dependency, not indirect")
			}
			break
		}
	}

	if !foundSystray {
		t.Fatal("fyne.io/systray not found in go.mod")
	}
}
