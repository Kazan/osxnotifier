package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestAboutDialogContent(t *testing.T) {
	t.Helper()

	if appName != "OSX Notifier" {
		t.Fatalf("expected appName to be %q, got %q", "OSX Notifier", appName)
	}

	if appVersion != "1.0.0" {
		t.Fatalf("expected appVersion to be %q, got %q", "1.0.0", appVersion)
	}

	if appDescription != "A macOS menu bar application" {
		t.Fatalf("expected appDescription to be %q, got %q", "A macOS menu bar application", appDescription)
	}

	message := fmt.Sprintf("%s\\n\\nVersion: %s\\n\\n%s", appName, appVersion, appDescription)
	expectedMessage := "OSX Notifier\\n\\nVersion: 1.0.0\\n\\nA macOS menu bar application"
	if message != expectedMessage {
		t.Fatalf("expected formatted message %q, got %q", expectedMessage, message)
	}
}

func TestShowAboutDialogNoParams(t *testing.T) {
	t.Helper()
	originalRunOsaScript := runOsaScript
	t.Cleanup(func() {
		runOsaScript = originalRunOsaScript
	})

	called := false
	var capturedScript string
	runOsaScript = func(script string) error {
		called = true
		capturedScript = script
		return nil
	}

	showAboutDialog()

	if !called {
		t.Fatal("expected showAboutDialog to invoke osascript runner")
	}

	if !strings.Contains(capturedScript, "display dialog") {
		t.Fatalf("expected script to contain display dialog command, got %q", capturedScript)
	}

	if !strings.Contains(capturedScript, appName) {
		t.Fatalf("expected script to contain appName %q, got %q", appName, capturedScript)
	}
}
